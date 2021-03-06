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

// AwsWafRuleGroupsGetter has a method to return a AwsWafRuleGroupInterface.
// A group's client should implement this interface.
type AwsWafRuleGroupsGetter interface {
	AwsWafRuleGroups(namespace string) AwsWafRuleGroupInterface
}

// AwsWafRuleGroupInterface has methods to work with AwsWafRuleGroup resources.
type AwsWafRuleGroupInterface interface {
	Create(*v1alpha1.AwsWafRuleGroup) (*v1alpha1.AwsWafRuleGroup, error)
	Update(*v1alpha1.AwsWafRuleGroup) (*v1alpha1.AwsWafRuleGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafRuleGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafRuleGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafRuleGroup, err error)
	AwsWafRuleGroupExpansion
}

// awsWafRuleGroups implements AwsWafRuleGroupInterface
type awsWafRuleGroups struct {
	client rest.Interface
	ns     string
}

// newAwsWafRuleGroups returns a AwsWafRuleGroups
func newAwsWafRuleGroups(c *ChronojamV1alpha1Client, namespace string) *awsWafRuleGroups {
	return &awsWafRuleGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsWafRuleGroup, and returns the corresponding awsWafRuleGroup object, and an error if there is any.
func (c *awsWafRuleGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafRuleGroup, err error) {
	result = &v1alpha1.AwsWafRuleGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafrulegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafRuleGroups that match those selectors.
func (c *awsWafRuleGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsWafRuleGroupList, err error) {
	result = &v1alpha1.AwsWafRuleGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafRuleGroups.
func (c *awsWafRuleGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awswafrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsWafRuleGroup and creates it.  Returns the server's representation of the awsWafRuleGroup, and an error, if there is any.
func (c *awsWafRuleGroups) Create(awsWafRuleGroup *v1alpha1.AwsWafRuleGroup) (result *v1alpha1.AwsWafRuleGroup, err error) {
	result = &v1alpha1.AwsWafRuleGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awswafrulegroups").
		Body(awsWafRuleGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafRuleGroup and updates it. Returns the server's representation of the awsWafRuleGroup, and an error, if there is any.
func (c *awsWafRuleGroups) Update(awsWafRuleGroup *v1alpha1.AwsWafRuleGroup) (result *v1alpha1.AwsWafRuleGroup, err error) {
	result = &v1alpha1.AwsWafRuleGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awswafrulegroups").
		Name(awsWafRuleGroup.Name).
		Body(awsWafRuleGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafRuleGroup and deletes it. Returns an error if one occurs.
func (c *awsWafRuleGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafrulegroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafRuleGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafrulegroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafRuleGroup.
func (c *awsWafRuleGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafRuleGroup, err error) {
	result = &v1alpha1.AwsWafRuleGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awswafrulegroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
