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

// FakeAwsElasticBeanstalkApplicationVersions implements AwsElasticBeanstalkApplicationVersionInterface
type FakeAwsElasticBeanstalkApplicationVersions struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awselasticbeanstalkapplicationversionsResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awselasticbeanstalkapplicationversions"}

var awselasticbeanstalkapplicationversionsKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsElasticBeanstalkApplicationVersion"}

// Get takes name of the awsElasticBeanstalkApplicationVersion, and returns the corresponding awsElasticBeanstalkApplicationVersion object, and an error if there is any.
func (c *FakeAwsElasticBeanstalkApplicationVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticBeanstalkApplicationVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awselasticbeanstalkapplicationversionsResource, c.ns, name), &v1alpha1.AwsElasticBeanstalkApplicationVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkApplicationVersion), err
}

// List takes label and field selectors, and returns the list of AwsElasticBeanstalkApplicationVersions that match those selectors.
func (c *FakeAwsElasticBeanstalkApplicationVersions) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticBeanstalkApplicationVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awselasticbeanstalkapplicationversionsResource, awselasticbeanstalkapplicationversionsKind, c.ns, opts), &v1alpha1.AwsElasticBeanstalkApplicationVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsElasticBeanstalkApplicationVersionList{}
	for _, item := range obj.(*v1alpha1.AwsElasticBeanstalkApplicationVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElasticBeanstalkApplicationVersions.
func (c *FakeAwsElasticBeanstalkApplicationVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awselasticbeanstalkapplicationversionsResource, c.ns, opts))

}

// Create takes the representation of a awsElasticBeanstalkApplicationVersion and creates it.  Returns the server's representation of the awsElasticBeanstalkApplicationVersion, and an error, if there is any.
func (c *FakeAwsElasticBeanstalkApplicationVersions) Create(awsElasticBeanstalkApplicationVersion *v1alpha1.AwsElasticBeanstalkApplicationVersion) (result *v1alpha1.AwsElasticBeanstalkApplicationVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awselasticbeanstalkapplicationversionsResource, c.ns, awsElasticBeanstalkApplicationVersion), &v1alpha1.AwsElasticBeanstalkApplicationVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkApplicationVersion), err
}

// Update takes the representation of a awsElasticBeanstalkApplicationVersion and updates it. Returns the server's representation of the awsElasticBeanstalkApplicationVersion, and an error, if there is any.
func (c *FakeAwsElasticBeanstalkApplicationVersions) Update(awsElasticBeanstalkApplicationVersion *v1alpha1.AwsElasticBeanstalkApplicationVersion) (result *v1alpha1.AwsElasticBeanstalkApplicationVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awselasticbeanstalkapplicationversionsResource, c.ns, awsElasticBeanstalkApplicationVersion), &v1alpha1.AwsElasticBeanstalkApplicationVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkApplicationVersion), err
}

// Delete takes name of the awsElasticBeanstalkApplicationVersion and deletes it. Returns an error if one occurs.
func (c *FakeAwsElasticBeanstalkApplicationVersions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awselasticbeanstalkapplicationversionsResource, c.ns, name), &v1alpha1.AwsElasticBeanstalkApplicationVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElasticBeanstalkApplicationVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awselasticbeanstalkapplicationversionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsElasticBeanstalkApplicationVersionList{})
	return err
}

// Patch applies the patch and returns the patched awsElasticBeanstalkApplicationVersion.
func (c *FakeAwsElasticBeanstalkApplicationVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticBeanstalkApplicationVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awselasticbeanstalkapplicationversionsResource, c.ns, name, data, subresources...), &v1alpha1.AwsElasticBeanstalkApplicationVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkApplicationVersion), err
}
