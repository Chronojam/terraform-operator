package awsautoscalinglifecyclehook

import (
    log "github.com/Sirupsen/logrus"
    "github.com/chronojam/terraform-operator/pkg/terraform"
    "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
)

const ResourceName="aws_autoscaling_lifecycle_hook"
type Handler struct{}

// Init is used for initialization logic
func (t *Handler) Init() error {
	return nil
}

// ObjectCreated is called when an object is created
func (t *Handler) ObjectCreated(obj interface{}) {
    o := obj.(*v1alpha1.AwsAutoscalingLifecycleHook)
	b, err := terraform.RenderToTerraform(o.Spec, ResourceName, string(o.GetUID()))
	if err != nil {
		log.Info(err)
	}

	log.Infof("%s", string(b))
}

// ObjectDeleted is called when an object is deleted
func (t *Handler) ObjectDeleted(obj interface{}) {
}

// ObjectUpdated is called when an object is updated
func (t *Handler) ObjectUpdated(objOld, objNew interface{}) {
}
