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
	RootBlockDevice              []AwsLaunchConfigurationSpecRootBlockDevice    `json:"root_block_device"`
	NamePrefix                   string                                         `json:"name_prefix"`
	ImageId                      string                                         `json:"image_id"`
	UserDataBase64               string                                         `json:"user_data_base64"`
	EbsOptimized                 bool                                           `json:"ebs_optimized"`
	EphemeralBlockDevice         AwsLaunchConfigurationSpecEphemeralBlockDevice `json:"ephemeral_block_device"`
	Name                         string                                         `json:"name"`
	UserData                     string                                         `json:"user_data"`
	AssociatePublicIpAddress     bool                                           `json:"associate_public_ip_address"`
	EnableMonitoring             bool                                           `json:"enable_monitoring"`
	EbsBlockDevice               AwsLaunchConfigurationSpecEbsBlockDevice       `json:"ebs_block_device"`
	SecurityGroups               string                                         `json:"security_groups"`
	VpcClassicLinkId             string                                         `json:"vpc_classic_link_id"`
	SpotPrice                    string                                         `json:"spot_price"`
	PlacementTenancy             string                                         `json:"placement_tenancy"`
	InstanceType                 string                                         `json:"instance_type"`
	IamInstanceProfile           string                                         `json:"iam_instance_profile"`
	KeyName                      string                                         `json:"key_name"`
	VpcClassicLinkSecurityGroups string                                         `json:"vpc_classic_link_security_groups"`
}

type AwsLaunchConfigurationSpecRootBlockDevice struct {
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
}

type AwsLaunchConfigurationSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsLaunchConfigurationSpecEbsBlockDevice struct {
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	Encrypted           bool   `json:"encrypted"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	NoDevice            bool   `json:"no_device"`
	Iops                int    `json:"iops"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLaunchConfigurationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLaunchConfiguration `json"items"`
}
