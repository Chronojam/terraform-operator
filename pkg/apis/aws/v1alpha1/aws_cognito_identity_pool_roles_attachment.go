// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoIdentityPoolRolesAttachment struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsCognitoIdentityPoolRolesAttachmentSpec `json:"spec"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpec struct {
	IdentityPoolId string                                                    `json:"identity_pool_id"`
	RoleMapping    AwsCognitoIdentityPoolRolesAttachmentSpecRoleMapping      `json:"role_mapping"`
	Roles          map[string]AwsCognitoIdentityPoolRolesAttachmentSpecRoles `json:"roles"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpecRoleMapping struct {
	IdentityProvider        string                                                            `json:"identity_provider"`
	AmbiguousRoleResolution string                                                            `json:"ambiguous_role_resolution"`
	MappingRule             []AwsCognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule `json:"mapping_rule"`
	Type                    string                                                            `json:"type"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule struct {
	Claim     string `json:"claim"`
	MatchType string `json:"match_type"`
	RoleArn   string `json:"role_arn"`
	Value     string `json:"value"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpecRoles struct {
	Authenticated   string `json:"authenticated"`
	Unauthenticated string `json:"unauthenticated"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoIdentityPoolRolesAttachmentList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsCognitoIdentityPoolRolesAttachment `json:"items"`
}
