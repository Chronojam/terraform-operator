// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInstance struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsInstanceSpec `json"spec"`
}

type AwsInstanceSpec struct {
	DisableApiTermination             bool                                 `json:"disable_api_termination"`
	PlacementGroup                    string                               `json:"placement_group"`
	PublicIp                          string                               `json:"public_ip"`
	InstanceInitiatedShutdownBehavior string                               `json:"instance_initiated_shutdown_behavior"`
	Monitoring                        bool                                 `json:"monitoring"`
	BlockDevice                       map[string]string                    `json:"block_device"`
	KeyName                           string                               `json:"key_name"`
	PrivateDns                        string                               `json:"private_dns"`
	Ami                               string                               `json:"ami"`
	NetworkInterface                  AwsInstanceSpecNetworkInterface      `json:"network_interface"`
	Tags                              map[string]string                    `json:"tags"`
	RootBlockDevice                   []AwsInstanceSpecRootBlockDevice     `json:"root_block_device"`
	UserData                          string                               `json:"user_data"`
	Ipv6AddressCount                  int                                  `json:"ipv6_address_count"`
	SecurityGroups                    string                               `json:"security_groups"`
	Ipv6Addresses                     []string                             `json:"ipv6_addresses"`
	EbsBlockDevice                    AwsInstanceSpecEbsBlockDevice        `json:"ebs_block_device"`
	EphemeralBlockDevice              AwsInstanceSpecEphemeralBlockDevice  `json:"ephemeral_block_device"`
	AssociatePublicIpAddress          bool                                 `json:"associate_public_ip_address"`
	SourceDestCheck                   bool                                 `json:"source_dest_check"`
	EbsOptimized                      bool                                 `json:"ebs_optimized"`
	IamInstanceProfile                string                               `json:"iam_instance_profile"`
	VolumeTags                        map[string]string                    `json:"volume_tags"`
	CreditSpecification               []AwsInstanceSpecCreditSpecification `json:"credit_specification"`
	InstanceType                      string                               `json:"instance_type"`
	UserDataBase64                    string                               `json:"user_data_base64"`
	PublicDns                         string                               `json:"public_dns"`
	NetworkInterfaceId                string                               `json:"network_interface_id"`
	PrimaryNetworkInterfaceId         string                               `json:"primary_network_interface_id"`
	InstanceState                     string                               `json:"instance_state"`
	GetPasswordData                   bool                                 `json:"get_password_data"`
	SubnetId                          string                               `json:"subnet_id"`
	PrivateIp                         string                               `json:"private_ip"`
	VpcSecurityGroupIds               string                               `json:"vpc_security_group_ids"`
	Tenancy                           string                               `json:"tenancy"`
	AvailabilityZone                  string                               `json:"availability_zone"`
	PasswordData                      string                               `json:"password_data"`
}

type AwsInstanceSpecNetworkInterface struct {
	DeviceIndex         int    `json:"device_index"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	NetworkInterfaceId  string `json:"network_interface_id"`
}

type AwsInstanceSpecRootBlockDevice struct {
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type AwsInstanceSpecEbsBlockDevice struct {
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
}

type AwsInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
	NoDevice    bool   `json:"no_device"`
}

type AwsInstanceSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInstanceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsInstance `json"items"`
}
