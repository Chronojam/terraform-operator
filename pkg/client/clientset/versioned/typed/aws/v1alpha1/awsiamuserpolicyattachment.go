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

// AwsIamUserPolicyAttachmentsGetter has a method to return a AwsIamUserPolicyAttachmentInterface.
// A group's client should implement this interface.
type AwsIamUserPolicyAttachmentsGetter interface {
	AwsIamUserPolicyAttachments(namespace string) AwsIamUserPolicyAttachmentInterface
}

// AwsIamUserPolicyAttachmentInterface has methods to work with AwsIamUserPolicyAttachment resources.
type AwsIamUserPolicyAttachmentInterface interface {
	Create(*v1alpha1.AwsIamUserPolicyAttachment) (*v1alpha1.AwsIamUserPolicyAttachment, error)
	Update(*v1alpha1.AwsIamUserPolicyAttachment) (*v1alpha1.AwsIamUserPolicyAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsIamUserPolicyAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsIamUserPolicyAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamUserPolicyAttachment, err error)
	AwsIamUserPolicyAttachmentExpansion
}

// awsIamUserPolicyAttachments implements AwsIamUserPolicyAttachmentInterface
type awsIamUserPolicyAttachments struct {
	client rest.Interface
	ns     string
}

// newAwsIamUserPolicyAttachments returns a AwsIamUserPolicyAttachments
func newAwsIamUserPolicyAttachments(c *ChronojamV1alpha1Client, namespace string) *awsIamUserPolicyAttachments {
	return &awsIamUserPolicyAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsIamUserPolicyAttachment, and returns the corresponding awsIamUserPolicyAttachment object, and an error if there is any.
func (c *awsIamUserPolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamUserPolicyAttachment, err error) {
	result = &v1alpha1.AwsIamUserPolicyAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamuserpolicyattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIamUserPolicyAttachments that match those selectors.
func (c *awsIamUserPolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsIamUserPolicyAttachmentList, err error) {
	result = &v1alpha1.AwsIamUserPolicyAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamuserpolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIamUserPolicyAttachments.
func (c *awsIamUserPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiamuserpolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsIamUserPolicyAttachment and creates it.  Returns the server's representation of the awsIamUserPolicyAttachment, and an error, if there is any.
func (c *awsIamUserPolicyAttachments) Create(awsIamUserPolicyAttachment *v1alpha1.AwsIamUserPolicyAttachment) (result *v1alpha1.AwsIamUserPolicyAttachment, err error) {
	result = &v1alpha1.AwsIamUserPolicyAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiamuserpolicyattachments").
		Body(awsIamUserPolicyAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIamUserPolicyAttachment and updates it. Returns the server's representation of the awsIamUserPolicyAttachment, and an error, if there is any.
func (c *awsIamUserPolicyAttachments) Update(awsIamUserPolicyAttachment *v1alpha1.AwsIamUserPolicyAttachment) (result *v1alpha1.AwsIamUserPolicyAttachment, err error) {
	result = &v1alpha1.AwsIamUserPolicyAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiamuserpolicyattachments").
		Name(awsIamUserPolicyAttachment.Name).
		Body(awsIamUserPolicyAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIamUserPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *awsIamUserPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamuserpolicyattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIamUserPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamuserpolicyattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIamUserPolicyAttachment.
func (c *awsIamUserPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamUserPolicyAttachment, err error) {
	result = &v1alpha1.AwsIamUserPolicyAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiamuserpolicyattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
