// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotFleetRequest struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSpotFleetRequestSpec `json"spec"`
}

type AwsSpotFleetRequestSpec struct {
	ReplaceUnhealthyInstances        bool                                       `json:"replace_unhealthy_instances"`
	ClientToken                      string                                     `json:"client_token"`
	SpotPrice                        string                                     `json:"spot_price"`
	SpotRequestState                 string                                     `json:"spot_request_state"`
	LaunchSpecification              AwsSpotFleetRequestSpecLaunchSpecification `json:"launch_specification"`
	TargetCapacity                   int                                        `json:"target_capacity"`
	AllocationStrategy               string                                     `json:"allocation_strategy"`
	ExcessCapacityTerminationPolicy  string                                     `json:"excess_capacity_termination_policy"`
	IamFleetRole                     string                                     `json:"iam_fleet_role"`
	TerminateInstancesWithExpiration bool                                       `json:"terminate_instances_with_expiration"`
	ValidUntil                       string                                     `json:"valid_until"`
	LoadBalancers                    string                                     `json:"load_balancers"`
	WaitForFulfillment               bool                                       `json:"wait_for_fulfillment"`
	InstanceInterruptionBehaviour    string                                     `json:"instance_interruption_behaviour"`
	ValidFrom                        string                                     `json:"valid_from"`
	TargetGroupArns                  string                                     `json:"target_group_arns"`
}

type AwsSpotFleetRequestSpecLaunchSpecification struct {
	RootBlockDevice          AwsSpotFleetRequestSpecLaunchSpecificationRootBlockDevice      `json:"root_block_device"`
	Monitoring               bool                                                           `json:"monitoring"`
	AvailabilityZone         string                                                         `json:"availability_zone"`
	VpcSecurityGroupIds      string                                                         `json:"vpc_security_group_ids"`
	IamInstanceProfile       string                                                         `json:"iam_instance_profile"`
	WeightedCapacity         string                                                         `json:"weighted_capacity"`
	KeyName                  string                                                         `json:"key_name"`
	SpotPrice                string                                                         `json:"spot_price"`
	SubnetId                 string                                                         `json:"subnet_id"`
	EphemeralBlockDevice     AwsSpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice `json:"ephemeral_block_device"`
	EbsOptimized             bool                                                           `json:"ebs_optimized"`
	Ami                      string                                                         `json:"ami"`
	PlacementGroup           string                                                         `json:"placement_group"`
	PlacementTenancy         string                                                         `json:"placement_tenancy"`
	UserData                 string                                                         `json:"user_data"`
	Tags                     map[string]string                                              `json:"tags"`
	AssociatePublicIpAddress bool                                                           `json:"associate_public_ip_address"`
	EbsBlockDevice           AwsSpotFleetRequestSpecLaunchSpecificationEbsBlockDevice       `json:"ebs_block_device"`
	InstanceType             string                                                         `json:"instance_type"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationEbsBlockDevice struct {
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotFleetRequestList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSpotFleetRequest `json"items"`
}
