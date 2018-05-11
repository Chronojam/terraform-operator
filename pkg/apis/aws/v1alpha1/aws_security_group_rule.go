// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroupRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSecurityGroupRuleSpec `json"spec"`
}

type AwsSecurityGroupRuleSpec struct {
	ToPort                int      `json:"to_port"`
	CidrBlocks            []string `json:"cidr_blocks"`
	Ipv6CidrBlocks        []string `json:"ipv6_cidr_blocks"`
	SourceSecurityGroupId string   `json:"source_security_group_id"`
	Self                  bool     `json:"self"`
	Type                  string   `json:"type"`
	FromPort              int      `json:"from_port"`
	Protocol              string   `json:"protocol"`
	PrefixListIds         []string `json:"prefix_list_ids"`
	SecurityGroupId       string   `json:"security_group_id"`
	Description           string   `json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroupRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSecurityGroupRule `json"items"`
}
