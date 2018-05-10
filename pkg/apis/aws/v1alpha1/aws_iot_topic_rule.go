// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotTopicRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIotTopicRuleSpec `json"spec"`
}

type AwsIotTopicRuleSpec struct {
	CloudwatchAlarm  AwsIotTopicRuleSpecCloudwatchAlarm  `json:"cloudwatch_alarm"`
	Elasticsearch    AwsIotTopicRuleSpecElasticsearch    `json:"elasticsearch"`
	Name             string                              `json:"name"`
	SqlVersion       string                              `json:"sql_version"`
	CloudwatchMetric AwsIotTopicRuleSpecCloudwatchMetric `json:"cloudwatch_metric"`
	Dynamodb         AwsIotTopicRuleSpecDynamodb         `json:"dynamodb"`
	Firehose         AwsIotTopicRuleSpecFirehose         `json:"firehose"`
	Arn              string                              `json:"arn"`
	Description      string                              `json:"description"`
	Enabled          bool                                `json:"enabled"`
	Kinesis          AwsIotTopicRuleSpecKinesis          `json:"kinesis"`
	Lambda           AwsIotTopicRuleSpecLambda           `json:"lambda"`
	Sqs              AwsIotTopicRuleSpecSqs              `json:"sqs"`
	Sql              string                              `json:"sql"`
	Republish        AwsIotTopicRuleSpecRepublish        `json:"republish"`
	S3               AwsIotTopicRuleSpecS3               `json:"s3"`
	Sns              AwsIotTopicRuleSpecSns              `json:"sns"`
}

type AwsIotTopicRuleSpecCloudwatchAlarm struct {
	AlarmName   string `json:"alarm_name"`
	RoleArn     string `json:"role_arn"`
	StateReason string `json:"state_reason"`
	StateValue  string `json:"state_value"`
}

type AwsIotTopicRuleSpecElasticsearch struct {
	RoleArn  string `json:"role_arn"`
	Type     string `json:"type"`
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
	Index    string `json:"index"`
}

type AwsIotTopicRuleSpecCloudwatchMetric struct {
	MetricName      string `json:"metric_name"`
	MetricNamespace string `json:"metric_namespace"`
	MetricTimestamp string `json:"metric_timestamp"`
	MetricUnit      string `json:"metric_unit"`
	MetricValue     string `json:"metric_value"`
	RoleArn         string `json:"role_arn"`
}

type AwsIotTopicRuleSpecDynamodb struct {
	RangeKeyValue string `json:"range_key_value"`
	TableName     string `json:"table_name"`
	HashKeyType   string `json:"hash_key_type"`
	HashKeyValue  string `json:"hash_key_value"`
	PayloadField  string `json:"payload_field"`
	RangeKeyField string `json:"range_key_field"`
	RangeKeyType  string `json:"range_key_type"`
	RoleArn       string `json:"role_arn"`
	HashKeyField  string `json:"hash_key_field"`
}

type AwsIotTopicRuleSpecFirehose struct {
	DeliveryStreamName string `json:"delivery_stream_name"`
	RoleArn            string `json:"role_arn"`
}

type AwsIotTopicRuleSpecKinesis struct {
	PartitionKey string `json:"partition_key"`
	RoleArn      string `json:"role_arn"`
	StreamName   string `json:"stream_name"`
}

type AwsIotTopicRuleSpecLambda struct {
	FunctionArn string `json:"function_arn"`
}

type AwsIotTopicRuleSpecSqs struct {
	QueueUrl  string `json:"queue_url"`
	RoleArn   string `json:"role_arn"`
	UseBase64 bool   `json:"use_base64"`
}

type AwsIotTopicRuleSpecRepublish struct {
	Topic   string `json:"topic"`
	RoleArn string `json:"role_arn"`
}

type AwsIotTopicRuleSpecS3 struct {
	BucketName string `json:"bucket_name"`
	Key        string `json:"key"`
	RoleArn    string `json:"role_arn"`
}

type AwsIotTopicRuleSpecSns struct {
	MessageFormat string `json:"message_format"`
	TargetArn     string `json:"target_arn"`
	RoleArn       string `json:"role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotTopicRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIotTopicRule `json"items"`
}
