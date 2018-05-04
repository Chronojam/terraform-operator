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

// AwsOpsworksPhpAppLayersGetter has a method to return a AwsOpsworksPhpAppLayerInterface.
// A group's client should implement this interface.
type AwsOpsworksPhpAppLayersGetter interface {
	AwsOpsworksPhpAppLayers(namespace string) AwsOpsworksPhpAppLayerInterface
}

// AwsOpsworksPhpAppLayerInterface has methods to work with AwsOpsworksPhpAppLayer resources.
type AwsOpsworksPhpAppLayerInterface interface {
	Create(*v1alpha1.AwsOpsworksPhpAppLayer) (*v1alpha1.AwsOpsworksPhpAppLayer, error)
	Update(*v1alpha1.AwsOpsworksPhpAppLayer) (*v1alpha1.AwsOpsworksPhpAppLayer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsOpsworksPhpAppLayer, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsOpsworksPhpAppLayerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksPhpAppLayer, err error)
	AwsOpsworksPhpAppLayerExpansion
}

// awsOpsworksPhpAppLayers implements AwsOpsworksPhpAppLayerInterface
type awsOpsworksPhpAppLayers struct {
	client rest.Interface
	ns     string
}

// newAwsOpsworksPhpAppLayers returns a AwsOpsworksPhpAppLayers
func newAwsOpsworksPhpAppLayers(c *AwsV1alpha1Client, namespace string) *awsOpsworksPhpAppLayers {
	return &awsOpsworksPhpAppLayers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsOpsworksPhpAppLayer, and returns the corresponding awsOpsworksPhpAppLayer object, and an error if there is any.
func (c *awsOpsworksPhpAppLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksPhpAppLayer, err error) {
	result = &v1alpha1.AwsOpsworksPhpAppLayer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksphpapplayers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOpsworksPhpAppLayers that match those selectors.
func (c *awsOpsworksPhpAppLayers) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksPhpAppLayerList, err error) {
	result = &v1alpha1.AwsOpsworksPhpAppLayerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksphpapplayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOpsworksPhpAppLayers.
func (c *awsOpsworksPhpAppLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksphpapplayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsOpsworksPhpAppLayer and creates it.  Returns the server's representation of the awsOpsworksPhpAppLayer, and an error, if there is any.
func (c *awsOpsworksPhpAppLayers) Create(awsOpsworksPhpAppLayer *v1alpha1.AwsOpsworksPhpAppLayer) (result *v1alpha1.AwsOpsworksPhpAppLayer, err error) {
	result = &v1alpha1.AwsOpsworksPhpAppLayer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsopsworksphpapplayers").
		Body(awsOpsworksPhpAppLayer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOpsworksPhpAppLayer and updates it. Returns the server's representation of the awsOpsworksPhpAppLayer, and an error, if there is any.
func (c *awsOpsworksPhpAppLayers) Update(awsOpsworksPhpAppLayer *v1alpha1.AwsOpsworksPhpAppLayer) (result *v1alpha1.AwsOpsworksPhpAppLayer, err error) {
	result = &v1alpha1.AwsOpsworksPhpAppLayer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsopsworksphpapplayers").
		Name(awsOpsworksPhpAppLayer.Name).
		Body(awsOpsworksPhpAppLayer).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOpsworksPhpAppLayer and deletes it. Returns an error if one occurs.
func (c *awsOpsworksPhpAppLayers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworksphpapplayers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOpsworksPhpAppLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworksphpapplayers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOpsworksPhpAppLayer.
func (c *awsOpsworksPhpAppLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksPhpAppLayer, err error) {
	result = &v1alpha1.AwsOpsworksPhpAppLayer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsopsworksphpapplayers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
