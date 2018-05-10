// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmi struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAmiSpec `json"spec"`
}

type AwsAmiSpec struct {
	ImageLocation        string                         `json:"image_location"`
	Tags                 map[string]string              `json:"tags"`
	ManageEbsSnapshots   bool                           `json:"manage_ebs_snapshots"`
	KernelId             string                         `json:"kernel_id"`
	RootDeviceName       string                         `json:"root_device_name"`
	EbsBlockDevice       AwsAmiSpecEbsBlockDevice       `json:"ebs_block_device"`
	SriovNetSupport      string                         `json:"sriov_net_support"`
	RamdiskId            string                         `json:"ramdisk_id"`
	RootSnapshotId       string                         `json:"root_snapshot_id"`
	VirtualizationType   string                         `json:"virtualization_type"`
	EphemeralBlockDevice AwsAmiSpecEphemeralBlockDevice `json:"ephemeral_block_device"`
	Architecture         string                         `json:"architecture"`
	Description          string                         `json:"description"`
	Name                 string                         `json:"name"`
}

type AwsAmiSpecEbsBlockDevice struct {
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
}

type AwsAmiSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAmi `json"items"`
}
