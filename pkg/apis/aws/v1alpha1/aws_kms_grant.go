// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsGrant struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsKmsGrantSpec `json"spec"`
}

type AwsKmsGrantSpec struct {
	Name                string                     `json:"name"`
	KeyId               string                     `json:"key_id"`
	RetiringPrincipal   string                     `json:"retiring_principal"`
	GrantCreationTokens string                     `json:"grant_creation_tokens"`
	RetireOnDelete      bool                       `json:"retire_on_delete"`
	GrantId             string                     `json:"grant_id"`
	GrantToken          string                     `json:"grant_token"`
	GranteePrincipal    string                     `json:"grantee_principal"`
	Operations          string                     `json:"operations"`
	Constraints         AwsKmsGrantSpecConstraints `json:"constraints"`
}

type AwsKmsGrantSpecConstraints struct {
	EncryptionContextEquals map[string]string `json:"encryption_context_equals"`
	EncryptionContextSubset map[string]string `json:"encryption_context_subset"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsGrantList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsKmsGrant `json"items"`
}
