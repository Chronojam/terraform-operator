// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbEventSubscription struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDbEventSubscriptionSpec `json"spec"`
}

type AwsDbEventSubscriptionSpec struct {
	CustomerAwsId   string            `json:"customer_aws_id"`
	SnsTopic        string            `json:"sns_topic"`
	EventCategories string            `json:"event_categories"`
	SourceIds       string            `json:"source_ids"`
	SourceType      string            `json:"source_type"`
	Arn             string            `json:"arn"`
	Name            string            `json:"name"`
	Enabled         bool              `json:"enabled"`
	Tags            map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbEventSubscriptionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDbEventSubscription `json"items"`
}
