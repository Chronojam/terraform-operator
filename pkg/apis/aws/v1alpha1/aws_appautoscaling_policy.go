// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAppautoscalingPolicySpec `json"spec"`
}

type AwsAppautoscalingPolicySpec struct {
	ResourceId                               string                                                                `json:"resource_id"`
	Cooldown                                 int                                                                   `json:"cooldown"`
	MetricAggregationType                    string                                                                `json:"metric_aggregation_type"`
	TargetTrackingScalingPolicyConfiguration []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"target_tracking_scaling_policy_configuration"`
	ScalableDimension                        string                                                                `json:"scalable_dimension"`
	ServiceNamespace                         string                                                                `json:"service_namespace"`
	StepAdjustment                           AwsAppautoscalingPolicySpecStepAdjustment                             `json:"step_adjustment"`
	Name                                     string                                                                `json:"name"`
	Alarms                                   []string                                                              `json:"alarms"`
	AdjustmentType                           string                                                                `json:"adjustment_type"`
	MinAdjustmentMagnitude                   int                                                                   `json:"min_adjustment_magnitude"`
	Arn                                      string                                                                `json:"arn"`
	PolicyType                               string                                                                `json:"policy_type"`
	StepScalingPolicyConfiguration           []AwsAppautoscalingPolicySpecStepScalingPolicyConfiguration           `json:"step_scaling_policy_configuration"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration struct {
	ScaleInCooldown               int                                                                                                `json:"scale_in_cooldown"`
	ScaleOutCooldown              int                                                                                                `json:"scale_out_cooldown"`
	TargetValue                   float64                                                                                            `json:"target_value"`
	CustomizedMetricSpecification []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification `json:"customized_metric_specification"`
	PredefinedMetricSpecification []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification `json:"predefined_metric_specification"`
	DisableScaleIn                bool                                                                                               `json:"disable_scale_in"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification struct {
	Dimensions AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions `json:"dimensions"`
	MetricName string                                                                                                     `json:"metric_name"`
	Namespace  string                                                                                                     `json:"namespace"`
	Statistic  string                                                                                                     `json:"statistic"`
	Unit       string                                                                                                     `json:"unit"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification struct {
	ResourceLabel        string `json:"resource_label"`
	PredefinedMetricType string `json:"predefined_metric_type"`
}

type AwsAppautoscalingPolicySpecStepAdjustment struct {
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
}

type AwsAppautoscalingPolicySpecStepScalingPolicyConfiguration struct {
	AdjustmentType         string                                                                  `json:"adjustment_type"`
	Cooldown               int                                                                     `json:"cooldown"`
	MetricAggregationType  string                                                                  `json:"metric_aggregation_type"`
	MinAdjustmentMagnitude int                                                                     `json:"min_adjustment_magnitude"`
	StepAdjustment         AwsAppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment `json:"step_adjustment"`
}

type AwsAppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment struct {
	MetricIntervalLowerBound float64 `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound float64 `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int     `json:"scaling_adjustment"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAppautoscalingPolicy `json"items"`
}
