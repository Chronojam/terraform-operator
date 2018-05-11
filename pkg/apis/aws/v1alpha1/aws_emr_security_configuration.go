// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrSecurityConfiguration struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsEmrSecurityConfigurationSpec `json:"spec"`
}

type AwsEmrSecurityConfigurationSpec struct {
	Configuration string `json:"configuration"`
	CreationDate  string `json:"creation_date"`
	Name          string `json:"name"`
	NamePrefix    string `json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrSecurityConfigurationList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsEmrSecurityConfiguration `json:"items"`
}
