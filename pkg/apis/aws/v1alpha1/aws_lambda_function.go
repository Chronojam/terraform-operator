// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaFunction struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsLambdaFunctionSpec `json:"spec"`
}

type AwsLambdaFunctionSpec struct {
	FunctionName                 string                                  `json:"function_name"`
	Handler                      string                                  `json:"handler"`
	MemorySize                   int                                     `json:"memory_size"`
	Runtime                      string                                  `json:"runtime"`
	S3Key                        string                                  `json:"s3_key"`
	DeadLetterConfig             []AwsLambdaFunctionSpecDeadLetterConfig `json:"dead_letter_config"`
	VpcConfig                    []AwsLambdaFunctionSpecVpcConfig        `json:"vpc_config"`
	Arn                          string                                  `json:"arn"`
	QualifiedArn                 string                                  `json:"qualified_arn"`
	SourceCodeHash               string                                  `json:"source_code_hash"`
	SourceCodeSize               int                                     `json:"source_code_size"`
	Environment                  []AwsLambdaFunctionSpecEnvironment      `json:"environment"`
	Timeout                      int                                     `json:"timeout"`
	Publish                      bool                                    `json:"publish"`
	Tags                         map[string]string                       `json:"tags"`
	TracingConfig                []AwsLambdaFunctionSpecTracingConfig    `json:"tracing_config"`
	KmsKeyArn                    string                                  `json:"kms_key_arn"`
	Description                  string                                  `json:"description"`
	ReservedConcurrentExecutions int                                     `json:"reserved_concurrent_executions"`
	InvokeArn                    string                                  `json:"invoke_arn"`
	Filename                     string                                  `json:"filename"`
	S3Bucket                     string                                  `json:"s3_bucket"`
	Version                      string                                  `json:"version"`
	LastModified                 string                                  `json:"last_modified"`
	S3ObjectVersion              string                                  `json:"s3_object_version"`
	Role                         string                                  `json:"role"`
}

type AwsLambdaFunctionSpecDeadLetterConfig struct {
	TargetArn string `json:"target_arn"`
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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaFunctionList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsLambdaFunction `json:"items"`
}
