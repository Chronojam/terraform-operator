// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbTable struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDynamodbTableSpec `json"spec"`
}

type AwsDynamodbTableSpec struct {
	HashKey              string                                     `json:"hash_key"`
	LocalSecondaryIndex  AwsDynamodbTableSpecLocalSecondaryIndex    `json:"local_secondary_index"`
	GlobalSecondaryIndex AwsDynamodbTableSpecGlobalSecondaryIndex   `json:"global_secondary_index"`
	ServerSideEncryption []AwsDynamodbTableSpecServerSideEncryption `json:"server_side_encryption"`
	WriteCapacity        int                                        `json:"write_capacity"`
	Ttl                  AwsDynamodbTableSpecTtl                    `json:"ttl"`
	StreamArn            string                                     `json:"stream_arn"`
	Tags                 map[string]string                          `json:"tags"`
	Attribute            AwsDynamodbTableSpecAttribute              `json:"attribute"`
	StreamEnabled        bool                                       `json:"stream_enabled"`
	StreamViewType       string                                     `json:"stream_view_type"`
	StreamLabel          string                                     `json:"stream_label"`
	PointInTimeRecovery  []AwsDynamodbTableSpecPointInTimeRecovery  `json:"point_in_time_recovery"`
	Arn                  string                                     `json:"arn"`
	Name                 string                                     `json:"name"`
	RangeKey             string                                     `json:"range_key"`
	ReadCapacity         int                                        `json:"read_capacity"`
}

type AwsDynamodbTableSpecLocalSecondaryIndex struct {
	NonKeyAttributes []string `json:"non_key_attributes"`
	Name             string   `json:"name"`
	RangeKey         string   `json:"range_key"`
	ProjectionType   string   `json:"projection_type"`
}

type AwsDynamodbTableSpecGlobalSecondaryIndex struct {
	NonKeyAttributes []string `json:"non_key_attributes"`
	Name             string   `json:"name"`
	WriteCapacity    int      `json:"write_capacity"`
	ReadCapacity     int      `json:"read_capacity"`
	HashKey          string   `json:"hash_key"`
	RangeKey         string   `json:"range_key"`
	ProjectionType   string   `json:"projection_type"`
}

type AwsDynamodbTableSpecServerSideEncryption struct {
	Enabled bool `json:"enabled"`
}

type AwsDynamodbTableSpecTtl struct {
	AttributeName string `json:"attribute_name"`
	Enabled       bool   `json:"enabled"`
}

type AwsDynamodbTableSpecAttribute struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AwsDynamodbTableSpecPointInTimeRecovery struct {
	Enabled bool `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbTableList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDynamodbTable `json"items"`
}
