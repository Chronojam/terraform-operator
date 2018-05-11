// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchMetricAlarm struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsCloudwatchMetricAlarmSpec `json:"spec"`
}

type AwsCloudwatchMetricAlarmSpec struct {
	TreatMissingData                  string            `json:"treat_missing_data"`
	EvaluateLowSampleCountPercentiles string            `json:"evaluate_low_sample_count_percentiles"`
	Threshold                         float64           `json:"threshold"`
	Dimensions                        map[string]string `json:"dimensions"`
	ExtendedStatistic                 string            `json:"extended_statistic"`
	Namespace                         string            `json:"namespace"`
	InsufficientDataActions           string            `json:"insufficient_data_actions"`
	AlarmDescription                  string            `json:"alarm_description"`
	DatapointsToAlarm                 int               `json:"datapoints_to_alarm"`
	OkActions                         string            `json:"ok_actions"`
	ComparisonOperator                string            `json:"comparison_operator"`
	AlarmActions                      string            `json:"alarm_actions"`
	MetricName                        string            `json:"metric_name"`
	Period                            int               `json:"period"`
	Statistic                         string            `json:"statistic"`
	ActionsEnabled                    bool              `json:"actions_enabled"`
	Unit                              string            `json:"unit"`
	AlarmName                         string            `json:"alarm_name"`
	EvaluationPeriods                 int               `json:"evaluation_periods"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchMetricAlarmList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsCloudwatchMetricAlarm `json:"items"`
}
