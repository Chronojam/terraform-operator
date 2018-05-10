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

// FakeAwsElasticBeanstalkEnvironments implements AwsElasticBeanstalkEnvironmentInterface
type FakeAwsElasticBeanstalkEnvironments struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awselasticbeanstalkenvironmentsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awselasticbeanstalkenvironments"}

var awselasticbeanstalkenvironmentsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsElasticBeanstalkEnvironment"}

// Get takes name of the awsElasticBeanstalkEnvironment, and returns the corresponding awsElasticBeanstalkEnvironment object, and an error if there is any.
func (c *FakeAwsElasticBeanstalkEnvironments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticBeanstalkEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awselasticbeanstalkenvironmentsResource, c.ns, name), &v1alpha1.AwsElasticBeanstalkEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkEnvironment), err
}

// List takes label and field selectors, and returns the list of AwsElasticBeanstalkEnvironments that match those selectors.
func (c *FakeAwsElasticBeanstalkEnvironments) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticBeanstalkEnvironmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awselasticbeanstalkenvironmentsResource, awselasticbeanstalkenvironmentsKind, c.ns, opts), &v1alpha1.AwsElasticBeanstalkEnvironmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsElasticBeanstalkEnvironmentList{}
	for _, item := range obj.(*v1alpha1.AwsElasticBeanstalkEnvironmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElasticBeanstalkEnvironments.
func (c *FakeAwsElasticBeanstalkEnvironments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awselasticbeanstalkenvironmentsResource, c.ns, opts))

}

// Create takes the representation of a awsElasticBeanstalkEnvironment and creates it.  Returns the server's representation of the awsElasticBeanstalkEnvironment, and an error, if there is any.
func (c *FakeAwsElasticBeanstalkEnvironments) Create(awsElasticBeanstalkEnvironment *v1alpha1.AwsElasticBeanstalkEnvironment) (result *v1alpha1.AwsElasticBeanstalkEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awselasticbeanstalkenvironmentsResource, c.ns, awsElasticBeanstalkEnvironment), &v1alpha1.AwsElasticBeanstalkEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkEnvironment), err
}

// Update takes the representation of a awsElasticBeanstalkEnvironment and updates it. Returns the server's representation of the awsElasticBeanstalkEnvironment, and an error, if there is any.
func (c *FakeAwsElasticBeanstalkEnvironments) Update(awsElasticBeanstalkEnvironment *v1alpha1.AwsElasticBeanstalkEnvironment) (result *v1alpha1.AwsElasticBeanstalkEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awselasticbeanstalkenvironmentsResource, c.ns, awsElasticBeanstalkEnvironment), &v1alpha1.AwsElasticBeanstalkEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkEnvironment), err
}

// Delete takes name of the awsElasticBeanstalkEnvironment and deletes it. Returns an error if one occurs.
func (c *FakeAwsElasticBeanstalkEnvironments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awselasticbeanstalkenvironmentsResource, c.ns, name), &v1alpha1.AwsElasticBeanstalkEnvironment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElasticBeanstalkEnvironments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awselasticbeanstalkenvironmentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsElasticBeanstalkEnvironmentList{})
	return err
}

// Patch applies the patch and returns the patched awsElasticBeanstalkEnvironment.
func (c *FakeAwsElasticBeanstalkEnvironments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticBeanstalkEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awselasticbeanstalkenvironmentsResource, c.ns, name, data, subresources...), &v1alpha1.AwsElasticBeanstalkEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkEnvironment), err
}
