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

// AwsDynamodbGlobalTablesGetter has a method to return a AwsDynamodbGlobalTableInterface.
// A group's client should implement this interface.
type AwsDynamodbGlobalTablesGetter interface {
	AwsDynamodbGlobalTables(namespace string) AwsDynamodbGlobalTableInterface
}

// AwsDynamodbGlobalTableInterface has methods to work with AwsDynamodbGlobalTable resources.
type AwsDynamodbGlobalTableInterface interface {
	Create(*v1alpha1.AwsDynamodbGlobalTable) (*v1alpha1.AwsDynamodbGlobalTable, error)
	Update(*v1alpha1.AwsDynamodbGlobalTable) (*v1alpha1.AwsDynamodbGlobalTable, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDynamodbGlobalTable, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDynamodbGlobalTableList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDynamodbGlobalTable, err error)
	AwsDynamodbGlobalTableExpansion
}

// awsDynamodbGlobalTables implements AwsDynamodbGlobalTableInterface
type awsDynamodbGlobalTables struct {
	client rest.Interface
	ns     string
}

// newAwsDynamodbGlobalTables returns a AwsDynamodbGlobalTables
func newAwsDynamodbGlobalTables(c *AwsV1alpha1Client, namespace string) *awsDynamodbGlobalTables {
	return &awsDynamodbGlobalTables{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDynamodbGlobalTable, and returns the corresponding awsDynamodbGlobalTable object, and an error if there is any.
func (c *awsDynamodbGlobalTables) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDynamodbGlobalTable, err error) {
	result = &v1alpha1.AwsDynamodbGlobalTable{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdynamodbglobaltables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDynamodbGlobalTables that match those selectors.
func (c *awsDynamodbGlobalTables) List(opts v1.ListOptions) (result *v1alpha1.AwsDynamodbGlobalTableList, err error) {
	result = &v1alpha1.AwsDynamodbGlobalTableList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdynamodbglobaltables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDynamodbGlobalTables.
func (c *awsDynamodbGlobalTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdynamodbglobaltables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDynamodbGlobalTable and creates it.  Returns the server's representation of the awsDynamodbGlobalTable, and an error, if there is any.
func (c *awsDynamodbGlobalTables) Create(awsDynamodbGlobalTable *v1alpha1.AwsDynamodbGlobalTable) (result *v1alpha1.AwsDynamodbGlobalTable, err error) {
	result = &v1alpha1.AwsDynamodbGlobalTable{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdynamodbglobaltables").
		Body(awsDynamodbGlobalTable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDynamodbGlobalTable and updates it. Returns the server's representation of the awsDynamodbGlobalTable, and an error, if there is any.
func (c *awsDynamodbGlobalTables) Update(awsDynamodbGlobalTable *v1alpha1.AwsDynamodbGlobalTable) (result *v1alpha1.AwsDynamodbGlobalTable, err error) {
	result = &v1alpha1.AwsDynamodbGlobalTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdynamodbglobaltables").
		Name(awsDynamodbGlobalTable.Name).
		Body(awsDynamodbGlobalTable).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDynamodbGlobalTable and deletes it. Returns an error if one occurs.
func (c *awsDynamodbGlobalTables) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdynamodbglobaltables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDynamodbGlobalTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdynamodbglobaltables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDynamodbGlobalTable.
func (c *awsDynamodbGlobalTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDynamodbGlobalTable, err error) {
	result = &v1alpha1.AwsDynamodbGlobalTable{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdynamodbglobaltables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
