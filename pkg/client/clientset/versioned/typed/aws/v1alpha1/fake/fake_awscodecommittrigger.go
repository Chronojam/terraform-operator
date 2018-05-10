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

// FakeAwsCodecommitTriggers implements AwsCodecommitTriggerInterface
type FakeAwsCodecommitTriggers struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awscodecommittriggersResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awscodecommittriggers"}

var awscodecommittriggersKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsCodecommitTrigger"}

// Get takes name of the awsCodecommitTrigger, and returns the corresponding awsCodecommitTrigger object, and an error if there is any.
func (c *FakeAwsCodecommitTriggers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCodecommitTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscodecommittriggersResource, c.ns, name), &v1alpha1.AwsCodecommitTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodecommitTrigger), err
}

// List takes label and field selectors, and returns the list of AwsCodecommitTriggers that match those selectors.
func (c *FakeAwsCodecommitTriggers) List(opts v1.ListOptions) (result *v1alpha1.AwsCodecommitTriggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscodecommittriggersResource, awscodecommittriggersKind, c.ns, opts), &v1alpha1.AwsCodecommitTriggerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCodecommitTriggerList{}
	for _, item := range obj.(*v1alpha1.AwsCodecommitTriggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCodecommitTriggers.
func (c *FakeAwsCodecommitTriggers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscodecommittriggersResource, c.ns, opts))

}

// Create takes the representation of a awsCodecommitTrigger and creates it.  Returns the server's representation of the awsCodecommitTrigger, and an error, if there is any.
func (c *FakeAwsCodecommitTriggers) Create(awsCodecommitTrigger *v1alpha1.AwsCodecommitTrigger) (result *v1alpha1.AwsCodecommitTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscodecommittriggersResource, c.ns, awsCodecommitTrigger), &v1alpha1.AwsCodecommitTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodecommitTrigger), err
}

// Update takes the representation of a awsCodecommitTrigger and updates it. Returns the server's representation of the awsCodecommitTrigger, and an error, if there is any.
func (c *FakeAwsCodecommitTriggers) Update(awsCodecommitTrigger *v1alpha1.AwsCodecommitTrigger) (result *v1alpha1.AwsCodecommitTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscodecommittriggersResource, c.ns, awsCodecommitTrigger), &v1alpha1.AwsCodecommitTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodecommitTrigger), err
}

// Delete takes name of the awsCodecommitTrigger and deletes it. Returns an error if one occurs.
func (c *FakeAwsCodecommitTriggers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscodecommittriggersResource, c.ns, name), &v1alpha1.AwsCodecommitTrigger{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCodecommitTriggers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscodecommittriggersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCodecommitTriggerList{})
	return err
}

// Patch applies the patch and returns the patched awsCodecommitTrigger.
func (c *FakeAwsCodecommitTriggers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodecommitTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscodecommittriggersResource, c.ns, name, data, subresources...), &v1alpha1.AwsCodecommitTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodecommitTrigger), err
}
