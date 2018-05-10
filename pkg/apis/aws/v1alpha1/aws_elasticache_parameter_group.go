// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheParameterGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticacheParameterGroupSpec `json"spec"`
}

type AwsElasticacheParameterGroupSpec struct {
	Name        string                                    `json:"name"`
	Family      string                                    `json:"family"`
	Description string                                    `json:"description"`
	Parameter   AwsElasticacheParameterGroupSpecParameter `json:"parameter"`
}

type AwsElasticacheParameterGroupSpecParameter struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheParameterGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticacheParameterGroup `json"items"`
}
