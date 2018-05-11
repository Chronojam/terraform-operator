// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafXssMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafXssMatchSetSpec `json"spec"`
}

type AwsWafXssMatchSetSpec struct {
	Name           string                              `json:"name"`
	XssMatchTuples AwsWafXssMatchSetSpecXssMatchTuples `json:"xss_match_tuples"`
}

type AwsWafXssMatchSetSpecXssMatchTuples struct {
	TextTransformation string                                          `json:"text_transformation"`
	FieldToMatch       AwsWafXssMatchSetSpecXssMatchTuplesFieldToMatch `json:"field_to_match"`
}

type AwsWafXssMatchSetSpecXssMatchTuplesFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafXssMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafXssMatchSet `json"items"`
}
