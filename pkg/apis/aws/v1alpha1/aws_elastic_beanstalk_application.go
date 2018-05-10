// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkApplication struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticBeanstalkApplicationSpec `json"spec"`
}

type AwsElasticBeanstalkApplicationSpec struct {
	AppversionLifecycle []AwsElasticBeanstalkApplicationSpecAppversionLifecycle `json:"appversion_lifecycle"`
	Name                string                                                  `json:"name"`
	Description         string                                                  `json:"description"`
}

type AwsElasticBeanstalkApplicationSpecAppversionLifecycle struct {
	DeleteSourceFromS3 bool   `json:"delete_source_from_s3"`
	ServiceRole        string `json:"service_role"`
	MaxAgeInDays       int    `json:"max_age_in_days"`
	MaxCount           int    `json:"max_count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkApplicationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticBeanstalkApplication `json"items"`
}
