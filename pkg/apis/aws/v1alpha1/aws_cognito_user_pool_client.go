// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolClient struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCognitoUserPoolClientSpec `json"spec"`
}

type AwsCognitoUserPoolClientSpec struct {
	AllowedOauthFlowsUserPoolClient bool     `json:"allowed_oauth_flows_user_pool_client"`
	AllowedOauthScopes              string   `json:"allowed_oauth_scopes"`
	DefaultRedirectUri              string   `json:"default_redirect_uri"`
	ClientSecret                    string   `json:"client_secret"`
	GenerateSecret                  bool     `json:"generate_secret"`
	UserPoolId                      string   `json:"user_pool_id"`
	ExplicitAuthFlows               string   `json:"explicit_auth_flows"`
	AllowedOauthFlows               string   `json:"allowed_oauth_flows"`
	Name                            string   `json:"name"`
	ReadAttributes                  string   `json:"read_attributes"`
	WriteAttributes                 string   `json:"write_attributes"`
	RefreshTokenValidity            int      `json:"refresh_token_validity"`
	CallbackUrls                    []string `json:"callback_urls"`
	LogoutUrls                      []string `json:"logout_urls"`
	SupportedIdentityProviders      []string `json:"supported_identity_providers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolClientList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCognitoUserPoolClient `json"items"`
}
