// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLbListenerRuleSpec `json"spec"`
}

type AwsLbListenerRuleSpec struct {
	Arn         string                         `json:"arn"`
	ListenerArn string                         `json:"listener_arn"`
	Priority    int                            `json:"priority"`
	Action      []AwsLbListenerRuleSpecAction  `json:"action"`
	Condition   AwsLbListenerRuleSpecCondition `json:"condition"`
}

type AwsLbListenerRuleSpecAction struct {
	TargetGroupArn string `json:"target_group_arn"`
	Type           string `json:"type"`
}

type AwsLbListenerRuleSpecCondition struct {
	Field  string   `json:"field"`
	Values []string `json:"values"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLbListenerRule `json"items"`
}
