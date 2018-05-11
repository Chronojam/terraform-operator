// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigDeliveryChannel struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsConfigDeliveryChannelSpec `json"spec"`
}

type AwsConfigDeliveryChannelSpec struct {
	S3KeyPrefix                string                                                   `json:"s3_key_prefix"`
	SnsTopicArn                string                                                   `json:"sns_topic_arn"`
	SnapshotDeliveryProperties []AwsConfigDeliveryChannelSpecSnapshotDeliveryProperties `json:"snapshot_delivery_properties"`
	Name                       string                                                   `json:"name"`
	S3BucketName               string                                                   `json:"s3_bucket_name"`
}

type AwsConfigDeliveryChannelSpecSnapshotDeliveryProperties struct {
	DeliveryFrequency string `json:"delivery_frequency"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigDeliveryChannelList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsConfigDeliveryChannel `json"items"`
}
