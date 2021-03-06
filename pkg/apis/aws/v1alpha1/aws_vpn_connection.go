// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnConnection struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsVpnConnectionSpec `json:"spec"`
}

type AwsVpnConnectionSpec struct {
	Tunnel2InsideCidr            string                           `json:"tunnel2_inside_cidr"`
	CustomerGatewayConfiguration string                           `json:"customer_gateway_configuration"`
	Tunnel1Address               string                           `json:"tunnel1_address"`
	Tunnel1InsideCidr            string                           `json:"tunnel1_inside_cidr"`
	Tags                         map[string]string                `json:"tags"`
	Tunnel1CgwInsideAddress      string                           `json:"tunnel1_cgw_inside_address"`
	Tunnel1VgwInsideAddress      string                           `json:"tunnel1_vgw_inside_address"`
	Tunnel2Address               string                           `json:"tunnel2_address"`
	VgwTelemetry                 AwsVpnConnectionSpecVgwTelemetry `json:"vgw_telemetry"`
	CustomerGatewayId            string                           `json:"customer_gateway_id"`
	Tunnel2PresharedKey          string                           `json:"tunnel2_preshared_key"`
	Tunnel1BgpAsn                string                           `json:"tunnel1_bgp_asn"`
	Tunnel2BgpHoldtime           int                              `json:"tunnel2_bgp_holdtime"`
	StaticRoutesOnly             bool                             `json:"static_routes_only"`
	Type                         string                           `json:"type"`
	Tunnel1PresharedKey          string                           `json:"tunnel1_preshared_key"`
	Tunnel1BgpHoldtime           int                              `json:"tunnel1_bgp_holdtime"`
	Tunnel2CgwInsideAddress      string                           `json:"tunnel2_cgw_inside_address"`
	Tunnel2VgwInsideAddress      string                           `json:"tunnel2_vgw_inside_address"`
	Tunnel2BgpAsn                string                           `json:"tunnel2_bgp_asn"`
	Routes                       AwsVpnConnectionSpecRoutes       `json:"routes"`
	VpnGatewayId                 string                           `json:"vpn_gateway_id"`
}

type AwsVpnConnectionSpecVgwTelemetry struct {
	LastStatusChange   string `json:"last_status_change"`
	OutsideIpAddress   string `json:"outside_ip_address"`
	Status             string `json:"status"`
	StatusMessage      string `json:"status_message"`
	AcceptedRouteCount int    `json:"accepted_route_count"`
}

type AwsVpnConnectionSpecRoutes struct {
	State                string `json:"state"`
	DestinationCidrBlock string `json:"destination_cidr_block"`
	Source               string `json:"source"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnConnectionList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsVpnConnection `json:"items"`
}
