// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStaticWebLayer struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsOpsworksStaticWebLayerSpec `json:"spec"`
}

type AwsOpsworksStaticWebLayerSpec struct {
	AutoHealing              bool                                   `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                                   `json:"install_updates_on_boot"`
	CustomSetupRecipes       []string                               `json:"custom_setup_recipes"`
	CustomDeployRecipes      []string                               `json:"custom_deploy_recipes"`
	CustomShutdownRecipes    []string                               `json:"custom_shutdown_recipes"`
	CustomJson               string                                 `json:"custom_json"`
	UseEbsOptimizedInstances bool                                   `json:"use_ebs_optimized_instances"`
	AutoAssignPublicIps      bool                                   `json:"auto_assign_public_ips"`
	CustomConfigureRecipes   []string                               `json:"custom_configure_recipes"`
	CustomUndeployRecipes    []string                               `json:"custom_undeploy_recipes"`
	InstanceShutdownTimeout  int                                    `json:"instance_shutdown_timeout"`
	Name                     string                                 `json:"name"`
	CustomInstanceProfileArn string                                 `json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds   string                                 `json:"custom_security_group_ids"`
	StackId                  string                                 `json:"stack_id"`
	EbsVolume                AwsOpsworksStaticWebLayerSpecEbsVolume `json:"ebs_volume"`
	AutoAssignElasticIps     bool                                   `json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer      string                                 `json:"elastic_load_balancer"`
	DrainElbOnShutdown       bool                                   `json:"drain_elb_on_shutdown"`
	SystemPackages           string                                 `json:"system_packages"`
}

type AwsOpsworksStaticWebLayerSpecEbsVolume struct {
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStaticWebLayerList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsOpsworksStaticWebLayer `json:"items"`
}
