// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPool struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCognitoUserPoolSpec `json"spec"`
}

type AwsCognitoUserPoolSpec struct {
	CreationDate                string                                              `json:"creation_date"`
	EmailConfiguration          []AwsCognitoUserPoolSpecEmailConfiguration          `json:"email_configuration"`
	EmailVerificationSubject    string                                              `json:"email_verification_subject"`
	Schema                      AwsCognitoUserPoolSpecSchema                        `json:"schema"`
	SmsConfiguration            []AwsCognitoUserPoolSpecSmsConfiguration            `json:"sms_configuration"`
	Tags                        map[string]string                                   `json:"tags"`
	Arn                         string                                              `json:"arn"`
	AutoVerifiedAttributes      string                                              `json:"auto_verified_attributes"`
	MfaConfiguration            string                                              `json:"mfa_configuration"`
	Name                        string                                              `json:"name"`
	AdminCreateUserConfig       []AwsCognitoUserPoolSpecAdminCreateUserConfig       `json:"admin_create_user_config"`
	AliasAttributes             string                                              `json:"alias_attributes"`
	EmailVerificationMessage    string                                              `json:"email_verification_message"`
	LambdaConfig                []AwsCognitoUserPoolSpecLambdaConfig                `json:"lambda_config"`
	SmsVerificationMessage      string                                              `json:"sms_verification_message"`
	UsernameAttributes          []string                                            `json:"username_attributes"`
	DeviceConfiguration         []AwsCognitoUserPoolSpecDeviceConfiguration         `json:"device_configuration"`
	LastModifiedDate            string                                              `json:"last_modified_date"`
	PasswordPolicy              []AwsCognitoUserPoolSpecPasswordPolicy              `json:"password_policy"`
	SmsAuthenticationMessage    string                                              `json:"sms_authentication_message"`
	VerificationMessageTemplate []AwsCognitoUserPoolSpecVerificationMessageTemplate `json:"verification_message_template"`
}

type AwsCognitoUserPoolSpecEmailConfiguration struct {
	ReplyToEmailAddress string `json:"reply_to_email_address"`
	SourceArn           string `json:"source_arn"`
}

type AwsCognitoUserPoolSpecSchema struct {
	NumberAttributeConstraints []AwsCognitoUserPoolSpecSchemaNumberAttributeConstraints `json:"number_attribute_constraints"`
	Required                   bool                                                     `json:"required"`
	StringAttributeConstraints []AwsCognitoUserPoolSpecSchemaStringAttributeConstraints `json:"string_attribute_constraints"`
	AttributeDataType          string                                                   `json:"attribute_data_type"`
	DeveloperOnlyAttribute     bool                                                     `json:"developer_only_attribute"`
	Mutable                    bool                                                     `json:"mutable"`
	Name                       string                                                   `json:"name"`
}

type AwsCognitoUserPoolSpecSchemaNumberAttributeConstraints struct {
	MinValue string `json:"min_value"`
	MaxValue string `json:"max_value"`
}

type AwsCognitoUserPoolSpecSchemaStringAttributeConstraints struct {
	MinLength string `json:"min_length"`
	MaxLength string `json:"max_length"`
}

type AwsCognitoUserPoolSpecSmsConfiguration struct {
	ExternalId   string `json:"external_id"`
	SnsCallerArn string `json:"sns_caller_arn"`
}

type AwsCognitoUserPoolSpecAdminCreateUserConfig struct {
	AllowAdminCreateUserOnly  bool                                                               `json:"allow_admin_create_user_only"`
	InviteMessageTemplate     []AwsCognitoUserPoolSpecAdminCreateUserConfigInviteMessageTemplate `json:"invite_message_template"`
	UnusedAccountValidityDays int                                                                `json:"unused_account_validity_days"`
}

type AwsCognitoUserPoolSpecAdminCreateUserConfigInviteMessageTemplate struct {
	EmailMessage string `json:"email_message"`
	EmailSubject string `json:"email_subject"`
	SmsMessage   string `json:"sms_message"`
}

type AwsCognitoUserPoolSpecLambdaConfig struct {
	DefineAuthChallenge         string `json:"define_auth_challenge"`
	PostConfirmation            string `json:"post_confirmation"`
	PreAuthentication           string `json:"pre_authentication"`
	PreTokenGeneration          string `json:"pre_token_generation"`
	UserMigration               string `json:"user_migration"`
	VerifyAuthChallengeResponse string `json:"verify_auth_challenge_response"`
	CreateAuthChallenge         string `json:"create_auth_challenge"`
	CustomMessage               string `json:"custom_message"`
	PostAuthentication          string `json:"post_authentication"`
	PreSignUp                   string `json:"pre_sign_up"`
}

type AwsCognitoUserPoolSpecDeviceConfiguration struct {
	DeviceOnlyRememberedOnUserPrompt bool `json:"device_only_remembered_on_user_prompt"`
	ChallengeRequiredOnNewDevice     bool `json:"challenge_required_on_new_device"`
}

type AwsCognitoUserPoolSpecPasswordPolicy struct {
	MinimumLength    int  `json:"minimum_length"`
	RequireLowercase bool `json:"require_lowercase"`
	RequireNumbers   bool `json:"require_numbers"`
	RequireSymbols   bool `json:"require_symbols"`
	RequireUppercase bool `json:"require_uppercase"`
}

type AwsCognitoUserPoolSpecVerificationMessageTemplate struct {
	EmailSubjectByLink string `json:"email_subject_by_link"`
	SmsMessage         string `json:"sms_message"`
	DefaultEmailOption string `json:"default_email_option"`
	EmailMessage       string `json:"email_message"`
	EmailMessageByLink string `json:"email_message_by_link"`
	EmailSubject       string `json:"email_subject"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCognitoUserPool `json"items"`
}
