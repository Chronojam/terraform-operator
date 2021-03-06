// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListener struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsLbListenerSpec `json:"spec"`
}

type AwsLbListenerSpec struct {
	DefaultAction   []AwsLbListenerSpecDefaultAction `json:"default_action"`
	Arn             string                           `json:"arn"`
	LoadBalancerArn string                           `json:"load_balancer_arn"`
	Port            int                              `json:"port"`
	Protocol        string                           `json:"protocol"`
	SslPolicy       string                           `json:"ssl_policy"`
	CertificateArn  string                           `json:"certificate_arn"`
}

type AwsLbListenerSpecDefaultAction struct {
	TargetGroupArn string `json:"target_group_arn"`
	Type           string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsLbListener `json:"items"`
}
