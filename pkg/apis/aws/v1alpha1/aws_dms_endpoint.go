// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsEndpoint struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDmsEndpointSpec `json"spec"`
}

type AwsDmsEndpointSpec struct {
	DatabaseName              string                              `json:"database_name"`
	EngineName                string                              `json:"engine_name"`
	SslMode                   string                              `json:"ssl_mode"`
	EndpointArn               string                              `json:"endpoint_arn"`
	ServiceAccessRole         string                              `json:"service_access_role"`
	EndpointType              string                              `json:"endpoint_type"`
	Port                      int                                 `json:"port"`
	ServerName                string                              `json:"server_name"`
	Tags                      map[string]string                   `json:"tags"`
	S3Settings                []AwsDmsEndpointSpecS3Settings      `json:"s3_settings"`
	CertificateArn            string                              `json:"certificate_arn"`
	KmsKeyArn                 string                              `json:"kms_key_arn"`
	Username                  string                              `json:"username"`
	MongodbSettings           []AwsDmsEndpointSpecMongodbSettings `json:"mongodb_settings"`
	EndpointId                string                              `json:"endpoint_id"`
	ExtraConnectionAttributes string                              `json:"extra_connection_attributes"`
	Password                  string                              `json:"password"`
}

type AwsDmsEndpointSpecS3Settings struct {
	CompressionType         string `json:"compression_type"`
	ServiceAccessRoleArn    string `json:"service_access_role_arn"`
	ExternalTableDefinition string `json:"external_table_definition"`
	CsvRowDelimiter         string `json:"csv_row_delimiter"`
	CsvDelimiter            string `json:"csv_delimiter"`
	BucketFolder            string `json:"bucket_folder"`
	BucketName              string `json:"bucket_name"`
}

type AwsDmsEndpointSpecMongodbSettings struct {
	AuthType          string `json:"auth_type"`
	AuthMechanism     string `json:"auth_mechanism"`
	NestingLevel      string `json:"nesting_level"`
	ExtractDocId      string `json:"extract_doc_id"`
	DocsToInvestigate string `json:"docs_to_investigate"`
	AuthSource        string `json:"auth_source"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsEndpointList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDmsEndpoint `json"items"`
}
