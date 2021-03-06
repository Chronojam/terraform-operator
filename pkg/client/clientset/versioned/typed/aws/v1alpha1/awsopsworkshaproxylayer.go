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

// AwsOpsworksHaproxyLayersGetter has a method to return a AwsOpsworksHaproxyLayerInterface.
// A group's client should implement this interface.
type AwsOpsworksHaproxyLayersGetter interface {
	AwsOpsworksHaproxyLayers(namespace string) AwsOpsworksHaproxyLayerInterface
}

// AwsOpsworksHaproxyLayerInterface has methods to work with AwsOpsworksHaproxyLayer resources.
type AwsOpsworksHaproxyLayerInterface interface {
	Create(*v1alpha1.AwsOpsworksHaproxyLayer) (*v1alpha1.AwsOpsworksHaproxyLayer, error)
	Update(*v1alpha1.AwsOpsworksHaproxyLayer) (*v1alpha1.AwsOpsworksHaproxyLayer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsOpsworksHaproxyLayer, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsOpsworksHaproxyLayerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksHaproxyLayer, err error)
	AwsOpsworksHaproxyLayerExpansion
}

// awsOpsworksHaproxyLayers implements AwsOpsworksHaproxyLayerInterface
type awsOpsworksHaproxyLayers struct {
	client rest.Interface
	ns     string
}

// newAwsOpsworksHaproxyLayers returns a AwsOpsworksHaproxyLayers
func newAwsOpsworksHaproxyLayers(c *ChronojamV1alpha1Client, namespace string) *awsOpsworksHaproxyLayers {
	return &awsOpsworksHaproxyLayers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsOpsworksHaproxyLayer, and returns the corresponding awsOpsworksHaproxyLayer object, and an error if there is any.
func (c *awsOpsworksHaproxyLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksHaproxyLayer, err error) {
	result = &v1alpha1.AwsOpsworksHaproxyLayer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworkshaproxylayers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOpsworksHaproxyLayers that match those selectors.
func (c *awsOpsworksHaproxyLayers) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksHaproxyLayerList, err error) {
	result = &v1alpha1.AwsOpsworksHaproxyLayerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworkshaproxylayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOpsworksHaproxyLayers.
func (c *awsOpsworksHaproxyLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworkshaproxylayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsOpsworksHaproxyLayer and creates it.  Returns the server's representation of the awsOpsworksHaproxyLayer, and an error, if there is any.
func (c *awsOpsworksHaproxyLayers) Create(awsOpsworksHaproxyLayer *v1alpha1.AwsOpsworksHaproxyLayer) (result *v1alpha1.AwsOpsworksHaproxyLayer, err error) {
	result = &v1alpha1.AwsOpsworksHaproxyLayer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsopsworkshaproxylayers").
		Body(awsOpsworksHaproxyLayer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOpsworksHaproxyLayer and updates it. Returns the server's representation of the awsOpsworksHaproxyLayer, and an error, if there is any.
func (c *awsOpsworksHaproxyLayers) Update(awsOpsworksHaproxyLayer *v1alpha1.AwsOpsworksHaproxyLayer) (result *v1alpha1.AwsOpsworksHaproxyLayer, err error) {
	result = &v1alpha1.AwsOpsworksHaproxyLayer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsopsworkshaproxylayers").
		Name(awsOpsworksHaproxyLayer.Name).
		Body(awsOpsworksHaproxyLayer).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOpsworksHaproxyLayer and deletes it. Returns an error if one occurs.
func (c *awsOpsworksHaproxyLayers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworkshaproxylayers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOpsworksHaproxyLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworkshaproxylayers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOpsworksHaproxyLayer.
func (c *awsOpsworksHaproxyLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksHaproxyLayer, err error) {
	result = &v1alpha1.AwsOpsworksHaproxyLayer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsopsworkshaproxylayers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
