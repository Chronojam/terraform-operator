// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53Zone struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRoute53ZoneSpec `json"spec"`
}

type AwsRoute53ZoneSpec struct {
	VpcId           string            `json:"vpc_id"`
	VpcRegion       string            `json:"vpc_region"`
	ZoneId          string            `json:"zone_id"`
	ForceDestroy    bool              `json:"force_destroy"`
	Name            string            `json:"name"`
	Comment         string            `json:"comment"`
	Tags            map[string]string `json:"tags"`
	DelegationSetId string            `json:"delegation_set_id"`
	NameServers     []string          `json:"name_servers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53ZoneList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRoute53Zone `json"items"`
}
