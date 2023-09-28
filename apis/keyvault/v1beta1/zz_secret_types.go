/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecretInitParameters struct {

	// Specifies the content type for the Key Vault Secret.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
	NotBeforeDate *string `json:"notBeforeDate,omitempty" tf:"not_before_date,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SecretObservation struct {

	// Specifies the content type for the Key Vault Secret.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The Key Vault Secret ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Key Vault where the Secret should be created. Changing this forces a new resource to be created.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
	NotBeforeDate *string `json:"notBeforeDate,omitempty" tf:"not_before_date,omitempty"`

	// The (Versioned) ID for this Key Vault Secret. This property points to a specific version of a Key Vault Secret, as such using this won't auto-rotate values if used in other Azure Services.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// The Versionless ID of the Key Vault Secret. This property allows other Azure Services (that support it) to auto-rotate their value when the Key Vault Secret is updated.
	ResourceVersionlessID *string `json:"resourceVersionlessId,omitempty" tf:"resource_versionless_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The current version of the Key Vault Secret.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The Base ID of the Key Vault Secret.
	VersionlessID *string `json:"versionlessId,omitempty" tf:"versionless_id,omitempty"`
}

type SecretParameters struct {

	// Specifies the content type for the Key Vault Secret.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
	// +kubebuilder:validation:Optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The ID of the Key Vault where the Secret should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Vault
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Vault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Vault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
	// +kubebuilder:validation:Optional
	NotBeforeDate *string `json:"notBeforeDate,omitempty" tf:"not_before_date,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the value of the Key Vault Secret.
	// +kubebuilder:validation:Optional
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

// SecretSpec defines the desired state of Secret
type SecretSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SecretInitParameters `json:"initProvider,omitempty"`
}

// SecretStatus defines the observed state of Secret.
type SecretStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Secret is the Schema for the Secrets API. Manages a Key Vault Secret.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.valueSecretRef)",message="spec.forProvider.valueSecretRef is a required parameter"
	Spec   SecretSpec   `json:"spec"`
	Status SecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretList contains a list of Secrets
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}

// Repository type metadata.
var (
	Secret_Kind             = "Secret"
	Secret_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Secret_Kind}.String()
	Secret_KindAPIVersion   = Secret_Kind + "." + CRDGroupVersion.String()
	Secret_GroupVersionKind = CRDGroupVersion.WithKind(Secret_Kind)
)

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}
