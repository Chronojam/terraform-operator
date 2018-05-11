// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroupAttachment struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsAlbTargetGroupAttachmentSpec `json:"spec"`
}

type AwsAlbTargetGroupAttachmentSpec struct {
	TargetGroupArn   string `json:"target_group_arn"`
	TargetId         string `json:"target_id"`
	Port             int    `json:"port"`
	AvailabilityZone string `json:"availability_zone"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroupAttachmentList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsAlbTargetGroupAttachment `json:"items"`
}
