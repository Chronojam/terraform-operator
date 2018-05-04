// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkEnvironment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticBeanstalkEnvironmentSpec `json"spec"`
}

type AwsElasticBeanstalkEnvironmentSpec struct {
	Application          string                                        `json:"application"`
	SolutionStackName    string                                        `json:"solution_stack_name"`
	Tags                 map[string]string                             `json:"tags"`
	Arn                  string                                        `json:"arn"`
	TemplateName         string                                        `json:"template_name"`
	Queues               []string                                      `json:"queues"`
	Triggers             []string                                      `json:"triggers"`
	Setting              AwsElasticBeanstalkEnvironmentSpecSetting     `json:"setting"`
	AllSettings          AwsElasticBeanstalkEnvironmentSpecAllSettings `json:"all_settings"`
	WaitForReadyTimeout  string                                        `json:"wait_for_ready_timeout"`
	Instances            []string                                      `json:"instances"`
	Name                 string                                        `json:"name"`
	Description          string                                        `json:"description"`
	VersionLabel         string                                        `json:"version_label"`
	CnamePrefix          string                                        `json:"cname_prefix"`
	LaunchConfigurations []string                                      `json:"launch_configurations"`
	LoadBalancers        []string                                      `json:"load_balancers"`
	Cname                string                                        `json:"cname"`
	Tier                 string                                        `json:"tier"`
	PollInterval         string                                        `json:"poll_interval"`
	AutoscalingGroups    []string                                      `json:"autoscaling_groups"`
}

type AwsElasticBeanstalkEnvironmentSpecSetting struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Resource  string `json:"resource"`
}

type AwsElasticBeanstalkEnvironmentSpecAllSettings struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Resource  string `json:"resource"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkEnvironmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticBeanstalkEnvironment `json"items"`
}
