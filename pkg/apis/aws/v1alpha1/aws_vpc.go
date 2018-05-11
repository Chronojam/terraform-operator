// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpc struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcSpec `json"spec"`
}

type AwsVpcSpec struct {
	EnableDnsSupport             bool              `json:"enable_dns_support"`
	EnableClassiclinkDnsSupport  bool              `json:"enable_classiclink_dns_support"`
	Ipv6CidrBlock                string            `json:"ipv6_cidr_block"`
	InstanceTenancy              string            `json:"instance_tenancy"`
	EnableDnsHostnames           bool              `json:"enable_dns_hostnames"`
	Ipv6AssociationId            string            `json:"ipv6_association_id"`
	AssignGeneratedIpv6CidrBlock bool              `json:"assign_generated_ipv6_cidr_block"`
	DefaultNetworkAclId          string            `json:"default_network_acl_id"`
	DefaultRouteTableId          string            `json:"default_route_table_id"`
	Tags                         map[string]string `json:"tags"`
	CidrBlock                    string            `json:"cidr_block"`
	EnableClassiclink            bool              `json:"enable_classiclink"`
	MainRouteTableId             string            `json:"main_route_table_id"`
	DhcpOptionsId                string            `json:"dhcp_options_id"`
	DefaultSecurityGroupId       string            `json:"default_security_group_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpc `json"items"`
}
