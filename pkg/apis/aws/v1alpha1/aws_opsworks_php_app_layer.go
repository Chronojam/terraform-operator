// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksPhpAppLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksPhpAppLayerSpec `json"spec"`
}

type AwsOpsworksPhpAppLayerSpec struct {
	CustomSetupRecipes       []string                            `json:"custom_setup_recipes"`
	CustomDeployRecipes      []string                            `json:"custom_deploy_recipes"`
	StackId                  string                              `json:"stack_id"`
	Name                     string                              `json:"name"`
	ElasticLoadBalancer      string                              `json:"elastic_load_balancer"`
	CustomJson               string                              `json:"custom_json"`
	AutoHealing              bool                                `json:"auto_healing"`
	InstanceShutdownTimeout  int                                 `json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances bool                                `json:"use_ebs_optimized_instances"`
	EbsVolume                AwsOpsworksPhpAppLayerSpecEbsVolume `json:"ebs_volume"`
	CustomUndeployRecipes    []string                            `json:"custom_undeploy_recipes"`
	CustomShutdownRecipes    []string                            `json:"custom_shutdown_recipes"`
	InstallUpdatesOnBoot     bool                                `json:"install_updates_on_boot"`
	SystemPackages           string                              `json:"system_packages"`
	AutoAssignElasticIps     bool                                `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                                `json:"auto_assign_public_ips"`
	CustomInstanceProfileArn string                              `json:"custom_instance_profile_arn"`
	CustomConfigureRecipes   []string                            `json:"custom_configure_recipes"`
	CustomSecurityGroupIds   string                              `json:"custom_security_group_ids"`
	DrainElbOnShutdown       bool                                `json:"drain_elb_on_shutdown"`
}

type AwsOpsworksPhpAppLayerSpecEbsVolume struct {
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksPhpAppLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksPhpAppLayer `json"items"`
}
