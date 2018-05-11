// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEip struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEipSpec `json"spec"`
}

type AwsEipSpec struct {
	NetworkInterface       string            `json:"network_interface"`
	Domain                 string            `json:"domain"`
	PublicIp               string            `json:"public_ip"`
	Tags                   map[string]string `json:"tags"`
	Vpc                    bool              `json:"vpc"`
	Instance               string            `json:"instance"`
	AllocationId           string            `json:"allocation_id"`
	AssociationId          string            `json:"association_id"`
	PrivateIp              string            `json:"private_ip"`
	AssociateWithPrivateIp string            `json:"associate_with_private_ip"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEipList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEip `json"items"`
}
