// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayModel struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayModelSpec `json"spec"`
}

type AwsApiGatewayModelSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Schema      string `json:"schema"`
	ContentType string `json:"content_type"`
	RestApiId   string `json:"rest_api_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayModelList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayModel `json"items"`
}
