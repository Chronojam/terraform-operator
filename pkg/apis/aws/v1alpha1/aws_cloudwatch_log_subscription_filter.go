// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogSubscriptionFilter struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchLogSubscriptionFilterSpec `json"spec"`
}

type AwsCloudwatchLogSubscriptionFilterSpec struct {
	FilterPattern  string `json:"filter_pattern"`
	LogGroupName   string `json:"log_group_name"`
	RoleArn        string `json:"role_arn"`
	Distribution   string `json:"distribution"`
	Name           string `json:"name"`
	DestinationArn string `json:"destination_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogSubscriptionFilterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchLogSubscriptionFilter `json"items"`
}
