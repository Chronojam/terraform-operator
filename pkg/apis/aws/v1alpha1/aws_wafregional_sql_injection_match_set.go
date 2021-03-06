// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalSqlInjectionMatchSet struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsWafregionalSqlInjectionMatchSetSpec `json:"spec"`
}

type AwsWafregionalSqlInjectionMatchSetSpec struct {
	Name                   string                                                       `json:"name"`
	SqlInjectionMatchTuple AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTuple `json:"sql_injection_match_tuple"`
}

type AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTuple struct {
	FieldToMatch       []AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch `json:"field_to_match"`
	TextTransformation string                                                                     `json:"text_transformation"`
}

type AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalSqlInjectionMatchSetList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsWafregionalSqlInjectionMatchSet `json:"items"`
}
