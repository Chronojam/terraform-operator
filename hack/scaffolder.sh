#!/bin/bash

READ_PATH=pkg/apis/aws/v1alpha1
WRITE_PATH=pkg/operator/aws/v1alpha1

#mkdir -p WRITE_PATH

HANDLER_CONTENT='package LC

import (
    log "github.com/Sirupsen/logrus"
    "github.com/chronojam/terraform-operator/pkg/terraform"
    "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
)

const ResourceName="RN"
type Handler struct{}

// Init is used for initialization logic
func (t *Handler) Init() error {
	return nil
}

// ObjectCreated is called when an object is created
func (t *Handler) ObjectCreated(obj interface{}) {
    o := obj.(*v1alpha1.ST)
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
}'

INFORMER_CONTENT='package LC

import (
    meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    csinf_v1 "github.com/chronojam/terraform-operator/pkg/client/informers/externalversions/aws/v1alpha1"
    cs "github.com/chronojam/terraform-operator/pkg/client/clientset/versioned"
    "k8s.io/client-go/tools/cache"
)

func DefaultInformer(resClient cs.Interface) cache.SharedIndexInformer {
    return csinf_v1.NewSTInformer(
			resClient,
			meta_v1.NamespaceAll,
			0,
			cache.Indexers{},
		)
}'


CRD_TEMPLATE="
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: PLURAL.chronojam.co.uk
spec:
  group: chronojam.co.uk
  version: v1alpha1
  names:
    kind: ST
    plural: PLURAL
  scope: Namespaced
"


function write_file {
    RP=$1
    FN=$(echo $1 | awk -F/ '{ print $NF }')
    ST=$(cat $1 | grep "struct" | head -n1 | grep -v "List" | grep -v "Spec" | awk '{print $2}')
    if [ "$ST" ==  "" ]; then
        return
    fi
    LC=$(echo "$ST" | awk '{ print tolower($1) }' )
    LC=$(echo $LC | tr -d '\n')
    DIR=$WRITE_PATH/$LC

    RN=$(echo $FN | sed -e 's/.go//g')

    PLURAL=$(.build/plural $LC)

    mkdir -p $DIR
    mkdir -p crds/

    echo "$HANDLER_CONTENT" > $DIR/handler.go
    echo "$INFORMER_CONTENT" > $DIR/informer.go
    echo "$CRD_TEMPLATE" > crds/$LC.yaml
    echo "RN==$RN"

    sed -i '' "s/LC/$LC/g" $DIR/handler.go $DIR/informer.go crds/$LC.yaml
    sed -i '' "s/ST/$ST/g" $DIR/handler.go $DIR/informer.go crds/$LC.yaml
    sed -i '' "s/PLURAL/$PLURAL/g" $DIR/handler.go $DIR/informer.go crds/$LC.yaml
    sed -i '' "s/RN/$RN/g" $DIR/handler.go $DIR/informer.go crds/$LC.yaml

    echo "Written:  $DIR/handler.go"
    echo "Written:  $DIR/informer.go"
    echo "Written:  crds/$LC.yaml"

}


for i in $(find $READ_PATH -iname "aws_*"); do
    write_file $i
done
