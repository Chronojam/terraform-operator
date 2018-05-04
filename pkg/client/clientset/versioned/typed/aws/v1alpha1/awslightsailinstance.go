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

// AwsLightsailInstancesGetter has a method to return a AwsLightsailInstanceInterface.
// A group's client should implement this interface.
type AwsLightsailInstancesGetter interface {
	AwsLightsailInstances(namespace string) AwsLightsailInstanceInterface
}

// AwsLightsailInstanceInterface has methods to work with AwsLightsailInstance resources.
type AwsLightsailInstanceInterface interface {
	Create(*v1alpha1.AwsLightsailInstance) (*v1alpha1.AwsLightsailInstance, error)
	Update(*v1alpha1.AwsLightsailInstance) (*v1alpha1.AwsLightsailInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsLightsailInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsLightsailInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLightsailInstance, err error)
	AwsLightsailInstanceExpansion
}

// awsLightsailInstances implements AwsLightsailInstanceInterface
type awsLightsailInstances struct {
	client rest.Interface
	ns     string
}

// newAwsLightsailInstances returns a AwsLightsailInstances
func newAwsLightsailInstances(c *AwsV1alpha1Client, namespace string) *awsLightsailInstances {
	return &awsLightsailInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsLightsailInstance, and returns the corresponding awsLightsailInstance object, and an error if there is any.
func (c *awsLightsailInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLightsailInstance, err error) {
	result = &v1alpha1.AwsLightsailInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslightsailinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLightsailInstances that match those selectors.
func (c *awsLightsailInstances) List(opts v1.ListOptions) (result *v1alpha1.AwsLightsailInstanceList, err error) {
	result = &v1alpha1.AwsLightsailInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslightsailinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLightsailInstances.
func (c *awsLightsailInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awslightsailinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsLightsailInstance and creates it.  Returns the server's representation of the awsLightsailInstance, and an error, if there is any.
func (c *awsLightsailInstances) Create(awsLightsailInstance *v1alpha1.AwsLightsailInstance) (result *v1alpha1.AwsLightsailInstance, err error) {
	result = &v1alpha1.AwsLightsailInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awslightsailinstances").
		Body(awsLightsailInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLightsailInstance and updates it. Returns the server's representation of the awsLightsailInstance, and an error, if there is any.
func (c *awsLightsailInstances) Update(awsLightsailInstance *v1alpha1.AwsLightsailInstance) (result *v1alpha1.AwsLightsailInstance, err error) {
	result = &v1alpha1.AwsLightsailInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awslightsailinstances").
		Name(awsLightsailInstance.Name).
		Body(awsLightsailInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLightsailInstance and deletes it. Returns an error if one occurs.
func (c *awsLightsailInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslightsailinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLightsailInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslightsailinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLightsailInstance.
func (c *awsLightsailInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLightsailInstance, err error) {
	result = &v1alpha1.AwsLightsailInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awslightsailinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
