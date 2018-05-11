// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAutoscalingPolicySpec `json"spec"`
}

type AwsAutoscalingPolicySpec struct {
	EstimatedInstanceWarmup     int                                                   `json:"estimated_instance_warmup"`
	MetricAggregationType       string                                                `json:"metric_aggregation_type"`
	StepAdjustment              AwsAutoscalingPolicySpecStepAdjustment                `json:"step_adjustment"`
	Arn                         string                                                `json:"arn"`
	AdjustmentType              string                                                `json:"adjustment_type"`
	AutoscalingGroupName        string                                                `json:"autoscaling_group_name"`
	MinAdjustmentMagnitude      int                                                   `json:"min_adjustment_magnitude"`
	MinAdjustmentStep           int                                                   `json:"min_adjustment_step"`
	ScalingAdjustment           int                                                   `json:"scaling_adjustment"`
	TargetTrackingConfiguration []AwsAutoscalingPolicySpecTargetTrackingConfiguration `json:"target_tracking_configuration"`
	Name                        string                                                `json:"name"`
	PolicyType                  string                                                `json:"policy_type"`
	Cooldown                    int                                                   `json:"cooldown"`
}

type AwsAutoscalingPolicySpecStepAdjustment struct {
	ScalingAdjustment        int    `json:"scaling_adjustment"`
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfiguration struct {
	PredefinedMetricSpecification []AwsAutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification `json:"predefined_metric_specification"`
	CustomizedMetricSpecification []AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification `json:"customized_metric_specification"`
	TargetValue                   float64                                                                            `json:"target_value"`
	DisableScaleIn                bool                                                                               `json:"disable_scale_in"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	ResourceLabel        string `json:"resource_label"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification struct {
	Unit            string                                                                                            `json:"unit"`
	MetricDimension []AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension `json:"metric_dimension"`
	MetricName      string                                                                                            `json:"metric_name"`
	Namespace       string                                                                                            `json:"namespace"`
	Statistic       string                                                                                            `json:"statistic"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAutoscalingPolicy `json"items"`
}
