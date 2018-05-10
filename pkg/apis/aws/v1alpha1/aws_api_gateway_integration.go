// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegration struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayIntegrationSpec `json"spec"`
}

type AwsApiGatewayIntegrationSpec struct {
	IntegrationHttpMethod   string            `json:"integration_http_method"`
	RequestTemplates        map[string]string `json:"request_templates"`
	RequestParametersInJson string            `json:"request_parameters_in_json"`
	ContentHandling         string            `json:"content_handling"`
	RestApiId               string            `json:"rest_api_id"`
	Type                    string            `json:"type"`
	ConnectionType          string            `json:"connection_type"`
	Credentials             string            `json:"credentials"`
	HttpMethod              string            `json:"http_method"`
	CacheNamespace          string            `json:"cache_namespace"`
	PassthroughBehavior     string            `json:"passthrough_behavior"`
	CacheKeyParameters      string            `json:"cache_key_parameters"`
	ResourceId              string            `json:"resource_id"`
	ConnectionId            string            `json:"connection_id"`
	Uri                     string            `json:"uri"`
	RequestParameters       map[string]string `json:"request_parameters"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegrationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayIntegration `json"items"`
}
