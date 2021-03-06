// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroup struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsAlbTargetGroupSpec `json:"spec"`
}

type AwsAlbTargetGroupSpec struct {
	Arn                 string                             `json:"arn"`
	DeregistrationDelay int                                `json:"deregistration_delay"`
	ProxyProtocolV2     bool                               `json:"proxy_protocol_v2"`
	TargetType          string                             `json:"target_type"`
	Stickiness          []AwsAlbTargetGroupSpecStickiness  `json:"stickiness"`
	VpcId               string                             `json:"vpc_id"`
	HealthCheck         []AwsAlbTargetGroupSpecHealthCheck `json:"health_check"`
	Tags                map[string]string                  `json:"tags"`
	ArnSuffix           string                             `json:"arn_suffix"`
	Name                string                             `json:"name"`
	NamePrefix          string                             `json:"name_prefix"`
	Port                int                                `json:"port"`
	Protocol            string                             `json:"protocol"`
}

type AwsAlbTargetGroupSpecStickiness struct {
	Enabled        bool   `json:"enabled"`
	Type           string `json:"type"`
	CookieDuration int    `json:"cookie_duration"`
}

type AwsAlbTargetGroupSpecHealthCheck struct {
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Interval           int    `json:"interval"`
	Path               string `json:"path"`
	Port               string `json:"port"`
	Protocol           string `json:"protocol"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroupList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsAlbTargetGroup `json:"items"`
}
