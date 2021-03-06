// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointServiceAllowedPrincipal struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsVpcEndpointServiceAllowedPrincipalSpec `json:"spec"`
}

type AwsVpcEndpointServiceAllowedPrincipalSpec struct {
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`
	PrincipalArn         string `json:"principal_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointServiceAllowedPrincipalList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsVpcEndpointServiceAllowedPrincipal `json:"items"`
}
