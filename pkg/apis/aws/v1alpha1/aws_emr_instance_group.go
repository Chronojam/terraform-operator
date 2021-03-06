// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrInstanceGroup struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsEmrInstanceGroupSpec `json:"spec"`
}

type AwsEmrInstanceGroupSpec struct {
	RunningInstanceCount int                              `json:"running_instance_count"`
	Status               string                           `json:"status"`
	Name                 string                           `json:"name"`
	EbsOptimized         bool                             `json:"ebs_optimized"`
	EbsConfig            AwsEmrInstanceGroupSpecEbsConfig `json:"ebs_config"`
	ClusterId            string                           `json:"cluster_id"`
	InstanceType         string                           `json:"instance_type"`
	InstanceCount        int                              `json:"instance_count"`
}

type AwsEmrInstanceGroupSpecEbsConfig struct {
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int    `json:"volumes_per_instance"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrInstanceGroupList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsEmrInstanceGroup `json:"items"`
}
