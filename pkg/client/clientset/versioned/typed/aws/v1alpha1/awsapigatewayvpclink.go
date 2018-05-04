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

// AwsApiGatewayVpcLinksGetter has a method to return a AwsApiGatewayVpcLinkInterface.
// A group's client should implement this interface.
type AwsApiGatewayVpcLinksGetter interface {
	AwsApiGatewayVpcLinks(namespace string) AwsApiGatewayVpcLinkInterface
}

// AwsApiGatewayVpcLinkInterface has methods to work with AwsApiGatewayVpcLink resources.
type AwsApiGatewayVpcLinkInterface interface {
	Create(*v1alpha1.AwsApiGatewayVpcLink) (*v1alpha1.AwsApiGatewayVpcLink, error)
	Update(*v1alpha1.AwsApiGatewayVpcLink) (*v1alpha1.AwsApiGatewayVpcLink, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayVpcLink, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayVpcLinkList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayVpcLink, err error)
	AwsApiGatewayVpcLinkExpansion
}

// awsApiGatewayVpcLinks implements AwsApiGatewayVpcLinkInterface
type awsApiGatewayVpcLinks struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayVpcLinks returns a AwsApiGatewayVpcLinks
func newAwsApiGatewayVpcLinks(c *AwsV1alpha1Client, namespace string) *awsApiGatewayVpcLinks {
	return &awsApiGatewayVpcLinks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayVpcLink, and returns the corresponding awsApiGatewayVpcLink object, and an error if there is any.
func (c *awsApiGatewayVpcLinks) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayVpcLink, err error) {
	result = &v1alpha1.AwsApiGatewayVpcLink{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayvpclinks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayVpcLinks that match those selectors.
func (c *awsApiGatewayVpcLinks) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayVpcLinkList, err error) {
	result = &v1alpha1.AwsApiGatewayVpcLinkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayvpclinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayVpcLinks.
func (c *awsApiGatewayVpcLinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayvpclinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayVpcLink and creates it.  Returns the server's representation of the awsApiGatewayVpcLink, and an error, if there is any.
func (c *awsApiGatewayVpcLinks) Create(awsApiGatewayVpcLink *v1alpha1.AwsApiGatewayVpcLink) (result *v1alpha1.AwsApiGatewayVpcLink, err error) {
	result = &v1alpha1.AwsApiGatewayVpcLink{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewayvpclinks").
		Body(awsApiGatewayVpcLink).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayVpcLink and updates it. Returns the server's representation of the awsApiGatewayVpcLink, and an error, if there is any.
func (c *awsApiGatewayVpcLinks) Update(awsApiGatewayVpcLink *v1alpha1.AwsApiGatewayVpcLink) (result *v1alpha1.AwsApiGatewayVpcLink, err error) {
	result = &v1alpha1.AwsApiGatewayVpcLink{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewayvpclinks").
		Name(awsApiGatewayVpcLink.Name).
		Body(awsApiGatewayVpcLink).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayVpcLink and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayVpcLinks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayvpclinks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayVpcLinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayvpclinks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayVpcLink.
func (c *awsApiGatewayVpcLinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayVpcLink, err error) {
	result = &v1alpha1.AwsApiGatewayVpcLink{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewayvpclinks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
