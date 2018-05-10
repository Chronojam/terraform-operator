package controller

import log "github.com/Sirupsen/logrus"

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

type DummyHandler struct{}

// Init is used for initialization logic
func (t *DummyHandler) Init() error {
	log.Info("DummyHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *DummyHandler) ObjectCreated(obj interface{}) {
	log.Info("DummyHandler.ObjectCreated")
}

// ObjectDeleted is called when an object is deleted
func (t *DummyHandler) ObjectDeleted(obj interface{}) {
	log.Info("DummyHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *DummyHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("DummyHandler.ObjectUpdated")
}
