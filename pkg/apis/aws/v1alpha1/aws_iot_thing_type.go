// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotThingType struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIotThingTypeSpec `json"spec"`
}

type AwsIotThingTypeSpec struct {
	Arn        string                          `json:"arn"`
	Name       string                          `json:"name"`
	Properties []AwsIotThingTypeSpecProperties `json:"properties"`
	Deprecated bool                            `json:"deprecated"`
}

type AwsIotThingTypeSpecProperties struct {
	Description          string `json:"description"`
	SearchableAttributes string `json:"searchable_attributes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotThingTypeList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIotThingType `json"items"`
}
