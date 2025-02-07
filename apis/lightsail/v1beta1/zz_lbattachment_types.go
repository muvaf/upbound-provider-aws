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

type LBAttachmentObservation struct {

	// A combination of attributes to create a unique id: lb_name,instance_name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LBAttachmentParameters struct {

	// The name of the instance to attach to the load balancer.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta1.Instance
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Reference to a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameRef *v1.Reference `json:"instanceNameRef,omitempty" tf:"-"`

	// Selector for a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameSelector *v1.Selector `json:"instanceNameSelector,omitempty" tf:"-"`

	// The name of the Lightsail load balancer.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta1.LB
	// +kubebuilder:validation:Optional
	LBName *string `json:"lbName,omitempty" tf:"lb_name,omitempty"`

	// Reference to a LB in lightsail to populate lbName.
	// +kubebuilder:validation:Optional
	LBNameRef *v1.Reference `json:"lbNameRef,omitempty" tf:"-"`

	// Selector for a LB in lightsail to populate lbName.
	// +kubebuilder:validation:Optional
	LBNameSelector *v1.Selector `json:"lbNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LBAttachmentSpec defines the desired state of LBAttachment
type LBAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBAttachmentParameters `json:"forProvider"`
}

// LBAttachmentStatus defines the observed state of LBAttachment.
type LBAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBAttachment is the Schema for the LBAttachments API. Attaches a Lightsail Instance to a Lightsail Load Balancer
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBAttachmentSpec   `json:"spec"`
	Status            LBAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBAttachmentList contains a list of LBAttachments
type LBAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBAttachment `json:"items"`
}

// Repository type metadata.
var (
	LBAttachment_Kind             = "LBAttachment"
	LBAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBAttachment_Kind}.String()
	LBAttachment_KindAPIVersion   = LBAttachment_Kind + "." + CRDGroupVersion.String()
	LBAttachment_GroupVersionKind = CRDGroupVersion.WithKind(LBAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&LBAttachment{}, &LBAttachmentList{})
}
