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

// AwsGuarddutyThreatintelsetsGetter has a method to return a AwsGuarddutyThreatintelsetInterface.
// A group's client should implement this interface.
type AwsGuarddutyThreatintelsetsGetter interface {
	AwsGuarddutyThreatintelsets(namespace string) AwsGuarddutyThreatintelsetInterface
}

// AwsGuarddutyThreatintelsetInterface has methods to work with AwsGuarddutyThreatintelset resources.
type AwsGuarddutyThreatintelsetInterface interface {
	Create(*v1alpha1.AwsGuarddutyThreatintelset) (*v1alpha1.AwsGuarddutyThreatintelset, error)
	Update(*v1alpha1.AwsGuarddutyThreatintelset) (*v1alpha1.AwsGuarddutyThreatintelset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsGuarddutyThreatintelset, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsGuarddutyThreatintelsetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGuarddutyThreatintelset, err error)
	AwsGuarddutyThreatintelsetExpansion
}

// awsGuarddutyThreatintelsets implements AwsGuarddutyThreatintelsetInterface
type awsGuarddutyThreatintelsets struct {
	client rest.Interface
	ns     string
}

// newAwsGuarddutyThreatintelsets returns a AwsGuarddutyThreatintelsets
func newAwsGuarddutyThreatintelsets(c *ChronojamV1alpha1Client, namespace string) *awsGuarddutyThreatintelsets {
	return &awsGuarddutyThreatintelsets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsGuarddutyThreatintelset, and returns the corresponding awsGuarddutyThreatintelset object, and an error if there is any.
func (c *awsGuarddutyThreatintelsets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGuarddutyThreatintelset, err error) {
	result = &v1alpha1.AwsGuarddutyThreatintelset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsguarddutythreatintelsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsGuarddutyThreatintelsets that match those selectors.
func (c *awsGuarddutyThreatintelsets) List(opts v1.ListOptions) (result *v1alpha1.AwsGuarddutyThreatintelsetList, err error) {
	result = &v1alpha1.AwsGuarddutyThreatintelsetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsguarddutythreatintelsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsGuarddutyThreatintelsets.
func (c *awsGuarddutyThreatintelsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsguarddutythreatintelsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsGuarddutyThreatintelset and creates it.  Returns the server's representation of the awsGuarddutyThreatintelset, and an error, if there is any.
func (c *awsGuarddutyThreatintelsets) Create(awsGuarddutyThreatintelset *v1alpha1.AwsGuarddutyThreatintelset) (result *v1alpha1.AwsGuarddutyThreatintelset, err error) {
	result = &v1alpha1.AwsGuarddutyThreatintelset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsguarddutythreatintelsets").
		Body(awsGuarddutyThreatintelset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsGuarddutyThreatintelset and updates it. Returns the server's representation of the awsGuarddutyThreatintelset, and an error, if there is any.
func (c *awsGuarddutyThreatintelsets) Update(awsGuarddutyThreatintelset *v1alpha1.AwsGuarddutyThreatintelset) (result *v1alpha1.AwsGuarddutyThreatintelset, err error) {
	result = &v1alpha1.AwsGuarddutyThreatintelset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsguarddutythreatintelsets").
		Name(awsGuarddutyThreatintelset.Name).
		Body(awsGuarddutyThreatintelset).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsGuarddutyThreatintelset and deletes it. Returns an error if one occurs.
func (c *awsGuarddutyThreatintelsets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsguarddutythreatintelsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsGuarddutyThreatintelsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsguarddutythreatintelsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsGuarddutyThreatintelset.
func (c *awsGuarddutyThreatintelsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGuarddutyThreatintelset, err error) {
	result = &v1alpha1.AwsGuarddutyThreatintelset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsguarddutythreatintelsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
