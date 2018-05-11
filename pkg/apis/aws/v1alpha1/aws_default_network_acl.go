// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultNetworkAcl struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsDefaultNetworkAclSpec `json:"spec"`
}

type AwsDefaultNetworkAclSpec struct {
	VpcId               string                          `json:"vpc_id"`
	DefaultNetworkAclId string                          `json:"default_network_acl_id"`
	SubnetIds           string                          `json:"subnet_ids"`
	Ingress             AwsDefaultNetworkAclSpecIngress `json:"ingress"`
	Egress              AwsDefaultNetworkAclSpecEgress  `json:"egress"`
	Tags                map[string]string               `json:"tags"`
}

type AwsDefaultNetworkAclSpecIngress struct {
	ToPort        int    `json:"to_port"`
	RuleNo        int    `json:"rule_no"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	CidrBlock     string `json:"cidr_block"`
	IcmpType      int    `json:"icmp_type"`
	IcmpCode      int    `json:"icmp_code"`
	FromPort      int    `json:"from_port"`
	Action        string `json:"action"`
	Protocol      string `json:"protocol"`
}

type AwsDefaultNetworkAclSpecEgress struct {
	IcmpType      int    `json:"icmp_type"`
	FromPort      int    `json:"from_port"`
	ToPort        int    `json:"to_port"`
	RuleNo        int    `json:"rule_no"`
	Action        string `json:"action"`
	CidrBlock     string `json:"cidr_block"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	Protocol      string `json:"protocol"`
	IcmpCode      int    `json:"icmp_code"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultNetworkAclList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsDefaultNetworkAcl `json:"items"`
}
