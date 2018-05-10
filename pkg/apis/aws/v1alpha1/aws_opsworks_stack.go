// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStack struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksStackSpec `json"spec"`
}

type AwsOpsworksStackSpec struct {
	CustomCookbooksSource       []AwsOpsworksStackSpecCustomCookbooksSource `json:"custom_cookbooks_source"`
	DefaultAvailabilityZone     string                                      `json:"default_availability_zone"`
	Tags                        map[string]string                           `json:"tags"`
	AgentVersion                string                                      `json:"agent_version"`
	Name                        string                                      `json:"name"`
	ConfigurationManagerVersion string                                      `json:"configuration_manager_version"`
	DefaultSshKeyName           string                                      `json:"default_ssh_key_name"`
	DefaultSubnetId             string                                      `json:"default_subnet_id"`
	StackEndpoint               string                                      `json:"stack_endpoint"`
	Color                       string                                      `json:"color"`
	ConfigurationManagerName    string                                      `json:"configuration_manager_name"`
	CustomJson                  string                                      `json:"custom_json"`
	DefaultInstanceProfileArn   string                                      `json:"default_instance_profile_arn"`
	DefaultOs                   string                                      `json:"default_os"`
	DefaultRootDeviceType       string                                      `json:"default_root_device_type"`
	Arn                         string                                      `json:"arn"`
	Region                      string                                      `json:"region"`
	ServiceRoleArn              string                                      `json:"service_role_arn"`
	UseCustomCookbooks          bool                                        `json:"use_custom_cookbooks"`
	UseOpsworksSecurityGroups   bool                                        `json:"use_opsworks_security_groups"`
	VpcId                       string                                      `json:"vpc_id"`
	ManageBerkshelf             bool                                        `json:"manage_berkshelf"`
	BerkshelfVersion            string                                      `json:"berkshelf_version"`
	HostnameTheme               string                                      `json:"hostname_theme"`
}

type AwsOpsworksStackSpecCustomCookbooksSource struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
	Type     string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStackList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksStack `json"items"`
}
