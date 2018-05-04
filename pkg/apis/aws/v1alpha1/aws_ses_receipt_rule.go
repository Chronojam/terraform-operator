// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesReceiptRuleSpec `json"spec"`
}

type AwsSesReceiptRuleSpec struct {
	S3Action        AwsSesReceiptRuleSpecS3Action        `json:"s3_action"`
	SnsAction       AwsSesReceiptRuleSpecSnsAction       `json:"sns_action"`
	Enabled         bool                                 `json:"enabled"`
	AddHeaderAction AwsSesReceiptRuleSpecAddHeaderAction `json:"add_header_action"`
	BounceAction    AwsSesReceiptRuleSpecBounceAction    `json:"bounce_action"`
	Recipients      string                               `json:"recipients"`
	After           string                               `json:"after"`
	ScanEnabled     bool                                 `json:"scan_enabled"`
	WorkmailAction  AwsSesReceiptRuleSpecWorkmailAction  `json:"workmail_action"`
	RuleSetName     string                               `json:"rule_set_name"`
	TlsPolicy       string                               `json:"tls_policy"`
	LambdaAction    AwsSesReceiptRuleSpecLambdaAction    `json:"lambda_action"`
	StopAction      AwsSesReceiptRuleSpecStopAction      `json:"stop_action"`
	Name            string                               `json:"name"`
}

type AwsSesReceiptRuleSpecS3Action struct {
	BucketName      string `json:"bucket_name"`
	KmsKeyArn       string `json:"kms_key_arn"`
	ObjectKeyPrefix string `json:"object_key_prefix"`
	TopicArn        string `json:"topic_arn"`
	Position        int    `json:"position"`
}

type AwsSesReceiptRuleSpecSnsAction struct {
	Position int    `json:"position"`
	TopicArn string `json:"topic_arn"`
}

type AwsSesReceiptRuleSpecAddHeaderAction struct {
	Position    int    `json:"position"`
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
}

type AwsSesReceiptRuleSpecBounceAction struct {
	Message       string `json:"message"`
	Sender        string `json:"sender"`
	SmtpReplyCode string `json:"smtp_reply_code"`
	StatusCode    string `json:"status_code"`
	TopicArn      string `json:"topic_arn"`
	Position      int    `json:"position"`
}

type AwsSesReceiptRuleSpecWorkmailAction struct {
	OrganizationArn string `json:"organization_arn"`
	TopicArn        string `json:"topic_arn"`
	Position        int    `json:"position"`
}

type AwsSesReceiptRuleSpecLambdaAction struct {
	Position       int    `json:"position"`
	FunctionArn    string `json:"function_arn"`
	InvocationType string `json:"invocation_type"`
	TopicArn       string `json:"topic_arn"`
}

type AwsSesReceiptRuleSpecStopAction struct {
	Scope    string `json:"scope"`
	TopicArn string `json:"topic_arn"`
	Position int    `json:"position"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesReceiptRule `json"items"`
}
