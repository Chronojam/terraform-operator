// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmAssociationSpec `json"spec"`
}

type AwsSsmAssociationSpec struct {
	Targets            []AwsSsmAssociationSpecTargets        `json:"targets"`
	AssociationName    string                                `json:"association_name"`
	AssociationId      string                                `json:"association_id"`
	InstanceId         string                                `json:"instance_id"`
	DocumentVersion    string                                `json:"document_version"`
	ScheduleExpression string                                `json:"schedule_expression"`
	Name               string                                `json:"name"`
	Parameters         map[string]string                     `json:"parameters"`
	OutputLocation     []AwsSsmAssociationSpecOutputLocation `json:"output_location"`
}

type AwsSsmAssociationSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type AwsSsmAssociationSpecOutputLocation struct {
	S3BucketName string `json:"s3_bucket_name"`
	S3KeyPrefix  string `json:"s3_key_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmAssociation `json"items"`
}
