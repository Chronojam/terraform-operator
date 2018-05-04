// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLb struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLbSpec `json"spec"`
}

type AwsLbSpec struct {
	Tags                         map[string]string      `json:"tags"`
	ArnSuffix                    string                 `json:"arn_suffix"`
	NamePrefix                   string                 `json:"name_prefix"`
	SubnetMapping                AwsLbSpecSubnetMapping `json:"subnet_mapping"`
	AccessLogs                   []AwsLbSpecAccessLogs  `json:"access_logs"`
	EnableDeletionProtection     bool                   `json:"enable_deletion_protection"`
	IdleTimeout                  int                    `json:"idle_timeout"`
	EnableHttp2                  bool                   `json:"enable_http2"`
	ZoneId                       string                 `json:"zone_id"`
	Name                         string                 `json:"name"`
	Internal                     bool                   `json:"internal"`
	LoadBalancerType             string                 `json:"load_balancer_type"`
	Subnets                      string                 `json:"subnets"`
	DnsName                      string                 `json:"dns_name"`
	Arn                          string                 `json:"arn"`
	SecurityGroups               string                 `json:"security_groups"`
	EnableCrossZoneLoadBalancing bool                   `json:"enable_cross_zone_load_balancing"`
	IpAddressType                string                 `json:"ip_address_type"`
	VpcId                        string                 `json:"vpc_id"`
}

type AwsLbSpecSubnetMapping struct {
	SubnetId     string `json:"subnet_id"`
	AllocationId string `json:"allocation_id"`
}

type AwsLbSpecAccessLogs struct {
	Bucket  string `json:"bucket"`
	Prefix  string `json:"prefix"`
	Enabled bool   `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLb `json"items"`
}
