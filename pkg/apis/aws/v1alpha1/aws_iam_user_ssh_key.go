// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserSshKey struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsIamUserSshKeySpec `json:"spec"`
}

type AwsIamUserSshKeySpec struct {
	SshPublicKeyId string `json:"ssh_public_key_id"`
	Fingerprint    string `json:"fingerprint"`
	Username       string `json:"username"`
	PublicKey      string `json:"public_key"`
	Encoding       string `json:"encoding"`
	Status         string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserSshKeyList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsIamUserSshKey `json:"items"`
}
