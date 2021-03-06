// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserLoginProfile struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsIamUserLoginProfileSpec `json:"spec"`
}

type AwsIamUserLoginProfileSpec struct {
	User                  string `json:"user"`
	PgpKey                string `json:"pgp_key"`
	PasswordResetRequired bool   `json:"password_reset_required"`
	PasswordLength        int    `json:"password_length"`
	KeyFingerprint        string `json:"key_fingerprint"`
	EncryptedPassword     string `json:"encrypted_password"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserLoginProfileList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsIamUserLoginProfile `json:"items"`
}
