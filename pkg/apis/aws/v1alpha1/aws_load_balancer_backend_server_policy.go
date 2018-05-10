// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerBackendServerPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLoadBalancerBackendServerPolicySpec `json"spec"`
}

type AwsLoadBalancerBackendServerPolicySpec struct {
	InstancePort     int    `json:"instance_port"`
	LoadBalancerName string `json:"load_balancer_name"`
	PolicyNames      string `json:"policy_names"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerBackendServerPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLoadBalancerBackendServerPolicy `json"items"`
}
