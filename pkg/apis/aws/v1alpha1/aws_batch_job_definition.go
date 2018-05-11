// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchJobDefinition struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsBatchJobDefinitionSpec `json"spec"`
}

type AwsBatchJobDefinitionSpec struct {
	RetryStrategy       []AwsBatchJobDefinitionSpecRetryStrategy `json:"retry_strategy"`
	Timeout             []AwsBatchJobDefinitionSpecTimeout       `json:"timeout"`
	Type                string                                   `json:"type"`
	Revision            int                                      `json:"revision"`
	Arn                 string                                   `json:"arn"`
	Name                string                                   `json:"name"`
	ContainerProperties string                                   `json:"container_properties"`
	Parameters          map[string]string                        `json:"parameters"`
}

type AwsBatchJobDefinitionSpecRetryStrategy struct {
	Attempts int `json:"attempts"`
}

type AwsBatchJobDefinitionSpecTimeout struct {
	AttemptDurationSeconds int `json:"attempt_duration_seconds"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchJobDefinitionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsBatchJobDefinition `json"items"`
}
