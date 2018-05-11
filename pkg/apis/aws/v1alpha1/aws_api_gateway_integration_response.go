// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegrationResponse struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsApiGatewayIntegrationResponseSpec `json:"spec"`
}

type AwsApiGatewayIntegrationResponseSpec struct {
	HttpMethod               string            `json:"http_method"`
	SelectionPattern         string            `json:"selection_pattern"`
	ResponseParametersInJson string            `json:"response_parameters_in_json"`
	ContentHandling          string            `json:"content_handling"`
	RestApiId                string            `json:"rest_api_id"`
	ResourceId               string            `json:"resource_id"`
	StatusCode               string            `json:"status_code"`
	ResponseTemplates        map[string]string `json:"response_templates"`
	ResponseParameters       map[string]string `json:"response_parameters"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegrationResponseList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsApiGatewayIntegrationResponse `json:"items"`
}
