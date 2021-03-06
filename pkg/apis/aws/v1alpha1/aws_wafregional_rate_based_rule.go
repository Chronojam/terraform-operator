// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRateBasedRule struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsWafregionalRateBasedRuleSpec `json:"spec"`
}

type AwsWafregionalRateBasedRuleSpec struct {
	Predicate  AwsWafregionalRateBasedRuleSpecPredicate `json:"predicate"`
	RateKey    string                                   `json:"rate_key"`
	RateLimit  int                                      `json:"rate_limit"`
	Name       string                                   `json:"name"`
	MetricName string                                   `json:"metric_name"`
}

type AwsWafregionalRateBasedRuleSpecPredicate struct {
	Negated bool   `json:"negated"`
	DataId  string `json:"data_id"`
	Type    string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRateBasedRuleList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsWafregionalRateBasedRule `json:"items"`
}
