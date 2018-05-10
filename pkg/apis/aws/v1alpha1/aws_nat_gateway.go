// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNatGateway struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsNatGatewaySpec `json"spec"`
}

type AwsNatGatewaySpec struct {
	PrivateIp          string            `json:"private_ip"`
	PublicIp           string            `json:"public_ip"`
	Tags               map[string]string `json:"tags"`
	AllocationId       string            `json:"allocation_id"`
	SubnetId           string            `json:"subnet_id"`
	NetworkInterfaceId string            `json:"network_interface_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNatGatewayList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsNatGateway `json"items"`
}
