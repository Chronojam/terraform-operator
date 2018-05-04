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

// AwsCloud9EnvironmentEc2sGetter has a method to return a AwsCloud9EnvironmentEc2Interface.
// A group's client should implement this interface.
type AwsCloud9EnvironmentEc2sGetter interface {
	AwsCloud9EnvironmentEc2s(namespace string) AwsCloud9EnvironmentEc2Interface
}

// AwsCloud9EnvironmentEc2Interface has methods to work with AwsCloud9EnvironmentEc2 resources.
type AwsCloud9EnvironmentEc2Interface interface {
	Create(*v1alpha1.AwsCloud9EnvironmentEc2) (*v1alpha1.AwsCloud9EnvironmentEc2, error)
	Update(*v1alpha1.AwsCloud9EnvironmentEc2) (*v1alpha1.AwsCloud9EnvironmentEc2, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCloud9EnvironmentEc2, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCloud9EnvironmentEc2List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloud9EnvironmentEc2, err error)
	AwsCloud9EnvironmentEc2Expansion
}

// awsCloud9EnvironmentEc2s implements AwsCloud9EnvironmentEc2Interface
type awsCloud9EnvironmentEc2s struct {
	client rest.Interface
	ns     string
}

// newAwsCloud9EnvironmentEc2s returns a AwsCloud9EnvironmentEc2s
func newAwsCloud9EnvironmentEc2s(c *ChronojamV1alpha1Client, namespace string) *awsCloud9EnvironmentEc2s {
	return &awsCloud9EnvironmentEc2s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsCloud9EnvironmentEc2, and returns the corresponding awsCloud9EnvironmentEc2 object, and an error if there is any.
func (c *awsCloud9EnvironmentEc2s) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloud9EnvironmentEc2, err error) {
	result = &v1alpha1.AwsCloud9EnvironmentEc2{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscloud9environmentec2s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCloud9EnvironmentEc2s that match those selectors.
func (c *awsCloud9EnvironmentEc2s) List(opts v1.ListOptions) (result *v1alpha1.AwsCloud9EnvironmentEc2List, err error) {
	result = &v1alpha1.AwsCloud9EnvironmentEc2List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscloud9environmentec2s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCloud9EnvironmentEc2s.
func (c *awsCloud9EnvironmentEc2s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscloud9environmentec2s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsCloud9EnvironmentEc2 and creates it.  Returns the server's representation of the awsCloud9EnvironmentEc2, and an error, if there is any.
func (c *awsCloud9EnvironmentEc2s) Create(awsCloud9EnvironmentEc2 *v1alpha1.AwsCloud9EnvironmentEc2) (result *v1alpha1.AwsCloud9EnvironmentEc2, err error) {
	result = &v1alpha1.AwsCloud9EnvironmentEc2{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscloud9environmentec2s").
		Body(awsCloud9EnvironmentEc2).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCloud9EnvironmentEc2 and updates it. Returns the server's representation of the awsCloud9EnvironmentEc2, and an error, if there is any.
func (c *awsCloud9EnvironmentEc2s) Update(awsCloud9EnvironmentEc2 *v1alpha1.AwsCloud9EnvironmentEc2) (result *v1alpha1.AwsCloud9EnvironmentEc2, err error) {
	result = &v1alpha1.AwsCloud9EnvironmentEc2{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscloud9environmentec2s").
		Name(awsCloud9EnvironmentEc2.Name).
		Body(awsCloud9EnvironmentEc2).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCloud9EnvironmentEc2 and deletes it. Returns an error if one occurs.
func (c *awsCloud9EnvironmentEc2s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscloud9environmentec2s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCloud9EnvironmentEc2s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscloud9environmentec2s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCloud9EnvironmentEc2.
func (c *awsCloud9EnvironmentEc2s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloud9EnvironmentEc2, err error) {
	result = &v1alpha1.AwsCloud9EnvironmentEc2{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscloud9environmentec2s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
