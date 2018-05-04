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

// AwsGlueConnectionsGetter has a method to return a AwsGlueConnectionInterface.
// A group's client should implement this interface.
type AwsGlueConnectionsGetter interface {
	AwsGlueConnections(namespace string) AwsGlueConnectionInterface
}

// AwsGlueConnectionInterface has methods to work with AwsGlueConnection resources.
type AwsGlueConnectionInterface interface {
	Create(*v1alpha1.AwsGlueConnection) (*v1alpha1.AwsGlueConnection, error)
	Update(*v1alpha1.AwsGlueConnection) (*v1alpha1.AwsGlueConnection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsGlueConnection, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsGlueConnectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGlueConnection, err error)
	AwsGlueConnectionExpansion
}

// awsGlueConnections implements AwsGlueConnectionInterface
type awsGlueConnections struct {
	client rest.Interface
	ns     string
}

// newAwsGlueConnections returns a AwsGlueConnections
func newAwsGlueConnections(c *ChronojamV1alpha1Client, namespace string) *awsGlueConnections {
	return &awsGlueConnections{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsGlueConnection, and returns the corresponding awsGlueConnection object, and an error if there is any.
func (c *awsGlueConnections) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGlueConnection, err error) {
	result = &v1alpha1.AwsGlueConnection{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsglueconnections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsGlueConnections that match those selectors.
func (c *awsGlueConnections) List(opts v1.ListOptions) (result *v1alpha1.AwsGlueConnectionList, err error) {
	result = &v1alpha1.AwsGlueConnectionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsglueconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsGlueConnections.
func (c *awsGlueConnections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsglueconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsGlueConnection and creates it.  Returns the server's representation of the awsGlueConnection, and an error, if there is any.
func (c *awsGlueConnections) Create(awsGlueConnection *v1alpha1.AwsGlueConnection) (result *v1alpha1.AwsGlueConnection, err error) {
	result = &v1alpha1.AwsGlueConnection{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsglueconnections").
		Body(awsGlueConnection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsGlueConnection and updates it. Returns the server's representation of the awsGlueConnection, and an error, if there is any.
func (c *awsGlueConnections) Update(awsGlueConnection *v1alpha1.AwsGlueConnection) (result *v1alpha1.AwsGlueConnection, err error) {
	result = &v1alpha1.AwsGlueConnection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsglueconnections").
		Name(awsGlueConnection.Name).
		Body(awsGlueConnection).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsGlueConnection and deletes it. Returns an error if one occurs.
func (c *awsGlueConnections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsglueconnections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsGlueConnections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsglueconnections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsGlueConnection.
func (c *awsGlueConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGlueConnection, err error) {
	result = &v1alpha1.AwsGlueConnection{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsglueconnections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
