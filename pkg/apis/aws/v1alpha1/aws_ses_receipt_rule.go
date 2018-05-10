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
	LambdaAction    AwsSesReceiptRuleSpecLambdaAction    `json:"lambda_action"`
	SnsAction       AwsSesReceiptRuleSpecSnsAction       `json:"sns_action"`
	RuleSetName     string                               `json:"rule_set_name"`
	ScanEnabled     bool                                 `json:"scan_enabled"`
	BounceAction    AwsSesReceiptRuleSpecBounceAction    `json:"bounce_action"`
	StopAction      AwsSesReceiptRuleSpecStopAction      `json:"stop_action"`
	Name            string                               `json:"name"`
	After           string                               `json:"after"`
	Enabled         bool                                 `json:"enabled"`
	Recipients      string                               `json:"recipients"`
	S3Action        AwsSesReceiptRuleSpecS3Action        `json:"s3_action"`
	WorkmailAction  AwsSesReceiptRuleSpecWorkmailAction  `json:"workmail_action"`
	TlsPolicy       string                               `json:"tls_policy"`
	AddHeaderAction AwsSesReceiptRuleSpecAddHeaderAction `json:"add_header_action"`
}

type AwsSesReceiptRuleSpecLambdaAction struct {
	FunctionArn    string `json:"function_arn"`
	InvocationType string `json:"invocation_type"`
	TopicArn       string `json:"topic_arn"`
	Position       int    `json:"position"`
}

type AwsSesReceiptRuleSpecSnsAction struct {
	TopicArn string `json:"topic_arn"`
	Position int    `json:"position"`
}

type AwsSesReceiptRuleSpecBounceAction struct {
	Sender        string `json:"sender"`
	SmtpReplyCode string `json:"smtp_reply_code"`
	StatusCode    string `json:"status_code"`
	TopicArn      string `json:"topic_arn"`
	Position      int    `json:"position"`
	Message       string `json:"message"`
}

type AwsSesReceiptRuleSpecStopAction struct {
	Scope    string `json:"scope"`
	TopicArn string `json:"topic_arn"`
	Position int    `json:"position"`
}

type AwsSesReceiptRuleSpecS3Action struct {
	BucketName      string `json:"bucket_name"`
	KmsKeyArn       string `json:"kms_key_arn"`
	ObjectKeyPrefix string `json:"object_key_prefix"`
	TopicArn        string `json:"topic_arn"`
	Position        int    `json:"position"`
}

type AwsSesReceiptRuleSpecWorkmailAction struct {
	Position        int    `json:"position"`
	OrganizationArn string `json:"organization_arn"`
	TopicArn        string `json:"topic_arn"`
}

type AwsSesReceiptRuleSpecAddHeaderAction struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
	Position    int    `json:"position"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesReceiptRule `json"items"`
}
