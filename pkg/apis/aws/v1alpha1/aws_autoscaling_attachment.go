// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingAttachment struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsAutoscalingAttachmentSpec `json:"spec"`
}

type AwsAutoscalingAttachmentSpec struct {
	AutoscalingGroupName string `json:"autoscaling_group_name"`
	Elb                  string `json:"elb"`
	AlbTargetGroupArn    string `json:"alb_target_group_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingAttachmentList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsAutoscalingAttachment `json:"items"`
}
