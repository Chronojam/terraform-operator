// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsMountTarget struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEfsMountTargetSpec `json"spec"`
}

type AwsEfsMountTargetSpec struct {
	FileSystemId       string `json:"file_system_id"`
	IpAddress          string `json:"ip_address"`
	SecurityGroups     string `json:"security_groups"`
	SubnetId           string `json:"subnet_id"`
	NetworkInterfaceId string `json:"network_interface_id"`
	DnsName            string `json:"dns_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsMountTargetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEfsMountTarget `json"items"`
}
