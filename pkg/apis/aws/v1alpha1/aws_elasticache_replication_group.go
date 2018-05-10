// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheReplicationGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticacheReplicationGroupSpec `json"spec"`
}

type AwsElasticacheReplicationGroupSpec struct {
	NodeType                     string                                          `json:"node_type"`
	SubnetGroupName              string                                          `json:"subnet_group_name"`
	SnapshotArns                 string                                          `json:"snapshot_arns"`
	Engine                       string                                          `json:"engine"`
	ClusterMode                  []AwsElasticacheReplicationGroupSpecClusterMode `json:"cluster_mode"`
	ParameterGroupName           string                                          `json:"parameter_group_name"`
	SnapshotWindow               string                                          `json:"snapshot_window"`
	NotificationTopicArn         string                                          `json:"notification_topic_arn"`
	TransitEncryptionEnabled     bool                                            `json:"transit_encryption_enabled"`
	SecurityGroupNames           string                                          `json:"security_group_names"`
	Port                         int                                             `json:"port"`
	AutoMinorVersionUpgrade      bool                                            `json:"auto_minor_version_upgrade"`
	NumberCacheClusters          int                                             `json:"number_cache_clusters"`
	ConfigurationEndpointAddress string                                          `json:"configuration_endpoint_address"`
	MaintenanceWindow            string                                          `json:"maintenance_window"`
	ReplicationGroupId           string                                          `json:"replication_group_id"`
	ReplicationGroupDescription  string                                          `json:"replication_group_description"`
	AtRestEncryptionEnabled      bool                                            `json:"at_rest_encryption_enabled"`
	AuthToken                    string                                          `json:"auth_token"`
	EngineVersion                string                                          `json:"engine_version"`
	SecurityGroupIds             string                                          `json:"security_group_ids"`
	SnapshotRetentionLimit       int                                             `json:"snapshot_retention_limit"`
	Tags                         map[string]string                               `json:"tags"`
	ApplyImmediately             bool                                            `json:"apply_immediately"`
	AutomaticFailoverEnabled     bool                                            `json:"automatic_failover_enabled"`
	AvailabilityZones            string                                          `json:"availability_zones"`
	SnapshotName                 string                                          `json:"snapshot_name"`
	PrimaryEndpointAddress       string                                          `json:"primary_endpoint_address"`
}

type AwsElasticacheReplicationGroupSpecClusterMode struct {
	NumNodeGroups        int `json:"num_node_groups"`
	ReplicasPerNodeGroup int `json:"replicas_per_node_group"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheReplicationGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticacheReplicationGroup `json"items"`
}
