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

// AwsCloudwatchEventRulesGetter has a method to return a AwsCloudwatchEventRuleInterface.
// A group's client should implement this interface.
type AwsCloudwatchEventRulesGetter interface {
	AwsCloudwatchEventRules(namespace string) AwsCloudwatchEventRuleInterface
}

// AwsCloudwatchEventRuleInterface has methods to work with AwsCloudwatchEventRule resources.
type AwsCloudwatchEventRuleInterface interface {
	Create(*v1alpha1.AwsCloudwatchEventRule) (*v1alpha1.AwsCloudwatchEventRule, error)
	Update(*v1alpha1.AwsCloudwatchEventRule) (*v1alpha1.AwsCloudwatchEventRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCloudwatchEventRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCloudwatchEventRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchEventRule, err error)
	AwsCloudwatchEventRuleExpansion
}

// awsCloudwatchEventRules implements AwsCloudwatchEventRuleInterface
type awsCloudwatchEventRules struct {
	client rest.Interface
	ns     string
}

// newAwsCloudwatchEventRules returns a AwsCloudwatchEventRules
func newAwsCloudwatchEventRules(c *ChronojamV1alpha1Client, namespace string) *awsCloudwatchEventRules {
	return &awsCloudwatchEventRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsCloudwatchEventRule, and returns the corresponding awsCloudwatchEventRule object, and an error if there is any.
func (c *awsCloudwatchEventRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudwatchEventRule, err error) {
	result = &v1alpha1.AwsCloudwatchEventRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscloudwatcheventrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCloudwatchEventRules that match those selectors.
func (c *awsCloudwatchEventRules) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudwatchEventRuleList, err error) {
	result = &v1alpha1.AwsCloudwatchEventRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscloudwatcheventrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchEventRules.
func (c *awsCloudwatchEventRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscloudwatcheventrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsCloudwatchEventRule and creates it.  Returns the server's representation of the awsCloudwatchEventRule, and an error, if there is any.
func (c *awsCloudwatchEventRules) Create(awsCloudwatchEventRule *v1alpha1.AwsCloudwatchEventRule) (result *v1alpha1.AwsCloudwatchEventRule, err error) {
	result = &v1alpha1.AwsCloudwatchEventRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscloudwatcheventrules").
		Body(awsCloudwatchEventRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCloudwatchEventRule and updates it. Returns the server's representation of the awsCloudwatchEventRule, and an error, if there is any.
func (c *awsCloudwatchEventRules) Update(awsCloudwatchEventRule *v1alpha1.AwsCloudwatchEventRule) (result *v1alpha1.AwsCloudwatchEventRule, err error) {
	result = &v1alpha1.AwsCloudwatchEventRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscloudwatcheventrules").
		Name(awsCloudwatchEventRule.Name).
		Body(awsCloudwatchEventRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCloudwatchEventRule and deletes it. Returns an error if one occurs.
func (c *awsCloudwatchEventRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscloudwatcheventrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCloudwatchEventRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscloudwatcheventrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCloudwatchEventRule.
func (c *awsCloudwatchEventRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchEventRule, err error) {
	result = &v1alpha1.AwsCloudwatchEventRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscloudwatcheventrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
