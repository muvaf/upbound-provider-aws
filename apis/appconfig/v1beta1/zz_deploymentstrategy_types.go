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

type DeploymentStrategyObservation struct {

	// ARN of the AppConfig Deployment Strategy.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// AppConfig deployment strategy ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DeploymentStrategyParameters struct {

	// Total amount of time for a deployment to last. Minimum value of 0, maximum value of 1440.
	// +kubebuilder:validation:Required
	DeploymentDurationInMinutes *float64 `json:"deploymentDurationInMinutes" tf:"deployment_duration_in_minutes,omitempty"`

	// Description of the deployment strategy. Can be at most 1024 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amount of time AWS AppConfig monitors for alarms before considering the deployment to be complete and no longer eligible for automatic roll back. Minimum value of 0, maximum value of 1440.
	// +kubebuilder:validation:Optional
	FinalBakeTimeInMinutes *float64 `json:"finalBakeTimeInMinutes,omitempty" tf:"final_bake_time_in_minutes,omitempty"`

	// Percentage of targets to receive a deployed configuration during each interval. Minimum value of 1.0, maximum value of 100.0.
	// +kubebuilder:validation:Required
	GrowthFactor *float64 `json:"growthFactor" tf:"growth_factor,omitempty"`

	// Algorithm used to define how percentage grows over time. Valid value: LINEAR and EXPONENTIAL. Defaults to LINEAR.
	// +kubebuilder:validation:Optional
	GrowthType *string `json:"growthType,omitempty" tf:"growth_type,omitempty"`

	// Name for the deployment strategy. Must be between 1 and 64 characters in length.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Where to save the deployment strategy. Valid values: NONE and SSM_DOCUMENT.
	// +kubebuilder:validation:Required
	ReplicateTo *string `json:"replicateTo" tf:"replicate_to,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DeploymentStrategySpec defines the desired state of DeploymentStrategy
type DeploymentStrategySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentStrategyParameters `json:"forProvider"`
}

// DeploymentStrategyStatus defines the observed state of DeploymentStrategy.
type DeploymentStrategyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentStrategyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentStrategy is the Schema for the DeploymentStrategys API. Provides an AppConfig Deployment Strategy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DeploymentStrategy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentStrategySpec   `json:"spec"`
	Status            DeploymentStrategyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentStrategyList contains a list of DeploymentStrategys
type DeploymentStrategyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentStrategy `json:"items"`
}

// Repository type metadata.
var (
	DeploymentStrategy_Kind             = "DeploymentStrategy"
	DeploymentStrategy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeploymentStrategy_Kind}.String()
	DeploymentStrategy_KindAPIVersion   = DeploymentStrategy_Kind + "." + CRDGroupVersion.String()
	DeploymentStrategy_GroupVersionKind = CRDGroupVersion.WithKind(DeploymentStrategy_Kind)
)

func init() {
	SchemeBuilder.Register(&DeploymentStrategy{}, &DeploymentStrategyList{})
}
