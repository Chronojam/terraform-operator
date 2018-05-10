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

// AwsWafWebAclsGetter has a method to return a AwsWafWebAclInterface.
// A group's client should implement this interface.
type AwsWafWebAclsGetter interface {
	AwsWafWebAcls(namespace string) AwsWafWebAclInterface
}

// AwsWafWebAclInterface has methods to work with AwsWafWebAcl resources.
type AwsWafWebAclInterface interface {
	Create(*v1alpha1.AwsWafWebAcl) (*v1alpha1.AwsWafWebAcl, error)
	Update(*v1alpha1.AwsWafWebAcl) (*v1alpha1.AwsWafWebAcl, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafWebAcl, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafWebAclList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafWebAcl, err error)
	AwsWafWebAclExpansion
}

// awsWafWebAcls implements AwsWafWebAclInterface
type awsWafWebAcls struct {
	client rest.Interface
	ns     string
}

// newAwsWafWebAcls returns a AwsWafWebAcls
func newAwsWafWebAcls(c *ChronojamV1alpha1Client, namespace string) *awsWafWebAcls {
	return &awsWafWebAcls{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsWafWebAcl, and returns the corresponding awsWafWebAcl object, and an error if there is any.
func (c *awsWafWebAcls) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafWebAcl, err error) {
	result = &v1alpha1.AwsWafWebAcl{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafwebacls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafWebAcls that match those selectors.
func (c *awsWafWebAcls) List(opts v1.ListOptions) (result *v1alpha1.AwsWafWebAclList, err error) {
	result = &v1alpha1.AwsWafWebAclList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafwebacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafWebAcls.
func (c *awsWafWebAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awswafwebacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsWafWebAcl and creates it.  Returns the server's representation of the awsWafWebAcl, and an error, if there is any.
func (c *awsWafWebAcls) Create(awsWafWebAcl *v1alpha1.AwsWafWebAcl) (result *v1alpha1.AwsWafWebAcl, err error) {
	result = &v1alpha1.AwsWafWebAcl{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awswafwebacls").
		Body(awsWafWebAcl).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafWebAcl and updates it. Returns the server's representation of the awsWafWebAcl, and an error, if there is any.
func (c *awsWafWebAcls) Update(awsWafWebAcl *v1alpha1.AwsWafWebAcl) (result *v1alpha1.AwsWafWebAcl, err error) {
	result = &v1alpha1.AwsWafWebAcl{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awswafwebacls").
		Name(awsWafWebAcl.Name).
		Body(awsWafWebAcl).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafWebAcl and deletes it. Returns an error if one occurs.
func (c *awsWafWebAcls) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafwebacls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafWebAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafwebacls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafWebAcl.
func (c *awsWafWebAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafWebAcl, err error) {
	result = &v1alpha1.AwsWafWebAcl{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awswafwebacls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
