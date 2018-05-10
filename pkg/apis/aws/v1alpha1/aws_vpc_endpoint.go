// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpoint struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcEndpointSpec `json"spec"`
}

type AwsVpcEndpointSpec struct {
	CidrBlocks          []string                     `json:"cidr_blocks"`
	ServiceName         string                       `json:"service_name"`
	RouteTableIds       string                       `json:"route_table_ids"`
	SubnetIds           string                       `json:"subnet_ids"`
	PrivateDnsEnabled   bool                         `json:"private_dns_enabled"`
	VpcId               string                       `json:"vpc_id"`
	Policy              string                       `json:"policy"`
	State               string                       `json:"state"`
	PrefixListId        string                       `json:"prefix_list_id"`
	DnsEntry            []AwsVpcEndpointSpecDnsEntry `json:"dns_entry"`
	AutoAccept          bool                         `json:"auto_accept"`
	VpcEndpointType     string                       `json:"vpc_endpoint_type"`
	SecurityGroupIds    string                       `json:"security_group_ids"`
	NetworkInterfaceIds string                       `json:"network_interface_ids"`
}

type AwsVpcEndpointSpecDnsEntry struct {
	DnsName      string `json:"dns_name"`
	HostedZoneId string `json:"hosted_zone_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcEndpoint `json"items"`
}
