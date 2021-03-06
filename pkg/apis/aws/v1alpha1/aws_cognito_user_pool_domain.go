// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolDomain struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsCognitoUserPoolDomainSpec `json:"spec"`
}

type AwsCognitoUserPoolDomainSpec struct {
	S3Bucket                  string `json:"s3_bucket"`
	Version                   string `json:"version"`
	Domain                    string `json:"domain"`
	UserPoolId                string `json:"user_pool_id"`
	AwsAccountId              string `json:"aws_account_id"`
	CloudfrontDistributionArn string `json:"cloudfront_distribution_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolDomainList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsCognitoUserPoolDomain `json:"items"`
}
