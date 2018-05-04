/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AwsAutoscalingNotificationLister helps list AwsAutoscalingNotifications.
type AwsAutoscalingNotificationLister interface {
	// List lists all AwsAutoscalingNotifications in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAutoscalingNotification, err error)
	// AwsAutoscalingNotifications returns an object that can list and get AwsAutoscalingNotifications.
	AwsAutoscalingNotifications(namespace string) AwsAutoscalingNotificationNamespaceLister
	AwsAutoscalingNotificationListerExpansion
}

// awsAutoscalingNotificationLister implements the AwsAutoscalingNotificationLister interface.
type awsAutoscalingNotificationLister struct {
	indexer cache.Indexer
}

// NewAwsAutoscalingNotificationLister returns a new AwsAutoscalingNotificationLister.
func NewAwsAutoscalingNotificationLister(indexer cache.Indexer) AwsAutoscalingNotificationLister {
	return &awsAutoscalingNotificationLister{indexer: indexer}
}

// List lists all AwsAutoscalingNotifications in the indexer.
func (s *awsAutoscalingNotificationLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAutoscalingNotification, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAutoscalingNotification))
	})
	return ret, err
}

// AwsAutoscalingNotifications returns an object that can list and get AwsAutoscalingNotifications.
func (s *awsAutoscalingNotificationLister) AwsAutoscalingNotifications(namespace string) AwsAutoscalingNotificationNamespaceLister {
	return awsAutoscalingNotificationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsAutoscalingNotificationNamespaceLister helps list and get AwsAutoscalingNotifications.
type AwsAutoscalingNotificationNamespaceLister interface {
	// List lists all AwsAutoscalingNotifications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAutoscalingNotification, err error)
	// Get retrieves the AwsAutoscalingNotification from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsAutoscalingNotification, error)
	AwsAutoscalingNotificationNamespaceListerExpansion
}

// awsAutoscalingNotificationNamespaceLister implements the AwsAutoscalingNotificationNamespaceLister
// interface.
type awsAutoscalingNotificationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsAutoscalingNotifications in the indexer for a given namespace.
func (s awsAutoscalingNotificationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAutoscalingNotification, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAutoscalingNotification))
	})
	return ret, err
}

// Get retrieves the AwsAutoscalingNotification from the indexer for a given namespace and name.
func (s awsAutoscalingNotificationNamespaceLister) Get(name string) (*v1alpha1.AwsAutoscalingNotification, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsautoscalingnotification"), name)
	}
	return obj.(*v1alpha1.AwsAutoscalingNotification), nil
}
