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
	EcsClusterArn          string                                           `json:"ecs_cluster_arn"`
	Status                 string                                           `json:"status"`
	State                  string                                           `json:"state"`
	ComputeResources       []AwsBatchComputeEnvironmentSpecComputeResources `json:"compute_resources"`
	ServiceRole            string                                           `json:"service_role"`
	Type                   string                                           `json:"type"`
	Arn                    string                                           `json:"arn"`
	StatusReason           string                                           `json:"status_reason"`
	ComputeEnvironmentName string                                           `json:"compute_environment_name"`
}

type AwsBatchComputeEnvironmentSpecComputeResources struct {
	BidPercentage    int               `json:"bid_percentage"`
	DesiredVcpus     int               `json:"desired_vcpus"`
	Ec2KeyPair       string            `json:"ec2_key_pair"`
	InstanceType     string            `json:"instance_type"`
	MinVcpus         int               `json:"min_vcpus"`
	SecurityGroupIds string            `json:"security_group_ids"`
	Type             string            `json:"type"`
	ImageId          string            `json:"image_id"`
	InstanceRole     string            `json:"instance_role"`
	MaxVcpus         int               `json:"max_vcpus"`
	SpotIamFleetRole string            `json:"spot_iam_fleet_role"`
	Subnets          string            `json:"subnets"`
	Tags             map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchComputeEnvironmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsBatchComputeEnvironment `json"items"`
}
