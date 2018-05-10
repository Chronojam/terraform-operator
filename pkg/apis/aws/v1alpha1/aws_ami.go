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
	Name                 string                         `json:"name"`
	RootDeviceName       string                         `json:"root_device_name"`
	RootSnapshotId       string                         `json:"root_snapshot_id"`
	EbsBlockDevice       AwsAmiSpecEbsBlockDevice       `json:"ebs_block_device"`
	ImageLocation        string                         `json:"image_location"`
	Description          string                         `json:"description"`
	RamdiskId            string                         `json:"ramdisk_id"`
	SriovNetSupport      string                         `json:"sriov_net_support"`
	Architecture         string                         `json:"architecture"`
	VirtualizationType   string                         `json:"virtualization_type"`
	EphemeralBlockDevice AwsAmiSpecEphemeralBlockDevice `json:"ephemeral_block_device"`
	ManageEbsSnapshots   bool                           `json:"manage_ebs_snapshots"`
	KernelId             string                         `json:"kernel_id"`
	Tags                 map[string]string              `json:"tags"`
}

type AwsAmiSpecEbsBlockDevice struct {
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
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
