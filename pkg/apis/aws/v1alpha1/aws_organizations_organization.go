// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsOrganization struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOrganizationsOrganizationSpec `json"spec"`
}

type AwsOrganizationsOrganizationSpec struct {
	Arn                string `json:"arn"`
	MasterAccountArn   string `json:"master_account_arn"`
	MasterAccountEmail string `json:"master_account_email"`
	MasterAccountId    string `json:"master_account_id"`
	FeatureSet         string `json:"feature_set"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsOrganizationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOrganizationsOrganization `json"items"`
}
