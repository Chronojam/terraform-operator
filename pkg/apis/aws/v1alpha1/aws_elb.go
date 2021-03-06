// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElb struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsElbSpec `json:"spec"`
}

type AwsElbSpec struct {
	NamePrefix                string                  `json:"name_prefix"`
	SourceSecurityGroupId     string                  `json:"source_security_group_id"`
	ConnectionDraining        bool                    `json:"connection_draining"`
	AccessLogs                []AwsElbSpecAccessLogs  `json:"access_logs"`
	Name                      string                  `json:"name"`
	AvailabilityZones         string                  `json:"availability_zones"`
	SourceSecurityGroup       string                  `json:"source_security_group"`
	Subnets                   string                  `json:"subnets"`
	Listener                  AwsElbSpecListener      `json:"listener"`
	Internal                  bool                    `json:"internal"`
	Instances                 string                  `json:"instances"`
	ConnectionDrainingTimeout int                     `json:"connection_draining_timeout"`
	Tags                      map[string]string       `json:"tags"`
	DnsName                   string                  `json:"dns_name"`
	ZoneId                    string                  `json:"zone_id"`
	Arn                       string                  `json:"arn"`
	CrossZoneLoadBalancing    bool                    `json:"cross_zone_load_balancing"`
	SecurityGroups            string                  `json:"security_groups"`
	IdleTimeout               int                     `json:"idle_timeout"`
	HealthCheck               []AwsElbSpecHealthCheck `json:"health_check"`
}

type AwsElbSpecAccessLogs struct {
	Interval     int    `json:"interval"`
	Bucket       string `json:"bucket"`
	BucketPrefix string `json:"bucket_prefix"`
	Enabled      bool   `json:"enabled"`
}

type AwsElbSpecListener struct {
	LbPort           int    `json:"lb_port"`
	LbProtocol       string `json:"lb_protocol"`
	SslCertificateId string `json:"ssl_certificate_id"`
	InstancePort     int    `json:"instance_port"`
	InstanceProtocol string `json:"instance_protocol"`
}

type AwsElbSpecHealthCheck struct {
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Target             string `json:"target"`
	Interval           int    `json:"interval"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElbList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsElb `json:"items"`
}
