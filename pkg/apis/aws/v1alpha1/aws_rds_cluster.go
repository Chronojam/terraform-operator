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
	Endpoint                         string                      `json:"endpoint"`
	PreferredMaintenanceWindow       string                      `json:"preferred_maintenance_window"`
	AvailabilityZones                string                      `json:"availability_zones"`
	PreferredBackupWindow            string                      `json:"preferred_backup_window"`
	IamDatabaseAuthenticationEnabled bool                        `json:"iam_database_authentication_enabled"`
	DbSubnetGroupName                string                      `json:"db_subnet_group_name"`
	MasterUsername                   string                      `json:"master_username"`
	Port                             int                         `json:"port"`
	ClusterResourceId                string                      `json:"cluster_resource_id"`
	ReaderEndpoint                   string                      `json:"reader_endpoint"`
	IamRoles                         string                      `json:"iam_roles"`
	SourceRegion                     string                      `json:"source_region"`
	HostedZoneId                     string                      `json:"hosted_zone_id"`
	ClusterMembers                   string                      `json:"cluster_members"`
	DatabaseName                     string                      `json:"database_name"`
	StorageEncrypted                 bool                        `json:"storage_encrypted"`
	FinalSnapshotIdentifier          string                      `json:"final_snapshot_identifier"`
	MasterPassword                   string                      `json:"master_password"`
	ApplyImmediately                 bool                        `json:"apply_immediately"`
	VpcSecurityGroupIds              string                      `json:"vpc_security_group_ids"`
	ClusterIdentifierPrefix          string                      `json:"cluster_identifier_prefix"`
	SnapshotIdentifier               string                      `json:"snapshot_identifier"`
	BackupRetentionPeriod            int                         `json:"backup_retention_period"`
	KmsKeyId                         string                      `json:"kms_key_id"`
	Tags                             map[string]string           `json:"tags"`
	DbClusterParameterGroupName      string                      `json:"db_cluster_parameter_group_name"`
	Engine                           string                      `json:"engine"`
	EngineVersion                    string                      `json:"engine_version"`
	S3Import                         []AwsRdsClusterSpecS3Import `json:"s3_import"`
	SkipFinalSnapshot                bool                        `json:"skip_final_snapshot"`
	ReplicationSourceIdentifier      string                      `json:"replication_source_identifier"`
	ClusterIdentifier                string                      `json:"cluster_identifier"`
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
