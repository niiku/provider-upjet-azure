// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SubscriptionInitParameters struct {

	// The ID of the API which should be assigned to this Subscription. Changing this forces a new resource to be created.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Determines whether tracing can be enabled. Defaults to true.
	AllowTracing *bool `json:"allowTracing,omitempty" tf:"allow_tracing,omitempty"`

	// The primary subscription key to use for the subscription.
	PrimaryKeySecretRef *v1.SecretKeySelector `json:"primaryKeySecretRef,omitempty" tf:"-"`

	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Product
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Reference to a Product in apimanagement to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDRef *v1.Reference `json:"productIdRef,omitempty" tf:"-"`

	// Selector for a Product in apimanagement to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDSelector *v1.Selector `json:"productIdSelector,omitempty" tf:"-"`

	// The secondary subscription key to use for the subscription.
	SecondaryKeySecretRef *v1.SecretKeySelector `json:"secondaryKeySecretRef,omitempty" tf:"-"`

	// The state of this Subscription. Possible values are active, cancelled, expired, rejected, submitted and suspended. Defaults to submitted.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in apimanagement to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in apimanagement to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

type SubscriptionObservation struct {

	// The ID of the API which should be assigned to this Subscription. Changing this forces a new resource to be created.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Determines whether tracing can be enabled. Defaults to true.
	AllowTracing *bool `json:"allowTracing,omitempty" tf:"allow_tracing,omitempty"`

	// The ID of the API Management Subscription.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The state of this Subscription. Possible values are active, cancelled, expired, rejected, submitted and suspended. Defaults to submitted.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type SubscriptionParameters struct {

	// The ID of the API which should be assigned to this Subscription. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta2.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// Determines whether tracing can be enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	AllowTracing *bool `json:"allowTracing,omitempty" tf:"allow_tracing,omitempty"`

	// The primary subscription key to use for the subscription.
	// +kubebuilder:validation:Optional
	PrimaryKeySecretRef *v1.SecretKeySelector `json:"primaryKeySecretRef,omitempty" tf:"-"`

	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Product
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Reference to a Product in apimanagement to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDRef *v1.Reference `json:"productIdRef,omitempty" tf:"-"`

	// Selector for a Product in apimanagement to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDSelector *v1.Selector `json:"productIdSelector,omitempty" tf:"-"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The secondary subscription key to use for the subscription.
	// +kubebuilder:validation:Optional
	SecondaryKeySecretRef *v1.SecretKeySelector `json:"secondaryKeySecretRef,omitempty" tf:"-"`

	// The state of this Subscription. Possible values are active, cancelled, expired, rejected, submitted and suspended. Defaults to submitted.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in apimanagement to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in apimanagement to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SubscriptionInitParameters `json:"initProvider,omitempty"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Subscription is the Schema for the Subscriptions API. Manages a Subscription within a API Management Service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionSpec   `json:"spec"`
	Status            SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}
