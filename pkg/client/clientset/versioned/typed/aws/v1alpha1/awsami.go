/*
Copyright 2018 The Kubernetes Authors.

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

// AwsAmisGetter has a method to return a AwsAmiInterface.
// A group's client should implement this interface.
type AwsAmisGetter interface {
	AwsAmis(namespace string) AwsAmiInterface
}

// AwsAmiInterface has methods to work with AwsAmi resources.
type AwsAmiInterface interface {
	Create(*v1alpha1.AwsAmi) (*v1alpha1.AwsAmi, error)
	Update(*v1alpha1.AwsAmi) (*v1alpha1.AwsAmi, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAmi, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAmiList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAmi, err error)
	AwsAmiExpansion
}

// awsAmis implements AwsAmiInterface
type awsAmis struct {
	client rest.Interface
	ns     string
}

// newAwsAmis returns a AwsAmis
func newAwsAmis(c *ChronojamV1alpha1Client, namespace string) *awsAmis {
	return &awsAmis{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsAmi, and returns the corresponding awsAmi object, and an error if there is any.
func (c *awsAmis) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAmi, err error) {
	result = &v1alpha1.AwsAmi{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsamis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAmis that match those selectors.
func (c *awsAmis) List(opts v1.ListOptions) (result *v1alpha1.AwsAmiList, err error) {
	result = &v1alpha1.AwsAmiList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsamis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAmis.
func (c *awsAmis) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsamis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsAmi and creates it.  Returns the server's representation of the awsAmi, and an error, if there is any.
func (c *awsAmis) Create(awsAmi *v1alpha1.AwsAmi) (result *v1alpha1.AwsAmi, err error) {
	result = &v1alpha1.AwsAmi{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsamis").
		Body(awsAmi).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAmi and updates it. Returns the server's representation of the awsAmi, and an error, if there is any.
func (c *awsAmis) Update(awsAmi *v1alpha1.AwsAmi) (result *v1alpha1.AwsAmi, err error) {
	result = &v1alpha1.AwsAmi{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsamis").
		Name(awsAmi.Name).
		Body(awsAmi).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAmi and deletes it. Returns an error if one occurs.
func (c *awsAmis) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsamis").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAmis) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsamis").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAmi.
func (c *awsAmis) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAmi, err error) {
	result = &v1alpha1.AwsAmi{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsamis").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
