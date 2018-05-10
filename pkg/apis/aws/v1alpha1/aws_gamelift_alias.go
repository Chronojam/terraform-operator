// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGameliftAlias struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGameliftAliasSpec `json"spec"`
}

type AwsGameliftAliasSpec struct {
	Name            string                                `json:"name"`
	Description     string                                `json:"description"`
	RoutingStrategy []AwsGameliftAliasSpecRoutingStrategy `json:"routing_strategy"`
	Arn             string                                `json:"arn"`
}

type AwsGameliftAliasSpecRoutingStrategy struct {
	FleetId string `json:"fleet_id"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGameliftAliasList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGameliftAlias `json"items"`
}
