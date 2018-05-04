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

// AwsAutoscalingSchedulesGetter has a method to return a AwsAutoscalingScheduleInterface.
// A group's client should implement this interface.
type AwsAutoscalingSchedulesGetter interface {
	AwsAutoscalingSchedules(namespace string) AwsAutoscalingScheduleInterface
}

// AwsAutoscalingScheduleInterface has methods to work with AwsAutoscalingSchedule resources.
type AwsAutoscalingScheduleInterface interface {
	Create(*v1alpha1.AwsAutoscalingSchedule) (*v1alpha1.AwsAutoscalingSchedule, error)
	Update(*v1alpha1.AwsAutoscalingSchedule) (*v1alpha1.AwsAutoscalingSchedule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAutoscalingSchedule, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAutoscalingScheduleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingSchedule, err error)
	AwsAutoscalingScheduleExpansion
}

// awsAutoscalingSchedules implements AwsAutoscalingScheduleInterface
type awsAutoscalingSchedules struct {
	client rest.Interface
	ns     string
}

// newAwsAutoscalingSchedules returns a AwsAutoscalingSchedules
func newAwsAutoscalingSchedules(c *AwsV1alpha1Client, namespace string) *awsAutoscalingSchedules {
	return &awsAutoscalingSchedules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsAutoscalingSchedule, and returns the corresponding awsAutoscalingSchedule object, and an error if there is any.
func (c *awsAutoscalingSchedules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAutoscalingSchedule, err error) {
	result = &v1alpha1.AwsAutoscalingSchedule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalingschedules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAutoscalingSchedules that match those selectors.
func (c *awsAutoscalingSchedules) List(opts v1.ListOptions) (result *v1alpha1.AwsAutoscalingScheduleList, err error) {
	result = &v1alpha1.AwsAutoscalingScheduleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalingschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAutoscalingSchedules.
func (c *awsAutoscalingSchedules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalingschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsAutoscalingSchedule and creates it.  Returns the server's representation of the awsAutoscalingSchedule, and an error, if there is any.
func (c *awsAutoscalingSchedules) Create(awsAutoscalingSchedule *v1alpha1.AwsAutoscalingSchedule) (result *v1alpha1.AwsAutoscalingSchedule, err error) {
	result = &v1alpha1.AwsAutoscalingSchedule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsautoscalingschedules").
		Body(awsAutoscalingSchedule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAutoscalingSchedule and updates it. Returns the server's representation of the awsAutoscalingSchedule, and an error, if there is any.
func (c *awsAutoscalingSchedules) Update(awsAutoscalingSchedule *v1alpha1.AwsAutoscalingSchedule) (result *v1alpha1.AwsAutoscalingSchedule, err error) {
	result = &v1alpha1.AwsAutoscalingSchedule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsautoscalingschedules").
		Name(awsAutoscalingSchedule.Name).
		Body(awsAutoscalingSchedule).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAutoscalingSchedule and deletes it. Returns an error if one occurs.
func (c *awsAutoscalingSchedules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsautoscalingschedules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAutoscalingSchedules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsautoscalingschedules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAutoscalingSchedule.
func (c *awsAutoscalingSchedules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingSchedule, err error) {
	result = &v1alpha1.AwsAutoscalingSchedule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsautoscalingschedules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
