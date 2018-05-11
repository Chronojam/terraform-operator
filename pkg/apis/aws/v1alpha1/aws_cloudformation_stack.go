// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudformationStack struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudformationStackSpec `json"spec"`
}

type AwsCloudformationStackSpec struct {
	Tags             map[string]string `json:"tags"`
	TemplateUrl      string            `json:"template_url"`
	DisableRollback  bool              `json:"disable_rollback"`
	Parameters       map[string]string `json:"parameters"`
	PolicyBody       string            `json:"policy_body"`
	TimeoutInMinutes int               `json:"timeout_in_minutes"`
	Capabilities     string            `json:"capabilities"`
	PolicyUrl        string            `json:"policy_url"`
	IamRoleArn       string            `json:"iam_role_arn"`
	Name             string            `json:"name"`
	TemplateBody     string            `json:"template_body"`
	Outputs          map[string]string `json:"outputs"`
	NotificationArns string            `json:"notification_arns"`
	OnFailure        string            `json:"on_failure"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudformationStackList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudformationStack `json"items"`
}
