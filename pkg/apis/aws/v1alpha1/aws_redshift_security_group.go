// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftSecurityGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRedshiftSecurityGroupSpec `json"spec"`
}

type AwsRedshiftSecurityGroupSpec struct {
	Name        string                              `json:"name"`
	Description string                              `json:"description"`
	Ingress     AwsRedshiftSecurityGroupSpecIngress `json:"ingress"`
}

type AwsRedshiftSecurityGroupSpecIngress struct {
	Cidr                 string `json:"cidr"`
	SecurityGroupName    string `json:"security_group_name"`
	SecurityGroupOwnerId string `json:"security_group_owner_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftSecurityGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRedshiftSecurityGroup `json"items"`
}
