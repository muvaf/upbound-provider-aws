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

type CognitoMemberDefinitionObservation struct {
}

type CognitoMemberDefinitionParameters struct {

	// An identifier for an application client. You must create the app client ID using Amazon Cognito.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolClient
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// An identifier for a user group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserGroup *string `json:"userGroup,omitempty" tf:"user_group,omitempty"`

	// Reference to a UserGroup in cognitoidp to populate userGroup.
	// +kubebuilder:validation:Optional
	UserGroupRef *v1.Reference `json:"userGroupRef,omitempty" tf:"-"`

	// Selector for a UserGroup in cognitoidp to populate userGroup.
	// +kubebuilder:validation:Optional
	UserGroupSelector *v1.Selector `json:"userGroupSelector,omitempty" tf:"-"`

	// An identifier for a user pool. The user pool must be in the same region as the service that you are calling.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolDomain
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("user_pool_id",false)
	// +kubebuilder:validation:Optional
	UserPool *string `json:"userPool,omitempty" tf:"user_pool,omitempty"`

	// Reference to a UserPoolDomain in cognitoidp to populate userPool.
	// +kubebuilder:validation:Optional
	UserPoolRef *v1.Reference `json:"userPoolRef,omitempty" tf:"-"`

	// Selector for a UserPoolDomain in cognitoidp to populate userPool.
	// +kubebuilder:validation:Optional
	UserPoolSelector *v1.Selector `json:"userPoolSelector,omitempty" tf:"-"`
}

type MemberDefinitionObservation struct {
}

type MemberDefinitionParameters struct {

	// The Amazon Cognito user group that is part of the work team. See Cognito Member Definition details below.
	// +kubebuilder:validation:Optional
	CognitoMemberDefinition []CognitoMemberDefinitionParameters `json:"cognitoMemberDefinition,omitempty" tf:"cognito_member_definition,omitempty"`

	// A list user groups that exist in your OIDC Identity Provider (IdP). One to ten groups can be used to create a single private work team. See Cognito Member Definition details below.
	// +kubebuilder:validation:Optional
	OidcMemberDefinition []OidcMemberDefinitionParameters `json:"oidcMemberDefinition,omitempty" tf:"oidc_member_definition,omitempty"`
}

type NotificationConfigurationObservation struct {
}

type NotificationConfigurationParameters struct {

	// The ARN for the SNS topic to which notifications should be published.
	// +kubebuilder:validation:Optional
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
}

type OidcMemberDefinitionObservation struct {
}

type OidcMemberDefinitionParameters struct {

	// A list of comma separated strings that identifies user groups in your OIDC IdP. Each user group is made up of a group of private workers.
	// +kubebuilder:validation:Required
	Groups []*string `json:"groups" tf:"groups,omitempty"`
}

type WorkteamObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this Workteam.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the Workteam.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The subdomain for your OIDC Identity Provider.
	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WorkteamParameters struct {

	// A description of the work team.
	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use cognito_member_definition. For workforces created using your own OIDC identity provider (IdP) use oidc_member_definition. Do not provide input for both of these parameters in a single request. see Member Definition details below.
	// +kubebuilder:validation:Required
	MemberDefinition []MemberDefinitionParameters `json:"memberDefinition" tf:"member_definition,omitempty"`

	// Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
	// +kubebuilder:validation:Optional
	NotificationConfiguration []NotificationConfigurationParameters `json:"notificationConfiguration,omitempty" tf:"notification_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The name of the Workteam (must be unique).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta1.Workforce
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkforceName *string `json:"workforceName,omitempty" tf:"workforce_name,omitempty"`

	// Reference to a Workforce in sagemaker to populate workforceName.
	// +kubebuilder:validation:Optional
	WorkforceNameRef *v1.Reference `json:"workforceNameRef,omitempty" tf:"-"`

	// Selector for a Workforce in sagemaker to populate workforceName.
	// +kubebuilder:validation:Optional
	WorkforceNameSelector *v1.Selector `json:"workforceNameSelector,omitempty" tf:"-"`
}

// WorkteamSpec defines the desired state of Workteam
type WorkteamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkteamParameters `json:"forProvider"`
}

// WorkteamStatus defines the observed state of Workteam.
type WorkteamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkteamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workteam is the Schema for the Workteams API. Provides a SageMaker Workteam resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Workteam struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkteamSpec   `json:"spec"`
	Status            WorkteamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkteamList contains a list of Workteams
type WorkteamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workteam `json:"items"`
}

// Repository type metadata.
var (
	Workteam_Kind             = "Workteam"
	Workteam_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workteam_Kind}.String()
	Workteam_KindAPIVersion   = Workteam_Kind + "." + CRDGroupVersion.String()
	Workteam_GroupVersionKind = CRDGroupVersion.WithKind(Workteam_Kind)
)

func init() {
	SchemeBuilder.Register(&Workteam{}, &WorkteamList{})
}
