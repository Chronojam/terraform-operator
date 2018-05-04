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

package fake

import (
	v1alpha1 "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAwsElasticBeanstalkConfigurationTemplates implements AwsElasticBeanstalkConfigurationTemplateInterface
type FakeAwsElasticBeanstalkConfigurationTemplates struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awselasticbeanstalkconfigurationtemplatesResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awselasticbeanstalkconfigurationtemplates"}

var awselasticbeanstalkconfigurationtemplatesKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsElasticBeanstalkConfigurationTemplate"}

// Get takes name of the awsElasticBeanstalkConfigurationTemplate, and returns the corresponding awsElasticBeanstalkConfigurationTemplate object, and an error if there is any.
func (c *FakeAwsElasticBeanstalkConfigurationTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awselasticbeanstalkconfigurationtemplatesResource, c.ns, name), &v1alpha1.AwsElasticBeanstalkConfigurationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkConfigurationTemplate), err
}

// List takes label and field selectors, and returns the list of AwsElasticBeanstalkConfigurationTemplates that match those selectors.
func (c *FakeAwsElasticBeanstalkConfigurationTemplates) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awselasticbeanstalkconfigurationtemplatesResource, awselasticbeanstalkconfigurationtemplatesKind, c.ns, opts), &v1alpha1.AwsElasticBeanstalkConfigurationTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsElasticBeanstalkConfigurationTemplateList{}
	for _, item := range obj.(*v1alpha1.AwsElasticBeanstalkConfigurationTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElasticBeanstalkConfigurationTemplates.
func (c *FakeAwsElasticBeanstalkConfigurationTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awselasticbeanstalkconfigurationtemplatesResource, c.ns, opts))

}

// Create takes the representation of a awsElasticBeanstalkConfigurationTemplate and creates it.  Returns the server's representation of the awsElasticBeanstalkConfigurationTemplate, and an error, if there is any.
func (c *FakeAwsElasticBeanstalkConfigurationTemplates) Create(awsElasticBeanstalkConfigurationTemplate *v1alpha1.AwsElasticBeanstalkConfigurationTemplate) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awselasticbeanstalkconfigurationtemplatesResource, c.ns, awsElasticBeanstalkConfigurationTemplate), &v1alpha1.AwsElasticBeanstalkConfigurationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkConfigurationTemplate), err
}

// Update takes the representation of a awsElasticBeanstalkConfigurationTemplate and updates it. Returns the server's representation of the awsElasticBeanstalkConfigurationTemplate, and an error, if there is any.
func (c *FakeAwsElasticBeanstalkConfigurationTemplates) Update(awsElasticBeanstalkConfigurationTemplate *v1alpha1.AwsElasticBeanstalkConfigurationTemplate) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awselasticbeanstalkconfigurationtemplatesResource, c.ns, awsElasticBeanstalkConfigurationTemplate), &v1alpha1.AwsElasticBeanstalkConfigurationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkConfigurationTemplate), err
}

// Delete takes name of the awsElasticBeanstalkConfigurationTemplate and deletes it. Returns an error if one occurs.
func (c *FakeAwsElasticBeanstalkConfigurationTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awselasticbeanstalkconfigurationtemplatesResource, c.ns, name), &v1alpha1.AwsElasticBeanstalkConfigurationTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElasticBeanstalkConfigurationTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awselasticbeanstalkconfigurationtemplatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsElasticBeanstalkConfigurationTemplateList{})
	return err
}

// Patch applies the patch and returns the patched awsElasticBeanstalkConfigurationTemplate.
func (c *FakeAwsElasticBeanstalkConfigurationTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awselasticbeanstalkconfigurationtemplatesResource, c.ns, name, data, subresources...), &v1alpha1.AwsElasticBeanstalkConfigurationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkConfigurationTemplate), err
}
