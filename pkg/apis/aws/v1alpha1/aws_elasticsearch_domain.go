// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticsearchDomain struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsElasticsearchDomainSpec `json:"spec"`
}

type AwsElasticsearchDomainSpec struct {
	ClusterConfig        []AwsElasticsearchDomainSpecClusterConfig      `json:"cluster_config"`
	SnapshotOptions      []AwsElasticsearchDomainSpecSnapshotOptions    `json:"snapshot_options"`
	VpcOptions           []AwsElasticsearchDomainSpecVpcOptions         `json:"vpc_options"`
	AccessPolicies       string                                         `json:"access_policies"`
	DomainId             string                                         `json:"domain_id"`
	LogPublishingOptions AwsElasticsearchDomainSpecLogPublishingOptions `json:"log_publishing_options"`
	Tags                 map[string]string                              `json:"tags"`
	DomainName           string                                         `json:"domain_name"`
	EncryptAtRest        []AwsElasticsearchDomainSpecEncryptAtRest      `json:"encrypt_at_rest"`
	ElasticsearchVersion string                                         `json:"elasticsearch_version"`
	EbsOptions           []AwsElasticsearchDomainSpecEbsOptions         `json:"ebs_options"`
	Arn                  string                                         `json:"arn"`
	Endpoint             string                                         `json:"endpoint"`
	KibanaEndpoint       string                                         `json:"kibana_endpoint"`
	AdvancedOptions      map[string]string                              `json:"advanced_options"`
}

type AwsElasticsearchDomainSpecClusterConfig struct {
	DedicatedMasterCount   int    `json:"dedicated_master_count"`
	DedicatedMasterEnabled bool   `json:"dedicated_master_enabled"`
	DedicatedMasterType    string `json:"dedicated_master_type"`
	InstanceCount          int    `json:"instance_count"`
	InstanceType           string `json:"instance_type"`
	ZoneAwarenessEnabled   bool   `json:"zone_awareness_enabled"`
}

type AwsElasticsearchDomainSpecSnapshotOptions struct {
	AutomatedSnapshotStartHour int `json:"automated_snapshot_start_hour"`
}

type AwsElasticsearchDomainSpecVpcOptions struct {
	AvailabilityZones string `json:"availability_zones"`
	SecurityGroupIds  string `json:"security_group_ids"`
	SubnetIds         string `json:"subnet_ids"`
	VpcId             string `json:"vpc_id"`
}

type AwsElasticsearchDomainSpecLogPublishingOptions struct {
	LogType               string `json:"log_type"`
	CloudwatchLogGroupArn string `json:"cloudwatch_log_group_arn"`
	Enabled               bool   `json:"enabled"`
}

type AwsElasticsearchDomainSpecEncryptAtRest struct {
	Enabled  bool   `json:"enabled"`
	KmsKeyId string `json:"kms_key_id"`
}

type AwsElasticsearchDomainSpecEbsOptions struct {
	EbsEnabled bool   `json:"ebs_enabled"`
	Iops       int    `json:"iops"`
	VolumeSize int    `json:"volume_size"`
	VolumeType string `json:"volume_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticsearchDomainList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsElasticsearchDomain `json:"items"`
}
