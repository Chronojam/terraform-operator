// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayRestApi struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayRestApiSpec `json"spec"`
}

type AwsApiGatewayRestApiSpec struct {
	BinaryMediaTypes       []string `json:"binary_media_types"`
	Body                   string   `json:"body"`
	MinimumCompressionSize int      `json:"minimum_compression_size"`
	RootResourceId         string   `json:"root_resource_id"`
	CreatedDate            string   `json:"created_date"`
	Name                   string   `json:"name"`
	Description            string   `json:"description"`
	Policy                 string   `json:"policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayRestApiList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayRestApi `json"items"`
}
