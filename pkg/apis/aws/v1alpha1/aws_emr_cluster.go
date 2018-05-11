// GENERATED BY TF-GENERATOR, DO NOT EDIT.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrCluster struct {
	meta_v1.TypeMeta   `json",inline"`
	meta_v1.ObjectMeta `json"metadata,omitempty"`
	Spec               AwsEmrClusterSpec `json"spec"`
}

type AwsEmrClusterSpec struct {
	Ec2Attributes               []AwsEmrClusterSpecEc2Attributes      `json:"ec2_attributes"`
	CoreInstanceType            string                                `json:"core_instance_type"`
	CoreInstanceCount           int                                   `json:"core_instance_count"`
	ClusterState                string                                `json:"cluster_state"`
	LogUri                      string                                `json:"log_uri"`
	KeepJobFlowAliveWhenNoSteps bool                                  `json:"keep_job_flow_alive_when_no_steps"`
	Applications                string                                `json:"applications"`
	TerminationProtection       bool                                  `json:"termination_protection"`
	Tags                        map[string]string                     `json:"tags"`
	Configurations              string                                `json:"configurations"`
	ScaleDownBehavior           string                                `json:"scale_down_behavior"`
	VisibleToAllUsers           bool                                  `json:"visible_to_all_users"`
	EbsRootVolumeSize           int                                   `json:"ebs_root_volume_size"`
	Name                        string                                `json:"name"`
	MasterPublicDns             string                                `json:"master_public_dns"`
	KerberosAttributes          []AwsEmrClusterSpecKerberosAttributes `json:"kerberos_attributes"`
	BootstrapAction             AwsEmrClusterSpecBootstrapAction      `json:"bootstrap_action"`
	ServiceRole                 string                                `json:"service_role"`
	AutoscalingRole             string                                `json:"autoscaling_role"`
	CustomAmiId                 string                                `json:"custom_ami_id"`
	ReleaseLabel                string                                `json:"release_label"`
	MasterInstanceType          string                                `json:"master_instance_type"`
	InstanceGroup               AwsEmrClusterSpecInstanceGroup        `json:"instance_group"`
	Step                        []AwsEmrClusterSpecStep               `json:"step"`
	SecurityConfiguration       string                                `json:"security_configuration"`
}

type AwsEmrClusterSpecEc2Attributes struct {
	EmrManagedMasterSecurityGroup  string `json:"emr_managed_master_security_group"`
	EmrManagedSlaveSecurityGroup   string `json:"emr_managed_slave_security_group"`
	InstanceProfile                string `json:"instance_profile"`
	ServiceAccessSecurityGroup     string `json:"service_access_security_group"`
	KeyName                        string `json:"key_name"`
	SubnetId                       string `json:"subnet_id"`
	AdditionalMasterSecurityGroups string `json:"additional_master_security_groups"`
	AdditionalSlaveSecurityGroups  string `json:"additional_slave_security_groups"`
}

type AwsEmrClusterSpecKerberosAttributes struct {
	KdcAdminPassword                 string `json:"kdc_admin_password"`
	Realm                            string `json:"realm"`
	AdDomainJoinPassword             string `json:"ad_domain_join_password"`
	AdDomainJoinUser                 string `json:"ad_domain_join_user"`
	CrossRealmTrustPrincipalPassword string `json:"cross_realm_trust_principal_password"`
}

type AwsEmrClusterSpecBootstrapAction struct {
	Name string   `json:"name"`
	Path string   `json:"path"`
	Args []string `json:"args"`
}

type AwsEmrClusterSpecInstanceGroup struct {
	InstanceCount     int                                     `json:"instance_count"`
	AutoscalingPolicy string                                  `json:"autoscaling_policy"`
	InstanceRole      string                                  `json:"instance_role"`
	InstanceType      string                                  `json:"instance_type"`
	Name              string                                  `json:"name"`
	BidPrice          string                                  `json:"bid_price"`
	EbsConfig         AwsEmrClusterSpecInstanceGroupEbsConfig `json:"ebs_config"`
}

type AwsEmrClusterSpecInstanceGroupEbsConfig struct {
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int    `json:"volumes_per_instance"`
}

type AwsEmrClusterSpecStep struct {
	ActionOnFailure string                               `json:"action_on_failure"`
	HadoopJarStep   []AwsEmrClusterSpecStepHadoopJarStep `json:"hadoop_jar_step"`
	Name            string                               `json:"name"`
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
