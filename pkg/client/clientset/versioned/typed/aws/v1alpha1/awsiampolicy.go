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

// AwsIamPoliciesGetter has a method to return a AwsIamPolicyInterface.
// A group's client should implement this interface.
type AwsIamPoliciesGetter interface {
	AwsIamPolicies(namespace string) AwsIamPolicyInterface
}

// AwsIamPolicyInterface has methods to work with AwsIamPolicy resources.
type AwsIamPolicyInterface interface {
	Create(*v1alpha1.AwsIamPolicy) (*v1alpha1.AwsIamPolicy, error)
	Update(*v1alpha1.AwsIamPolicy) (*v1alpha1.AwsIamPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsIamPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsIamPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamPolicy, err error)
	AwsIamPolicyExpansion
}

// awsIamPolicies implements AwsIamPolicyInterface
type awsIamPolicies struct {
	client rest.Interface
	ns     string
}

// newAwsIamPolicies returns a AwsIamPolicies
func newAwsIamPolicies(c *ChronojamV1alpha1Client, namespace string) *awsIamPolicies {
	return &awsIamPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsIamPolicy, and returns the corresponding awsIamPolicy object, and an error if there is any.
func (c *awsIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamPolicy, err error) {
	result = &v1alpha1.AwsIamPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIamPolicies that match those selectors.
func (c *awsIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsIamPolicyList, err error) {
	result = &v1alpha1.AwsIamPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIamPolicies.
func (c *awsIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsIamPolicy and creates it.  Returns the server's representation of the awsIamPolicy, and an error, if there is any.
func (c *awsIamPolicies) Create(awsIamPolicy *v1alpha1.AwsIamPolicy) (result *v1alpha1.AwsIamPolicy, err error) {
	result = &v1alpha1.AwsIamPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiampolicies").
		Body(awsIamPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIamPolicy and updates it. Returns the server's representation of the awsIamPolicy, and an error, if there is any.
func (c *awsIamPolicies) Update(awsIamPolicy *v1alpha1.AwsIamPolicy) (result *v1alpha1.AwsIamPolicy, err error) {
	result = &v1alpha1.AwsIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiampolicies").
		Name(awsIamPolicy.Name).
		Body(awsIamPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIamPolicy and deletes it. Returns an error if one occurs.
func (c *awsIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiampolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiampolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIamPolicy.
func (c *awsIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamPolicy, err error) {
	result = &v1alpha1.AwsIamPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiampolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
