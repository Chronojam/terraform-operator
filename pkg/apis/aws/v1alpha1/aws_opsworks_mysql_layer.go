// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksMysqlLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksMysqlLayerSpec `json"spec"`
}

type AwsOpsworksMysqlLayerSpec struct {
	AutoAssignPublicIps        bool                               `json:"auto_assign_public_ips"`
	CustomDeployRecipes        []string                           `json:"custom_deploy_recipes"`
	CustomUndeployRecipes      []string                           `json:"custom_undeploy_recipes"`
	AutoHealing                bool                               `json:"auto_healing"`
	UseEbsOptimizedInstances   bool                               `json:"use_ebs_optimized_instances"`
	CustomInstanceProfileArn   string                             `json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds     string                             `json:"custom_security_group_ids"`
	InstallUpdatesOnBoot       bool                               `json:"install_updates_on_boot"`
	DrainElbOnShutdown         bool                               `json:"drain_elb_on_shutdown"`
	SystemPackages             string                             `json:"system_packages"`
	RootPassword               string                             `json:"root_password"`
	ElasticLoadBalancer        string                             `json:"elastic_load_balancer"`
	CustomSetupRecipes         []string                           `json:"custom_setup_recipes"`
	CustomShutdownRecipes      []string                           `json:"custom_shutdown_recipes"`
	CustomJson                 string                             `json:"custom_json"`
	AutoAssignElasticIps       bool                               `json:"auto_assign_elastic_ips"`
	CustomConfigureRecipes     []string                           `json:"custom_configure_recipes"`
	InstanceShutdownTimeout    int                                `json:"instance_shutdown_timeout"`
	StackId                    string                             `json:"stack_id"`
	EbsVolume                  AwsOpsworksMysqlLayerSpecEbsVolume `json:"ebs_volume"`
	Name                       string                             `json:"name"`
	RootPasswordOnAllInstances bool                               `json:"root_password_on_all_instances"`
}

type AwsOpsworksMysqlLayerSpecEbsVolume struct {
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksMysqlLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksMysqlLayer `json"items"`
}
