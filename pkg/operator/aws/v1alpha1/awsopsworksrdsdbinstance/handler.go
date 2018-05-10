package awsopsworksrdsdbinstance

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
}
