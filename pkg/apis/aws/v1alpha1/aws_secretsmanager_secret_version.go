// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecretsmanagerSecretVersion struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSecretsmanagerSecretVersionSpec `json"spec"`
}

type AwsSecretsmanagerSecretVersionSpec struct {
	VersionStages string `json:"version_stages"`
	SecretId      string `json:"secret_id"`
	SecretString  string `json:"secret_string"`
	VersionId     string `json:"version_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecretsmanagerSecretVersionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSecretsmanagerSecretVersion `json"items"`
}
