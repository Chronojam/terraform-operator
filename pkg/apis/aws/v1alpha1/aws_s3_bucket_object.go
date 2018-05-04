// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketObject struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsS3BucketObjectSpec `json"spec"`
}

type AwsS3BucketObjectSpec struct {
	Acl                  string            `json:"acl"`
	CacheControl         string            `json:"cache_control"`
	ContentBase64        string            `json:"content_base64"`
	Bucket               string            `json:"bucket"`
	Key                  string            `json:"key"`
	KmsKeyId             string            `json:"kms_key_id"`
	Etag                 string            `json:"etag"`
	WebsiteRedirect      string            `json:"website_redirect"`
	Content              string            `json:"content"`
	ServerSideEncryption string            `json:"server_side_encryption"`
	VersionId            string            `json:"version_id"`
	Tags                 map[string]string `json:"tags"`
	ContentDisposition   string            `json:"content_disposition"`
	ContentEncoding      string            `json:"content_encoding"`
	ContentLanguage      string            `json:"content_language"`
	ContentType          string            `json:"content_type"`
	Source               string            `json:"source"`
	StorageClass         string            `json:"storage_class"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketObjectList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsS3BucketObject `json"items"`
}
