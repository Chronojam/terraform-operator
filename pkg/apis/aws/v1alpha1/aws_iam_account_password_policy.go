// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccountPasswordPolicy struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               AwsIamAccountPasswordPolicySpec `json:"spec"`
}

type AwsIamAccountPasswordPolicySpec struct {
	ExpirePasswords            bool `json:"expire_passwords"`
	RequireLowercaseCharacters bool `json:"require_lowercase_characters"`
	RequireNumbers             bool `json:"require_numbers"`
	RequireUppercaseCharacters bool `json:"require_uppercase_characters"`
	AllowUsersToChangePassword bool `json:"allow_users_to_change_password"`
	MaxPasswordAge             int  `json:"max_password_age"`
	MinimumPasswordLength      int  `json:"minimum_password_length"`
	PasswordReusePrevention    int  `json:"password_reuse_prevention"`
	RequireSymbols             bool `json:"require_symbols"`
	HardExpiry                 bool `json:"hard_expiry"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccountPasswordPolicyList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`
	Items            []AwsIamAccountPasswordPolicy `json:"items"`
}
