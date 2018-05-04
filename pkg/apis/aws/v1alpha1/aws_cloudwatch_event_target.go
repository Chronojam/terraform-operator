// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventTarget struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchEventTargetSpec `json"spec"`
}

type AwsCloudwatchEventTargetSpec struct {
	Input             string                                          `json:"input"`
	RunCommandTargets []AwsCloudwatchEventTargetSpecRunCommandTargets `json:"run_command_targets"`
	EcsTarget         []AwsCloudwatchEventTargetSpecEcsTarget         `json:"ecs_target"`
	BatchTarget       []AwsCloudwatchEventTargetSpecBatchTarget       `json:"batch_target"`
	SqsTarget         []AwsCloudwatchEventTargetSpecSqsTarget         `json:"sqs_target"`
	Rule              string                                          `json:"rule"`
	TargetId          string                                          `json:"target_id"`
	Arn               string                                          `json:"arn"`
	InputTransformer  []AwsCloudwatchEventTargetSpecInputTransformer  `json:"input_transformer"`
	InputPath         string                                          `json:"input_path"`
	RoleArn           string                                          `json:"role_arn"`
	KinesisTarget     []AwsCloudwatchEventTargetSpecKinesisTarget     `json:"kinesis_target"`
}

type AwsCloudwatchEventTargetSpecRunCommandTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type AwsCloudwatchEventTargetSpecEcsTarget struct {
	TaskCount         int    `json:"task_count"`
	TaskDefinitionArn string `json:"task_definition_arn"`
}

type AwsCloudwatchEventTargetSpecBatchTarget struct {
	JobName       string `json:"job_name"`
	ArraySize     int    `json:"array_size"`
	JobAttempts   int    `json:"job_attempts"`
	JobDefinition string `json:"job_definition"`
}

type AwsCloudwatchEventTargetSpecSqsTarget struct {
	MessageGroupId string `json:"message_group_id"`
}

type AwsCloudwatchEventTargetSpecInputTransformer struct {
	InputPaths    map[string]string `json:"input_paths"`
	InputTemplate string            `json:"input_template"`
}

type AwsCloudwatchEventTargetSpecKinesisTarget struct {
	PartitionKeyPath string `json:"partition_key_path"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventTargetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchEventTarget `json"items"`
}
