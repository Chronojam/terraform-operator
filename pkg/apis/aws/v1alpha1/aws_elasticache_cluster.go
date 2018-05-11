// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheCluster struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticacheClusterSpec `json"spec"`
}

type AwsElasticacheClusterSpec struct {
	NodeType               string                                `json:"node_type"`
	EngineVersion          string                                `json:"engine_version"`
	ParameterGroupName     string                                `json:"parameter_group_name"`
	SnapshotName           string                                `json:"snapshot_name"`
	MaintenanceWindow      string                                `json:"maintenance_window"`
	NotificationTopicArn   string                                `json:"notification_topic_arn"`
	NumCacheNodes          int                                   `json:"num_cache_nodes"`
	CacheNodes             []AwsElasticacheClusterSpecCacheNodes `json:"cache_nodes"`
	AvailabilityZones      string                                `json:"availability_zones"`
	SecurityGroupNames     string                                `json:"security_group_names"`
	SnapshotWindow         string                                `json:"snapshot_window"`
	SnapshotRetentionLimit int                                   `json:"snapshot_retention_limit"`
	ApplyImmediately       bool                                  `json:"apply_immediately"`
	Tags                   map[string]string                     `json:"tags"`
	ConfigurationEndpoint  string                                `json:"configuration_endpoint"`
	ReplicationGroupId     string                                `json:"replication_group_id"`
	AvailabilityZone       string                                `json:"availability_zone"`
	Engine                 string                                `json:"engine"`
	SubnetGroupName        string                                `json:"subnet_group_name"`
	SecurityGroupIds       string                                `json:"security_group_ids"`
	SnapshotArns           string                                `json:"snapshot_arns"`
	Port                   int                                   `json:"port"`
	ClusterId              string                                `json:"cluster_id"`
	AzMode                 string                                `json:"az_mode"`
	ClusterAddress         string                                `json:"cluster_address"`
}

type AwsElasticacheClusterSpecCacheNodes struct {
	Id               string `json:"id"`
	Address          string `json:"address"`
	Port             int    `json:"port"`
	AvailabilityZone string `json:"availability_zone"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheClusterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticacheCluster `json"items"`
}
