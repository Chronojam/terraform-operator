// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsFileSystem struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEfsFileSystemSpec `json"spec"`
}

type AwsEfsFileSystemSpec struct {
	CreationToken   string            `json:"creation_token"`
	ReferenceName   string            `json:"reference_name"`
	PerformanceMode string            `json:"performance_mode"`
	Encrypted       bool              `json:"encrypted"`
	KmsKeyId        string            `json:"kms_key_id"`
	DnsName         string            `json:"dns_name"`
	Tags            map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsFileSystemList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEfsFileSystem `json"items"`
}
