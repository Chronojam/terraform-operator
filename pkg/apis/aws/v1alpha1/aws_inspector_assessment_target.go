// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorAssessmentTarget struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsInspectorAssessmentTargetSpec `json:"spec"`
}

type AwsInspectorAssessmentTargetSpec struct {
	Arn              string `json:"arn"`
	ResourceGroupArn string `json:"resource_group_arn"`
	Name             string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorAssessmentTargetList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsInspectorAssessmentTarget `json:"items"`
}
