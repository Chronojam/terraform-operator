// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalIpset struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsWafregionalIpsetSpec `json:"spec"`
}

type AwsWafregionalIpsetSpec struct {
	Name            string                                 `json:"name"`
	IpSetDescriptor AwsWafregionalIpsetSpecIpSetDescriptor `json:"ip_set_descriptor"`
}

type AwsWafregionalIpsetSpecIpSetDescriptor struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalIpsetList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsWafregionalIpset `json:"items"`
}
