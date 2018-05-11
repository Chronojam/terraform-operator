// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRole struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamRoleSpec `json"spec"`
}

type AwsIamRoleSpec struct {
	Arn                 string `json:"arn"`
	UniqueId            string `json:"unique_id"`
	Path                string `json:"path"`
	Description         string `json:"description"`
	AssumeRolePolicy    string `json:"assume_role_policy"`
	MaxSessionDuration  int    `json:"max_session_duration"`
	Name                string `json:"name"`
	NamePrefix          string `json:"name_prefix"`
	ForceDetachPolicies bool   `json:"force_detach_policies"`
	CreateDate          string `json:"create_date"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRoleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamRole `json"items"`
}
