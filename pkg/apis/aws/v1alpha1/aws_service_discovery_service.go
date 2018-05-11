// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryService struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsServiceDiscoveryServiceSpec `json"spec"`
}

type AwsServiceDiscoveryServiceSpec struct {
	Name                    string                                                  `json:"name"`
	Description             string                                                  `json:"description"`
	DnsConfig               []AwsServiceDiscoveryServiceSpecDnsConfig               `json:"dns_config"`
	HealthCheckConfig       []AwsServiceDiscoveryServiceSpecHealthCheckConfig       `json:"health_check_config"`
	HealthCheckCustomConfig []AwsServiceDiscoveryServiceSpecHealthCheckCustomConfig `json:"health_check_custom_config"`
	Arn                     string                                                  `json:"arn"`
}

type AwsServiceDiscoveryServiceSpecDnsConfig struct {
	NamespaceId   string                                              `json:"namespace_id"`
	DnsRecords    []AwsServiceDiscoveryServiceSpecDnsConfigDnsRecords `json:"dns_records"`
	RoutingPolicy string                                              `json:"routing_policy"`
}

type AwsServiceDiscoveryServiceSpecDnsConfigDnsRecords struct {
	Ttl  int    `json:"ttl"`
	Type string `json:"type"`
}

type AwsServiceDiscoveryServiceSpecHealthCheckConfig struct {
	ResourcePath     string `json:"resource_path"`
	Type             string `json:"type"`
	FailureThreshold int    `json:"failure_threshold"`
}

type AwsServiceDiscoveryServiceSpecHealthCheckCustomConfig struct {
	FailureThreshold int `json:"failure_threshold"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryServiceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsServiceDiscoveryService `json"items"`
}
