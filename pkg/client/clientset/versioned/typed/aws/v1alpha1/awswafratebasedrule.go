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

// AwsWafRateBasedRulesGetter has a method to return a AwsWafRateBasedRuleInterface.
// A group's client should implement this interface.
type AwsWafRateBasedRulesGetter interface {
	AwsWafRateBasedRules(namespace string) AwsWafRateBasedRuleInterface
}

// AwsWafRateBasedRuleInterface has methods to work with AwsWafRateBasedRule resources.
type AwsWafRateBasedRuleInterface interface {
	Create(*v1alpha1.AwsWafRateBasedRule) (*v1alpha1.AwsWafRateBasedRule, error)
	Update(*v1alpha1.AwsWafRateBasedRule) (*v1alpha1.AwsWafRateBasedRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafRateBasedRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafRateBasedRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafRateBasedRule, err error)
	AwsWafRateBasedRuleExpansion
}

// awsWafRateBasedRules implements AwsWafRateBasedRuleInterface
type awsWafRateBasedRules struct {
	client rest.Interface
	ns     string
}

// newAwsWafRateBasedRules returns a AwsWafRateBasedRules
func newAwsWafRateBasedRules(c *AwsV1alpha1Client, namespace string) *awsWafRateBasedRules {
	return &awsWafRateBasedRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsWafRateBasedRule, and returns the corresponding awsWafRateBasedRule object, and an error if there is any.
func (c *awsWafRateBasedRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafRateBasedRule, err error) {
	result = &v1alpha1.AwsWafRateBasedRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafratebasedrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafRateBasedRules that match those selectors.
func (c *awsWafRateBasedRules) List(opts v1.ListOptions) (result *v1alpha1.AwsWafRateBasedRuleList, err error) {
	result = &v1alpha1.AwsWafRateBasedRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafratebasedrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafRateBasedRules.
func (c *awsWafRateBasedRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awswafratebasedrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsWafRateBasedRule and creates it.  Returns the server's representation of the awsWafRateBasedRule, and an error, if there is any.
func (c *awsWafRateBasedRules) Create(awsWafRateBasedRule *v1alpha1.AwsWafRateBasedRule) (result *v1alpha1.AwsWafRateBasedRule, err error) {
	result = &v1alpha1.AwsWafRateBasedRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awswafratebasedrules").
		Body(awsWafRateBasedRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafRateBasedRule and updates it. Returns the server's representation of the awsWafRateBasedRule, and an error, if there is any.
func (c *awsWafRateBasedRules) Update(awsWafRateBasedRule *v1alpha1.AwsWafRateBasedRule) (result *v1alpha1.AwsWafRateBasedRule, err error) {
	result = &v1alpha1.AwsWafRateBasedRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awswafratebasedrules").
		Name(awsWafRateBasedRule.Name).
		Body(awsWafRateBasedRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafRateBasedRule and deletes it. Returns an error if one occurs.
func (c *awsWafRateBasedRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafratebasedrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafRateBasedRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafratebasedrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafRateBasedRule.
func (c *awsWafRateBasedRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafRateBasedRule, err error) {
	result = &v1alpha1.AwsWafRateBasedRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awswafratebasedrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
