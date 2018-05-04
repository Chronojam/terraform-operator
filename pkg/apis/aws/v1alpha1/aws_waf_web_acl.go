// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafWebAcl struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafWebAclSpec `json"spec"`
}

type AwsWafWebAclSpec struct {
	Name          string                        `json:"name"`
	DefaultAction AwsWafWebAclSpecDefaultAction `json:"default_action"`
	MetricName    string                        `json:"metric_name"`
	Rules         AwsWafWebAclSpecRules         `json:"rules"`
}

type AwsWafWebAclSpecDefaultAction struct {
	Type string `json:"type"`
}

type AwsWafWebAclSpecRules struct {
	Action   AwsWafWebAclSpecRulesAction `json:"action"`
	Priority int                         `json:"priority"`
	Type     string                      `json:"type"`
	RuleId   string                      `json:"rule_id"`
}

type AwsWafWebAclSpecRulesAction struct {
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafWebAclList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafWebAcl `json"items"`
}
