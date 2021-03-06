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

// AwsGuarddutyDetectorsGetter has a method to return a AwsGuarddutyDetectorInterface.
// A group's client should implement this interface.
type AwsGuarddutyDetectorsGetter interface {
	AwsGuarddutyDetectors(namespace string) AwsGuarddutyDetectorInterface
}

// AwsGuarddutyDetectorInterface has methods to work with AwsGuarddutyDetector resources.
type AwsGuarddutyDetectorInterface interface {
	Create(*v1alpha1.AwsGuarddutyDetector) (*v1alpha1.AwsGuarddutyDetector, error)
	Update(*v1alpha1.AwsGuarddutyDetector) (*v1alpha1.AwsGuarddutyDetector, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsGuarddutyDetector, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsGuarddutyDetectorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGuarddutyDetector, err error)
	AwsGuarddutyDetectorExpansion
}

// awsGuarddutyDetectors implements AwsGuarddutyDetectorInterface
type awsGuarddutyDetectors struct {
	client rest.Interface
	ns     string
}

// newAwsGuarddutyDetectors returns a AwsGuarddutyDetectors
func newAwsGuarddutyDetectors(c *ChronojamV1alpha1Client, namespace string) *awsGuarddutyDetectors {
	return &awsGuarddutyDetectors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsGuarddutyDetector, and returns the corresponding awsGuarddutyDetector object, and an error if there is any.
func (c *awsGuarddutyDetectors) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGuarddutyDetector, err error) {
	result = &v1alpha1.AwsGuarddutyDetector{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsguarddutydetectors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsGuarddutyDetectors that match those selectors.
func (c *awsGuarddutyDetectors) List(opts v1.ListOptions) (result *v1alpha1.AwsGuarddutyDetectorList, err error) {
	result = &v1alpha1.AwsGuarddutyDetectorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsguarddutydetectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsGuarddutyDetectors.
func (c *awsGuarddutyDetectors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsguarddutydetectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsGuarddutyDetector and creates it.  Returns the server's representation of the awsGuarddutyDetector, and an error, if there is any.
func (c *awsGuarddutyDetectors) Create(awsGuarddutyDetector *v1alpha1.AwsGuarddutyDetector) (result *v1alpha1.AwsGuarddutyDetector, err error) {
	result = &v1alpha1.AwsGuarddutyDetector{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsguarddutydetectors").
		Body(awsGuarddutyDetector).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsGuarddutyDetector and updates it. Returns the server's representation of the awsGuarddutyDetector, and an error, if there is any.
func (c *awsGuarddutyDetectors) Update(awsGuarddutyDetector *v1alpha1.AwsGuarddutyDetector) (result *v1alpha1.AwsGuarddutyDetector, err error) {
	result = &v1alpha1.AwsGuarddutyDetector{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsguarddutydetectors").
		Name(awsGuarddutyDetector.Name).
		Body(awsGuarddutyDetector).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsGuarddutyDetector and deletes it. Returns an error if one occurs.
func (c *awsGuarddutyDetectors) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsguarddutydetectors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsGuarddutyDetectors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsguarddutydetectors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsGuarddutyDetector.
func (c *awsGuarddutyDetectors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGuarddutyDetector, err error) {
	result = &v1alpha1.AwsGuarddutyDetector{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsguarddutydetectors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
