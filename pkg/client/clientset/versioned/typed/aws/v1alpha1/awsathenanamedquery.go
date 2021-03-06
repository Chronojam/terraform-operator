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

// AwsAthenaNamedQueriesGetter has a method to return a AwsAthenaNamedQueryInterface.
// A group's client should implement this interface.
type AwsAthenaNamedQueriesGetter interface {
	AwsAthenaNamedQueries(namespace string) AwsAthenaNamedQueryInterface
}

// AwsAthenaNamedQueryInterface has methods to work with AwsAthenaNamedQuery resources.
type AwsAthenaNamedQueryInterface interface {
	Create(*v1alpha1.AwsAthenaNamedQuery) (*v1alpha1.AwsAthenaNamedQuery, error)
	Update(*v1alpha1.AwsAthenaNamedQuery) (*v1alpha1.AwsAthenaNamedQuery, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAthenaNamedQuery, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAthenaNamedQueryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAthenaNamedQuery, err error)
	AwsAthenaNamedQueryExpansion
}

// awsAthenaNamedQueries implements AwsAthenaNamedQueryInterface
type awsAthenaNamedQueries struct {
	client rest.Interface
	ns     string
}

// newAwsAthenaNamedQueries returns a AwsAthenaNamedQueries
func newAwsAthenaNamedQueries(c *ChronojamV1alpha1Client, namespace string) *awsAthenaNamedQueries {
	return &awsAthenaNamedQueries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsAthenaNamedQuery, and returns the corresponding awsAthenaNamedQuery object, and an error if there is any.
func (c *awsAthenaNamedQueries) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAthenaNamedQuery, err error) {
	result = &v1alpha1.AwsAthenaNamedQuery{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsathenanamedqueries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAthenaNamedQueries that match those selectors.
func (c *awsAthenaNamedQueries) List(opts v1.ListOptions) (result *v1alpha1.AwsAthenaNamedQueryList, err error) {
	result = &v1alpha1.AwsAthenaNamedQueryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsathenanamedqueries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAthenaNamedQueries.
func (c *awsAthenaNamedQueries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsathenanamedqueries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsAthenaNamedQuery and creates it.  Returns the server's representation of the awsAthenaNamedQuery, and an error, if there is any.
func (c *awsAthenaNamedQueries) Create(awsAthenaNamedQuery *v1alpha1.AwsAthenaNamedQuery) (result *v1alpha1.AwsAthenaNamedQuery, err error) {
	result = &v1alpha1.AwsAthenaNamedQuery{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsathenanamedqueries").
		Body(awsAthenaNamedQuery).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAthenaNamedQuery and updates it. Returns the server's representation of the awsAthenaNamedQuery, and an error, if there is any.
func (c *awsAthenaNamedQueries) Update(awsAthenaNamedQuery *v1alpha1.AwsAthenaNamedQuery) (result *v1alpha1.AwsAthenaNamedQuery, err error) {
	result = &v1alpha1.AwsAthenaNamedQuery{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsathenanamedqueries").
		Name(awsAthenaNamedQuery.Name).
		Body(awsAthenaNamedQuery).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAthenaNamedQuery and deletes it. Returns an error if one occurs.
func (c *awsAthenaNamedQueries) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsathenanamedqueries").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAthenaNamedQueries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsathenanamedqueries").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAthenaNamedQuery.
func (c *awsAthenaNamedQueries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAthenaNamedQuery, err error) {
	result = &v1alpha1.AwsAthenaNamedQuery{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsathenanamedqueries").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
