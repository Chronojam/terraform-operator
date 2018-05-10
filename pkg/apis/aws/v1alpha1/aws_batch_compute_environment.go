// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchComputeEnvironment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsBatchComputeEnvironmentSpec `json"spec"`
}

type AwsBatchComputeEnvironmentSpec struct {
	EccClusterArn          string                                           `json:"ecc_cluster_arn"`
	Status                 string                                           `json:"status"`
	StatusReason           string                                           `json:"status_reason"`
	ComputeEnvironmentName string                                           `json:"compute_environment_name"`
	State                  string                                           `json:"state"`
	Type                   string                                           `json:"type"`
	Arn                    string                                           `json:"arn"`
	ComputeResources       []AwsBatchComputeEnvironmentSpecComputeResources `json:"compute_resources"`
	ServiceRole            string                                           `json:"service_role"`
	EcsClusterArn          string                                           `json:"ecs_cluster_arn"`
}

type AwsBatchComputeEnvironmentSpecComputeResources struct {
	Ec2KeyPair       string            `json:"ec2_key_pair"`
	ImageId          string            `json:"image_id"`
	InstanceRole     string            `json:"instance_role"`
	InstanceType     string            `json:"instance_type"`
	MinVcpus         int               `json:"min_vcpus"`
	SecurityGroupIds string            `json:"security_group_ids"`
	BidPercentage    int               `json:"bid_percentage"`
	DesiredVcpus     int               `json:"desired_vcpus"`
	SpotIamFleetRole string            `json:"spot_iam_fleet_role"`
	Tags             map[string]string `json:"tags"`
	Type             string            `json:"type"`
	MaxVcpus         int               `json:"max_vcpus"`
	Subnets          string            `json:"subnets"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchComputeEnvironmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsBatchComputeEnvironment `json"items"`
}
