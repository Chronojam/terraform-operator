// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowTask struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsSsmMaintenanceWindowTaskSpec `json:"spec"`
}

type AwsSsmMaintenanceWindowTaskSpec struct {
	MaxConcurrency string                                          `json:"max_concurrency"`
	TaskType       string                                          `json:"task_type"`
	TaskArn        string                                          `json:"task_arn"`
	Priority       int                                             `json:"priority"`
	LoggingInfo    []AwsSsmMaintenanceWindowTaskSpecLoggingInfo    `json:"logging_info"`
	TaskParameters []AwsSsmMaintenanceWindowTaskSpecTaskParameters `json:"task_parameters"`
	WindowId       string                                          `json:"window_id"`
	MaxErrors      string                                          `json:"max_errors"`
	ServiceRoleArn string                                          `json:"service_role_arn"`
	Targets        []AwsSsmMaintenanceWindowTaskSpecTargets        `json:"targets"`
}

type AwsSsmMaintenanceWindowTaskSpecLoggingInfo struct {
	S3BucketName   string `json:"s3_bucket_name"`
	S3Region       string `json:"s3_region"`
	S3BucketPrefix string `json:"s3_bucket_prefix"`
}

type AwsSsmMaintenanceWindowTaskSpecTaskParameters struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type AwsSsmMaintenanceWindowTaskSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowTaskList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsSsmMaintenanceWindowTask `json:"items"`
}
