// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationSubnetGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDmsReplicationSubnetGroupSpec `json"spec"`
}

type AwsDmsReplicationSubnetGroupSpec struct {
	ReplicationSubnetGroupArn         string            `json:"replication_subnet_group_arn"`
	ReplicationSubnetGroupDescription string            `json:"replication_subnet_group_description"`
	ReplicationSubnetGroupId          string            `json:"replication_subnet_group_id"`
	SubnetIds                         string            `json:"subnet_ids"`
	Tags                              map[string]string `json:"tags"`
	VpcId                             string            `json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationSubnetGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDmsReplicationSubnetGroup `json"items"`
}
