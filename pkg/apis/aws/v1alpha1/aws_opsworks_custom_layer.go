// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksCustomLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksCustomLayerSpec `json"spec"`
}

type AwsOpsworksCustomLayerSpec struct {
	StackId                  string                              `json:"stack_id"`
	Name                     string                              `json:"name"`
	CustomInstanceProfileArn string                              `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string                            `json:"custom_setup_recipes"`
	CustomConfigureRecipes   []string                            `json:"custom_configure_recipes"`
	CustomJson               string                              `json:"custom_json"`
	AutoHealing              bool                                `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                                `json:"install_updates_on_boot"`
	AutoAssignPublicIps      bool                                `json:"auto_assign_public_ips"`
	CustomDeployRecipes      []string                            `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                            `json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds   string                              `json:"custom_security_group_ids"`
	ShortName                string                              `json:"short_name"`
	AutoAssignElasticIps     bool                                `json:"auto_assign_elastic_ips"`
	CustomShutdownRecipes    []string                            `json:"custom_shutdown_recipes"`
	InstanceShutdownTimeout  int                                 `json:"instance_shutdown_timeout"`
	DrainElbOnShutdown       bool                                `json:"drain_elb_on_shutdown"`
	SystemPackages           string                              `json:"system_packages"`
	EbsVolume                AwsOpsworksCustomLayerSpecEbsVolume `json:"ebs_volume"`
	ElasticLoadBalancer      string                              `json:"elastic_load_balancer"`
	UseEbsOptimizedInstances bool                                `json:"use_ebs_optimized_instances"`
}

type AwsOpsworksCustomLayerSpecEbsVolume struct {
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksCustomLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksCustomLayer `json"items"`
}
