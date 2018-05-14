package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	cs "github.com/chronojam/terraform-operator/pkg/client/clientset/versioned"
	operator "github.com/chronojam/terraform-operator/pkg/operator"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsacmcertificate"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsacmcertificatevalidation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsacmpcacertificateauthority"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsalb"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsalbtargetgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsalbtargetgroupattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsami"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsamicopy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsamifrominstance"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsamilaunchpermission"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayaccount"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayapikey"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayauthorizer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaybasepathmapping"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayclientcertificate"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaydeployment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaydocumentationpart"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaydocumentationversion"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaydomainname"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaygatewayresponse"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayintegration"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayintegrationresponse"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaymethod"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaymethodresponse"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaymethodsettings"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaymodel"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayrequestvalidator"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayresource"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayrestapi"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewaystage"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayusageplan"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayusageplankey"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsapigatewayvpclink"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsappautoscalingpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsappautoscalingscheduledaction"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsappautoscalingtarget"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsappcookiestickinesspolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsappsyncdatasource"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsappsyncgraphqlapi"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsathenadatabase"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsathenanamedquery"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsautoscalingattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsautoscalinggroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsautoscalinglifecyclehook"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsautoscalingnotification"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsautoscalingpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsautoscalingschedule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsbatchcomputeenvironment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsbatchjobdefinition"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsbatchjobqueue"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsbudgetsbudget"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloud9environmentec2"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudformationstack"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudfrontdistribution"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudfrontoriginaccessidentity"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudtrail"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatchdashboard"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatcheventpermission"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatcheventrule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatcheventtarget"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatchlogdestination"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatchlogdestinationpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatchloggroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatchlogmetricfilter"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatchlogresourcepolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatchlogstream"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatchlogsubscriptionfilter"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscloudwatchmetricalarm"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscodebuildproject"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscodecommitrepository"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscodecommittrigger"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscodedeployapp"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscodedeploydeploymentconfig"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscodedeploydeploymentgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscodepipeline"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscognitoidentitypool"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscognitoidentitypoolrolesattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscognitousergroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscognitouserpool"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscognitouserpoolclient"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscognitouserpooldomain"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsconfigconfigrule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsconfigconfigurationrecorder"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsconfigconfigurationrecorderstatus"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsconfigdeliverychannel"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awscustomergateway"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdaxcluster"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdaxparametergroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdaxsubnetgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdbeventsubscription"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdbinstance"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdboptiongroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdbparametergroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdbsecuritygroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdbsnapshot"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdbsubnetgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdefaultnetworkacl"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdefaultroutetable"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdefaultsecuritygroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdefaultsubnet"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdefaultvpc"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdefaultvpcdhcpoptions"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdevicefarmproject"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdirectoryserviceconditionalforwarder"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdirectoryservicedirectory"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdmscertificate"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdmsendpoint"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdmsreplicationinstance"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdmsreplicationsubnetgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdmsreplicationtask"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdxconnection"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdxconnectionassociation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdxlag"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdynamodbglobaltable"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdynamodbtable"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsdynamodbtableitem"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsebssnapshot"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsebsvolume"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsecrlifecyclepolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsecrrepository"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsecrrepositorypolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsecscluster"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsecsservice"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsecstaskdefinition"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsefsfilesystem"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsefsmounttarget"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsegressonlyinternetgateway"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awseip"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awseipassociation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticachecluster"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticacheparametergroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticachereplicationgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticachesecuritygroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticachesubnetgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticbeanstalkapplication"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticbeanstalkapplicationversion"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticbeanstalkconfigurationtemplate"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticbeanstalkenvironment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticsearchdomain"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselasticsearchdomainpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselastictranscoderpipeline"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselastictranscoderpreset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselb"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awselbattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsemrcluster"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsemrinstancegroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsemrsecurityconfiguration"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsflowlog"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsgameliftalias"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsgameliftbuild"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsgameliftfleet"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsglaciervault"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsgluecatalogdatabase"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsgluecatalogtable"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsglueconnection"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsgluejob"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsguarddutydetector"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsguarddutyipset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsguarddutymember"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsguarddutythreatintelset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamaccesskey"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamaccountalias"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamaccountpasswordpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamgroupmembership"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamgrouppolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamgrouppolicyattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiaminstanceprofile"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamopenidconnectprovider"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiampolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiampolicyattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamrole"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamrolepolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamrolepolicyattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamsamlprovider"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamservercertificate"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamservicelinkedrole"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamuser"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamusergroupmembership"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamuserloginprofile"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamuserpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamuserpolicyattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiamusersshkey"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsinspectorassessmenttarget"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsinspectorassessmenttemplate"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsinspectorresourcegroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsinstance"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsinternetgateway"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiotcertificate"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiotpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiotthing"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiotthingtype"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiottopicrule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awskeypair"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awskinesisfirehosedeliverystream"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awskinesisstream"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awskmsalias"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awskmsgrant"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awskmskey"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslambdaalias"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslambdaeventsourcemapping"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslambdafunction"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslambdapermission"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslaunchconfiguration"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslaunchtemplate"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslb"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslbcookiestickinesspolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslbsslnegotiationpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslbtargetgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslbtargetgroupattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslightsaildomain"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslightsailinstance"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslightsailkeypair"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslightsailstaticip"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awslightsailstaticipattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsloadbalancerbackendserverpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsloadbalancerpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsmainroutetableassociation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsmediastorecontainer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsmqbroker"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsmqconfiguration"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsnatgateway"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsnetworkacl"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsnetworkaclrule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsnetworkinterface"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsnetworkinterfaceattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsnetworkinterfacesgattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksapplication"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworkscustomlayer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksganglialayer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworkshaproxylayer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksinstance"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksjavaapplayer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksmemcachedlayer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksmysqllayer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksnodejsapplayer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworkspermission"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksphpapplayer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksrailsapplayer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksrdsdbinstance"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksstack"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksstaticweblayer"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsopsworksuserprofile"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsorganizationsaccount"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsorganizationsorganization"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsorganizationspolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsorganizationspolicyattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsplacementgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsproxyprotocolpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsrdscluster"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsrdsclusterinstance"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsrdsclusterparametergroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsredshiftcluster"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsredshiftparametergroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsredshiftsecuritygroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsredshiftsubnetgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsroute"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsroute53delegationset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsroute53healthcheck"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsroute53querylog"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsroute53record"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsroute53zone"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsroute53zoneassociation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsroutetable"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsroutetableassociation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awss3bucket"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awss3bucketmetric"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awss3bucketnotification"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awss3bucketobject"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awss3bucketpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssecretsmanagersecret"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssecretsmanagersecretversion"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssecuritygroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssecuritygrouprule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsservicecatalogportfolio"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsservicediscoveryprivatednsnamespace"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsservicediscoverypublicdnsnamespace"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsservicediscoveryservice"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssesactivereceiptruleset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssesconfigurationset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssesdomaindkim"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssesdomainidentity"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssesdomainidentityverification"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssesdomainmailfrom"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsseseventdestination"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssesidentitynotificationtopic"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssesreceiptfilter"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssesreceiptrule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssesreceiptruleset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssestemplate"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssfnactivity"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssfnstatemachine"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssimpledbdomain"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssnapshotcreatevolumepermission"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssnsplatformapplication"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssnstopic"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssnstopicpolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssnstopicsubscription"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsspotdatafeedsubscription"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsspotfleetrequest"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsspotinstancerequest"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssqsqueue"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssqsqueuepolicy"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsssmactivation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsssmassociation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsssmdocument"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsssmmaintenancewindow"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsssmmaintenancewindowtarget"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsssmmaintenancewindowtask"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsssmparameter"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsssmpatchbaseline"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsssmpatchgroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsssmresourcedatasync"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awssubnet"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvolumeattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpc"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcdhcpoptions"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcdhcpoptionsassociation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcendpoint"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcendpointconnectionnotification"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcendpointroutetableassociation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcendpointservice"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcendpointserviceallowedprincipal"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcendpointsubnetassociation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcpeeringconnection"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcpeeringconnectionaccepter"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpcpeeringconnectionoptions"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpnconnection"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpnconnectionroute"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpngateway"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpngatewayattachment"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsvpngatewayroutepropagation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafbytematchset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafgeomatchset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafipset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafratebasedrule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregexmatchset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregexpatternset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalbytematchset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalgeomatchset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalipset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalratebasedrule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalregexmatchset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalregexpatternset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalrule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalrulegroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalsizeconstraintset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalsqlinjectionmatchset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalwebacl"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalwebaclassociation"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafregionalxssmatchset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafrule"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafrulegroup"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafsizeconstraintset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafsqlinjectionmatchset"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafwebacl"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awswafxssmatchset"
)

func main() {
	client, resClient := operator.GetKubernetesClient()
	stopCh := make(chan struct{})
	defer close(stopCh)
	informers, err := constructInformers(resClient)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for inf, handler := range informers {
		queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
		inf.AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				key, err := cache.MetaNamespaceKeyFunc(obj)
				if err == nil {
					queue.Add(key)
				}
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				key, err := cache.MetaNamespaceKeyFunc(newObj)
				if err == nil {
					queue.Add(key)
				}
			},
			DeleteFunc: func(obj interface{}) {
				key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
				if err == nil {
					queue.Add(key)
				}
			},
		})
		op := operator.Controller{
			Logger:    log.NewEntry(log.New()),
			ClientSet: client,
			Informer:  inf,
			Queue:     queue,
			Handler:   handler,
		}
		go op.Run(stopCh)
	}

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)
	<-sigTerm
}

func constructInformers(resClient cs.Interface) (map[cache.SharedIndexInformer]operator.Handler, error) {
	return map[cache.SharedIndexInformer]operator.Handler{
		awsacmcertificate.DefaultInformer(resClient):                        &awsacmcertificate.Handler{},
		awsacmcertificatevalidation.DefaultInformer(resClient):              &awsacmcertificatevalidation.Handler{},
		awsacmpcacertificateauthority.DefaultInformer(resClient):            &awsacmpcacertificateauthority.Handler{},
		awsalb.DefaultInformer(resClient):                                   &awsalb.Handler{},
		awsalbtargetgroup.DefaultInformer(resClient):                        &awsalbtargetgroup.Handler{},
		awsalbtargetgroupattachment.DefaultInformer(resClient):              &awsalbtargetgroupattachment.Handler{},
		awsami.DefaultInformer(resClient):                                   &awsami.Handler{},
		awsamicopy.DefaultInformer(resClient):                               &awsamicopy.Handler{},
		awsamifrominstance.DefaultInformer(resClient):                       &awsamifrominstance.Handler{},
		awsamilaunchpermission.DefaultInformer(resClient):                   &awsamilaunchpermission.Handler{},
		awsapigatewayaccount.DefaultInformer(resClient):                     &awsapigatewayaccount.Handler{},
		awsapigatewayapikey.DefaultInformer(resClient):                      &awsapigatewayapikey.Handler{},
		awsapigatewayauthorizer.DefaultInformer(resClient):                  &awsapigatewayauthorizer.Handler{},
		awsapigatewaybasepathmapping.DefaultInformer(resClient):             &awsapigatewaybasepathmapping.Handler{},
		awsapigatewayclientcertificate.DefaultInformer(resClient):           &awsapigatewayclientcertificate.Handler{},
		awsapigatewaydeployment.DefaultInformer(resClient):                  &awsapigatewaydeployment.Handler{},
		awsapigatewaydocumentationpart.DefaultInformer(resClient):           &awsapigatewaydocumentationpart.Handler{},
		awsapigatewaydocumentationversion.DefaultInformer(resClient):        &awsapigatewaydocumentationversion.Handler{},
		awsapigatewaydomainname.DefaultInformer(resClient):                  &awsapigatewaydomainname.Handler{},
		awsapigatewaygatewayresponse.DefaultInformer(resClient):             &awsapigatewaygatewayresponse.Handler{},
		awsapigatewayintegration.DefaultInformer(resClient):                 &awsapigatewayintegration.Handler{},
		awsapigatewayintegrationresponse.DefaultInformer(resClient):         &awsapigatewayintegrationresponse.Handler{},
		awsapigatewaymethod.DefaultInformer(resClient):                      &awsapigatewaymethod.Handler{},
		awsapigatewaymethodresponse.DefaultInformer(resClient):              &awsapigatewaymethodresponse.Handler{},
		awsapigatewaymethodsettings.DefaultInformer(resClient):              &awsapigatewaymethodsettings.Handler{},
		awsapigatewaymodel.DefaultInformer(resClient):                       &awsapigatewaymodel.Handler{},
		awsapigatewayrequestvalidator.DefaultInformer(resClient):            &awsapigatewayrequestvalidator.Handler{},
		awsapigatewayresource.DefaultInformer(resClient):                    &awsapigatewayresource.Handler{},
		awsapigatewayrestapi.DefaultInformer(resClient):                     &awsapigatewayrestapi.Handler{},
		awsapigatewaystage.DefaultInformer(resClient):                       &awsapigatewaystage.Handler{},
		awsapigatewayusageplan.DefaultInformer(resClient):                   &awsapigatewayusageplan.Handler{},
		awsapigatewayusageplankey.DefaultInformer(resClient):                &awsapigatewayusageplankey.Handler{},
		awsapigatewayvpclink.DefaultInformer(resClient):                     &awsapigatewayvpclink.Handler{},
		awsappautoscalingpolicy.DefaultInformer(resClient):                  &awsappautoscalingpolicy.Handler{},
		awsappautoscalingscheduledaction.DefaultInformer(resClient):         &awsappautoscalingscheduledaction.Handler{},
		awsappautoscalingtarget.DefaultInformer(resClient):                  &awsappautoscalingtarget.Handler{},
		awsappcookiestickinesspolicy.DefaultInformer(resClient):             &awsappcookiestickinesspolicy.Handler{},
		awsappsyncdatasource.DefaultInformer(resClient):                     &awsappsyncdatasource.Handler{},
		awsappsyncgraphqlapi.DefaultInformer(resClient):                     &awsappsyncgraphqlapi.Handler{},
		awsathenadatabase.DefaultInformer(resClient):                        &awsathenadatabase.Handler{},
		awsathenanamedquery.DefaultInformer(resClient):                      &awsathenanamedquery.Handler{},
		awsautoscalingattachment.DefaultInformer(resClient):                 &awsautoscalingattachment.Handler{},
		awsautoscalinggroup.DefaultInformer(resClient):                      &awsautoscalinggroup.Handler{},
		awsautoscalinglifecyclehook.DefaultInformer(resClient):              &awsautoscalinglifecyclehook.Handler{},
		awsautoscalingnotification.DefaultInformer(resClient):               &awsautoscalingnotification.Handler{},
		awsautoscalingpolicy.DefaultInformer(resClient):                     &awsautoscalingpolicy.Handler{},
		awsautoscalingschedule.DefaultInformer(resClient):                   &awsautoscalingschedule.Handler{},
		awsbatchcomputeenvironment.DefaultInformer(resClient):               &awsbatchcomputeenvironment.Handler{},
		awsbatchjobdefinition.DefaultInformer(resClient):                    &awsbatchjobdefinition.Handler{},
		awsbatchjobqueue.DefaultInformer(resClient):                         &awsbatchjobqueue.Handler{},
		awsbudgetsbudget.DefaultInformer(resClient):                         &awsbudgetsbudget.Handler{},
		awscloud9environmentec2.DefaultInformer(resClient):                  &awscloud9environmentec2.Handler{},
		awscloudformationstack.DefaultInformer(resClient):                   &awscloudformationstack.Handler{},
		awscloudfrontdistribution.DefaultInformer(resClient):                &awscloudfrontdistribution.Handler{},
		awscloudfrontoriginaccessidentity.DefaultInformer(resClient):        &awscloudfrontoriginaccessidentity.Handler{},
		awscloudtrail.DefaultInformer(resClient):                            &awscloudtrail.Handler{},
		awscloudwatchdashboard.DefaultInformer(resClient):                   &awscloudwatchdashboard.Handler{},
		awscloudwatcheventpermission.DefaultInformer(resClient):             &awscloudwatcheventpermission.Handler{},
		awscloudwatcheventrule.DefaultInformer(resClient):                   &awscloudwatcheventrule.Handler{},
		awscloudwatcheventtarget.DefaultInformer(resClient):                 &awscloudwatcheventtarget.Handler{},
		awscloudwatchlogdestination.DefaultInformer(resClient):              &awscloudwatchlogdestination.Handler{},
		awscloudwatchlogdestinationpolicy.DefaultInformer(resClient):        &awscloudwatchlogdestinationpolicy.Handler{},
		awscloudwatchloggroup.DefaultInformer(resClient):                    &awscloudwatchloggroup.Handler{},
		awscloudwatchlogmetricfilter.DefaultInformer(resClient):             &awscloudwatchlogmetricfilter.Handler{},
		awscloudwatchlogresourcepolicy.DefaultInformer(resClient):           &awscloudwatchlogresourcepolicy.Handler{},
		awscloudwatchlogstream.DefaultInformer(resClient):                   &awscloudwatchlogstream.Handler{},
		awscloudwatchlogsubscriptionfilter.DefaultInformer(resClient):       &awscloudwatchlogsubscriptionfilter.Handler{},
		awscloudwatchmetricalarm.DefaultInformer(resClient):                 &awscloudwatchmetricalarm.Handler{},
		awscodebuildproject.DefaultInformer(resClient):                      &awscodebuildproject.Handler{},
		awscodecommitrepository.DefaultInformer(resClient):                  &awscodecommitrepository.Handler{},
		awscodecommittrigger.DefaultInformer(resClient):                     &awscodecommittrigger.Handler{},
		awscodedeployapp.DefaultInformer(resClient):                         &awscodedeployapp.Handler{},
		awscodedeploydeploymentconfig.DefaultInformer(resClient):            &awscodedeploydeploymentconfig.Handler{},
		awscodedeploydeploymentgroup.DefaultInformer(resClient):             &awscodedeploydeploymentgroup.Handler{},
		awscodepipeline.DefaultInformer(resClient):                          &awscodepipeline.Handler{},
		awscognitoidentitypool.DefaultInformer(resClient):                   &awscognitoidentitypool.Handler{},
		awscognitoidentitypoolrolesattachment.DefaultInformer(resClient):    &awscognitoidentitypoolrolesattachment.Handler{},
		awscognitousergroup.DefaultInformer(resClient):                      &awscognitousergroup.Handler{},
		awscognitouserpool.DefaultInformer(resClient):                       &awscognitouserpool.Handler{},
		awscognitouserpoolclient.DefaultInformer(resClient):                 &awscognitouserpoolclient.Handler{},
		awscognitouserpooldomain.DefaultInformer(resClient):                 &awscognitouserpooldomain.Handler{},
		awsconfigconfigrule.DefaultInformer(resClient):                      &awsconfigconfigrule.Handler{},
		awsconfigconfigurationrecorder.DefaultInformer(resClient):           &awsconfigconfigurationrecorder.Handler{},
		awsconfigconfigurationrecorderstatus.DefaultInformer(resClient):     &awsconfigconfigurationrecorderstatus.Handler{},
		awsconfigdeliverychannel.DefaultInformer(resClient):                 &awsconfigdeliverychannel.Handler{},
		awscustomergateway.DefaultInformer(resClient):                       &awscustomergateway.Handler{},
		awsdaxcluster.DefaultInformer(resClient):                            &awsdaxcluster.Handler{},
		awsdaxparametergroup.DefaultInformer(resClient):                     &awsdaxparametergroup.Handler{},
		awsdaxsubnetgroup.DefaultInformer(resClient):                        &awsdaxsubnetgroup.Handler{},
		awsdbeventsubscription.DefaultInformer(resClient):                   &awsdbeventsubscription.Handler{},
		awsdbinstance.DefaultInformer(resClient):                            &awsdbinstance.Handler{},
		awsdboptiongroup.DefaultInformer(resClient):                         &awsdboptiongroup.Handler{},
		awsdbparametergroup.DefaultInformer(resClient):                      &awsdbparametergroup.Handler{},
		awsdbsecuritygroup.DefaultInformer(resClient):                       &awsdbsecuritygroup.Handler{},
		awsdbsnapshot.DefaultInformer(resClient):                            &awsdbsnapshot.Handler{},
		awsdbsubnetgroup.DefaultInformer(resClient):                         &awsdbsubnetgroup.Handler{},
		awsdefaultnetworkacl.DefaultInformer(resClient):                     &awsdefaultnetworkacl.Handler{},
		awsdefaultroutetable.DefaultInformer(resClient):                     &awsdefaultroutetable.Handler{},
		awsdefaultsecuritygroup.DefaultInformer(resClient):                  &awsdefaultsecuritygroup.Handler{},
		awsdefaultsubnet.DefaultInformer(resClient):                         &awsdefaultsubnet.Handler{},
		awsdefaultvpc.DefaultInformer(resClient):                            &awsdefaultvpc.Handler{},
		awsdefaultvpcdhcpoptions.DefaultInformer(resClient):                 &awsdefaultvpcdhcpoptions.Handler{},
		awsdevicefarmproject.DefaultInformer(resClient):                     &awsdevicefarmproject.Handler{},
		awsdirectoryserviceconditionalforwarder.DefaultInformer(resClient):  &awsdirectoryserviceconditionalforwarder.Handler{},
		awsdirectoryservicedirectory.DefaultInformer(resClient):             &awsdirectoryservicedirectory.Handler{},
		awsdmscertificate.DefaultInformer(resClient):                        &awsdmscertificate.Handler{},
		awsdmsendpoint.DefaultInformer(resClient):                           &awsdmsendpoint.Handler{},
		awsdmsreplicationinstance.DefaultInformer(resClient):                &awsdmsreplicationinstance.Handler{},
		awsdmsreplicationsubnetgroup.DefaultInformer(resClient):             &awsdmsreplicationsubnetgroup.Handler{},
		awsdmsreplicationtask.DefaultInformer(resClient):                    &awsdmsreplicationtask.Handler{},
		awsdxconnection.DefaultInformer(resClient):                          &awsdxconnection.Handler{},
		awsdxconnectionassociation.DefaultInformer(resClient):               &awsdxconnectionassociation.Handler{},
		awsdxlag.DefaultInformer(resClient):                                 &awsdxlag.Handler{},
		awsdynamodbglobaltable.DefaultInformer(resClient):                   &awsdynamodbglobaltable.Handler{},
		awsdynamodbtable.DefaultInformer(resClient):                         &awsdynamodbtable.Handler{},
		awsdynamodbtableitem.DefaultInformer(resClient):                     &awsdynamodbtableitem.Handler{},
		awsebssnapshot.DefaultInformer(resClient):                           &awsebssnapshot.Handler{},
		awsebsvolume.DefaultInformer(resClient):                             &awsebsvolume.Handler{},
		awsecrlifecyclepolicy.DefaultInformer(resClient):                    &awsecrlifecyclepolicy.Handler{},
		awsecrrepository.DefaultInformer(resClient):                         &awsecrrepository.Handler{},
		awsecrrepositorypolicy.DefaultInformer(resClient):                   &awsecrrepositorypolicy.Handler{},
		awsecscluster.DefaultInformer(resClient):                            &awsecscluster.Handler{},
		awsecsservice.DefaultInformer(resClient):                            &awsecsservice.Handler{},
		awsecstaskdefinition.DefaultInformer(resClient):                     &awsecstaskdefinition.Handler{},
		awsefsfilesystem.DefaultInformer(resClient):                         &awsefsfilesystem.Handler{},
		awsefsmounttarget.DefaultInformer(resClient):                        &awsefsmounttarget.Handler{},
		awsegressonlyinternetgateway.DefaultInformer(resClient):             &awsegressonlyinternetgateway.Handler{},
		awseip.DefaultInformer(resClient):                                   &awseip.Handler{},
		awseipassociation.DefaultInformer(resClient):                        &awseipassociation.Handler{},
		awselasticachecluster.DefaultInformer(resClient):                    &awselasticachecluster.Handler{},
		awselasticacheparametergroup.DefaultInformer(resClient):             &awselasticacheparametergroup.Handler{},
		awselasticachereplicationgroup.DefaultInformer(resClient):           &awselasticachereplicationgroup.Handler{},
		awselasticachesecuritygroup.DefaultInformer(resClient):              &awselasticachesecuritygroup.Handler{},
		awselasticachesubnetgroup.DefaultInformer(resClient):                &awselasticachesubnetgroup.Handler{},
		awselasticbeanstalkapplication.DefaultInformer(resClient):           &awselasticbeanstalkapplication.Handler{},
		awselasticbeanstalkapplicationversion.DefaultInformer(resClient):    &awselasticbeanstalkapplicationversion.Handler{},
		awselasticbeanstalkconfigurationtemplate.DefaultInformer(resClient): &awselasticbeanstalkconfigurationtemplate.Handler{},
		awselasticbeanstalkenvironment.DefaultInformer(resClient):           &awselasticbeanstalkenvironment.Handler{},
		awselasticsearchdomain.DefaultInformer(resClient):                   &awselasticsearchdomain.Handler{},
		awselasticsearchdomainpolicy.DefaultInformer(resClient):             &awselasticsearchdomainpolicy.Handler{},
		awselastictranscoderpipeline.DefaultInformer(resClient):             &awselastictranscoderpipeline.Handler{},
		awselastictranscoderpreset.DefaultInformer(resClient):               &awselastictranscoderpreset.Handler{},
		awselb.DefaultInformer(resClient):                                   &awselb.Handler{},
		awselbattachment.DefaultInformer(resClient):                         &awselbattachment.Handler{},
		awsemrcluster.DefaultInformer(resClient):                            &awsemrcluster.Handler{},
		awsemrinstancegroup.DefaultInformer(resClient):                      &awsemrinstancegroup.Handler{},
		awsemrsecurityconfiguration.DefaultInformer(resClient):              &awsemrsecurityconfiguration.Handler{},
		awsflowlog.DefaultInformer(resClient):                               &awsflowlog.Handler{},
		awsgameliftalias.DefaultInformer(resClient):                         &awsgameliftalias.Handler{},
		awsgameliftbuild.DefaultInformer(resClient):                         &awsgameliftbuild.Handler{},
		awsgameliftfleet.DefaultInformer(resClient):                         &awsgameliftfleet.Handler{},
		awsglaciervault.DefaultInformer(resClient):                          &awsglaciervault.Handler{},
		awsgluecatalogdatabase.DefaultInformer(resClient):                   &awsgluecatalogdatabase.Handler{},
		awsgluecatalogtable.DefaultInformer(resClient):                      &awsgluecatalogtable.Handler{},
		awsglueconnection.DefaultInformer(resClient):                        &awsglueconnection.Handler{},
		awsgluejob.DefaultInformer(resClient):                               &awsgluejob.Handler{},
		awsguarddutydetector.DefaultInformer(resClient):                     &awsguarddutydetector.Handler{},
		awsguarddutyipset.DefaultInformer(resClient):                        &awsguarddutyipset.Handler{},
		awsguarddutymember.DefaultInformer(resClient):                       &awsguarddutymember.Handler{},
		awsguarddutythreatintelset.DefaultInformer(resClient):               &awsguarddutythreatintelset.Handler{},
		awsiamaccesskey.DefaultInformer(resClient):                          &awsiamaccesskey.Handler{},
		awsiamaccountalias.DefaultInformer(resClient):                       &awsiamaccountalias.Handler{},
		awsiamaccountpasswordpolicy.DefaultInformer(resClient):              &awsiamaccountpasswordpolicy.Handler{},
		awsiamgroup.DefaultInformer(resClient):                              &awsiamgroup.Handler{},
		awsiamgroupmembership.DefaultInformer(resClient):                    &awsiamgroupmembership.Handler{},
		awsiamgrouppolicy.DefaultInformer(resClient):                        &awsiamgrouppolicy.Handler{},
		awsiamgrouppolicyattachment.DefaultInformer(resClient):              &awsiamgrouppolicyattachment.Handler{},
		awsiaminstanceprofile.DefaultInformer(resClient):                    &awsiaminstanceprofile.Handler{},
		awsiamopenidconnectprovider.DefaultInformer(resClient):              &awsiamopenidconnectprovider.Handler{},
		awsiampolicy.DefaultInformer(resClient):                             &awsiampolicy.Handler{},
		awsiampolicyattachment.DefaultInformer(resClient):                   &awsiampolicyattachment.Handler{},
		awsiamrole.DefaultInformer(resClient):                               &awsiamrole.Handler{},
		awsiamrolepolicy.DefaultInformer(resClient):                         &awsiamrolepolicy.Handler{},
		awsiamrolepolicyattachment.DefaultInformer(resClient):               &awsiamrolepolicyattachment.Handler{},
		awsiamsamlprovider.DefaultInformer(resClient):                       &awsiamsamlprovider.Handler{},
		awsiamservercertificate.DefaultInformer(resClient):                  &awsiamservercertificate.Handler{},
		awsiamservicelinkedrole.DefaultInformer(resClient):                  &awsiamservicelinkedrole.Handler{},
		awsiamuser.DefaultInformer(resClient):                               &awsiamuser.Handler{},
		awsiamusergroupmembership.DefaultInformer(resClient):                &awsiamusergroupmembership.Handler{},
		awsiamuserloginprofile.DefaultInformer(resClient):                   &awsiamuserloginprofile.Handler{},
		awsiamuserpolicy.DefaultInformer(resClient):                         &awsiamuserpolicy.Handler{},
		awsiamuserpolicyattachment.DefaultInformer(resClient):               &awsiamuserpolicyattachment.Handler{},
		awsiamusersshkey.DefaultInformer(resClient):                         &awsiamusersshkey.Handler{},
		awsinspectorassessmenttarget.DefaultInformer(resClient):             &awsinspectorassessmenttarget.Handler{},
		awsinspectorassessmenttemplate.DefaultInformer(resClient):           &awsinspectorassessmenttemplate.Handler{},
		awsinspectorresourcegroup.DefaultInformer(resClient):                &awsinspectorresourcegroup.Handler{},
		awsinstance.DefaultInformer(resClient):                              &awsinstance.Handler{},
		awsinternetgateway.DefaultInformer(resClient):                       &awsinternetgateway.Handler{},
		awsiotcertificate.DefaultInformer(resClient):                        &awsiotcertificate.Handler{},
		awsiotpolicy.DefaultInformer(resClient):                             &awsiotpolicy.Handler{},
		awsiotthing.DefaultInformer(resClient):                              &awsiotthing.Handler{},
		awsiotthingtype.DefaultInformer(resClient):                          &awsiotthingtype.Handler{},
		awsiottopicrule.DefaultInformer(resClient):                          &awsiottopicrule.Handler{},
		awskeypair.DefaultInformer(resClient):                               &awskeypair.Handler{},
		awskinesisfirehosedeliverystream.DefaultInformer(resClient):         &awskinesisfirehosedeliverystream.Handler{},
		awskinesisstream.DefaultInformer(resClient):                         &awskinesisstream.Handler{},
		awskmsalias.DefaultInformer(resClient):                              &awskmsalias.Handler{},
		awskmsgrant.DefaultInformer(resClient):                              &awskmsgrant.Handler{},
		awskmskey.DefaultInformer(resClient):                                &awskmskey.Handler{},
		awslambdaalias.DefaultInformer(resClient):                           &awslambdaalias.Handler{},
		awslambdaeventsourcemapping.DefaultInformer(resClient):              &awslambdaeventsourcemapping.Handler{},
		awslambdafunction.DefaultInformer(resClient):                        &awslambdafunction.Handler{},
		awslambdapermission.DefaultInformer(resClient):                      &awslambdapermission.Handler{},
		awslaunchconfiguration.DefaultInformer(resClient):                   &awslaunchconfiguration.Handler{},
		awslaunchtemplate.DefaultInformer(resClient):                        &awslaunchtemplate.Handler{},
		awslb.DefaultInformer(resClient):                                    &awslb.Handler{},
		awslbcookiestickinesspolicy.DefaultInformer(resClient):              &awslbcookiestickinesspolicy.Handler{},
		awslbsslnegotiationpolicy.DefaultInformer(resClient):                &awslbsslnegotiationpolicy.Handler{},
		awslbtargetgroup.DefaultInformer(resClient):                         &awslbtargetgroup.Handler{},
		awslbtargetgroupattachment.DefaultInformer(resClient):               &awslbtargetgroupattachment.Handler{},
		awslightsaildomain.DefaultInformer(resClient):                       &awslightsaildomain.Handler{},
		awslightsailinstance.DefaultInformer(resClient):                     &awslightsailinstance.Handler{},
		awslightsailkeypair.DefaultInformer(resClient):                      &awslightsailkeypair.Handler{},
		awslightsailstaticip.DefaultInformer(resClient):                     &awslightsailstaticip.Handler{},
		awslightsailstaticipattachment.DefaultInformer(resClient):           &awslightsailstaticipattachment.Handler{},
		awsloadbalancerbackendserverpolicy.DefaultInformer(resClient):       &awsloadbalancerbackendserverpolicy.Handler{},
		awsloadbalancerpolicy.DefaultInformer(resClient):                    &awsloadbalancerpolicy.Handler{},
		awsmainroutetableassociation.DefaultInformer(resClient):             &awsmainroutetableassociation.Handler{},
		awsmediastorecontainer.DefaultInformer(resClient):                   &awsmediastorecontainer.Handler{},
		awsmqbroker.DefaultInformer(resClient):                              &awsmqbroker.Handler{},
		awsmqconfiguration.DefaultInformer(resClient):                       &awsmqconfiguration.Handler{},
		awsnatgateway.DefaultInformer(resClient):                            &awsnatgateway.Handler{},
		awsnetworkacl.DefaultInformer(resClient):                            &awsnetworkacl.Handler{},
		awsnetworkaclrule.DefaultInformer(resClient):                        &awsnetworkaclrule.Handler{},
		awsnetworkinterface.DefaultInformer(resClient):                      &awsnetworkinterface.Handler{},
		awsnetworkinterfaceattachment.DefaultInformer(resClient):            &awsnetworkinterfaceattachment.Handler{},
		awsnetworkinterfacesgattachment.DefaultInformer(resClient):          &awsnetworkinterfacesgattachment.Handler{},
		awsopsworksapplication.DefaultInformer(resClient):                   &awsopsworksapplication.Handler{},
		awsopsworkscustomlayer.DefaultInformer(resClient):                   &awsopsworkscustomlayer.Handler{},
		awsopsworksganglialayer.DefaultInformer(resClient):                  &awsopsworksganglialayer.Handler{},
		awsopsworkshaproxylayer.DefaultInformer(resClient):                  &awsopsworkshaproxylayer.Handler{},
		awsopsworksinstance.DefaultInformer(resClient):                      &awsopsworksinstance.Handler{},
		awsopsworksjavaapplayer.DefaultInformer(resClient):                  &awsopsworksjavaapplayer.Handler{},
		awsopsworksmemcachedlayer.DefaultInformer(resClient):                &awsopsworksmemcachedlayer.Handler{},
		awsopsworksmysqllayer.DefaultInformer(resClient):                    &awsopsworksmysqllayer.Handler{},
		awsopsworksnodejsapplayer.DefaultInformer(resClient):                &awsopsworksnodejsapplayer.Handler{},
		awsopsworkspermission.DefaultInformer(resClient):                    &awsopsworkspermission.Handler{},
		awsopsworksphpapplayer.DefaultInformer(resClient):                   &awsopsworksphpapplayer.Handler{},
		awsopsworksrailsapplayer.DefaultInformer(resClient):                 &awsopsworksrailsapplayer.Handler{},
		awsopsworksrdsdbinstance.DefaultInformer(resClient):                 &awsopsworksrdsdbinstance.Handler{},
		awsopsworksstack.DefaultInformer(resClient):                         &awsopsworksstack.Handler{},
		awsopsworksstaticweblayer.DefaultInformer(resClient):                &awsopsworksstaticweblayer.Handler{},
		awsopsworksuserprofile.DefaultInformer(resClient):                   &awsopsworksuserprofile.Handler{},
		awsorganizationsaccount.DefaultInformer(resClient):                  &awsorganizationsaccount.Handler{},
		awsorganizationsorganization.DefaultInformer(resClient):             &awsorganizationsorganization.Handler{},
		awsorganizationspolicy.DefaultInformer(resClient):                   &awsorganizationspolicy.Handler{},
		awsorganizationspolicyattachment.DefaultInformer(resClient):         &awsorganizationspolicyattachment.Handler{},
		awsplacementgroup.DefaultInformer(resClient):                        &awsplacementgroup.Handler{},
		awsproxyprotocolpolicy.DefaultInformer(resClient):                   &awsproxyprotocolpolicy.Handler{},
		awsrdscluster.DefaultInformer(resClient):                            &awsrdscluster.Handler{},
		awsrdsclusterinstance.DefaultInformer(resClient):                    &awsrdsclusterinstance.Handler{},
		awsrdsclusterparametergroup.DefaultInformer(resClient):              &awsrdsclusterparametergroup.Handler{},
		awsredshiftcluster.DefaultInformer(resClient):                       &awsredshiftcluster.Handler{},
		awsredshiftparametergroup.DefaultInformer(resClient):                &awsredshiftparametergroup.Handler{},
		awsredshiftsecuritygroup.DefaultInformer(resClient):                 &awsredshiftsecuritygroup.Handler{},
		awsredshiftsubnetgroup.DefaultInformer(resClient):                   &awsredshiftsubnetgroup.Handler{},
		awsroute.DefaultInformer(resClient):                                 &awsroute.Handler{},
		awsroute53delegationset.DefaultInformer(resClient):                  &awsroute53delegationset.Handler{},
		awsroute53healthcheck.DefaultInformer(resClient):                    &awsroute53healthcheck.Handler{},
		awsroute53querylog.DefaultInformer(resClient):                       &awsroute53querylog.Handler{},
		awsroute53record.DefaultInformer(resClient):                         &awsroute53record.Handler{},
		awsroute53zone.DefaultInformer(resClient):                           &awsroute53zone.Handler{},
		awsroute53zoneassociation.DefaultInformer(resClient):                &awsroute53zoneassociation.Handler{},
		awsroutetable.DefaultInformer(resClient):                            &awsroutetable.Handler{},
		awsroutetableassociation.DefaultInformer(resClient):                 &awsroutetableassociation.Handler{},
		awss3bucket.DefaultInformer(resClient):                              &awss3bucket.Handler{},
		awss3bucketmetric.DefaultInformer(resClient):                        &awss3bucketmetric.Handler{},
		awss3bucketnotification.DefaultInformer(resClient):                  &awss3bucketnotification.Handler{},
		awss3bucketobject.DefaultInformer(resClient):                        &awss3bucketobject.Handler{},
		awss3bucketpolicy.DefaultInformer(resClient):                        &awss3bucketpolicy.Handler{},
		awssecretsmanagersecret.DefaultInformer(resClient):                  &awssecretsmanagersecret.Handler{},
		awssecretsmanagersecretversion.DefaultInformer(resClient):           &awssecretsmanagersecretversion.Handler{},
		awssecuritygroup.DefaultInformer(resClient):                         &awssecuritygroup.Handler{},
		awssecuritygrouprule.DefaultInformer(resClient):                     &awssecuritygrouprule.Handler{},
		awsservicecatalogportfolio.DefaultInformer(resClient):               &awsservicecatalogportfolio.Handler{},
		awsservicediscoveryprivatednsnamespace.DefaultInformer(resClient):   &awsservicediscoveryprivatednsnamespace.Handler{},
		awsservicediscoverypublicdnsnamespace.DefaultInformer(resClient):    &awsservicediscoverypublicdnsnamespace.Handler{},
		awsservicediscoveryservice.DefaultInformer(resClient):               &awsservicediscoveryservice.Handler{},
		awssesactivereceiptruleset.DefaultInformer(resClient):               &awssesactivereceiptruleset.Handler{},
		awssesconfigurationset.DefaultInformer(resClient):                   &awssesconfigurationset.Handler{},
		awssesdomaindkim.DefaultInformer(resClient):                         &awssesdomaindkim.Handler{},
		awssesdomainidentity.DefaultInformer(resClient):                     &awssesdomainidentity.Handler{},
		awssesdomainidentityverification.DefaultInformer(resClient):         &awssesdomainidentityverification.Handler{},
		awssesdomainmailfrom.DefaultInformer(resClient):                     &awssesdomainmailfrom.Handler{},
		awsseseventdestination.DefaultInformer(resClient):                   &awsseseventdestination.Handler{},
		awssesidentitynotificationtopic.DefaultInformer(resClient):          &awssesidentitynotificationtopic.Handler{},
		awssesreceiptfilter.DefaultInformer(resClient):                      &awssesreceiptfilter.Handler{},
		awssesreceiptrule.DefaultInformer(resClient):                        &awssesreceiptrule.Handler{},
		awssesreceiptruleset.DefaultInformer(resClient):                     &awssesreceiptruleset.Handler{},
		awssestemplate.DefaultInformer(resClient):                           &awssestemplate.Handler{},
		awssfnactivity.DefaultInformer(resClient):                           &awssfnactivity.Handler{},
		awssfnstatemachine.DefaultInformer(resClient):                       &awssfnstatemachine.Handler{},
		awssimpledbdomain.DefaultInformer(resClient):                        &awssimpledbdomain.Handler{},
		awssnapshotcreatevolumepermission.DefaultInformer(resClient):        &awssnapshotcreatevolumepermission.Handler{},
		awssnsplatformapplication.DefaultInformer(resClient):                &awssnsplatformapplication.Handler{},
		awssnstopic.DefaultInformer(resClient):                              &awssnstopic.Handler{},
		awssnstopicpolicy.DefaultInformer(resClient):                        &awssnstopicpolicy.Handler{},
		awssnstopicsubscription.DefaultInformer(resClient):                  &awssnstopicsubscription.Handler{},
		awsspotdatafeedsubscription.DefaultInformer(resClient):              &awsspotdatafeedsubscription.Handler{},
		awsspotfleetrequest.DefaultInformer(resClient):                      &awsspotfleetrequest.Handler{},
		awsspotinstancerequest.DefaultInformer(resClient):                   &awsspotinstancerequest.Handler{},
		awssqsqueue.DefaultInformer(resClient):                              &awssqsqueue.Handler{},
		awssqsqueuepolicy.DefaultInformer(resClient):                        &awssqsqueuepolicy.Handler{},
		awsssmactivation.DefaultInformer(resClient):                         &awsssmactivation.Handler{},
		awsssmassociation.DefaultInformer(resClient):                        &awsssmassociation.Handler{},
		awsssmdocument.DefaultInformer(resClient):                           &awsssmdocument.Handler{},
		awsssmmaintenancewindow.DefaultInformer(resClient):                  &awsssmmaintenancewindow.Handler{},
		awsssmmaintenancewindowtarget.DefaultInformer(resClient):            &awsssmmaintenancewindowtarget.Handler{},
		awsssmmaintenancewindowtask.DefaultInformer(resClient):              &awsssmmaintenancewindowtask.Handler{},
		awsssmparameter.DefaultInformer(resClient):                          &awsssmparameter.Handler{},
		awsssmpatchbaseline.DefaultInformer(resClient):                      &awsssmpatchbaseline.Handler{},
		awsssmpatchgroup.DefaultInformer(resClient):                         &awsssmpatchgroup.Handler{},
		awsssmresourcedatasync.DefaultInformer(resClient):                   &awsssmresourcedatasync.Handler{},
		awssubnet.DefaultInformer(resClient):                                &awssubnet.Handler{},
		awsvolumeattachment.DefaultInformer(resClient):                      &awsvolumeattachment.Handler{},
		awsvpc.DefaultInformer(resClient):                                   &awsvpc.Handler{},
		awsvpcdhcpoptions.DefaultInformer(resClient):                        &awsvpcdhcpoptions.Handler{},
		awsvpcdhcpoptionsassociation.DefaultInformer(resClient):             &awsvpcdhcpoptionsassociation.Handler{},
		awsvpcendpoint.DefaultInformer(resClient):                           &awsvpcendpoint.Handler{},
		awsvpcendpointconnectionnotification.DefaultInformer(resClient):     &awsvpcendpointconnectionnotification.Handler{},
		awsvpcendpointroutetableassociation.DefaultInformer(resClient):      &awsvpcendpointroutetableassociation.Handler{},
		awsvpcendpointservice.DefaultInformer(resClient):                    &awsvpcendpointservice.Handler{},
		awsvpcendpointserviceallowedprincipal.DefaultInformer(resClient):    &awsvpcendpointserviceallowedprincipal.Handler{},
		awsvpcendpointsubnetassociation.DefaultInformer(resClient):          &awsvpcendpointsubnetassociation.Handler{},
		awsvpcpeeringconnection.DefaultInformer(resClient):                  &awsvpcpeeringconnection.Handler{},
		awsvpcpeeringconnectionaccepter.DefaultInformer(resClient):          &awsvpcpeeringconnectionaccepter.Handler{},
		awsvpcpeeringconnectionoptions.DefaultInformer(resClient):           &awsvpcpeeringconnectionoptions.Handler{},
		awsvpnconnection.DefaultInformer(resClient):                         &awsvpnconnection.Handler{},
		awsvpnconnectionroute.DefaultInformer(resClient):                    &awsvpnconnectionroute.Handler{},
		awsvpngateway.DefaultInformer(resClient):                            &awsvpngateway.Handler{},
		awsvpngatewayattachment.DefaultInformer(resClient):                  &awsvpngatewayattachment.Handler{},
		awsvpngatewayroutepropagation.DefaultInformer(resClient):            &awsvpngatewayroutepropagation.Handler{},
		awswafbytematchset.DefaultInformer(resClient):                       &awswafbytematchset.Handler{},
		awswafgeomatchset.DefaultInformer(resClient):                        &awswafgeomatchset.Handler{},
		awswafipset.DefaultInformer(resClient):                              &awswafipset.Handler{},
		awswafratebasedrule.DefaultInformer(resClient):                      &awswafratebasedrule.Handler{},
		awswafregexmatchset.DefaultInformer(resClient):                      &awswafregexmatchset.Handler{},
		awswafregexpatternset.DefaultInformer(resClient):                    &awswafregexpatternset.Handler{},
		awswafregionalbytematchset.DefaultInformer(resClient):               &awswafregionalbytematchset.Handler{},
		awswafregionalgeomatchset.DefaultInformer(resClient):                &awswafregionalgeomatchset.Handler{},
		awswafregionalipset.DefaultInformer(resClient):                      &awswafregionalipset.Handler{},
		awswafregionalratebasedrule.DefaultInformer(resClient):              &awswafregionalratebasedrule.Handler{},
		awswafregionalregexmatchset.DefaultInformer(resClient):              &awswafregionalregexmatchset.Handler{},
		awswafregionalregexpatternset.DefaultInformer(resClient):            &awswafregionalregexpatternset.Handler{},
		awswafregionalrule.DefaultInformer(resClient):                       &awswafregionalrule.Handler{},
		awswafregionalrulegroup.DefaultInformer(resClient):                  &awswafregionalrulegroup.Handler{},
		awswafregionalsizeconstraintset.DefaultInformer(resClient):          &awswafregionalsizeconstraintset.Handler{},
		awswafregionalsqlinjectionmatchset.DefaultInformer(resClient):       &awswafregionalsqlinjectionmatchset.Handler{},
		awswafregionalwebacl.DefaultInformer(resClient):                     &awswafregionalwebacl.Handler{},
		awswafregionalwebaclassociation.DefaultInformer(resClient):          &awswafregionalwebaclassociation.Handler{},
		awswafregionalxssmatchset.DefaultInformer(resClient):                &awswafregionalxssmatchset.Handler{},
		awswafrule.DefaultInformer(resClient):                               &awswafrule.Handler{},
		awswafrulegroup.DefaultInformer(resClient):                          &awswafrulegroup.Handler{},
		awswafsizeconstraintset.DefaultInformer(resClient):                  &awswafsizeconstraintset.Handler{},
		awswafsqlinjectionmatchset.DefaultInformer(resClient):               &awswafsqlinjectionmatchset.Handler{},
		awswafwebacl.DefaultInformer(resClient):                             &awswafwebacl.Handler{},
		awswafxssmatchset.DefaultInformer(resClient):                        &awswafxssmatchset.Handler{},
	}, nil
}
