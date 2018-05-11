// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingLifecycleHook struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsAutoscalingLifecycleHookSpec `json:"spec"`
}

type AwsAutoscalingLifecycleHookSpec struct {
	Name                  string `json:"name"`
	AutoscalingGroupName  string `json:"autoscaling_group_name"`
	DefaultResult         string `json:"default_result"`
	HeartbeatTimeout      int    `json:"heartbeat_timeout"`
	LifecycleTransition   string `json:"lifecycle_transition"`
	NotificationMetadata  string `json:"notification_metadata"`
	NotificationTargetArn string `json:"notification_target_arn"`
	RoleArn               string `json:"role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingLifecycleHookList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsAutoscalingLifecycleHook `json:"items"`
}
