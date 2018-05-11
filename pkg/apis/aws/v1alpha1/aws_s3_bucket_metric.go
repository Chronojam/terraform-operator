// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketMetric struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsS3BucketMetricSpec `json"spec"`
}

type AwsS3BucketMetricSpec struct {
	Bucket string                        `json:"bucket"`
	Filter []AwsS3BucketMetricSpecFilter `json:"filter"`
	Name   string                        `json:"name"`
}

type AwsS3BucketMetricSpecFilter struct {
	Tags   map[string]string `json:"tags"`
	Prefix string            `json:"prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketMetricList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsS3BucketMetric `json"items"`
}
