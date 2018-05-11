// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksApplication struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksApplicationSpec `json"spec"`
}

type AwsOpsworksApplicationSpec struct {
	ShortName              string                                       `json:"short_name"`
	DataSourceType         string                                       `json:"data_source_type"`
	DataSourceArn          string                                       `json:"data_source_arn"`
	EnableSsl              bool                                         `json:"enable_ssl"`
	SslConfiguration       []AwsOpsworksApplicationSpecSslConfiguration `json:"ssl_configuration"`
	StackId                string                                       `json:"stack_id"`
	AutoBundleOnDeploy     string                                       `json:"auto_bundle_on_deploy"`
	Environment            AwsOpsworksApplicationSpecEnvironment        `json:"environment"`
	Name                   string                                       `json:"name"`
	Type                   string                                       `json:"type"`
	DocumentRoot           string                                       `json:"document_root"`
	RailsEnv               string                                       `json:"rails_env"`
	AppSource              []AwsOpsworksApplicationSpecAppSource        `json:"app_source"`
	AwsFlowRubySettings    string                                       `json:"aws_flow_ruby_settings"`
	DataSourceDatabaseName string                                       `json:"data_source_database_name"`
	Description            string                                       `json:"description"`
	Domains                []string                                     `json:"domains"`
}

type AwsOpsworksApplicationSpecSslConfiguration struct {
	Chain       string `json:"chain"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

type AwsOpsworksApplicationSpecEnvironment struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Secure bool   `json:"secure"`
}

type AwsOpsworksApplicationSpecAppSource struct {
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksApplicationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksApplication `json"items"`
}
