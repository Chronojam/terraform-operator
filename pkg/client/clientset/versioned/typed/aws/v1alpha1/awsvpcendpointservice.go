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

// AwsVpcEndpointServicesGetter has a method to return a AwsVpcEndpointServiceInterface.
// A group's client should implement this interface.
type AwsVpcEndpointServicesGetter interface {
	AwsVpcEndpointServices(namespace string) AwsVpcEndpointServiceInterface
}

// AwsVpcEndpointServiceInterface has methods to work with AwsVpcEndpointService resources.
type AwsVpcEndpointServiceInterface interface {
	Create(*v1alpha1.AwsVpcEndpointService) (*v1alpha1.AwsVpcEndpointService, error)
	Update(*v1alpha1.AwsVpcEndpointService) (*v1alpha1.AwsVpcEndpointService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsVpcEndpointService, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsVpcEndpointServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcEndpointService, err error)
	AwsVpcEndpointServiceExpansion
}

// awsVpcEndpointServices implements AwsVpcEndpointServiceInterface
type awsVpcEndpointServices struct {
	client rest.Interface
	ns     string
}

// newAwsVpcEndpointServices returns a AwsVpcEndpointServices
func newAwsVpcEndpointServices(c *AwsV1alpha1Client, namespace string) *awsVpcEndpointServices {
	return &awsVpcEndpointServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsVpcEndpointService, and returns the corresponding awsVpcEndpointService object, and an error if there is any.
func (c *awsVpcEndpointServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpcEndpointService, err error) {
	result = &v1alpha1.AwsVpcEndpointService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsvpcendpointservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsVpcEndpointServices that match those selectors.
func (c *awsVpcEndpointServices) List(opts v1.ListOptions) (result *v1alpha1.AwsVpcEndpointServiceList, err error) {
	result = &v1alpha1.AwsVpcEndpointServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsvpcendpointservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsVpcEndpointServices.
func (c *awsVpcEndpointServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsvpcendpointservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsVpcEndpointService and creates it.  Returns the server's representation of the awsVpcEndpointService, and an error, if there is any.
func (c *awsVpcEndpointServices) Create(awsVpcEndpointService *v1alpha1.AwsVpcEndpointService) (result *v1alpha1.AwsVpcEndpointService, err error) {
	result = &v1alpha1.AwsVpcEndpointService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsvpcendpointservices").
		Body(awsVpcEndpointService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsVpcEndpointService and updates it. Returns the server's representation of the awsVpcEndpointService, and an error, if there is any.
func (c *awsVpcEndpointServices) Update(awsVpcEndpointService *v1alpha1.AwsVpcEndpointService) (result *v1alpha1.AwsVpcEndpointService, err error) {
	result = &v1alpha1.AwsVpcEndpointService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsvpcendpointservices").
		Name(awsVpcEndpointService.Name).
		Body(awsVpcEndpointService).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsVpcEndpointService and deletes it. Returns an error if one occurs.
func (c *awsVpcEndpointServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsvpcendpointservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsVpcEndpointServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsvpcendpointservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsVpcEndpointService.
func (c *awsVpcEndpointServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcEndpointService, err error) {
	result = &v1alpha1.AwsVpcEndpointService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsvpcendpointservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
