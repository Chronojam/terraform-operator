// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLaunchConfiguration struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLaunchConfigurationSpec `json"spec"`
}

type AwsLaunchConfigurationSpec struct {
	ImageId                      string                                         `json:"image_id"`
	VpcClassicLinkSecurityGroups string                                         `json:"vpc_classic_link_security_groups"`
	EbsBlockDevice               AwsLaunchConfigurationSpecEbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice         AwsLaunchConfigurationSpecEphemeralBlockDevice `json:"ephemeral_block_device"`
	Name                         string                                         `json:"name"`
	KeyName                      string                                         `json:"key_name"`
	PlacementTenancy             string                                         `json:"placement_tenancy"`
	RootBlockDevice              []AwsLaunchConfigurationSpecRootBlockDevice    `json:"root_block_device"`
	NamePrefix                   string                                         `json:"name_prefix"`
	InstanceType                 string                                         `json:"instance_type"`
	IamInstanceProfile           string                                         `json:"iam_instance_profile"`
	SecurityGroups               string                                         `json:"security_groups"`
	AssociatePublicIpAddress     bool                                           `json:"associate_public_ip_address"`
	UserData                     string                                         `json:"user_data"`
	UserDataBase64               string                                         `json:"user_data_base64"`
	VpcClassicLinkId             string                                         `json:"vpc_classic_link_id"`
	SpotPrice                    string                                         `json:"spot_price"`
	EbsOptimized                 bool                                           `json:"ebs_optimized"`
	EnableMonitoring             bool                                           `json:"enable_monitoring"`
}

type AwsLaunchConfigurationSpecEbsBlockDevice struct {
	VolumeType          string `json:"volume_type"`
	Encrypted           bool   `json:"encrypted"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	NoDevice            bool   `json:"no_device"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
}

type AwsLaunchConfigurationSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsLaunchConfigurationSpecRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLaunchConfigurationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLaunchConfiguration `json"items"`
}
