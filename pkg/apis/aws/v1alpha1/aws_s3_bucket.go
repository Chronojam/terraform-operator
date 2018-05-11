// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3Bucket struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsS3BucketSpec `json:"spec"`
}

type AwsS3BucketSpec struct {
	Arn                               string                                             `json:"arn"`
	ReplicationConfiguration          []AwsS3BucketSpecReplicationConfiguration          `json:"replication_configuration"`
	BucketPrefix                      string                                             `json:"bucket_prefix"`
	Website                           []AwsS3BucketSpecWebsite                           `json:"website"`
	Logging                           AwsS3BucketSpecLogging                             `json:"logging"`
	ServerSideEncryptionConfiguration []AwsS3BucketSpecServerSideEncryptionConfiguration `json:"server_side_encryption_configuration"`
	BucketDomainName                  string                                             `json:"bucket_domain_name"`
	Policy                            string                                             `json:"policy"`
	WebsiteEndpoint                   string                                             `json:"website_endpoint"`
	WebsiteDomain                     string                                             `json:"website_domain"`
	LifecycleRule                     []AwsS3BucketSpecLifecycleRule                     `json:"lifecycle_rule"`
	ForceDestroy                      bool                                               `json:"force_destroy"`
	RequestPayer                      string                                             `json:"request_payer"`
	Acl                               string                                             `json:"acl"`
	CorsRule                          []AwsS3BucketSpecCorsRule                          `json:"cors_rule"`
	HostedZoneId                      string                                             `json:"hosted_zone_id"`
	Region                            string                                             `json:"region"`
	Versioning                        []AwsS3BucketSpecVersioning                        `json:"versioning"`
	AccelerationStatus                string                                             `json:"acceleration_status"`
	Tags                              map[string]string                                  `json:"tags"`
	Bucket                            string                                             `json:"bucket"`
}

type AwsS3BucketSpecReplicationConfiguration struct {
	Role  string                                       `json:"role"`
	Rules AwsS3BucketSpecReplicationConfigurationRules `json:"rules"`
}

type AwsS3BucketSpecReplicationConfigurationRules struct {
	SourceSelectionCriteria AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria `json:"source_selection_criteria"`
	Prefix                  string                                                              `json:"prefix"`
	Status                  string                                                              `json:"status"`
	Id                      string                                                              `json:"id"`
	Destination             AwsS3BucketSpecReplicationConfigurationRulesDestination             `json:"destination"`
}

type AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria struct {
	SseKmsEncryptedObjects AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects `json:"sse_kms_encrypted_objects"`
}

type AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	Enabled bool `json:"enabled"`
}

type AwsS3BucketSpecReplicationConfigurationRulesDestination struct {
	Bucket          string `json:"bucket"`
	StorageClass    string `json:"storage_class"`
	ReplicaKmsKeyId string `json:"replica_kms_key_id"`
}

type AwsS3BucketSpecWebsite struct {
	RedirectAllRequestsTo string `json:"redirect_all_requests_to"`
	RoutingRules          string `json:"routing_rules"`
	IndexDocument         string `json:"index_document"`
	ErrorDocument         string `json:"error_document"`
}

type AwsS3BucketSpecLogging struct {
	TargetBucket string `json:"target_bucket"`
	TargetPrefix string `json:"target_prefix"`
}

type AwsS3BucketSpecServerSideEncryptionConfiguration struct {
	Rule []AwsS3BucketSpecServerSideEncryptionConfigurationRule `json:"rule"`
}

type AwsS3BucketSpecServerSideEncryptionConfigurationRule struct {
	ApplyServerSideEncryptionByDefault []AwsS3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault `json:"apply_server_side_encryption_by_default"`
}

type AwsS3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault struct {
	KmsMasterKeyId string `json:"kms_master_key_id"`
	SseAlgorithm   string `json:"sse_algorithm"`
}

type AwsS3BucketSpecLifecycleRule struct {
	Expiration                         AwsS3BucketSpecLifecycleRuleExpiration                  `json:"expiration"`
	NoncurrentVersionTransition        AwsS3BucketSpecLifecycleRuleNoncurrentVersionTransition `json:"noncurrent_version_transition"`
	Enabled                            bool                                                    `json:"enabled"`
	AbortIncompleteMultipartUploadDays int                                                     `json:"abort_incomplete_multipart_upload_days"`
	Tags                               map[string]string                                       `json:"tags"`
	NoncurrentVersionExpiration        AwsS3BucketSpecLifecycleRuleNoncurrentVersionExpiration `json:"noncurrent_version_expiration"`
	Transition                         AwsS3BucketSpecLifecycleRuleTransition                  `json:"transition"`
	Id                                 string                                                  `json:"id"`
	Prefix                             string                                                  `json:"prefix"`
}

type AwsS3BucketSpecLifecycleRuleExpiration struct {
	ExpiredObjectDeleteMarker bool   `json:"expired_object_delete_marker"`
	Date                      string `json:"date"`
	Days                      int    `json:"days"`
}

type AwsS3BucketSpecLifecycleRuleNoncurrentVersionTransition struct {
	Days         int    `json:"days"`
	StorageClass string `json:"storage_class"`
}

type AwsS3BucketSpecLifecycleRuleNoncurrentVersionExpiration struct {
	Days int `json:"days"`
}

type AwsS3BucketSpecLifecycleRuleTransition struct {
	Days         int    `json:"days"`
	StorageClass string `json:"storage_class"`
	Date         string `json:"date"`
}

type AwsS3BucketSpecCorsRule struct {
	MaxAgeSeconds  int      `json:"max_age_seconds"`
	AllowedHeaders []string `json:"allowed_headers"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedOrigins []string `json:"allowed_origins"`
	ExposeHeaders  []string `json:"expose_headers"`
}

type AwsS3BucketSpecVersioning struct {
	Enabled   bool `json:"enabled"`
	MfaDelete bool `json:"mfa_delete"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsS3Bucket `json:"items"`
}
