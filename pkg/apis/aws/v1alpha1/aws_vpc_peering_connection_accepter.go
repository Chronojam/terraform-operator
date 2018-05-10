// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcPeeringConnectionAccepter struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcPeeringConnectionAccepterSpec `json"spec"`
}

type AwsVpcPeeringConnectionAccepterSpec struct {
	AutoAccept             bool                                         `json:"auto_accept"`
	AcceptStatus           string                                       `json:"accept_status"`
	Requester              AwsVpcPeeringConnectionAccepterSpecRequester `json:"requester"`
	Accepter               AwsVpcPeeringConnectionAccepterSpecAccepter  `json:"accepter"`
	Tags                   map[string]string                            `json:"tags"`
	VpcPeeringConnectionId string                                       `json:"vpc_peering_connection_id"`
	VpcId                  string                                       `json:"vpc_id"`
	PeerVpcId              string                                       `json:"peer_vpc_id"`
	PeerOwnerId            string                                       `json:"peer_owner_id"`
	PeerRegion             string                                       `json:"peer_region"`
}

type AwsVpcPeeringConnectionAccepterSpecRequester struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionAccepterSpecAccepter struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcPeeringConnectionAccepterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcPeeringConnectionAccepter `json"items"`
}
