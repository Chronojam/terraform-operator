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

// AwsApiGatewayUsagePlansGetter has a method to return a AwsApiGatewayUsagePlanInterface.
// A group's client should implement this interface.
type AwsApiGatewayUsagePlansGetter interface {
	AwsApiGatewayUsagePlans(namespace string) AwsApiGatewayUsagePlanInterface
}

// AwsApiGatewayUsagePlanInterface has methods to work with AwsApiGatewayUsagePlan resources.
type AwsApiGatewayUsagePlanInterface interface {
	Create(*v1alpha1.AwsApiGatewayUsagePlan) (*v1alpha1.AwsApiGatewayUsagePlan, error)
	Update(*v1alpha1.AwsApiGatewayUsagePlan) (*v1alpha1.AwsApiGatewayUsagePlan, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayUsagePlan, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayUsagePlanList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayUsagePlan, err error)
	AwsApiGatewayUsagePlanExpansion
}

// awsApiGatewayUsagePlans implements AwsApiGatewayUsagePlanInterface
type awsApiGatewayUsagePlans struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayUsagePlans returns a AwsApiGatewayUsagePlans
func newAwsApiGatewayUsagePlans(c *ChronojamV1alpha1Client, namespace string) *awsApiGatewayUsagePlans {
	return &awsApiGatewayUsagePlans{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayUsagePlan, and returns the corresponding awsApiGatewayUsagePlan object, and an error if there is any.
func (c *awsApiGatewayUsagePlans) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayUsagePlan, err error) {
	result = &v1alpha1.AwsApiGatewayUsagePlan{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayusageplans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayUsagePlans that match those selectors.
func (c *awsApiGatewayUsagePlans) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayUsagePlanList, err error) {
	result = &v1alpha1.AwsApiGatewayUsagePlanList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayusageplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayUsagePlans.
func (c *awsApiGatewayUsagePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayusageplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayUsagePlan and creates it.  Returns the server's representation of the awsApiGatewayUsagePlan, and an error, if there is any.
func (c *awsApiGatewayUsagePlans) Create(awsApiGatewayUsagePlan *v1alpha1.AwsApiGatewayUsagePlan) (result *v1alpha1.AwsApiGatewayUsagePlan, err error) {
	result = &v1alpha1.AwsApiGatewayUsagePlan{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewayusageplans").
		Body(awsApiGatewayUsagePlan).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayUsagePlan and updates it. Returns the server's representation of the awsApiGatewayUsagePlan, and an error, if there is any.
func (c *awsApiGatewayUsagePlans) Update(awsApiGatewayUsagePlan *v1alpha1.AwsApiGatewayUsagePlan) (result *v1alpha1.AwsApiGatewayUsagePlan, err error) {
	result = &v1alpha1.AwsApiGatewayUsagePlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewayusageplans").
		Name(awsApiGatewayUsagePlan.Name).
		Body(awsApiGatewayUsagePlan).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayUsagePlan and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayUsagePlans) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayusageplans").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayUsagePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayusageplans").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayUsagePlan.
func (c *awsApiGatewayUsagePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayUsagePlan, err error) {
	result = &v1alpha1.AwsApiGatewayUsagePlan{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewayusageplans").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
