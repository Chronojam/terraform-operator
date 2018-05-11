// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlb struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsAlbSpec `json:"spec"`
}

type AwsAlbSpec struct {
	EnableDeletionProtection     bool                    `json:"enable_deletion_protection"`
	IdleTimeout                  int                     `json:"idle_timeout"`
	IpAddressType                string                  `json:"ip_address_type"`
	Arn                          string                  `json:"arn"`
	NamePrefix                   string                  `json:"name_prefix"`
	SubnetMapping                AwsAlbSpecSubnetMapping `json:"subnet_mapping"`
	AccessLogs                   []AwsAlbSpecAccessLogs  `json:"access_logs"`
	EnableCrossZoneLoadBalancing bool                    `json:"enable_cross_zone_load_balancing"`
	ZoneId                       string                  `json:"zone_id"`
	DnsName                      string                  `json:"dns_name"`
	ArnSuffix                    string                  `json:"arn_suffix"`
	Subnets                      string                  `json:"subnets"`
	EnableHttp2                  bool                    `json:"enable_http2"`
	VpcId                        string                  `json:"vpc_id"`
	Tags                         map[string]string       `json:"tags"`
	Name                         string                  `json:"name"`
	Internal                     bool                    `json:"internal"`
	LoadBalancerType             string                  `json:"load_balancer_type"`
	SecurityGroups               string                  `json:"security_groups"`
}

type AwsAlbSpecSubnetMapping struct {
	SubnetId     string `json:"subnet_id"`
	AllocationId string `json:"allocation_id"`
}

type AwsAlbSpecAccessLogs struct {
	Prefix  string `json:"prefix"`
	Enabled bool   `json:"enabled"`
	Bucket  string `json:"bucket"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsAlb `json:"items"`
}
