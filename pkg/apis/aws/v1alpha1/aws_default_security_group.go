// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultSecurityGroup struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsDefaultSecurityGroupSpec `json:"spec"`
}

type AwsDefaultSecurityGroupSpec struct {
	Arn                 string                             `json:"arn"`
	OwnerId             string                             `json:"owner_id"`
	Tags                map[string]string                  `json:"tags"`
	RevokeRulesOnDelete bool                               `json:"revoke_rules_on_delete"`
	Name                string                             `json:"name"`
	VpcId               string                             `json:"vpc_id"`
	Ingress             AwsDefaultSecurityGroupSpecIngress `json:"ingress"`
	Egress              AwsDefaultSecurityGroupSpecEgress  `json:"egress"`
}

type AwsDefaultSecurityGroupSpecIngress struct {
	ToPort         int      `json:"to_port"`
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	SecurityGroups string   `json:"security_groups"`
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
}

type AwsDefaultSecurityGroupSpecEgress struct {
	Protocol       string   `json:"protocol"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
	ToPort         int      `json:"to_port"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	SecurityGroups string   `json:"security_groups"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultSecurityGroupList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsDefaultSecurityGroup `json:"items"`
}
