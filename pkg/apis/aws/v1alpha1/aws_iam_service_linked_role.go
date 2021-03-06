// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamServiceLinkedRole struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsIamServiceLinkedRoleSpec `json:"spec"`
}

type AwsIamServiceLinkedRoleSpec struct {
	Path           string `json:"path"`
	Arn            string `json:"arn"`
	CreateDate     string `json:"create_date"`
	UniqueId       string `json:"unique_id"`
	CustomSuffix   string `json:"custom_suffix"`
	Description    string `json:"description"`
	AwsServiceName string `json:"aws_service_name"`
	Name           string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamServiceLinkedRoleList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsIamServiceLinkedRole `json:"items"`
}
