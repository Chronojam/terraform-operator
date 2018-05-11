// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafSizeConstraintSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafSizeConstraintSetSpec `json"spec"`
}

type AwsWafSizeConstraintSetSpec struct {
	Name            string                                     `json:"name"`
	SizeConstraints AwsWafSizeConstraintSetSpecSizeConstraints `json:"size_constraints"`
}

type AwsWafSizeConstraintSetSpecSizeConstraints struct {
	FieldToMatch       AwsWafSizeConstraintSetSpecSizeConstraintsFieldToMatch `json:"field_to_match"`
	ComparisonOperator string                                                 `json:"comparison_operator"`
	Size               int                                                    `json:"size"`
	TextTransformation string                                                 `json:"text_transformation"`
}

type AwsWafSizeConstraintSetSpecSizeConstraintsFieldToMatch struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafSizeConstraintSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafSizeConstraintSet `json"items"`
}
