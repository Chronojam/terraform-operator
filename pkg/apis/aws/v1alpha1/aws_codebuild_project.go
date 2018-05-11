// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodebuildProject struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCodebuildProjectSpec `json"spec"`
}

type AwsCodebuildProjectSpec struct {
	VpcConfig     []AwsCodebuildProjectSpecVpcConfig `json:"vpc_config"`
	Cache         []AwsCodebuildProjectSpecCache     `json:"cache"`
	Description   string                             `json:"description"`
	BuildTimeout  int                                `json:"build_timeout"`
	Tags          map[string]string                  `json:"tags"`
	ServiceRole   string                             `json:"service_role"`
	Source        AwsCodebuildProjectSpecSource      `json:"source"`
	Timeout       int                                `json:"timeout"`
	Artifacts     AwsCodebuildProjectSpecArtifacts   `json:"artifacts"`
	EncryptionKey string                             `json:"encryption_key"`
	Environment   AwsCodebuildProjectSpecEnvironment `json:"environment"`
	Name          string                             `json:"name"`
}

type AwsCodebuildProjectSpecVpcConfig struct {
	Subnets          string `json:"subnets"`
	SecurityGroupIds string `json:"security_group_ids"`
	VpcId            string `json:"vpc_id"`
}

type AwsCodebuildProjectSpecCache struct {
	Type     string `json:"type"`
	Location string `json:"location"`
}

type AwsCodebuildProjectSpecSource struct {
	Auth      AwsCodebuildProjectSpecSourceAuth `json:"auth"`
	Buildspec string                            `json:"buildspec"`
	Location  string                            `json:"location"`
	Type      string                            `json:"type"`
}

type AwsCodebuildProjectSpecSourceAuth struct {
	Resource string `json:"resource"`
	Type     string `json:"type"`
}

type AwsCodebuildProjectSpecArtifacts struct {
	Name          string `json:"name"`
	Location      string `json:"location"`
	NamespaceType string `json:"namespace_type"`
	Packaging     string `json:"packaging"`
	Path          string `json:"path"`
	Type          string `json:"type"`
}

type AwsCodebuildProjectSpecEnvironment struct {
	ComputeType         string                                                  `json:"compute_type"`
	EnvironmentVariable []AwsCodebuildProjectSpecEnvironmentEnvironmentVariable `json:"environment_variable"`
	Image               string                                                  `json:"image"`
	Type                string                                                  `json:"type"`
	PrivilegedMode      bool                                                    `json:"privileged_mode"`
}

type AwsCodebuildProjectSpecEnvironmentEnvironmentVariable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodebuildProjectList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCodebuildProject `json"items"`
}
