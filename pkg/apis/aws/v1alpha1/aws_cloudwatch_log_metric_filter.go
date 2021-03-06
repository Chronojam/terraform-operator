// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogMetricFilter struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsCloudwatchLogMetricFilterSpec `json:"spec"`
}

type AwsCloudwatchLogMetricFilterSpec struct {
	Name                 string                                                 `json:"name"`
	Pattern              string                                                 `json:"pattern"`
	LogGroupName         string                                                 `json:"log_group_name"`
	MetricTransformation []AwsCloudwatchLogMetricFilterSpecMetricTransformation `json:"metric_transformation"`
}

type AwsCloudwatchLogMetricFilterSpecMetricTransformation struct {
	Name         string  `json:"name"`
	Namespace    string  `json:"namespace"`
	Value        string  `json:"value"`
	DefaultValue float64 `json:"default_value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogMetricFilterList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsCloudwatchLogMetricFilter `json:"items"`
}
