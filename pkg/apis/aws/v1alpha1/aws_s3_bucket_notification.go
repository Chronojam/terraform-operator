// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketNotification struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsS3BucketNotificationSpec `json"spec"`
}

type AwsS3BucketNotificationSpec struct {
	Queue          []AwsS3BucketNotificationSpecQueue          `json:"queue"`
	LambdaFunction []AwsS3BucketNotificationSpecLambdaFunction `json:"lambda_function"`
	Bucket         string                                      `json:"bucket"`
	Topic          []AwsS3BucketNotificationSpecTopic          `json:"topic"`
}

type AwsS3BucketNotificationSpecQueue struct {
	FilterPrefix string `json:"filter_prefix"`
	FilterSuffix string `json:"filter_suffix"`
	QueueArn     string `json:"queue_arn"`
	Events       string `json:"events"`
	Id           string `json:"id"`
}

type AwsS3BucketNotificationSpecLambdaFunction struct {
	FilterSuffix      string `json:"filter_suffix"`
	LambdaFunctionArn string `json:"lambda_function_arn"`
	Events            string `json:"events"`
	Id                string `json:"id"`
	FilterPrefix      string `json:"filter_prefix"`
}

type AwsS3BucketNotificationSpecTopic struct {
	Events       string `json:"events"`
	Id           string `json:"id"`
	FilterPrefix string `json:"filter_prefix"`
	FilterSuffix string `json:"filter_suffix"`
	TopicArn     string `json:"topic_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketNotificationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsS3BucketNotification `json"items"`
}
