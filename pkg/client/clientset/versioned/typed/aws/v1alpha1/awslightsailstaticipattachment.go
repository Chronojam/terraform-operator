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

// AwsLightsailStaticIpAttachmentsGetter has a method to return a AwsLightsailStaticIpAttachmentInterface.
// A group's client should implement this interface.
type AwsLightsailStaticIpAttachmentsGetter interface {
	AwsLightsailStaticIpAttachments(namespace string) AwsLightsailStaticIpAttachmentInterface
}

// AwsLightsailStaticIpAttachmentInterface has methods to work with AwsLightsailStaticIpAttachment resources.
type AwsLightsailStaticIpAttachmentInterface interface {
	Create(*v1alpha1.AwsLightsailStaticIpAttachment) (*v1alpha1.AwsLightsailStaticIpAttachment, error)
	Update(*v1alpha1.AwsLightsailStaticIpAttachment) (*v1alpha1.AwsLightsailStaticIpAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsLightsailStaticIpAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsLightsailStaticIpAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLightsailStaticIpAttachment, err error)
	AwsLightsailStaticIpAttachmentExpansion
}

// awsLightsailStaticIpAttachments implements AwsLightsailStaticIpAttachmentInterface
type awsLightsailStaticIpAttachments struct {
	client rest.Interface
	ns     string
}

// newAwsLightsailStaticIpAttachments returns a AwsLightsailStaticIpAttachments
func newAwsLightsailStaticIpAttachments(c *ChronojamV1alpha1Client, namespace string) *awsLightsailStaticIpAttachments {
	return &awsLightsailStaticIpAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsLightsailStaticIpAttachment, and returns the corresponding awsLightsailStaticIpAttachment object, and an error if there is any.
func (c *awsLightsailStaticIpAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLightsailStaticIpAttachment, err error) {
	result = &v1alpha1.AwsLightsailStaticIpAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslightsailstaticipattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLightsailStaticIpAttachments that match those selectors.
func (c *awsLightsailStaticIpAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsLightsailStaticIpAttachmentList, err error) {
	result = &v1alpha1.AwsLightsailStaticIpAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslightsailstaticipattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLightsailStaticIpAttachments.
func (c *awsLightsailStaticIpAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awslightsailstaticipattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsLightsailStaticIpAttachment and creates it.  Returns the server's representation of the awsLightsailStaticIpAttachment, and an error, if there is any.
func (c *awsLightsailStaticIpAttachments) Create(awsLightsailStaticIpAttachment *v1alpha1.AwsLightsailStaticIpAttachment) (result *v1alpha1.AwsLightsailStaticIpAttachment, err error) {
	result = &v1alpha1.AwsLightsailStaticIpAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awslightsailstaticipattachments").
		Body(awsLightsailStaticIpAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLightsailStaticIpAttachment and updates it. Returns the server's representation of the awsLightsailStaticIpAttachment, and an error, if there is any.
func (c *awsLightsailStaticIpAttachments) Update(awsLightsailStaticIpAttachment *v1alpha1.AwsLightsailStaticIpAttachment) (result *v1alpha1.AwsLightsailStaticIpAttachment, err error) {
	result = &v1alpha1.AwsLightsailStaticIpAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awslightsailstaticipattachments").
		Name(awsLightsailStaticIpAttachment.Name).
		Body(awsLightsailStaticIpAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLightsailStaticIpAttachment and deletes it. Returns an error if one occurs.
func (c *awsLightsailStaticIpAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslightsailstaticipattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLightsailStaticIpAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslightsailstaticipattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLightsailStaticIpAttachment.
func (c *awsLightsailStaticIpAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLightsailStaticIpAttachment, err error) {
	result = &v1alpha1.AwsLightsailStaticIpAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awslightsailstaticipattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
