// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSfnStateMachine struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSfnStateMachineSpec `json"spec"`
}

type AwsSfnStateMachineSpec struct {
	Name         string `json:"name"`
	RoleArn      string `json:"role_arn"`
	CreationDate string `json:"creation_date"`
	Status       string `json:"status"`
	Definition   string `json:"definition"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSfnStateMachineList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSfnStateMachine `json"items"`
}
