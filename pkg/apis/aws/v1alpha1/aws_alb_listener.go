// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbListener struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsAlbListenerSpec `json:"spec"`
}

type AwsAlbListenerSpec struct {
	DefaultAction   []AwsAlbListenerSpecDefaultAction `json:"default_action"`
	Arn             string                            `json:"arn"`
	LoadBalancerArn string                            `json:"load_balancer_arn"`
	Port            int                               `json:"port"`
	Protocol        string                            `json:"protocol"`
	SslPolicy       string                            `json:"ssl_policy"`
	CertificateArn  string                            `json:"certificate_arn"`
}

type AwsAlbListenerSpecDefaultAction struct {
	TargetGroupArn string `json:"target_group_arn"`
	Type           string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbListenerList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsAlbListener `json:"items"`
}
