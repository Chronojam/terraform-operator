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

// AwsWafregionalRuleGroupsGetter has a method to return a AwsWafregionalRuleGroupInterface.
// A group's client should implement this interface.
type AwsWafregionalRuleGroupsGetter interface {
	AwsWafregionalRuleGroups(namespace string) AwsWafregionalRuleGroupInterface
}

// AwsWafregionalRuleGroupInterface has methods to work with AwsWafregionalRuleGroup resources.
type AwsWafregionalRuleGroupInterface interface {
	Create(*v1alpha1.AwsWafregionalRuleGroup) (*v1alpha1.AwsWafregionalRuleGroup, error)
	Update(*v1alpha1.AwsWafregionalRuleGroup) (*v1alpha1.AwsWafregionalRuleGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafregionalRuleGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafregionalRuleGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalRuleGroup, err error)
	AwsWafregionalRuleGroupExpansion
}

// awsWafregionalRuleGroups implements AwsWafregionalRuleGroupInterface
type awsWafregionalRuleGroups struct {
	client rest.Interface
	ns     string
}

// newAwsWafregionalRuleGroups returns a AwsWafregionalRuleGroups
func newAwsWafregionalRuleGroups(c *ChronojamV1alpha1Client, namespace string) *awsWafregionalRuleGroups {
	return &awsWafregionalRuleGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsWafregionalRuleGroup, and returns the corresponding awsWafregionalRuleGroup object, and an error if there is any.
func (c *awsWafregionalRuleGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafregionalRuleGroup, err error) {
	result = &v1alpha1.AwsWafregionalRuleGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafregionalrulegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafregionalRuleGroups that match those selectors.
func (c *awsWafregionalRuleGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsWafregionalRuleGroupList, err error) {
	result = &v1alpha1.AwsWafregionalRuleGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafregionalrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafregionalRuleGroups.
func (c *awsWafregionalRuleGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awswafregionalrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsWafregionalRuleGroup and creates it.  Returns the server's representation of the awsWafregionalRuleGroup, and an error, if there is any.
func (c *awsWafregionalRuleGroups) Create(awsWafregionalRuleGroup *v1alpha1.AwsWafregionalRuleGroup) (result *v1alpha1.AwsWafregionalRuleGroup, err error) {
	result = &v1alpha1.AwsWafregionalRuleGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awswafregionalrulegroups").
		Body(awsWafregionalRuleGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafregionalRuleGroup and updates it. Returns the server's representation of the awsWafregionalRuleGroup, and an error, if there is any.
func (c *awsWafregionalRuleGroups) Update(awsWafregionalRuleGroup *v1alpha1.AwsWafregionalRuleGroup) (result *v1alpha1.AwsWafregionalRuleGroup, err error) {
	result = &v1alpha1.AwsWafregionalRuleGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awswafregionalrulegroups").
		Name(awsWafregionalRuleGroup.Name).
		Body(awsWafregionalRuleGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafregionalRuleGroup and deletes it. Returns an error if one occurs.
func (c *awsWafregionalRuleGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafregionalrulegroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafregionalRuleGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafregionalrulegroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafregionalRuleGroup.
func (c *awsWafregionalRuleGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalRuleGroup, err error) {
	result = &v1alpha1.AwsWafregionalRuleGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awswafregionalrulegroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
