// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowTarget struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsSsmMaintenanceWindowTargetSpec `json:"spec"`
}

type AwsSsmMaintenanceWindowTargetSpec struct {
	WindowId         string                                     `json:"window_id"`
	ResourceType     string                                     `json:"resource_type"`
	Targets          []AwsSsmMaintenanceWindowTargetSpecTargets `json:"targets"`
	OwnerInformation string                                     `json:"owner_information"`
}

type AwsSsmMaintenanceWindowTargetSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowTargetList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsSsmMaintenanceWindowTarget `json:"items"`
}
