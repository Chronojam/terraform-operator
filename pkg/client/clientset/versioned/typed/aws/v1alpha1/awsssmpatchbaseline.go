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

// AwsSsmPatchBaselinesGetter has a method to return a AwsSsmPatchBaselineInterface.
// A group's client should implement this interface.
type AwsSsmPatchBaselinesGetter interface {
	AwsSsmPatchBaselines(namespace string) AwsSsmPatchBaselineInterface
}

// AwsSsmPatchBaselineInterface has methods to work with AwsSsmPatchBaseline resources.
type AwsSsmPatchBaselineInterface interface {
	Create(*v1alpha1.AwsSsmPatchBaseline) (*v1alpha1.AwsSsmPatchBaseline, error)
	Update(*v1alpha1.AwsSsmPatchBaseline) (*v1alpha1.AwsSsmPatchBaseline, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSsmPatchBaseline, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSsmPatchBaselineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSsmPatchBaseline, err error)
	AwsSsmPatchBaselineExpansion
}

// awsSsmPatchBaselines implements AwsSsmPatchBaselineInterface
type awsSsmPatchBaselines struct {
	client rest.Interface
	ns     string
}

// newAwsSsmPatchBaselines returns a AwsSsmPatchBaselines
func newAwsSsmPatchBaselines(c *ChronojamV1alpha1Client, namespace string) *awsSsmPatchBaselines {
	return &awsSsmPatchBaselines{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSsmPatchBaseline, and returns the corresponding awsSsmPatchBaseline object, and an error if there is any.
func (c *awsSsmPatchBaselines) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSsmPatchBaseline, err error) {
	result = &v1alpha1.AwsSsmPatchBaseline{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsssmpatchbaselines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSsmPatchBaselines that match those selectors.
func (c *awsSsmPatchBaselines) List(opts v1.ListOptions) (result *v1alpha1.AwsSsmPatchBaselineList, err error) {
	result = &v1alpha1.AwsSsmPatchBaselineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsssmpatchbaselines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSsmPatchBaselines.
func (c *awsSsmPatchBaselines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsssmpatchbaselines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSsmPatchBaseline and creates it.  Returns the server's representation of the awsSsmPatchBaseline, and an error, if there is any.
func (c *awsSsmPatchBaselines) Create(awsSsmPatchBaseline *v1alpha1.AwsSsmPatchBaseline) (result *v1alpha1.AwsSsmPatchBaseline, err error) {
	result = &v1alpha1.AwsSsmPatchBaseline{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsssmpatchbaselines").
		Body(awsSsmPatchBaseline).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSsmPatchBaseline and updates it. Returns the server's representation of the awsSsmPatchBaseline, and an error, if there is any.
func (c *awsSsmPatchBaselines) Update(awsSsmPatchBaseline *v1alpha1.AwsSsmPatchBaseline) (result *v1alpha1.AwsSsmPatchBaseline, err error) {
	result = &v1alpha1.AwsSsmPatchBaseline{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsssmpatchbaselines").
		Name(awsSsmPatchBaseline.Name).
		Body(awsSsmPatchBaseline).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSsmPatchBaseline and deletes it. Returns an error if one occurs.
func (c *awsSsmPatchBaselines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsssmpatchbaselines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSsmPatchBaselines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsssmpatchbaselines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSsmPatchBaseline.
func (c *awsSsmPatchBaselines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSsmPatchBaseline, err error) {
	result = &v1alpha1.AwsSsmPatchBaseline{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsssmpatchbaselines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
