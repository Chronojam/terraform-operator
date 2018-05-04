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
	HealthCheckGracePeriod int                                         `json:"health_check_grace_period"`
	TargetGroupArns        string                                      `json:"target_group_arns"`
	Tag                    AwsAutoscalingGroupSpecTag                  `json:"tag"`
	MinSize                int                                         `json:"min_size"`
	ProtectFromScaleIn     bool                                        `json:"protect_from_scale_in"`
	InitialLifecycleHook   AwsAutoscalingGroupSpecInitialLifecycleHook `json:"initial_lifecycle_hook"`
	NamePrefix             string                                      `json:"name_prefix"`
	DesiredCapacity        int                                         `json:"desired_capacity"`
	TerminationPolicies    []string                                    `json:"termination_policies"`
	WaitForCapacityTimeout string                                      `json:"wait_for_capacity_timeout"`
	MetricsGranularity     string                                      `json:"metrics_granularity"`
	HealthCheckType        string                                      `json:"health_check_type"`
	WaitForElbCapacity     int                                         `json:"wait_for_elb_capacity"`
	Name                   string                                      `json:"name"`
	LaunchConfiguration    string                                      `json:"launch_configuration"`
	AvailabilityZones      string                                      `json:"availability_zones"`
	PlacementGroup         string                                      `json:"placement_group"`
	LoadBalancers          string                                      `json:"load_balancers"`
	VpcZoneIdentifier      string                                      `json:"vpc_zone_identifier"`
	EnabledMetrics         string                                      `json:"enabled_metrics"`
	Arn                    string                                      `json:"arn"`
	ServiceLinkedRoleArn   string                                      `json:"service_linked_role_arn"`
	MinElbCapacity         int                                         `json:"min_elb_capacity"`
	MaxSize                int                                         `json:"max_size"`
	LaunchTemplate         []AwsAutoscalingGroupSpecLaunchTemplate     `json:"launch_template"`
	DefaultCooldown        int                                         `json:"default_cooldown"`
	ForceDelete            bool                                        `json:"force_delete"`
	SuspendedProcesses     string                                      `json:"suspended_processes"`
	Tags                   []map[string]string                         `json:"tags"`
}

type AwsAutoscalingGroupSpecTag struct {
	Key               string `json:"key"`
	Value             string `json:"value"`
	PropagateAtLaunch bool   `json:"propagate_at_launch"`
}

type AwsAutoscalingGroupSpecInitialLifecycleHook struct {
	RoleArn               string `json:"role_arn"`
	Name                  string `json:"name"`
	DefaultResult         string `json:"default_result"`
	HeartbeatTimeout      int    `json:"heartbeat_timeout"`
	LifecycleTransition   string `json:"lifecycle_transition"`
	NotificationMetadata  string `json:"notification_metadata"`
	NotificationTargetArn string `json:"notification_target_arn"`
}

type AwsAutoscalingGroupSpecLaunchTemplate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAutoscalingGroup `json"items"`
}
