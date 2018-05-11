// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegrationResponse struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayIntegrationResponseSpec `json"spec"`
}

type AwsApiGatewayIntegrationResponseSpec struct {
	ResponseTemplates        map[string]string `json:"response_templates"`
	ResponseParameters       map[string]string `json:"response_parameters"`
	ResponseParametersInJson string            `json:"response_parameters_in_json"`
	ContentHandling          string            `json:"content_handling"`
	RestApiId                string            `json:"rest_api_id"`
	ResourceId               string            `json:"resource_id"`
	SelectionPattern         string            `json:"selection_pattern"`
	HttpMethod               string            `json:"http_method"`
	StatusCode               string            `json:"status_code"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegrationResponseList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayIntegrationResponse `json"items"`
}
