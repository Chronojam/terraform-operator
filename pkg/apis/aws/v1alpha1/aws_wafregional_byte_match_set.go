// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalByteMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalByteMatchSetSpec `json"spec"`
}

type AwsWafregionalByteMatchSetSpec struct {
	Name           string                                       `json:"name"`
	ByteMatchTuple AwsWafregionalByteMatchSetSpecByteMatchTuple `json:"byte_match_tuple"`
}

type AwsWafregionalByteMatchSetSpecByteMatchTuple struct {
	PositionalConstraint string                                                   `json:"positional_constraint"`
	TargetString         string                                                   `json:"target_string"`
	TextTransformation   string                                                   `json:"text_transformation"`
	FieldToMatch         AwsWafregionalByteMatchSetSpecByteMatchTupleFieldToMatch `json:"field_to_match"`
}

type AwsWafregionalByteMatchSetSpecByteMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalByteMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalByteMatchSet `json"items"`
}
