// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheReplicationGroup struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsElasticacheReplicationGroupSpec `json:"spec"`
}

type AwsElasticacheReplicationGroupSpec struct {
	ParameterGroupName           string                                          `json:"parameter_group_name"`
	SubnetGroupName              string                                          `json:"subnet_group_name"`
	ApplyImmediately             bool                                            `json:"apply_immediately"`
	ConfigurationEndpointAddress string                                          `json:"configuration_endpoint_address"`
	AuthToken                    string                                          `json:"auth_token"`
	Engine                       string                                          `json:"engine"`
	EngineVersion                string                                          `json:"engine_version"`
	SecurityGroupIds             string                                          `json:"security_group_ids"`
	NotificationTopicArn         string                                          `json:"notification_topic_arn"`
	Tags                         map[string]string                               `json:"tags"`
	NodeType                     string                                          `json:"node_type"`
	SecurityGroupNames           string                                          `json:"security_group_names"`
	SnapshotWindow               string                                          `json:"snapshot_window"`
	ReplicationGroupId           string                                          `json:"replication_group_id"`
	ClusterMode                  []AwsElasticacheReplicationGroupSpecClusterMode `json:"cluster_mode"`
	AvailabilityZones            string                                          `json:"availability_zones"`
	MaintenanceWindow            string                                          `json:"maintenance_window"`
	Port                         int                                             `json:"port"`
	SnapshotName                 string                                          `json:"snapshot_name"`
	AutoMinorVersionUpgrade      bool                                            `json:"auto_minor_version_upgrade"`
	NumberCacheClusters          int                                             `json:"number_cache_clusters"`
	TransitEncryptionEnabled     bool                                            `json:"transit_encryption_enabled"`
	SnapshotArns                 string                                          `json:"snapshot_arns"`
	SnapshotRetentionLimit       int                                             `json:"snapshot_retention_limit"`
	ReplicationGroupDescription  string                                          `json:"replication_group_description"`
	AtRestEncryptionEnabled      bool                                            `json:"at_rest_encryption_enabled"`
	AutomaticFailoverEnabled     bool                                            `json:"automatic_failover_enabled"`
	PrimaryEndpointAddress       string                                          `json:"primary_endpoint_address"`
}

type AwsElasticacheReplicationGroupSpecClusterMode struct {
	ReplicasPerNodeGroup int `json:"replicas_per_node_group"`
	NumNodeGroups        int `json:"num_node_groups"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheReplicationGroupList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsElasticacheReplicationGroup `json:"items"`
}
