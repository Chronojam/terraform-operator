#!/bin/bash

READ_PATH=pkg/apis/aws/v1alpha1
WRITE_PATH=pkg/operator/aws/v1alpha1

#mkdir -p WRITE_PATH

HANDLER_CONTENT='package LC

import (
    log "github.com/Sirupsen/logrus"
)

type Handler struct{}

// Init is used for initialization logic
func (t *Handler) Init() error {
	log.Info("Handler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *Handler) ObjectCreated(obj interface{}) {
	log.Info("Handler.ObjectCreated")
}

// ObjectDeleted is called when an object is deleted
func (t *Handler) ObjectDeleted(obj interface{}) {
	log.Info("Handler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *Handler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("Handler.ObjectUpdated")
}'

INFORMER_CONTENT='package LC

import (
    meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    csinf_v1 "github.com/trussle/terraform-operator/pkg/client/informers/externalversions/aws/v1"
    cs "github.com/trussle/terraform-operator/pkg/client/clientset/versioned"
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

function write_file {
    RP=$1
    FN=$(echo $1 | awk -F/ '{ print $NF }')
    ST=$(cat $1 | grep "struct" | head -n1 | grep -v "List" | grep -v "Spec" | awk '{print $2}')
    if [ "$ST" ==  "" ]; then
        return
    fi
    LC=$(echo "$ST" | awk '{ print tolower($1) }' )
    LC=$(echo $LC | tr -d '\n')
    echo $LC
    DIR=$WRITE_PATH/$LC

    mkdir -p $DIR

    echo "$HANDLER_CONTENT" > $DIR/handler.go
    echo "$INFORMER_CONTENT" > $DIR/informer.go

    echo "$DIR/handler.go"

    sed -i '' "s/LC/$LC/g" $DIR/handler.go $DIR/informer.go
    sed -i '' "s/ST/$ST/g" $DIR/handler.go $DIR/informer.go

    echo "Written:  $DIR/handler.go"
    echo "Written:  $DIR/informer.go"
}


for i in $(find $READ_PATH -iname "aws_*"); do
    write_file $i
done
