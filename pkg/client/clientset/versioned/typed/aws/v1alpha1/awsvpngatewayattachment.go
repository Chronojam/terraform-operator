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

// AwsVpnGatewayAttachmentsGetter has a method to return a AwsVpnGatewayAttachmentInterface.
// A group's client should implement this interface.
type AwsVpnGatewayAttachmentsGetter interface {
	AwsVpnGatewayAttachments(namespace string) AwsVpnGatewayAttachmentInterface
}

// AwsVpnGatewayAttachmentInterface has methods to work with AwsVpnGatewayAttachment resources.
type AwsVpnGatewayAttachmentInterface interface {
	Create(*v1alpha1.AwsVpnGatewayAttachment) (*v1alpha1.AwsVpnGatewayAttachment, error)
	Update(*v1alpha1.AwsVpnGatewayAttachment) (*v1alpha1.AwsVpnGatewayAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsVpnGatewayAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsVpnGatewayAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpnGatewayAttachment, err error)
	AwsVpnGatewayAttachmentExpansion
}

// awsVpnGatewayAttachments implements AwsVpnGatewayAttachmentInterface
type awsVpnGatewayAttachments struct {
	client rest.Interface
	ns     string
}

// newAwsVpnGatewayAttachments returns a AwsVpnGatewayAttachments
func newAwsVpnGatewayAttachments(c *ChronojamV1alpha1Client, namespace string) *awsVpnGatewayAttachments {
	return &awsVpnGatewayAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsVpnGatewayAttachment, and returns the corresponding awsVpnGatewayAttachment object, and an error if there is any.
func (c *awsVpnGatewayAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpnGatewayAttachment, err error) {
	result = &v1alpha1.AwsVpnGatewayAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsvpngatewayattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsVpnGatewayAttachments that match those selectors.
func (c *awsVpnGatewayAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsVpnGatewayAttachmentList, err error) {
	result = &v1alpha1.AwsVpnGatewayAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsvpngatewayattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsVpnGatewayAttachments.
func (c *awsVpnGatewayAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsvpngatewayattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsVpnGatewayAttachment and creates it.  Returns the server's representation of the awsVpnGatewayAttachment, and an error, if there is any.
func (c *awsVpnGatewayAttachments) Create(awsVpnGatewayAttachment *v1alpha1.AwsVpnGatewayAttachment) (result *v1alpha1.AwsVpnGatewayAttachment, err error) {
	result = &v1alpha1.AwsVpnGatewayAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsvpngatewayattachments").
		Body(awsVpnGatewayAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsVpnGatewayAttachment and updates it. Returns the server's representation of the awsVpnGatewayAttachment, and an error, if there is any.
func (c *awsVpnGatewayAttachments) Update(awsVpnGatewayAttachment *v1alpha1.AwsVpnGatewayAttachment) (result *v1alpha1.AwsVpnGatewayAttachment, err error) {
	result = &v1alpha1.AwsVpnGatewayAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsvpngatewayattachments").
		Name(awsVpnGatewayAttachment.Name).
		Body(awsVpnGatewayAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsVpnGatewayAttachment and deletes it. Returns an error if one occurs.
func (c *awsVpnGatewayAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsvpngatewayattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsVpnGatewayAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsvpngatewayattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsVpnGatewayAttachment.
func (c *awsVpnGatewayAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpnGatewayAttachment, err error) {
	result = &v1alpha1.AwsVpnGatewayAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsvpngatewayattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
