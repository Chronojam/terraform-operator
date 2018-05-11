// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAutoscalingGroupSpec `json"spec"`
}

type AwsAutoscalingGroupSpec struct {
	MinSize                int                                         `json:"min_size"`
	ForceDelete            bool                                        `json:"force_delete"`
	HealthCheckType        string                                      `json:"health_check_type"`
	VpcZoneIdentifier      string                                      `json:"vpc_zone_identifier"`
	TargetGroupArns        string                                      `json:"target_group_arns"`
	InitialLifecycleHook   AwsAutoscalingGroupSpecInitialLifecycleHook `json:"initial_lifecycle_hook"`
	Tags                   []map[string]string                         `json:"tags"`
	Name                   string                                      `json:"name"`
	DefaultCooldown        int                                         `json:"default_cooldown"`
	MetricsGranularity     string                                      `json:"metrics_granularity"`
	WaitForElbCapacity     int                                         `json:"wait_for_elb_capacity"`
	MinElbCapacity         int                                         `json:"min_elb_capacity"`
	LoadBalancers          string                                      `json:"load_balancers"`
	Arn                    string                                      `json:"arn"`
	ServiceLinkedRoleArn   string                                      `json:"service_linked_role_arn"`
	NamePrefix             string                                      `json:"name_prefix"`
	LaunchConfiguration    string                                      `json:"launch_configuration"`
	AvailabilityZones      string                                      `json:"availability_zones"`
	ProtectFromScaleIn     bool                                        `json:"protect_from_scale_in"`
	LaunchTemplate         []AwsAutoscalingGroupSpecLaunchTemplate     `json:"launch_template"`
	HealthCheckGracePeriod int                                         `json:"health_check_grace_period"`
	TerminationPolicies    []string                                    `json:"termination_policies"`
	WaitForCapacityTimeout string                                      `json:"wait_for_capacity_timeout"`
	DesiredCapacity        int                                         `json:"desired_capacity"`
	MaxSize                int                                         `json:"max_size"`
	PlacementGroup         string                                      `json:"placement_group"`
	EnabledMetrics         string                                      `json:"enabled_metrics"`
	SuspendedProcesses     string                                      `json:"suspended_processes"`
	Tag                    AwsAutoscalingGroupSpecTag                  `json:"tag"`
}

type AwsAutoscalingGroupSpecInitialLifecycleHook struct {
	HeartbeatTimeout      int    `json:"heartbeat_timeout"`
	LifecycleTransition   string `json:"lifecycle_transition"`
	NotificationMetadata  string `json:"notification_metadata"`
	NotificationTargetArn string `json:"notification_target_arn"`
	RoleArn               string `json:"role_arn"`
	Name                  string `json:"name"`
	DefaultResult         string `json:"default_result"`
}

type AwsAutoscalingGroupSpecLaunchTemplate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type AwsAutoscalingGroupSpecTag struct {
	Key               string `json:"key"`
	Value             string `json:"value"`
	PropagateAtLaunch bool   `json:"propagate_at_launch"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAutoscalingGroup `json"items"`
}
