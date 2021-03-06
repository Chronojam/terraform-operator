// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroup struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsSecurityGroupSpec `json:"spec"`
}

type AwsSecurityGroupSpec struct {
	OwnerId             string                      `json:"owner_id"`
	Tags                map[string]string           `json:"tags"`
	RevokeRulesOnDelete bool                        `json:"revoke_rules_on_delete"`
	Name                string                      `json:"name"`
	NamePrefix          string                      `json:"name_prefix"`
	Description         string                      `json:"description"`
	Egress              AwsSecurityGroupSpecEgress  `json:"egress"`
	Arn                 string                      `json:"arn"`
	VpcId               string                      `json:"vpc_id"`
	Ingress             AwsSecurityGroupSpecIngress `json:"ingress"`
}

type AwsSecurityGroupSpecEgress struct {
	ToPort         int      `json:"to_port"`
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	SecurityGroups string   `json:"security_groups"`
}

type AwsSecurityGroupSpecIngress struct {
	ToPort         int      `json:"to_port"`
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	SecurityGroups string   `json:"security_groups"`
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroupList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsSecurityGroup `json:"items"`
}
