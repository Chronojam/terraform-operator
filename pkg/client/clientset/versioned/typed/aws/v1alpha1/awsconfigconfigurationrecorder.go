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

// AwsConfigConfigurationRecordersGetter has a method to return a AwsConfigConfigurationRecorderInterface.
// A group's client should implement this interface.
type AwsConfigConfigurationRecordersGetter interface {
	AwsConfigConfigurationRecorders(namespace string) AwsConfigConfigurationRecorderInterface
}

// AwsConfigConfigurationRecorderInterface has methods to work with AwsConfigConfigurationRecorder resources.
type AwsConfigConfigurationRecorderInterface interface {
	Create(*v1alpha1.AwsConfigConfigurationRecorder) (*v1alpha1.AwsConfigConfigurationRecorder, error)
	Update(*v1alpha1.AwsConfigConfigurationRecorder) (*v1alpha1.AwsConfigConfigurationRecorder, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsConfigConfigurationRecorder, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsConfigConfigurationRecorderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsConfigConfigurationRecorder, err error)
	AwsConfigConfigurationRecorderExpansion
}

// awsConfigConfigurationRecorders implements AwsConfigConfigurationRecorderInterface
type awsConfigConfigurationRecorders struct {
	client rest.Interface
	ns     string
}

// newAwsConfigConfigurationRecorders returns a AwsConfigConfigurationRecorders
func newAwsConfigConfigurationRecorders(c *ChronojamV1alpha1Client, namespace string) *awsConfigConfigurationRecorders {
	return &awsConfigConfigurationRecorders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsConfigConfigurationRecorder, and returns the corresponding awsConfigConfigurationRecorder object, and an error if there is any.
func (c *awsConfigConfigurationRecorders) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsConfigConfigurationRecorder, err error) {
	result = &v1alpha1.AwsConfigConfigurationRecorder{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigconfigurationrecorders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsConfigConfigurationRecorders that match those selectors.
func (c *awsConfigConfigurationRecorders) List(opts v1.ListOptions) (result *v1alpha1.AwsConfigConfigurationRecorderList, err error) {
	result = &v1alpha1.AwsConfigConfigurationRecorderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigconfigurationrecorders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsConfigConfigurationRecorders.
func (c *awsConfigConfigurationRecorders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigconfigurationrecorders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsConfigConfigurationRecorder and creates it.  Returns the server's representation of the awsConfigConfigurationRecorder, and an error, if there is any.
func (c *awsConfigConfigurationRecorders) Create(awsConfigConfigurationRecorder *v1alpha1.AwsConfigConfigurationRecorder) (result *v1alpha1.AwsConfigConfigurationRecorder, err error) {
	result = &v1alpha1.AwsConfigConfigurationRecorder{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsconfigconfigurationrecorders").
		Body(awsConfigConfigurationRecorder).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsConfigConfigurationRecorder and updates it. Returns the server's representation of the awsConfigConfigurationRecorder, and an error, if there is any.
func (c *awsConfigConfigurationRecorders) Update(awsConfigConfigurationRecorder *v1alpha1.AwsConfigConfigurationRecorder) (result *v1alpha1.AwsConfigConfigurationRecorder, err error) {
	result = &v1alpha1.AwsConfigConfigurationRecorder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsconfigconfigurationrecorders").
		Name(awsConfigConfigurationRecorder.Name).
		Body(awsConfigConfigurationRecorder).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsConfigConfigurationRecorder and deletes it. Returns an error if one occurs.
func (c *awsConfigConfigurationRecorders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsconfigconfigurationrecorders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsConfigConfigurationRecorders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsconfigconfigurationrecorders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsConfigConfigurationRecorder.
func (c *awsConfigConfigurationRecorders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsConfigConfigurationRecorder, err error) {
	result = &v1alpha1.AwsConfigConfigurationRecorder{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsconfigconfigurationrecorders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
