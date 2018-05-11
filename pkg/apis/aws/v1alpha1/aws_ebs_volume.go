// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEbsVolume struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEbsVolumeSpec `json"spec"`
}

type AwsEbsVolumeSpec struct {
	Encrypted        bool              `json:"encrypted"`
	Size             int               `json:"size"`
	Tags             map[string]string `json:"tags"`
	Arn              string            `json:"arn"`
	AvailabilityZone string            `json:"availability_zone"`
	SnapshotId       string            `json:"snapshot_id"`
	Type             string            `json:"type"`
	Iops             int               `json:"iops"`
	KmsKeyId         string            `json:"kms_key_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEbsVolumeList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEbsVolume `json"items"`
}
