// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayStage struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayStageSpec `json"spec"`
}

type AwsApiGatewayStageSpec struct {
	CacheClusterSize     string                                    `json:"cache_cluster_size"`
	ClientCertificateId  string                                    `json:"client_certificate_id"`
	DeploymentId         string                                    `json:"deployment_id"`
	Description          string                                    `json:"description"`
	ExecutionArn         string                                    `json:"execution_arn"`
	InvokeUrl            string                                    `json:"invoke_url"`
	CacheClusterEnabled  bool                                      `json:"cache_cluster_enabled"`
	DocumentationVersion string                                    `json:"documentation_version"`
	RestApiId            string                                    `json:"rest_api_id"`
	StageName            string                                    `json:"stage_name"`
	Variables            map[string]string                         `json:"variables"`
	Tags                 map[string]string                         `json:"tags"`
	AccessLogSettings    []AwsApiGatewayStageSpecAccessLogSettings `json:"access_log_settings"`
}

type AwsApiGatewayStageSpecAccessLogSettings struct {
	Format         string `json:"format"`
	DestinationArn string `json:"destination_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayStageList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayStage `json"items"`
}
