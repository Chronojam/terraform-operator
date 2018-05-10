// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRouteSpec `json"spec"`
}

type AwsRouteSpec struct {
	DestinationPrefixListId  string `json:"destination_prefix_list_id"`
	GatewayId                string `json:"gateway_id"`
	EgressOnlyGatewayId      string `json:"egress_only_gateway_id"`
	NatGatewayId             string `json:"nat_gateway_id"`
	InstanceOwnerId          string `json:"instance_owner_id"`
	State                    string `json:"state"`
	DestinationCidrBlock     string `json:"destination_cidr_block"`
	DestinationIpv6CidrBlock string `json:"destination_ipv6_cidr_block"`
	InstanceId               string `json:"instance_id"`
	NetworkInterfaceId       string `json:"network_interface_id"`
	Origin                   string `json:"origin"`
	RouteTableId             string `json:"route_table_id"`
	VpcPeeringConnectionId   string `json:"vpc_peering_connection_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRouteList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRoute `json"items"`
}
