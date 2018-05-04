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

// AwsSecretsmanagerSecretsGetter has a method to return a AwsSecretsmanagerSecretInterface.
// A group's client should implement this interface.
type AwsSecretsmanagerSecretsGetter interface {
	AwsSecretsmanagerSecrets(namespace string) AwsSecretsmanagerSecretInterface
}

// AwsSecretsmanagerSecretInterface has methods to work with AwsSecretsmanagerSecret resources.
type AwsSecretsmanagerSecretInterface interface {
	Create(*v1alpha1.AwsSecretsmanagerSecret) (*v1alpha1.AwsSecretsmanagerSecret, error)
	Update(*v1alpha1.AwsSecretsmanagerSecret) (*v1alpha1.AwsSecretsmanagerSecret, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSecretsmanagerSecret, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSecretsmanagerSecretList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSecretsmanagerSecret, err error)
	AwsSecretsmanagerSecretExpansion
}

// awsSecretsmanagerSecrets implements AwsSecretsmanagerSecretInterface
type awsSecretsmanagerSecrets struct {
	client rest.Interface
	ns     string
}

// newAwsSecretsmanagerSecrets returns a AwsSecretsmanagerSecrets
func newAwsSecretsmanagerSecrets(c *AwsV1alpha1Client, namespace string) *awsSecretsmanagerSecrets {
	return &awsSecretsmanagerSecrets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSecretsmanagerSecret, and returns the corresponding awsSecretsmanagerSecret object, and an error if there is any.
func (c *awsSecretsmanagerSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSecretsmanagerSecret, err error) {
	result = &v1alpha1.AwsSecretsmanagerSecret{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssecretsmanagersecrets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSecretsmanagerSecrets that match those selectors.
func (c *awsSecretsmanagerSecrets) List(opts v1.ListOptions) (result *v1alpha1.AwsSecretsmanagerSecretList, err error) {
	result = &v1alpha1.AwsSecretsmanagerSecretList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssecretsmanagersecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSecretsmanagerSecrets.
func (c *awsSecretsmanagerSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssecretsmanagersecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSecretsmanagerSecret and creates it.  Returns the server's representation of the awsSecretsmanagerSecret, and an error, if there is any.
func (c *awsSecretsmanagerSecrets) Create(awsSecretsmanagerSecret *v1alpha1.AwsSecretsmanagerSecret) (result *v1alpha1.AwsSecretsmanagerSecret, err error) {
	result = &v1alpha1.AwsSecretsmanagerSecret{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssecretsmanagersecrets").
		Body(awsSecretsmanagerSecret).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSecretsmanagerSecret and updates it. Returns the server's representation of the awsSecretsmanagerSecret, and an error, if there is any.
func (c *awsSecretsmanagerSecrets) Update(awsSecretsmanagerSecret *v1alpha1.AwsSecretsmanagerSecret) (result *v1alpha1.AwsSecretsmanagerSecret, err error) {
	result = &v1alpha1.AwsSecretsmanagerSecret{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssecretsmanagersecrets").
		Name(awsSecretsmanagerSecret.Name).
		Body(awsSecretsmanagerSecret).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSecretsmanagerSecret and deletes it. Returns an error if one occurs.
func (c *awsSecretsmanagerSecrets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssecretsmanagersecrets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSecretsmanagerSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssecretsmanagersecrets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSecretsmanagerSecret.
func (c *awsSecretsmanagerSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSecretsmanagerSecret, err error) {
	result = &v1alpha1.AwsSecretsmanagerSecret{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssecretsmanagersecrets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
