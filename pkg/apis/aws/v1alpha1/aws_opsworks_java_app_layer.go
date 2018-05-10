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
	ElasticLoadBalancer      string                               `json:"elastic_load_balancer"`
	CustomUndeployRecipes    []string                             `json:"custom_undeploy_recipes"`
	AutoHealing              bool                                 `json:"auto_healing"`
	StackId                  string                               `json:"stack_id"`
	JvmType                  string                               `json:"jvm_type"`
	JvmVersion               string                               `json:"jvm_version"`
	AutoAssignElasticIps     bool                                 `json:"auto_assign_elastic_ips"`
	CustomConfigureRecipes   []string                             `json:"custom_configure_recipes"`
	SystemPackages           string                               `json:"system_packages"`
	JvmOptions               string                               `json:"jvm_options"`
	AppServerVersion         string                               `json:"app_server_version"`
	CustomShutdownRecipes    []string                             `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds   string                               `json:"custom_security_group_ids"`
	CustomJson               string                               `json:"custom_json"`
	InstanceShutdownTimeout  int                                  `json:"instance_shutdown_timeout"`
	DrainElbOnShutdown       bool                                 `json:"drain_elb_on_shutdown"`
	EbsVolume                AwsOpsworksJavaAppLayerSpecEbsVolume `json:"ebs_volume"`
	Name                     string                               `json:"name"`
	AppServer                string                               `json:"app_server"`
	AutoAssignPublicIps      bool                                 `json:"auto_assign_public_ips"`
	CustomInstanceProfileArn string                               `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string                             `json:"custom_setup_recipes"`
	CustomDeployRecipes      []string                             `json:"custom_deploy_recipes"`
	InstallUpdatesOnBoot     bool                                 `json:"install_updates_on_boot"`
	UseEbsOptimizedInstances bool                                 `json:"use_ebs_optimized_instances"`
}

type AwsOpsworksJavaAppLayerSpecEbsVolume struct {
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksJavaAppLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksJavaAppLayer `json"items"`
}
