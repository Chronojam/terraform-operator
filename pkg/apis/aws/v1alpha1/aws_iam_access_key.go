// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccessKey struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsIamAccessKeySpec `json:"spec"`
}

type AwsIamAccessKeySpec struct {
	Status          string `json:"status"`
	Secret          string `json:"secret"`
	SesSmtpPassword string `json:"ses_smtp_password"`
	PgpKey          string `json:"pgp_key"`
	KeyFingerprint  string `json:"key_fingerprint"`
	EncryptedSecret string `json:"encrypted_secret"`
	User            string `json:"user"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccessKeyList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsIamAccessKey `json:"items"`
}
