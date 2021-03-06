// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethodResponse struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsApiGatewayMethodResponseSpec `json:"spec"`
}

type AwsApiGatewayMethodResponseSpec struct {
	HttpMethod               string            `json:"http_method"`
	StatusCode               string            `json:"status_code"`
	ResponseModels           map[string]string `json:"response_models"`
	ResponseParameters       map[string]string `json:"response_parameters"`
	ResponseParametersInJson string            `json:"response_parameters_in_json"`
	RestApiId                string            `json:"rest_api_id"`
	ResourceId               string            `json:"resource_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethodResponseList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsApiGatewayMethodResponse `json:"items"`
}
