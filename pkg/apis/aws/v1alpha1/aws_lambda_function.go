// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaFunction struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLambdaFunctionSpec `json"spec"`
}

type AwsLambdaFunctionSpec struct {
	FunctionName                 string                                  `json:"function_name"`
	MemorySize                   int                                     `json:"memory_size"`
	Role                         string                                  `json:"role"`
	VpcConfig                    []AwsLambdaFunctionSpecVpcConfig        `json:"vpc_config"`
	Arn                          string                                  `json:"arn"`
	LastModified                 string                                  `json:"last_modified"`
	SourceCodeSize               int                                     `json:"source_code_size"`
	S3Bucket                     string                                  `json:"s3_bucket"`
	S3ObjectVersion              string                                  `json:"s3_object_version"`
	QualifiedArn                 string                                  `json:"qualified_arn"`
	SourceCodeHash               string                                  `json:"source_code_hash"`
	Environment                  []AwsLambdaFunctionSpecEnvironment      `json:"environment"`
	KmsKeyArn                    string                                  `json:"kms_key_arn"`
	S3Key                        string                                  `json:"s3_key"`
	Runtime                      string                                  `json:"runtime"`
	Timeout                      int                                     `json:"timeout"`
	InvokeArn                    string                                  `json:"invoke_arn"`
	TracingConfig                []AwsLambdaFunctionSpecTracingConfig    `json:"tracing_config"`
	Tags                         map[string]string                       `json:"tags"`
	Handler                      string                                  `json:"handler"`
	Description                  string                                  `json:"description"`
	DeadLetterConfig             []AwsLambdaFunctionSpecDeadLetterConfig `json:"dead_letter_config"`
	ReservedConcurrentExecutions int                                     `json:"reserved_concurrent_executions"`
	Publish                      bool                                    `json:"publish"`
	Version                      string                                  `json:"version"`
	Filename                     string                                  `json:"filename"`
}

type AwsLambdaFunctionSpecVpcConfig struct {
	SubnetIds        string `json:"subnet_ids"`
	SecurityGroupIds string `json:"security_group_ids"`
	VpcId            string `json:"vpc_id"`
}

type AwsLambdaFunctionSpecEnvironment struct {
	Variables map[string]string `json:"variables"`
}

type AwsLambdaFunctionSpecTracingConfig struct {
	Mode string `json:"mode"`
}

type AwsLambdaFunctionSpecDeadLetterConfig struct {
	TargetArn string `json:"target_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaFunctionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLambdaFunction `json"items"`
}
