// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMqConfiguration struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsMqConfigurationSpec `json"spec"`
}

type AwsMqConfigurationSpec struct {
	Arn            string `json:"arn"`
	Data           string `json:"data"`
	Description    string `json:"description"`
	EngineType     string `json:"engine_type"`
	EngineVersion  string `json:"engine_version"`
	Name           string `json:"name"`
	LatestRevision int    `json:"latest_revision"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMqConfigurationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsMqConfiguration `json"items"`
}
