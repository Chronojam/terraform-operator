package awsiampolicy

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	"github.com/chronojam/terraform-operator/pkg/terraform"
)

const ResourceName = "aws_iam_policy"

type Handler struct{}

// Init is used for initialization logic
func (t *Handler) Init() error {
	return nil
}

// ObjectCreated is called when an object is created
func (t *Handler) ObjectCreated(obj interface{}) {
	o := obj.(*v1alpha1.AwsIamPolicy)
	uid := string(o.GetUID())
	b, err := terraform.RenderToTerraform(o.Spec, ResourceName, uid)
	if err != nil {
		log.Info(err)
	}
	err = terraform.WriteToFile(b, fmt.Sprintf("%s-%s", ResourceName, uid))
	if err != nil {
		log.Info(err)
	}
}

// ObjectDeleted is called when an object is deleted
func (t *Handler) ObjectDeleted(obj interface{}) {
}

// ObjectUpdated is called when an object is updated
func (t *Handler) ObjectUpdated(objOld, objNew interface{}) {
}
