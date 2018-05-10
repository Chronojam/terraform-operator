// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueConnection struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGlueConnectionSpec `json"spec"`
}

type AwsGlueConnectionSpec struct {
	ConnectionProperties           map[string]string                                     `json:"connection_properties"`
	ConnectionType                 string                                                `json:"connection_type"`
	Description                    string                                                `json:"description"`
	MatchCriteria                  []string                                              `json:"match_criteria"`
	Name                           string                                                `json:"name"`
	PhysicalConnectionRequirements []AwsGlueConnectionSpecPhysicalConnectionRequirements `json:"physical_connection_requirements"`
	CatalogId                      string                                                `json:"catalog_id"`
}

type AwsGlueConnectionSpecPhysicalConnectionRequirements struct {
	SecurityGroupIdList []string `json:"security_group_id_list"`
	SubnetId            string   `json:"subnet_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueConnectionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGlueConnection `json"items"`
}
