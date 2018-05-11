// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbCookieStickinessPolicy struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsLbCookieStickinessPolicySpec `json:"spec"`
}

type AwsLbCookieStickinessPolicySpec struct {
	CookieExpirationPeriod int    `json:"cookie_expiration_period"`
	Name                   string `json:"name"`
	LoadBalancer           string `json:"load_balancer"`
	LbPort                 int    `json:"lb_port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbCookieStickinessPolicyList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsLbCookieStickinessPolicy `json:"items"`
}
