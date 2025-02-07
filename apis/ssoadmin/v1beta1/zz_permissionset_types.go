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

type PermissionSetObservation struct {

	// The Amazon Resource Name (ARN) of the Permission Set.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The date the Permission Set was created in RFC3339 format.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// The Amazon Resource Names (ARNs) of the Permission Set and SSO Instance, separated by a comma (,).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type PermissionSetParameters struct {

	// The description of the Permission Set.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	// +kubebuilder:validation:Required
	InstanceArn *string `json:"instanceArn" tf:"instance_arn,omitempty"`

	// The name of the Permission Set.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The relay state URL used to redirect users within the application during the federation authentication process.
	// +kubebuilder:validation:Optional
	RelayState *string `json:"relayState,omitempty" tf:"relay_state,omitempty"`

	// The length of time that the application user sessions are valid in the ISO-8601 standard. Default: PT1H.
	// +kubebuilder:validation:Optional
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PermissionSetSpec defines the desired state of PermissionSet
type PermissionSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionSetParameters `json:"forProvider"`
}

// PermissionSetStatus defines the observed state of PermissionSet.
type PermissionSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionSet is the Schema for the PermissionSets API. Manages a Single Sign-On (SSO) Permission Set
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PermissionSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionSetSpec   `json:"spec"`
	Status            PermissionSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionSetList contains a list of PermissionSets
type PermissionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PermissionSet `json:"items"`
}

// Repository type metadata.
var (
	PermissionSet_Kind             = "PermissionSet"
	PermissionSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PermissionSet_Kind}.String()
	PermissionSet_KindAPIVersion   = PermissionSet_Kind + "." + CRDGroupVersion.String()
	PermissionSet_GroupVersionKind = CRDGroupVersion.WithKind(PermissionSet_Kind)
)

func init() {
	SchemeBuilder.Register(&PermissionSet{}, &PermissionSetList{})
}
