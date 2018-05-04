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

// FakeAwsRoute53HealthChecks implements AwsRoute53HealthCheckInterface
type FakeAwsRoute53HealthChecks struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsroute53healthchecksResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsroute53healthchecks"}

var awsroute53healthchecksKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsRoute53HealthCheck"}

// Get takes name of the awsRoute53HealthCheck, and returns the corresponding awsRoute53HealthCheck object, and an error if there is any.
func (c *FakeAwsRoute53HealthChecks) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRoute53HealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsroute53healthchecksResource, c.ns, name), &v1alpha1.AwsRoute53HealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRoute53HealthCheck), err
}

// List takes label and field selectors, and returns the list of AwsRoute53HealthChecks that match those selectors.
func (c *FakeAwsRoute53HealthChecks) List(opts v1.ListOptions) (result *v1alpha1.AwsRoute53HealthCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsroute53healthchecksResource, awsroute53healthchecksKind, c.ns, opts), &v1alpha1.AwsRoute53HealthCheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsRoute53HealthCheckList{}
	for _, item := range obj.(*v1alpha1.AwsRoute53HealthCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsRoute53HealthChecks.
func (c *FakeAwsRoute53HealthChecks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsroute53healthchecksResource, c.ns, opts))

}

// Create takes the representation of a awsRoute53HealthCheck and creates it.  Returns the server's representation of the awsRoute53HealthCheck, and an error, if there is any.
func (c *FakeAwsRoute53HealthChecks) Create(awsRoute53HealthCheck *v1alpha1.AwsRoute53HealthCheck) (result *v1alpha1.AwsRoute53HealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsroute53healthchecksResource, c.ns, awsRoute53HealthCheck), &v1alpha1.AwsRoute53HealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRoute53HealthCheck), err
}

// Update takes the representation of a awsRoute53HealthCheck and updates it. Returns the server's representation of the awsRoute53HealthCheck, and an error, if there is any.
func (c *FakeAwsRoute53HealthChecks) Update(awsRoute53HealthCheck *v1alpha1.AwsRoute53HealthCheck) (result *v1alpha1.AwsRoute53HealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsroute53healthchecksResource, c.ns, awsRoute53HealthCheck), &v1alpha1.AwsRoute53HealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRoute53HealthCheck), err
}

// Delete takes name of the awsRoute53HealthCheck and deletes it. Returns an error if one occurs.
func (c *FakeAwsRoute53HealthChecks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsroute53healthchecksResource, c.ns, name), &v1alpha1.AwsRoute53HealthCheck{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsRoute53HealthChecks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsroute53healthchecksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsRoute53HealthCheckList{})
	return err
}

// Patch applies the patch and returns the patched awsRoute53HealthCheck.
func (c *FakeAwsRoute53HealthChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRoute53HealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsroute53healthchecksResource, c.ns, name, data, subresources...), &v1alpha1.AwsRoute53HealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRoute53HealthCheck), err
}
