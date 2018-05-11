// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecretsmanagerSecret struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsSecretsmanagerSecretSpec `json:"spec"`
}

type AwsSecretsmanagerSecretSpec struct {
	KmsKeyId             string                                     `json:"kms_key_id"`
	RotationEnabled      bool                                       `json:"rotation_enabled"`
	RotationRules        []AwsSecretsmanagerSecretSpecRotationRules `json:"rotation_rules"`
	Arn                  string                                     `json:"arn"`
	Description          string                                     `json:"description"`
	Name                 string                                     `json:"name"`
	RecoveryWindowInDays int                                        `json:"recovery_window_in_days"`
	RotationLambdaArn    string                                     `json:"rotation_lambda_arn"`
	Tags                 map[string]string                          `json:"tags"`
}

type AwsSecretsmanagerSecretSpecRotationRules struct {
	AutomaticallyAfterDays int `json:"automatically_after_days"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecretsmanagerSecretList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsSecretsmanagerSecret `json:"items"`
}
