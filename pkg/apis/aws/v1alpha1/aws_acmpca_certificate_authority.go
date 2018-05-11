// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAcmpcaCertificateAuthority struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAcmpcaCertificateAuthoritySpec `json"spec"`
}

type AwsAcmpcaCertificateAuthoritySpec struct {
	RevocationConfiguration           []AwsAcmpcaCertificateAuthoritySpecRevocationConfiguration           `json:"revocation_configuration"`
	Arn                               string                                                               `json:"arn"`
	Certificate                       string                                                               `json:"certificate"`
	CertificateAuthorityConfiguration []AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration `json:"certificate_authority_configuration"`
	CertificateChain                  string                                                               `json:"certificate_chain"`
	CertificateSigningRequest         string                                                               `json:"certificate_signing_request"`
	NotAfter                          string                                                               `json:"not_after"`
	NotBefore                         string                                                               `json:"not_before"`
	Status                            string                                                               `json:"status"`
	Tags                              map[string]string                                                    `json:"tags"`
	Enabled                           bool                                                                 `json:"enabled"`
	Serial                            string                                                               `json:"serial"`
	Type                              string                                                               `json:"type"`
}

type AwsAcmpcaCertificateAuthoritySpecRevocationConfiguration struct {
	CrlConfiguration []AwsAcmpcaCertificateAuthoritySpecRevocationConfigurationCrlConfiguration `json:"crl_configuration"`
}

type AwsAcmpcaCertificateAuthoritySpecRevocationConfigurationCrlConfiguration struct {
	CustomCname      string `json:"custom_cname"`
	Enabled          bool   `json:"enabled"`
	ExpirationInDays int    `json:"expiration_in_days"`
	S3BucketName     string `json:"s3_bucket_name"`
}

type AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration struct {
	KeyAlgorithm     string                                                                      `json:"key_algorithm"`
	SigningAlgorithm string                                                                      `json:"signing_algorithm"`
	Subject          []AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfigurationSubject `json:"subject"`
}

type AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfigurationSubject struct {
	Surname                    string `json:"surname"`
	CommonName                 string `json:"common_name"`
	Organization               string `json:"organization"`
	OrganizationalUnit         string `json:"organizational_unit"`
	State                      string `json:"state"`
	Initials                   string `json:"initials"`
	Locality                   string `json:"locality"`
	Pseudonym                  string `json:"pseudonym"`
	Title                      string `json:"title"`
	Country                    string `json:"country"`
	DistinguishedNameQualifier string `json:"distinguished_name_qualifier"`
	GenerationQualifier        string `json:"generation_qualifier"`
	GivenName                  string `json:"given_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAcmpcaCertificateAuthorityList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAcmpcaCertificateAuthority `json"items"`
}
