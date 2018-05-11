// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsKey struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsKmsKeySpec `json:"spec"`
}

type AwsKmsKeySpec struct {
	Tags                 map[string]string `json:"tags"`
	KeyId                string            `json:"key_id"`
	Description          string            `json:"description"`
	KeyUsage             string            `json:"key_usage"`
	Policy               string            `json:"policy"`
	IsEnabled            bool              `json:"is_enabled"`
	DeletionWindowInDays int               `json:"deletion_window_in_days"`
	Arn                  string            `json:"arn"`
	EnableKeyRotation    bool              `json:"enable_key_rotation"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsKeyList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsKmsKey `json:"items"`
}
