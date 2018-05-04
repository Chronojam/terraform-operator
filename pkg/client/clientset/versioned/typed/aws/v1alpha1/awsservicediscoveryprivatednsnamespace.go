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

// AwsServiceDiscoveryPrivateDnsNamespacesGetter has a method to return a AwsServiceDiscoveryPrivateDnsNamespaceInterface.
// A group's client should implement this interface.
type AwsServiceDiscoveryPrivateDnsNamespacesGetter interface {
	AwsServiceDiscoveryPrivateDnsNamespaces(namespace string) AwsServiceDiscoveryPrivateDnsNamespaceInterface
}

// AwsServiceDiscoveryPrivateDnsNamespaceInterface has methods to work with AwsServiceDiscoveryPrivateDnsNamespace resources.
type AwsServiceDiscoveryPrivateDnsNamespaceInterface interface {
	Create(*v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace) (*v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace, error)
	Update(*v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace) (*v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsServiceDiscoveryPrivateDnsNamespaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace, err error)
	AwsServiceDiscoveryPrivateDnsNamespaceExpansion
}

// awsServiceDiscoveryPrivateDnsNamespaces implements AwsServiceDiscoveryPrivateDnsNamespaceInterface
type awsServiceDiscoveryPrivateDnsNamespaces struct {
	client rest.Interface
	ns     string
}

// newAwsServiceDiscoveryPrivateDnsNamespaces returns a AwsServiceDiscoveryPrivateDnsNamespaces
func newAwsServiceDiscoveryPrivateDnsNamespaces(c *ChronojamV1alpha1Client, namespace string) *awsServiceDiscoveryPrivateDnsNamespaces {
	return &awsServiceDiscoveryPrivateDnsNamespaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsServiceDiscoveryPrivateDnsNamespace, and returns the corresponding awsServiceDiscoveryPrivateDnsNamespace object, and an error if there is any.
func (c *awsServiceDiscoveryPrivateDnsNamespaces) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace, err error) {
	result = &v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsservicediscoveryprivatednsnamespaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsServiceDiscoveryPrivateDnsNamespaces that match those selectors.
func (c *awsServiceDiscoveryPrivateDnsNamespaces) List(opts v1.ListOptions) (result *v1alpha1.AwsServiceDiscoveryPrivateDnsNamespaceList, err error) {
	result = &v1alpha1.AwsServiceDiscoveryPrivateDnsNamespaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsservicediscoveryprivatednsnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsServiceDiscoveryPrivateDnsNamespaces.
func (c *awsServiceDiscoveryPrivateDnsNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsservicediscoveryprivatednsnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsServiceDiscoveryPrivateDnsNamespace and creates it.  Returns the server's representation of the awsServiceDiscoveryPrivateDnsNamespace, and an error, if there is any.
func (c *awsServiceDiscoveryPrivateDnsNamespaces) Create(awsServiceDiscoveryPrivateDnsNamespace *v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace) (result *v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace, err error) {
	result = &v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsservicediscoveryprivatednsnamespaces").
		Body(awsServiceDiscoveryPrivateDnsNamespace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsServiceDiscoveryPrivateDnsNamespace and updates it. Returns the server's representation of the awsServiceDiscoveryPrivateDnsNamespace, and an error, if there is any.
func (c *awsServiceDiscoveryPrivateDnsNamespaces) Update(awsServiceDiscoveryPrivateDnsNamespace *v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace) (result *v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace, err error) {
	result = &v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsservicediscoveryprivatednsnamespaces").
		Name(awsServiceDiscoveryPrivateDnsNamespace.Name).
		Body(awsServiceDiscoveryPrivateDnsNamespace).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsServiceDiscoveryPrivateDnsNamespace and deletes it. Returns an error if one occurs.
func (c *awsServiceDiscoveryPrivateDnsNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsservicediscoveryprivatednsnamespaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsServiceDiscoveryPrivateDnsNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsservicediscoveryprivatednsnamespaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsServiceDiscoveryPrivateDnsNamespace.
func (c *awsServiceDiscoveryPrivateDnsNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace, err error) {
	result = &v1alpha1.AwsServiceDiscoveryPrivateDnsNamespace{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsservicediscoveryprivatednsnamespaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
