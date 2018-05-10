// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSqsQueue struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSqsQueueSpec `json"spec"`
}

type AwsSqsQueueSpec struct {
	NamePrefix                   string            `json:"name_prefix"`
	VisibilityTimeoutSeconds     int               `json:"visibility_timeout_seconds"`
	RedrivePolicy                string            `json:"redrive_policy"`
	Arn                          string            `json:"arn"`
	FifoQueue                    bool              `json:"fifo_queue"`
	KmsDataKeyReusePeriodSeconds int               `json:"kms_data_key_reuse_period_seconds"`
	Name                         string            `json:"name"`
	DelaySeconds                 int               `json:"delay_seconds"`
	MaxMessageSize               int               `json:"max_message_size"`
	ReceiveWaitTimeSeconds       int               `json:"receive_wait_time_seconds"`
	Policy                       string            `json:"policy"`
	ContentBasedDeduplication    bool              `json:"content_based_deduplication"`
	KmsMasterKeyId               string            `json:"kms_master_key_id"`
	Tags                         map[string]string `json:"tags"`
	MessageRetentionSeconds      int               `json:"message_retention_seconds"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSqsQueueList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSqsQueue `json"items"`
}
