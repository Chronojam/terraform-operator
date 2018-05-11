// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudtrail struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudtrailSpec `json"spec"`
}

type AwsCloudtrailSpec struct {
	S3KeyPrefix                string                           `json:"s3_key_prefix"`
	EnableLogFileValidation    bool                             `json:"enable_log_file_validation"`
	Tags                       map[string]string                `json:"tags"`
	S3BucketName               string                           `json:"s3_bucket_name"`
	IsMultiRegionTrail         bool                             `json:"is_multi_region_trail"`
	EnableLogging              bool                             `json:"enable_logging"`
	CloudWatchLogsRoleArn      string                           `json:"cloud_watch_logs_role_arn"`
	SnsTopicName               string                           `json:"sns_topic_name"`
	HomeRegion                 string                           `json:"home_region"`
	Name                       string                           `json:"name"`
	IncludeGlobalServiceEvents bool                             `json:"include_global_service_events"`
	KmsKeyId                   string                           `json:"kms_key_id"`
	EventSelector              []AwsCloudtrailSpecEventSelector `json:"event_selector"`
	Arn                        string                           `json:"arn"`
	CloudWatchLogsGroupArn     string                           `json:"cloud_watch_logs_group_arn"`
}

type AwsCloudtrailSpecEventSelector struct {
	ReadWriteType           string                                       `json:"read_write_type"`
	IncludeManagementEvents bool                                         `json:"include_management_events"`
	DataResource            []AwsCloudtrailSpecEventSelectorDataResource `json:"data_resource"`
}

type AwsCloudtrailSpecEventSelectorDataResource struct {
	Values []string `json:"values"`
	Type   string   `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudtrailList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudtrail `json"items"`
}
