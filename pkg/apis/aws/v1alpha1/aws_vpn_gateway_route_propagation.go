// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayRoutePropagation struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsVpnGatewayRoutePropagationSpec `json:"spec"`
}

type AwsVpnGatewayRoutePropagationSpec struct {
	VpnGatewayId string `json:"vpn_gateway_id"`
	RouteTableId string `json:"route_table_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayRoutePropagationList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsVpnGatewayRoutePropagation `json:"items"`
}
