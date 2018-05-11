// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsTaskDefinition struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsEcsTaskDefinitionSpec `json:"spec"`
}

type AwsEcsTaskDefinitionSpec struct {
	Memory                  string                                       `json:"memory"`
	NetworkMode             string                                       `json:"network_mode"`
	Volume                  AwsEcsTaskDefinitionSpecVolume               `json:"volume"`
	PlacementConstraints    AwsEcsTaskDefinitionSpecPlacementConstraints `json:"placement_constraints"`
	Arn                     string                                       `json:"arn"`
	Cpu                     string                                       `json:"cpu"`
	Revision                int                                          `json:"revision"`
	ExecutionRoleArn        string                                       `json:"execution_role_arn"`
	RequiresCompatibilities string                                       `json:"requires_compatibilities"`
	Family                  string                                       `json:"family"`
	ContainerDefinitions    string                                       `json:"container_definitions"`
	TaskRoleArn             string                                       `json:"task_role_arn"`
}

type AwsEcsTaskDefinitionSpecVolume struct {
	Name     string `json:"name"`
	HostPath string `json:"host_path"`
}

type AwsEcsTaskDefinitionSpecPlacementConstraints struct {
	Type       string `json:"type"`
	Expression string `json:"expression"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsTaskDefinitionList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsEcsTaskDefinition `json:"items"`
}
