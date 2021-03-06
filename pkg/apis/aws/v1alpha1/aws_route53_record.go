// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53Record struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsRoute53RecordSpec `json:"spec"`
}

type AwsRoute53RecordSpec struct {
	HealthCheckId                 string                                         `json:"health_check_id"`
	AllowOverwrite                bool                                           `json:"allow_overwrite"`
	GeolocationRoutingPolicy      []AwsRoute53RecordSpecGeolocationRoutingPolicy `json:"geolocation_routing_policy"`
	WeightedRoutingPolicy         []AwsRoute53RecordSpecWeightedRoutingPolicy    `json:"weighted_routing_policy"`
	LatencyRoutingPolicy          []AwsRoute53RecordSpecLatencyRoutingPolicy     `json:"latency_routing_policy"`
	ZoneId                        string                                         `json:"zone_id"`
	Weight                        int                                            `json:"weight"`
	Failover                      string                                         `json:"failover"`
	FailoverRoutingPolicy         []AwsRoute53RecordSpecFailoverRoutingPolicy    `json:"failover_routing_policy"`
	Name                          string                                         `json:"name"`
	Fqdn                          string                                         `json:"fqdn"`
	SetIdentifier                 string                                         `json:"set_identifier"`
	Alias                         AwsRoute53RecordSpecAlias                      `json:"alias"`
	MultivalueAnswerRoutingPolicy bool                                           `json:"multivalue_answer_routing_policy"`
	Records                       string                                         `json:"records"`
	Type                          string                                         `json:"type"`
	Ttl                           int                                            `json:"ttl"`
}

type AwsRoute53RecordSpecGeolocationRoutingPolicy struct {
	Continent   string `json:"continent"`
	Country     string `json:"country"`
	Subdivision string `json:"subdivision"`
}

type AwsRoute53RecordSpecWeightedRoutingPolicy struct {
	Weight int `json:"weight"`
}

type AwsRoute53RecordSpecLatencyRoutingPolicy struct {
	Region string `json:"region"`
}

type AwsRoute53RecordSpecFailoverRoutingPolicy struct {
	Type string `json:"type"`
}

type AwsRoute53RecordSpecAlias struct {
	ZoneId               string `json:"zone_id"`
	Name                 string `json:"name"`
	EvaluateTargetHealth bool   `json:"evaluate_target_health"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53RecordList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsRoute53Record `json:"items"`
}
