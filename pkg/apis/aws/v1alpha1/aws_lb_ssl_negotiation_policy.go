// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbSslNegotiationPolicy struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsLbSslNegotiationPolicySpec `json:"spec"`
}

type AwsLbSslNegotiationPolicySpec struct {
	Name         string                                 `json:"name"`
	LoadBalancer string                                 `json:"load_balancer"`
	LbPort       int                                    `json:"lb_port"`
	Attribute    AwsLbSslNegotiationPolicySpecAttribute `json:"attribute"`
}

type AwsLbSslNegotiationPolicySpecAttribute struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbSslNegotiationPolicyList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsLbSslNegotiationPolicy `json:"items"`
}
