// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationInstance struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsDmsReplicationInstanceSpec `json:"spec"`
}

type AwsDmsReplicationInstanceSpec struct {
	AutoMinorVersionUpgrade       bool              `json:"auto_minor_version_upgrade"`
	MultiAz                       bool              `json:"multi_az"`
	PubliclyAccessible            bool              `json:"publicly_accessible"`
	ReplicationInstanceArn        string            `json:"replication_instance_arn"`
	ReplicationInstanceId         string            `json:"replication_instance_id"`
	ReplicationInstancePrivateIps []string          `json:"replication_instance_private_ips"`
	ReplicationSubnetGroupId      string            `json:"replication_subnet_group_id"`
	ApplyImmediately              bool              `json:"apply_immediately"`
	VpcSecurityGroupIds           string            `json:"vpc_security_group_ids"`
	KmsKeyArn                     string            `json:"kms_key_arn"`
	ReplicationInstancePublicIps  []string          `json:"replication_instance_public_ips"`
	EngineVersion                 string            `json:"engine_version"`
	PreferredMaintenanceWindow    string            `json:"preferred_maintenance_window"`
	AvailabilityZone              string            `json:"availability_zone"`
	ReplicationInstanceClass      string            `json:"replication_instance_class"`
	Tags                          map[string]string `json:"tags"`
	AllocatedStorage              int               `json:"allocated_storage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationInstanceList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsDmsReplicationInstance `json:"items"`
}
