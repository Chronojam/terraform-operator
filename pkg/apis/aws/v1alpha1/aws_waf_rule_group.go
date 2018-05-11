// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRuleGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafRuleGroupSpec `json"spec"`
}

type AwsWafRuleGroupSpec struct {
	Name          string                           `json:"name"`
	MetricName    string                           `json:"metric_name"`
	ActivatedRule AwsWafRuleGroupSpecActivatedRule `json:"activated_rule"`
}

type AwsWafRuleGroupSpecActivatedRule struct {
	Action   []AwsWafRuleGroupSpecActivatedRuleAction `json:"action"`
	Priority int                                      `json:"priority"`
	RuleId   string                                   `json:"rule_id"`
	Type     string                                   `json:"type"`
}

type AwsWafRuleGroupSpecActivatedRuleAction struct {
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRuleGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafRuleGroup `json"items"`
}
