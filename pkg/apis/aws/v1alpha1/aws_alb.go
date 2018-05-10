// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlb struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAlbSpec `json"spec"`
}

type AwsAlbSpec struct {
	SecurityGroups               string                  `json:"security_groups"`
	AccessLogs                   []AwsAlbSpecAccessLogs  `json:"access_logs"`
	EnableHttp2                  bool                    `json:"enable_http2"`
	VpcId                        string                  `json:"vpc_id"`
	ArnSuffix                    string                  `json:"arn_suffix"`
	NamePrefix                   string                  `json:"name_prefix"`
	Subnets                      string                  `json:"subnets"`
	SubnetMapping                AwsAlbSpecSubnetMapping `json:"subnet_mapping"`
	Tags                         map[string]string       `json:"tags"`
	Arn                          string                  `json:"arn"`
	Internal                     bool                    `json:"internal"`
	LoadBalancerType             string                  `json:"load_balancer_type"`
	IdleTimeout                  int                     `json:"idle_timeout"`
	EnableCrossZoneLoadBalancing bool                    `json:"enable_cross_zone_load_balancing"`
	DnsName                      string                  `json:"dns_name"`
	Name                         string                  `json:"name"`
	EnableDeletionProtection     bool                    `json:"enable_deletion_protection"`
	IpAddressType                string                  `json:"ip_address_type"`
	ZoneId                       string                  `json:"zone_id"`
}

type AwsAlbSpecAccessLogs struct {
	Bucket  string `json:"bucket"`
	Prefix  string `json:"prefix"`
	Enabled bool   `json:"enabled"`
}

type AwsAlbSpecSubnetMapping struct {
	SubnetId     string `json:"subnet_id"`
	AllocationId string `json:"allocation_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAlb `json"items"`
}
