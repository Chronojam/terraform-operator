// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontOriginAccessIdentity struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsCloudfrontOriginAccessIdentitySpec `json:"spec"`
}

type AwsCloudfrontOriginAccessIdentitySpec struct {
	CloudfrontAccessIdentityPath string `json:"cloudfront_access_identity_path"`
	Etag                         string `json:"etag"`
	IamArn                       string `json:"iam_arn"`
	S3CanonicalUserId            string `json:"s3_canonical_user_id"`
	Comment                      string `json:"comment"`
	CallerReference              string `json:"caller_reference"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontOriginAccessIdentityList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsCloudfrontOriginAccessIdentity `json:"items"`
}
