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

// AwsConfigDeliveryChannelsGetter has a method to return a AwsConfigDeliveryChannelInterface.
// A group's client should implement this interface.
type AwsConfigDeliveryChannelsGetter interface {
	AwsConfigDeliveryChannels(namespace string) AwsConfigDeliveryChannelInterface
}

// AwsConfigDeliveryChannelInterface has methods to work with AwsConfigDeliveryChannel resources.
type AwsConfigDeliveryChannelInterface interface {
	Create(*v1alpha1.AwsConfigDeliveryChannel) (*v1alpha1.AwsConfigDeliveryChannel, error)
	Update(*v1alpha1.AwsConfigDeliveryChannel) (*v1alpha1.AwsConfigDeliveryChannel, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsConfigDeliveryChannel, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsConfigDeliveryChannelList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsConfigDeliveryChannel, err error)
	AwsConfigDeliveryChannelExpansion
}

// awsConfigDeliveryChannels implements AwsConfigDeliveryChannelInterface
type awsConfigDeliveryChannels struct {
	client rest.Interface
	ns     string
}

// newAwsConfigDeliveryChannels returns a AwsConfigDeliveryChannels
func newAwsConfigDeliveryChannels(c *ChronojamV1alpha1Client, namespace string) *awsConfigDeliveryChannels {
	return &awsConfigDeliveryChannels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsConfigDeliveryChannel, and returns the corresponding awsConfigDeliveryChannel object, and an error if there is any.
func (c *awsConfigDeliveryChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsConfigDeliveryChannel, err error) {
	result = &v1alpha1.AwsConfigDeliveryChannel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigdeliverychannels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsConfigDeliveryChannels that match those selectors.
func (c *awsConfigDeliveryChannels) List(opts v1.ListOptions) (result *v1alpha1.AwsConfigDeliveryChannelList, err error) {
	result = &v1alpha1.AwsConfigDeliveryChannelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigdeliverychannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsConfigDeliveryChannels.
func (c *awsConfigDeliveryChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigdeliverychannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsConfigDeliveryChannel and creates it.  Returns the server's representation of the awsConfigDeliveryChannel, and an error, if there is any.
func (c *awsConfigDeliveryChannels) Create(awsConfigDeliveryChannel *v1alpha1.AwsConfigDeliveryChannel) (result *v1alpha1.AwsConfigDeliveryChannel, err error) {
	result = &v1alpha1.AwsConfigDeliveryChannel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsconfigdeliverychannels").
		Body(awsConfigDeliveryChannel).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsConfigDeliveryChannel and updates it. Returns the server's representation of the awsConfigDeliveryChannel, and an error, if there is any.
func (c *awsConfigDeliveryChannels) Update(awsConfigDeliveryChannel *v1alpha1.AwsConfigDeliveryChannel) (result *v1alpha1.AwsConfigDeliveryChannel, err error) {
	result = &v1alpha1.AwsConfigDeliveryChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsconfigdeliverychannels").
		Name(awsConfigDeliveryChannel.Name).
		Body(awsConfigDeliveryChannel).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsConfigDeliveryChannel and deletes it. Returns an error if one occurs.
func (c *awsConfigDeliveryChannels) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsconfigdeliverychannels").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsConfigDeliveryChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsconfigdeliverychannels").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsConfigDeliveryChannel.
func (c *awsConfigDeliveryChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsConfigDeliveryChannel, err error) {
	result = &v1alpha1.AwsConfigDeliveryChannel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsconfigdeliverychannels").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
