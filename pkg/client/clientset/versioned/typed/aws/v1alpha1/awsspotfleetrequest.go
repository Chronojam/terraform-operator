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

// AwsSpotFleetRequestsGetter has a method to return a AwsSpotFleetRequestInterface.
// A group's client should implement this interface.
type AwsSpotFleetRequestsGetter interface {
	AwsSpotFleetRequests(namespace string) AwsSpotFleetRequestInterface
}

// AwsSpotFleetRequestInterface has methods to work with AwsSpotFleetRequest resources.
type AwsSpotFleetRequestInterface interface {
	Create(*v1alpha1.AwsSpotFleetRequest) (*v1alpha1.AwsSpotFleetRequest, error)
	Update(*v1alpha1.AwsSpotFleetRequest) (*v1alpha1.AwsSpotFleetRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSpotFleetRequest, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSpotFleetRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSpotFleetRequest, err error)
	AwsSpotFleetRequestExpansion
}

// awsSpotFleetRequests implements AwsSpotFleetRequestInterface
type awsSpotFleetRequests struct {
	client rest.Interface
	ns     string
}

// newAwsSpotFleetRequests returns a AwsSpotFleetRequests
func newAwsSpotFleetRequests(c *ChronojamV1alpha1Client, namespace string) *awsSpotFleetRequests {
	return &awsSpotFleetRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSpotFleetRequest, and returns the corresponding awsSpotFleetRequest object, and an error if there is any.
func (c *awsSpotFleetRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSpotFleetRequest, err error) {
	result = &v1alpha1.AwsSpotFleetRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsspotfleetrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSpotFleetRequests that match those selectors.
func (c *awsSpotFleetRequests) List(opts v1.ListOptions) (result *v1alpha1.AwsSpotFleetRequestList, err error) {
	result = &v1alpha1.AwsSpotFleetRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsspotfleetrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSpotFleetRequests.
func (c *awsSpotFleetRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsspotfleetrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSpotFleetRequest and creates it.  Returns the server's representation of the awsSpotFleetRequest, and an error, if there is any.
func (c *awsSpotFleetRequests) Create(awsSpotFleetRequest *v1alpha1.AwsSpotFleetRequest) (result *v1alpha1.AwsSpotFleetRequest, err error) {
	result = &v1alpha1.AwsSpotFleetRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsspotfleetrequests").
		Body(awsSpotFleetRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSpotFleetRequest and updates it. Returns the server's representation of the awsSpotFleetRequest, and an error, if there is any.
func (c *awsSpotFleetRequests) Update(awsSpotFleetRequest *v1alpha1.AwsSpotFleetRequest) (result *v1alpha1.AwsSpotFleetRequest, err error) {
	result = &v1alpha1.AwsSpotFleetRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsspotfleetrequests").
		Name(awsSpotFleetRequest.Name).
		Body(awsSpotFleetRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSpotFleetRequest and deletes it. Returns an error if one occurs.
func (c *awsSpotFleetRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsspotfleetrequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSpotFleetRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsspotfleetrequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSpotFleetRequest.
func (c *awsSpotFleetRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSpotFleetRequest, err error) {
	result = &v1alpha1.AwsSpotFleetRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsspotfleetrequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
