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

type AuthenticationObservation struct {

	// The intended audience to receive authentication tokens for the service. The default value is https://dicom.azurehealthcareapis.azure.com
	Audience []*string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
	// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
	Authority *string `json:"authority,omitempty" tf:"authority,omitempty"`
}

type AuthenticationParameters struct {
}

type HealthcareDICOMServiceObservation struct {

	// The authentication block as defined below.
	Authentication []AuthenticationObservation `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// The ID of the Healthcare DICOM Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	PrivateEndpoint []PrivateEndpointObservation `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	// The url of the Healthcare DICOM Services.
	ServiceURL *string `json:"serviceUrl,omitempty" tf:"service_url,omitempty"`
}

type HealthcareDICOMServiceParameters struct {

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Healthcare DICOM Service should be created. Changing this forces a new Healthcare DICOM Service to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Whether to enabled public networks when data plane traffic coming from public networks while private endpoint is enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the id of the Healthcare Workspace where the Healthcare DICOM Service should exist. Changing this forces a new Healthcare DICOM Service to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/healthcareapis/v1beta1.HealthcareWorkspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a HealthcareWorkspace in healthcareapis to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a HealthcareWorkspace in healthcareapis to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

type IdentityObservation struct {

	// The ID of the Healthcare DICOM Service.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The ID of the Healthcare DICOM Service.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// A list of User Assigned Identity IDs which should be assigned to this Healthcare DICOM service.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The type of identity used for the Healthcare DICOM service. Possible values are UserAssigned, SystemAssigned and SystemAssigned, UserAssigned. If UserAssigned is set, an identity_ids must be set as well.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type PrivateEndpointObservation struct {

	// The ID of the Healthcare DICOM Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Healthcare DICOM Service. Changing this forces a new Healthcare DICOM Service to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PrivateEndpointParameters struct {
}

// HealthcareDICOMServiceSpec defines the desired state of HealthcareDICOMService
type HealthcareDICOMServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthcareDICOMServiceParameters `json:"forProvider"`
}

// HealthcareDICOMServiceStatus defines the observed state of HealthcareDICOMService.
type HealthcareDICOMServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthcareDICOMServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareDICOMService is the Schema for the HealthcareDICOMServices API. Manages a Healthcare DICOM (Digital Imaging and Communications in Medicine) Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HealthcareDICOMService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HealthcareDICOMServiceSpec   `json:"spec"`
	Status            HealthcareDICOMServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareDICOMServiceList contains a list of HealthcareDICOMServices
type HealthcareDICOMServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthcareDICOMService `json:"items"`
}

// Repository type metadata.
var (
	HealthcareDICOMService_Kind             = "HealthcareDICOMService"
	HealthcareDICOMService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthcareDICOMService_Kind}.String()
	HealthcareDICOMService_KindAPIVersion   = HealthcareDICOMService_Kind + "." + CRDGroupVersion.String()
	HealthcareDICOMService_GroupVersionKind = CRDGroupVersion.WithKind(HealthcareDICOMService_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthcareDICOMService{}, &HealthcareDICOMServiceList{})
}
