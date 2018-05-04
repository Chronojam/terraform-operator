// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchEventRuleSpec `json"spec"`
}

type AwsCloudwatchEventRuleSpec struct {
	IsEnabled          bool   `json:"is_enabled"`
	Arn                string `json:"arn"`
	Name               string `json:"name"`
	NamePrefix         string `json:"name_prefix"`
	ScheduleExpression string `json:"schedule_expression"`
	EventPattern       string `json:"event_pattern"`
	Description        string `json:"description"`
	RoleArn            string `json:"role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchEventRule `json"items"`
}
