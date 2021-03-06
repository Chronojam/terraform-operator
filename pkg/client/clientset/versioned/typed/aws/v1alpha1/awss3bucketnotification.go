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

// AwsS3BucketNotificationsGetter has a method to return a AwsS3BucketNotificationInterface.
// A group's client should implement this interface.
type AwsS3BucketNotificationsGetter interface {
	AwsS3BucketNotifications(namespace string) AwsS3BucketNotificationInterface
}

// AwsS3BucketNotificationInterface has methods to work with AwsS3BucketNotification resources.
type AwsS3BucketNotificationInterface interface {
	Create(*v1alpha1.AwsS3BucketNotification) (*v1alpha1.AwsS3BucketNotification, error)
	Update(*v1alpha1.AwsS3BucketNotification) (*v1alpha1.AwsS3BucketNotification, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsS3BucketNotification, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsS3BucketNotificationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsS3BucketNotification, err error)
	AwsS3BucketNotificationExpansion
}

// awsS3BucketNotifications implements AwsS3BucketNotificationInterface
type awsS3BucketNotifications struct {
	client rest.Interface
	ns     string
}

// newAwsS3BucketNotifications returns a AwsS3BucketNotifications
func newAwsS3BucketNotifications(c *ChronojamV1alpha1Client, namespace string) *awsS3BucketNotifications {
	return &awsS3BucketNotifications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsS3BucketNotification, and returns the corresponding awsS3BucketNotification object, and an error if there is any.
func (c *awsS3BucketNotifications) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsS3BucketNotification, err error) {
	result = &v1alpha1.AwsS3BucketNotification{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awss3bucketnotifications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsS3BucketNotifications that match those selectors.
func (c *awsS3BucketNotifications) List(opts v1.ListOptions) (result *v1alpha1.AwsS3BucketNotificationList, err error) {
	result = &v1alpha1.AwsS3BucketNotificationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awss3bucketnotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsS3BucketNotifications.
func (c *awsS3BucketNotifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awss3bucketnotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsS3BucketNotification and creates it.  Returns the server's representation of the awsS3BucketNotification, and an error, if there is any.
func (c *awsS3BucketNotifications) Create(awsS3BucketNotification *v1alpha1.AwsS3BucketNotification) (result *v1alpha1.AwsS3BucketNotification, err error) {
	result = &v1alpha1.AwsS3BucketNotification{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awss3bucketnotifications").
		Body(awsS3BucketNotification).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsS3BucketNotification and updates it. Returns the server's representation of the awsS3BucketNotification, and an error, if there is any.
func (c *awsS3BucketNotifications) Update(awsS3BucketNotification *v1alpha1.AwsS3BucketNotification) (result *v1alpha1.AwsS3BucketNotification, err error) {
	result = &v1alpha1.AwsS3BucketNotification{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awss3bucketnotifications").
		Name(awsS3BucketNotification.Name).
		Body(awsS3BucketNotification).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsS3BucketNotification and deletes it. Returns an error if one occurs.
func (c *awsS3BucketNotifications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awss3bucketnotifications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsS3BucketNotifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awss3bucketnotifications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsS3BucketNotification.
func (c *awsS3BucketNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsS3BucketNotification, err error) {
	result = &v1alpha1.AwsS3BucketNotification{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awss3bucketnotifications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
