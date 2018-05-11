// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudformationStack struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsCloudformationStackSpec `json:"spec"`
}

type AwsCloudformationStackSpec struct {
	TimeoutInMinutes int               `json:"timeout_in_minutes"`
	Tags             map[string]string `json:"tags"`
	TemplateUrl      string            `json:"template_url"`
	OnFailure        string            `json:"on_failure"`
	Parameters       map[string]string `json:"parameters"`
	PolicyBody       string            `json:"policy_body"`
	PolicyUrl        string            `json:"policy_url"`
	TemplateBody     string            `json:"template_body"`
	Capabilities     string            `json:"capabilities"`
	Name             string            `json:"name"`
	NotificationArns string            `json:"notification_arns"`
	Outputs          map[string]string `json:"outputs"`
	IamRoleArn       string            `json:"iam_role_arn"`
	DisableRollback  bool              `json:"disable_rollback"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudformationStackList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsCloudformationStack `json:"items"`
}
