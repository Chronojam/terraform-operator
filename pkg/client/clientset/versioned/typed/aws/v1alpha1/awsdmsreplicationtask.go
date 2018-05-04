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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	scheme "github.com/chronojam/terraform-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsDmsReplicationTasksGetter has a method to return a AwsDmsReplicationTaskInterface.
// A group's client should implement this interface.
type AwsDmsReplicationTasksGetter interface {
	AwsDmsReplicationTasks(namespace string) AwsDmsReplicationTaskInterface
}

// AwsDmsReplicationTaskInterface has methods to work with AwsDmsReplicationTask resources.
type AwsDmsReplicationTaskInterface interface {
	Create(*v1alpha1.AwsDmsReplicationTask) (*v1alpha1.AwsDmsReplicationTask, error)
	Update(*v1alpha1.AwsDmsReplicationTask) (*v1alpha1.AwsDmsReplicationTask, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDmsReplicationTask, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDmsReplicationTaskList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDmsReplicationTask, err error)
	AwsDmsReplicationTaskExpansion
}

// awsDmsReplicationTasks implements AwsDmsReplicationTaskInterface
type awsDmsReplicationTasks struct {
	client rest.Interface
	ns     string
}

// newAwsDmsReplicationTasks returns a AwsDmsReplicationTasks
func newAwsDmsReplicationTasks(c *AwsV1alpha1Client, namespace string) *awsDmsReplicationTasks {
	return &awsDmsReplicationTasks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDmsReplicationTask, and returns the corresponding awsDmsReplicationTask object, and an error if there is any.
func (c *awsDmsReplicationTasks) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDmsReplicationTask, err error) {
	result = &v1alpha1.AwsDmsReplicationTask{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdmsreplicationtasks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDmsReplicationTasks that match those selectors.
func (c *awsDmsReplicationTasks) List(opts v1.ListOptions) (result *v1alpha1.AwsDmsReplicationTaskList, err error) {
	result = &v1alpha1.AwsDmsReplicationTaskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdmsreplicationtasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDmsReplicationTasks.
func (c *awsDmsReplicationTasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdmsreplicationtasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDmsReplicationTask and creates it.  Returns the server's representation of the awsDmsReplicationTask, and an error, if there is any.
func (c *awsDmsReplicationTasks) Create(awsDmsReplicationTask *v1alpha1.AwsDmsReplicationTask) (result *v1alpha1.AwsDmsReplicationTask, err error) {
	result = &v1alpha1.AwsDmsReplicationTask{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdmsreplicationtasks").
		Body(awsDmsReplicationTask).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDmsReplicationTask and updates it. Returns the server's representation of the awsDmsReplicationTask, and an error, if there is any.
func (c *awsDmsReplicationTasks) Update(awsDmsReplicationTask *v1alpha1.AwsDmsReplicationTask) (result *v1alpha1.AwsDmsReplicationTask, err error) {
	result = &v1alpha1.AwsDmsReplicationTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdmsreplicationtasks").
		Name(awsDmsReplicationTask.Name).
		Body(awsDmsReplicationTask).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDmsReplicationTask and deletes it. Returns an error if one occurs.
func (c *awsDmsReplicationTasks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdmsreplicationtasks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDmsReplicationTasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdmsreplicationtasks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDmsReplicationTask.
func (c *awsDmsReplicationTasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDmsReplicationTask, err error) {
	result = &v1alpha1.AwsDmsReplicationTask{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdmsreplicationtasks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
