// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbOptionGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDbOptionGroupSpec `json"spec"`
}

type AwsDbOptionGroupSpec struct {
	Arn                    string                     `json:"arn"`
	Name                   string                     `json:"name"`
	NamePrefix             string                     `json:"name_prefix"`
	EngineName             string                     `json:"engine_name"`
	MajorEngineVersion     string                     `json:"major_engine_version"`
	OptionGroupDescription string                     `json:"option_group_description"`
	Option                 AwsDbOptionGroupSpecOption `json:"option"`
	Tags                   map[string]string          `json:"tags"`
}

type AwsDbOptionGroupSpecOption struct {
	OptionName                  string                                   `json:"option_name"`
	OptionSettings              AwsDbOptionGroupSpecOptionOptionSettings `json:"option_settings"`
	Port                        int                                      `json:"port"`
	DbSecurityGroupMemberships  string                                   `json:"db_security_group_memberships"`
	VpcSecurityGroupMemberships string                                   `json:"vpc_security_group_memberships"`
	Version                     string                                   `json:"version"`
}

type AwsDbOptionGroupSpecOptionOptionSettings struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbOptionGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDbOptionGroup `json"items"`
}
