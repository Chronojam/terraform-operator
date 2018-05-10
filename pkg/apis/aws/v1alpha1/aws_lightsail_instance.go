// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailInstance struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLightsailInstanceSpec `json"spec"`
}

type AwsLightsailInstanceSpec struct {
	KeyPairName      string `json:"key_pair_name"`
	UserData         string `json:"user_data"`
	Username         string `json:"username"`
	Name             string `json:"name"`
	BlueprintId      string `json:"blueprint_id"`
	IsStaticIp       bool   `json:"is_static_ip"`
	PublicIpAddress  string `json:"public_ip_address"`
	AvailabilityZone string `json:"availability_zone"`
	CreatedAt        string `json:"created_at"`
	BundleId         string `json:"bundle_id"`
	RamSize          int    `json:"ram_size"`
	Ipv6Address      string `json:"ipv6_address"`
	PrivateIpAddress string `json:"private_ip_address"`
	Arn              string `json:"arn"`
	CpuCount         int    `json:"cpu_count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailInstanceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLightsailInstance `json"items"`
}
