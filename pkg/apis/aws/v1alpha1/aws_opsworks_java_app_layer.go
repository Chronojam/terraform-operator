// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksJavaAppLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksJavaAppLayerSpec `json"spec"`
}

type AwsOpsworksJavaAppLayerSpec struct {
	SystemPackages           string                               `json:"system_packages"`
	AppServerVersion         string                               `json:"app_server_version"`
	CustomSetupRecipes       []string                             `json:"custom_setup_recipes"`
	InstallUpdatesOnBoot     bool                                 `json:"install_updates_on_boot"`
	CustomUndeployRecipes    []string                             `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool                                 `json:"drain_elb_on_shutdown"`
	StackId                  string                               `json:"stack_id"`
	Name                     string                               `json:"name"`
	AppServer                string                               `json:"app_server"`
	AutoAssignElasticIps     bool                                 `json:"auto_assign_elastic_ips"`
	CustomConfigureRecipes   []string                             `json:"custom_configure_recipes"`
	CustomShutdownRecipes    []string                             `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds   string                               `json:"custom_security_group_ids"`
	InstanceShutdownTimeout  int                                  `json:"instance_shutdown_timeout"`
	JvmVersion               string                               `json:"jvm_version"`
	JvmOptions               string                               `json:"jvm_options"`
	AutoAssignPublicIps      bool                                 `json:"auto_assign_public_ips"`
	CustomDeployRecipes      []string                             `json:"custom_deploy_recipes"`
	CustomJson               string                               `json:"custom_json"`
	AutoHealing              bool                                 `json:"auto_healing"`
	UseEbsOptimizedInstances bool                                 `json:"use_ebs_optimized_instances"`
	EbsVolume                AwsOpsworksJavaAppLayerSpecEbsVolume `json:"ebs_volume"`
	JvmType                  string                               `json:"jvm_type"`
	CustomInstanceProfileArn string                               `json:"custom_instance_profile_arn"`
	ElasticLoadBalancer      string                               `json:"elastic_load_balancer"`
}

type AwsOpsworksJavaAppLayerSpecEbsVolume struct {
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksJavaAppLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksJavaAppLayer `json"items"`
}
