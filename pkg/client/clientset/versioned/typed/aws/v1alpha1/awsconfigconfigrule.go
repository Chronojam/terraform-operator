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

// AwsConfigConfigRulesGetter has a method to return a AwsConfigConfigRuleInterface.
// A group's client should implement this interface.
type AwsConfigConfigRulesGetter interface {
	AwsConfigConfigRules(namespace string) AwsConfigConfigRuleInterface
}

// AwsConfigConfigRuleInterface has methods to work with AwsConfigConfigRule resources.
type AwsConfigConfigRuleInterface interface {
	Create(*v1alpha1.AwsConfigConfigRule) (*v1alpha1.AwsConfigConfigRule, error)
	Update(*v1alpha1.AwsConfigConfigRule) (*v1alpha1.AwsConfigConfigRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsConfigConfigRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsConfigConfigRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsConfigConfigRule, err error)
	AwsConfigConfigRuleExpansion
}

// awsConfigConfigRules implements AwsConfigConfigRuleInterface
type awsConfigConfigRules struct {
	client rest.Interface
	ns     string
}

// newAwsConfigConfigRules returns a AwsConfigConfigRules
func newAwsConfigConfigRules(c *AwsV1alpha1Client, namespace string) *awsConfigConfigRules {
	return &awsConfigConfigRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsConfigConfigRule, and returns the corresponding awsConfigConfigRule object, and an error if there is any.
func (c *awsConfigConfigRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsConfigConfigRule, err error) {
	result = &v1alpha1.AwsConfigConfigRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigconfigrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsConfigConfigRules that match those selectors.
func (c *awsConfigConfigRules) List(opts v1.ListOptions) (result *v1alpha1.AwsConfigConfigRuleList, err error) {
	result = &v1alpha1.AwsConfigConfigRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigconfigrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsConfigConfigRules.
func (c *awsConfigConfigRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigconfigrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsConfigConfigRule and creates it.  Returns the server's representation of the awsConfigConfigRule, and an error, if there is any.
func (c *awsConfigConfigRules) Create(awsConfigConfigRule *v1alpha1.AwsConfigConfigRule) (result *v1alpha1.AwsConfigConfigRule, err error) {
	result = &v1alpha1.AwsConfigConfigRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsconfigconfigrules").
		Body(awsConfigConfigRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsConfigConfigRule and updates it. Returns the server's representation of the awsConfigConfigRule, and an error, if there is any.
func (c *awsConfigConfigRules) Update(awsConfigConfigRule *v1alpha1.AwsConfigConfigRule) (result *v1alpha1.AwsConfigConfigRule, err error) {
	result = &v1alpha1.AwsConfigConfigRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsconfigconfigrules").
		Name(awsConfigConfigRule.Name).
		Body(awsConfigConfigRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsConfigConfigRule and deletes it. Returns an error if one occurs.
func (c *awsConfigConfigRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsconfigconfigrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsConfigConfigRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsconfigconfigrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsConfigConfigRule.
func (c *awsConfigConfigRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsConfigConfigRule, err error) {
	result = &v1alpha1.AwsConfigConfigRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsconfigconfigrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
