// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppsyncDatasource struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAppsyncDatasourceSpec `json"spec"`
}

type AwsAppsyncDatasourceSpec struct {
	Name                string                                        `json:"name"`
	DynamodbConfig      []AwsAppsyncDatasourceSpecDynamodbConfig      `json:"dynamodb_config"`
	ElasticsearchConfig []AwsAppsyncDatasourceSpecElasticsearchConfig `json:"elasticsearch_config"`
	LambdaConfig        []AwsAppsyncDatasourceSpecLambdaConfig        `json:"lambda_config"`
	Arn                 string                                        `json:"arn"`
	ApiId               string                                        `json:"api_id"`
	Description         string                                        `json:"description"`
	ServiceRoleArn      string                                        `json:"service_role_arn"`
	Type                string                                        `json:"type"`
}

type AwsAppsyncDatasourceSpecDynamodbConfig struct {
	UseCallerCredentials bool   `json:"use_caller_credentials"`
	Region               string `json:"region"`
	TableName            string `json:"table_name"`
}

type AwsAppsyncDatasourceSpecElasticsearchConfig struct {
	Region   string `json:"region"`
	Endpoint string `json:"endpoint"`
}

type AwsAppsyncDatasourceSpecLambdaConfig struct {
	FunctionArn string `json:"function_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppsyncDatasourceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAppsyncDatasource `json"items"`
}
