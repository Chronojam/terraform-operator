// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointService struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcEndpointServiceSpec `json"spec"`
}

type AwsVpcEndpointServiceSpec struct {
	BaseEndpointDnsNames    string `json:"base_endpoint_dns_names"`
	AcceptanceRequired      bool   `json:"acceptance_required"`
	AllowedPrincipals       string `json:"allowed_principals"`
	State                   string `json:"state"`
	ServiceName             string `json:"service_name"`
	NetworkLoadBalancerArns string `json:"network_load_balancer_arns"`
	ServiceType             string `json:"service_type"`
	AvailabilityZones       string `json:"availability_zones"`
	PrivateDnsName          string `json:"private_dns_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointServiceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcEndpointService `json"items"`
}
