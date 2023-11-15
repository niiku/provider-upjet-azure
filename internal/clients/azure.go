package clients

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/terraform-provider-azurerm/xpprovider"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/provider-azure/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal Azure credentials as JSON"
	errSubscriptionIDNotSet = "subscription ID must be set in ProviderConfig when credential source is InjectedIdentity, OIDCTokenFile or Upbound"
	errTenantIDNotSet       = "tenant ID must be set in ProviderConfig when credential source is InjectedIdentity, OIDCTokenFile or Upbound"
	errClientIDNotSet       = "client ID must be set in ProviderConfig when credential source is OIDCTokenFile or Upbound"
	// Azure service principal credentials file JSON keys
	keyAzureSubscriptionID = "subscriptionId"
	keyAzureClientID       = "clientId"
	keyAzureClientSecret   = "clientSecret"
	keyAzureClientCert     = "clientCertificate"
	keyAzureClientCertPass = "clientCertificatePassword"
	keyAzureTenantID       = "tenantId"
	// Terraform Provider configuration block keys
	keyTerraformFeatures        = "features"
	keySkipProviderRegistration = "skip_provider_registration"
	keyUseMSI                   = "use_msi"
	keyClientID                 = "client_id"
	keySubscriptionID           = "subscription_id"
	keyTenantID                 = "tenant_id"
	keyMSIEndpoint              = "msi_endpoint"
	keyClientSecret             = "client_secret"
	keyClientCert               = "client_certificate"
	keyClientCertPassword       = "client_certificate_password"
	keyEnvironment              = "environment"
	keyOidcTokenFilePath        = "oidc_token_file_path"
	keyUseOIDC                  = "use_oidc"
	// Default OidcTokenFilePath
	defaultOidcTokenFilePath = "/var/run/secrets/azure/tokens/azure-identity-token"
)

var (
	credentialsSourceUserAssignedManagedIdentity   xpv1.CredentialsSource = "UserAssignedManagedIdentity"
	credentialsSourceSystemAssignedManagedIdentity xpv1.CredentialsSource = "SystemAssignedManagedIdentity"
	credentialsSourceOIDCTokenFile                 xpv1.CredentialsSource = "OIDCTokenFile"
	credentialsSourceUpbound                       xpv1.CredentialsSource = "Upbound"

	upboundProviderIdentityTokenFile = "/var/run/secrets/upbound.io/provider/token"
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(version, providerSource, providerVersion string, scheduler terraform.ProviderScheduler) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg xpresource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
			Scheduler: scheduler,
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := xpresource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		ps.Configuration = map[string]interface{}{
			keyTerraformFeatures: struct{}{},
			// Terraform AzureRM provider tries to register all resource providers
			// in Azure just in case if the provider of the resource you're
			// trying to create is not registered and the returned error is
			// ambiguous. However, this requires service principal to have provider
			// registration permissions which are irrelevant in most contexts.
			// For details, see https://github.com/upbound/provider-azure/issues/104
			keySkipProviderRegistration: true,
		}

		var err error
		switch pc.Spec.Credentials.Source { //nolint:exhaustive
		case credentialsSourceSystemAssignedManagedIdentity, credentialsSourceUserAssignedManagedIdentity:
			err = msiAuth(pc, &ps)
		case credentialsSourceOIDCTokenFile:
			err = oidcAuth(pc, &ps)
		case credentialsSourceUpbound:
			err = upboundAuth(pc, &ps)
		default:
			err = spAuth(ctx, pc, &ps, client)
		}
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "failed to prepare terraform.Setup")
		}

		return ps, errors.Wrap(configureNoForkAzureClient(ctx, pc, &ps), "failed to configure the no-fork Azure client")
	}
}

func configureNoForkAzureClient(ctx context.Context, pc *v1beta1.ProviderConfig, ps *terraform.Setup) error {
	cb := xpprovider.AzureClientBuilder{}
	switch pc.Spec.Credentials.Source { //nolint:exhaustive
	case xpv1.CredentialsSourceSecret:
		cb.SubscriptionID = ps.Configuration[keySubscriptionID].(string)
		cb.AuthConfig = &auth.Credentials{
			ClientID:                              ps.Configuration[keyClientID].(string),
			TenantID:                              ps.Configuration[keyTenantID].(string),
			ClientSecret:                          ps.Configuration[keyClientSecret].(string),
			EnableAuthenticatingUsingClientSecret: true,
		}
	}
	// TODO: we need to check how to prepare environment's
	// authorization context, this may differ especially if
	// we are preparing an environment inside an EKS cluster, etc.
	cb.AuthConfig.Environment = *environments.AzurePublic()
	c, err := cb.GetClient(context.WithoutCancel(ctx))
	if err != nil {
		return errors.Wrap(err, "failed to build the Terraform Azure provider client")
	}
	ps.Meta = c
	return nil
}

func spAuth(ctx context.Context, pc *v1beta1.ProviderConfig, ps *terraform.Setup, client client.Client) error {
	data, err := xpresource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return errors.Wrap(err, errExtractCredentials)
	}
	data = []byte(strings.TrimSpace(string(data)))
	azureCreds := map[string]string{}
	if err := json.Unmarshal(data, &azureCreds); err != nil {
		return errors.Wrap(err, errUnmarshalCredentials)
	}
	// set credentials configuration
	ps.Configuration[keySubscriptionID] = azureCreds[keyAzureSubscriptionID]
	ps.Configuration[keyTenantID] = azureCreds[keyAzureTenantID]
	ps.Configuration[keyClientID] = azureCreds[keyAzureClientID]
	ps.Configuration[keyClientSecret] = azureCreds[keyAzureClientSecret]
	if clientCert, ok := azureCreds[keyAzureClientCert]; ok {
		ps.Configuration[keyClientCert] = clientCert
		if clientCertPass, passwordOk := azureCreds[keyAzureClientCertPass]; passwordOk {
			ps.Configuration[keyClientCertPassword] = clientCertPass
		}
	}
	if pc.Spec.SubscriptionID != nil {
		ps.Configuration[keySubscriptionID] = *pc.Spec.SubscriptionID
	}
	if pc.Spec.TenantID != nil {
		ps.Configuration[keyTenantID] = *pc.Spec.TenantID
	}
	if pc.Spec.ClientID != nil {
		ps.Configuration[keyClientID] = *pc.Spec.ClientID
	}
	if pc.Spec.Environment != nil {
		ps.Configuration[keyEnvironment] = *pc.Spec.Environment
	}
	return nil
}

func msiAuth(pc *v1beta1.ProviderConfig, ps *terraform.Setup) error {
	if pc.Spec.SubscriptionID == nil || len(*pc.Spec.SubscriptionID) == 0 {
		return errors.New(errSubscriptionIDNotSet)
	}
	if pc.Spec.TenantID == nil || len(*pc.Spec.TenantID) == 0 {
		return errors.New(errTenantIDNotSet)
	}
	ps.Configuration[keySubscriptionID] = *pc.Spec.SubscriptionID
	ps.Configuration[keyTenantID] = *pc.Spec.TenantID
	ps.Configuration[keyUseMSI] = "true"
	if pc.Spec.MSIEndpoint != nil {
		ps.Configuration[keyMSIEndpoint] = *pc.Spec.MSIEndpoint
	}
	if pc.Spec.ClientID != nil {
		ps.Configuration[keyClientID] = *pc.Spec.ClientID
	}
	if pc.Spec.Environment != nil {
		ps.Configuration[keyEnvironment] = *pc.Spec.Environment
	}
	return nil
}

func oidcAuth(pc *v1beta1.ProviderConfig, ps *terraform.Setup) error {
	if pc.Spec.SubscriptionID == nil || len(*pc.Spec.SubscriptionID) == 0 {
		return errors.New(errSubscriptionIDNotSet)
	}
	if pc.Spec.TenantID == nil || len(*pc.Spec.TenantID) == 0 {
		return errors.New(errTenantIDNotSet)
	}
	if pc.Spec.ClientID == nil || len(*pc.Spec.ClientID) == 0 {
		return errors.New(errClientIDNotSet)
	}
	// OIDC Token File Path defaults to a projected-volume path mounted in the pod running in the AKS cluster, when workload identity is enabled on the pod.
	ps.Configuration[keyOidcTokenFilePath] = defaultOidcTokenFilePath
	if pc.Spec.OidcTokenFilePath != nil {
		ps.Configuration[keyOidcTokenFilePath] = *pc.Spec.OidcTokenFilePath
	}
	ps.Configuration[keySubscriptionID] = *pc.Spec.SubscriptionID
	ps.Configuration[keyTenantID] = *pc.Spec.TenantID
	ps.Configuration[keyClientID] = *pc.Spec.ClientID
	ps.Configuration[keyUseOIDC] = "true"
	return nil

}

func upboundAuth(pc *v1beta1.ProviderConfig, ps *terraform.Setup) error {
	if pc.Spec.SubscriptionID == nil || len(*pc.Spec.SubscriptionID) == 0 {
		return errors.New(errSubscriptionIDNotSet)
	}
	if pc.Spec.TenantID == nil || len(*pc.Spec.TenantID) == 0 {
		return errors.New(errTenantIDNotSet)
	}
	if pc.Spec.ClientID == nil || len(*pc.Spec.ClientID) == 0 {
		return errors.New(errClientIDNotSet)
	}
	ps.Configuration[keyOidcTokenFilePath] = upboundProviderIdentityTokenFile
	ps.Configuration[keySubscriptionID] = *pc.Spec.SubscriptionID
	ps.Configuration[keyTenantID] = *pc.Spec.TenantID
	ps.Configuration[keyClientID] = *pc.Spec.ClientID
	ps.Configuration[keyUseOIDC] = "true"
	return nil

}
