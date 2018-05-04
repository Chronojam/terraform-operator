// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontDistribution struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudfrontDistributionSpec `json"spec"`
}

type AwsCloudfrontDistributionSpec struct {
	RetainOnDelete              bool                                                `json:"retain_on_delete"`
	LastModifiedTime            string                                              `json:"last_modified_time"`
	HostedZoneId                string                                              `json:"hosted_zone_id"`
	Comment                     string                                              `json:"comment"`
	ViewerCertificate           AwsCloudfrontDistributionSpecViewerCertificate      `json:"viewer_certificate"`
	CallerReference             string                                              `json:"caller_reference"`
	ActiveTrustedSigners        map[string]string                                   `json:"active_trusted_signers"`
	DomainName                  string                                              `json:"domain_name"`
	IsIpv6Enabled               bool                                                `json:"is_ipv6_enabled"`
	Tags                        map[string]string                                   `json:"tags"`
	LoggingConfig               AwsCloudfrontDistributionSpecLoggingConfig          `json:"logging_config"`
	Restrictions                AwsCloudfrontDistributionSpecRestrictions           `json:"restrictions"`
	DefaultCacheBehavior        AwsCloudfrontDistributionSpecDefaultCacheBehavior   `json:"default_cache_behavior"`
	CustomErrorResponse         AwsCloudfrontDistributionSpecCustomErrorResponse    `json:"custom_error_response"`
	DefaultRootObject           string                                              `json:"default_root_object"`
	HttpVersion                 string                                              `json:"http_version"`
	Status                      string                                              `json:"status"`
	Arn                         string                                              `json:"arn"`
	OrderedCacheBehavior        []AwsCloudfrontDistributionSpecOrderedCacheBehavior `json:"ordered_cache_behavior"`
	Origin                      AwsCloudfrontDistributionSpecOrigin                 `json:"origin"`
	InProgressValidationBatches int                                                 `json:"in_progress_validation_batches"`
	Etag                        string                                              `json:"etag"`
	Aliases                     string                                              `json:"aliases"`
	CacheBehavior               AwsCloudfrontDistributionSpecCacheBehavior          `json:"cache_behavior"`
	WebAclId                    string                                              `json:"web_acl_id"`
	Enabled                     bool                                                `json:"enabled"`
	PriceClass                  string                                              `json:"price_class"`
}

type AwsCloudfrontDistributionSpecViewerCertificate struct {
	CloudfrontDefaultCertificate bool   `json:"cloudfront_default_certificate"`
	IamCertificateId             string `json:"iam_certificate_id"`
	MinimumProtocolVersion       string `json:"minimum_protocol_version"`
	SslSupportMethod             string `json:"ssl_support_method"`
	AcmCertificateArn            string `json:"acm_certificate_arn"`
}

type AwsCloudfrontDistributionSpecLoggingConfig struct {
	Bucket         string `json:"bucket"`
	IncludeCookies bool   `json:"include_cookies"`
	Prefix         string `json:"prefix"`
}

type AwsCloudfrontDistributionSpecRestrictions struct {
	GeoRestriction AwsCloudfrontDistributionSpecRestrictionsGeoRestriction `json:"geo_restriction"`
}

type AwsCloudfrontDistributionSpecRestrictionsGeoRestriction struct {
	Locations       []string `json:"locations"`
	RestrictionType string   `json:"restriction_type"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehavior struct {
	MinTtl                    int                                                                        `json:"min_ttl"`
	ViewerProtocolPolicy      string                                                                     `json:"viewer_protocol_policy"`
	Compress                  bool                                                                       `json:"compress"`
	ForwardedValues           AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues           `json:"forwarded_values"`
	MaxTtl                    int                                                                        `json:"max_ttl"`
	FieldLevelEncryptionId    string                                                                     `json:"field_level_encryption_id"`
	LambdaFunctionAssociation AwsCloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation `json:"lambda_function_association"`
	SmoothStreaming           bool                                                                       `json:"smooth_streaming"`
	TargetOriginId            string                                                                     `json:"target_origin_id"`
	TrustedSigners            []string                                                                   `json:"trusted_signers"`
	AllowedMethods            []string                                                                   `json:"allowed_methods"`
	CachedMethods             []string                                                                   `json:"cached_methods"`
	DefaultTtl                int                                                                        `json:"default_ttl"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues struct {
	Cookies              AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies `json:"cookies"`
	Headers              []string                                                                `json:"headers"`
	QueryString          bool                                                                    `json:"query_string"`
	QueryStringCacheKeys []string                                                                `json:"query_string_cache_keys"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation struct {
	EventType string `json:"event_type"`
	LambdaArn string `json:"lambda_arn"`
}

type AwsCloudfrontDistributionSpecCustomErrorResponse struct {
	ErrorCachingMinTtl int    `json:"error_caching_min_ttl"`
	ErrorCode          int    `json:"error_code"`
	ResponseCode       int    `json:"response_code"`
	ResponsePagePath   string `json:"response_page_path"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehavior struct {
	LambdaFunctionAssociation AwsCloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation `json:"lambda_function_association"`
	PathPattern               string                                                                     `json:"path_pattern"`
	ForwardedValues           AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues           `json:"forwarded_values"`
	TrustedSigners            []string                                                                   `json:"trusted_signers"`
	ViewerProtocolPolicy      string                                                                     `json:"viewer_protocol_policy"`
	Compress                  bool                                                                       `json:"compress"`
	DefaultTtl                int                                                                        `json:"default_ttl"`
	FieldLevelEncryptionId    string                                                                     `json:"field_level_encryption_id"`
	MaxTtl                    int                                                                        `json:"max_ttl"`
	MinTtl                    int                                                                        `json:"min_ttl"`
	SmoothStreaming           bool                                                                       `json:"smooth_streaming"`
	CachedMethods             string                                                                     `json:"cached_methods"`
	TargetOriginId            string                                                                     `json:"target_origin_id"`
	AllowedMethods            string                                                                     `json:"allowed_methods"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation struct {
	EventType string `json:"event_type"`
	LambdaArn string `json:"lambda_arn"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues struct {
	Cookies              AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies `json:"cookies"`
	Headers              []string                                                                `json:"headers"`
	QueryString          bool                                                                    `json:"query_string"`
	QueryStringCacheKeys []string                                                                `json:"query_string_cache_keys"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type AwsCloudfrontDistributionSpecOrigin struct {
	CustomOriginConfig AwsCloudfrontDistributionSpecOriginCustomOriginConfig `json:"custom_origin_config"`
	DomainName         string                                                `json:"domain_name"`
	CustomHeader       AwsCloudfrontDistributionSpecOriginCustomHeader       `json:"custom_header"`
	OriginId           string                                                `json:"origin_id"`
	OriginPath         string                                                `json:"origin_path"`
	S3OriginConfig     AwsCloudfrontDistributionSpecOriginS3OriginConfig     `json:"s3_origin_config"`
}

type AwsCloudfrontDistributionSpecOriginCustomOriginConfig struct {
	HttpPort               int      `json:"http_port"`
	HttpsPort              int      `json:"https_port"`
	OriginKeepaliveTimeout int      `json:"origin_keepalive_timeout"`
	OriginReadTimeout      int      `json:"origin_read_timeout"`
	OriginProtocolPolicy   string   `json:"origin_protocol_policy"`
	OriginSslProtocols     []string `json:"origin_ssl_protocols"`
}

type AwsCloudfrontDistributionSpecOriginCustomHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsCloudfrontDistributionSpecOriginS3OriginConfig struct {
	OriginAccessIdentity string `json:"origin_access_identity"`
}

type AwsCloudfrontDistributionSpecCacheBehavior struct {
	CachedMethods             []string                                                            `json:"cached_methods"`
	DefaultTtl                int                                                                 `json:"default_ttl"`
	FieldLevelEncryptionId    string                                                              `json:"field_level_encryption_id"`
	SmoothStreaming           bool                                                                `json:"smooth_streaming"`
	TargetOriginId            string                                                              `json:"target_origin_id"`
	Compress                  bool                                                                `json:"compress"`
	LambdaFunctionAssociation AwsCloudfrontDistributionSpecCacheBehaviorLambdaFunctionAssociation `json:"lambda_function_association"`
	TrustedSigners            []string                                                            `json:"trusted_signers"`
	AllowedMethods            []string                                                            `json:"allowed_methods"`
	ForwardedValues           AwsCloudfrontDistributionSpecCacheBehaviorForwardedValues           `json:"forwarded_values"`
	MaxTtl                    int                                                                 `json:"max_ttl"`
	MinTtl                    int                                                                 `json:"min_ttl"`
	PathPattern               string                                                              `json:"path_pattern"`
	ViewerProtocolPolicy      string                                                              `json:"viewer_protocol_policy"`
}

type AwsCloudfrontDistributionSpecCacheBehaviorLambdaFunctionAssociation struct {
	EventType string `json:"event_type"`
	LambdaArn string `json:"lambda_arn"`
}

type AwsCloudfrontDistributionSpecCacheBehaviorForwardedValues struct {
	Headers              []string                                                         `json:"headers"`
	QueryString          bool                                                             `json:"query_string"`
	QueryStringCacheKeys []string                                                         `json:"query_string_cache_keys"`
	Cookies              AwsCloudfrontDistributionSpecCacheBehaviorForwardedValuesCookies `json:"cookies"`
}

type AwsCloudfrontDistributionSpecCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontDistributionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudfrontDistribution `json"items"`
}
