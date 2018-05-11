// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSubnet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSubnetSpec `json"spec"`
}

type AwsSubnetSpec struct {
	Ipv6CidrBlock               string            `json:"ipv6_cidr_block"`
	AvailabilityZone            string            `json:"availability_zone"`
	MapPublicIpOnLaunch         bool              `json:"map_public_ip_on_launch"`
	AssignIpv6AddressOnCreation bool              `json:"assign_ipv6_address_on_creation"`
	Ipv6CidrBlockAssociationId  string            `json:"ipv6_cidr_block_association_id"`
	Tags                        map[string]string `json:"tags"`
	VpcId                       string            `json:"vpc_id"`
	CidrBlock                   string            `json:"cidr_block"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSubnetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSubnet `json"items"`
}
