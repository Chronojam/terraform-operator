// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailKeyPair struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLightsailKeyPairSpec `json"spec"`
}

type AwsLightsailKeyPairSpec struct {
	NamePrefix           string `json:"name_prefix"`
	Arn                  string `json:"arn"`
	Fingerprint          string `json:"fingerprint"`
	PublicKey            string `json:"public_key"`
	EncryptedPrivateKey  string `json:"encrypted_private_key"`
	Name                 string `json:"name"`
	PrivateKey           string `json:"private_key"`
	EncryptedFingerprint string `json:"encrypted_fingerprint"`
	PgpKey               string `json:"pgp_key"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailKeyPairList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLightsailKeyPair `json"items"`
}
