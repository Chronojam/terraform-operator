// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53HealthCheck struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRoute53HealthCheckSpec `json"spec"`
}

type AwsRoute53HealthCheckSpec struct {
	Fqdn                         string            `json:"fqdn"`
	CloudwatchAlarmRegion        string            `json:"cloudwatch_alarm_region"`
	CloudwatchAlarmName          string            `json:"cloudwatch_alarm_name"`
	RequestInterval              int               `json:"request_interval"`
	SearchString                 string            `json:"search_string"`
	ChildHealthThreshold         int               `json:"child_health_threshold"`
	Tags                         map[string]string `json:"tags"`
	IpAddress                    string            `json:"ip_address"`
	ResourcePath                 string            `json:"resource_path"`
	InsufficientDataHealthStatus string            `json:"insufficient_data_health_status"`
	InvertHealthcheck            bool              `json:"invert_healthcheck"`
	MeasureLatency               bool              `json:"measure_latency"`
	ChildHealthchecks            string            `json:"child_healthchecks"`
	ReferenceName                string            `json:"reference_name"`
	EnableSni                    bool              `json:"enable_sni"`
	Type                         string            `json:"type"`
	FailureThreshold             int               `json:"failure_threshold"`
	Port                         int               `json:"port"`
	Regions                      string            `json:"regions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53HealthCheckList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRoute53HealthCheck `json"items"`
}
