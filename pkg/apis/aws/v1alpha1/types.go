package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayAccount struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayAccountSpec `json"spec"`
}

type AwsApiGatewayAccountSpec struct {
	CloudwatchRoleArn string                                     `json:"cloudwatch_role_arn"`
	ThrottleSettings  []AwsApiGatewayAccountSpecThrottleSettings `json:"throttle_settings"`
}

type AwsApiGatewayAccountSpecThrottleSettings struct {
	BurstLimit int     `json:"burst_limit"`
	RateLimit  float64 `json:"rate_limit"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayAccountList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayAccount `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksMysqlLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksMysqlLayerSpec `json"spec"`
}

type AwsOpsworksMysqlLayerSpec struct {
	ElasticLoadBalancer        string                             `json:"elastic_load_balancer"`
	CustomConfigureRecipes     []string                           `json:"custom_configure_recipes"`
	CustomShutdownRecipes      []string                           `json:"custom_shutdown_recipes"`
	StackId                    string                             `json:"stack_id"`
	Name                       string                             `json:"name"`
	CustomInstanceProfileArn   string                             `json:"custom_instance_profile_arn"`
	CustomSetupRecipes         []string                           `json:"custom_setup_recipes"`
	CustomDeployRecipes        []string                           `json:"custom_deploy_recipes"`
	CustomUndeployRecipes      []string                           `json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds     string                             `json:"custom_security_group_ids"`
	CustomJson                 string                             `json:"custom_json"`
	AutoAssignPublicIps        bool                               `json:"auto_assign_public_ips"`
	AutoHealing                bool                               `json:"auto_healing"`
	InstanceShutdownTimeout    int                                `json:"instance_shutdown_timeout"`
	DrainElbOnShutdown         bool                               `json:"drain_elb_on_shutdown"`
	EbsVolume                  AwsOpsworksMysqlLayerSpecEbsVolume `json:"ebs_volume"`
	RootPasswordOnAllInstances bool                               `json:"root_password_on_all_instances"`
	AutoAssignElasticIps       bool                               `json:"auto_assign_elastic_ips"`
	InstallUpdatesOnBoot       bool                               `json:"install_updates_on_boot"`
	SystemPackages             string                             `json:"system_packages"`
	UseEbsOptimizedInstances   bool                               `json:"use_ebs_optimized_instances"`
	RootPassword               string                             `json:"root_password"`
}

type AwsOpsworksMysqlLayerSpecEbsVolume struct {
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksMysqlLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksMysqlLayer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53ZoneAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRoute53ZoneAssociationSpec `json"spec"`
}

type AwsRoute53ZoneAssociationSpec struct {
	ZoneId    string `json:"zone_id"`
	VpcId     string `json:"vpc_id"`
	VpcRegion string `json:"vpc_region"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53ZoneAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRoute53ZoneAssociation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainIdentityVerification struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesDomainIdentityVerificationSpec `json"spec"`
}

type AwsSesDomainIdentityVerificationSpec struct {
	Arn    string `json:"arn"`
	Domain string `json:"domain"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainIdentityVerificationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesDomainIdentityVerification `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesTemplate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesTemplateSpec `json"spec"`
}

type AwsSesTemplateSpec struct {
	Name    string `json:"name"`
	Html    string `json:"html"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesTemplateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesTemplate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppCookieStickinessPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAppCookieStickinessPolicySpec `json"spec"`
}

type AwsAppCookieStickinessPolicySpec struct {
	Name         string `json:"name"`
	LoadBalancer string `json:"load_balancer"`
	LbPort       int    `json:"lb_port"`
	CookieName   string `json:"cookie_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppCookieStickinessPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAppCookieStickinessPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsClusterInstance struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRdsClusterInstanceSpec `json"spec"`
}

type AwsRdsClusterInstanceSpec struct {
	DbSubnetGroupName           string            `json:"db_subnet_group_name"`
	Endpoint                    string            `json:"endpoint"`
	PubliclyAccessible          bool              `json:"publicly_accessible"`
	MonitoringRoleArn           string            `json:"monitoring_role_arn"`
	PreferredMaintenanceWindow  string            `json:"preferred_maintenance_window"`
	PromotionTier               int               `json:"promotion_tier"`
	IdentifierPrefix            string            `json:"identifier_prefix"`
	AutoMinorVersionUpgrade     bool              `json:"auto_minor_version_upgrade"`
	PreferredBackupWindow       string            `json:"preferred_backup_window"`
	PerformanceInsightsEnabled  bool              `json:"performance_insights_enabled"`
	PerformanceInsightsKmsKeyId string            `json:"performance_insights_kms_key_id"`
	EngineVersion               string            `json:"engine_version"`
	Port                        int               `json:"port"`
	ApplyImmediately            bool              `json:"apply_immediately"`
	AvailabilityZone            string            `json:"availability_zone"`
	Writer                      bool              `json:"writer"`
	ClusterIdentifier           string            `json:"cluster_identifier"`
	InstanceClass               string            `json:"instance_class"`
	Engine                      string            `json:"engine"`
	DbParameterGroupName        string            `json:"db_parameter_group_name"`
	KmsKeyId                    string            `json:"kms_key_id"`
	StorageEncrypted            bool              `json:"storage_encrypted"`
	DbiResourceId               string            `json:"dbi_resource_id"`
	Identifier                  string            `json:"identifier"`
	Tags                        map[string]string `json:"tags"`
	MonitoringInterval          int               `json:"monitoring_interval"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsClusterInstanceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRdsClusterInstance `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultRouteTable struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDefaultRouteTableSpec `json"spec"`
}

type AwsDefaultRouteTableSpec struct {
	Tags                map[string]string             `json:"tags"`
	DefaultRouteTableId string                        `json:"default_route_table_id"`
	VpcId               string                        `json:"vpc_id"`
	PropagatingVgws     string                        `json:"propagating_vgws"`
	Route               AwsDefaultRouteTableSpecRoute `json:"route"`
}

type AwsDefaultRouteTableSpecRoute struct {
	Ipv6CidrBlock          string `json:"ipv6_cidr_block"`
	EgressOnlyGatewayId    string `json:"egress_only_gateway_id"`
	GatewayId              string `json:"gateway_id"`
	InstanceId             string `json:"instance_id"`
	NatGatewayId           string `json:"nat_gateway_id"`
	VpcPeeringConnectionId string `json:"vpc_peering_connection_id"`
	NetworkInterfaceId     string `json:"network_interface_id"`
	CidrBlock              string `json:"cidr_block"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultRouteTableList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDefaultRouteTable `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbTargetGroupAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLbTargetGroupAttachmentSpec `json"spec"`
}

type AwsLbTargetGroupAttachmentSpec struct {
	AvailabilityZone string `json:"availability_zone"`
	TargetGroupArn   string `json:"target_group_arn"`
	TargetId         string `json:"target_id"`
	Port             int    `json:"port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbTargetGroupAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLbTargetGroupAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchLogGroupSpec `json"spec"`
}

type AwsCloudwatchLogGroupSpec struct {
	Name            string            `json:"name"`
	NamePrefix      string            `json:"name_prefix"`
	RetentionInDays int               `json:"retention_in_days"`
	KmsKeyId        string            `json:"kms_key_id"`
	Arn             string            `json:"arn"`
	Tags            map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchLogGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbSecurityGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDbSecurityGroupSpec `json"spec"`
}

type AwsDbSecurityGroupSpec struct {
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Ingress     AwsDbSecurityGroupSpecIngress `json:"ingress"`
	Tags        map[string]string             `json:"tags"`
	Arn         string                        `json:"arn"`
}

type AwsDbSecurityGroupSpecIngress struct {
	Cidr                 string `json:"cidr"`
	SecurityGroupName    string `json:"security_group_name"`
	SecurityGroupId      string `json:"security_group_id"`
	SecurityGroupOwnerId string `json:"security_group_owner_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbSecurityGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDbSecurityGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElastictranscoderPipeline struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElastictranscoderPipelineSpec `json"spec"`
}

type AwsElastictranscoderPipelineSpec struct {
	Notifications              AwsElastictranscoderPipelineSpecNotifications              `json:"notifications"`
	ThumbnailConfig            AwsElastictranscoderPipelineSpecThumbnailConfig            `json:"thumbnail_config"`
	ThumbnailConfigPermissions AwsElastictranscoderPipelineSpecThumbnailConfigPermissions `json:"thumbnail_config_permissions"`
	Arn                        string                                                     `json:"arn"`
	ContentConfigPermissions   AwsElastictranscoderPipelineSpecContentConfigPermissions   `json:"content_config_permissions"`
	InputBucket                string                                                     `json:"input_bucket"`
	Name                       string                                                     `json:"name"`
	OutputBucket               string                                                     `json:"output_bucket"`
	Role                       string                                                     `json:"role"`
	AwsKmsKeyArn               string                                                     `json:"aws_kms_key_arn"`
	ContentConfig              AwsElastictranscoderPipelineSpecContentConfig              `json:"content_config"`
}

type AwsElastictranscoderPipelineSpecNotifications struct {
	Completed   string `json:"completed"`
	Error       string `json:"error"`
	Progressing string `json:"progressing"`
	Warning     string `json:"warning"`
}

type AwsElastictranscoderPipelineSpecThumbnailConfig struct {
	Bucket       string `json:"bucket"`
	StorageClass string `json:"storage_class"`
}

type AwsElastictranscoderPipelineSpecThumbnailConfigPermissions struct {
	Access      []string `json:"access"`
	Grantee     string   `json:"grantee"`
	GranteeType string   `json:"grantee_type"`
}

type AwsElastictranscoderPipelineSpecContentConfigPermissions struct {
	Access      []string `json:"access"`
	Grantee     string   `json:"grantee"`
	GranteeType string   `json:"grantee_type"`
}

type AwsElastictranscoderPipelineSpecContentConfig struct {
	Bucket       string `json:"bucket"`
	StorageClass string `json:"storage_class"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElastictranscoderPipelineList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElastictranscoderPipeline `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotFleetRequest struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSpotFleetRequestSpec `json"spec"`
}

type AwsSpotFleetRequestSpec struct {
	ReplaceUnhealthyInstances        bool                                       `json:"replace_unhealthy_instances"`
	TargetCapacity                   int                                        `json:"target_capacity"`
	SpotPrice                        string                                     `json:"spot_price"`
	TerminateInstancesWithExpiration bool                                       `json:"terminate_instances_with_expiration"`
	ValidFrom                        string                                     `json:"valid_from"`
	LaunchSpecification              AwsSpotFleetRequestSpecLaunchSpecification `json:"launch_specification"`
	AllocationStrategy               string                                     `json:"allocation_strategy"`
	ExcessCapacityTerminationPolicy  string                                     `json:"excess_capacity_termination_policy"`
	ValidUntil                       string                                     `json:"valid_until"`
	SpotRequestState                 string                                     `json:"spot_request_state"`
	WaitForFulfillment               bool                                       `json:"wait_for_fulfillment"`
	IamFleetRole                     string                                     `json:"iam_fleet_role"`
	InstanceInterruptionBehaviour    string                                     `json:"instance_interruption_behaviour"`
	ClientToken                      string                                     `json:"client_token"`
	LoadBalancers                    string                                     `json:"load_balancers"`
	TargetGroupArns                  string                                     `json:"target_group_arns"`
}

type AwsSpotFleetRequestSpecLaunchSpecification struct {
	Monitoring               bool                                                           `json:"monitoring"`
	UserData                 string                                                         `json:"user_data"`
	AvailabilityZone         string                                                         `json:"availability_zone"`
	Tags                     map[string]string                                              `json:"tags"`
	RootBlockDevice          AwsSpotFleetRequestSpecLaunchSpecificationRootBlockDevice      `json:"root_block_device"`
	Ami                      string                                                         `json:"ami"`
	InstanceType             string                                                         `json:"instance_type"`
	KeyName                  string                                                         `json:"key_name"`
	SpotPrice                string                                                         `json:"spot_price"`
	SubnetId                 string                                                         `json:"subnet_id"`
	AssociatePublicIpAddress bool                                                           `json:"associate_public_ip_address"`
	EbsBlockDevice           AwsSpotFleetRequestSpecLaunchSpecificationEbsBlockDevice       `json:"ebs_block_device"`
	EbsOptimized             bool                                                           `json:"ebs_optimized"`
	IamInstanceProfile       string                                                         `json:"iam_instance_profile"`
	PlacementTenancy         string                                                         `json:"placement_tenancy"`
	WeightedCapacity         string                                                         `json:"weighted_capacity"`
	VpcSecurityGroupIds      string                                                         `json:"vpc_security_group_ids"`
	EphemeralBlockDevice     AwsSpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice `json:"ephemeral_block_device"`
	PlacementGroup           string                                                         `json:"placement_group"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationEbsBlockDevice struct {
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
}

type AwsSpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotFleetRequestList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSpotFleetRequest `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAlbTargetGroupSpec `json"spec"`
}

type AwsAlbTargetGroupSpec struct {
	VpcId               string                             `json:"vpc_id"`
	DeregistrationDelay int                                `json:"deregistration_delay"`
	Tags                map[string]string                  `json:"tags"`
	ArnSuffix           string                             `json:"arn_suffix"`
	NamePrefix          string                             `json:"name_prefix"`
	Port                int                                `json:"port"`
	TargetType          string                             `json:"target_type"`
	Stickiness          []AwsAlbTargetGroupSpecStickiness  `json:"stickiness"`
	HealthCheck         []AwsAlbTargetGroupSpecHealthCheck `json:"health_check"`
	Arn                 string                             `json:"arn"`
	Name                string                             `json:"name"`
	Protocol            string                             `json:"protocol"`
}

type AwsAlbTargetGroupSpecStickiness struct {
	Type           string `json:"type"`
	CookieDuration int    `json:"cookie_duration"`
	Enabled        bool   `json:"enabled"`
}

type AwsAlbTargetGroupSpecHealthCheck struct {
	Path               string `json:"path"`
	Port               string `json:"port"`
	Protocol           string `json:"protocol"`
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Interval           int    `json:"interval"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAlbTargetGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontDistribution struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudfrontDistributionSpec `json"spec"`
}

type AwsCloudfrontDistributionSpec struct {
	DefaultCacheBehavior        AwsCloudfrontDistributionSpecDefaultCacheBehavior   `json:"default_cache_behavior"`
	Restrictions                AwsCloudfrontDistributionSpecRestrictions           `json:"restrictions"`
	HostedZoneId                string                                              `json:"hosted_zone_id"`
	CallerReference             string                                              `json:"caller_reference"`
	LastModifiedTime            string                                              `json:"last_modified_time"`
	Comment                     string                                              `json:"comment"`
	PriceClass                  string                                              `json:"price_class"`
	CustomErrorResponse         AwsCloudfrontDistributionSpecCustomErrorResponse    `json:"custom_error_response"`
	IsIpv6Enabled               bool                                                `json:"is_ipv6_enabled"`
	CacheBehavior               AwsCloudfrontDistributionSpecCacheBehavior          `json:"cache_behavior"`
	HttpVersion                 string                                              `json:"http_version"`
	LoggingConfig               AwsCloudfrontDistributionSpecLoggingConfig          `json:"logging_config"`
	ActiveTrustedSigners        map[string]string                                   `json:"active_trusted_signers"`
	DomainName                  string                                              `json:"domain_name"`
	Tags                        map[string]string                                   `json:"tags"`
	Origin                      AwsCloudfrontDistributionSpecOrigin                 `json:"origin"`
	ViewerCertificate           AwsCloudfrontDistributionSpecViewerCertificate      `json:"viewer_certificate"`
	InProgressValidationBatches int                                                 `json:"in_progress_validation_batches"`
	DefaultRootObject           string                                              `json:"default_root_object"`
	Enabled                     bool                                                `json:"enabled"`
	WebAclId                    string                                              `json:"web_acl_id"`
	Etag                        string                                              `json:"etag"`
	Arn                         string                                              `json:"arn"`
	Aliases                     string                                              `json:"aliases"`
	OrderedCacheBehavior        []AwsCloudfrontDistributionSpecOrderedCacheBehavior `json:"ordered_cache_behavior"`
	Status                      string                                              `json:"status"`
	RetainOnDelete              bool                                                `json:"retain_on_delete"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehavior struct {
	AllowedMethods            []string                                                                   `json:"allowed_methods"`
	LambdaFunctionAssociation AwsCloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation `json:"lambda_function_association"`
	MaxTtl                    int                                                                        `json:"max_ttl"`
	SmoothStreaming           bool                                                                       `json:"smooth_streaming"`
	TrustedSigners            []string                                                                   `json:"trusted_signers"`
	ViewerProtocolPolicy      string                                                                     `json:"viewer_protocol_policy"`
	CachedMethods             []string                                                                   `json:"cached_methods"`
	Compress                  bool                                                                       `json:"compress"`
	DefaultTtl                int                                                                        `json:"default_ttl"`
	FieldLevelEncryptionId    string                                                                     `json:"field_level_encryption_id"`
	ForwardedValues           AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues           `json:"forwarded_values"`
	MinTtl                    int                                                                        `json:"min_ttl"`
	TargetOriginId            string                                                                     `json:"target_origin_id"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation struct {
	EventType string `json:"event_type"`
	LambdaArn string `json:"lambda_arn"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues struct {
	QueryString          bool                                                                    `json:"query_string"`
	QueryStringCacheKeys []string                                                                `json:"query_string_cache_keys"`
	Cookies              AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies `json:"cookies"`
	Headers              []string                                                                `json:"headers"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type AwsCloudfrontDistributionSpecRestrictions struct {
	GeoRestriction AwsCloudfrontDistributionSpecRestrictionsGeoRestriction `json:"geo_restriction"`
}

type AwsCloudfrontDistributionSpecRestrictionsGeoRestriction struct {
	Locations       []string `json:"locations"`
	RestrictionType string   `json:"restriction_type"`
}

type AwsCloudfrontDistributionSpecCustomErrorResponse struct {
	ErrorCachingMinTtl int    `json:"error_caching_min_ttl"`
	ErrorCode          int    `json:"error_code"`
	ResponseCode       int    `json:"response_code"`
	ResponsePagePath   string `json:"response_page_path"`
}

type AwsCloudfrontDistributionSpecCacheBehavior struct {
	Compress                  bool                                                                `json:"compress"`
	FieldLevelEncryptionId    string                                                              `json:"field_level_encryption_id"`
	PathPattern               string                                                              `json:"path_pattern"`
	TargetOriginId            string                                                              `json:"target_origin_id"`
	ViewerProtocolPolicy      string                                                              `json:"viewer_protocol_policy"`
	AllowedMethods            []string                                                            `json:"allowed_methods"`
	ForwardedValues           AwsCloudfrontDistributionSpecCacheBehaviorForwardedValues           `json:"forwarded_values"`
	LambdaFunctionAssociation AwsCloudfrontDistributionSpecCacheBehaviorLambdaFunctionAssociation `json:"lambda_function_association"`
	MaxTtl                    int                                                                 `json:"max_ttl"`
	MinTtl                    int                                                                 `json:"min_ttl"`
	TrustedSigners            []string                                                            `json:"trusted_signers"`
	CachedMethods             []string                                                            `json:"cached_methods"`
	DefaultTtl                int                                                                 `json:"default_ttl"`
	SmoothStreaming           bool                                                                `json:"smooth_streaming"`
}

type AwsCloudfrontDistributionSpecCacheBehaviorForwardedValues struct {
	Cookies              AwsCloudfrontDistributionSpecCacheBehaviorForwardedValuesCookies `json:"cookies"`
	Headers              []string                                                         `json:"headers"`
	QueryString          bool                                                             `json:"query_string"`
	QueryStringCacheKeys []string                                                         `json:"query_string_cache_keys"`
}

type AwsCloudfrontDistributionSpecCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type AwsCloudfrontDistributionSpecCacheBehaviorLambdaFunctionAssociation struct {
	EventType string `json:"event_type"`
	LambdaArn string `json:"lambda_arn"`
}

type AwsCloudfrontDistributionSpecLoggingConfig struct {
	Bucket         string `json:"bucket"`
	IncludeCookies bool   `json:"include_cookies"`
	Prefix         string `json:"prefix"`
}

type AwsCloudfrontDistributionSpecOrigin struct {
	OriginPath         string                                                `json:"origin_path"`
	S3OriginConfig     AwsCloudfrontDistributionSpecOriginS3OriginConfig     `json:"s3_origin_config"`
	CustomOriginConfig AwsCloudfrontDistributionSpecOriginCustomOriginConfig `json:"custom_origin_config"`
	DomainName         string                                                `json:"domain_name"`
	CustomHeader       AwsCloudfrontDistributionSpecOriginCustomHeader       `json:"custom_header"`
	OriginId           string                                                `json:"origin_id"`
}

type AwsCloudfrontDistributionSpecOriginS3OriginConfig struct {
	OriginAccessIdentity string `json:"origin_access_identity"`
}

type AwsCloudfrontDistributionSpecOriginCustomOriginConfig struct {
	HttpPort               int      `json:"http_port"`
	HttpsPort              int      `json:"https_port"`
	OriginKeepaliveTimeout int      `json:"origin_keepalive_timeout"`
	OriginReadTimeout      int      `json:"origin_read_timeout"`
	OriginProtocolPolicy   string   `json:"origin_protocol_policy"`
	OriginSslProtocols     []string `json:"origin_ssl_protocols"`
}

type AwsCloudfrontDistributionSpecOriginCustomHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsCloudfrontDistributionSpecViewerCertificate struct {
	CloudfrontDefaultCertificate bool   `json:"cloudfront_default_certificate"`
	IamCertificateId             string `json:"iam_certificate_id"`
	MinimumProtocolVersion       string `json:"minimum_protocol_version"`
	SslSupportMethod             string `json:"ssl_support_method"`
	AcmCertificateArn            string `json:"acm_certificate_arn"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehavior struct {
	MinTtl                    int                                                                        `json:"min_ttl"`
	AllowedMethods            string                                                                     `json:"allowed_methods"`
	FieldLevelEncryptionId    string                                                                     `json:"field_level_encryption_id"`
	ForwardedValues           AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues           `json:"forwarded_values"`
	LambdaFunctionAssociation AwsCloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation `json:"lambda_function_association"`
	TargetOriginId            string                                                                     `json:"target_origin_id"`
	ViewerProtocolPolicy      string                                                                     `json:"viewer_protocol_policy"`
	CachedMethods             string                                                                     `json:"cached_methods"`
	PathPattern               string                                                                     `json:"path_pattern"`
	MaxTtl                    int                                                                        `json:"max_ttl"`
	DefaultTtl                int                                                                        `json:"default_ttl"`
	SmoothStreaming           bool                                                                       `json:"smooth_streaming"`
	TrustedSigners            []string                                                                   `json:"trusted_signers"`
	Compress                  bool                                                                       `json:"compress"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues struct {
	Cookies              AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies `json:"cookies"`
	Headers              []string                                                                `json:"headers"`
	QueryString          bool                                                                    `json:"query_string"`
	QueryStringCacheKeys []string                                                                `json:"query_string_cache_keys"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation struct {
	EventType string `json:"event_type"`
	LambdaArn string `json:"lambda_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontDistributionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudfrontDistribution `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcDhcpOptionsAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcDhcpOptionsAssociationSpec `json"spec"`
}

type AwsVpcDhcpOptionsAssociationSpec struct {
	DhcpOptionsId string `json:"dhcp_options_id"`
	VpcId         string `json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcDhcpOptionsAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcDhcpOptionsAssociation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointSubnetAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcEndpointSubnetAssociationSpec `json"spec"`
}

type AwsVpcEndpointSubnetAssociationSpec struct {
	VpcEndpointId string `json:"vpc_endpoint_id"`
	SubnetId      string `json:"subnet_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointSubnetAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcEndpointSubnetAssociation `json"items"`
}

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
	DeviceConfiguration         []AwsCognitoUserPoolSpecDeviceConfiguration         `json:"device_configuration"`
	LambdaConfig                []AwsCognitoUserPoolSpecLambdaConfig                `json:"lambda_config"`
	LastModifiedDate            string                                              `json:"last_modified_date"`
	MfaConfiguration            string                                              `json:"mfa_configuration"`
	PasswordPolicy              []AwsCognitoUserPoolSpecPasswordPolicy              `json:"password_policy"`
	VerificationMessageTemplate []AwsCognitoUserPoolSpecVerificationMessageTemplate `json:"verification_message_template"`
	AdminCreateUserConfig       []AwsCognitoUserPoolSpecAdminCreateUserConfig       `json:"admin_create_user_config"`
	Arn                         string                                              `json:"arn"`
	Tags                        map[string]string                                   `json:"tags"`
	AutoVerifiedAttributes      string                                              `json:"auto_verified_attributes"`
	EmailConfiguration          []AwsCognitoUserPoolSpecEmailConfiguration          `json:"email_configuration"`
	EmailVerificationSubject    string                                              `json:"email_verification_subject"`
	EmailVerificationMessage    string                                              `json:"email_verification_message"`
	Name                        string                                              `json:"name"`
	Schema                      AwsCognitoUserPoolSpecSchema                        `json:"schema"`
	SmsConfiguration            []AwsCognitoUserPoolSpecSmsConfiguration            `json:"sms_configuration"`
	SmsVerificationMessage      string                                              `json:"sms_verification_message"`
	AliasAttributes             string                                              `json:"alias_attributes"`
	SmsAuthenticationMessage    string                                              `json:"sms_authentication_message"`
	UsernameAttributes          []string                                            `json:"username_attributes"`
}

type AwsCognitoUserPoolSpecDeviceConfiguration struct {
	ChallengeRequiredOnNewDevice     bool `json:"challenge_required_on_new_device"`
	DeviceOnlyRememberedOnUserPrompt bool `json:"device_only_remembered_on_user_prompt"`
}

type AwsCognitoUserPoolSpecLambdaConfig struct {
	PostConfirmation            string `json:"post_confirmation"`
	PreTokenGeneration          string `json:"pre_token_generation"`
	CreateAuthChallenge         string `json:"create_auth_challenge"`
	CustomMessage               string `json:"custom_message"`
	DefineAuthChallenge         string `json:"define_auth_challenge"`
	PostAuthentication          string `json:"post_authentication"`
	PreAuthentication           string `json:"pre_authentication"`
	PreSignUp                   string `json:"pre_sign_up"`
	UserMigration               string `json:"user_migration"`
	VerifyAuthChallengeResponse string `json:"verify_auth_challenge_response"`
}

type AwsCognitoUserPoolSpecPasswordPolicy struct {
	RequireNumbers   bool `json:"require_numbers"`
	RequireSymbols   bool `json:"require_symbols"`
	RequireUppercase bool `json:"require_uppercase"`
	MinimumLength    int  `json:"minimum_length"`
	RequireLowercase bool `json:"require_lowercase"`
}

type AwsCognitoUserPoolSpecVerificationMessageTemplate struct {
	DefaultEmailOption string `json:"default_email_option"`
	EmailMessage       string `json:"email_message"`
	EmailMessageByLink string `json:"email_message_by_link"`
	EmailSubject       string `json:"email_subject"`
	EmailSubjectByLink string `json:"email_subject_by_link"`
	SmsMessage         string `json:"sms_message"`
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

type AwsCognitoUserPoolSpecEmailConfiguration struct {
	ReplyToEmailAddress string `json:"reply_to_email_address"`
	SourceArn           string `json:"source_arn"`
}

type AwsCognitoUserPoolSpecSchema struct {
	AttributeDataType          string                                                   `json:"attribute_data_type"`
	DeveloperOnlyAttribute     bool                                                     `json:"developer_only_attribute"`
	Mutable                    bool                                                     `json:"mutable"`
	Name                       string                                                   `json:"name"`
	NumberAttributeConstraints []AwsCognitoUserPoolSpecSchemaNumberAttributeConstraints `json:"number_attribute_constraints"`
	Required                   bool                                                     `json:"required"`
	StringAttributeConstraints []AwsCognitoUserPoolSpecSchemaStringAttributeConstraints `json:"string_attribute_constraints"`
}

type AwsCognitoUserPoolSpecSchemaNumberAttributeConstraints struct {
	MinValue string `json:"min_value"`
	MaxValue string `json:"max_value"`
}

type AwsCognitoUserPoolSpecSchemaStringAttributeConstraints struct {
	MaxLength string `json:"max_length"`
	MinLength string `json:"min_length"`
}

type AwsCognitoUserPoolSpecSmsConfiguration struct {
	ExternalId   string `json:"external_id"`
	SnsCallerArn string `json:"sns_caller_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCognitoUserPool `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchDashboard struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchDashboardSpec `json"spec"`
}

type AwsCloudwatchDashboardSpec struct {
	DashboardArn  string `json:"dashboard_arn"`
	DashboardBody string `json:"dashboard_body"`
	DashboardName string `json:"dashboard_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchDashboardList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchDashboard `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKinesisStream struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsKinesisStreamSpec `json"spec"`
}

type AwsKinesisStreamSpec struct {
	RetentionPeriod   int               `json:"retention_period"`
	ShardLevelMetrics string            `json:"shard_level_metrics"`
	EncryptionType    string            `json:"encryption_type"`
	KmsKeyId          string            `json:"kms_key_id"`
	Arn               string            `json:"arn"`
	Tags              map[string]string `json:"tags"`
	Name              string            `json:"name"`
	ShardCount        int               `json:"shard_count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKinesisStreamList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsKinesisStream `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMqBroker struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsMqBrokerSpec `json"spec"`
}

type AwsMqBrokerSpec struct {
	EngineVersion              string                                      `json:"engine_version"`
	ApplyImmediately           bool                                        `json:"apply_immediately"`
	Configuration              []AwsMqBrokerSpecConfiguration              `json:"configuration"`
	EngineType                 string                                      `json:"engine_type"`
	PubliclyAccessible         bool                                        `json:"publicly_accessible"`
	AutoMinorVersionUpgrade    bool                                        `json:"auto_minor_version_upgrade"`
	DeploymentMode             string                                      `json:"deployment_mode"`
	MaintenanceWindowStartTime []AwsMqBrokerSpecMaintenanceWindowStartTime `json:"maintenance_window_start_time"`
	SecurityGroups             string                                      `json:"security_groups"`
	User                       AwsMqBrokerSpecUser                         `json:"user"`
	Instances                  []AwsMqBrokerSpecInstances                  `json:"instances"`
	BrokerName                 string                                      `json:"broker_name"`
	HostInstanceType           string                                      `json:"host_instance_type"`
	SubnetIds                  string                                      `json:"subnet_ids"`
	Arn                        string                                      `json:"arn"`
}

type AwsMqBrokerSpecConfiguration struct {
	Id       string `json:"id"`
	Revision int    `json:"revision"`
}

type AwsMqBrokerSpecMaintenanceWindowStartTime struct {
	TimeOfDay string `json:"time_of_day"`
	TimeZone  string `json:"time_zone"`
	DayOfWeek string `json:"day_of_week"`
}

type AwsMqBrokerSpecUser struct {
	Username      string `json:"username"`
	ConsoleAccess bool   `json:"console_access"`
	Groups        string `json:"groups"`
	Password      string `json:"password"`
}

type AwsMqBrokerSpecInstances struct {
	ConsoleUrl string   `json:"console_url"`
	Endpoints  []string `json:"endpoints"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMqBrokerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsMqBroker `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStack struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksStackSpec `json"spec"`
}

type AwsOpsworksStackSpec struct {
	ConfigurationManagerVersion string                                      `json:"configuration_manager_version"`
	CustomCookbooksSource       []AwsOpsworksStackSpecCustomCookbooksSource `json:"custom_cookbooks_source"`
	DefaultAvailabilityZone     string                                      `json:"default_availability_zone"`
	DefaultSshKeyName           string                                      `json:"default_ssh_key_name"`
	Tags                        map[string]string                           `json:"tags"`
	UseCustomCookbooks          bool                                        `json:"use_custom_cookbooks"`
	AgentVersion                string                                      `json:"agent_version"`
	DefaultInstanceProfileArn   string                                      `json:"default_instance_profile_arn"`
	ConfigurationManagerName    string                                      `json:"configuration_manager_name"`
	ManageBerkshelf             bool                                        `json:"manage_berkshelf"`
	DefaultRootDeviceType       string                                      `json:"default_root_device_type"`
	UseOpsworksSecurityGroups   bool                                        `json:"use_opsworks_security_groups"`
	Arn                         string                                      `json:"arn"`
	Region                      string                                      `json:"region"`
	DefaultOs                   string                                      `json:"default_os"`
	HostnameTheme               string                                      `json:"hostname_theme"`
	Name                        string                                      `json:"name"`
	Color                       string                                      `json:"color"`
	CustomJson                  string                                      `json:"custom_json"`
	DefaultSubnetId             string                                      `json:"default_subnet_id"`
	VpcId                       string                                      `json:"vpc_id"`
	StackEndpoint               string                                      `json:"stack_endpoint"`
	ServiceRoleArn              string                                      `json:"service_role_arn"`
	BerkshelfVersion            string                                      `json:"berkshelf_version"`
}

type AwsOpsworksStackSpecCustomCookbooksSource struct {
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStackList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksStack `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftSubnetGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRedshiftSubnetGroupSpec `json"spec"`
}

type AwsRedshiftSubnetGroupSpec struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	SubnetIds   string            `json:"subnet_ids"`
	Tags        map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftSubnetGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRedshiftSubnetGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethodResponse struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayMethodResponseSpec `json"spec"`
}

type AwsApiGatewayMethodResponseSpec struct {
	RestApiId                string            `json:"rest_api_id"`
	ResourceId               string            `json:"resource_id"`
	HttpMethod               string            `json:"http_method"`
	StatusCode               string            `json:"status_code"`
	ResponseModels           map[string]string `json:"response_models"`
	ResponseParameters       map[string]string `json:"response_parameters"`
	ResponseParametersInJson string            `json:"response_parameters_in_json"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethodResponseList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayMethodResponse `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsCertificate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDmsCertificateSpec `json"spec"`
}

type AwsDmsCertificateSpec struct {
	CertificatePem    string `json:"certificate_pem"`
	CertificateWallet string `json:"certificate_wallet"`
	CertificateArn    string `json:"certificate_arn"`
	CertificateId     string `json:"certificate_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsCertificateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDmsCertificate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbGlobalTable struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDynamodbGlobalTableSpec `json"spec"`
}

type AwsDynamodbGlobalTableSpec struct {
	Name    string                            `json:"name"`
	Replica AwsDynamodbGlobalTableSpecReplica `json:"replica"`
	Arn     string                            `json:"arn"`
}

type AwsDynamodbGlobalTableSpecReplica struct {
	RegionName string `json:"region_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbGlobalTableList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDynamodbGlobalTable `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketObject struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsS3BucketObjectSpec `json"spec"`
}

type AwsS3BucketObjectSpec struct {
	ContentBase64        string            `json:"content_base64"`
	StorageClass         string            `json:"storage_class"`
	Etag                 string            `json:"etag"`
	Bucket               string            `json:"bucket"`
	ContentDisposition   string            `json:"content_disposition"`
	ContentType          string            `json:"content_type"`
	WebsiteRedirect      string            `json:"website_redirect"`
	ContentEncoding      string            `json:"content_encoding"`
	Key                  string            `json:"key"`
	Tags                 map[string]string `json:"tags"`
	Acl                  string            `json:"acl"`
	ContentLanguage      string            `json:"content_language"`
	ServerSideEncryption string            `json:"server_side_encryption"`
	KmsKeyId             string            `json:"kms_key_id"`
	VersionId            string            `json:"version_id"`
	CacheControl         string            `json:"cache_control"`
	Source               string            `json:"source"`
	Content              string            `json:"content"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketObjectList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsS3BucketObject `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchJobDefinition struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsBatchJobDefinitionSpec `json"spec"`
}

type AwsBatchJobDefinitionSpec struct {
	RetryStrategy       []AwsBatchJobDefinitionSpecRetryStrategy `json:"retry_strategy"`
	Timeout             []AwsBatchJobDefinitionSpecTimeout       `json:"timeout"`
	Type                string                                   `json:"type"`
	Revision            int                                      `json:"revision"`
	Arn                 string                                   `json:"arn"`
	Name                string                                   `json:"name"`
	ContainerProperties string                                   `json:"container_properties"`
	Parameters          map[string]string                        `json:"parameters"`
}

type AwsBatchJobDefinitionSpecRetryStrategy struct {
	Attempts int `json:"attempts"`
}

type AwsBatchJobDefinitionSpecTimeout struct {
	AttemptDurationSeconds int `json:"attempt_duration_seconds"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchJobDefinitionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsBatchJobDefinition `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailInstance struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLightsailInstanceSpec `json"spec"`
}

type AwsLightsailInstanceSpec struct {
	Username         string `json:"username"`
	Name             string `json:"name"`
	UserData         string `json:"user_data"`
	RamSize          int    `json:"ram_size"`
	Ipv6Address      string `json:"ipv6_address"`
	AvailabilityZone string `json:"availability_zone"`
	KeyPairName      string `json:"key_pair_name"`
	PrivateIpAddress string `json:"private_ip_address"`
	BundleId         string `json:"bundle_id"`
	CreatedAt        string `json:"created_at"`
	CpuCount         int    `json:"cpu_count"`
	BlueprintId      string `json:"blueprint_id"`
	Arn              string `json:"arn"`
	IsStaticIp       bool   `json:"is_static_ip"`
	PublicIpAddress  string `json:"public_ip_address"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailInstanceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLightsailInstance `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailKeyPair struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLightsailKeyPairSpec `json"spec"`
}

type AwsLightsailKeyPairSpec struct {
	PgpKey               string `json:"pgp_key"`
	Fingerprint          string `json:"fingerprint"`
	PublicKey            string `json:"public_key"`
	PrivateKey           string `json:"private_key"`
	Name                 string `json:"name"`
	NamePrefix           string `json:"name_prefix"`
	Arn                  string `json:"arn"`
	EncryptedFingerprint string `json:"encrypted_fingerprint"`
	EncryptedPrivateKey  string `json:"encrypted_private_key"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailKeyPairList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLightsailKeyPair `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingNotification struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAutoscalingNotificationSpec `json"spec"`
}

type AwsAutoscalingNotificationSpec struct {
	TopicArn      string `json:"topic_arn"`
	GroupNames    string `json:"group_names"`
	Notifications string `json:"notifications"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingNotificationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAutoscalingNotification `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventPermission struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchEventPermissionSpec `json"spec"`
}

type AwsCloudwatchEventPermissionSpec struct {
	Principal   string `json:"principal"`
	StatementId string `json:"statement_id"`
	Action      string `json:"action"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventPermissionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchEventPermission `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogDestination struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchLogDestinationSpec `json"spec"`
}

type AwsCloudwatchLogDestinationSpec struct {
	Name      string `json:"name"`
	RoleArn   string `json:"role_arn"`
	TargetArn string `json:"target_arn"`
	Arn       string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogDestinationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchLogDestination `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoIdentityPool struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCognitoIdentityPoolSpec `json"spec"`
}

type AwsCognitoIdentityPoolSpec struct {
	IdentityPoolName               string                                             `json:"identity_pool_name"`
	CognitoIdentityProviders       AwsCognitoIdentityPoolSpecCognitoIdentityProviders `json:"cognito_identity_providers"`
	DeveloperProviderName          string                                             `json:"developer_provider_name"`
	AllowUnauthenticatedIdentities bool                                               `json:"allow_unauthenticated_identities"`
	OpenidConnectProviderArns      []string                                           `json:"openid_connect_provider_arns"`
	SamlProviderArns               []string                                           `json:"saml_provider_arns"`
	SupportedLoginProviders        map[string]string                                  `json:"supported_login_providers"`
}

type AwsCognitoIdentityPoolSpecCognitoIdentityProviders struct {
	ClientId             string `json:"client_id"`
	ProviderName         string `json:"provider_name"`
	ServerSideTokenCheck bool   `json:"server_side_token_check"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoIdentityPoolList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCognitoIdentityPool `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDaxSubnetGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDaxSubnetGroupSpec `json"spec"`
}

type AwsDaxSubnetGroupSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SubnetIds   string `json:"subnet_ids"`
	VpcId       string `json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDaxSubnetGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDaxSubnetGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailDomain struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLightsailDomainSpec `json"spec"`
}

type AwsLightsailDomainSpec struct {
	DomainName string `json:"domain_name"`
	Arn        string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailDomainList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLightsailDomain `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3Bucket struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsS3BucketSpec `json"spec"`
}

type AwsS3BucketSpec struct {
	Arn                               string                                             `json:"arn"`
	Acl                               string                                             `json:"acl"`
	Versioning                        []AwsS3BucketSpecVersioning                        `json:"versioning"`
	AccelerationStatus                string                                             `json:"acceleration_status"`
	ForceDestroy                      bool                                               `json:"force_destroy"`
	Bucket                            string                                             `json:"bucket"`
	BucketPrefix                      string                                             `json:"bucket_prefix"`
	Region                            string                                             `json:"region"`
	WebsiteEndpoint                   string                                             `json:"website_endpoint"`
	Tags                              map[string]string                                  `json:"tags"`
	Website                           []AwsS3BucketSpecWebsite                           `json:"website"`
	HostedZoneId                      string                                             `json:"hosted_zone_id"`
	Logging                           AwsS3BucketSpecLogging                             `json:"logging"`
	ReplicationConfiguration          []AwsS3BucketSpecReplicationConfiguration          `json:"replication_configuration"`
	LifecycleRule                     []AwsS3BucketSpecLifecycleRule                     `json:"lifecycle_rule"`
	RequestPayer                      string                                             `json:"request_payer"`
	ServerSideEncryptionConfiguration []AwsS3BucketSpecServerSideEncryptionConfiguration `json:"server_side_encryption_configuration"`
	BucketDomainName                  string                                             `json:"bucket_domain_name"`
	Policy                            string                                             `json:"policy"`
	CorsRule                          []AwsS3BucketSpecCorsRule                          `json:"cors_rule"`
	WebsiteDomain                     string                                             `json:"website_domain"`
}

type AwsS3BucketSpecVersioning struct {
	Enabled   bool `json:"enabled"`
	MfaDelete bool `json:"mfa_delete"`
}

type AwsS3BucketSpecWebsite struct {
	IndexDocument         string `json:"index_document"`
	ErrorDocument         string `json:"error_document"`
	RedirectAllRequestsTo string `json:"redirect_all_requests_to"`
	RoutingRules          string `json:"routing_rules"`
}

type AwsS3BucketSpecLogging struct {
	TargetBucket string `json:"target_bucket"`
	TargetPrefix string `json:"target_prefix"`
}

type AwsS3BucketSpecReplicationConfiguration struct {
	Role  string                                       `json:"role"`
	Rules AwsS3BucketSpecReplicationConfigurationRules `json:"rules"`
}

type AwsS3BucketSpecReplicationConfigurationRules struct {
	SourceSelectionCriteria AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria `json:"source_selection_criteria"`
	Prefix                  string                                                              `json:"prefix"`
	Status                  string                                                              `json:"status"`
	Id                      string                                                              `json:"id"`
	Destination             AwsS3BucketSpecReplicationConfigurationRulesDestination             `json:"destination"`
}

type AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria struct {
	SseKmsEncryptedObjects AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects `json:"sse_kms_encrypted_objects"`
}

type AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	Enabled bool `json:"enabled"`
}

type AwsS3BucketSpecReplicationConfigurationRulesDestination struct {
	Bucket          string `json:"bucket"`
	StorageClass    string `json:"storage_class"`
	ReplicaKmsKeyId string `json:"replica_kms_key_id"`
}

type AwsS3BucketSpecLifecycleRule struct {
	Id                                 string                                                  `json:"id"`
	Transition                         AwsS3BucketSpecLifecycleRuleTransition                  `json:"transition"`
	NoncurrentVersionTransition        AwsS3BucketSpecLifecycleRuleNoncurrentVersionTransition `json:"noncurrent_version_transition"`
	AbortIncompleteMultipartUploadDays int                                                     `json:"abort_incomplete_multipart_upload_days"`
	Expiration                         AwsS3BucketSpecLifecycleRuleExpiration                  `json:"expiration"`
	NoncurrentVersionExpiration        AwsS3BucketSpecLifecycleRuleNoncurrentVersionExpiration `json:"noncurrent_version_expiration"`
	Prefix                             string                                                  `json:"prefix"`
	Tags                               map[string]string                                       `json:"tags"`
	Enabled                            bool                                                    `json:"enabled"`
}

type AwsS3BucketSpecLifecycleRuleTransition struct {
	Date         string `json:"date"`
	Days         int    `json:"days"`
	StorageClass string `json:"storage_class"`
}

type AwsS3BucketSpecLifecycleRuleNoncurrentVersionTransition struct {
	Days         int    `json:"days"`
	StorageClass string `json:"storage_class"`
}

type AwsS3BucketSpecLifecycleRuleExpiration struct {
	Days                      int    `json:"days"`
	ExpiredObjectDeleteMarker bool   `json:"expired_object_delete_marker"`
	Date                      string `json:"date"`
}

type AwsS3BucketSpecLifecycleRuleNoncurrentVersionExpiration struct {
	Days int `json:"days"`
}

type AwsS3BucketSpecServerSideEncryptionConfiguration struct {
	Rule []AwsS3BucketSpecServerSideEncryptionConfigurationRule `json:"rule"`
}

type AwsS3BucketSpecServerSideEncryptionConfigurationRule struct {
	ApplyServerSideEncryptionByDefault []AwsS3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault `json:"apply_server_side_encryption_by_default"`
}

type AwsS3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault struct {
	SseAlgorithm   string `json:"sse_algorithm"`
	KmsMasterKeyId string `json:"kms_master_key_id"`
}

type AwsS3BucketSpecCorsRule struct {
	MaxAgeSeconds  int      `json:"max_age_seconds"`
	AllowedHeaders []string `json:"allowed_headers"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedOrigins []string `json:"allowed_origins"`
	ExposeHeaders  []string `json:"expose_headers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsS3Bucket `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsS3BucketPolicySpec `json"spec"`
}

type AwsS3BucketPolicySpec struct {
	Bucket string `json:"bucket"`
	Policy string `json:"policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsS3BucketPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmActivation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmActivationSpec `json"spec"`
}

type AwsSsmActivationSpec struct {
	Expired           string `json:"expired"`
	ExpirationDate    string `json:"expiration_date"`
	IamRole           string `json:"iam_role"`
	RegistrationLimit int    `json:"registration_limit"`
	RegistrationCount int    `json:"registration_count"`
	ActivationCode    string `json:"activation_code"`
	Name              string `json:"name"`
	Description       string `json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmActivationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmActivation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowTask struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmMaintenanceWindowTaskSpec `json"spec"`
}

type AwsSsmMaintenanceWindowTaskSpec struct {
	TaskParameters []AwsSsmMaintenanceWindowTaskSpecTaskParameters `json:"task_parameters"`
	MaxConcurrency string                                          `json:"max_concurrency"`
	MaxErrors      string                                          `json:"max_errors"`
	TaskArn        string                                          `json:"task_arn"`
	Targets        []AwsSsmMaintenanceWindowTaskSpecTargets        `json:"targets"`
	Priority       int                                             `json:"priority"`
	WindowId       string                                          `json:"window_id"`
	TaskType       string                                          `json:"task_type"`
	ServiceRoleArn string                                          `json:"service_role_arn"`
	LoggingInfo    []AwsSsmMaintenanceWindowTaskSpecLoggingInfo    `json:"logging_info"`
}

type AwsSsmMaintenanceWindowTaskSpecTaskParameters struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type AwsSsmMaintenanceWindowTaskSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type AwsSsmMaintenanceWindowTaskSpecLoggingInfo struct {
	S3Region       string `json:"s3_region"`
	S3BucketPrefix string `json:"s3_bucket_prefix"`
	S3BucketName   string `json:"s3_bucket_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowTaskList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmMaintenanceWindowTask `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDocumentationVersion struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayDocumentationVersionSpec `json"spec"`
}

type AwsApiGatewayDocumentationVersionSpec struct {
	Version     string `json:"version"`
	RestApiId   string `json:"rest_api_id"`
	Description string `json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDocumentationVersionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayDocumentationVersion `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbInstance struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDbInstanceSpec `json"spec"`
}

type AwsDbInstanceSpec struct {
	MonitoringRoleArn                string                      `json:"monitoring_role_arn"`
	IamDatabaseAuthenticationEnabled bool                        `json:"iam_database_authentication_enabled"`
	Password                         string                      `json:"password"`
	Engine                           string                      `json:"engine"`
	PubliclyAccessible               bool                        `json:"publicly_accessible"`
	S3Import                         []AwsDbInstanceSpecS3Import `json:"s3_import"`
	ParameterGroupName               string                      `json:"parameter_group_name"`
	HostedZoneId                     string                      `json:"hosted_zone_id"`
	EnabledCloudwatchLogsExports     []string                    `json:"enabled_cloudwatch_logs_exports"`
	Port                             int                         `json:"port"`
	SecurityGroupNames               string                      `json:"security_group_names"`
	FinalSnapshotIdentifier          string                      `json:"final_snapshot_identifier"`
	CopyTagsToSnapshot               bool                        `json:"copy_tags_to_snapshot"`
	KmsKeyId                         string                      `json:"kms_key_id"`
	StorageType                      string                      `json:"storage_type"`
	BackupRetentionPeriod            int                         `json:"backup_retention_period"`
	SkipFinalSnapshot                bool                        `json:"skip_final_snapshot"`
	OptionGroupName                  string                      `json:"option_group_name"`
	ResourceId                       string                      `json:"resource_id"`
	Status                           string                      `json:"status"`
	MonitoringInterval               int                         `json:"monitoring_interval"`
	Arn                              string                      `json:"arn"`
	CharacterSetName                 string                      `json:"character_set_name"`
	StorageEncrypted                 bool                        `json:"storage_encrypted"`
	Identifier                       string                      `json:"identifier"`
	BackupWindow                     string                      `json:"backup_window"`
	VpcSecurityGroupIds              string                      `json:"vpc_security_group_ids"`
	Tags                             map[string]string           `json:"tags"`
	AllocatedStorage                 int                         `json:"allocated_storage"`
	ReplicateSourceDb                string                      `json:"replicate_source_db"`
	AutoMinorVersionUpgrade          bool                        `json:"auto_minor_version_upgrade"`
	Timezone                         string                      `json:"timezone"`
	CaCertIdentifier                 string                      `json:"ca_cert_identifier"`
	Username                         string                      `json:"username"`
	IdentifierPrefix                 string                      `json:"identifier_prefix"`
	LicenseModel                     string                      `json:"license_model"`
	DbSubnetGroupName                string                      `json:"db_subnet_group_name"`
	Endpoint                         string                      `json:"endpoint"`
	Replicas                         []string                    `json:"replicas"`
	EngineVersion                    string                      `json:"engine_version"`
	AvailabilityZone                 string                      `json:"availability_zone"`
	MaintenanceWindow                string                      `json:"maintenance_window"`
	ApplyImmediately                 bool                        `json:"apply_immediately"`
	AllowMajorVersionUpgrade         bool                        `json:"allow_major_version_upgrade"`
	Name                             string                      `json:"name"`
	InstanceClass                    string                      `json:"instance_class"`
	Iops                             int                         `json:"iops"`
	MultiAz                          bool                        `json:"multi_az"`
	Address                          string                      `json:"address"`
	SnapshotIdentifier               string                      `json:"snapshot_identifier"`
}

type AwsDbInstanceSpecS3Import struct {
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbInstanceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDbInstance `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcrRepository struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEcrRepositorySpec `json"spec"`
}

type AwsEcrRepositorySpec struct {
	Name          string `json:"name"`
	Arn           string `json:"arn"`
	RegistryId    string `json:"registry_id"`
	RepositoryUrl string `json:"repository_url"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcrRepositoryList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEcrRepository `json"items"`
}

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

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamUserPolicySpec `json"spec"`
}

type AwsIamUserPolicySpec struct {
	NamePrefix string `json:"name_prefix"`
	User       string `json:"user"`
	Policy     string `json:"policy"`
	Name       string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamUserPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerListenerPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLoadBalancerListenerPolicySpec `json"spec"`
}

type AwsLoadBalancerListenerPolicySpec struct {
	LoadBalancerPort int    `json:"load_balancer_port"`
	LoadBalancerName string `json:"load_balancer_name"`
	PolicyNames      string `json:"policy_names"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerListenerPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLoadBalancerListenerPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodedeployApp struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCodedeployAppSpec `json"spec"`
}

type AwsCodedeployAppSpec struct {
	UniqueId string `json:"unique_id"`
	Name     string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodedeployAppList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCodedeployApp `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamSamlProvider struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamSamlProviderSpec `json"spec"`
}

type AwsIamSamlProviderSpec struct {
	SamlMetadataDocument string `json:"saml_metadata_document"`
	Arn                  string `json:"arn"`
	ValidUntil           string `json:"valid_until"`
	Name                 string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamSamlProviderList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamSamlProvider `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcPeeringConnectionOptions struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcPeeringConnectionOptionsSpec `json"spec"`
}

type AwsVpcPeeringConnectionOptionsSpec struct {
	VpcPeeringConnectionId string                                      `json:"vpc_peering_connection_id"`
	Accepter               AwsVpcPeeringConnectionOptionsSpecAccepter  `json:"accepter"`
	Requester              AwsVpcPeeringConnectionOptionsSpecRequester `json:"requester"`
}

type AwsVpcPeeringConnectionOptionsSpecAccepter struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionOptionsSpecRequester struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcPeeringConnectionOptionsList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcPeeringConnectionOptions `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayResource struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayResourceSpec `json"spec"`
}

type AwsApiGatewayResourceSpec struct {
	RestApiId string `json:"rest_api_id"`
	ParentId  string `json:"parent_id"`
	PathPart  string `json:"path_part"`
	Path      string `json:"path"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayResourceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayResource `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodepipeline struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCodepipelineSpec `json"spec"`
}

type AwsCodepipelineSpec struct {
	Name          string                             `json:"name"`
	RoleArn       string                             `json:"role_arn"`
	ArtifactStore []AwsCodepipelineSpecArtifactStore `json:"artifact_store"`
	Stage         []AwsCodepipelineSpecStage         `json:"stage"`
	Arn           string                             `json:"arn"`
}

type AwsCodepipelineSpecArtifactStore struct {
	Location      string                                          `json:"location"`
	Type          string                                          `json:"type"`
	EncryptionKey []AwsCodepipelineSpecArtifactStoreEncryptionKey `json:"encryption_key"`
}

type AwsCodepipelineSpecArtifactStoreEncryptionKey struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type AwsCodepipelineSpecStage struct {
	Action []AwsCodepipelineSpecStageAction `json:"action"`
	Name   string                           `json:"name"`
}

type AwsCodepipelineSpecStageAction struct {
	Name            string            `json:"name"`
	RoleArn         string            `json:"role_arn"`
	RunOrder        int               `json:"run_order"`
	Configuration   map[string]string `json:"configuration"`
	Category        string            `json:"category"`
	Owner           string            `json:"owner"`
	Version         string            `json:"version"`
	Provider        string            `json:"provider"`
	InputArtifacts  []string          `json:"input_artifacts"`
	OutputArtifacts []string          `json:"output_artifacts"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodepipelineList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCodepipeline `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsFileSystem struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEfsFileSystemSpec `json"spec"`
}

type AwsEfsFileSystemSpec struct {
	CreationToken   string            `json:"creation_token"`
	ReferenceName   string            `json:"reference_name"`
	PerformanceMode string            `json:"performance_mode"`
	Encrypted       bool              `json:"encrypted"`
	KmsKeyId        string            `json:"kms_key_id"`
	DnsName         string            `json:"dns_name"`
	Tags            map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsFileSystemList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEfsFileSystem `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLaunchConfiguration struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLaunchConfigurationSpec `json"spec"`
}

type AwsLaunchConfigurationSpec struct {
	NamePrefix                   string                                         `json:"name_prefix"`
	IamInstanceProfile           string                                         `json:"iam_instance_profile"`
	VpcClassicLinkId             string                                         `json:"vpc_classic_link_id"`
	EphemeralBlockDevice         AwsLaunchConfigurationSpecEphemeralBlockDevice `json:"ephemeral_block_device"`
	RootBlockDevice              []AwsLaunchConfigurationSpecRootBlockDevice    `json:"root_block_device"`
	VpcClassicLinkSecurityGroups string                                         `json:"vpc_classic_link_security_groups"`
	PlacementTenancy             string                                         `json:"placement_tenancy"`
	Name                         string                                         `json:"name"`
	ImageId                      string                                         `json:"image_id"`
	InstanceType                 string                                         `json:"instance_type"`
	KeyName                      string                                         `json:"key_name"`
	SecurityGroups               string                                         `json:"security_groups"`
	EbsBlockDevice               AwsLaunchConfigurationSpecEbsBlockDevice       `json:"ebs_block_device"`
	UserData                     string                                         `json:"user_data"`
	AssociatePublicIpAddress     bool                                           `json:"associate_public_ip_address"`
	SpotPrice                    string                                         `json:"spot_price"`
	EbsOptimized                 bool                                           `json:"ebs_optimized"`
	EnableMonitoring             bool                                           `json:"enable_monitoring"`
	UserDataBase64               string                                         `json:"user_data_base64"`
}

type AwsLaunchConfigurationSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsLaunchConfigurationSpecRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsLaunchConfigurationSpecEbsBlockDevice struct {
	DeviceName          string `json:"device_name"`
	NoDevice            bool   `json:"no_device"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	Encrypted           bool   `json:"encrypted"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLaunchConfigurationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLaunchConfiguration `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsTopicSubscription struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSnsTopicSubscriptionSpec `json"spec"`
}

type AwsSnsTopicSubscriptionSpec struct {
	EndpointAutoConfirms         bool   `json:"endpoint_auto_confirms"`
	TopicArn                     string `json:"topic_arn"`
	RawMessageDelivery           bool   `json:"raw_message_delivery"`
	FilterPolicy                 string `json:"filter_policy"`
	Protocol                     string `json:"protocol"`
	Endpoint                     string `json:"endpoint"`
	ConfirmationTimeoutInMinutes int    `json:"confirmation_timeout_in_minutes"`
	DeliveryPolicy               string `json:"delivery_policy"`
	Arn                          string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsTopicSubscriptionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSnsTopicSubscription `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcPeeringConnection struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcPeeringConnectionSpec `json"spec"`
}

type AwsVpcPeeringConnectionSpec struct {
	Tags         map[string]string                    `json:"tags"`
	PeerVpcId    string                               `json:"peer_vpc_id"`
	AutoAccept   bool                                 `json:"auto_accept"`
	AcceptStatus string                               `json:"accept_status"`
	PeerRegion   string                               `json:"peer_region"`
	Accepter     AwsVpcPeeringConnectionSpecAccepter  `json:"accepter"`
	PeerOwnerId  string                               `json:"peer_owner_id"`
	VpcId        string                               `json:"vpc_id"`
	Requester    AwsVpcPeeringConnectionSpecRequester `json:"requester"`
}

type AwsVpcPeeringConnectionSpecAccepter struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionSpecRequester struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcPeeringConnectionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcPeeringConnection `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDirectoryServiceDirectory struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDirectoryServiceDirectorySpec `json"spec"`
}

type AwsDirectoryServiceDirectorySpec struct {
	ConnectSettings []AwsDirectoryServiceDirectorySpecConnectSettings `json:"connect_settings"`
	Edition         string                                            `json:"edition"`
	ShortName       string                                            `json:"short_name"`
	Tags            map[string]string                                 `json:"tags"`
	AccessUrl       string                                            `json:"access_url"`
	SecurityGroupId string                                            `json:"security_group_id"`
	Name            string                                            `json:"name"`
	Alias           string                                            `json:"alias"`
	Description     string                                            `json:"description"`
	VpcSettings     []AwsDirectoryServiceDirectorySpecVpcSettings     `json:"vpc_settings"`
	EnableSso       bool                                              `json:"enable_sso"`
	Type            string                                            `json:"type"`
	Password        string                                            `json:"password"`
	Size            string                                            `json:"size"`
	DnsIpAddresses  string                                            `json:"dns_ip_addresses"`
}

type AwsDirectoryServiceDirectorySpecConnectSettings struct {
	CustomerUsername string `json:"customer_username"`
	CustomerDnsIps   string `json:"customer_dns_ips"`
	SubnetIds        string `json:"subnet_ids"`
	VpcId            string `json:"vpc_id"`
}

type AwsDirectoryServiceDirectorySpecVpcSettings struct {
	VpcId     string `json:"vpc_id"`
	SubnetIds string `json:"subnet_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDirectoryServiceDirectoryList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDirectoryServiceDirectory `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElbAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElbAttachmentSpec `json"spec"`
}

type AwsElbAttachmentSpec struct {
	Elb      string `json:"elb"`
	Instance string `json:"instance"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElbAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElbAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsPlacementGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsPlacementGroupSpec `json"spec"`
}

type AwsPlacementGroupSpec struct {
	Name     string `json:"name"`
	Strategy string `json:"strategy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsPlacementGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsPlacementGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryPrivateDnsNamespace struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsServiceDiscoveryPrivateDnsNamespaceSpec `json"spec"`
}

type AwsServiceDiscoveryPrivateDnsNamespaceSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Vpc         string `json:"vpc"`
	Arn         string `json:"arn"`
	HostedZone  string `json:"hosted_zone"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryPrivateDnsNamespaceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsServiceDiscoveryPrivateDnsNamespace `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalWebAclAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalWebAclAssociationSpec `json"spec"`
}

type AwsWafregionalWebAclAssociationSpec struct {
	ResourceArn string `json:"resource_arn"`
	WebAclId    string `json:"web_acl_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalWebAclAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalWebAclAssociation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultVpcDhcpOptions struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDefaultVpcDhcpOptionsSpec `json"spec"`
}

type AwsDefaultVpcDhcpOptionsSpec struct {
	NetbiosNameServers []string          `json:"netbios_name_servers"`
	Tags               map[string]string `json:"tags"`
	DomainName         string            `json:"domain_name"`
	DomainNameServers  string            `json:"domain_name_servers"`
	NtpServers         string            `json:"ntp_servers"`
	NetbiosNodeType    string            `json:"netbios_node_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultVpcDhcpOptionsList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDefaultVpcDhcpOptions `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultSubnet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDefaultSubnetSpec `json"spec"`
}

type AwsDefaultSubnetSpec struct {
	AssignIpv6AddressOnCreation bool              `json:"assign_ipv6_address_on_creation"`
	Ipv6CidrBlockAssociationId  string            `json:"ipv6_cidr_block_association_id"`
	Tags                        map[string]string `json:"tags"`
	VpcId                       string            `json:"vpc_id"`
	CidrBlock                   string            `json:"cidr_block"`
	Ipv6CidrBlock               string            `json:"ipv6_cidr_block"`
	AvailabilityZone            string            `json:"availability_zone"`
	MapPublicIpOnLaunch         bool              `json:"map_public_ip_on_launch"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultSubnetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDefaultSubnet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultVpc struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDefaultVpcSpec `json"spec"`
}

type AwsDefaultVpcSpec struct {
	InstanceTenancy              string            `json:"instance_tenancy"`
	EnableClassiclinkDnsSupport  bool              `json:"enable_classiclink_dns_support"`
	DefaultNetworkAclId          string            `json:"default_network_acl_id"`
	Ipv6AssociationId            string            `json:"ipv6_association_id"`
	CidrBlock                    string            `json:"cidr_block"`
	EnableClassiclink            bool              `json:"enable_classiclink"`
	DhcpOptionsId                string            `json:"dhcp_options_id"`
	Tags                         map[string]string `json:"tags"`
	EnableDnsHostnames           bool              `json:"enable_dns_hostnames"`
	AssignGeneratedIpv6CidrBlock bool              `json:"assign_generated_ipv6_cidr_block"`
	MainRouteTableId             string            `json:"main_route_table_id"`
	EnableDnsSupport             bool              `json:"enable_dns_support"`
	DefaultSecurityGroupId       string            `json:"default_security_group_id"`
	DefaultRouteTableId          string            `json:"default_route_table_id"`
	Ipv6CidrBlock                string            `json:"ipv6_cidr_block"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultVpcList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDefaultVpc `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationSubnetGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDmsReplicationSubnetGroupSpec `json"spec"`
}

type AwsDmsReplicationSubnetGroupSpec struct {
	ReplicationSubnetGroupArn         string            `json:"replication_subnet_group_arn"`
	ReplicationSubnetGroupDescription string            `json:"replication_subnet_group_description"`
	ReplicationSubnetGroupId          string            `json:"replication_subnet_group_id"`
	SubnetIds                         string            `json:"subnet_ids"`
	Tags                              map[string]string `json:"tags"`
	VpcId                             string            `json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationSubnetGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDmsReplicationSubnetGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEbsSnapshot struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEbsSnapshotSpec `json"spec"`
}

type AwsEbsSnapshotSpec struct {
	DataEncryptionKeyId string            `json:"data_encryption_key_id"`
	Tags                map[string]string `json:"tags"`
	VolumeId            string            `json:"volume_id"`
	Encrypted           bool              `json:"encrypted"`
	KmsKeyId            string            `json:"kms_key_id"`
	VolumeSize          int               `json:"volume_size"`
	Description         string            `json:"description"`
	OwnerId             string            `json:"owner_id"`
	OwnerAlias          string            `json:"owner_alias"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEbsSnapshotList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEbsSnapshot `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkConfigurationTemplate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticBeanstalkConfigurationTemplateSpec `json"spec"`
}

type AwsElasticBeanstalkConfigurationTemplateSpec struct {
	Setting           AwsElasticBeanstalkConfigurationTemplateSpecSetting `json:"setting"`
	SolutionStackName string                                              `json:"solution_stack_name"`
	Name              string                                              `json:"name"`
	Application       string                                              `json:"application"`
	Description       string                                              `json:"description"`
	EnvironmentId     string                                              `json:"environment_id"`
}

type AwsElasticBeanstalkConfigurationTemplateSpecSetting struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Resource  string `json:"resource"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkConfigurationTemplateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticBeanstalkConfigurationTemplate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamOpenidConnectProvider struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamOpenidConnectProviderSpec `json"spec"`
}

type AwsIamOpenidConnectProviderSpec struct {
	Arn            string   `json:"arn"`
	Url            string   `json:"url"`
	ClientIdList   []string `json:"client_id_list"`
	ThumbprintList []string `json:"thumbprint_list"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamOpenidConnectProviderList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamOpenidConnectProvider `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKinesisFirehoseDeliveryStream struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsKinesisFirehoseDeliveryStreamSpec `json"spec"`
}

type AwsKinesisFirehoseDeliveryStreamSpec struct {
	Name                       string                                                           `json:"name"`
	KinesisSourceConfiguration []AwsKinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration `json:"kinesis_source_configuration"`
	S3Configuration            []AwsKinesisFirehoseDeliveryStreamSpecS3Configuration            `json:"s3_configuration"`
	ExtendedS3Configuration    []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3Configuration    `json:"extended_s3_configuration"`
	RedshiftConfiguration      []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfiguration      `json:"redshift_configuration"`
	SplunkConfiguration        []AwsKinesisFirehoseDeliveryStreamSpecSplunkConfiguration        `json:"splunk_configuration"`
	Destination                string                                                           `json:"destination"`
	ElasticsearchConfiguration []AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration `json:"elasticsearch_configuration"`
	Arn                        string                                                           `json:"arn"`
	VersionId                  string                                                           `json:"version_id"`
	DestinationId              string                                                           `json:"destination_id"`
}

type AwsKinesisFirehoseDeliveryStreamSpecKinesisSourceConfiguration struct {
	KinesisStreamArn string `json:"kinesis_stream_arn"`
	RoleArn          string `json:"role_arn"`
}

type AwsKinesisFirehoseDeliveryStreamSpecS3Configuration struct {
	CloudwatchLoggingOptions AwsKinesisFirehoseDeliveryStreamSpecS3ConfigurationCloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
	BucketArn                string                                                                      `json:"bucket_arn"`
	BufferSize               int                                                                         `json:"buffer_size"`
	BufferInterval           int                                                                         `json:"buffer_interval"`
	CompressionFormat        string                                                                      `json:"compression_format"`
	KmsKeyArn                string                                                                      `json:"kms_key_arn"`
	RoleArn                  string                                                                      `json:"role_arn"`
	Prefix                   string                                                                      `json:"prefix"`
}

type AwsKinesisFirehoseDeliveryStreamSpecS3ConfigurationCloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3Configuration struct {
	CloudwatchLoggingOptions AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationCloudwatchLoggingOptions  `json:"cloudwatch_logging_options"`
	BufferInterval           int                                                                                  `json:"buffer_interval"`
	RoleArn                  string                                                                               `json:"role_arn"`
	CompressionFormat        string                                                                               `json:"compression_format"`
	KmsKeyArn                string                                                                               `json:"kms_key_arn"`
	Prefix                   string                                                                               `json:"prefix"`
	S3BackupMode             string                                                                               `json:"s3_backup_mode"`
	S3BackupConfiguration    []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration   `json:"s3_backup_configuration"`
	ProcessingConfiguration  []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration `json:"processing_configuration"`
	BucketArn                string                                                                               `json:"bucket_arn"`
	BufferSize               int                                                                                  `json:"buffer_size"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationCloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfiguration struct {
	RoleArn                  string                                                                                                   `json:"role_arn"`
	Prefix                   string                                                                                                   `json:"prefix"`
	CloudwatchLoggingOptions AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
	BucketArn                string                                                                                                   `json:"bucket_arn"`
	BufferSize               int                                                                                                      `json:"buffer_size"`
	BufferInterval           int                                                                                                      `json:"buffer_interval"`
	CompressionFormat        string                                                                                                   `json:"compression_format"`
	KmsKeyArn                string                                                                                                   `json:"kms_key_arn"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfiguration struct {
	Enabled    bool                                                                                           `json:"enabled"`
	Processors []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors `json:"processors"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessors struct {
	Parameters []AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters `json:"parameters"`
	Type       string                                                                                                   `json:"type"`
}

type AwsKinesisFirehoseDeliveryStreamSpecExtendedS3ConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfiguration struct {
	Username                 string                                                                             `json:"username"`
	Password                 string                                                                             `json:"password"`
	S3BackupConfiguration    []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration   `json:"s3_backup_configuration"`
	RetryDuration            int                                                                                `json:"retry_duration"`
	DataTableColumns         string                                                                             `json:"data_table_columns"`
	DataTableName            string                                                                             `json:"data_table_name"`
	ClusterJdbcurl           string                                                                             `json:"cluster_jdbcurl"`
	ProcessingConfiguration  []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration `json:"processing_configuration"`
	RoleArn                  string                                                                             `json:"role_arn"`
	S3BackupMode             string                                                                             `json:"s3_backup_mode"`
	CopyOptions              string                                                                             `json:"copy_options"`
	CloudwatchLoggingOptions AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationCloudwatchLoggingOptions  `json:"cloudwatch_logging_options"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfiguration struct {
	CompressionFormat        string                                                                                                 `json:"compression_format"`
	KmsKeyArn                string                                                                                                 `json:"kms_key_arn"`
	RoleArn                  string                                                                                                 `json:"role_arn"`
	Prefix                   string                                                                                                 `json:"prefix"`
	CloudwatchLoggingOptions AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
	BucketArn                string                                                                                                 `json:"bucket_arn"`
	BufferSize               int                                                                                                    `json:"buffer_size"`
	BufferInterval           int                                                                                                    `json:"buffer_interval"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions struct {
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
	Enabled       bool   `json:"enabled"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfiguration struct {
	Enabled    bool                                                                                         `json:"enabled"`
	Processors []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors `json:"processors"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessors struct {
	Parameters []AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters `json:"parameters"`
	Type       string                                                                                                 `json:"type"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type AwsKinesisFirehoseDeliveryStreamSpecRedshiftConfigurationCloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecSplunkConfiguration struct {
	ProcessingConfiguration  []AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration `json:"processing_configuration"`
	HecAcknowledgmentTimeout int                                                                              `json:"hec_acknowledgment_timeout"`
	HecEndpoint              string                                                                           `json:"hec_endpoint"`
	HecEndpointType          string                                                                           `json:"hec_endpoint_type"`
	HecToken                 string                                                                           `json:"hec_token"`
	S3BackupMode             string                                                                           `json:"s3_backup_mode"`
	RetryDuration            int                                                                              `json:"retry_duration"`
	CloudwatchLoggingOptions AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationCloudwatchLoggingOptions  `json:"cloudwatch_logging_options"`
}

type AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfiguration struct {
	Enabled    bool                                                                                       `json:"enabled"`
	Processors []AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors `json:"processors"`
}

type AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessors struct {
	Parameters []AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters `json:"parameters"`
	Type       string                                                                                               `json:"type"`
}

type AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterValue string `json:"parameter_value"`
	ParameterName  string `json:"parameter_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecSplunkConfigurationCloudwatchLoggingOptions struct {
	LogStreamName string `json:"log_stream_name"`
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfiguration struct {
	BufferingSize            int                                                                                     `json:"buffering_size"`
	IndexName                string                                                                                  `json:"index_name"`
	S3BackupMode             string                                                                                  `json:"s3_backup_mode"`
	CloudwatchLoggingOptions AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationCloudwatchLoggingOptions  `json:"cloudwatch_logging_options"`
	ProcessingConfiguration  []AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration `json:"processing_configuration"`
	TypeName                 string                                                                                  `json:"type_name"`
	BufferingInterval        int                                                                                     `json:"buffering_interval"`
	DomainArn                string                                                                                  `json:"domain_arn"`
	IndexRotationPeriod      string                                                                                  `json:"index_rotation_period"`
	RetryDuration            int                                                                                     `json:"retry_duration"`
	RoleArn                  string                                                                                  `json:"role_arn"`
}

type AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationCloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfiguration struct {
	Processors []AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors `json:"processors"`
	Enabled    bool                                                                                              `json:"enabled"`
}

type AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessors struct {
	Parameters []AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters `json:"parameters"`
	Type       string                                                                                                      `json:"type"`
}

type AwsKinesisFirehoseDeliveryStreamSpecElasticsearchConfigurationProcessingConfigurationProcessorsParameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKinesisFirehoseDeliveryStreamList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsKinesisFirehoseDeliveryStream `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMainRouteTableAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsMainRouteTableAssociationSpec `json"spec"`
}

type AwsMainRouteTableAssociationSpec struct {
	VpcId                string `json:"vpc_id"`
	RouteTableId         string `json:"route_table_id"`
	OriginalRouteTableId string `json:"original_route_table_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMainRouteTableAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsMainRouteTableAssociation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRegexMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalRegexMatchSetSpec `json"spec"`
}

type AwsWafregionalRegexMatchSetSpec struct {
	Name            string                                         `json:"name"`
	RegexMatchTuple AwsWafregionalRegexMatchSetSpecRegexMatchTuple `json:"regex_match_tuple"`
}

type AwsWafregionalRegexMatchSetSpecRegexMatchTuple struct {
	FieldToMatch       []AwsWafregionalRegexMatchSetSpecRegexMatchTupleFieldToMatch `json:"field_to_match"`
	RegexPatternSetId  string                                                       `json:"regex_pattern_set_id"`
	TextTransformation string                                                       `json:"text_transformation"`
}

type AwsWafregionalRegexMatchSetSpecRegexMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRegexMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalRegexMatchSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDevicefarmProject struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDevicefarmProjectSpec `json"spec"`
}

type AwsDevicefarmProjectSpec struct {
	Arn  string `json:"arn"`
	Name string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDevicefarmProjectList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDevicefarmProject `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserPolicyAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamUserPolicyAttachmentSpec `json"spec"`
}

type AwsIamUserPolicyAttachmentSpec struct {
	PolicyArn string `json:"policy_arn"`
	User      string `json:"user"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserPolicyAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamUserPolicyAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLaunchTemplate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLaunchTemplateSpec `json"spec"`
}

type AwsLaunchTemplateSpec struct {
	LatestVersion                     int                                             `json:"latest_version"`
	UserData                          string                                          `json:"user_data"`
	VpcSecurityGroupIds               string                                          `json:"vpc_security_group_ids"`
	DefaultVersion                    int                                             `json:"default_version"`
	CreditSpecification               []AwsLaunchTemplateSpecCreditSpecification      `json:"credit_specification"`
	ElasticGpuSpecifications          []AwsLaunchTemplateSpecElasticGpuSpecifications `json:"elastic_gpu_specifications"`
	KernelId                          string                                          `json:"kernel_id"`
	KeyName                           string                                          `json:"key_name"`
	RamDiskId                         string                                          `json:"ram_disk_id"`
	SecurityGroupNames                string                                          `json:"security_group_names"`
	Description                       string                                          `json:"description"`
	EbsOptimized                      bool                                            `json:"ebs_optimized"`
	IamInstanceProfile                []AwsLaunchTemplateSpecIamInstanceProfile       `json:"iam_instance_profile"`
	InstanceInitiatedShutdownBehavior string                                          `json:"instance_initiated_shutdown_behavior"`
	InstanceMarketOptions             []AwsLaunchTemplateSpecInstanceMarketOptions    `json:"instance_market_options"`
	NetworkInterfaces                 []AwsLaunchTemplateSpecNetworkInterfaces        `json:"network_interfaces"`
	Monitoring                        []AwsLaunchTemplateSpecMonitoring               `json:"monitoring"`
	Placement                         []AwsLaunchTemplateSpecPlacement                `json:"placement"`
	Name                              string                                          `json:"name"`
	NamePrefix                        string                                          `json:"name_prefix"`
	BlockDeviceMappings               []AwsLaunchTemplateSpecBlockDeviceMappings      `json:"block_device_mappings"`
	DisableApiTermination             bool                                            `json:"disable_api_termination"`
	ImageId                           string                                          `json:"image_id"`
	InstanceType                      string                                          `json:"instance_type"`
	TagSpecifications                 []AwsLaunchTemplateSpecTagSpecifications        `json:"tag_specifications"`
}

type AwsLaunchTemplateSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type AwsLaunchTemplateSpecElasticGpuSpecifications struct {
	Type string `json:"type"`
}

type AwsLaunchTemplateSpecIamInstanceProfile struct {
	Name string `json:"name"`
	Arn  string `json:"arn"`
}

type AwsLaunchTemplateSpecInstanceMarketOptions struct {
	MarketType  string                                                  `json:"market_type"`
	SpotOptions []AwsLaunchTemplateSpecInstanceMarketOptionsSpotOptions `json:"spot_options"`
}

type AwsLaunchTemplateSpecInstanceMarketOptionsSpotOptions struct {
	SpotInstanceType             string `json:"spot_instance_type"`
	ValidUntil                   string `json:"valid_until"`
	BlockDurationMinutes         int    `json:"block_duration_minutes"`
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior"`
	MaxPrice                     string `json:"max_price"`
}

type AwsLaunchTemplateSpecNetworkInterfaces struct {
	Description              string `json:"description"`
	Ipv6AddressCount         int    `json:"ipv6_address_count"`
	Ipv6Addresses            string `json:"ipv6_addresses"`
	PrivateIpAddress         string `json:"private_ip_address"`
	SubnetId                 string `json:"subnet_id"`
	AssociatePublicIpAddress bool   `json:"associate_public_ip_address"`
	DeleteOnTermination      bool   `json:"delete_on_termination"`
	DeviceIndex              int    `json:"device_index"`
	SecurityGroups           string `json:"security_groups"`
	NetworkInterfaceId       string `json:"network_interface_id"`
	Ipv4Addresses            string `json:"ipv4_addresses"`
	Ipv4AddressCount         int    `json:"ipv4_address_count"`
}

type AwsLaunchTemplateSpecMonitoring struct {
	Enabled bool `json:"enabled"`
}

type AwsLaunchTemplateSpecPlacement struct {
	Affinity         string `json:"affinity"`
	AvailabilityZone string `json:"availability_zone"`
	GroupName        string `json:"group_name"`
	HostId           string `json:"host_id"`
	SpreadDomain     string `json:"spread_domain"`
	Tenancy          string `json:"tenancy"`
}

type AwsLaunchTemplateSpecBlockDeviceMappings struct {
	NoDevice    string                                        `json:"no_device"`
	VirtualName string                                        `json:"virtual_name"`
	Ebs         []AwsLaunchTemplateSpecBlockDeviceMappingsEbs `json:"ebs"`
	DeviceName  string                                        `json:"device_name"`
}

type AwsLaunchTemplateSpecBlockDeviceMappingsEbs struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsLaunchTemplateSpecTagSpecifications struct {
	ResourceType string            `json:"resource_type"`
	Tags         map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLaunchTemplateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLaunchTemplate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRegexPatternSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafRegexPatternSetSpec `json"spec"`
}

type AwsWafRegexPatternSetSpec struct {
	Name                string `json:"name"`
	RegexPatternStrings string `json:"regex_pattern_strings"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRegexPatternSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafRegexPatternSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiCopy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAmiCopySpec `json"spec"`
}

type AwsAmiCopySpec struct {
	KernelId             string                             `json:"kernel_id"`
	Encrypted            bool                               `json:"encrypted"`
	EbsBlockDevice       AwsAmiCopySpecEbsBlockDevice       `json:"ebs_block_device"`
	Tags                 map[string]string                  `json:"tags"`
	SourceAmiId          string                             `json:"source_ami_id"`
	SourceAmiRegion      string                             `json:"source_ami_region"`
	ImageLocation        string                             `json:"image_location"`
	Name                 string                             `json:"name"`
	RootDeviceName       string                             `json:"root_device_name"`
	SriovNetSupport      string                             `json:"sriov_net_support"`
	ManageEbsSnapshots   bool                               `json:"manage_ebs_snapshots"`
	KmsKeyId             string                             `json:"kms_key_id"`
	Architecture         string                             `json:"architecture"`
	RamdiskId            string                             `json:"ramdisk_id"`
	VirtualizationType   string                             `json:"virtualization_type"`
	EphemeralBlockDevice AwsAmiCopySpecEphemeralBlockDevice `json:"ephemeral_block_device"`
	Description          string                             `json:"description"`
	RootSnapshotId       string                             `json:"root_snapshot_id"`
}

type AwsAmiCopySpecEbsBlockDevice struct {
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
}

type AwsAmiCopySpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiCopyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAmiCopy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAutoscalingAttachmentSpec `json"spec"`
}

type AwsAutoscalingAttachmentSpec struct {
	AutoscalingGroupName string `json:"autoscaling_group_name"`
	Elb                  string `json:"elb"`
	AlbTargetGroupArn    string `json:"alb_target_group_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAutoscalingAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogStream struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchLogStreamSpec `json"spec"`
}

type AwsCloudwatchLogStreamSpec struct {
	LogGroupName string `json:"log_group_name"`
	Arn          string `json:"arn"`
	Name         string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogStreamList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchLogStream `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorAssessmentTarget struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsInspectorAssessmentTargetSpec `json"spec"`
}

type AwsInspectorAssessmentTargetSpec struct {
	Name             string `json:"name"`
	Arn              string `json:"arn"`
	ResourceGroupArn string `json:"resource_group_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorAssessmentTargetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsInspectorAssessmentTarget `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSecurityGroupSpec `json"spec"`
}

type AwsSecurityGroupSpec struct {
	Name                string                      `json:"name"`
	VpcId               string                      `json:"vpc_id"`
	Ingress             AwsSecurityGroupSpecIngress `json:"ingress"`
	RevokeRulesOnDelete bool                        `json:"revoke_rules_on_delete"`
	NamePrefix          string                      `json:"name_prefix"`
	Description         string                      `json:"description"`
	Egress              AwsSecurityGroupSpecEgress  `json:"egress"`
	Arn                 string                      `json:"arn"`
	OwnerId             string                      `json:"owner_id"`
	Tags                map[string]string           `json:"tags"`
}

type AwsSecurityGroupSpecIngress struct {
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	SecurityGroups string   `json:"security_groups"`
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
	ToPort         int      `json:"to_port"`
}

type AwsSecurityGroupSpecEgress struct {
	ToPort         int      `json:"to_port"`
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	SecurityGroups string   `json:"security_groups"`
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
	Self           bool     `json:"self"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSecurityGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListener struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLbListenerSpec `json"spec"`
}

type AwsLbListenerSpec struct {
	DefaultAction   []AwsLbListenerSpecDefaultAction `json:"default_action"`
	Arn             string                           `json:"arn"`
	LoadBalancerArn string                           `json:"load_balancer_arn"`
	Port            int                              `json:"port"`
	Protocol        string                           `json:"protocol"`
	SslPolicy       string                           `json:"ssl_policy"`
	CertificateArn  string                           `json:"certificate_arn"`
}

type AwsLbListenerSpecDefaultAction struct {
	TargetGroupArn string `json:"target_group_arn"`
	Type           string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLbListener `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayClientCertificate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayClientCertificateSpec `json"spec"`
}

type AwsApiGatewayClientCertificateSpec struct {
	Description           string `json:"description"`
	CreatedDate           string `json:"created_date"`
	ExpirationDate        string `json:"expiration_date"`
	PemEncodedCertificate string `json:"pem_encoded_certificate"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayClientCertificateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayClientCertificate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDeployment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayDeploymentSpec `json"spec"`
}

type AwsApiGatewayDeploymentSpec struct {
	ExecutionArn     string            `json:"execution_arn"`
	RestApiId        string            `json:"rest_api_id"`
	StageName        string            `json:"stage_name"`
	Description      string            `json:"description"`
	StageDescription string            `json:"stage_description"`
	Variables        map[string]string `json:"variables"`
	CreatedDate      string            `json:"created_date"`
	InvokeUrl        string            `json:"invoke_url"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDeploymentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayDeployment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodecommitRepository struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCodecommitRepositorySpec `json"spec"`
}

type AwsCodecommitRepositorySpec struct {
	RepositoryId   string `json:"repository_id"`
	CloneUrlHttp   string `json:"clone_url_http"`
	CloneUrlSsh    string `json:"clone_url_ssh"`
	DefaultBranch  string `json:"default_branch"`
	RepositoryName string `json:"repository_name"`
	Description    string `json:"description"`
	Arn            string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodecommitRepositoryList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCodecommitRepository `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcrLifecyclePolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEcrLifecyclePolicySpec `json"spec"`
}

type AwsEcrLifecyclePolicySpec struct {
	Repository string `json:"repository"`
	Policy     string `json:"policy"`
	RegistryId string `json:"registry_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcrLifecyclePolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEcrLifecyclePolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpc struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcSpec `json"spec"`
}

type AwsVpcSpec struct {
	InstanceTenancy              string            `json:"instance_tenancy"`
	DefaultNetworkAclId          string            `json:"default_network_acl_id"`
	DefaultSecurityGroupId       string            `json:"default_security_group_id"`
	DefaultRouteTableId          string            `json:"default_route_table_id"`
	Tags                         map[string]string `json:"tags"`
	EnableDnsHostnames           bool              `json:"enable_dns_hostnames"`
	EnableDnsSupport             bool              `json:"enable_dns_support"`
	Ipv6AssociationId            string            `json:"ipv6_association_id"`
	EnableClassiclink            bool              `json:"enable_classiclink"`
	MainRouteTableId             string            `json:"main_route_table_id"`
	Ipv6CidrBlock                string            `json:"ipv6_cidr_block"`
	CidrBlock                    string            `json:"cidr_block"`
	EnableClassiclinkDnsSupport  bool              `json:"enable_classiclink_dns_support"`
	AssignGeneratedIpv6CidrBlock bool              `json:"assign_generated_ipv6_cidr_block"`
	DhcpOptionsId                string            `json:"dhcp_options_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpc `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRegexMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafRegexMatchSetSpec `json"spec"`
}

type AwsWafRegexMatchSetSpec struct {
	Name            string                                 `json:"name"`
	RegexMatchTuple AwsWafRegexMatchSetSpecRegexMatchTuple `json:"regex_match_tuple"`
}

type AwsWafRegexMatchSetSpecRegexMatchTuple struct {
	FieldToMatch       []AwsWafRegexMatchSetSpecRegexMatchTupleFieldToMatch `json:"field_to_match"`
	RegexPatternSetId  string                                               `json:"regex_pattern_set_id"`
	TextTransformation string                                               `json:"text_transformation"`
}

type AwsWafRegexMatchSetSpecRegexMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRegexMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafRegexMatchSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingLifecycleHook struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAutoscalingLifecycleHookSpec `json"spec"`
}

type AwsAutoscalingLifecycleHookSpec struct {
	DefaultResult         string `json:"default_result"`
	HeartbeatTimeout      int    `json:"heartbeat_timeout"`
	LifecycleTransition   string `json:"lifecycle_transition"`
	NotificationMetadata  string `json:"notification_metadata"`
	NotificationTargetArn string `json:"notification_target_arn"`
	RoleArn               string `json:"role_arn"`
	Name                  string `json:"name"`
	AutoscalingGroupName  string `json:"autoscaling_group_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingLifecycleHookList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAutoscalingLifecycleHook `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGameliftBuild struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGameliftBuildSpec `json"spec"`
}

type AwsGameliftBuildSpec struct {
	Version         string                                `json:"version"`
	Name            string                                `json:"name"`
	OperatingSystem string                                `json:"operating_system"`
	StorageLocation []AwsGameliftBuildSpecStorageLocation `json:"storage_location"`
}

type AwsGameliftBuildSpecStorageLocation struct {
	Bucket  string `json:"bucket"`
	Key     string `json:"key"`
	RoleArn string `json:"role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGameliftBuildList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGameliftBuild `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesReceiptRuleSpec `json"spec"`
}

type AwsSesReceiptRuleSpec struct {
	AddHeaderAction AwsSesReceiptRuleSpecAddHeaderAction `json:"add_header_action"`
	StopAction      AwsSesReceiptRuleSpecStopAction      `json:"stop_action"`
	Name            string                               `json:"name"`
	ScanEnabled     bool                                 `json:"scan_enabled"`
	TlsPolicy       string                               `json:"tls_policy"`
	BounceAction    AwsSesReceiptRuleSpecBounceAction    `json:"bounce_action"`
	Enabled         bool                                 `json:"enabled"`
	SnsAction       AwsSesReceiptRuleSpecSnsAction       `json:"sns_action"`
	WorkmailAction  AwsSesReceiptRuleSpecWorkmailAction  `json:"workmail_action"`
	RuleSetName     string                               `json:"rule_set_name"`
	After           string                               `json:"after"`
	Recipients      string                               `json:"recipients"`
	LambdaAction    AwsSesReceiptRuleSpecLambdaAction    `json:"lambda_action"`
	S3Action        AwsSesReceiptRuleSpecS3Action        `json:"s3_action"`
}

type AwsSesReceiptRuleSpecAddHeaderAction struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
	Position    int    `json:"position"`
}

type AwsSesReceiptRuleSpecStopAction struct {
	Scope    string `json:"scope"`
	TopicArn string `json:"topic_arn"`
	Position int    `json:"position"`
}

type AwsSesReceiptRuleSpecBounceAction struct {
	Position      int    `json:"position"`
	Message       string `json:"message"`
	Sender        string `json:"sender"`
	SmtpReplyCode string `json:"smtp_reply_code"`
	StatusCode    string `json:"status_code"`
	TopicArn      string `json:"topic_arn"`
}

type AwsSesReceiptRuleSpecSnsAction struct {
	TopicArn string `json:"topic_arn"`
	Position int    `json:"position"`
}

type AwsSesReceiptRuleSpecWorkmailAction struct {
	TopicArn        string `json:"topic_arn"`
	Position        int    `json:"position"`
	OrganizationArn string `json:"organization_arn"`
}

type AwsSesReceiptRuleSpecLambdaAction struct {
	TopicArn       string `json:"topic_arn"`
	Position       int    `json:"position"`
	FunctionArn    string `json:"function_arn"`
	InvocationType string `json:"invocation_type"`
}

type AwsSesReceiptRuleSpecS3Action struct {
	BucketName      string `json:"bucket_name"`
	KmsKeyArn       string `json:"kms_key_arn"`
	ObjectKeyPrefix string `json:"object_key_prefix"`
	TopicArn        string `json:"topic_arn"`
	Position        int    `json:"position"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesReceiptRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSubnet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSubnetSpec `json"spec"`
}

type AwsSubnetSpec struct {
	CidrBlock                   string            `json:"cidr_block"`
	Ipv6CidrBlock               string            `json:"ipv6_cidr_block"`
	AvailabilityZone            string            `json:"availability_zone"`
	MapPublicIpOnLaunch         bool              `json:"map_public_ip_on_launch"`
	AssignIpv6AddressOnCreation bool              `json:"assign_ipv6_address_on_creation"`
	Ipv6CidrBlockAssociationId  string            `json:"ipv6_cidr_block_association_id"`
	Tags                        map[string]string `json:"tags"`
	VpcId                       string            `json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSubnetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSubnet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcDhcpOptions struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcDhcpOptionsSpec `json"spec"`
}

type AwsVpcDhcpOptionsSpec struct {
	NetbiosNodeType    string            `json:"netbios_node_type"`
	NetbiosNameServers []string          `json:"netbios_name_servers"`
	Tags               map[string]string `json:"tags"`
	DomainName         string            `json:"domain_name"`
	DomainNameServers  []string          `json:"domain_name_servers"`
	NtpServers         []string          `json:"ntp_servers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcDhcpOptionsList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcDhcpOptions `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayAuthorizer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayAuthorizerSpec `json"spec"`
}

type AwsApiGatewayAuthorizerSpec struct {
	Type                         string `json:"type"`
	IdentityValidationExpression string `json:"identity_validation_expression"`
	IdentitySource               string `json:"identity_source"`
	Name                         string `json:"name"`
	RestApiId                    string `json:"rest_api_id"`
	ProviderArns                 string `json:"provider_arns"`
	AuthorizerUri                string `json:"authorizer_uri"`
	AuthorizerCredentials        string `json:"authorizer_credentials"`
	AuthorizerResultTtlInSeconds int    `json:"authorizer_result_ttl_in_seconds"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayAuthorizerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayAuthorizer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloud9EnvironmentEc2 struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloud9EnvironmentEc2Spec `json"spec"`
}

type AwsCloud9EnvironmentEc2Spec struct {
	OwnerArn                 string `json:"owner_arn"`
	SubnetId                 string `json:"subnet_id"`
	Arn                      string `json:"arn"`
	Type                     string `json:"type"`
	Name                     string `json:"name"`
	InstanceType             string `json:"instance_type"`
	AutomaticStopTimeMinutes int    `json:"automatic_stop_time_minutes"`
	Description              string `json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloud9EnvironmentEc2List struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloud9EnvironmentEc2 `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaEventSourceMapping struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLambdaEventSourceMappingSpec `json"spec"`
}

type AwsLambdaEventSourceMappingSpec struct {
	FunctionName          string `json:"function_name"`
	StartingPosition      string `json:"starting_position"`
	BatchSize             int    `json:"batch_size"`
	Enabled               bool   `json:"enabled"`
	LastModified          string `json:"last_modified"`
	Uuid                  string `json:"uuid"`
	EventSourceArn        string `json:"event_source_arn"`
	FunctionArn           string `json:"function_arn"`
	LastProcessingResult  string `json:"last_processing_result"`
	State                 string `json:"state"`
	StateTransitionReason string `json:"state_transition_reason"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaEventSourceMappingList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLambdaEventSourceMapping `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsOrganization struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOrganizationsOrganizationSpec `json"spec"`
}

type AwsOrganizationsOrganizationSpec struct {
	MasterAccountEmail string `json:"master_account_email"`
	MasterAccountId    string `json:"master_account_id"`
	FeatureSet         string `json:"feature_set"`
	Arn                string `json:"arn"`
	MasterAccountArn   string `json:"master_account_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsOrganizationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOrganizationsOrganization `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketMetric struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsS3BucketMetricSpec `json"spec"`
}

type AwsS3BucketMetricSpec struct {
	Bucket string                        `json:"bucket"`
	Filter []AwsS3BucketMetricSpecFilter `json:"filter"`
	Name   string                        `json:"name"`
}

type AwsS3BucketMetricSpecFilter struct {
	Prefix string            `json:"prefix"`
	Tags   map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketMetricList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsS3BucketMetric `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafRuleSpec `json"spec"`
}

type AwsWafRuleSpec struct {
	Name       string                   `json:"name"`
	MetricName string                   `json:"metric_name"`
	Predicates AwsWafRuleSpecPredicates `json:"predicates"`
}

type AwsWafRuleSpecPredicates struct {
	Negated bool   `json:"negated"`
	DataId  string `json:"data_id"`
	Type    string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxLag struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDxLagSpec `json"spec"`
}

type AwsDxLagSpec struct {
	Arn                  string            `json:"arn"`
	Name                 string            `json:"name"`
	ConnectionsBandwidth string            `json:"connections_bandwidth"`
	Location             string            `json:"location"`
	NumberOfConnections  int               `json:"number_of_connections"`
	ForceDestroy         bool              `json:"force_destroy"`
	Tags                 map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxLagList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDxLag `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyThreatintelset struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGuarddutyThreatintelsetSpec `json"spec"`
}

type AwsGuarddutyThreatintelsetSpec struct {
	Format     string `json:"format"`
	Location   string `json:"location"`
	Activate   bool   `json:"activate"`
	DetectorId string `json:"detector_id"`
	Name       string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyThreatintelsetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGuarddutyThreatintelset `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamGroupPolicySpec `json"spec"`
}

type AwsIamGroupPolicySpec struct {
	Policy     string `json:"policy"`
	Name       string `json:"name"`
	NamePrefix string `json:"name_prefix"`
	Group      string `json:"group"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamGroupPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksApplication struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksApplicationSpec `json"spec"`
}

type AwsOpsworksApplicationSpec struct {
	Name                   string                                       `json:"name"`
	Type                   string                                       `json:"type"`
	DocumentRoot           string                                       `json:"document_root"`
	Domains                []string                                     `json:"domains"`
	Environment            AwsOpsworksApplicationSpecEnvironment        `json:"environment"`
	EnableSsl              bool                                         `json:"enable_ssl"`
	AwsFlowRubySettings    string                                       `json:"aws_flow_ruby_settings"`
	DataSourceArn          string                                       `json:"data_source_arn"`
	StackId                string                                       `json:"stack_id"`
	RailsEnv               string                                       `json:"rails_env"`
	AutoBundleOnDeploy     string                                       `json:"auto_bundle_on_deploy"`
	DataSourceType         string                                       `json:"data_source_type"`
	Description            string                                       `json:"description"`
	ShortName              string                                       `json:"short_name"`
	AppSource              []AwsOpsworksApplicationSpecAppSource        `json:"app_source"`
	DataSourceDatabaseName string                                       `json:"data_source_database_name"`
	SslConfiguration       []AwsOpsworksApplicationSpecSslConfiguration `json:"ssl_configuration"`
}

type AwsOpsworksApplicationSpecEnvironment struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Secure bool   `json:"secure"`
}

type AwsOpsworksApplicationSpecAppSource struct {
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
}

type AwsOpsworksApplicationSpecSslConfiguration struct {
	PrivateKey  string `json:"private_key"`
	Chain       string `json:"chain"`
	Certificate string `json:"certificate"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksApplicationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksApplication `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnapshotCreateVolumePermission struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSnapshotCreateVolumePermissionSpec `json"spec"`
}

type AwsSnapshotCreateVolumePermissionSpec struct {
	SnapshotId string `json:"snapshot_id"`
	AccountId  string `json:"account_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnapshotCreateVolumePermissionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSnapshotCreateVolumePermission `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRateBasedRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafRateBasedRuleSpec `json"spec"`
}

type AwsWafRateBasedRuleSpec struct {
	Name       string                            `json:"name"`
	MetricName string                            `json:"metric_name"`
	Predicates AwsWafRateBasedRuleSpecPredicates `json:"predicates"`
	RateKey    string                            `json:"rate_key"`
	RateLimit  int                               `json:"rate_limit"`
}

type AwsWafRateBasedRuleSpecPredicates struct {
	Negated bool   `json:"negated"`
	DataId  string `json:"data_id"`
	Type    string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRateBasedRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafRateBasedRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEip struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEipSpec `json"spec"`
}

type AwsEipSpec struct {
	PrivateIp              string            `json:"private_ip"`
	AssociateWithPrivateIp string            `json:"associate_with_private_ip"`
	Tags                   map[string]string `json:"tags"`
	Instance               string            `json:"instance"`
	NetworkInterface       string            `json:"network_interface"`
	AllocationId           string            `json:"allocation_id"`
	PublicIp               string            `json:"public_ip"`
	Vpc                    bool              `json:"vpc"`
	AssociationId          string            `json:"association_id"`
	Domain                 string            `json:"domain"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEipList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEip `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElastictranscoderPreset struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElastictranscoderPresetSpec `json"spec"`
}

type AwsElastictranscoderPresetSpec struct {
	Arn               string                                          `json:"arn"`
	Audio             AwsElastictranscoderPresetSpecAudio             `json:"audio"`
	Description       string                                          `json:"description"`
	Name              string                                          `json:"name"`
	VideoWatermarks   AwsElastictranscoderPresetSpecVideoWatermarks   `json:"video_watermarks"`
	VideoCodecOptions map[string]string                               `json:"video_codec_options"`
	AudioCodecOptions AwsElastictranscoderPresetSpecAudioCodecOptions `json:"audio_codec_options"`
	Container         string                                          `json:"container"`
	Thumbnails        AwsElastictranscoderPresetSpecThumbnails        `json:"thumbnails"`
	Type              string                                          `json:"type"`
	Video             AwsElastictranscoderPresetSpecVideo             `json:"video"`
}

type AwsElastictranscoderPresetSpecAudio struct {
	SampleRate       string `json:"sample_rate"`
	AudioPackingMode string `json:"audio_packing_mode"`
	BitRate          string `json:"bit_rate"`
	Channels         string `json:"channels"`
	Codec            string `json:"codec"`
}

type AwsElastictranscoderPresetSpecVideoWatermarks struct {
	HorizontalAlign  string `json:"horizontal_align"`
	HorizontalOffset string `json:"horizontal_offset"`
	MaxWidth         string `json:"max_width"`
	SizingPolicy     string `json:"sizing_policy"`
	VerticalAlign    string `json:"vertical_align"`
	Id               string `json:"id"`
	MaxHeight        string `json:"max_height"`
	Opacity          string `json:"opacity"`
	Target           string `json:"target"`
	VerticalOffset   string `json:"vertical_offset"`
}

type AwsElastictranscoderPresetSpecAudioCodecOptions struct {
	BitDepth string `json:"bit_depth"`
	BitOrder string `json:"bit_order"`
	Profile  string `json:"profile"`
	Signed   string `json:"signed"`
}

type AwsElastictranscoderPresetSpecThumbnails struct {
	PaddingPolicy string `json:"padding_policy"`
	Resolution    string `json:"resolution"`
	SizingPolicy  string `json:"sizing_policy"`
	AspectRatio   string `json:"aspect_ratio"`
	Format        string `json:"format"`
	Interval      string `json:"interval"`
	MaxHeight     string `json:"max_height"`
	MaxWidth      string `json:"max_width"`
}

type AwsElastictranscoderPresetSpecVideo struct {
	Resolution         string `json:"resolution"`
	SizingPolicy       string `json:"sizing_policy"`
	AspectRatio        string `json:"aspect_ratio"`
	Codec              string `json:"codec"`
	MaxHeight          string `json:"max_height"`
	FrameRate          string `json:"frame_rate"`
	KeyframesMaxDist   string `json:"keyframes_max_dist"`
	MaxFrameRate       string `json:"max_frame_rate"`
	MaxWidth           string `json:"max_width"`
	PaddingPolicy      string `json:"padding_policy"`
	BitRate            string `json:"bit_rate"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	FixedGop           string `json:"fixed_gop"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElastictranscoderPresetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElastictranscoderPreset `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaAlias struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLambdaAliasSpec `json"spec"`
}

type AwsLambdaAliasSpec struct {
	Description     string `json:"description"`
	FunctionName    string `json:"function_name"`
	FunctionVersion string `json:"function_version"`
	Name            string `json:"name"`
	Arn             string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaAliasList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLambdaAlias `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalSqlInjectionMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalSqlInjectionMatchSetSpec `json"spec"`
}

type AwsWafregionalSqlInjectionMatchSetSpec struct {
	Name                   string                                                       `json:"name"`
	SqlInjectionMatchTuple AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTuple `json:"sql_injection_match_tuple"`
}

type AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTuple struct {
	FieldToMatch       []AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch `json:"field_to_match"`
	TextTransformation string                                                                     `json:"text_transformation"`
}

type AwsWafregionalSqlInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalSqlInjectionMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalSqlInjectionMatchSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainMailFrom struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesDomainMailFromSpec `json"spec"`
}

type AwsSesDomainMailFromSpec struct {
	Domain              string `json:"domain"`
	MailFromDomain      string `json:"mail_from_domain"`
	BehaviorOnMxFailure string `json:"behavior_on_mx_failure"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainMailFromList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesDomainMailFrom `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigConfigurationRecorderStatus struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsConfigConfigurationRecorderStatusSpec `json"spec"`
}

type AwsConfigConfigurationRecorderStatusSpec struct {
	Name      string `json:"name"`
	IsEnabled bool   `json:"is_enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigConfigurationRecorderStatusList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsConfigConfigurationRecorderStatus `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodecommitTrigger struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCodecommitTriggerSpec `json"spec"`
}

type AwsCodecommitTriggerSpec struct {
	RepositoryName  string                          `json:"repository_name"`
	ConfigurationId string                          `json:"configuration_id"`
	Trigger         AwsCodecommitTriggerSpecTrigger `json:"trigger"`
}

type AwsCodecommitTriggerSpecTrigger struct {
	Branches       []string `json:"branches"`
	Events         []string `json:"events"`
	Name           string   `json:"name"`
	DestinationArn string   `json:"destination_arn"`
	CustomData     string   `json:"custom_data"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodecommitTriggerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCodecommitTrigger `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationInstance struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDmsReplicationInstanceSpec `json"spec"`
}

type AwsDmsReplicationInstanceSpec struct {
	AutoMinorVersionUpgrade       bool              `json:"auto_minor_version_upgrade"`
	MultiAz                       bool              `json:"multi_az"`
	PubliclyAccessible            bool              `json:"publicly_accessible"`
	ReplicationInstanceId         string            `json:"replication_instance_id"`
	EngineVersion                 string            `json:"engine_version"`
	KmsKeyArn                     string            `json:"kms_key_arn"`
	ReplicationInstanceArn        string            `json:"replication_instance_arn"`
	ReplicationInstancePublicIps  []string          `json:"replication_instance_public_ips"`
	ReplicationSubnetGroupId      string            `json:"replication_subnet_group_id"`
	ApplyImmediately              bool              `json:"apply_immediately"`
	AllocatedStorage              int               `json:"allocated_storage"`
	AvailabilityZone              string            `json:"availability_zone"`
	PreferredMaintenanceWindow    string            `json:"preferred_maintenance_window"`
	ReplicationInstanceClass      string            `json:"replication_instance_class"`
	ReplicationInstancePrivateIps []string          `json:"replication_instance_private_ips"`
	Tags                          map[string]string `json:"tags"`
	VpcSecurityGroupIds           string            `json:"vpc_security_group_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationInstanceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDmsReplicationInstance `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsMountTarget struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEfsMountTargetSpec `json"spec"`
}

type AwsEfsMountTargetSpec struct {
	DnsName            string `json:"dns_name"`
	FileSystemId       string `json:"file_system_id"`
	IpAddress          string `json:"ip_address"`
	SecurityGroups     string `json:"security_groups"`
	SubnetId           string `json:"subnet_id"`
	NetworkInterfaceId string `json:"network_interface_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsMountTargetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEfsMountTarget `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaPermission struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLambdaPermissionSpec `json"spec"`
}

type AwsLambdaPermissionSpec struct {
	Principal         string `json:"principal"`
	Qualifier         string `json:"qualifier"`
	SourceAccount     string `json:"source_account"`
	SourceArn         string `json:"source_arn"`
	StatementId       string `json:"statement_id"`
	StatementIdPrefix string `json:"statement_id_prefix"`
	Action            string `json:"action"`
	FunctionName      string `json:"function_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaPermissionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLambdaPermission `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksRdsDbInstance struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksRdsDbInstanceSpec `json"spec"`
}

type AwsOpsworksRdsDbInstanceSpec struct {
	DbUser           string `json:"db_user"`
	StackId          string `json:"stack_id"`
	RdsDbInstanceArn string `json:"rds_db_instance_arn"`
	DbPassword       string `json:"db_password"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksRdsDbInstanceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksRdsDbInstance `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53QueryLog struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRoute53QueryLogSpec `json"spec"`
}

type AwsRoute53QueryLogSpec struct {
	ZoneId                string `json:"zone_id"`
	CloudwatchLogGroupArn string `json:"cloudwatch_log_group_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53QueryLogList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRoute53QueryLog `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpnGatewayAttachmentSpec `json"spec"`
}

type AwsVpnGatewayAttachmentSpec struct {
	VpcId        string `json:"vpc_id"`
	VpnGatewayId string `json:"vpn_gateway_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpnGatewayAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayVpcLink struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayVpcLinkSpec `json"spec"`
}

type AwsApiGatewayVpcLinkSpec struct {
	TargetArns  string `json:"target_arns"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayVpcLinkList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayVpcLink `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsCluster struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEcsClusterSpec `json"spec"`
}

type AwsEcsClusterSpec struct {
	Name string `json:"name"`
	Arn  string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsClusterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEcsCluster `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticsearchDomainPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticsearchDomainPolicySpec `json"spec"`
}

type AwsElasticsearchDomainPolicySpec struct {
	DomainName     string `json:"domain_name"`
	AccessPolicies string `json:"access_policies"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticsearchDomainPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticsearchDomainPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUser struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamUserSpec `json"spec"`
}

type AwsIamUserSpec struct {
	Arn          string `json:"arn"`
	UniqueId     string `json:"unique_id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	ForceDestroy bool   `json:"force_destroy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamUser `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorResourceGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsInspectorResourceGroupSpec `json"spec"`
}

type AwsInspectorResourceGroupSpec struct {
	Tags map[string]string `json:"tags"`
	Arn  string            `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorResourceGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsInspectorResourceGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultNetworkAcl struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDefaultNetworkAclSpec `json"spec"`
}

type AwsDefaultNetworkAclSpec struct {
	SubnetIds           string                          `json:"subnet_ids"`
	Ingress             AwsDefaultNetworkAclSpecIngress `json:"ingress"`
	Egress              AwsDefaultNetworkAclSpecEgress  `json:"egress"`
	Tags                map[string]string               `json:"tags"`
	VpcId               string                          `json:"vpc_id"`
	DefaultNetworkAclId string                          `json:"default_network_acl_id"`
}

type AwsDefaultNetworkAclSpecIngress struct {
	CidrBlock     string `json:"cidr_block"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	IcmpType      int    `json:"icmp_type"`
	IcmpCode      int    `json:"icmp_code"`
	FromPort      int    `json:"from_port"`
	ToPort        int    `json:"to_port"`
	RuleNo        int    `json:"rule_no"`
	Protocol      string `json:"protocol"`
	Action        string `json:"action"`
}

type AwsDefaultNetworkAclSpecEgress struct {
	RuleNo        int    `json:"rule_no"`
	Action        string `json:"action"`
	CidrBlock     string `json:"cidr_block"`
	IcmpType      int    `json:"icmp_type"`
	IcmpCode      int    `json:"icmp_code"`
	FromPort      int    `json:"from_port"`
	ToPort        int    `json:"to_port"`
	Protocol      string `json:"protocol"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultNetworkAclList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDefaultNetworkAcl `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheCluster struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticacheClusterSpec `json"spec"`
}

type AwsElasticacheClusterSpec struct {
	AzMode                 string                                `json:"az_mode"`
	AvailabilityZone       string                                `json:"availability_zone"`
	SnapshotWindow         string                                `json:"snapshot_window"`
	MaintenanceWindow      string                                `json:"maintenance_window"`
	SnapshotRetentionLimit int                                   `json:"snapshot_retention_limit"`
	SnapshotArns           string                                `json:"snapshot_arns"`
	NotificationTopicArn   string                                `json:"notification_topic_arn"`
	Tags                   map[string]string                     `json:"tags"`
	NumCacheNodes          int                                   `json:"num_cache_nodes"`
	ClusterAddress         string                                `json:"cluster_address"`
	Engine                 string                                `json:"engine"`
	SubnetGroupName        string                                `json:"subnet_group_name"`
	SecurityGroupIds       string                                `json:"security_group_ids"`
	ParameterGroupName     string                                `json:"parameter_group_name"`
	ClusterId              string                                `json:"cluster_id"`
	ConfigurationEndpoint  string                                `json:"configuration_endpoint"`
	ReplicationGroupId     string                                `json:"replication_group_id"`
	AvailabilityZones      string                                `json:"availability_zones"`
	NodeType               string                                `json:"node_type"`
	EngineVersion          string                                `json:"engine_version"`
	ApplyImmediately       bool                                  `json:"apply_immediately"`
	CacheNodes             []AwsElasticacheClusterSpecCacheNodes `json:"cache_nodes"`
	SecurityGroupNames     string                                `json:"security_group_names"`
	SnapshotName           string                                `json:"snapshot_name"`
	Port                   int                                   `json:"port"`
}

type AwsElasticacheClusterSpecCacheNodes struct {
	Id               string `json:"id"`
	Address          string `json:"address"`
	Port             int    `json:"port"`
	AvailabilityZone string `json:"availability_zone"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheClusterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticacheCluster `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamGroupSpec `json"spec"`
}

type AwsIamGroupSpec struct {
	Arn      string `json:"arn"`
	UniqueId string `json:"unique_id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRolePolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamRolePolicySpec `json"spec"`
}

type AwsIamRolePolicySpec struct {
	Policy     string `json:"policy"`
	Name       string `json:"name"`
	NamePrefix string `json:"name_prefix"`
	Role       string `json:"role"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRolePolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamRolePolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDirectoryServiceConditionalForwarder struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDirectoryServiceConditionalForwarderSpec `json"spec"`
}

type AwsDirectoryServiceConditionalForwarderSpec struct {
	DnsIps           []string `json:"dns_ips"`
	RemoteDomainName string   `json:"remote_domain_name"`
	DirectoryId      string   `json:"directory_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDirectoryServiceConditionalForwarderList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDirectoryServiceConditionalForwarder `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbTable struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDynamodbTableSpec `json"spec"`
}

type AwsDynamodbTableSpec struct {
	RangeKey             string                                     `json:"range_key"`
	Attribute            AwsDynamodbTableSpecAttribute              `json:"attribute"`
	LocalSecondaryIndex  AwsDynamodbTableSpecLocalSecondaryIndex    `json:"local_secondary_index"`
	StreamEnabled        bool                                       `json:"stream_enabled"`
	Name                 string                                     `json:"name"`
	HashKey              string                                     `json:"hash_key"`
	ReadCapacity         int                                        `json:"read_capacity"`
	StreamArn            string                                     `json:"stream_arn"`
	WriteCapacity        int                                        `json:"write_capacity"`
	GlobalSecondaryIndex AwsDynamodbTableSpecGlobalSecondaryIndex   `json:"global_secondary_index"`
	StreamViewType       string                                     `json:"stream_view_type"`
	StreamLabel          string                                     `json:"stream_label"`
	PointInTimeRecovery  []AwsDynamodbTableSpecPointInTimeRecovery  `json:"point_in_time_recovery"`
	Arn                  string                                     `json:"arn"`
	Ttl                  AwsDynamodbTableSpecTtl                    `json:"ttl"`
	ServerSideEncryption []AwsDynamodbTableSpecServerSideEncryption `json:"server_side_encryption"`
	Tags                 map[string]string                          `json:"tags"`
}

type AwsDynamodbTableSpecAttribute struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AwsDynamodbTableSpecLocalSecondaryIndex struct {
	ProjectionType   string   `json:"projection_type"`
	NonKeyAttributes []string `json:"non_key_attributes"`
	Name             string   `json:"name"`
	RangeKey         string   `json:"range_key"`
}

type AwsDynamodbTableSpecGlobalSecondaryIndex struct {
	Name             string   `json:"name"`
	WriteCapacity    int      `json:"write_capacity"`
	ReadCapacity     int      `json:"read_capacity"`
	HashKey          string   `json:"hash_key"`
	RangeKey         string   `json:"range_key"`
	ProjectionType   string   `json:"projection_type"`
	NonKeyAttributes []string `json:"non_key_attributes"`
}

type AwsDynamodbTableSpecPointInTimeRecovery struct {
	Enabled bool `json:"enabled"`
}

type AwsDynamodbTableSpecTtl struct {
	AttributeName string `json:"attribute_name"`
	Enabled       bool   `json:"enabled"`
}

type AwsDynamodbTableSpecServerSideEncryption struct {
	Enabled bool `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbTableList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDynamodbTable `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyMember struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGuarddutyMemberSpec `json"spec"`
}

type AwsGuarddutyMemberSpec struct {
	Email      string `json:"email"`
	AccountId  string `json:"account_id"`
	DetectorId string `json:"detector_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyMemberList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGuarddutyMember `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNatGateway struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsNatGatewaySpec `json"spec"`
}

type AwsNatGatewaySpec struct {
	NetworkInterfaceId string            `json:"network_interface_id"`
	PrivateIp          string            `json:"private_ip"`
	PublicIp           string            `json:"public_ip"`
	Tags               map[string]string `json:"tags"`
	AllocationId       string            `json:"allocation_id"`
	SubnetId           string            `json:"subnet_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNatGatewayList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsNatGateway `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSimpledbDomain struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSimpledbDomainSpec `json"spec"`
}

type AwsSimpledbDomainSpec struct {
	Name string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSimpledbDomainList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSimpledbDomain `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudformationStack struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudformationStackSpec `json"spec"`
}

type AwsCloudformationStackSpec struct {
	TemplateUrl      string            `json:"template_url"`
	Parameters       map[string]string `json:"parameters"`
	Outputs          map[string]string `json:"outputs"`
	Name             string            `json:"name"`
	Capabilities     string            `json:"capabilities"`
	NotificationArns string            `json:"notification_arns"`
	TemplateBody     string            `json:"template_body"`
	PolicyBody       string            `json:"policy_body"`
	PolicyUrl        string            `json:"policy_url"`
	Tags             map[string]string `json:"tags"`
	DisableRollback  bool              `json:"disable_rollback"`
	OnFailure        string            `json:"on_failure"`
	TimeoutInMinutes int               `json:"timeout_in_minutes"`
	IamRoleArn       string            `json:"iam_role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudformationStackList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudformationStack `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigConfigRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsConfigConfigRuleSpec `json"spec"`
}

type AwsConfigConfigRuleSpec struct {
	MaximumExecutionFrequency string                          `json:"maximum_execution_frequency"`
	Scope                     []AwsConfigConfigRuleSpecScope  `json:"scope"`
	Source                    []AwsConfigConfigRuleSpecSource `json:"source"`
	Name                      string                          `json:"name"`
	RuleId                    string                          `json:"rule_id"`
	Arn                       string                          `json:"arn"`
	Description               string                          `json:"description"`
	InputParameters           string                          `json:"input_parameters"`
}

type AwsConfigConfigRuleSpecScope struct {
	ComplianceResourceId    string `json:"compliance_resource_id"`
	ComplianceResourceTypes string `json:"compliance_resource_types"`
	TagKey                  string `json:"tag_key"`
	TagValue                string `json:"tag_value"`
}

type AwsConfigConfigRuleSpecSource struct {
	Owner            string                                    `json:"owner"`
	SourceDetail     AwsConfigConfigRuleSpecSourceSourceDetail `json:"source_detail"`
	SourceIdentifier string                                    `json:"source_identifier"`
}

type AwsConfigConfigRuleSpecSourceSourceDetail struct {
	MaximumExecutionFrequency string `json:"maximum_execution_frequency"`
	MessageType               string `json:"message_type"`
	EventSource               string `json:"event_source"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigConfigRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsConfigConfigRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDaxCluster struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDaxClusterSpec `json"spec"`
}

type AwsDaxClusterSpec struct {
	Arn                   string                   `json:"arn"`
	NodeType              string                   `json:"node_type"`
	NotificationTopicArn  string                   `json:"notification_topic_arn"`
	Tags                  map[string]string        `json:"tags"`
	IamRoleArn            string                   `json:"iam_role_arn"`
	SubnetGroupName       string                   `json:"subnet_group_name"`
	Port                  int                      `json:"port"`
	Nodes                 []AwsDaxClusterSpecNodes `json:"nodes"`
	ReplicationFactor     int                      `json:"replication_factor"`
	Description           string                   `json:"description"`
	ParameterGroupName    string                   `json:"parameter_group_name"`
	MaintenanceWindow     string                   `json:"maintenance_window"`
	ConfigurationEndpoint string                   `json:"configuration_endpoint"`
	ClusterName           string                   `json:"cluster_name"`
	AvailabilityZones     string                   `json:"availability_zones"`
	SecurityGroupIds      string                   `json:"security_group_ids"`
	ClusterAddress        string                   `json:"cluster_address"`
}

type AwsDaxClusterSpecNodes struct {
	Id               string `json:"id"`
	Address          string `json:"address"`
	Port             int    `json:"port"`
	AvailabilityZone string `json:"availability_zone"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDaxClusterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDaxCluster `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamInstanceProfile struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamInstanceProfileSpec `json"spec"`
}

type AwsIamInstanceProfileSpec struct {
	Role       string `json:"role"`
	Arn        string `json:"arn"`
	CreateDate string `json:"create_date"`
	UniqueId   string `json:"unique_id"`
	Name       string `json:"name"`
	NamePrefix string `json:"name_prefix"`
	Path       string `json:"path"`
	Roles      string `json:"roles"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamInstanceProfileList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamInstanceProfile `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksNodejsAppLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksNodejsAppLayerSpec `json"spec"`
}

type AwsOpsworksNodejsAppLayerSpec struct {
	AutoHealing              bool                                   `json:"auto_healing"`
	DrainElbOnShutdown       bool                                   `json:"drain_elb_on_shutdown"`
	CustomSetupRecipes       []string                               `json:"custom_setup_recipes"`
	ElasticLoadBalancer      string                                 `json:"elastic_load_balancer"`
	CustomDeployRecipes      []string                               `json:"custom_deploy_recipes"`
	NodejsVersion            string                                 `json:"nodejs_version"`
	CustomInstanceProfileArn string                                 `json:"custom_instance_profile_arn"`
	CustomUndeployRecipes    []string                               `json:"custom_undeploy_recipes"`
	CustomShutdownRecipes    []string                               `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds   string                                 `json:"custom_security_group_ids"`
	InstanceShutdownTimeout  int                                    `json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances bool                                   `json:"use_ebs_optimized_instances"`
	Name                     string                                 `json:"name"`
	AutoAssignElasticIps     bool                                   `json:"auto_assign_elastic_ips"`
	CustomConfigureRecipes   []string                               `json:"custom_configure_recipes"`
	CustomJson               string                                 `json:"custom_json"`
	InstallUpdatesOnBoot     bool                                   `json:"install_updates_on_boot"`
	SystemPackages           string                                 `json:"system_packages"`
	StackId                  string                                 `json:"stack_id"`
	EbsVolume                AwsOpsworksNodejsAppLayerSpecEbsVolume `json:"ebs_volume"`
	AutoAssignPublicIps      bool                                   `json:"auto_assign_public_ips"`
}

type AwsOpsworksNodejsAppLayerSpecEbsVolume struct {
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksNodejsAppLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksNodejsAppLayer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegration struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayIntegrationSpec `json"spec"`
}

type AwsApiGatewayIntegrationSpec struct {
	RequestParameters       map[string]string `json:"request_parameters"`
	IntegrationHttpMethod   string            `json:"integration_http_method"`
	RequestTemplates        map[string]string `json:"request_templates"`
	ContentHandling         string            `json:"content_handling"`
	CacheKeyParameters      string            `json:"cache_key_parameters"`
	RestApiId               string            `json:"rest_api_id"`
	ConnectionId            string            `json:"connection_id"`
	ResourceId              string            `json:"resource_id"`
	Type                    string            `json:"type"`
	ConnectionType          string            `json:"connection_type"`
	Uri                     string            `json:"uri"`
	Credentials             string            `json:"credentials"`
	RequestParametersInJson string            `json:"request_parameters_in_json"`
	PassthroughBehavior     string            `json:"passthrough_behavior"`
	CacheNamespace          string            `json:"cache_namespace"`
	HttpMethod              string            `json:"http_method"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegrationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayIntegration `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsEndpoint struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDmsEndpointSpec `json"spec"`
}

type AwsDmsEndpointSpec struct {
	ServerName                string                              `json:"server_name"`
	EndpointType              string                              `json:"endpoint_type"`
	ExtraConnectionAttributes string                              `json:"extra_connection_attributes"`
	KmsKeyArn                 string                              `json:"kms_key_arn"`
	SslMode                   string                              `json:"ssl_mode"`
	Username                  string                              `json:"username"`
	MongodbSettings           []AwsDmsEndpointSpecMongodbSettings `json:"mongodb_settings"`
	CertificateArn            string                              `json:"certificate_arn"`
	EndpointArn               string                              `json:"endpoint_arn"`
	Password                  string                              `json:"password"`
	Port                      int                                 `json:"port"`
	Tags                      map[string]string                   `json:"tags"`
	DatabaseName              string                              `json:"database_name"`
	EndpointId                string                              `json:"endpoint_id"`
	ServiceAccessRole         string                              `json:"service_access_role"`
	EngineName                string                              `json:"engine_name"`
}

type AwsDmsEndpointSpecMongodbSettings struct {
	DocsToInvestigate string `json:"docs_to_investigate"`
	AuthSource        string `json:"auth_source"`
	AuthType          string `json:"auth_type"`
	AuthMechanism     string `json:"auth_mechanism"`
	NestingLevel      string `json:"nesting_level"`
	ExtractDocId      string `json:"extract_doc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsEndpointList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDmsEndpoint `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccountAlias struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamAccountAliasSpec `json"spec"`
}

type AwsIamAccountAliasSpec struct {
	AccountAlias string `json:"account_alias"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccountAliasList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamAccountAlias `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInternetGateway struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsInternetGatewaySpec `json"spec"`
}

type AwsInternetGatewaySpec struct {
	VpcId string            `json:"vpc_id"`
	Tags  map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInternetGatewayList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsInternetGateway `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmDocument struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmDocumentSpec `json"spec"`
}

type AwsSsmDocumentSpec struct {
	Name           string                                   `json:"name"`
	CreatedDate    string                                   `json:"created_date"`
	Owner          string                                   `json:"owner"`
	Parameter      []AwsSsmDocumentSpecParameter            `json:"parameter"`
	Permissions    map[string]AwsSsmDocumentSpecPermissions `json:"permissions"`
	Arn            string                                   `json:"arn"`
	SchemaVersion  string                                   `json:"schema_version"`
	HashType       string                                   `json:"hash_type"`
	LatestVersion  string                                   `json:"latest_version"`
	PlatformTypes  []string                                 `json:"platform_types"`
	Content        string                                   `json:"content"`
	DocumentFormat string                                   `json:"document_format"`
	DocumentType   string                                   `json:"document_type"`
	DefaultVersion string                                   `json:"default_version"`
	Description    string                                   `json:"description"`
	Status         string                                   `json:"status"`
	Hash           string                                   `json:"hash"`
}

type AwsSsmDocumentSpecParameter struct {
	Name         string `json:"name"`
	DefaultValue string `json:"default_value"`
	Description  string `json:"description"`
	Type         string `json:"type"`
}

type AwsSsmDocumentSpecPermissions struct {
	Type       string `json:"type"`
	AccountIds string `json:"account_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmDocumentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmDocument `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLb struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLbSpec `json"spec"`
}

type AwsLbSpec struct {
	SecurityGroups               string                 `json:"security_groups"`
	EnableHttp2                  bool                   `json:"enable_http2"`
	VpcId                        string                 `json:"vpc_id"`
	ZoneId                       string                 `json:"zone_id"`
	DnsName                      string                 `json:"dns_name"`
	Name                         string                 `json:"name"`
	NamePrefix                   string                 `json:"name_prefix"`
	Internal                     bool                   `json:"internal"`
	IpAddressType                string                 `json:"ip_address_type"`
	Tags                         map[string]string      `json:"tags"`
	ArnSuffix                    string                 `json:"arn_suffix"`
	AccessLogs                   []AwsLbSpecAccessLogs  `json:"access_logs"`
	EnableCrossZoneLoadBalancing bool                   `json:"enable_cross_zone_load_balancing"`
	SubnetMapping                AwsLbSpecSubnetMapping `json:"subnet_mapping"`
	EnableDeletionProtection     bool                   `json:"enable_deletion_protection"`
	IdleTimeout                  int                    `json:"idle_timeout"`
	Arn                          string                 `json:"arn"`
	LoadBalancerType             string                 `json:"load_balancer_type"`
	Subnets                      string                 `json:"subnets"`
}

type AwsLbSpecAccessLogs struct {
	Bucket  string `json:"bucket"`
	Prefix  string `json:"prefix"`
	Enabled bool   `json:"enabled"`
}

type AwsLbSpecSubnetMapping struct {
	SubnetId     string `json:"subnet_id"`
	AllocationId string `json:"allocation_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLb `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalGeoMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalGeoMatchSetSpec `json"spec"`
}

type AwsWafregionalGeoMatchSetSpec struct {
	Name               string                                          `json:"name"`
	GeoMatchConstraint AwsWafregionalGeoMatchSetSpecGeoMatchConstraint `json:"geo_match_constraint"`
}

type AwsWafregionalGeoMatchSetSpecGeoMatchConstraint struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalGeoMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalGeoMatchSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingScheduledAction struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAppautoscalingScheduledActionSpec `json"spec"`
}

type AwsAppautoscalingScheduledActionSpec struct {
	ServiceNamespace     string                                                     `json:"service_namespace"`
	ScalableTargetAction []AwsAppautoscalingScheduledActionSpecScalableTargetAction `json:"scalable_target_action"`
	StartTime            string                                                     `json:"start_time"`
	Arn                  string                                                     `json:"arn"`
	Name                 string                                                     `json:"name"`
	ScalableDimension    string                                                     `json:"scalable_dimension"`
	Schedule             string                                                     `json:"schedule"`
	EndTime              string                                                     `json:"end_time"`
	ResourceId           string                                                     `json:"resource_id"`
}

type AwsAppautoscalingScheduledActionSpecScalableTargetAction struct {
	MaxCapacity int `json:"max_capacity"`
	MinCapacity int `json:"min_capacity"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingScheduledActionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAppautoscalingScheduledAction `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheReplicationGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticacheReplicationGroupSpec `json"spec"`
}

type AwsElasticacheReplicationGroupSpec struct {
	AutomaticFailoverEnabled     bool                                            `json:"automatic_failover_enabled"`
	ReplicationGroupDescription  string                                          `json:"replication_group_description"`
	ConfigurationEndpointAddress string                                          `json:"configuration_endpoint_address"`
	SecurityGroupIds             string                                          `json:"security_group_ids"`
	SnapshotWindow               string                                          `json:"snapshot_window"`
	NotificationTopicArn         string                                          `json:"notification_topic_arn"`
	Engine                       string                                          `json:"engine"`
	SnapshotName                 string                                          `json:"snapshot_name"`
	AvailabilityZones            string                                          `json:"availability_zones"`
	AtRestEncryptionEnabled      bool                                            `json:"at_rest_encryption_enabled"`
	ReplicationGroupId           string                                          `json:"replication_group_id"`
	TransitEncryptionEnabled     bool                                            `json:"transit_encryption_enabled"`
	NodeType                     string                                          `json:"node_type"`
	SnapshotRetentionLimit       int                                             `json:"snapshot_retention_limit"`
	Tags                         map[string]string                               `json:"tags"`
	AuthToken                    string                                          `json:"auth_token"`
	ParameterGroupName           string                                          `json:"parameter_group_name"`
	SnapshotArns                 string                                          `json:"snapshot_arns"`
	NumberCacheClusters          int                                             `json:"number_cache_clusters"`
	PrimaryEndpointAddress       string                                          `json:"primary_endpoint_address"`
	ClusterMode                  []AwsElasticacheReplicationGroupSpecClusterMode `json:"cluster_mode"`
	SecurityGroupNames           string                                          `json:"security_group_names"`
	Port                         int                                             `json:"port"`
	AutoMinorVersionUpgrade      bool                                            `json:"auto_minor_version_upgrade"`
	ApplyImmediately             bool                                            `json:"apply_immediately"`
	EngineVersion                string                                          `json:"engine_version"`
	SubnetGroupName              string                                          `json:"subnet_group_name"`
	MaintenanceWindow            string                                          `json:"maintenance_window"`
}

type AwsElasticacheReplicationGroupSpecClusterMode struct {
	ReplicasPerNodeGroup int `json:"replicas_per_node_group"`
	NumNodeGroups        int `json:"num_node_groups"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheReplicationGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticacheReplicationGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGameliftFleet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGameliftFleetSpec `json"spec"`
}

type AwsGameliftFleetSpec struct {
	OperatingSystem                string                                            `json:"operating_system"`
	ResourceCreationLimitPolicy    []AwsGameliftFleetSpecResourceCreationLimitPolicy `json:"resource_creation_limit_policy"`
	LogPaths                       []string                                          `json:"log_paths"`
	NewGameSessionProtectionPolicy string                                            `json:"new_game_session_protection_policy"`
	Ec2InstanceType                string                                            `json:"ec2_instance_type"`
	Name                           string                                            `json:"name"`
	Description                    string                                            `json:"description"`
	Ec2InboundPermission           []AwsGameliftFleetSpecEc2InboundPermission        `json:"ec2_inbound_permission"`
	MetricGroups                   []string                                          `json:"metric_groups"`
	RuntimeConfiguration           []AwsGameliftFleetSpecRuntimeConfiguration        `json:"runtime_configuration"`
	Arn                            string                                            `json:"arn"`
	BuildId                        string                                            `json:"build_id"`
}

type AwsGameliftFleetSpecResourceCreationLimitPolicy struct {
	NewGameSessionsPerCreator int `json:"new_game_sessions_per_creator"`
	PolicyPeriodInMinutes     int `json:"policy_period_in_minutes"`
}

type AwsGameliftFleetSpecEc2InboundPermission struct {
	ToPort   int    `json:"to_port"`
	FromPort int    `json:"from_port"`
	IpRange  string `json:"ip_range"`
	Protocol string `json:"protocol"`
}

type AwsGameliftFleetSpecRuntimeConfiguration struct {
	GameSessionActivationTimeoutSeconds int                                                     `json:"game_session_activation_timeout_seconds"`
	MaxConcurrentGameSessionActivations int                                                     `json:"max_concurrent_game_session_activations"`
	ServerProcess                       []AwsGameliftFleetSpecRuntimeConfigurationServerProcess `json:"server_process"`
}

type AwsGameliftFleetSpecRuntimeConfigurationServerProcess struct {
	ConcurrentExecutions int    `json:"concurrent_executions"`
	LaunchPath           string `json:"launch_path"`
	Parameters           string `json:"parameters"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGameliftFleetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGameliftFleet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftCluster struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRedshiftClusterSpec `json"spec"`
}

type AwsRedshiftClusterSpec struct {
	ClusterParameterGroupName        string                               `json:"cluster_parameter_group_name"`
	NumberOfNodes                    int                                  `json:"number_of_nodes"`
	PubliclyAccessible               bool                                 `json:"publicly_accessible"`
	FinalSnapshotIdentifier          string                               `json:"final_snapshot_identifier"`
	ClusterRevisionNumber            string                               `json:"cluster_revision_number"`
	ClusterIdentifier                string                               `json:"cluster_identifier"`
	ClusterPublicKey                 string                               `json:"cluster_public_key"`
	S3KeyPrefix                      string                               `json:"s3_key_prefix"`
	ClusterType                      string                               `json:"cluster_type"`
	ClusterVersion                   string                               `json:"cluster_version"`
	IamRoles                         string                               `json:"iam_roles"`
	VpcSecurityGroupIds              string                               `json:"vpc_security_group_ids"`
	SnapshotCopy                     []AwsRedshiftClusterSpecSnapshotCopy `json:"snapshot_copy"`
	Logging                          []AwsRedshiftClusterSpecLogging      `json:"logging"`
	EnableLogging                    bool                                 `json:"enable_logging"`
	MasterPassword                   string                               `json:"master_password"`
	ClusterSubnetGroupName           string                               `json:"cluster_subnet_group_name"`
	AvailabilityZone                 string                               `json:"availability_zone"`
	EnhancedVpcRouting               bool                                 `json:"enhanced_vpc_routing"`
	ElasticIp                        string                               `json:"elastic_ip"`
	SnapshotClusterIdentifier        string                               `json:"snapshot_cluster_identifier"`
	OwnerAccount                     string                               `json:"owner_account"`
	NodeType                         string                               `json:"node_type"`
	AllowVersionUpgrade              bool                                 `json:"allow_version_upgrade"`
	Encrypted                        bool                                 `json:"encrypted"`
	DatabaseName                     string                               `json:"database_name"`
	AutomatedSnapshotRetentionPeriod int                                  `json:"automated_snapshot_retention_period"`
	Port                             int                                  `json:"port"`
	SnapshotIdentifier               string                               `json:"snapshot_identifier"`
	Tags                             map[string]string                    `json:"tags"`
	MasterUsername                   string                               `json:"master_username"`
	PreferredMaintenanceWindow       string                               `json:"preferred_maintenance_window"`
	KmsKeyId                         string                               `json:"kms_key_id"`
	SkipFinalSnapshot                bool                                 `json:"skip_final_snapshot"`
	Endpoint                         string                               `json:"endpoint"`
	BucketName                       string                               `json:"bucket_name"`
	ClusterSecurityGroups            string                               `json:"cluster_security_groups"`
}

type AwsRedshiftClusterSpecSnapshotCopy struct {
	DestinationRegion string `json:"destination_region"`
	RetentionPeriod   int    `json:"retention_period"`
	GrantName         string `json:"grant_name"`
}

type AwsRedshiftClusterSpecLogging struct {
	Enable      bool   `json:"enable"`
	BucketName  string `json:"bucket_name"`
	S3KeyPrefix string `json:"s3_key_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftClusterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRedshiftCluster `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptRuleSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesReceiptRuleSetSpec `json"spec"`
}

type AwsSesReceiptRuleSetSpec struct {
	RuleSetName string `json:"rule_set_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptRuleSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesReceiptRuleSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafIpset struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafIpsetSpec `json"spec"`
}

type AwsWafIpsetSpec struct {
	Name             string                          `json:"name"`
	IpSetDescriptors AwsWafIpsetSpecIpSetDescriptors `json:"ip_set_descriptors"`
}

type AwsWafIpsetSpecIpSetDescriptors struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafIpsetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafIpset `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsClusterParameterGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRdsClusterParameterGroupSpec `json"spec"`
}

type AwsRdsClusterParameterGroupSpec struct {
	Tags        map[string]string                        `json:"tags"`
	Arn         string                                   `json:"arn"`
	Name        string                                   `json:"name"`
	NamePrefix  string                                   `json:"name_prefix"`
	Family      string                                   `json:"family"`
	Description string                                   `json:"description"`
	Parameter   AwsRdsClusterParameterGroupSpecParameter `json:"parameter"`
}

type AwsRdsClusterParameterGroupSpecParameter struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	ApplyMethod string `json:"apply_method"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsClusterParameterGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRdsClusterParameterGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsTopic struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSnsTopicSpec `json"spec"`
}

type AwsSnsTopicSpec struct {
	SqsFailureFeedbackRoleArn            string `json:"sqs_failure_feedback_role_arn"`
	HttpSuccessFeedbackSampleRate        int    `json:"http_success_feedback_sample_rate"`
	LambdaSuccessFeedbackRoleArn         string `json:"lambda_success_feedback_role_arn"`
	SqsSuccessFeedbackRoleArn            string `json:"sqs_success_feedback_role_arn"`
	NamePrefix                           string `json:"name_prefix"`
	SqsSuccessFeedbackSampleRate         int    `json:"sqs_success_feedback_sample_rate"`
	Arn                                  string `json:"arn"`
	ApplicationSuccessFeedbackSampleRate int    `json:"application_success_feedback_sample_rate"`
	HttpFailureFeedbackRoleArn           string `json:"http_failure_feedback_role_arn"`
	Name                                 string `json:"name"`
	Policy                               string `json:"policy"`
	ApplicationSuccessFeedbackRoleArn    string `json:"application_success_feedback_role_arn"`
	HttpSuccessFeedbackRoleArn           string `json:"http_success_feedback_role_arn"`
	LambdaSuccessFeedbackSampleRate      int    `json:"lambda_success_feedback_sample_rate"`
	LambdaFailureFeedbackRoleArn         string `json:"lambda_failure_feedback_role_arn"`
	DisplayName                          string `json:"display_name"`
	DeliveryPolicy                       string `json:"delivery_policy"`
	ApplicationFailureFeedbackRoleArn    string `json:"application_failure_feedback_role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsTopicList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSnsTopic `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAcmCertificate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAcmCertificateSpec `json"spec"`
}

type AwsAcmCertificateSpec struct {
	SubjectAlternativeNames []string                                       `json:"subject_alternative_names"`
	ValidationMethod        string                                         `json:"validation_method"`
	Arn                     string                                         `json:"arn"`
	DomainValidationOptions []AwsAcmCertificateSpecDomainValidationOptions `json:"domain_validation_options"`
	ValidationEmails        []string                                       `json:"validation_emails"`
	Tags                    map[string]string                              `json:"tags"`
	DomainName              string                                         `json:"domain_name"`
}

type AwsAcmCertificateSpecDomainValidationOptions struct {
	DomainName          string `json:"domain_name"`
	ResourceRecordName  string `json:"resource_record_name"`
	ResourceRecordType  string `json:"resource_record_type"`
	ResourceRecordValue string `json:"resource_record_value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAcmCertificateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAcmCertificate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamPolicySpec `json"spec"`
}

type AwsIamPolicySpec struct {
	Arn         string `json:"arn"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Policy      string `json:"policy"`
	Name        string `json:"name"`
	NamePrefix  string `json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksJavaAppLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksJavaAppLayerSpec `json"spec"`
}

type AwsOpsworksJavaAppLayerSpec struct {
	JvmType                  string                               `json:"jvm_type"`
	AutoAssignElasticIps     bool                                 `json:"auto_assign_elastic_ips"`
	CustomSetupRecipes       []string                             `json:"custom_setup_recipes"`
	CustomDeployRecipes      []string                             `json:"custom_deploy_recipes"`
	CustomShutdownRecipes    []string                             `json:"custom_shutdown_recipes"`
	CustomJson               string                               `json:"custom_json"`
	DrainElbOnShutdown       bool                                 `json:"drain_elb_on_shutdown"`
	StackId                  string                               `json:"stack_id"`
	EbsVolume                AwsOpsworksJavaAppLayerSpecEbsVolume `json:"ebs_volume"`
	AutoAssignPublicIps      bool                                 `json:"auto_assign_public_ips"`
	CustomInstanceProfileArn string                               `json:"custom_instance_profile_arn"`
	ElasticLoadBalancer      string                               `json:"elastic_load_balancer"`
	CustomConfigureRecipes   []string                             `json:"custom_configure_recipes"`
	InstallUpdatesOnBoot     bool                                 `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  int                                  `json:"instance_shutdown_timeout"`
	JvmVersion               string                               `json:"jvm_version"`
	JvmOptions               string                               `json:"jvm_options"`
	CustomUndeployRecipes    []string                             `json:"custom_undeploy_recipes"`
	SystemPackages           string                               `json:"system_packages"`
	Name                     string                               `json:"name"`
	CustomSecurityGroupIds   string                               `json:"custom_security_group_ids"`
	AutoHealing              bool                                 `json:"auto_healing"`
	UseEbsOptimizedInstances bool                                 `json:"use_ebs_optimized_instances"`
	AppServer                string                               `json:"app_server"`
	AppServerVersion         string                               `json:"app_server_version"`
}

type AwsOpsworksJavaAppLayerSpecEbsVolume struct {
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksJavaAppLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksJavaAppLayer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmResourceDataSync struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmResourceDataSyncSpec `json"spec"`
}

type AwsSsmResourceDataSyncSpec struct {
	Name          string                                    `json:"name"`
	S3Destination []AwsSsmResourceDataSyncSpecS3Destination `json:"s3_destination"`
}

type AwsSsmResourceDataSyncSpecS3Destination struct {
	BucketName string `json:"bucket_name"`
	Prefix     string `json:"prefix"`
	Region     string `json:"region"`
	SyncFormat string `json:"sync_format"`
	KmsKeyArn  string `json:"kms_key_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmResourceDataSyncList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmResourceDataSync `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcPeeringConnectionAccepter struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcPeeringConnectionAccepterSpec `json"spec"`
}

type AwsVpcPeeringConnectionAccepterSpec struct {
	AcceptStatus           string                                       `json:"accept_status"`
	PeerVpcId              string                                       `json:"peer_vpc_id"`
	Accepter               AwsVpcPeeringConnectionAccepterSpecAccepter  `json:"accepter"`
	Requester              AwsVpcPeeringConnectionAccepterSpecRequester `json:"requester"`
	Tags                   map[string]string                            `json:"tags"`
	VpcPeeringConnectionId string                                       `json:"vpc_peering_connection_id"`
	VpcId                  string                                       `json:"vpc_id"`
	PeerOwnerId            string                                       `json:"peer_owner_id"`
	PeerRegion             string                                       `json:"peer_region"`
	AutoAccept             bool                                         `json:"auto_accept"`
}

type AwsVpcPeeringConnectionAccepterSpecAccepter struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type AwsVpcPeeringConnectionAccepterSpecRequester struct {
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

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayUsagePlan struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayUsagePlanSpec `json"spec"`
}

type AwsApiGatewayUsagePlanSpec struct {
	Description      string                                     `json:"description"`
	ApiStages        []AwsApiGatewayUsagePlanSpecApiStages      `json:"api_stages"`
	QuotaSettings    AwsApiGatewayUsagePlanSpecQuotaSettings    `json:"quota_settings"`
	ThrottleSettings AwsApiGatewayUsagePlanSpecThrottleSettings `json:"throttle_settings"`
	ProductCode      string                                     `json:"product_code"`
	Name             string                                     `json:"name"`
}

type AwsApiGatewayUsagePlanSpecApiStages struct {
	ApiId string `json:"api_id"`
	Stage string `json:"stage"`
}

type AwsApiGatewayUsagePlanSpecQuotaSettings struct {
	Period string `json:"period"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type AwsApiGatewayUsagePlanSpecThrottleSettings struct {
	BurstLimit int     `json:"burst_limit"`
	RateLimit  float64 `json:"rate_limit"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayUsagePlanList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayUsagePlan `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRolePolicyAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamRolePolicyAttachmentSpec `json"spec"`
}

type AwsIamRolePolicyAttachmentSpec struct {
	Role      string `json:"role"`
	PolicyArn string `json:"policy_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRolePolicyAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamRolePolicyAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLoadBalancerPolicySpec `json"spec"`
}

type AwsLoadBalancerPolicySpec struct {
	PolicyAttribute  AwsLoadBalancerPolicySpecPolicyAttribute `json:"policy_attribute"`
	LoadBalancerName string                                   `json:"load_balancer_name"`
	PolicyName       string                                   `json:"policy_name"`
	PolicyTypeName   string                                   `json:"policy_type_name"`
}

type AwsLoadBalancerPolicySpecPolicyAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLoadBalancerPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmPatchGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmPatchGroupSpec `json"spec"`
}

type AwsSsmPatchGroupSpec struct {
	BaselineId string `json:"baseline_id"`
	PatchGroup string `json:"patch_group"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmPatchGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmPatchGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmParameter struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmParameterSpec `json"spec"`
}

type AwsSsmParameterSpec struct {
	Type           string            `json:"type"`
	Arn            string            `json:"arn"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	Overwrite      bool              `json:"overwrite"`
	AllowedPattern string            `json:"allowed_pattern"`
	Tags           map[string]string `json:"tags"`
	Value          string            `json:"value"`
	KeyId          string            `json:"key_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmParameterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmParameter `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpoint struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcEndpointSpec `json"spec"`
}

type AwsVpcEndpointSpec struct {
	ServiceName         string                       `json:"service_name"`
	Policy              string                       `json:"policy"`
	SubnetIds           string                       `json:"subnet_ids"`
	State               string                       `json:"state"`
	VpcId               string                       `json:"vpc_id"`
	VpcEndpointType     string                       `json:"vpc_endpoint_type"`
	RouteTableIds       string                       `json:"route_table_ids"`
	NetworkInterfaceIds string                       `json:"network_interface_ids"`
	PrefixListId        string                       `json:"prefix_list_id"`
	CidrBlocks          []string                     `json:"cidr_blocks"`
	AutoAccept          bool                         `json:"auto_accept"`
	SecurityGroupIds    string                       `json:"security_group_ids"`
	PrivateDnsEnabled   bool                         `json:"private_dns_enabled"`
	DnsEntry            []AwsVpcEndpointSpecDnsEntry `json:"dns_entry"`
}

type AwsVpcEndpointSpecDnsEntry struct {
	DnsName      string `json:"dns_name"`
	HostedZoneId string `json:"hosted_zone_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcEndpoint `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointService struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcEndpointServiceSpec `json"spec"`
}

type AwsVpcEndpointServiceSpec struct {
	AcceptanceRequired      bool   `json:"acceptance_required"`
	NetworkLoadBalancerArns string `json:"network_load_balancer_arns"`
	PrivateDnsName          string `json:"private_dns_name"`
	BaseEndpointDnsNames    string `json:"base_endpoint_dns_names"`
	AvailabilityZones       string `json:"availability_zones"`
	AllowedPrincipals       string `json:"allowed_principals"`
	State                   string `json:"state"`
	ServiceName             string `json:"service_name"`
	ServiceType             string `json:"service_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointServiceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcEndpointService `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchComputeEnvironment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsBatchComputeEnvironmentSpec `json"spec"`
}

type AwsBatchComputeEnvironmentSpec struct {
	StatusReason           string                                           `json:"status_reason"`
	State                  string                                           `json:"state"`
	Type                   string                                           `json:"type"`
	Arn                    string                                           `json:"arn"`
	EccClusterArn          string                                           `json:"ecc_cluster_arn"`
	Status                 string                                           `json:"status"`
	ComputeEnvironmentName string                                           `json:"compute_environment_name"`
	ComputeResources       []AwsBatchComputeEnvironmentSpecComputeResources `json:"compute_resources"`
	ServiceRole            string                                           `json:"service_role"`
	EcsClusterArn          string                                           `json:"ecs_cluster_arn"`
}

type AwsBatchComputeEnvironmentSpecComputeResources struct {
	MinVcpus         int               `json:"min_vcpus"`
	SecurityGroupIds string            `json:"security_group_ids"`
	SpotIamFleetRole string            `json:"spot_iam_fleet_role"`
	BidPercentage    int               `json:"bid_percentage"`
	DesiredVcpus     int               `json:"desired_vcpus"`
	Ec2KeyPair       string            `json:"ec2_key_pair"`
	ImageId          string            `json:"image_id"`
	InstanceRole     string            `json:"instance_role"`
	Subnets          string            `json:"subnets"`
	InstanceType     string            `json:"instance_type"`
	MaxVcpus         int               `json:"max_vcpus"`
	Tags             map[string]string `json:"tags"`
	Type             string            `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchComputeEnvironmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsBatchComputeEnvironment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontOriginAccessIdentity struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudfrontOriginAccessIdentitySpec `json"spec"`
}

type AwsCloudfrontOriginAccessIdentitySpec struct {
	Comment                      string `json:"comment"`
	CallerReference              string `json:"caller_reference"`
	CloudfrontAccessIdentityPath string `json:"cloudfront_access_identity_path"`
	Etag                         string `json:"etag"`
	IamArn                       string `json:"iam_arn"`
	S3CanonicalUserId            string `json:"s3_canonical_user_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontOriginAccessIdentityList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudfrontOriginAccessIdentity `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbEventSubscription struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDbEventSubscriptionSpec `json"spec"`
}

type AwsDbEventSubscriptionSpec struct {
	Arn             string            `json:"arn"`
	SnsTopic        string            `json:"sns_topic"`
	SourceIds       string            `json:"source_ids"`
	Tags            map[string]string `json:"tags"`
	CustomerAwsId   string            `json:"customer_aws_id"`
	Name            string            `json:"name"`
	EventCategories string            `json:"event_categories"`
	SourceType      string            `json:"source_type"`
	Enabled         bool              `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbEventSubscriptionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDbEventSubscription `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftSecurityGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRedshiftSecurityGroupSpec `json"spec"`
}

type AwsRedshiftSecurityGroupSpec struct {
	Name        string                              `json:"name"`
	Description string                              `json:"description"`
	Ingress     AwsRedshiftSecurityGroupSpecIngress `json:"ingress"`
}

type AwsRedshiftSecurityGroupSpecIngress struct {
	Cidr                 string `json:"cidr"`
	SecurityGroupName    string `json:"security_group_name"`
	SecurityGroupOwnerId string `json:"security_group_owner_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftSecurityGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRedshiftSecurityGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRouteTable struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRouteTableSpec `json"spec"`
}

type AwsRouteTableSpec struct {
	VpcId           string                 `json:"vpc_id"`
	Tags            map[string]string      `json:"tags"`
	PropagatingVgws string                 `json:"propagating_vgws"`
	Route           AwsRouteTableSpecRoute `json:"route"`
}

type AwsRouteTableSpecRoute struct {
	Ipv6CidrBlock          string `json:"ipv6_cidr_block"`
	EgressOnlyGatewayId    string `json:"egress_only_gateway_id"`
	GatewayId              string `json:"gateway_id"`
	InstanceId             string `json:"instance_id"`
	NatGatewayId           string `json:"nat_gateway_id"`
	VpcPeeringConnectionId string `json:"vpc_peering_connection_id"`
	NetworkInterfaceId     string `json:"network_interface_id"`
	CidrBlock              string `json:"cidr_block"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRouteTableList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRouteTable `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRouteTableAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRouteTableAssociationSpec `json"spec"`
}

type AwsRouteTableAssociationSpec struct {
	SubnetId     string `json:"subnet_id"`
	RouteTableId string `json:"route_table_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRouteTableAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRouteTableAssociation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainDkim struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesDomainDkimSpec `json"spec"`
}

type AwsSesDomainDkimSpec struct {
	Domain     string   `json:"domain"`
	DkimTokens []string `json:"dkim_tokens"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainDkimList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesDomainDkim `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsPlatformApplication struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSnsPlatformApplicationSpec `json"spec"`
}

type AwsSnsPlatformApplicationSpec struct {
	Name                         string `json:"name"`
	EventEndpointCreatedTopicArn string `json:"event_endpoint_created_topic_arn"`
	EventEndpointUpdatedTopicArn string `json:"event_endpoint_updated_topic_arn"`
	FailureFeedbackRoleArn       string `json:"failure_feedback_role_arn"`
	PlatformPrincipal            string `json:"platform_principal"`
	SuccessFeedbackSampleRate    string `json:"success_feedback_sample_rate"`
	Platform                     string `json:"platform"`
	PlatformCredential           string `json:"platform_credential"`
	Arn                          string `json:"arn"`
	EventDeliveryFailureTopicArn string `json:"event_delivery_failure_topic_arn"`
	EventEndpointDeletedTopicArn string `json:"event_endpoint_deleted_topic_arn"`
	SuccessFeedbackRoleArn       string `json:"success_feedback_role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsPlatformApplicationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSnsPlatformApplication `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchEventRuleSpec `json"spec"`
}

type AwsCloudwatchEventRuleSpec struct {
	ScheduleExpression string `json:"schedule_expression"`
	EventPattern       string `json:"event_pattern"`
	Description        string `json:"description"`
	RoleArn            string `json:"role_arn"`
	IsEnabled          bool   `json:"is_enabled"`
	Arn                string `json:"arn"`
	Name               string `json:"name"`
	NamePrefix         string `json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchEventRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogSubscriptionFilter struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchLogSubscriptionFilterSpec `json"spec"`
}

type AwsCloudwatchLogSubscriptionFilterSpec struct {
	FilterPattern  string `json:"filter_pattern"`
	LogGroupName   string `json:"log_group_name"`
	RoleArn        string `json:"role_arn"`
	Distribution   string `json:"distribution"`
	Name           string `json:"name"`
	DestinationArn string `json:"destination_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogSubscriptionFilterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchLogSubscriptionFilter `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamServiceLinkedRole struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamServiceLinkedRoleSpec `json"spec"`
}

type AwsIamServiceLinkedRoleSpec struct {
	CustomSuffix   string `json:"custom_suffix"`
	Description    string `json:"description"`
	AwsServiceName string `json:"aws_service_name"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	Arn            string `json:"arn"`
	CreateDate     string `json:"create_date"`
	UniqueId       string `json:"unique_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamServiceLinkedRoleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamServiceLinkedRole `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIotPolicySpec `json"spec"`
}

type AwsIotPolicySpec struct {
	DefaultVersionId string `json:"default_version_id"`
	Name             string `json:"name"`
	Policy           string `json:"policy"`
	Arn              string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIotPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksRailsAppLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksRailsAppLayerSpec `json"spec"`
}

type AwsOpsworksRailsAppLayerSpec struct {
	ElasticLoadBalancer      string                                `json:"elastic_load_balancer"`
	CustomShutdownRecipes    []string                              `json:"custom_shutdown_recipes"`
	AutoHealing              bool                                  `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                                  `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  int                                   `json:"instance_shutdown_timeout"`
	SystemPackages           string                                `json:"system_packages"`
	EbsVolume                AwsOpsworksRailsAppLayerSpecEbsVolume `json:"ebs_volume"`
	AutoAssignElasticIps     bool                                  `json:"auto_assign_elastic_ips"`
	RubygemsVersion          string                                `json:"rubygems_version"`
	BundlerVersion           string                                `json:"bundler_version"`
	PassengerVersion         string                                `json:"passenger_version"`
	CustomSecurityGroupIds   string                                `json:"custom_security_group_ids"`
	UseEbsOptimizedInstances bool                                  `json:"use_ebs_optimized_instances"`
	CustomConfigureRecipes   []string                              `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                              `json:"custom_deploy_recipes"`
	CustomJson               string                                `json:"custom_json"`
	Name                     string                                `json:"name"`
	ManageBundler            bool                                  `json:"manage_bundler"`
	RubyVersion              string                                `json:"ruby_version"`
	AppServer                string                                `json:"app_server"`
	AutoAssignPublicIps      bool                                  `json:"auto_assign_public_ips"`
	CustomSetupRecipes       []string                              `json:"custom_setup_recipes"`
	CustomUndeployRecipes    []string                              `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool                                  `json:"drain_elb_on_shutdown"`
	StackId                  string                                `json:"stack_id"`
	CustomInstanceProfileArn string                                `json:"custom_instance_profile_arn"`
}

type AwsOpsworksRailsAppLayerSpecEbsVolume struct {
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksRailsAppLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksRailsAppLayer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksUserProfile struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksUserProfileSpec `json"spec"`
}

type AwsOpsworksUserProfileSpec struct {
	SshUsername         string `json:"ssh_username"`
	SshPublicKey        string `json:"ssh_public_key"`
	UserArn             string `json:"user_arn"`
	AllowSelfManagement bool   `json:"allow_self_management"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksUserProfileList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksUserProfile `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsGrant struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsKmsGrantSpec `json"spec"`
}

type AwsKmsGrantSpec struct {
	KeyId               string                     `json:"key_id"`
	GranteePrincipal    string                     `json:"grantee_principal"`
	Operations          string                     `json:"operations"`
	Constraints         AwsKmsGrantSpecConstraints `json:"constraints"`
	RetiringPrincipal   string                     `json:"retiring_principal"`
	GrantCreationTokens string                     `json:"grant_creation_tokens"`
	RetireOnDelete      bool                       `json:"retire_on_delete"`
	Name                string                     `json:"name"`
	GrantId             string                     `json:"grant_id"`
	GrantToken          string                     `json:"grant_token"`
}

type AwsKmsGrantSpecConstraints struct {
	EncryptionContextSubset map[string]string `json:"encryption_context_subset"`
	EncryptionContextEquals map[string]string `json:"encryption_context_equals"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsGrantList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsKmsGrant `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultSecurityGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDefaultSecurityGroupSpec `json"spec"`
}

type AwsDefaultSecurityGroupSpec struct {
	Ingress             AwsDefaultSecurityGroupSpecIngress `json:"ingress"`
	Egress              AwsDefaultSecurityGroupSpecEgress  `json:"egress"`
	Tags                map[string]string                  `json:"tags"`
	Arn                 string                             `json:"arn"`
	OwnerId             string                             `json:"owner_id"`
	RevokeRulesOnDelete bool                               `json:"revoke_rules_on_delete"`
	Name                string                             `json:"name"`
	VpcId               string                             `json:"vpc_id"`
}

type AwsDefaultSecurityGroupSpecIngress struct {
	FromPort       int      `json:"from_port"`
	ToPort         int      `json:"to_port"`
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	SecurityGroups string   `json:"security_groups"`
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
}

type AwsDefaultSecurityGroupSpecEgress struct {
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
	Protocol       string   `json:"protocol"`
	CidrBlocks     []string `json:"cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	SecurityGroups string   `json:"security_groups"`
	Self           bool     `json:"self"`
	ToPort         int      `json:"to_port"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultSecurityGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDefaultSecurityGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethod struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayMethodSpec `json"spec"`
}

type AwsApiGatewayMethodSpec struct {
	RestApiId               string            `json:"rest_api_id"`
	ResourceId              string            `json:"resource_id"`
	RequestParametersInJson string            `json:"request_parameters_in_json"`
	RequestParameters       map[string]string `json:"request_parameters"`
	RequestValidatorId      string            `json:"request_validator_id"`
	HttpMethod              string            `json:"http_method"`
	Authorization           string            `json:"authorization"`
	AuthorizerId            string            `json:"authorizer_id"`
	ApiKeyRequired          bool              `json:"api_key_required"`
	RequestModels           map[string]string `json:"request_models"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethodList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayMethod `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingTarget struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAppautoscalingTargetSpec `json"spec"`
}

type AwsAppautoscalingTargetSpec struct {
	MaxCapacity       int    `json:"max_capacity"`
	MinCapacity       int    `json:"min_capacity"`
	ResourceId        string `json:"resource_id"`
	RoleArn           string `json:"role_arn"`
	ScalableDimension string `json:"scalable_dimension"`
	ServiceNamespace  string `json:"service_namespace"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingTargetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAppautoscalingTarget `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAutoscalingPolicySpec `json"spec"`
}

type AwsAutoscalingPolicySpec struct {
	AutoscalingGroupName        string                                                `json:"autoscaling_group_name"`
	MinAdjustmentStep           int                                                   `json:"min_adjustment_step"`
	ScalingAdjustment           int                                                   `json:"scaling_adjustment"`
	StepAdjustment              AwsAutoscalingPolicySpecStepAdjustment                `json:"step_adjustment"`
	TargetTrackingConfiguration []AwsAutoscalingPolicySpecTargetTrackingConfiguration `json:"target_tracking_configuration"`
	Name                        string                                                `json:"name"`
	AdjustmentType              string                                                `json:"adjustment_type"`
	PolicyType                  string                                                `json:"policy_type"`
	Cooldown                    int                                                   `json:"cooldown"`
	EstimatedInstanceWarmup     int                                                   `json:"estimated_instance_warmup"`
	MetricAggregationType       string                                                `json:"metric_aggregation_type"`
	MinAdjustmentMagnitude      int                                                   `json:"min_adjustment_magnitude"`
	Arn                         string                                                `json:"arn"`
}

type AwsAutoscalingPolicySpecStepAdjustment struct {
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfiguration struct {
	PredefinedMetricSpecification []AwsAutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification `json:"predefined_metric_specification"`
	CustomizedMetricSpecification []AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification `json:"customized_metric_specification"`
	TargetValue                   float64                                                                            `json:"target_value"`
	DisableScaleIn                bool                                                                               `json:"disable_scale_in"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	ResourceLabel        string `json:"resource_label"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification struct {
	MetricDimension []AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension `json:"metric_dimension"`
	MetricName      string                                                                                            `json:"metric_name"`
	Namespace       string                                                                                            `json:"namespace"`
	Statistic       string                                                                                            `json:"statistic"`
	Unit            string                                                                                            `json:"unit"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAutoscalingPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchMetricAlarm struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchMetricAlarmSpec `json"spec"`
}

type AwsCloudwatchMetricAlarmSpec struct {
	Statistic                         string            `json:"statistic"`
	TreatMissingData                  string            `json:"treat_missing_data"`
	EvaluationPeriods                 int               `json:"evaluation_periods"`
	Namespace                         string            `json:"namespace"`
	ActionsEnabled                    bool              `json:"actions_enabled"`
	AlarmDescription                  string            `json:"alarm_description"`
	DatapointsToAlarm                 int               `json:"datapoints_to_alarm"`
	OkActions                         string            `json:"ok_actions"`
	ExtendedStatistic                 string            `json:"extended_statistic"`
	MetricName                        string            `json:"metric_name"`
	Period                            int               `json:"period"`
	InsufficientDataActions           string            `json:"insufficient_data_actions"`
	Threshold                         float64           `json:"threshold"`
	Dimensions                        map[string]string `json:"dimensions"`
	AlarmActions                      string            `json:"alarm_actions"`
	Unit                              string            `json:"unit"`
	EvaluateLowSampleCountPercentiles string            `json:"evaluate_low_sample_count_percentiles"`
	AlarmName                         string            `json:"alarm_name"`
	ComparisonOperator                string            `json:"comparison_operator"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchMetricAlarmList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchMetricAlarm `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsTaskDefinition struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEcsTaskDefinitionSpec `json"spec"`
}

type AwsEcsTaskDefinitionSpec struct {
	Memory                  string                                       `json:"memory"`
	PlacementConstraints    AwsEcsTaskDefinitionSpecPlacementConstraints `json:"placement_constraints"`
	Arn                     string                                       `json:"arn"`
	Family                  string                                       `json:"family"`
	Revision                int                                          `json:"revision"`
	ContainerDefinitions    string                                       `json:"container_definitions"`
	TaskRoleArn             string                                       `json:"task_role_arn"`
	Cpu                     string                                       `json:"cpu"`
	ExecutionRoleArn        string                                       `json:"execution_role_arn"`
	NetworkMode             string                                       `json:"network_mode"`
	Volume                  AwsEcsTaskDefinitionSpecVolume               `json:"volume"`
	RequiresCompatibilities string                                       `json:"requires_compatibilities"`
}

type AwsEcsTaskDefinitionSpecPlacementConstraints struct {
	Type       string `json:"type"`
	Expression string `json:"expression"`
}

type AwsEcsTaskDefinitionSpecVolume struct {
	Name     string `json:"name"`
	HostPath string `json:"host_path"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsTaskDefinitionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEcsTaskDefinition `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserGroupMembership struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamUserGroupMembershipSpec `json"spec"`
}

type AwsIamUserGroupMembershipSpec struct {
	Groups string `json:"groups"`
	User   string `json:"user"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserGroupMembershipList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamUserGroupMembership `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalByteMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalByteMatchSetSpec `json"spec"`
}

type AwsWafregionalByteMatchSetSpec struct {
	Name           string                                       `json:"name"`
	ByteMatchTuple AwsWafregionalByteMatchSetSpecByteMatchTuple `json:"byte_match_tuple"`
}

type AwsWafregionalByteMatchSetSpecByteMatchTuple struct {
	FieldToMatch         AwsWafregionalByteMatchSetSpecByteMatchTupleFieldToMatch `json:"field_to_match"`
	PositionalConstraint string                                                   `json:"positional_constraint"`
	TargetString         string                                                   `json:"target_string"`
	TextTransformation   string                                                   `json:"text_transformation"`
}

type AwsWafregionalByteMatchSetSpecByteMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalByteMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalByteMatchSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayStage struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayStageSpec `json"spec"`
}

type AwsApiGatewayStageSpec struct {
	CacheClusterEnabled  bool                                      `json:"cache_cluster_enabled"`
	ClientCertificateId  string                                    `json:"client_certificate_id"`
	DocumentationVersion string                                    `json:"documentation_version"`
	ExecutionArn         string                                    `json:"execution_arn"`
	Variables            map[string]string                         `json:"variables"`
	StageName            string                                    `json:"stage_name"`
	Tags                 map[string]string                         `json:"tags"`
	AccessLogSettings    []AwsApiGatewayStageSpecAccessLogSettings `json:"access_log_settings"`
	CacheClusterSize     string                                    `json:"cache_cluster_size"`
	DeploymentId         string                                    `json:"deployment_id"`
	Description          string                                    `json:"description"`
	InvokeUrl            string                                    `json:"invoke_url"`
	RestApiId            string                                    `json:"rest_api_id"`
}

type AwsApiGatewayStageSpecAccessLogSettings struct {
	Format         string `json:"format"`
	DestinationArn string `json:"destination_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayStageList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayStage `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsService struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEcsServiceSpec `json"spec"`
}

type AwsEcsServiceSpec struct {
	ServiceRegistries               AwsEcsServiceSpecServiceRegistries          `json:"service_registries"`
	Name                            string                                      `json:"name"`
	DeploymentMinimumHealthyPercent int                                         `json:"deployment_minimum_healthy_percent"`
	LoadBalancer                    AwsEcsServiceSpecLoadBalancer               `json:"load_balancer"`
	IamRole                         string                                      `json:"iam_role"`
	DeploymentMaximumPercent        int                                         `json:"deployment_maximum_percent"`
	NetworkConfiguration            []AwsEcsServiceSpecNetworkConfiguration     `json:"network_configuration"`
	OrderedPlacementStrategy        []AwsEcsServiceSpecOrderedPlacementStrategy `json:"ordered_placement_strategy"`
	TaskDefinition                  string                                      `json:"task_definition"`
	LaunchType                      string                                      `json:"launch_type"`
	HealthCheckGracePeriodSeconds   int                                         `json:"health_check_grace_period_seconds"`
	PlacementStrategy               AwsEcsServiceSpecPlacementStrategy          `json:"placement_strategy"`
	PlacementConstraints            AwsEcsServiceSpecPlacementConstraints       `json:"placement_constraints"`
	Cluster                         string                                      `json:"cluster"`
	DesiredCount                    int                                         `json:"desired_count"`
}

type AwsEcsServiceSpecServiceRegistries struct {
	Port        int    `json:"port"`
	RegistryArn string `json:"registry_arn"`
}

type AwsEcsServiceSpecLoadBalancer struct {
	ElbName        string `json:"elb_name"`
	TargetGroupArn string `json:"target_group_arn"`
	ContainerName  string `json:"container_name"`
	ContainerPort  int    `json:"container_port"`
}

type AwsEcsServiceSpecNetworkConfiguration struct {
	SecurityGroups string `json:"security_groups"`
	Subnets        string `json:"subnets"`
	AssignPublicIp bool   `json:"assign_public_ip"`
}

type AwsEcsServiceSpecOrderedPlacementStrategy struct {
	Type  string `json:"type"`
	Field string `json:"field"`
}

type AwsEcsServiceSpecPlacementStrategy struct {
	Type  string `json:"type"`
	Field string `json:"field"`
}

type AwsEcsServiceSpecPlacementConstraints struct {
	Type       string `json:"type"`
	Expression string `json:"expression"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsServiceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEcsService `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53Zone struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRoute53ZoneSpec `json"spec"`
}

type AwsRoute53ZoneSpec struct {
	Name            string            `json:"name"`
	VpcId           string            `json:"vpc_id"`
	ZoneId          string            `json:"zone_id"`
	DelegationSetId string            `json:"delegation_set_id"`
	NameServers     []string          `json:"name_servers"`
	Comment         string            `json:"comment"`
	VpcRegion       string            `json:"vpc_region"`
	Tags            map[string]string `json:"tags"`
	ForceDestroy    bool              `json:"force_destroy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53ZoneList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRoute53Zone `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigConfigurationRecorder struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsConfigConfigurationRecorderSpec `json"spec"`
}

type AwsConfigConfigurationRecorderSpec struct {
	Name           string                                             `json:"name"`
	RoleArn        string                                             `json:"role_arn"`
	RecordingGroup []AwsConfigConfigurationRecorderSpecRecordingGroup `json:"recording_group"`
}

type AwsConfigConfigurationRecorderSpecRecordingGroup struct {
	ResourceTypes              string `json:"resource_types"`
	AllSupported               bool   `json:"all_supported"`
	IncludeGlobalResourceTypes bool   `json:"include_global_resource_types"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigConfigurationRecorderList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsConfigConfigurationRecorder `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbCookieStickinessPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLbCookieStickinessPolicySpec `json"spec"`
}

type AwsLbCookieStickinessPolicySpec struct {
	LbPort                 int    `json:"lb_port"`
	CookieExpirationPeriod int    `json:"cookie_expiration_period"`
	Name                   string `json:"name"`
	LoadBalancer           string `json:"load_balancer"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbCookieStickinessPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLbCookieStickinessPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerBackendServerPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLoadBalancerBackendServerPolicySpec `json"spec"`
}

type AwsLoadBalancerBackendServerPolicySpec struct {
	LoadBalancerName string `json:"load_balancer_name"`
	PolicyNames      string `json:"policy_names"`
	InstancePort     int    `json:"instance_port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerBackendServerPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLoadBalancerBackendServerPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53HealthCheck struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRoute53HealthCheckSpec `json"spec"`
}

type AwsRoute53HealthCheckSpec struct {
	InvertHealthcheck            bool              `json:"invert_healthcheck"`
	MeasureLatency               bool              `json:"measure_latency"`
	ReferenceName                string            `json:"reference_name"`
	IpAddress                    string            `json:"ip_address"`
	ResourcePath                 string            `json:"resource_path"`
	CloudwatchAlarmName          string            `json:"cloudwatch_alarm_name"`
	EnableSni                    bool              `json:"enable_sni"`
	Type                         string            `json:"type"`
	RequestInterval              int               `json:"request_interval"`
	ChildHealthchecks            string            `json:"child_healthchecks"`
	ChildHealthThreshold         int               `json:"child_health_threshold"`
	InsufficientDataHealthStatus string            `json:"insufficient_data_health_status"`
	Regions                      string            `json:"regions"`
	Tags                         map[string]string `json:"tags"`
	FailureThreshold             int               `json:"failure_threshold"`
	Fqdn                         string            `json:"fqdn"`
	Port                         int               `json:"port"`
	SearchString                 string            `json:"search_string"`
	CloudwatchAlarmRegion        string            `json:"cloudwatch_alarm_region"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53HealthCheckList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRoute53HealthCheck `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointServiceAllowedPrincipal struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcEndpointServiceAllowedPrincipalSpec `json"spec"`
}

type AwsVpcEndpointServiceAllowedPrincipalSpec struct {
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`
	PrincipalArn         string `json:"principal_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointServiceAllowedPrincipalList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcEndpointServiceAllowedPrincipal `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalWebAcl struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalWebAclSpec `json"spec"`
}

type AwsWafregionalWebAclSpec struct {
	Name          string                                  `json:"name"`
	DefaultAction []AwsWafregionalWebAclSpecDefaultAction `json:"default_action"`
	MetricName    string                                  `json:"metric_name"`
	Rule          AwsWafregionalWebAclSpecRule            `json:"rule"`
}

type AwsWafregionalWebAclSpecDefaultAction struct {
	Type string `json:"type"`
}

type AwsWafregionalWebAclSpecRule struct {
	Action   []AwsWafregionalWebAclSpecRuleAction `json:"action"`
	Priority int                                  `json:"priority"`
	RuleId   string                               `json:"rule_id"`
}

type AwsWafregionalWebAclSpecRuleAction struct {
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalWebAclList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalWebAcl `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCustomerGateway struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCustomerGatewaySpec `json"spec"`
}

type AwsCustomerGatewaySpec struct {
	BgpAsn    int               `json:"bgp_asn"`
	IpAddress string            `json:"ip_address"`
	Type      string            `json:"type"`
	Tags      map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCustomerGatewayList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCustomerGateway `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcrRepositoryPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEcrRepositoryPolicySpec `json"spec"`
}

type AwsEcrRepositoryPolicySpec struct {
	Repository string `json:"repository"`
	Policy     string `json:"policy"`
	RegistryId string `json:"registry_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcrRepositoryPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEcrRepositoryPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesActiveReceiptRuleSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesActiveReceiptRuleSetSpec `json"spec"`
}

type AwsSesActiveReceiptRuleSetSpec struct {
	RuleSetName string `json:"rule_set_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesActiveReceiptRuleSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesActiveReceiptRuleSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbListenerCertificate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAlbListenerCertificateSpec `json"spec"`
}

type AwsAlbListenerCertificateSpec struct {
	CertificateArn string `json:"certificate_arn"`
	ListenerArn    string `json:"listener_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbListenerCertificateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAlbListenerCertificate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerCertificate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLbListenerCertificateSpec `json"spec"`
}

type AwsLbListenerCertificateSpec struct {
	ListenerArn    string `json:"listener_arn"`
	CertificateArn string `json:"certificate_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerCertificateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLbListenerCertificate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogMetricFilter struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchLogMetricFilterSpec `json"spec"`
}

type AwsCloudwatchLogMetricFilterSpec struct {
	Name                 string                                                 `json:"name"`
	Pattern              string                                                 `json:"pattern"`
	LogGroupName         string                                                 `json:"log_group_name"`
	MetricTransformation []AwsCloudwatchLogMetricFilterSpecMetricTransformation `json:"metric_transformation"`
}

type AwsCloudwatchLogMetricFilterSpecMetricTransformation struct {
	Name         string  `json:"name"`
	Namespace    string  `json:"namespace"`
	Value        string  `json:"value"`
	DefaultValue float64 `json:"default_value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogMetricFilterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchLogMetricFilter `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEgressOnlyInternetGateway struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEgressOnlyInternetGatewaySpec `json"spec"`
}

type AwsEgressOnlyInternetGatewaySpec struct {
	VpcId string `json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEgressOnlyInternetGatewayList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEgressOnlyInternetGateway `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorAssessmentTemplate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsInspectorAssessmentTemplateSpec `json"spec"`
}

type AwsInspectorAssessmentTemplateSpec struct {
	Duration         int    `json:"duration"`
	RulesPackageArns string `json:"rules_package_arns"`
	Name             string `json:"name"`
	TargetArn        string `json:"target_arn"`
	Arn              string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorAssessmentTemplateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsInspectorAssessmentTemplate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafXssMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafXssMatchSetSpec `json"spec"`
}

type AwsWafXssMatchSetSpec struct {
	Name           string                              `json:"name"`
	XssMatchTuples AwsWafXssMatchSetSpecXssMatchTuples `json:"xss_match_tuples"`
}

type AwsWafXssMatchSetSpecXssMatchTuples struct {
	FieldToMatch       AwsWafXssMatchSetSpecXssMatchTuplesFieldToMatch `json:"field_to_match"`
	TextTransformation string                                          `json:"text_transformation"`
}

type AwsWafXssMatchSetSpecXssMatchTuplesFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafXssMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafXssMatchSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLbListenerRuleSpec `json"spec"`
}

type AwsLbListenerRuleSpec struct {
	Arn         string                         `json:"arn"`
	ListenerArn string                         `json:"listener_arn"`
	Priority    int                            `json:"priority"`
	Action      []AwsLbListenerRuleSpecAction  `json:"action"`
	Condition   AwsLbListenerRuleSpecCondition `json:"condition"`
}

type AwsLbListenerRuleSpecAction struct {
	TargetGroupArn string `json:"target_group_arn"`
	Type           string `json:"type"`
}

type AwsLbListenerRuleSpecCondition struct {
	Field  string   `json:"field"`
	Values []string `json:"values"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLbListenerRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsTopicPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSnsTopicPolicySpec `json"spec"`
}

type AwsSnsTopicPolicySpec struct {
	Arn    string `json:"arn"`
	Policy string `json:"policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsTopicPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSnsTopicPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRuleGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafRuleGroupSpec `json"spec"`
}

type AwsWafRuleGroupSpec struct {
	ActivatedRule AwsWafRuleGroupSpecActivatedRule `json:"activated_rule"`
	Name          string                           `json:"name"`
	MetricName    string                           `json:"metric_name"`
}

type AwsWafRuleGroupSpecActivatedRule struct {
	Type     string                                   `json:"type"`
	Action   []AwsWafRuleGroupSpecActivatedRuleAction `json:"action"`
	Priority int                                      `json:"priority"`
	RuleId   string                                   `json:"rule_id"`
}

type AwsWafRuleGroupSpecActivatedRuleAction struct {
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRuleGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafRuleGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEbsVolume struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEbsVolumeSpec `json"spec"`
}

type AwsEbsVolumeSpec struct {
	AvailabilityZone string            `json:"availability_zone"`
	KmsKeyId         string            `json:"kms_key_id"`
	Size             int               `json:"size"`
	SnapshotId       string            `json:"snapshot_id"`
	Type             string            `json:"type"`
	Arn              string            `json:"arn"`
	Encrypted        bool              `json:"encrypted"`
	Iops             int               `json:"iops"`
	Tags             map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEbsVolumeList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEbsVolume `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticsearchDomain struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticsearchDomainSpec `json"spec"`
}

type AwsElasticsearchDomainSpec struct {
	SnapshotOptions      []AwsElasticsearchDomainSpecSnapshotOptions    `json:"snapshot_options"`
	VpcOptions           []AwsElasticsearchDomainSpecVpcOptions         `json:"vpc_options"`
	Arn                  string                                         `json:"arn"`
	DomainName           string                                         `json:"domain_name"`
	Endpoint             string                                         `json:"endpoint"`
	EncryptAtRest        []AwsElasticsearchDomainSpecEncryptAtRest      `json:"encrypt_at_rest"`
	AccessPolicies       string                                         `json:"access_policies"`
	DomainId             string                                         `json:"domain_id"`
	ElasticsearchVersion string                                         `json:"elasticsearch_version"`
	AdvancedOptions      map[string]string                              `json:"advanced_options"`
	EbsOptions           []AwsElasticsearchDomainSpecEbsOptions         `json:"ebs_options"`
	ClusterConfig        []AwsElasticsearchDomainSpecClusterConfig      `json:"cluster_config"`
	LogPublishingOptions AwsElasticsearchDomainSpecLogPublishingOptions `json:"log_publishing_options"`
	Tags                 map[string]string                              `json:"tags"`
	KibanaEndpoint       string                                         `json:"kibana_endpoint"`
}

type AwsElasticsearchDomainSpecSnapshotOptions struct {
	AutomatedSnapshotStartHour int `json:"automated_snapshot_start_hour"`
}

type AwsElasticsearchDomainSpecVpcOptions struct {
	AvailabilityZones string `json:"availability_zones"`
	SecurityGroupIds  string `json:"security_group_ids"`
	SubnetIds         string `json:"subnet_ids"`
	VpcId             string `json:"vpc_id"`
}

type AwsElasticsearchDomainSpecEncryptAtRest struct {
	Enabled  bool   `json:"enabled"`
	KmsKeyId string `json:"kms_key_id"`
}

type AwsElasticsearchDomainSpecEbsOptions struct {
	EbsEnabled bool   `json:"ebs_enabled"`
	Iops       int    `json:"iops"`
	VolumeSize int    `json:"volume_size"`
	VolumeType string `json:"volume_type"`
}

type AwsElasticsearchDomainSpecClusterConfig struct {
	DedicatedMasterCount   int    `json:"dedicated_master_count"`
	DedicatedMasterEnabled bool   `json:"dedicated_master_enabled"`
	DedicatedMasterType    string `json:"dedicated_master_type"`
	InstanceCount          int    `json:"instance_count"`
	InstanceType           string `json:"instance_type"`
	ZoneAwarenessEnabled   bool   `json:"zone_awareness_enabled"`
}

type AwsElasticsearchDomainSpecLogPublishingOptions struct {
	LogType               string `json:"log_type"`
	CloudwatchLogGroupArn string `json:"cloudwatch_log_group_arn"`
	Enabled               bool   `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticsearchDomainList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticsearchDomain `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrInstanceGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEmrInstanceGroupSpec `json"spec"`
}

type AwsEmrInstanceGroupSpec struct {
	EbsOptimized         bool                             `json:"ebs_optimized"`
	EbsConfig            AwsEmrInstanceGroupSpecEbsConfig `json:"ebs_config"`
	ClusterId            string                           `json:"cluster_id"`
	InstanceType         string                           `json:"instance_type"`
	InstanceCount        int                              `json:"instance_count"`
	RunningInstanceCount int                              `json:"running_instance_count"`
	Status               string                           `json:"status"`
	Name                 string                           `json:"name"`
}

type AwsEmrInstanceGroupSpecEbsConfig struct {
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int    `json:"volumes_per_instance"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrInstanceGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEmrInstanceGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterface struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsNetworkInterfaceSpec `json"spec"`
}

type AwsNetworkInterfaceSpec struct {
	PrivateDnsName  string                            `json:"private_dns_name"`
	SecurityGroups  string                            `json:"security_groups"`
	SourceDestCheck bool                              `json:"source_dest_check"`
	Description     string                            `json:"description"`
	Attachment      AwsNetworkInterfaceSpecAttachment `json:"attachment"`
	Tags            map[string]string                 `json:"tags"`
	SubnetId        string                            `json:"subnet_id"`
	PrivateIps      string                            `json:"private_ips"`
	PrivateIpsCount int                               `json:"private_ips_count"`
	PrivateIp       string                            `json:"private_ip"`
}

type AwsNetworkInterfaceSpecAttachment struct {
	Instance     string `json:"instance"`
	DeviceIndex  int    `json:"device_index"`
	AttachmentId string `json:"attachment_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterfaceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsNetworkInterface `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsCluster struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRdsClusterSpec `json"spec"`
}

type AwsRdsClusterSpec struct {
	Endpoint                         string                      `json:"endpoint"`
	MasterPassword                   string                      `json:"master_password"`
	PreferredMaintenanceWindow       string                      `json:"preferred_maintenance_window"`
	KmsKeyId                         string                      `json:"kms_key_id"`
	ClusterIdentifierPrefix          string                      `json:"cluster_identifier_prefix"`
	HostedZoneId                     string                      `json:"hosted_zone_id"`
	FinalSnapshotIdentifier          string                      `json:"final_snapshot_identifier"`
	SnapshotIdentifier               string                      `json:"snapshot_identifier"`
	IamRoles                         string                      `json:"iam_roles"`
	DbClusterParameterGroupName      string                      `json:"db_cluster_parameter_group_name"`
	S3Import                         []AwsRdsClusterSpecS3Import `json:"s3_import"`
	Port                             int                         `json:"port"`
	BackupRetentionPeriod            int                         `json:"backup_retention_period"`
	ClusterResourceId                string                      `json:"cluster_resource_id"`
	DatabaseName                     string                      `json:"database_name"`
	ReaderEndpoint                   string                      `json:"reader_endpoint"`
	EngineVersion                    string                      `json:"engine_version"`
	MasterUsername                   string                      `json:"master_username"`
	PreferredBackupWindow            string                      `json:"preferred_backup_window"`
	ClusterIdentifier                string                      `json:"cluster_identifier"`
	ClusterMembers                   string                      `json:"cluster_members"`
	DbSubnetGroupName                string                      `json:"db_subnet_group_name"`
	AvailabilityZones                string                      `json:"availability_zones"`
	ApplyImmediately                 bool                        `json:"apply_immediately"`
	VpcSecurityGroupIds              string                      `json:"vpc_security_group_ids"`
	IamDatabaseAuthenticationEnabled bool                        `json:"iam_database_authentication_enabled"`
	SourceRegion                     string                      `json:"source_region"`
	Tags                             map[string]string           `json:"tags"`
	StorageEncrypted                 bool                        `json:"storage_encrypted"`
	SkipFinalSnapshot                bool                        `json:"skip_final_snapshot"`
	ReplicationSourceIdentifier      string                      `json:"replication_source_identifier"`
	Engine                           string                      `json:"engine"`
}

type AwsRdsClusterSpecS3Import struct {
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsClusterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRdsCluster `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindow struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmMaintenanceWindowSpec `json"spec"`
}

type AwsSsmMaintenanceWindowSpec struct {
	Name                     string `json:"name"`
	Schedule                 string `json:"schedule"`
	Duration                 int    `json:"duration"`
	Cutoff                   int    `json:"cutoff"`
	AllowUnassociatedTargets bool   `json:"allow_unassociated_targets"`
	Enabled                  bool   `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmMaintenanceWindow `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksHaproxyLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksHaproxyLayerSpec `json"spec"`
}

type AwsOpsworksHaproxyLayerSpec struct {
	StackId                  string                               `json:"stack_id"`
	StatsPassword            string                               `json:"stats_password"`
	StatsUser                string                               `json:"stats_user"`
	CustomDeployRecipes      []string                             `json:"custom_deploy_recipes"`
	CustomSecurityGroupIds   string                               `json:"custom_security_group_ids"`
	AutoHealing              bool                                 `json:"auto_healing"`
	InstallUpdatesOnBoot     bool                                 `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  int                                  `json:"instance_shutdown_timeout"`
	SystemPackages           string                               `json:"system_packages"`
	HealthcheckMethod        string                               `json:"healthcheck_method"`
	AutoAssignElasticIps     bool                                 `json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn string                               `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string                             `json:"custom_setup_recipes"`
	CustomConfigureRecipes   []string                             `json:"custom_configure_recipes"`
	CustomJson               string                               `json:"custom_json"`
	ElasticLoadBalancer      string                               `json:"elastic_load_balancer"`
	CustomShutdownRecipes    []string                             `json:"custom_shutdown_recipes"`
	DrainElbOnShutdown       bool                                 `json:"drain_elb_on_shutdown"`
	HealthcheckUrl           string                               `json:"healthcheck_url"`
	StatsEnabled             bool                                 `json:"stats_enabled"`
	StatsUrl                 string                               `json:"stats_url"`
	AutoAssignPublicIps      bool                                 `json:"auto_assign_public_ips"`
	CustomUndeployRecipes    []string                             `json:"custom_undeploy_recipes"`
	UseEbsOptimizedInstances bool                                 `json:"use_ebs_optimized_instances"`
	EbsVolume                AwsOpsworksHaproxyLayerSpecEbsVolume `json:"ebs_volume"`
	Name                     string                               `json:"name"`
}

type AwsOpsworksHaproxyLayerSpecEbsVolume struct {
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksHaproxyLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksHaproxyLayer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmAssociationSpec `json"spec"`
}

type AwsSsmAssociationSpec struct {
	AssociationName    string                                `json:"association_name"`
	InstanceId         string                                `json:"instance_id"`
	DocumentVersion    string                                `json:"document_version"`
	Name               string                                `json:"name"`
	AssociationId      string                                `json:"association_id"`
	Parameters         map[string]string                     `json:"parameters"`
	ScheduleExpression string                                `json:"schedule_expression"`
	OutputLocation     []AwsSsmAssociationSpecOutputLocation `json:"output_location"`
	Targets            []AwsSsmAssociationSpecTargets        `json:"targets"`
}

type AwsSsmAssociationSpecOutputLocation struct {
	S3BucketName string `json:"s3_bucket_name"`
	S3KeyPrefix  string `json:"s3_key_prefix"`
}

type AwsSsmAssociationSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmAssociation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafByteMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafByteMatchSetSpec `json"spec"`
}

type AwsWafByteMatchSetSpec struct {
	Name            string                                `json:"name"`
	ByteMatchTuples AwsWafByteMatchSetSpecByteMatchTuples `json:"byte_match_tuples"`
}

type AwsWafByteMatchSetSpecByteMatchTuples struct {
	FieldToMatch         AwsWafByteMatchSetSpecByteMatchTuplesFieldToMatch `json:"field_to_match"`
	PositionalConstraint string                                            `json:"positional_constraint"`
	TargetString         string                                            `json:"target_string"`
	TextTransformation   string                                            `json:"text_transformation"`
}

type AwsWafByteMatchSetSpecByteMatchTuplesFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafByteMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafByteMatchSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAppautoscalingPolicySpec `json"spec"`
}

type AwsAppautoscalingPolicySpec struct {
	Arn                                      string                                                                `json:"arn"`
	ResourceId                               string                                                                `json:"resource_id"`
	AdjustmentType                           string                                                                `json:"adjustment_type"`
	Cooldown                                 int                                                                   `json:"cooldown"`
	MinAdjustmentMagnitude                   int                                                                   `json:"min_adjustment_magnitude"`
	Name                                     string                                                                `json:"name"`
	Alarms                                   []string                                                              `json:"alarms"`
	PolicyType                               string                                                                `json:"policy_type"`
	ScalableDimension                        string                                                                `json:"scalable_dimension"`
	StepAdjustment                           AwsAppautoscalingPolicySpecStepAdjustment                             `json:"step_adjustment"`
	TargetTrackingScalingPolicyConfiguration []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"target_tracking_scaling_policy_configuration"`
	ServiceNamespace                         string                                                                `json:"service_namespace"`
	StepScalingPolicyConfiguration           []AwsAppautoscalingPolicySpecStepScalingPolicyConfiguration           `json:"step_scaling_policy_configuration"`
	MetricAggregationType                    string                                                                `json:"metric_aggregation_type"`
}

type AwsAppautoscalingPolicySpecStepAdjustment struct {
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration struct {
	ScaleInCooldown               int                                                                                                `json:"scale_in_cooldown"`
	ScaleOutCooldown              int                                                                                                `json:"scale_out_cooldown"`
	TargetValue                   float64                                                                                            `json:"target_value"`
	CustomizedMetricSpecification []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification `json:"customized_metric_specification"`
	PredefinedMetricSpecification []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification `json:"predefined_metric_specification"`
	DisableScaleIn                bool                                                                                               `json:"disable_scale_in"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification struct {
	Namespace  string                                                                                                     `json:"namespace"`
	Statistic  string                                                                                                     `json:"statistic"`
	Unit       string                                                                                                     `json:"unit"`
	Dimensions AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions `json:"dimensions"`
	MetricName string                                                                                                     `json:"metric_name"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	ResourceLabel        string `json:"resource_label"`
}

type AwsAppautoscalingPolicySpecStepScalingPolicyConfiguration struct {
	AdjustmentType         string                                                                  `json:"adjustment_type"`
	Cooldown               int                                                                     `json:"cooldown"`
	MetricAggregationType  string                                                                  `json:"metric_aggregation_type"`
	MinAdjustmentMagnitude int                                                                     `json:"min_adjustment_magnitude"`
	StepAdjustment         AwsAppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment `json:"step_adjustment"`
}

type AwsAppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment struct {
	MetricIntervalLowerBound float64 `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound float64 `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int     `json:"scaling_adjustment"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAppautoscalingPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAthenaNamedQuery struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAthenaNamedQuerySpec `json"spec"`
}

type AwsAthenaNamedQuerySpec struct {
	Database    string `json:"database"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Query       string `json:"query"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAthenaNamedQueryList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAthenaNamedQuery `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53Record struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRoute53RecordSpec `json"spec"`
}

type AwsRoute53RecordSpec struct {
	ZoneId                        string                                         `json:"zone_id"`
	SetIdentifier                 string                                         `json:"set_identifier"`
	HealthCheckId                 string                                         `json:"health_check_id"`
	AllowOverwrite                bool                                           `json:"allow_overwrite"`
	Alias                         AwsRoute53RecordSpecAlias                      `json:"alias"`
	GeolocationRoutingPolicy      []AwsRoute53RecordSpecGeolocationRoutingPolicy `json:"geolocation_routing_policy"`
	MultivalueAnswerRoutingPolicy bool                                           `json:"multivalue_answer_routing_policy"`
	Records                       string                                         `json:"records"`
	Name                          string                                         `json:"name"`
	Ttl                           int                                            `json:"ttl"`
	Weight                        int                                            `json:"weight"`
	FailoverRoutingPolicy         []AwsRoute53RecordSpecFailoverRoutingPolicy    `json:"failover_routing_policy"`
	LatencyRoutingPolicy          []AwsRoute53RecordSpecLatencyRoutingPolicy     `json:"latency_routing_policy"`
	WeightedRoutingPolicy         []AwsRoute53RecordSpecWeightedRoutingPolicy    `json:"weighted_routing_policy"`
	Fqdn                          string                                         `json:"fqdn"`
	Type                          string                                         `json:"type"`
	Failover                      string                                         `json:"failover"`
}

type AwsRoute53RecordSpecAlias struct {
	ZoneId               string `json:"zone_id"`
	Name                 string `json:"name"`
	EvaluateTargetHealth bool   `json:"evaluate_target_health"`
}

type AwsRoute53RecordSpecGeolocationRoutingPolicy struct {
	Continent   string `json:"continent"`
	Country     string `json:"country"`
	Subdivision string `json:"subdivision"`
}

type AwsRoute53RecordSpecFailoverRoutingPolicy struct {
	Type string `json:"type"`
}

type AwsRoute53RecordSpecLatencyRoutingPolicy struct {
	Region string `json:"region"`
}

type AwsRoute53RecordSpecWeightedRoutingPolicy struct {
	Weight int `json:"weight"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53RecordList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRoute53Record `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesConfigurationSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesConfigurationSetSpec `json"spec"`
}

type AwsSesConfigurationSetSpec struct {
	Name string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesConfigurationSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesConfigurationSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServicecatalogPortfolio struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsServicecatalogPortfolioSpec `json"spec"`
}

type AwsServicecatalogPortfolioSpec struct {
	Description  string            `json:"description"`
	ProviderName string            `json:"provider_name"`
	Tags         map[string]string `json:"tags"`
	Arn          string            `json:"arn"`
	CreatedTime  string            `json:"created_time"`
	Name         string            `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServicecatalogPortfolioList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsServicecatalogPortfolio `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppsyncDatasource struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAppsyncDatasourceSpec `json"spec"`
}

type AwsAppsyncDatasourceSpec struct {
	Type                string                                        `json:"type"`
	DynamodbConfig      []AwsAppsyncDatasourceSpecDynamodbConfig      `json:"dynamodb_config"`
	LambdaConfig        []AwsAppsyncDatasourceSpecLambdaConfig        `json:"lambda_config"`
	ApiId               string                                        `json:"api_id"`
	Name                string                                        `json:"name"`
	Description         string                                        `json:"description"`
	ElasticsearchConfig []AwsAppsyncDatasourceSpecElasticsearchConfig `json:"elasticsearch_config"`
	ServiceRoleArn      string                                        `json:"service_role_arn"`
	Arn                 string                                        `json:"arn"`
}

type AwsAppsyncDatasourceSpecDynamodbConfig struct {
	Region               string `json:"region"`
	TableName            string `json:"table_name"`
	UseCallerCredentials bool   `json:"use_caller_credentials"`
}

type AwsAppsyncDatasourceSpecLambdaConfig struct {
	FunctionArn string `json:"function_arn"`
}

type AwsAppsyncDatasourceSpecElasticsearchConfig struct {
	Region   string `json:"region"`
	Endpoint string `json:"endpoint"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppsyncDatasourceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAppsyncDatasource `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBudgetsBudget struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsBudgetsBudgetSpec `json"spec"`
}

type AwsBudgetsBudgetSpec struct {
	TimeUnit        string                          `json:"time_unit"`
	CostFilters     map[string]string               `json:"cost_filters"`
	AccountId       string                          `json:"account_id"`
	LimitAmount     string                          `json:"limit_amount"`
	LimitUnit       string                          `json:"limit_unit"`
	CostTypes       []AwsBudgetsBudgetSpecCostTypes `json:"cost_types"`
	TimePeriodStart string                          `json:"time_period_start"`
	TimePeriodEnd   string                          `json:"time_period_end"`
	Name            string                          `json:"name"`
	NamePrefix      string                          `json:"name_prefix"`
	BudgetType      string                          `json:"budget_type"`
}

type AwsBudgetsBudgetSpecCostTypes struct {
	UseBlended               bool `json:"use_blended"`
	IncludeDiscount          bool `json:"include_discount"`
	IncludeSupport           bool `json:"include_support"`
	IncludeUpfront           bool `json:"include_upfront"`
	IncludeRefund            bool `json:"include_refund"`
	IncludeSubscription      bool `json:"include_subscription"`
	IncludeTax               bool `json:"include_tax"`
	UseAmortized             bool `json:"use_amortized"`
	IncludeCredit            bool `json:"include_credit"`
	IncludeOtherSubscription bool `json:"include_other_subscription"`
	IncludeRecurring         bool `json:"include_recurring"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBudgetsBudgetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsBudgetsBudget `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogResourcePolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchLogResourcePolicySpec `json"spec"`
}

type AwsCloudwatchLogResourcePolicySpec struct {
	PolicyName     string `json:"policy_name"`
	PolicyDocument string `json:"policy_document"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogResourcePolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchLogResourcePolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkAclRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsNetworkAclRuleSpec `json"spec"`
}

type AwsNetworkAclRuleSpec struct {
	CidrBlock     string `json:"cidr_block"`
	IcmpCode      string `json:"icmp_code"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	FromPort      int    `json:"from_port"`
	ToPort        int    `json:"to_port"`
	NetworkAclId  string `json:"network_acl_id"`
	RuleNumber    int    `json:"rule_number"`
	Egress        bool   `json:"egress"`
	Protocol      string `json:"protocol"`
	RuleAction    string `json:"rule_action"`
	IcmpType      string `json:"icmp_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkAclRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsNetworkAclRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnConnectionRoute struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpnConnectionRouteSpec `json"spec"`
}

type AwsVpnConnectionRouteSpec struct {
	DestinationCidrBlock string `json:"destination_cidr_block"`
	VpnConnectionId      string `json:"vpn_connection_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnConnectionRouteList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpnConnectionRoute `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafSizeConstraintSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafSizeConstraintSetSpec `json"spec"`
}

type AwsWafSizeConstraintSetSpec struct {
	SizeConstraints AwsWafSizeConstraintSetSpecSizeConstraints `json:"size_constraints"`
	Name            string                                     `json:"name"`
}

type AwsWafSizeConstraintSetSpecSizeConstraints struct {
	FieldToMatch       AwsWafSizeConstraintSetSpecSizeConstraintsFieldToMatch `json:"field_to_match"`
	ComparisonOperator string                                                 `json:"comparison_operator"`
	Size               int                                                    `json:"size"`
	TextTransformation string                                                 `json:"text_transformation"`
}

type AwsWafSizeConstraintSetSpecSizeConstraintsFieldToMatch struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafSizeConstraintSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafSizeConstraintSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayApiKey struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayApiKeySpec `json"spec"`
}

type AwsApiGatewayApiKeySpec struct {
	Name            string                          `json:"name"`
	Description     string                          `json:"description"`
	Enabled         bool                            `json:"enabled"`
	StageKey        AwsApiGatewayApiKeySpecStageKey `json:"stage_key"`
	CreatedDate     string                          `json:"created_date"`
	LastUpdatedDate string                          `json:"last_updated_date"`
	Value           string                          `json:"value"`
}

type AwsApiGatewayApiKeySpecStageKey struct {
	RestApiId string `json:"rest_api_id"`
	StageName string `json:"stage_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayApiKeyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayApiKey `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDocumentationPart struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayDocumentationPartSpec `json"spec"`
}

type AwsApiGatewayDocumentationPartSpec struct {
	Location   []AwsApiGatewayDocumentationPartSpecLocation `json:"location"`
	Properties string                                       `json:"properties"`
	RestApiId  string                                       `json:"rest_api_id"`
}

type AwsApiGatewayDocumentationPartSpecLocation struct {
	Method     string `json:"method"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	StatusCode string `json:"status_code"`
	Type       string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDocumentationPartList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayDocumentationPart `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethodSettings struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayMethodSettingsSpec `json"spec"`
}

type AwsApiGatewayMethodSettingsSpec struct {
	StageName  string                                    `json:"stage_name"`
	MethodPath string                                    `json:"method_path"`
	Settings   []AwsApiGatewayMethodSettingsSpecSettings `json:"settings"`
	RestApiId  string                                    `json:"rest_api_id"`
}

type AwsApiGatewayMethodSettingsSpecSettings struct {
	MetricsEnabled                         bool    `json:"metrics_enabled"`
	ThrottlingBurstLimit                   int     `json:"throttling_burst_limit"`
	CacheDataEncrypted                     bool    `json:"cache_data_encrypted"`
	RequireAuthorizationForCacheControl    bool    `json:"require_authorization_for_cache_control"`
	LoggingLevel                           string  `json:"logging_level"`
	DataTraceEnabled                       bool    `json:"data_trace_enabled"`
	ThrottlingRateLimit                    float64 `json:"throttling_rate_limit"`
	CachingEnabled                         bool    `json:"caching_enabled"`
	CacheTtlInSeconds                      int     `json:"cache_ttl_in_seconds"`
	UnauthorizedCacheControlHeaderStrategy string  `json:"unauthorized_cache_control_header_strategy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethodSettingsList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayMethodSettings `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayUsagePlanKey struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayUsagePlanKeySpec `json"spec"`
}

type AwsApiGatewayUsagePlanKeySpec struct {
	KeyId       string `json:"key_id"`
	KeyType     string `json:"key_type"`
	UsagePlanId string `json:"usage_plan_id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayUsagePlanKeyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayUsagePlanKey `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryService struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsServiceDiscoveryServiceSpec `json"spec"`
}

type AwsServiceDiscoveryServiceSpec struct {
	Name                    string                                                  `json:"name"`
	Description             string                                                  `json:"description"`
	DnsConfig               []AwsServiceDiscoveryServiceSpecDnsConfig               `json:"dns_config"`
	HealthCheckConfig       []AwsServiceDiscoveryServiceSpecHealthCheckConfig       `json:"health_check_config"`
	HealthCheckCustomConfig []AwsServiceDiscoveryServiceSpecHealthCheckCustomConfig `json:"health_check_custom_config"`
	Arn                     string                                                  `json:"arn"`
}

type AwsServiceDiscoveryServiceSpecDnsConfig struct {
	NamespaceId   string                                              `json:"namespace_id"`
	DnsRecords    []AwsServiceDiscoveryServiceSpecDnsConfigDnsRecords `json:"dns_records"`
	RoutingPolicy string                                              `json:"routing_policy"`
}

type AwsServiceDiscoveryServiceSpecDnsConfigDnsRecords struct {
	Ttl  int    `json:"ttl"`
	Type string `json:"type"`
}

type AwsServiceDiscoveryServiceSpecHealthCheckConfig struct {
	Type             string `json:"type"`
	FailureThreshold int    `json:"failure_threshold"`
	ResourcePath     string `json:"resource_path"`
}

type AwsServiceDiscoveryServiceSpecHealthCheckCustomConfig struct {
	FailureThreshold int `json:"failure_threshold"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryServiceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsServiceDiscoveryService `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSqsQueue struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSqsQueueSpec `json"spec"`
}

type AwsSqsQueueSpec struct {
	Arn                          string            `json:"arn"`
	NamePrefix                   string            `json:"name_prefix"`
	ReceiveWaitTimeSeconds       int               `json:"receive_wait_time_seconds"`
	Policy                       string            `json:"policy"`
	KmsDataKeyReusePeriodSeconds int               `json:"kms_data_key_reuse_period_seconds"`
	MaxMessageSize               int               `json:"max_message_size"`
	DelaySeconds                 int               `json:"delay_seconds"`
	FifoQueue                    bool              `json:"fifo_queue"`
	ContentBasedDeduplication    bool              `json:"content_based_deduplication"`
	KmsMasterKeyId               string            `json:"kms_master_key_id"`
	Tags                         map[string]string `json:"tags"`
	Name                         string            `json:"name"`
	VisibilityTimeoutSeconds     int               `json:"visibility_timeout_seconds"`
	RedrivePolicy                string            `json:"redrive_policy"`
	MessageRetentionSeconds      int               `json:"message_retention_seconds"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSqsQueueList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSqsQueue `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSfnStateMachine struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSfnStateMachineSpec `json"spec"`
}

type AwsSfnStateMachineSpec struct {
	Status       string `json:"status"`
	Definition   string `json:"definition"`
	Name         string `json:"name"`
	RoleArn      string `json:"role_arn"`
	CreationDate string `json:"creation_date"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSfnStateMachineList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSfnStateMachine `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalSizeConstraintSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalSizeConstraintSetSpec `json"spec"`
}

type AwsWafregionalSizeConstraintSetSpec struct {
	Name            string                                             `json:"name"`
	SizeConstraints AwsWafregionalSizeConstraintSetSpecSizeConstraints `json:"size_constraints"`
}

type AwsWafregionalSizeConstraintSetSpecSizeConstraints struct {
	FieldToMatch       AwsWafregionalSizeConstraintSetSpecSizeConstraintsFieldToMatch `json:"field_to_match"`
	ComparisonOperator string                                                         `json:"comparison_operator"`
	Size               int                                                            `json:"size"`
	TextTransformation string                                                         `json:"text_transformation"`
}

type AwsWafregionalSizeConstraintSetSpecSizeConstraintsFieldToMatch struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalSizeConstraintSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalSizeConstraintSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDomainName struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayDomainNameSpec `json"spec"`
}

type AwsApiGatewayDomainNameSpec struct {
	CertificateBody       string `json:"certificate_body"`
	CertificateName       string `json:"certificate_name"`
	DomainName            string `json:"domain_name"`
	CloudfrontDomainName  string `json:"cloudfront_domain_name"`
	CertificateChain      string `json:"certificate_chain"`
	CertificatePrivateKey string `json:"certificate_private_key"`
	CertificateArn        string `json:"certificate_arn"`
	CertificateUploadDate string `json:"certificate_upload_date"`
	CloudfrontZoneId      string `json:"cloudfront_zone_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDomainNameList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayDomainName `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationTask struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDmsReplicationTaskSpec `json"spec"`
}

type AwsDmsReplicationTaskSpec struct {
	TargetEndpointArn       string            `json:"target_endpoint_arn"`
	CdcStartTime            string            `json:"cdc_start_time"`
	MigrationType           string            `json:"migration_type"`
	ReplicationTaskArn      string            `json:"replication_task_arn"`
	SourceEndpointArn       string            `json:"source_endpoint_arn"`
	TableMappings           string            `json:"table_mappings"`
	ReplicationInstanceArn  string            `json:"replication_instance_arn"`
	ReplicationTaskId       string            `json:"replication_task_id"`
	ReplicationTaskSettings string            `json:"replication_task_settings"`
	Tags                    map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationTaskList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDmsReplicationTask `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxConnection struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDxConnectionSpec `json"spec"`
}

type AwsDxConnectionSpec struct {
	Bandwidth string            `json:"bandwidth"`
	Location  string            `json:"location"`
	Tags      map[string]string `json:"tags"`
	Arn       string            `json:"arn"`
	Name      string            `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxConnectionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDxConnection `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsProxyProtocolPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsProxyProtocolPolicySpec `json"spec"`
}

type AwsProxyProtocolPolicySpec struct {
	InstancePorts string `json:"instance_ports"`
	LoadBalancer  string `json:"load_balancer"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsProxyProtocolPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsProxyProtocolPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRouteSpec `json"spec"`
}

type AwsRouteSpec struct {
	InstanceId               string `json:"instance_id"`
	NetworkInterfaceId       string `json:"network_interface_id"`
	DestinationCidrBlock     string `json:"destination_cidr_block"`
	DestinationIpv6CidrBlock string `json:"destination_ipv6_cidr_block"`
	DestinationPrefixListId  string `json:"destination_prefix_list_id"`
	GatewayId                string `json:"gateway_id"`
	EgressOnlyGatewayId      string `json:"egress_only_gateway_id"`
	NatGatewayId             string `json:"nat_gateway_id"`
	Origin                   string `json:"origin"`
	State                    string `json:"state"`
	InstanceOwnerId          string `json:"instance_owner_id"`
	RouteTableId             string `json:"route_table_id"`
	VpcPeeringConnectionId   string `json:"vpc_peering_connection_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRouteList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRoute `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptFilter struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesReceiptFilterSpec `json"spec"`
}

type AwsSesReceiptFilterSpec struct {
	Name   string `json:"name"`
	Cidr   string `json:"cidr"`
	Policy string `json:"policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptFilterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesReceiptFilter `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRateBasedRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalRateBasedRuleSpec `json"spec"`
}

type AwsWafregionalRateBasedRuleSpec struct {
	Name       string                                   `json:"name"`
	MetricName string                                   `json:"metric_name"`
	Predicate  AwsWafregionalRateBasedRuleSpecPredicate `json:"predicate"`
	RateKey    string                                   `json:"rate_key"`
	RateLimit  int                                      `json:"rate_limit"`
}

type AwsWafregionalRateBasedRuleSpecPredicate struct {
	Negated bool   `json:"negated"`
	DataId  string `json:"data_id"`
	Type    string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRateBasedRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalRateBasedRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalXssMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalXssMatchSetSpec `json"spec"`
}

type AwsWafregionalXssMatchSetSpec struct {
	Name          string                                     `json:"name"`
	XssMatchTuple AwsWafregionalXssMatchSetSpecXssMatchTuple `json:"xss_match_tuple"`
}

type AwsWafregionalXssMatchSetSpecXssMatchTuple struct {
	FieldToMatch       AwsWafregionalXssMatchSetSpecXssMatchTupleFieldToMatch `json:"field_to_match"`
	TextTransformation string                                                 `json:"text_transformation"`
}

type AwsWafregionalXssMatchSetSpecXssMatchTupleFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalXssMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalXssMatchSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkEnvironment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticBeanstalkEnvironmentSpec `json"spec"`
}

type AwsElasticBeanstalkEnvironmentSpec struct {
	WaitForReadyTimeout  string                                        `json:"wait_for_ready_timeout"`
	CnamePrefix          string                                        `json:"cname_prefix"`
	SolutionStackName    string                                        `json:"solution_stack_name"`
	AllSettings          AwsElasticBeanstalkEnvironmentSpecAllSettings `json:"all_settings"`
	Instances            []string                                      `json:"instances"`
	VersionLabel         string                                        `json:"version_label"`
	Tier                 string                                        `json:"tier"`
	Application          string                                        `json:"application"`
	Cname                string                                        `json:"cname"`
	Setting              AwsElasticBeanstalkEnvironmentSpecSetting     `json:"setting"`
	TemplateName         string                                        `json:"template_name"`
	AutoscalingGroups    []string                                      `json:"autoscaling_groups"`
	LaunchConfigurations []string                                      `json:"launch_configurations"`
	Arn                  string                                        `json:"arn"`
	Name                 string                                        `json:"name"`
	Triggers             []string                                      `json:"triggers"`
	Tags                 map[string]string                             `json:"tags"`
	LoadBalancers        []string                                      `json:"load_balancers"`
	Queues               []string                                      `json:"queues"`
	Description          string                                        `json:"description"`
	PollInterval         string                                        `json:"poll_interval"`
}

type AwsElasticBeanstalkEnvironmentSpecAllSettings struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Resource  string `json:"resource"`
}

type AwsElasticBeanstalkEnvironmentSpecSetting struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Resource  string `json:"resource"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkEnvironmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticBeanstalkEnvironment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueJob struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGlueJobSpec `json"spec"`
}

type AwsGlueJobSpec struct {
	ExecutionProperty []AwsGlueJobSpecExecutionProperty `json:"execution_property"`
	MaxRetries        int                               `json:"max_retries"`
	AllocatedCapacity int                               `json:"allocated_capacity"`
	Command           []AwsGlueJobSpecCommand           `json:"command"`
	Connections       []string                          `json:"connections"`
	DefaultArguments  map[string]string                 `json:"default_arguments"`
	Description       string                            `json:"description"`
	Name              string                            `json:"name"`
	RoleArn           string                            `json:"role_arn"`
}

type AwsGlueJobSpecExecutionProperty struct {
	MaxConcurrentRuns int `json:"max_concurrent_runs"`
}

type AwsGlueJobSpecCommand struct {
	Name           string `json:"name"`
	ScriptLocation string `json:"script_location"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueJobList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGlueJob `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupPolicyAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamGroupPolicyAttachmentSpec `json"spec"`
}

type AwsIamGroupPolicyAttachmentSpec struct {
	Group     string `json:"group"`
	PolicyArn string `json:"policy_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupPolicyAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamGroupPolicyAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamPolicyAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamPolicyAttachmentSpec `json"spec"`
}

type AwsIamPolicyAttachmentSpec struct {
	Users     string `json:"users"`
	Roles     string `json:"roles"`
	Groups    string `json:"groups"`
	PolicyArn string `json:"policy_arn"`
	Name      string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamPolicyAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamPolicyAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOrganizationsPolicySpec `json"spec"`
}

type AwsOrganizationsPolicySpec struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Arn         string `json:"arn"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOrganizationsPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVolumeAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVolumeAttachmentSpec `json"spec"`
}

type AwsVolumeAttachmentSpec struct {
	SkipDestroy bool   `json:"skip_destroy"`
	DeviceName  string `json:"device_name"`
	InstanceId  string `json:"instance_id"`
	VolumeId    string `json:"volume_id"`
	ForceDetach bool   `json:"force_detach"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVolumeAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVolumeAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailStaticIp struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLightsailStaticIpSpec `json"spec"`
}

type AwsLightsailStaticIpSpec struct {
	Name        string `json:"name"`
	IpAddress   string `json:"ip_address"`
	Arn         string `json:"arn"`
	SupportCode string `json:"support_code"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailStaticIpList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLightsailStaticIp `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53DelegationSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRoute53DelegationSetSpec `json"spec"`
}

type AwsRoute53DelegationSetSpec struct {
	ReferenceName string   `json:"reference_name"`
	NameServers   []string `json:"name_servers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53DelegationSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRoute53DelegationSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowTarget struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmMaintenanceWindowTargetSpec `json"spec"`
}

type AwsSsmMaintenanceWindowTargetSpec struct {
	WindowId         string                                     `json:"window_id"`
	ResourceType     string                                     `json:"resource_type"`
	Targets          []AwsSsmMaintenanceWindowTargetSpecTargets `json:"targets"`
	OwnerInformation string                                     `json:"owner_information"`
}

type AwsSsmMaintenanceWindowTargetSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowTargetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmMaintenanceWindowTarget `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayRoutePropagation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpnGatewayRoutePropagationSpec `json"spec"`
}

type AwsVpnGatewayRoutePropagationSpec struct {
	VpnGatewayId string `json:"vpn_gateway_id"`
	RouteTableId string `json:"route_table_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayRoutePropagationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpnGatewayRoutePropagation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlb struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAlbSpec `json"spec"`
}

type AwsAlbSpec struct {
	NamePrefix                   string                  `json:"name_prefix"`
	SecurityGroups               string                  `json:"security_groups"`
	AccessLogs                   []AwsAlbSpecAccessLogs  `json:"access_logs"`
	Tags                         map[string]string       `json:"tags"`
	Arn                          string                  `json:"arn"`
	Name                         string                  `json:"name"`
	IdleTimeout                  int                     `json:"idle_timeout"`
	DnsName                      string                  `json:"dns_name"`
	LoadBalancerType             string                  `json:"load_balancer_type"`
	EnableDeletionProtection     bool                    `json:"enable_deletion_protection"`
	VpcId                        string                  `json:"vpc_id"`
	ArnSuffix                    string                  `json:"arn_suffix"`
	EnableCrossZoneLoadBalancing bool                    `json:"enable_cross_zone_load_balancing"`
	SubnetMapping                AwsAlbSpecSubnetMapping `json:"subnet_mapping"`
	EnableHttp2                  bool                    `json:"enable_http2"`
	IpAddressType                string                  `json:"ip_address_type"`
	ZoneId                       string                  `json:"zone_id"`
	Internal                     bool                    `json:"internal"`
	Subnets                      string                  `json:"subnets"`
}

type AwsAlbSpecAccessLogs struct {
	Enabled bool   `json:"enabled"`
	Bucket  string `json:"bucket"`
	Prefix  string `json:"prefix"`
}

type AwsAlbSpecSubnetMapping struct {
	AllocationId string `json:"allocation_id"`
	SubnetId     string `json:"subnet_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAlb `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppsyncGraphqlApi struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAppsyncGraphqlApiSpec `json"spec"`
}

type AwsAppsyncGraphqlApiSpec struct {
	AuthenticationType string                                   `json:"authentication_type"`
	Name               string                                   `json:"name"`
	UserPoolConfig     []AwsAppsyncGraphqlApiSpecUserPoolConfig `json:"user_pool_config"`
	Arn                string                                   `json:"arn"`
}

type AwsAppsyncGraphqlApiSpecUserPoolConfig struct {
	AppIdClientRegex string `json:"app_id_client_regex"`
	AwsRegion        string `json:"aws_region"`
	DefaultAction    string `json:"default_action"`
	UserPoolId       string `json:"user_pool_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppsyncGraphqlApiList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAppsyncGraphqlApi `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkAcl struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsNetworkAclSpec `json"spec"`
}

type AwsNetworkAclSpec struct {
	Tags      map[string]string        `json:"tags"`
	VpcId     string                   `json:"vpc_id"`
	SubnetId  string                   `json:"subnet_id"`
	SubnetIds string                   `json:"subnet_ids"`
	Ingress   AwsNetworkAclSpecIngress `json:"ingress"`
	Egress    AwsNetworkAclSpecEgress  `json:"egress"`
}

type AwsNetworkAclSpecIngress struct {
	FromPort      int    `json:"from_port"`
	ToPort        int    `json:"to_port"`
	Action        string `json:"action"`
	Protocol      string `json:"protocol"`
	RuleNo        int    `json:"rule_no"`
	CidrBlock     string `json:"cidr_block"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	IcmpType      int    `json:"icmp_type"`
	IcmpCode      int    `json:"icmp_code"`
}

type AwsNetworkAclSpecEgress struct {
	RuleNo        int    `json:"rule_no"`
	Protocol      string `json:"protocol"`
	CidrBlock     string `json:"cidr_block"`
	IcmpType      int    `json:"icmp_type"`
	FromPort      int    `json:"from_port"`
	ToPort        int    `json:"to_port"`
	Action        string `json:"action"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	IcmpCode      int    `json:"icmp_code"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkAclList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsNetworkAcl `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterfaceAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsNetworkInterfaceAttachmentSpec `json"spec"`
}

type AwsNetworkInterfaceAttachmentSpec struct {
	NetworkInterfaceId string `json:"network_interface_id"`
	AttachmentId       string `json:"attachment_id"`
	Status             string `json:"status"`
	DeviceIndex        int    `json:"device_index"`
	InstanceId         string `json:"instance_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterfaceAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsNetworkInterfaceAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksCustomLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksCustomLayerSpec `json"spec"`
}

type AwsOpsworksCustomLayerSpec struct {
	AutoAssignPublicIps      bool                                `json:"auto_assign_public_ips"`
	UseEbsOptimizedInstances bool                                `json:"use_ebs_optimized_instances"`
	ShortName                string                              `json:"short_name"`
	ElasticLoadBalancer      string                              `json:"elastic_load_balancer"`
	CustomDeployRecipes      []string                            `json:"custom_deploy_recipes"`
	CustomShutdownRecipes    []string                            `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds   string                              `json:"custom_security_group_ids"`
	CustomJson               string                              `json:"custom_json"`
	SystemPackages           string                              `json:"system_packages"`
	AutoAssignElasticIps     bool                                `json:"auto_assign_elastic_ips"`
	AutoHealing              bool                                `json:"auto_healing"`
	InstanceShutdownTimeout  int                                 `json:"instance_shutdown_timeout"`
	Name                     string                              `json:"name"`
	StackId                  string                              `json:"stack_id"`
	EbsVolume                AwsOpsworksCustomLayerSpecEbsVolume `json:"ebs_volume"`
	CustomInstanceProfileArn string                              `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string                            `json:"custom_setup_recipes"`
	CustomConfigureRecipes   []string                            `json:"custom_configure_recipes"`
	CustomUndeployRecipes    []string                            `json:"custom_undeploy_recipes"`
	InstallUpdatesOnBoot     bool                                `json:"install_updates_on_boot"`
	DrainElbOnShutdown       bool                                `json:"drain_elb_on_shutdown"`
}

type AwsOpsworksCustomLayerSpecEbsVolume struct {
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksCustomLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksCustomLayer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterfaceSgAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsNetworkInterfaceSgAttachmentSpec `json"spec"`
}

type AwsNetworkInterfaceSgAttachmentSpec struct {
	SecurityGroupId    string `json:"security_group_id"`
	NetworkInterfaceId string `json:"network_interface_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterfaceSgAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsNetworkInterfaceSgAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSfnActivity struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSfnActivitySpec `json"spec"`
}

type AwsSfnActivitySpec struct {
	Name         string `json:"name"`
	CreationDate string `json:"creation_date"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSfnActivityList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSfnActivity `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbTargetGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLbTargetGroupSpec `json"spec"`
}

type AwsLbTargetGroupSpec struct {
	Arn                 string                            `json:"arn"`
	ArnSuffix           string                            `json:"arn_suffix"`
	NamePrefix          string                            `json:"name_prefix"`
	Port                int                               `json:"port"`
	HealthCheck         []AwsLbTargetGroupSpecHealthCheck `json:"health_check"`
	Tags                map[string]string                 `json:"tags"`
	Name                string                            `json:"name"`
	Protocol            string                            `json:"protocol"`
	VpcId               string                            `json:"vpc_id"`
	DeregistrationDelay int                               `json:"deregistration_delay"`
	TargetType          string                            `json:"target_type"`
	Stickiness          []AwsLbTargetGroupSpecStickiness  `json:"stickiness"`
}

type AwsLbTargetGroupSpecHealthCheck struct {
	Protocol           string `json:"protocol"`
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Interval           int    `json:"interval"`
	Path               string `json:"path"`
	Port               string `json:"port"`
}

type AwsLbTargetGroupSpecStickiness struct {
	Type           string `json:"type"`
	CookieDuration int    `json:"cookie_duration"`
	Enabled        bool   `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbTargetGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLbTargetGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiLaunchPermission struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAmiLaunchPermissionSpec `json"spec"`
}

type AwsAmiLaunchPermissionSpec struct {
	ImageId   string `json:"image_id"`
	AccountId string `json:"account_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiLaunchPermissionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAmiLaunchPermission `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueConnection struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGlueConnectionSpec `json"spec"`
}

type AwsGlueConnectionSpec struct {
	Name                           string                                                `json:"name"`
	PhysicalConnectionRequirements []AwsGlueConnectionSpecPhysicalConnectionRequirements `json:"physical_connection_requirements"`
	CatalogId                      string                                                `json:"catalog_id"`
	ConnectionProperties           map[string]string                                     `json:"connection_properties"`
	ConnectionType                 string                                                `json:"connection_type"`
	Description                    string                                                `json:"description"`
	MatchCriteria                  []string                                              `json:"match_criteria"`
}

type AwsGlueConnectionSpecPhysicalConnectionRequirements struct {
	SecurityGroupIdList []string `json:"security_group_id_list"`
	SubnetId            string   `json:"subnet_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueConnectionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGlueConnection `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccessKey struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamAccessKeySpec `json"spec"`
}

type AwsIamAccessKeySpec struct {
	SesSmtpPassword string `json:"ses_smtp_password"`
	PgpKey          string `json:"pgp_key"`
	KeyFingerprint  string `json:"key_fingerprint"`
	EncryptedSecret string `json:"encrypted_secret"`
	User            string `json:"user"`
	Status          string `json:"status"`
	Secret          string `json:"secret"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccessKeyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamAccessKey `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsAccount struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOrganizationsAccountSpec `json"spec"`
}

type AwsOrganizationsAccountSpec struct {
	Arn                    string `json:"arn"`
	JoinedMethod           string `json:"joined_method"`
	JoinedTimestamp        string `json:"joined_timestamp"`
	Status                 string `json:"status"`
	Name                   string `json:"name"`
	Email                  string `json:"email"`
	IamUserAccessToBilling string `json:"iam_user_access_to_billing"`
	RoleName               string `json:"role_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsAccountList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOrganizationsAccount `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbListener struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAlbListenerSpec `json"spec"`
}

type AwsAlbListenerSpec struct {
	SslPolicy       string                            `json:"ssl_policy"`
	CertificateArn  string                            `json:"certificate_arn"`
	DefaultAction   []AwsAlbListenerSpecDefaultAction `json:"default_action"`
	Arn             string                            `json:"arn"`
	LoadBalancerArn string                            `json:"load_balancer_arn"`
	Port            int                               `json:"port"`
	Protocol        string                            `json:"protocol"`
}

type AwsAlbListenerSpecDefaultAction struct {
	TargetGroupArn string `json:"target_group_arn"`
	Type           string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbListenerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAlbListener `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbListenerRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAlbListenerRuleSpec `json"spec"`
}

type AwsAlbListenerRuleSpec struct {
	Action      []AwsAlbListenerRuleSpecAction  `json:"action"`
	Condition   AwsAlbListenerRuleSpecCondition `json:"condition"`
	Arn         string                          `json:"arn"`
	ListenerArn string                          `json:"listener_arn"`
	Priority    int                             `json:"priority"`
}

type AwsAlbListenerRuleSpecAction struct {
	TargetGroupArn string `json:"target_group_arn"`
	Type           string `json:"type"`
}

type AwsAlbListenerRuleSpecCondition struct {
	Field  string   `json:"field"`
	Values []string `json:"values"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbListenerRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAlbListenerRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAcmCertificateValidation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAcmCertificateValidationSpec `json"spec"`
}

type AwsAcmCertificateValidationSpec struct {
	ValidationRecordFqdns string `json:"validation_record_fqdns"`
	CertificateArn        string `json:"certificate_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAcmCertificateValidationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAcmCertificateValidation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodedeployDeploymentGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCodedeployDeploymentGroupSpec `json"spec"`
}

type AwsCodedeployDeploymentGroupSpec struct {
	ServiceRoleArn              string                                                      `json:"service_role_arn"`
	AlarmConfiguration          []AwsCodedeployDeploymentGroupSpecAlarmConfiguration        `json:"alarm_configuration"`
	AutoRollbackConfiguration   []AwsCodedeployDeploymentGroupSpecAutoRollbackConfiguration `json:"auto_rollback_configuration"`
	AutoscalingGroups           string                                                      `json:"autoscaling_groups"`
	DeploymentConfigName        string                                                      `json:"deployment_config_name"`
	Ec2TagFilter                AwsCodedeployDeploymentGroupSpecEc2TagFilter                `json:"ec2_tag_filter"`
	TriggerConfiguration        AwsCodedeployDeploymentGroupSpecTriggerConfiguration        `json:"trigger_configuration"`
	AppName                     string                                                      `json:"app_name"`
	DeploymentStyle             []AwsCodedeployDeploymentGroupSpecDeploymentStyle           `json:"deployment_style"`
	BlueGreenDeploymentConfig   []AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfig `json:"blue_green_deployment_config"`
	LoadBalancerInfo            []AwsCodedeployDeploymentGroupSpecLoadBalancerInfo          `json:"load_balancer_info"`
	DeploymentGroupName         string                                                      `json:"deployment_group_name"`
	Ec2TagSet                   AwsCodedeployDeploymentGroupSpecEc2TagSet                   `json:"ec2_tag_set"`
	OnPremisesInstanceTagFilter AwsCodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter `json:"on_premises_instance_tag_filter"`
}

type AwsCodedeployDeploymentGroupSpecAlarmConfiguration struct {
	Alarms                 string `json:"alarms"`
	Enabled                bool   `json:"enabled"`
	IgnorePollAlarmFailure bool   `json:"ignore_poll_alarm_failure"`
}

type AwsCodedeployDeploymentGroupSpecAutoRollbackConfiguration struct {
	Enabled bool   `json:"enabled"`
	Events  string `json:"events"`
}

type AwsCodedeployDeploymentGroupSpecEc2TagFilter struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AwsCodedeployDeploymentGroupSpecTriggerConfiguration struct {
	TriggerEvents    string `json:"trigger_events"`
	TriggerName      string `json:"trigger_name"`
	TriggerTargetArn string `json:"trigger_target_arn"`
}

type AwsCodedeployDeploymentGroupSpecDeploymentStyle struct {
	DeploymentOption string `json:"deployment_option"`
	DeploymentType   string `json:"deployment_type"`
}

type AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfig struct {
	GreenFleetProvisioningOption              []AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption              `json:"green_fleet_provisioning_option"`
	TerminateBlueInstancesOnDeploymentSuccess []AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess `json:"terminate_blue_instances_on_deployment_success"`
	DeploymentReadyOption                     []AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption                     `json:"deployment_ready_option"`
}

type AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption struct {
	Action string `json:"action"`
}

type AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess struct {
	Action                       string `json:"action"`
	TerminationWaitTimeInMinutes int    `json:"termination_wait_time_in_minutes"`
}

type AwsCodedeployDeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption struct {
	ActionOnTimeout   string `json:"action_on_timeout"`
	WaitTimeInMinutes int    `json:"wait_time_in_minutes"`
}

type AwsCodedeployDeploymentGroupSpecLoadBalancerInfo struct {
	ElbInfo         AwsCodedeployDeploymentGroupSpecLoadBalancerInfoElbInfo         `json:"elb_info"`
	TargetGroupInfo AwsCodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupInfo `json:"target_group_info"`
}

type AwsCodedeployDeploymentGroupSpecLoadBalancerInfoElbInfo struct {
	Name string `json:"name"`
}

type AwsCodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupInfo struct {
	Name string `json:"name"`
}

type AwsCodedeployDeploymentGroupSpecEc2TagSet struct {
	Ec2TagFilter AwsCodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter `json:"ec2_tag_filter"`
}

type AwsCodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AwsCodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodedeployDeploymentGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCodedeployDeploymentGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRole struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamRoleSpec `json"spec"`
}

type AwsIamRoleSpec struct {
	MaxSessionDuration  int    `json:"max_session_duration"`
	Arn                 string `json:"arn"`
	NamePrefix          string `json:"name_prefix"`
	Description         string `json:"description"`
	AssumeRolePolicy    string `json:"assume_role_policy"`
	ForceDetachPolicies bool   `json:"force_detach_policies"`
	CreateDate          string `json:"create_date"`
	UniqueId            string `json:"unique_id"`
	Name                string `json:"name"`
	Path                string `json:"path"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRoleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamRole `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksInstance struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksInstanceSpec `json"spec"`
}

type AwsOpsworksInstanceSpec struct {
	AvailabilityZone         string                                      `json:"availability_zone"`
	EbsOptimized             bool                                        `json:"ebs_optimized"`
	InfrastructureClass      string                                      `json:"infrastructure_class"`
	ReportedOsFamily         string                                      `json:"reported_os_family"`
	ReportedOsName           string                                      `json:"reported_os_name"`
	SshKeyName               string                                      `json:"ssh_key_name"`
	AgentVersion             string                                      `json:"agent_version"`
	Os                       string                                      `json:"os"`
	RegisteredBy             string                                      `json:"registered_by"`
	ReportedAgentVersion     string                                      `json:"reported_agent_version"`
	Ec2InstanceId            string                                      `json:"ec2_instance_id"`
	DeleteEbs                bool                                        `json:"delete_ebs"`
	ElasticIp                string                                      `json:"elastic_ip"`
	LayerIds                 []string                                    `json:"layer_ids"`
	Platform                 string                                      `json:"platform"`
	StackId                  string                                      `json:"stack_id"`
	Architecture             string                                      `json:"architecture"`
	SshHostRsaKeyFingerprint string                                      `json:"ssh_host_rsa_key_fingerprint"`
	Tenancy                  string                                      `json:"tenancy"`
	EbsBlockDevice           AwsOpsworksInstanceSpecEbsBlockDevice       `json:"ebs_block_device"`
	ReportedOsVersion        string                                      `json:"reported_os_version"`
	DeleteEip                bool                                        `json:"delete_eip"`
	EcsClusterArn            string                                      `json:"ecs_cluster_arn"`
	RootDeviceVolumeId       string                                      `json:"root_device_volume_id"`
	State                    string                                      `json:"state"`
	Status                   string                                      `json:"status"`
	SubnetId                 string                                      `json:"subnet_id"`
	EphemeralBlockDevice     AwsOpsworksInstanceSpecEphemeralBlockDevice `json:"ephemeral_block_device"`
	AmiId                    string                                      `json:"ami_id"`
	RootBlockDevice          AwsOpsworksInstanceSpecRootBlockDevice      `json:"root_block_device"`
	PublicIp                 string                                      `json:"public_ip"`
	SecurityGroupIds         []string                                    `json:"security_group_ids"`
	SshHostDsaKeyFingerprint string                                      `json:"ssh_host_dsa_key_fingerprint"`
	VirtualizationType       string                                      `json:"virtualization_type"`
	InstanceProfileArn       string                                      `json:"instance_profile_arn"`
	LastServiceErrorId       string                                      `json:"last_service_error_id"`
	PrivateIp                string                                      `json:"private_ip"`
	PublicDns                string                                      `json:"public_dns"`
	Hostname                 string                                      `json:"hostname"`
	CreatedAt                string                                      `json:"created_at"`
	InstallUpdatesOnBoot     bool                                        `json:"install_updates_on_boot"`
	InstanceType             string                                      `json:"instance_type"`
	PrivateDns               string                                      `json:"private_dns"`
	RootDeviceType           string                                      `json:"root_device_type"`
	AutoScalingType          string                                      `json:"auto_scaling_type"`
}

type AwsOpsworksInstanceSpecEbsBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsOpsworksInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsOpsworksInstanceSpecRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksInstanceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksInstance `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketNotification struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsS3BucketNotificationSpec `json"spec"`
}

type AwsS3BucketNotificationSpec struct {
	Bucket         string                                      `json:"bucket"`
	Topic          []AwsS3BucketNotificationSpecTopic          `json:"topic"`
	Queue          []AwsS3BucketNotificationSpecQueue          `json:"queue"`
	LambdaFunction []AwsS3BucketNotificationSpecLambdaFunction `json:"lambda_function"`
}

type AwsS3BucketNotificationSpecTopic struct {
	Events       string `json:"events"`
	Id           string `json:"id"`
	FilterPrefix string `json:"filter_prefix"`
	FilterSuffix string `json:"filter_suffix"`
	TopicArn     string `json:"topic_arn"`
}

type AwsS3BucketNotificationSpecQueue struct {
	QueueArn     string `json:"queue_arn"`
	Events       string `json:"events"`
	Id           string `json:"id"`
	FilterPrefix string `json:"filter_prefix"`
	FilterSuffix string `json:"filter_suffix"`
}

type AwsS3BucketNotificationSpecLambdaFunction struct {
	FilterSuffix      string `json:"filter_suffix"`
	LambdaFunctionArn string `json:"lambda_function_arn"`
	Events            string `json:"events"`
	Id                string `json:"id"`
	FilterPrefix      string `json:"filter_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketNotificationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsS3BucketNotification `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafGeoMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafGeoMatchSetSpec `json"spec"`
}

type AwsWafGeoMatchSetSpec struct {
	GeoMatchConstraint AwsWafGeoMatchSetSpecGeoMatchConstraint `json:"geo_match_constraint"`
	Name               string                                  `json:"name"`
}

type AwsWafGeoMatchSetSpecGeoMatchConstraint struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafGeoMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafGeoMatchSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegrationResponse struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayIntegrationResponseSpec `json"spec"`
}

type AwsApiGatewayIntegrationResponseSpec struct {
	RestApiId                string            `json:"rest_api_id"`
	StatusCode               string            `json:"status_code"`
	ResponseTemplates        map[string]string `json:"response_templates"`
	ResponseParameters       map[string]string `json:"response_parameters"`
	ResponseParametersInJson string            `json:"response_parameters_in_json"`
	ResourceId               string            `json:"resource_id"`
	HttpMethod               string            `json:"http_method"`
	SelectionPattern         string            `json:"selection_pattern"`
	ContentHandling          string            `json:"content_handling"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegrationResponseList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayIntegrationResponse `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventTarget struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchEventTargetSpec `json"spec"`
}

type AwsCloudwatchEventTargetSpec struct {
	Arn               string                                          `json:"arn"`
	Input             string                                          `json:"input"`
	RunCommandTargets []AwsCloudwatchEventTargetSpecRunCommandTargets `json:"run_command_targets"`
	BatchTarget       []AwsCloudwatchEventTargetSpecBatchTarget       `json:"batch_target"`
	KinesisTarget     []AwsCloudwatchEventTargetSpecKinesisTarget     `json:"kinesis_target"`
	SqsTarget         []AwsCloudwatchEventTargetSpecSqsTarget         `json:"sqs_target"`
	InputTransformer  []AwsCloudwatchEventTargetSpecInputTransformer  `json:"input_transformer"`
	TargetId          string                                          `json:"target_id"`
	InputPath         string                                          `json:"input_path"`
	RoleArn           string                                          `json:"role_arn"`
	EcsTarget         []AwsCloudwatchEventTargetSpecEcsTarget         `json:"ecs_target"`
	Rule              string                                          `json:"rule"`
}

type AwsCloudwatchEventTargetSpecRunCommandTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type AwsCloudwatchEventTargetSpecBatchTarget struct {
	JobDefinition string `json:"job_definition"`
	JobName       string `json:"job_name"`
	ArraySize     int    `json:"array_size"`
	JobAttempts   int    `json:"job_attempts"`
}

type AwsCloudwatchEventTargetSpecKinesisTarget struct {
	PartitionKeyPath string `json:"partition_key_path"`
}

type AwsCloudwatchEventTargetSpecSqsTarget struct {
	MessageGroupId string `json:"message_group_id"`
}

type AwsCloudwatchEventTargetSpecInputTransformer struct {
	InputPaths    map[string]string `json:"input_paths"`
	InputTemplate string            `json:"input_template"`
}

type AwsCloudwatchEventTargetSpecEcsTarget struct {
	TaskCount         int    `json:"task_count"`
	TaskDefinitionArn string `json:"task_definition_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventTargetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchEventTarget `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCognitoUserGroupSpec `json"spec"`
}

type AwsCognitoUserGroupSpec struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Precedence  int    `json:"precedence"`
	RoleArn     string `json:"role_arn"`
	UserPoolId  string `json:"user_pool_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCognitoUserGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbTableItem struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDynamodbTableItemSpec `json"spec"`
}

type AwsDynamodbTableItemSpec struct {
	TableName string `json:"table_name"`
	HashKey   string `json:"hash_key"`
	RangeKey  string `json:"range_key"`
	Item      string `json:"item"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbTableItemList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDynamodbTableItem `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalIpset struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalIpsetSpec `json"spec"`
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
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalIpset `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalRuleSpec `json"spec"`
}

type AwsWafregionalRuleSpec struct {
	Name       string                          `json:"name"`
	MetricName string                          `json:"metric_name"`
	Predicate  AwsWafregionalRuleSpecPredicate `json:"predicate"`
}

type AwsWafregionalRuleSpecPredicate struct {
	Type    string `json:"type"`
	Negated bool   `json:"negated"`
	DataId  string `json:"data_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayRestApi struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayRestApiSpec `json"spec"`
}

type AwsApiGatewayRestApiSpec struct {
	MinimumCompressionSize int      `json:"minimum_compression_size"`
	RootResourceId         string   `json:"root_resource_id"`
	CreatedDate            string   `json:"created_date"`
	Name                   string   `json:"name"`
	Description            string   `json:"description"`
	Policy                 string   `json:"policy"`
	BinaryMediaTypes       []string `json:"binary_media_types"`
	Body                   string   `json:"body"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayRestApiList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayRestApi `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogDestinationPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudwatchLogDestinationPolicySpec `json"spec"`
}

type AwsCloudwatchLogDestinationPolicySpec struct {
	AccessPolicy    string `json:"access_policy"`
	DestinationName string `json:"destination_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogDestinationPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudwatchLogDestinationPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInstance struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsInstanceSpec `json"spec"`
}

type AwsInstanceSpec struct {
	IamInstanceProfile                string                               `json:"iam_instance_profile"`
	PasswordData                      string                               `json:"password_data"`
	PrimaryNetworkInterfaceId         string                               `json:"primary_network_interface_id"`
	InstanceInitiatedShutdownBehavior string                               `json:"instance_initiated_shutdown_behavior"`
	BlockDevice                       map[string]string                    `json:"block_device"`
	AssociatePublicIpAddress          bool                                 `json:"associate_public_ip_address"`
	KeyName                           string                               `json:"key_name"`
	SubnetId                          string                               `json:"subnet_id"`
	Ipv6AddressCount                  int                                  `json:"ipv6_address_count"`
	Tags                              map[string]string                    `json:"tags"`
	EbsBlockDevice                    AwsInstanceSpecEbsBlockDevice        `json:"ebs_block_device"`
	Ami                               string                               `json:"ami"`
	PrivateIp                         string                               `json:"private_ip"`
	NetworkInterfaceId                string                               `json:"network_interface_id"`
	VolumeTags                        map[string]string                    `json:"volume_tags"`
	EphemeralBlockDevice              AwsInstanceSpecEphemeralBlockDevice  `json:"ephemeral_block_device"`
	AvailabilityZone                  string                               `json:"availability_zone"`
	PlacementGroup                    string                               `json:"placement_group"`
	VpcSecurityGroupIds               string                               `json:"vpc_security_group_ids"`
	PublicDns                         string                               `json:"public_dns"`
	NetworkInterface                  AwsInstanceSpecNetworkInterface      `json:"network_interface"`
	PublicIp                          string                               `json:"public_ip"`
	Monitoring                        bool                                 `json:"monitoring"`
	InstanceType                      string                               `json:"instance_type"`
	SourceDestCheck                   bool                                 `json:"source_dest_check"`
	UserData                          string                               `json:"user_data"`
	UserDataBase64                    string                               `json:"user_data_base64"`
	SecurityGroups                    string                               `json:"security_groups"`
	RootBlockDevice                   []AwsInstanceSpecRootBlockDevice     `json:"root_block_device"`
	CreditSpecification               []AwsInstanceSpecCreditSpecification `json:"credit_specification"`
	InstanceState                     string                               `json:"instance_state"`
	EbsOptimized                      bool                                 `json:"ebs_optimized"`
	DisableApiTermination             bool                                 `json:"disable_api_termination"`
	GetPasswordData                   bool                                 `json:"get_password_data"`
	PrivateDns                        string                               `json:"private_dns"`
	Ipv6Addresses                     []string                             `json:"ipv6_addresses"`
	Tenancy                           string                               `json:"tenancy"`
}

type AwsInstanceSpecEbsBlockDevice struct {
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type AwsInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
	NoDevice    bool   `json:"no_device"`
}

type AwsInstanceSpecNetworkInterface struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	NetworkInterfaceId  string `json:"network_interface_id"`
	DeviceIndex         int    `json:"device_index"`
}

type AwsInstanceSpecRootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
}

type AwsInstanceSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInstanceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsInstance `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksMemcachedLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksMemcachedLayerSpec `json"spec"`
}

type AwsOpsworksMemcachedLayerSpec struct {
	InstanceShutdownTimeout  int                                    `json:"instance_shutdown_timeout"`
	DrainElbOnShutdown       bool                                   `json:"drain_elb_on_shutdown"`
	SystemPackages           string                                 `json:"system_packages"`
	StackId                  string                                 `json:"stack_id"`
	UseEbsOptimizedInstances bool                                   `json:"use_ebs_optimized_instances"`
	AutoAssignPublicIps      bool                                   `json:"auto_assign_public_ips"`
	CustomConfigureRecipes   []string                               `json:"custom_configure_recipes"`
	CustomSecurityGroupIds   string                                 `json:"custom_security_group_ids"`
	EbsVolume                AwsOpsworksMemcachedLayerSpecEbsVolume `json:"ebs_volume"`
	AllocatedMemory          int                                    `json:"allocated_memory"`
	CustomDeployRecipes      []string                               `json:"custom_deploy_recipes"`
	InstallUpdatesOnBoot     bool                                   `json:"install_updates_on_boot"`
	Name                     string                                 `json:"name"`
	AutoAssignElasticIps     bool                                   `json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer      string                                 `json:"elastic_load_balancer"`
	CustomShutdownRecipes    []string                               `json:"custom_shutdown_recipes"`
	CustomJson               string                                 `json:"custom_json"`
	AutoHealing              bool                                   `json:"auto_healing"`
	CustomInstanceProfileArn string                                 `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string                               `json:"custom_setup_recipes"`
	CustomUndeployRecipes    []string                               `json:"custom_undeploy_recipes"`
}

type AwsOpsworksMemcachedLayerSpecEbsVolume struct {
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksMemcachedLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksMemcachedLayer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesEventDestination struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesEventDestinationSpec `json"spec"`
}

type AwsSesEventDestinationSpec struct {
	SnsDestination        AwsSesEventDestinationSpecSnsDestination        `json:"sns_destination"`
	Name                  string                                          `json:"name"`
	ConfigurationSetName  string                                          `json:"configuration_set_name"`
	Enabled               bool                                            `json:"enabled"`
	MatchingTypes         string                                          `json:"matching_types"`
	CloudwatchDestination AwsSesEventDestinationSpecCloudwatchDestination `json:"cloudwatch_destination"`
	KinesisDestination    AwsSesEventDestinationSpecKinesisDestination    `json:"kinesis_destination"`
}

type AwsSesEventDestinationSpecSnsDestination struct {
	TopicArn string `json:"topic_arn"`
}

type AwsSesEventDestinationSpecCloudwatchDestination struct {
	DefaultValue  string `json:"default_value"`
	DimensionName string `json:"dimension_name"`
	ValueSource   string `json:"value_source"`
}

type AwsSesEventDestinationSpecKinesisDestination struct {
	StreamArn string `json:"stream_arn"`
	RoleArn   string `json:"role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesEventDestinationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesEventDestination `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserSshKey struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamUserSshKeySpec `json"spec"`
}

type AwsIamUserSshKeySpec struct {
	Status         string `json:"status"`
	SshPublicKeyId string `json:"ssh_public_key_id"`
	Fingerprint    string `json:"fingerprint"`
	Username       string `json:"username"`
	PublicKey      string `json:"public_key"`
	Encoding       string `json:"encoding"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserSshKeyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamUserSshKey `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotCertificate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIotCertificateSpec `json"spec"`
}

type AwsIotCertificateSpec struct {
	Csr    string `json:"csr"`
	Active bool   `json:"active"`
	Arn    string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotCertificateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIotCertificate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiFromInstance struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAmiFromInstanceSpec `json"spec"`
}

type AwsAmiFromInstanceSpec struct {
	ImageLocation         string                                     `json:"image_location"`
	Architecture          string                                     `json:"architecture"`
	RootDeviceName        string                                     `json:"root_device_name"`
	SriovNetSupport       string                                     `json:"sriov_net_support"`
	EphemeralBlockDevice  AwsAmiFromInstanceSpecEphemeralBlockDevice `json:"ephemeral_block_device"`
	RamdiskId             string                                     `json:"ramdisk_id"`
	Tags                  map[string]string                          `json:"tags"`
	ManageEbsSnapshots    bool                                       `json:"manage_ebs_snapshots"`
	Description           string                                     `json:"description"`
	KernelId              string                                     `json:"kernel_id"`
	Name                  string                                     `json:"name"`
	RootSnapshotId        string                                     `json:"root_snapshot_id"`
	EbsBlockDevice        AwsAmiFromInstanceSpecEbsBlockDevice       `json:"ebs_block_device"`
	SourceInstanceId      string                                     `json:"source_instance_id"`
	VirtualizationType    string                                     `json:"virtualization_type"`
	SnapshotWithoutReboot bool                                       `json:"snapshot_without_reboot"`
}

type AwsAmiFromInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsAmiFromInstanceSpecEbsBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiFromInstanceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAmiFromInstance `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayModel struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayModelSpec `json"spec"`
}

type AwsApiGatewayModelSpec struct {
	Schema      string `json:"schema"`
	ContentType string `json:"content_type"`
	RestApiId   string `json:"rest_api_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayModelList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayModel `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbParameterGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDbParameterGroupSpec `json"spec"`
}

type AwsDbParameterGroupSpec struct {
	Description string                           `json:"description"`
	Parameter   AwsDbParameterGroupSpecParameter `json:"parameter"`
	Tags        map[string]string                `json:"tags"`
	Arn         string                           `json:"arn"`
	Name        string                           `json:"name"`
	NamePrefix  string                           `json:"name_prefix"`
	Family      string                           `json:"family"`
}

type AwsDbParameterGroupSpecParameter struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	ApplyMethod string `json:"apply_method"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbParameterGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDbParameterGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlacierVault struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGlacierVaultSpec `json"spec"`
}

type AwsGlacierVaultSpec struct {
	AccessPolicy string                            `json:"access_policy"`
	Notification []AwsGlacierVaultSpecNotification `json:"notification"`
	Tags         map[string]string                 `json:"tags"`
	Name         string                            `json:"name"`
	Location     string                            `json:"location"`
	Arn          string                            `json:"arn"`
}

type AwsGlacierVaultSpecNotification struct {
	Events   string `json:"events"`
	SnsTopic string `json:"sns_topic"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlacierVaultList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGlacierVault `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccountPasswordPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamAccountPasswordPolicySpec `json"spec"`
}

type AwsIamAccountPasswordPolicySpec struct {
	HardExpiry                 bool `json:"hard_expiry"`
	MaxPasswordAge             int  `json:"max_password_age"`
	MinimumPasswordLength      int  `json:"minimum_password_length"`
	RequireSymbols             bool `json:"require_symbols"`
	RequireUppercaseCharacters bool `json:"require_uppercase_characters"`
	AllowUsersToChangePassword bool `json:"allow_users_to_change_password"`
	ExpirePasswords            bool `json:"expire_passwords"`
	PasswordReusePrevention    int  `json:"password_reuse_prevention"`
	RequireLowercaseCharacters bool `json:"require_lowercase_characters"`
	RequireNumbers             bool `json:"require_numbers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccountPasswordPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamAccountPasswordPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupMembership struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamGroupMembershipSpec `json"spec"`
}

type AwsIamGroupMembershipSpec struct {
	Name  string `json:"name"`
	Users string `json:"users"`
	Group string `json:"group"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupMembershipList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamGroupMembership `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftParameterGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsRedshiftParameterGroupSpec `json"spec"`
}

type AwsRedshiftParameterGroupSpec struct {
	Family      string                                 `json:"family"`
	Description string                                 `json:"description"`
	Parameter   AwsRedshiftParameterGroupSpecParameter `json:"parameter"`
	Name        string                                 `json:"name"`
}

type AwsRedshiftParameterGroupSpecParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftParameterGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsRedshiftParameterGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotDatafeedSubscription struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSpotDatafeedSubscriptionSpec `json"spec"`
}

type AwsSpotDatafeedSubscriptionSpec struct {
	Bucket string `json:"bucket"`
	Prefix string `json:"prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotDatafeedSubscriptionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSpotDatafeedSubscription `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGateway struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpnGatewaySpec `json"spec"`
}

type AwsVpnGatewaySpec struct {
	AvailabilityZone string            `json:"availability_zone"`
	AmazonSideAsn    string            `json:"amazon_side_asn"`
	VpcId            string            `json:"vpc_id"`
	Tags             map[string]string `json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpnGateway `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRegexPatternSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalRegexPatternSetSpec `json"spec"`
}

type AwsWafregionalRegexPatternSetSpec struct {
	Name                string `json:"name"`
	RegexPatternStrings string `json:"regex_pattern_strings"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRegexPatternSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalRegexPatternSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecretsmanagerSecretVersion struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSecretsmanagerSecretVersionSpec `json"spec"`
}

type AwsSecretsmanagerSecretVersionSpec struct {
	SecretString  string `json:"secret_string"`
	VersionId     string `json:"version_id"`
	VersionStages string `json:"version_stages"`
	SecretId      string `json:"secret_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecretsmanagerSecretVersionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSecretsmanagerSecretVersion `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotInstanceRequest struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSpotInstanceRequestSpec `json"spec"`
}

type AwsSpotInstanceRequestSpec struct {
	ValidFrom                         string                                          `json:"valid_from"`
	GetPasswordData                   bool                                            `json:"get_password_data"`
	WaitForFulfillment                bool                                            `json:"wait_for_fulfillment"`
	SpotInstanceId                    string                                          `json:"spot_instance_id"`
	PlacementGroup                    string                                          `json:"placement_group"`
	DisableApiTermination             bool                                            `json:"disable_api_termination"`
	EbsBlockDevice                    AwsSpotInstanceRequestSpecEbsBlockDevice        `json:"ebs_block_device"`
	LaunchGroup                       string                                          `json:"launch_group"`
	SubnetId                          string                                          `json:"subnet_id"`
	InstanceInitiatedShutdownBehavior string                                          `json:"instance_initiated_shutdown_behavior"`
	Ipv6AddressCount                  int                                             `json:"ipv6_address_count"`
	UserDataBase64                    string                                          `json:"user_data_base64"`
	NetworkInterfaceId                string                                          `json:"network_interface_id"`
	PrivateDns                        string                                          `json:"private_dns"`
	IamInstanceProfile                string                                          `json:"iam_instance_profile"`
	Tenancy                           string                                          `json:"tenancy"`
	Ami                               string                                          `json:"ami"`
	AvailabilityZone                  string                                          `json:"availability_zone"`
	SourceDestCheck                   bool                                            `json:"source_dest_check"`
	CreditSpecification               []AwsSpotInstanceRequestSpecCreditSpecification `json:"credit_specification"`
	SpotBidStatus                     string                                          `json:"spot_bid_status"`
	BlockDurationMinutes              int                                             `json:"block_duration_minutes"`
	SpotRequestState                  string                                          `json:"spot_request_state"`
	InstanceInterruptionBehaviour     string                                          `json:"instance_interruption_behaviour"`
	ValidUntil                        string                                          `json:"valid_until"`
	PublicIp                          string                                          `json:"public_ip"`
	Tags                              map[string]string                               `json:"tags"`
	BlockDevice                       map[string]string                               `json:"block_device"`
	Monitoring                        bool                                            `json:"monitoring"`
	Ipv6Addresses                     []string                                        `json:"ipv6_addresses"`
	VolumeTags                        map[string]string                               `json:"volume_tags"`
	EphemeralBlockDevice              AwsSpotInstanceRequestSpecEphemeralBlockDevice  `json:"ephemeral_block_device"`
	RootBlockDevice                   []AwsSpotInstanceRequestSpecRootBlockDevice     `json:"root_block_device"`
	PrivateIp                         string                                          `json:"private_ip"`
	UserData                          string                                          `json:"user_data"`
	EbsOptimized                      bool                                            `json:"ebs_optimized"`
	NetworkInterface                  AwsSpotInstanceRequestSpecNetworkInterface      `json:"network_interface"`
	InstanceState                     string                                          `json:"instance_state"`
	PasswordData                      string                                          `json:"password_data"`
	SecurityGroups                    string                                          `json:"security_groups"`
	PrimaryNetworkInterfaceId         string                                          `json:"primary_network_interface_id"`
	VpcSecurityGroupIds               string                                          `json:"vpc_security_group_ids"`
	PublicDns                         string                                          `json:"public_dns"`
	SpotPrice                         string                                          `json:"spot_price"`
	SpotType                          string                                          `json:"spot_type"`
	AssociatePublicIpAddress          bool                                            `json:"associate_public_ip_address"`
	InstanceType                      string                                          `json:"instance_type"`
	KeyName                           string                                          `json:"key_name"`
}

type AwsSpotInstanceRequestSpecEbsBlockDevice struct {
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
}

type AwsSpotInstanceRequestSpecCreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type AwsSpotInstanceRequestSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
	NoDevice    bool   `json:"no_device"`
}

type AwsSpotInstanceRequestSpecRootBlockDevice struct {
	VolumeType          string `json:"volume_type"`
	VolumeId            string `json:"volume_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	VolumeSize          int    `json:"volume_size"`
}

type AwsSpotInstanceRequestSpecNetworkInterface struct {
	NetworkInterfaceId  string `json:"network_interface_id"`
	DeviceIndex         int    `json:"device_index"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotInstanceRequestList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSpotInstanceRequest `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolDomain struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCognitoUserPoolDomainSpec `json"spec"`
}

type AwsCognitoUserPoolDomainSpec struct {
	AwsAccountId              string `json:"aws_account_id"`
	CloudfrontDistributionArn string `json:"cloudfront_distribution_arn"`
	S3Bucket                  string `json:"s3_bucket"`
	Version                   string `json:"version"`
	Domain                    string `json:"domain"`
	UserPoolId                string `json:"user_pool_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolDomainList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCognitoUserPoolDomain `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDaxParameterGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDaxParameterGroupSpec `json"spec"`
}

type AwsDaxParameterGroupSpec struct {
	Name        string                             `json:"name"`
	Description string                             `json:"description"`
	Parameters  AwsDaxParameterGroupSpecParameters `json:"parameters"`
}

type AwsDaxParameterGroupSpecParameters struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDaxParameterGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDaxParameterGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheSubnetGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticacheSubnetGroupSpec `json"spec"`
}

type AwsElasticacheSubnetGroupSpec struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	SubnetIds   string `json:"subnet_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheSubnetGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticacheSubnetGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyDetector struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGuarddutyDetectorSpec `json"spec"`
}

type AwsGuarddutyDetectorSpec struct {
	Enable    bool   `json:"enable"`
	AccountId string `json:"account_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyDetectorList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGuarddutyDetector `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksGangliaLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksGangliaLayerSpec `json"spec"`
}

type AwsOpsworksGangliaLayerSpec struct {
	ElasticLoadBalancer      string                               `json:"elastic_load_balancer"`
	CustomSecurityGroupIds   string                               `json:"custom_security_group_ids"`
	CustomJson               string                               `json:"custom_json"`
	InstanceShutdownTimeout  int                                  `json:"instance_shutdown_timeout"`
	Url                      string                               `json:"url"`
	CustomShutdownRecipes    []string                             `json:"custom_shutdown_recipes"`
	DrainElbOnShutdown       bool                                 `json:"drain_elb_on_shutdown"`
	StackId                  string                               `json:"stack_id"`
	EbsVolume                AwsOpsworksGangliaLayerSpecEbsVolume `json:"ebs_volume"`
	Name                     string                               `json:"name"`
	Username                 string                               `json:"username"`
	AutoAssignElasticIps     bool                                 `json:"auto_assign_elastic_ips"`
	CustomSetupRecipes       []string                             `json:"custom_setup_recipes"`
	CustomDeployRecipes      []string                             `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                             `json:"custom_undeploy_recipes"`
	AutoHealing              bool                                 `json:"auto_healing"`
	Password                 string                               `json:"password"`
	AutoAssignPublicIps      bool                                 `json:"auto_assign_public_ips"`
	CustomInstanceProfileArn string                               `json:"custom_instance_profile_arn"`
	CustomConfigureRecipes   []string                             `json:"custom_configure_recipes"`
	InstallUpdatesOnBoot     bool                                 `json:"install_updates_on_boot"`
	SystemPackages           string                               `json:"system_packages"`
	UseEbsOptimizedInstances bool                                 `json:"use_ebs_optimized_instances"`
}

type AwsOpsworksGangliaLayerSpecEbsVolume struct {
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksGangliaLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksGangliaLayer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsPolicyAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOrganizationsPolicyAttachmentSpec `json"spec"`
}

type AwsOrganizationsPolicyAttachmentSpec struct {
	PolicyId string `json:"policy_id"`
	TargetId string `json:"target_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsPolicyAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOrganizationsPolicyAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRuleGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafregionalRuleGroupSpec `json"spec"`
}

type AwsWafregionalRuleGroupSpec struct {
	MetricName    string                                   `json:"metric_name"`
	ActivatedRule AwsWafregionalRuleGroupSpecActivatedRule `json:"activated_rule"`
	Name          string                                   `json:"name"`
}

type AwsWafregionalRuleGroupSpecActivatedRule struct {
	Action   []AwsWafregionalRuleGroupSpecActivatedRuleAction `json:"action"`
	Priority int                                              `json:"priority"`
	RuleId   string                                           `json:"rule_id"`
	Type     string                                           `json:"type"`
}

type AwsWafregionalRuleGroupSpecActivatedRuleAction struct {
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRuleGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafregionalRuleGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryPublicDnsNamespace struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsServiceDiscoveryPublicDnsNamespaceSpec `json"spec"`
}

type AwsServiceDiscoveryPublicDnsNamespaceSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Arn         string `json:"arn"`
	HostedZone  string `json:"hosted_zone"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryPublicDnsNamespaceList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsServiceDiscoveryPublicDnsNamespace `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchJobQueue struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsBatchJobQueueSpec `json"spec"`
}

type AwsBatchJobQueueSpec struct {
	ComputeEnvironments []string `json:"compute_environments"`
	Name                string   `json:"name"`
	Priority            int      `json:"priority"`
	State               string   `json:"state"`
	Arn                 string   `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchJobQueueList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsBatchJobQueue `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayBasePathMapping struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayBasePathMappingSpec `json"spec"`
}

type AwsApiGatewayBasePathMappingSpec struct {
	ApiId      string `json:"api_id"`
	BasePath   string `json:"base_path"`
	StageName  string `json:"stage_name"`
	DomainName string `json:"domain_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayBasePathMappingList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayBasePathMapping `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudtrail struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCloudtrailSpec `json"spec"`
}

type AwsCloudtrailSpec struct {
	CloudWatchLogsGroupArn     string                           `json:"cloud_watch_logs_group_arn"`
	EnableLogFileValidation    bool                             `json:"enable_log_file_validation"`
	EventSelector              []AwsCloudtrailSpecEventSelector `json:"event_selector"`
	HomeRegion                 string                           `json:"home_region"`
	Tags                       map[string]string                `json:"tags"`
	CloudWatchLogsRoleArn      string                           `json:"cloud_watch_logs_role_arn"`
	IncludeGlobalServiceEvents bool                             `json:"include_global_service_events"`
	Arn                        string                           `json:"arn"`
	KmsKeyId                   string                           `json:"kms_key_id"`
	Name                       string                           `json:"name"`
	S3BucketName               string                           `json:"s3_bucket_name"`
	S3KeyPrefix                string                           `json:"s3_key_prefix"`
	IsMultiRegionTrail         bool                             `json:"is_multi_region_trail"`
	SnsTopicName               string                           `json:"sns_topic_name"`
	EnableLogging              bool                             `json:"enable_logging"`
}

type AwsCloudtrailSpecEventSelector struct {
	ReadWriteType           string                                       `json:"read_write_type"`
	IncludeManagementEvents bool                                         `json:"include_management_events"`
	DataResource            []AwsCloudtrailSpecEventSelectorDataResource `json:"data_resource"`
}

type AwsCloudtrailSpecEventSelectorDataResource struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudtrailList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCloudtrail `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserLoginProfile struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamUserLoginProfileSpec `json"spec"`
}

type AwsIamUserLoginProfileSpec struct {
	PasswordLength        int    `json:"password_length"`
	KeyFingerprint        string `json:"key_fingerprint"`
	EncryptedPassword     string `json:"encrypted_password"`
	User                  string `json:"user"`
	PgpKey                string `json:"pgp_key"`
	PasswordResetRequired bool   `json:"password_reset_required"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserLoginProfileList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamUserLoginProfile `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKeyPair struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsKeyPairSpec `json"spec"`
}

type AwsKeyPairSpec struct {
	Fingerprint   string `json:"fingerprint"`
	KeyName       string `json:"key_name"`
	KeyNamePrefix string `json:"key_name_prefix"`
	PublicKey     string `json:"public_key"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKeyPairList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsKeyPair `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksPermission struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksPermissionSpec `json"spec"`
}

type AwsOpsworksPermissionSpec struct {
	UserArn   string `json:"user_arn"`
	Level     string `json:"level"`
	StackId   string `json:"stack_id"`
	AllowSsh  bool   `json:"allow_ssh"`
	AllowSudo bool   `json:"allow_sudo"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksPermissionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksPermission `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecretsmanagerSecret struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSecretsmanagerSecretSpec `json"spec"`
}

type AwsSecretsmanagerSecretSpec struct {
	Name                 string                                     `json:"name"`
	RotationLambdaArn    string                                     `json:"rotation_lambda_arn"`
	RotationRules        []AwsSecretsmanagerSecretSpecRotationRules `json:"rotation_rules"`
	Tags                 map[string]string                          `json:"tags"`
	KmsKeyId             string                                     `json:"kms_key_id"`
	Description          string                                     `json:"description"`
	RecoveryWindowInDays int                                        `json:"recovery_window_in_days"`
	RotationEnabled      bool                                       `json:"rotation_enabled"`
	Arn                  string                                     `json:"arn"`
}

type AwsSecretsmanagerSecretSpecRotationRules struct {
	AutomaticallyAfterDays int `json:"automatically_after_days"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecretsmanagerSecretList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSecretsmanagerSecret `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigDeliveryChannel struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsConfigDeliveryChannelSpec `json"spec"`
}

type AwsConfigDeliveryChannelSpec struct {
	SnsTopicArn                string                                                   `json:"sns_topic_arn"`
	SnapshotDeliveryProperties []AwsConfigDeliveryChannelSpecSnapshotDeliveryProperties `json:"snapshot_delivery_properties"`
	Name                       string                                                   `json:"name"`
	S3BucketName               string                                                   `json:"s3_bucket_name"`
	S3KeyPrefix                string                                                   `json:"s3_key_prefix"`
}

type AwsConfigDeliveryChannelSpecSnapshotDeliveryProperties struct {
	DeliveryFrequency string `json:"delivery_frequency"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigDeliveryChannelList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsConfigDeliveryChannel `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheSecurityGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticacheSecurityGroupSpec `json"spec"`
}

type AwsElasticacheSecurityGroupSpec struct {
	Description        string `json:"description"`
	Name               string `json:"name"`
	SecurityGroupNames string `json:"security_group_names"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheSecurityGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticacheSecurityGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkApplicationVersion struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticBeanstalkApplicationVersionSpec `json"spec"`
}

type AwsElasticBeanstalkApplicationVersionSpec struct {
	Bucket      string `json:"bucket"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	ForceDelete bool   `json:"force_delete"`
	Application string `json:"application"`
	Description string `json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkApplicationVersionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticBeanstalkApplicationVersion `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbOptionGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDbOptionGroupSpec `json"spec"`
}

type AwsDbOptionGroupSpec struct {
	EngineName             string                     `json:"engine_name"`
	MajorEngineVersion     string                     `json:"major_engine_version"`
	OptionGroupDescription string                     `json:"option_group_description"`
	Option                 AwsDbOptionGroupSpecOption `json:"option"`
	Tags                   map[string]string          `json:"tags"`
	Arn                    string                     `json:"arn"`
	Name                   string                     `json:"name"`
	NamePrefix             string                     `json:"name_prefix"`
}

type AwsDbOptionGroupSpecOption struct {
	OptionSettings              AwsDbOptionGroupSpecOptionOptionSettings `json:"option_settings"`
	Port                        int                                      `json:"port"`
	DbSecurityGroupMemberships  string                                   `json:"db_security_group_memberships"`
	VpcSecurityGroupMemberships string                                   `json:"vpc_security_group_memberships"`
	Version                     string                                   `json:"version"`
	OptionName                  string                                   `json:"option_name"`
}

type AwsDbOptionGroupSpecOptionOptionSettings struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbOptionGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDbOptionGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxConnectionAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDxConnectionAssociationSpec `json"spec"`
}

type AwsDxConnectionAssociationSpec struct {
	ConnectionId string `json:"connection_id"`
	LagId        string `json:"lag_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxConnectionAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDxConnectionAssociation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrCluster struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEmrClusterSpec `json"spec"`
}

type AwsEmrClusterSpec struct {
	TerminationProtection       bool                                  `json:"termination_protection"`
	KeepJobFlowAliveWhenNoSteps bool                                  `json:"keep_job_flow_alive_when_no_steps"`
	Tags                        map[string]string                     `json:"tags"`
	Configurations              string                                `json:"configurations"`
	MasterPublicDns             string                                `json:"master_public_dns"`
	InstanceGroup               AwsEmrClusterSpecInstanceGroup        `json:"instance_group"`
	VisibleToAllUsers           bool                                  `json:"visible_to_all_users"`
	CustomAmiId                 string                                `json:"custom_ami_id"`
	MasterInstanceType          string                                `json:"master_instance_type"`
	Applications                string                                `json:"applications"`
	ServiceRole                 string                                `json:"service_role"`
	ScaleDownBehavior           string                                `json:"scale_down_behavior"`
	SecurityConfiguration       string                                `json:"security_configuration"`
	EbsRootVolumeSize           int                                   `json:"ebs_root_volume_size"`
	Name                        string                                `json:"name"`
	CoreInstanceType            string                                `json:"core_instance_type"`
	CoreInstanceCount           int                                   `json:"core_instance_count"`
	ClusterState                string                                `json:"cluster_state"`
	LogUri                      string                                `json:"log_uri"`
	Ec2Attributes               []AwsEmrClusterSpecEc2Attributes      `json:"ec2_attributes"`
	KerberosAttributes          []AwsEmrClusterSpecKerberosAttributes `json:"kerberos_attributes"`
	BootstrapAction             AwsEmrClusterSpecBootstrapAction      `json:"bootstrap_action"`
	ReleaseLabel                string                                `json:"release_label"`
	AutoscalingRole             string                                `json:"autoscaling_role"`
	Step                        []AwsEmrClusterSpecStep               `json:"step"`
}

type AwsEmrClusterSpecInstanceGroup struct {
	Name              string                                  `json:"name"`
	BidPrice          string                                  `json:"bid_price"`
	EbsConfig         AwsEmrClusterSpecInstanceGroupEbsConfig `json:"ebs_config"`
	InstanceCount     int                                     `json:"instance_count"`
	AutoscalingPolicy string                                  `json:"autoscaling_policy"`
	InstanceRole      string                                  `json:"instance_role"`
	InstanceType      string                                  `json:"instance_type"`
}

type AwsEmrClusterSpecInstanceGroupEbsConfig struct {
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int    `json:"volumes_per_instance"`
}

type AwsEmrClusterSpecEc2Attributes struct {
	AdditionalMasterSecurityGroups string `json:"additional_master_security_groups"`
	AdditionalSlaveSecurityGroups  string `json:"additional_slave_security_groups"`
	EmrManagedMasterSecurityGroup  string `json:"emr_managed_master_security_group"`
	EmrManagedSlaveSecurityGroup   string `json:"emr_managed_slave_security_group"`
	InstanceProfile                string `json:"instance_profile"`
	ServiceAccessSecurityGroup     string `json:"service_access_security_group"`
	KeyName                        string `json:"key_name"`
	SubnetId                       string `json:"subnet_id"`
}

type AwsEmrClusterSpecKerberosAttributes struct {
	CrossRealmTrustPrincipalPassword string `json:"cross_realm_trust_principal_password"`
	KdcAdminPassword                 string `json:"kdc_admin_password"`
	Realm                            string `json:"realm"`
	AdDomainJoinPassword             string `json:"ad_domain_join_password"`
	AdDomainJoinUser                 string `json:"ad_domain_join_user"`
}

type AwsEmrClusterSpecBootstrapAction struct {
	Args []string `json:"args"`
	Name string   `json:"name"`
	Path string   `json:"path"`
}

type AwsEmrClusterSpecStep struct {
	Name            string                               `json:"name"`
	ActionOnFailure string                               `json:"action_on_failure"`
	HadoopJarStep   []AwsEmrClusterSpecStepHadoopJarStep `json:"hadoop_jar_step"`
}

type AwsEmrClusterSpecStepHadoopJarStep struct {
	Args       []string          `json:"args"`
	Jar        string            `json:"jar"`
	MainClass  string            `json:"main_class"`
	Properties map[string]string `json:"properties"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrClusterList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEmrCluster `json"items"`
}

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
	CallbackUrls                    []string `json:"callback_urls"`
	SupportedIdentityProviders      []string `json:"supported_identity_providers"`
	ClientSecret                    string   `json:"client_secret"`
	RefreshTokenValidity            int      `json:"refresh_token_validity"`
	AllowedOauthFlows               string   `json:"allowed_oauth_flows"`
	DefaultRedirectUri              string   `json:"default_redirect_uri"`
	GenerateSecret                  bool     `json:"generate_secret"`
	ReadAttributes                  string   `json:"read_attributes"`
	WriteAttributes                 string   `json:"write_attributes"`
	LogoutUrls                      []string `json:"logout_urls"`
	Name                            string   `json:"name"`
	UserPoolId                      string   `json:"user_pool_id"`
	ExplicitAuthFlows               string   `json:"explicit_auth_flows"`
	AllowedOauthScopes              string   `json:"allowed_oauth_scopes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolClientList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCognitoUserPoolClient `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbSubnetGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDbSubnetGroupSpec `json"spec"`
}

type AwsDbSubnetGroupSpec struct {
	Description string            `json:"description"`
	SubnetIds   string            `json:"subnet_ids"`
	Tags        map[string]string `json:"tags"`
	Arn         string            `json:"arn"`
	Name        string            `json:"name"`
	NamePrefix  string            `json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbSubnetGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDbSubnetGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStaticWebLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksStaticWebLayerSpec `json"spec"`
}

type AwsOpsworksStaticWebLayerSpec struct {
	CustomConfigureRecipes   []string                               `json:"custom_configure_recipes"`
	AutoHealing              bool                                   `json:"auto_healing"`
	StackId                  string                                 `json:"stack_id"`
	CustomUndeployRecipes    []string                               `json:"custom_undeploy_recipes"`
	InstallUpdatesOnBoot     bool                                   `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  int                                    `json:"instance_shutdown_timeout"`
	CustomInstanceProfileArn string                                 `json:"custom_instance_profile_arn"`
	ElasticLoadBalancer      string                                 `json:"elastic_load_balancer"`
	CustomSetupRecipes       []string                               `json:"custom_setup_recipes"`
	DrainElbOnShutdown       bool                                   `json:"drain_elb_on_shutdown"`
	SystemPackages           string                                 `json:"system_packages"`
	Name                     string                                 `json:"name"`
	UseEbsOptimizedInstances bool                                   `json:"use_ebs_optimized_instances"`
	EbsVolume                AwsOpsworksStaticWebLayerSpecEbsVolume `json:"ebs_volume"`
	AutoAssignElasticIps     bool                                   `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                                   `json:"auto_assign_public_ips"`
	CustomDeployRecipes      []string                               `json:"custom_deploy_recipes"`
	CustomShutdownRecipes    []string                               `json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds   string                                 `json:"custom_security_group_ids"`
	CustomJson               string                                 `json:"custom_json"`
}

type AwsOpsworksStaticWebLayerSpecEbsVolume struct {
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Iops          int    `json:"iops"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStaticWebLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksStaticWebLayer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSqsQueuePolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSqsQueuePolicySpec `json"spec"`
}

type AwsSqsQueuePolicySpec struct {
	QueueUrl string `json:"queue_url"`
	Policy   string `json:"policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSqsQueuePolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSqsQueuePolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointRouteTableAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcEndpointRouteTableAssociationSpec `json"spec"`
}

type AwsVpcEndpointRouteTableAssociationSpec struct {
	VpcEndpointId string `json:"vpc_endpoint_id"`
	RouteTableId  string `json:"route_table_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointRouteTableAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcEndpointRouteTableAssociation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafSqlInjectionMatchSet struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafSqlInjectionMatchSetSpec `json"spec"`
}

type AwsWafSqlInjectionMatchSetSpec struct {
	Name                    string                                                `json:"name"`
	SqlInjectionMatchTuples AwsWafSqlInjectionMatchSetSpecSqlInjectionMatchTuples `json:"sql_injection_match_tuples"`
}

type AwsWafSqlInjectionMatchSetSpecSqlInjectionMatchTuples struct {
	FieldToMatch       AwsWafSqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch `json:"field_to_match"`
	TextTransformation string                                                            `json:"text_transformation"`
}

type AwsWafSqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafSqlInjectionMatchSetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafSqlInjectionMatchSet `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointConnectionNotification struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpcEndpointConnectionNotificationSpec `json"spec"`
}

type AwsVpcEndpointConnectionNotificationSpec struct {
	VpcEndpointServiceId      string `json:"vpc_endpoint_service_id"`
	VpcEndpointId             string `json:"vpc_endpoint_id"`
	ConnectionNotificationArn string `json:"connection_notification_arn"`
	ConnectionEvents          string `json:"connection_events"`
	State                     string `json:"state"`
	NotificationType          string `json:"notification_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointConnectionNotificationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpcEndpointConnectionNotification `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnConnection struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsVpnConnectionSpec `json"spec"`
}

type AwsVpnConnectionSpec struct {
	Tunnel2VgwInsideAddress      string                           `json:"tunnel2_vgw_inside_address"`
	Routes                       AwsVpnConnectionSpecRoutes       `json:"routes"`
	VgwTelemetry                 AwsVpnConnectionSpecVgwTelemetry `json:"vgw_telemetry"`
	CustomerGatewayId            string                           `json:"customer_gateway_id"`
	CustomerGatewayConfiguration string                           `json:"customer_gateway_configuration"`
	Tunnel1Address               string                           `json:"tunnel1_address"`
	Tunnel1BgpHoldtime           int                              `json:"tunnel1_bgp_holdtime"`
	Tunnel2Address               string                           `json:"tunnel2_address"`
	Tags                         map[string]string                `json:"tags"`
	Tunnel2CgwInsideAddress      string                           `json:"tunnel2_cgw_inside_address"`
	Tunnel1VgwInsideAddress      string                           `json:"tunnel1_vgw_inside_address"`
	Tunnel2BgpAsn                string                           `json:"tunnel2_bgp_asn"`
	VpnGatewayId                 string                           `json:"vpn_gateway_id"`
	Tunnel1PresharedKey          string                           `json:"tunnel1_preshared_key"`
	Tunnel2InsideCidr            string                           `json:"tunnel2_inside_cidr"`
	Tunnel2PresharedKey          string                           `json:"tunnel2_preshared_key"`
	Tunnel1CgwInsideAddress      string                           `json:"tunnel1_cgw_inside_address"`
	Type                         string                           `json:"type"`
	StaticRoutesOnly             bool                             `json:"static_routes_only"`
	Tunnel1InsideCidr            string                           `json:"tunnel1_inside_cidr"`
	Tunnel1BgpAsn                string                           `json:"tunnel1_bgp_asn"`
	Tunnel2BgpHoldtime           int                              `json:"tunnel2_bgp_holdtime"`
}

type AwsVpnConnectionSpecRoutes struct {
	DestinationCidrBlock string `json:"destination_cidr_block"`
	Source               string `json:"source"`
	State                string `json:"state"`
}

type AwsVpnConnectionSpecVgwTelemetry struct {
	LastStatusChange   string `json:"last_status_change"`
	OutsideIpAddress   string `json:"outside_ip_address"`
	Status             string `json:"status"`
	StatusMessage      string `json:"status_message"`
	AcceptedRouteCount int    `json:"accepted_route_count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnConnectionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsVpnConnection `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAthenaDatabase struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAthenaDatabaseSpec `json"spec"`
}

type AwsAthenaDatabaseSpec struct {
	Name         string `json:"name"`
	Bucket       string `json:"bucket"`
	ForceDestroy bool   `json:"force_destroy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAthenaDatabaseList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAthenaDatabase `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrSecurityConfiguration struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEmrSecurityConfigurationSpec `json"spec"`
}

type AwsEmrSecurityConfigurationSpec struct {
	Name          string `json:"name"`
	NamePrefix    string `json:"name_prefix"`
	Configuration string `json:"configuration"`
	CreationDate  string `json:"creation_date"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrSecurityConfigurationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEmrSecurityConfiguration `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyIpset struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGuarddutyIpsetSpec `json"spec"`
}

type AwsGuarddutyIpsetSpec struct {
	Activate   bool   `json:"activate"`
	DetectorId string `json:"detector_id"`
	Name       string `json:"name"`
	Format     string `json:"format"`
	Location   string `json:"location"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyIpsetList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGuarddutyIpset `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamServerCertificate struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIamServerCertificateSpec `json"spec"`
}

type AwsIamServerCertificateSpec struct {
	NamePrefix       string `json:"name_prefix"`
	Arn              string `json:"arn"`
	CertificateBody  string `json:"certificate_body"`
	CertificateChain string `json:"certificate_chain"`
	Path             string `json:"path"`
	PrivateKey       string `json:"private_key"`
	Name             string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamServerCertificateList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIamServerCertificate `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotThing struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIotThingSpec `json"spec"`
}

type AwsIotThingSpec struct {
	DefaultClientId string            `json:"default_client_id"`
	Version         int               `json:"version"`
	Arn             string            `json:"arn"`
	Name            string            `json:"name"`
	Attributes      map[string]string `json:"attributes"`
	ThingTypeName   string            `json:"thing_type_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotThingList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIotThing `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbSslNegotiationPolicy struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLbSslNegotiationPolicySpec `json"spec"`
}

type AwsLbSslNegotiationPolicySpec struct {
	LbPort       int                                    `json:"lb_port"`
	Attribute    AwsLbSslNegotiationPolicySpecAttribute `json:"attribute"`
	Name         string                                 `json:"name"`
	LoadBalancer string                                 `json:"load_balancer"`
}

type AwsLbSslNegotiationPolicySpecAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbSslNegotiationPolicyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLbSslNegotiationPolicy `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMqConfiguration struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsMqConfigurationSpec `json"spec"`
}

type AwsMqConfigurationSpec struct {
	LatestRevision int    `json:"latest_revision"`
	Arn            string `json:"arn"`
	Data           string `json:"data"`
	Description    string `json:"description"`
	EngineType     string `json:"engine_type"`
	EngineVersion  string `json:"engine_version"`
	Name           string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMqConfigurationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsMqConfiguration `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainIdentity struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesDomainIdentitySpec `json"spec"`
}

type AwsSesDomainIdentitySpec struct {
	Arn               string `json:"arn"`
	Domain            string `json:"domain"`
	VerificationToken string `json:"verification_token"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainIdentityList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesDomainIdentity `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmi struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAmiSpec `json"spec"`
}

type AwsAmiSpec struct {
	ManageEbsSnapshots   bool                           `json:"manage_ebs_snapshots"`
	SriovNetSupport      string                         `json:"sriov_net_support"`
	VirtualizationType   string                         `json:"virtualization_type"`
	RootDeviceName       string                         `json:"root_device_name"`
	RootSnapshotId       string                         `json:"root_snapshot_id"`
	Tags                 map[string]string              `json:"tags"`
	Description          string                         `json:"description"`
	Name                 string                         `json:"name"`
	KernelId             string                         `json:"kernel_id"`
	RamdiskId            string                         `json:"ramdisk_id"`
	EbsBlockDevice       AwsAmiSpecEbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice AwsAmiSpecEphemeralBlockDevice `json:"ephemeral_block_device"`
	ImageLocation        string                         `json:"image_location"`
	Architecture         string                         `json:"architecture"`
}

type AwsAmiSpecEbsBlockDevice struct {
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
}

type AwsAmiSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAmi `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayGatewayResponse struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayGatewayResponseSpec `json"spec"`
}

type AwsApiGatewayGatewayResponseSpec struct {
	RestApiId          string            `json:"rest_api_id"`
	ResponseType       string            `json:"response_type"`
	StatusCode         string            `json:"status_code"`
	ResponseTemplates  map[string]string `json:"response_templates"`
	ResponseParameters map[string]string `json:"response_parameters"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayGatewayResponseList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayGatewayResponse `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAutoscalingGroupSpec `json"spec"`
}

type AwsAutoscalingGroupSpec struct {
	LoadBalancers          string                                      `json:"load_balancers"`
	EnabledMetrics         string                                      `json:"enabled_metrics"`
	Tags                   []map[string]string                         `json:"tags"`
	LaunchTemplate         []AwsAutoscalingGroupSpecLaunchTemplate     `json:"launch_template"`
	AvailabilityZones      string                                      `json:"availability_zones"`
	WaitForCapacityTimeout string                                      `json:"wait_for_capacity_timeout"`
	MetricsGranularity     string                                      `json:"metrics_granularity"`
	Arn                    string                                      `json:"arn"`
	HealthCheckType        string                                      `json:"health_check_type"`
	TerminationPolicies    []string                                    `json:"termination_policies"`
	WaitForElbCapacity     int                                         `json:"wait_for_elb_capacity"`
	TargetGroupArns        string                                      `json:"target_group_arns"`
	ForceDelete            bool                                        `json:"force_delete"`
	SuspendedProcesses     string                                      `json:"suspended_processes"`
	ProtectFromScaleIn     bool                                        `json:"protect_from_scale_in"`
	DefaultCooldown        int                                         `json:"default_cooldown"`
	PlacementGroup         string                                      `json:"placement_group"`
	ServiceLinkedRoleArn   string                                      `json:"service_linked_role_arn"`
	LaunchConfiguration    string                                      `json:"launch_configuration"`
	DesiredCapacity        int                                         `json:"desired_capacity"`
	MinSize                int                                         `json:"min_size"`
	MinElbCapacity         int                                         `json:"min_elb_capacity"`
	MaxSize                int                                         `json:"max_size"`
	InitialLifecycleHook   AwsAutoscalingGroupSpecInitialLifecycleHook `json:"initial_lifecycle_hook"`
	NamePrefix             string                                      `json:"name_prefix"`
	Tag                    AwsAutoscalingGroupSpecTag                  `json:"tag"`
	Name                   string                                      `json:"name"`
	HealthCheckGracePeriod int                                         `json:"health_check_grace_period"`
	VpcZoneIdentifier      string                                      `json:"vpc_zone_identifier"`
}

type AwsAutoscalingGroupSpecLaunchTemplate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type AwsAutoscalingGroupSpecInitialLifecycleHook struct {
	RoleArn               string `json:"role_arn"`
	Name                  string `json:"name"`
	DefaultResult         string `json:"default_result"`
	HeartbeatTimeout      int    `json:"heartbeat_timeout"`
	LifecycleTransition   string `json:"lifecycle_transition"`
	NotificationMetadata  string `json:"notification_metadata"`
	NotificationTargetArn string `json:"notification_target_arn"`
}

type AwsAutoscalingGroupSpecTag struct {
	Key               string `json:"key"`
	Value             string `json:"value"`
	PropagateAtLaunch bool   `json:"propagate_at_launch"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAutoscalingGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodebuildProject struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCodebuildProjectSpec `json"spec"`
}

type AwsCodebuildProjectSpec struct {
	Cache         []AwsCodebuildProjectSpecCache     `json:"cache"`
	Description   string                             `json:"description"`
	EncryptionKey string                             `json:"encryption_key"`
	ServiceRole   string                             `json:"service_role"`
	Timeout       int                                `json:"timeout"`
	VpcConfig     []AwsCodebuildProjectSpecVpcConfig `json:"vpc_config"`
	Artifacts     AwsCodebuildProjectSpecArtifacts   `json:"artifacts"`
	Environment   AwsCodebuildProjectSpecEnvironment `json:"environment"`
	Name          string                             `json:"name"`
	Source        AwsCodebuildProjectSpecSource      `json:"source"`
	BuildTimeout  int                                `json:"build_timeout"`
	Tags          map[string]string                  `json:"tags"`
}

type AwsCodebuildProjectSpecCache struct {
	Location string `json:"location"`
	Type     string `json:"type"`
}

type AwsCodebuildProjectSpecVpcConfig struct {
	VpcId            string `json:"vpc_id"`
	Subnets          string `json:"subnets"`
	SecurityGroupIds string `json:"security_group_ids"`
}

type AwsCodebuildProjectSpecArtifacts struct {
	Name          string `json:"name"`
	Location      string `json:"location"`
	NamespaceType string `json:"namespace_type"`
	Packaging     string `json:"packaging"`
	Path          string `json:"path"`
	Type          string `json:"type"`
}

type AwsCodebuildProjectSpecEnvironment struct {
	EnvironmentVariable []AwsCodebuildProjectSpecEnvironmentEnvironmentVariable `json:"environment_variable"`
	Image               string                                                  `json:"image"`
	Type                string                                                  `json:"type"`
	PrivilegedMode      bool                                                    `json:"privileged_mode"`
	ComputeType         string                                                  `json:"compute_type"`
}

type AwsCodebuildProjectSpecEnvironmentEnvironmentVariable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsCodebuildProjectSpecSource struct {
	Auth      AwsCodebuildProjectSpecSourceAuth `json:"auth"`
	Buildspec string                            `json:"buildspec"`
	Location  string                            `json:"location"`
	Type      string                            `json:"type"`
}

type AwsCodebuildProjectSpecSourceAuth struct {
	Resource string `json:"resource"`
	Type     string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodebuildProjectList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCodebuildProject `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEipAssociation struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEipAssociationSpec `json"spec"`
}

type AwsEipAssociationSpec struct {
	AllocationId       string `json:"allocation_id"`
	AllowReassociation bool   `json:"allow_reassociation"`
	InstanceId         string `json:"instance_id"`
	NetworkInterfaceId string `json:"network_interface_id"`
	PrivateIpAddress   string `json:"private_ip_address"`
	PublicIp           string `json:"public_ip"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEipAssociationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsEipAssociation `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheParameterGroup struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticacheParameterGroupSpec `json"spec"`
}

type AwsElasticacheParameterGroupSpec struct {
	Family      string                                    `json:"family"`
	Description string                                    `json:"description"`
	Parameter   AwsElasticacheParameterGroupSpecParameter `json:"parameter"`
	Name        string                                    `json:"name"`
}

type AwsElasticacheParameterGroupSpecParameter struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheParameterGroupList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticacheParameterGroup `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafWebAcl struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsWafWebAclSpec `json"spec"`
}

type AwsWafWebAclSpec struct {
	Name          string                        `json:"name"`
	DefaultAction AwsWafWebAclSpecDefaultAction `json:"default_action"`
	MetricName    string                        `json:"metric_name"`
	Rules         AwsWafWebAclSpecRules         `json:"rules"`
}

type AwsWafWebAclSpecDefaultAction struct {
	Type string `json:"type"`
}

type AwsWafWebAclSpecRules struct {
	RuleId   string                      `json:"rule_id"`
	Action   AwsWafWebAclSpecRulesAction `json:"action"`
	Priority int                         `json:"priority"`
	Type     string                      `json:"type"`
}

type AwsWafWebAclSpecRulesAction struct {
	Type string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafWebAclList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsWafWebAcl `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodedeployDeploymentConfig struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCodedeployDeploymentConfigSpec `json"spec"`
}

type AwsCodedeployDeploymentConfigSpec struct {
	DeploymentConfigName string                                                 `json:"deployment_config_name"`
	MinimumHealthyHosts  []AwsCodedeployDeploymentConfigSpecMinimumHealthyHosts `json:"minimum_healthy_hosts"`
	DeploymentConfigId   string                                                 `json:"deployment_config_id"`
}

type AwsCodedeployDeploymentConfigSpecMinimumHealthyHosts struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodedeployDeploymentConfigList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCodedeployDeploymentConfig `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsFlowLog struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsFlowLogSpec `json"spec"`
}

type AwsFlowLogSpec struct {
	IamRoleArn   string `json:"iam_role_arn"`
	LogGroupName string `json:"log_group_name"`
	VpcId        string `json:"vpc_id"`
	SubnetId     string `json:"subnet_id"`
	EniId        string `json:"eni_id"`
	TrafficType  string `json:"traffic_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsFlowLogList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsFlowLog `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsAlias struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsKmsAliasSpec `json"spec"`
}

type AwsKmsAliasSpec struct {
	NamePrefix   string `json:"name_prefix"`
	TargetKeyId  string `json:"target_key_id"`
	TargetKeyArn string `json:"target_key_arn"`
	Arn          string `json:"arn"`
	Name         string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsAliasList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsKmsAlias `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesIdentityNotificationTopic struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSesIdentityNotificationTopicSpec `json"spec"`
}

type AwsSesIdentityNotificationTopicSpec struct {
	Identity         string `json:"identity"`
	TopicArn         string `json:"topic_arn"`
	NotificationType string `json:"notification_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesIdentityNotificationTopicList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSesIdentityNotificationTopic `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbSnapshot struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsDbSnapshotSpec `json"spec"`
}

type AwsDbSnapshotSpec struct {
	Engine                     string `json:"engine"`
	SourceDbSnapshotIdentifier string `json:"source_db_snapshot_identifier"`
	VpcId                      string `json:"vpc_id"`
	DbSnapshotArn              string `json:"db_snapshot_arn"`
	AvailabilityZone           string `json:"availability_zone"`
	Encrypted                  bool   `json:"encrypted"`
	EngineVersion              string `json:"engine_version"`
	KmsKeyId                   string `json:"kms_key_id"`
	OptionGroupName            string `json:"option_group_name"`
	DbSnapshotIdentifier       string `json:"db_snapshot_identifier"`
	AllocatedStorage           int    `json:"allocated_storage"`
	Port                       int    `json:"port"`
	SnapshotType               string `json:"snapshot_type"`
	Status                     string `json:"status"`
	StorageType                string `json:"storage_type"`
	DbInstanceIdentifier       string `json:"db_instance_identifier"`
	LicenseModel               string `json:"license_model"`
	SourceRegion               string `json:"source_region"`
	Iops                       int    `json:"iops"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbSnapshotList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsDbSnapshot `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkApplication struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElasticBeanstalkApplicationSpec `json"spec"`
}

type AwsElasticBeanstalkApplicationSpec struct {
	Name                string                                                  `json:"name"`
	Description         string                                                  `json:"description"`
	AppversionLifecycle []AwsElasticBeanstalkApplicationSpecAppversionLifecycle `json:"appversion_lifecycle"`
}

type AwsElasticBeanstalkApplicationSpecAppversionLifecycle struct {
	DeleteSourceFromS3 bool   `json:"delete_source_from_s3"`
	ServiceRole        string `json:"service_role"`
	MaxAgeInDays       int    `json:"max_age_in_days"`
	MaxCount           int    `json:"max_count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkApplicationList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElasticBeanstalkApplication `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotTopicRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIotTopicRuleSpec `json"spec"`
}

type AwsIotTopicRuleSpec struct {
	Firehose         AwsIotTopicRuleSpecFirehose         `json:"firehose"`
	Kinesis          AwsIotTopicRuleSpecKinesis          `json:"kinesis"`
	Lambda           AwsIotTopicRuleSpecLambda           `json:"lambda"`
	Name             string                              `json:"name"`
	Enabled          bool                                `json:"enabled"`
	CloudwatchAlarm  AwsIotTopicRuleSpecCloudwatchAlarm  `json:"cloudwatch_alarm"`
	S3               AwsIotTopicRuleSpecS3               `json:"s3"`
	Description      string                              `json:"description"`
	Elasticsearch    AwsIotTopicRuleSpecElasticsearch    `json:"elasticsearch"`
	Republish        AwsIotTopicRuleSpecRepublish        `json:"republish"`
	Sns              AwsIotTopicRuleSpecSns              `json:"sns"`
	Sqs              AwsIotTopicRuleSpecSqs              `json:"sqs"`
	Sql              string                              `json:"sql"`
	SqlVersion       string                              `json:"sql_version"`
	CloudwatchMetric AwsIotTopicRuleSpecCloudwatchMetric `json:"cloudwatch_metric"`
	Dynamodb         AwsIotTopicRuleSpecDynamodb         `json:"dynamodb"`
	Arn              string                              `json:"arn"`
}

type AwsIotTopicRuleSpecFirehose struct {
	RoleArn            string `json:"role_arn"`
	DeliveryStreamName string `json:"delivery_stream_name"`
}

type AwsIotTopicRuleSpecKinesis struct {
	PartitionKey string `json:"partition_key"`
	RoleArn      string `json:"role_arn"`
	StreamName   string `json:"stream_name"`
}

type AwsIotTopicRuleSpecLambda struct {
	FunctionArn string `json:"function_arn"`
}

type AwsIotTopicRuleSpecCloudwatchAlarm struct {
	AlarmName   string `json:"alarm_name"`
	RoleArn     string `json:"role_arn"`
	StateReason string `json:"state_reason"`
	StateValue  string `json:"state_value"`
}

type AwsIotTopicRuleSpecS3 struct {
	RoleArn    string `json:"role_arn"`
	BucketName string `json:"bucket_name"`
	Key        string `json:"key"`
}

type AwsIotTopicRuleSpecElasticsearch struct {
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
	Index    string `json:"index"`
	RoleArn  string `json:"role_arn"`
	Type     string `json:"type"`
}

type AwsIotTopicRuleSpecRepublish struct {
	RoleArn string `json:"role_arn"`
	Topic   string `json:"topic"`
}

type AwsIotTopicRuleSpecSns struct {
	MessageFormat string `json:"message_format"`
	TargetArn     string `json:"target_arn"`
	RoleArn       string `json:"role_arn"`
}

type AwsIotTopicRuleSpecSqs struct {
	QueueUrl  string `json:"queue_url"`
	RoleArn   string `json:"role_arn"`
	UseBase64 bool   `json:"use_base64"`
}

type AwsIotTopicRuleSpecCloudwatchMetric struct {
	MetricName      string `json:"metric_name"`
	MetricNamespace string `json:"metric_namespace"`
	MetricTimestamp string `json:"metric_timestamp"`
	MetricUnit      string `json:"metric_unit"`
	MetricValue     string `json:"metric_value"`
	RoleArn         string `json:"role_arn"`
}

type AwsIotTopicRuleSpecDynamodb struct {
	TableName     string `json:"table_name"`
	HashKeyField  string `json:"hash_key_field"`
	HashKeyValue  string `json:"hash_key_value"`
	HashKeyType   string `json:"hash_key_type"`
	PayloadField  string `json:"payload_field"`
	RangeKeyValue string `json:"range_key_value"`
	RangeKeyType  string `json:"range_key_type"`
	RoleArn       string `json:"role_arn"`
	RangeKeyField string `json:"range_key_field"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotTopicRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIotTopicRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroupRule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSecurityGroupRuleSpec `json"spec"`
}

type AwsSecurityGroupRuleSpec struct {
	Protocol              string   `json:"protocol"`
	Ipv6CidrBlocks        []string `json:"ipv6_cidr_blocks"`
	PrefixListIds         []string `json:"prefix_list_ids"`
	SecurityGroupId       string   `json:"security_group_id"`
	Self                  bool     `json:"self"`
	Type                  string   `json:"type"`
	FromPort              int      `json:"from_port"`
	SourceSecurityGroupId string   `json:"source_security_group_id"`
	Description           string   `json:"description"`
	ToPort                int      `json:"to_port"`
	CidrBlocks            []string `json:"cidr_blocks"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroupRuleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSecurityGroupRule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksPhpAppLayer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsOpsworksPhpAppLayerSpec `json"spec"`
}

type AwsOpsworksPhpAppLayerSpec struct {
	StackId                  string                              `json:"stack_id"`
	AutoAssignPublicIps      bool                                `json:"auto_assign_public_ips"`
	CustomDeployRecipes      []string                            `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string                            `json:"custom_undeploy_recipes"`
	AutoHealing              bool                                `json:"auto_healing"`
	EbsVolume                AwsOpsworksPhpAppLayerSpecEbsVolume `json:"ebs_volume"`
	CustomInstanceProfileArn string                              `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string                            `json:"custom_setup_recipes"`
	CustomSecurityGroupIds   string                              `json:"custom_security_group_ids"`
	SystemPackages           string                              `json:"system_packages"`
	Name                     string                              `json:"name"`
	AutoAssignElasticIps     bool                                `json:"auto_assign_elastic_ips"`
	CustomShutdownRecipes    []string                            `json:"custom_shutdown_recipes"`
	CustomJson               string                              `json:"custom_json"`
	InstallUpdatesOnBoot     bool                                `json:"install_updates_on_boot"`
	UseEbsOptimizedInstances bool                                `json:"use_ebs_optimized_instances"`
	ElasticLoadBalancer      string                              `json:"elastic_load_balancer"`
	CustomConfigureRecipes   []string                            `json:"custom_configure_recipes"`
	InstanceShutdownTimeout  int                                 `json:"instance_shutdown_timeout"`
	DrainElbOnShutdown       bool                                `json:"drain_elb_on_shutdown"`
}

type AwsOpsworksPhpAppLayerSpecEbsVolume struct {
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksPhpAppLayerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsOpsworksPhpAppLayer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmPatchBaseline struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsSsmPatchBaselineSpec `json"spec"`
}

type AwsSsmPatchBaselineSpec struct {
	OperatingSystem                string                                `json:"operating_system"`
	ApprovedPatchesComplianceLevel string                                `json:"approved_patches_compliance_level"`
	Name                           string                                `json:"name"`
	Description                    string                                `json:"description"`
	GlobalFilter                   []AwsSsmPatchBaselineSpecGlobalFilter `json:"global_filter"`
	ApprovalRule                   []AwsSsmPatchBaselineSpecApprovalRule `json:"approval_rule"`
	ApprovedPatches                string                                `json:"approved_patches"`
	RejectedPatches                string                                `json:"rejected_patches"`
}

type AwsSsmPatchBaselineSpecGlobalFilter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type AwsSsmPatchBaselineSpecApprovalRule struct {
	ApproveAfterDays int                                              `json:"approve_after_days"`
	ComplianceLevel  string                                           `json:"compliance_level"`
	PatchFilter      []AwsSsmPatchBaselineSpecApprovalRulePatchFilter `json:"patch_filter"`
}

type AwsSsmPatchBaselineSpecApprovalRulePatchFilter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmPatchBaselineList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsSsmPatchBaseline `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayRequestValidator struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsApiGatewayRequestValidatorSpec `json"spec"`
}

type AwsApiGatewayRequestValidatorSpec struct {
	ValidateRequestParameters bool   `json:"validate_request_parameters"`
	RestApiId                 string `json:"rest_api_id"`
	Name                      string `json:"name"`
	ValidateRequestBody       bool   `json:"validate_request_body"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayRequestValidatorList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsApiGatewayRequestValidator `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingSchedule struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAutoscalingScheduleSpec `json"spec"`
}

type AwsAutoscalingScheduleSpec struct {
	EndTime              string `json:"end_time"`
	MaxSize              int    `json:"max_size"`
	ScheduledActionName  string `json:"scheduled_action_name"`
	StartTime            string `json:"start_time"`
	Recurrence           string `json:"recurrence"`
	MinSize              int    `json:"min_size"`
	DesiredCapacity      int    `json:"desired_capacity"`
	Arn                  string `json:"arn"`
	AutoscalingGroupName string `json:"autoscaling_group_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingScheduleList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAutoscalingSchedule `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueCatalogDatabase struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsGlueCatalogDatabaseSpec `json"spec"`
}

type AwsGlueCatalogDatabaseSpec struct {
	CatalogId   string            `json:"catalog_id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	LocationUri string            `json:"location_uri"`
	Parameters  map[string]string `json:"parameters"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueCatalogDatabaseList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsGlueCatalogDatabase `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsKey struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsKmsKeySpec `json"spec"`
}

type AwsKmsKeySpec struct {
	KeyId                string            `json:"key_id"`
	Policy               string            `json:"policy"`
	EnableKeyRotation    bool              `json:"enable_key_rotation"`
	DeletionWindowInDays int               `json:"deletion_window_in_days"`
	Tags                 map[string]string `json:"tags"`
	Arn                  string            `json:"arn"`
	Description          string            `json:"description"`
	KeyUsage             string            `json:"key_usage"`
	IsEnabled            bool              `json:"is_enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsKeyList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsKmsKey `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaFunction struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLambdaFunctionSpec `json"spec"`
}

type AwsLambdaFunctionSpec struct {
	Timeout                      int                                     `json:"timeout"`
	VpcConfig                    []AwsLambdaFunctionSpecVpcConfig        `json:"vpc_config"`
	Environment                  []AwsLambdaFunctionSpecEnvironment      `json:"environment"`
	S3Bucket                     string                                  `json:"s3_bucket"`
	S3ObjectVersion              string                                  `json:"s3_object_version"`
	FunctionName                 string                                  `json:"function_name"`
	Runtime                      string                                  `json:"runtime"`
	Description                  string                                  `json:"description"`
	ReservedConcurrentExecutions int                                     `json:"reserved_concurrent_executions"`
	QualifiedArn                 string                                  `json:"qualified_arn"`
	Version                      string                                  `json:"version"`
	Arn                          string                                  `json:"arn"`
	InvokeArn                    string                                  `json:"invoke_arn"`
	SourceCodeSize               int                                     `json:"source_code_size"`
	Filename                     string                                  `json:"filename"`
	S3Key                        string                                  `json:"s3_key"`
	DeadLetterConfig             []AwsLambdaFunctionSpecDeadLetterConfig `json:"dead_letter_config"`
	Publish                      bool                                    `json:"publish"`
	KmsKeyArn                    string                                  `json:"kms_key_arn"`
	Tags                         map[string]string                       `json:"tags"`
	SourceCodeHash               string                                  `json:"source_code_hash"`
	TracingConfig                []AwsLambdaFunctionSpecTracingConfig    `json:"tracing_config"`
	Handler                      string                                  `json:"handler"`
	MemorySize                   int                                     `json:"memory_size"`
	Role                         string                                  `json:"role"`
	LastModified                 string                                  `json:"last_modified"`
}

type AwsLambdaFunctionSpecVpcConfig struct {
	SubnetIds        string `json:"subnet_ids"`
	SecurityGroupIds string `json:"security_group_ids"`
	VpcId            string `json:"vpc_id"`
}

type AwsLambdaFunctionSpecEnvironment struct {
	Variables map[string]string `json:"variables"`
}

type AwsLambdaFunctionSpecDeadLetterConfig struct {
	TargetArn string `json:"target_arn"`
}

type AwsLambdaFunctionSpecTracingConfig struct {
	Mode string `json:"mode"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaFunctionList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLambdaFunction `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMediaStoreContainer struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsMediaStoreContainerSpec `json"spec"`
}

type AwsMediaStoreContainerSpec struct {
	Name     string `json:"name"`
	Arn      string `json:"arn"`
	Endpoint string `json:"endpoint"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMediaStoreContainerList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsMediaStoreContainer `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroupAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsAlbTargetGroupAttachmentSpec `json"spec"`
}

type AwsAlbTargetGroupAttachmentSpec struct {
	Port             int    `json:"port"`
	AvailabilityZone string `json:"availability_zone"`
	TargetGroupArn   string `json:"target_group_arn"`
	TargetId         string `json:"target_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroupAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsAlbTargetGroupAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoIdentityPoolRolesAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsCognitoIdentityPoolRolesAttachmentSpec `json"spec"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpec struct {
	Roles          map[string]AwsCognitoIdentityPoolRolesAttachmentSpecRoles `json:"roles"`
	IdentityPoolId string                                                    `json:"identity_pool_id"`
	RoleMapping    AwsCognitoIdentityPoolRolesAttachmentSpecRoleMapping      `json:"role_mapping"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpecRoles struct {
	Authenticated   string `json:"authenticated"`
	Unauthenticated string `json:"unauthenticated"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpecRoleMapping struct {
	IdentityProvider        string                                                            `json:"identity_provider"`
	AmbiguousRoleResolution string                                                            `json:"ambiguous_role_resolution"`
	MappingRule             []AwsCognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule `json:"mapping_rule"`
	Type                    string                                                            `json:"type"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule struct {
	Value     string `json:"value"`
	Claim     string `json:"claim"`
	MatchType string `json:"match_type"`
	RoleArn   string `json:"role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoIdentityPoolRolesAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsCognitoIdentityPoolRolesAttachment `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElb struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsElbSpec `json"spec"`
}

type AwsElbSpec struct {
	Internal                  bool                    `json:"internal"`
	SourceSecurityGroupId     string                  `json:"source_security_group_id"`
	AccessLogs                []AwsElbSpecAccessLogs  `json:"access_logs"`
	NamePrefix                string                  `json:"name_prefix"`
	Arn                       string                  `json:"arn"`
	SourceSecurityGroup       string                  `json:"source_security_group"`
	ConnectionDrainingTimeout int                     `json:"connection_draining_timeout"`
	HealthCheck               []AwsElbSpecHealthCheck `json:"health_check"`
	Tags                      map[string]string       `json:"tags"`
	Name                      string                  `json:"name"`
	CrossZoneLoadBalancing    bool                    `json:"cross_zone_load_balancing"`
	AvailabilityZones         string                  `json:"availability_zones"`
	Subnets                   string                  `json:"subnets"`
	ConnectionDraining        bool                    `json:"connection_draining"`
	Instances                 string                  `json:"instances"`
	SecurityGroups            string                  `json:"security_groups"`
	IdleTimeout               int                     `json:"idle_timeout"`
	Listener                  AwsElbSpecListener      `json:"listener"`
	DnsName                   string                  `json:"dns_name"`
	ZoneId                    string                  `json:"zone_id"`
}

type AwsElbSpecAccessLogs struct {
	Enabled      bool   `json:"enabled"`
	Interval     int    `json:"interval"`
	Bucket       string `json:"bucket"`
	BucketPrefix string `json:"bucket_prefix"`
}

type AwsElbSpecHealthCheck struct {
	Interval           int    `json:"interval"`
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Target             string `json:"target"`
}

type AwsElbSpecListener struct {
	LbProtocol       string `json:"lb_protocol"`
	SslCertificateId string `json:"ssl_certificate_id"`
	InstancePort     int    `json:"instance_port"`
	InstanceProtocol string `json:"instance_protocol"`
	LbPort           int    `json:"lb_port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElbList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsElb `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotThingType struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsIotThingTypeSpec `json"spec"`
}

type AwsIotThingTypeSpec struct {
	Deprecated bool                            `json:"deprecated"`
	Arn        string                          `json:"arn"`
	Name       string                          `json:"name"`
	Properties []AwsIotThingTypeSpecProperties `json:"properties"`
}

type AwsIotThingTypeSpecProperties struct {
	Description          string `json:"description"`
	SearchableAttributes string `json:"searchable_attributes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotThingTypeList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsIotThingType `json"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailStaticIpAttachment struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsLightsailStaticIpAttachmentSpec `json"spec"`
}

type AwsLightsailStaticIpAttachmentSpec struct {
	StaticIpName string `json:"static_ip_name"`
	InstanceName string `json:"instance_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailStaticIpAttachmentList struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Items              []AwsLightsailStaticIpAttachment `json"items"`
}
