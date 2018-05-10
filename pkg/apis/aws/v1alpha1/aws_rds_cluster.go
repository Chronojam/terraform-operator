// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsCluster struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRdsClusterSpec `json"spec"`
}

type AwsRdsClusterSpec struct {
	HostedZoneId                     string                      `json:"hosted_zone_id"`
	FinalSnapshotIdentifier          string                      `json:"final_snapshot_identifier"`
	ClusterIdentifier                string                      `json:"cluster_identifier"`
	ClusterIdentifierPrefix          string                      `json:"cluster_identifier_prefix"`
	DatabaseName                     string                      `json:"database_name"`
	Engine                           string                      `json:"engine"`
	S3Import                         []AwsRdsClusterSpecS3Import `json:"s3_import"`
	MasterPassword                   string                      `json:"master_password"`
	SnapshotIdentifier               string                      `json:"snapshot_identifier"`
	ApplyImmediately                 bool                        `json:"apply_immediately"`
	KmsKeyId                         string                      `json:"kms_key_id"`
	DbClusterParameterGroupName      string                      `json:"db_cluster_parameter_group_name"`
	StorageEncrypted                 bool                        `json:"storage_encrypted"`
	Tags                             map[string]string           `json:"tags"`
	ClusterMembers                   string                      `json:"cluster_members"`
	EngineVersion                    string                      `json:"engine_version"`
	VpcSecurityGroupIds              string                      `json:"vpc_security_group_ids"`
	ClusterResourceId                string                      `json:"cluster_resource_id"`
	AvailabilityZones                string                      `json:"availability_zones"`
	MasterUsername                   string                      `json:"master_username"`
	Port                             int                         `json:"port"`
	IamDatabaseAuthenticationEnabled bool                        `json:"iam_database_authentication_enabled"`
	SourceRegion                     string                      `json:"source_region"`
	BackupRetentionPeriod            int                         `json:"backup_retention_period"`
	DbSubnetGroupName                string                      `json:"db_subnet_group_name"`
	PreferredBackupWindow            string                      `json:"preferred_backup_window"`
	PreferredMaintenanceWindow       string                      `json:"preferred_maintenance_window"`
	Endpoint                         string                      `json:"endpoint"`
	ReaderEndpoint                   string                      `json:"reader_endpoint"`
	SkipFinalSnapshot                bool                        `json:"skip_final_snapshot"`
	ReplicationSourceIdentifier      string                      `json:"replication_source_identifier"`
	IamRoles                         string                      `json:"iam_roles"`
}

type AwsRdsClusterSpecS3Import struct {
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsClusterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRdsCluster `json"items"`
}
