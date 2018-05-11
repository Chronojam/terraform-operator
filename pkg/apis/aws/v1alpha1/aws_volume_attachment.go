// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVolumeAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVolumeAttachmentSpec `json"spec"`
}

type AwsVolumeAttachmentSpec struct {
	SkipDestroy bool   `json:"skip_destroy"`
	DeviceName  string `json:"device_name"`
	InstanceId  string `json:"instance_id"`
	VolumeId    string `json:"volume_id"`
	ForceDetach bool   `json:"force_detach"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVolumeAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVolumeAttachment `json"items"`
}
