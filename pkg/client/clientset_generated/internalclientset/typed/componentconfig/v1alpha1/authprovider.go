/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kope.io/auth/pkg/apis/componentconfig/v1alpha1"
	scheme "kope.io/auth/pkg/client/clientset_generated/internalclientset/scheme"
)

// AuthProvidersGetter has a method to return a AuthProviderInterface.
// A group's client should implement this interface.
type AuthProvidersGetter interface {
	AuthProviders(namespace string) AuthProviderInterface
}

// AuthProviderInterface has methods to work with AuthProvider resources.
type AuthProviderInterface interface {
	Create(*v1alpha1.AuthProvider) (*v1alpha1.AuthProvider, error)
	Update(*v1alpha1.AuthProvider) (*v1alpha1.AuthProvider, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AuthProvider, error)
	List(opts v1.ListOptions) (*v1alpha1.AuthProviderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AuthProvider, err error)
	AuthProviderExpansion
}

// authProviders implements AuthProviderInterface
type authProviders struct {
	client rest.Interface
	ns     string
}

// newAuthProviders returns a AuthProviders
func newAuthProviders(c *ComponentconfigV1alpha1Client, namespace string) *authProviders {
	return &authProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a authProvider and creates it.  Returns the server's representation of the authProvider, and an error, if there is any.
func (c *authProviders) Create(authProvider *v1alpha1.AuthProvider) (result *v1alpha1.AuthProvider, err error) {
	result = &v1alpha1.AuthProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("authproviders").
		Body(authProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a authProvider and updates it. Returns the server's representation of the authProvider, and an error, if there is any.
func (c *authProviders) Update(authProvider *v1alpha1.AuthProvider) (result *v1alpha1.AuthProvider, err error) {
	result = &v1alpha1.AuthProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("authproviders").
		Name(authProvider.Name).
		Body(authProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the authProvider and deletes it. Returns an error if one occurs.
func (c *authProviders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("authproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *authProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("authproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the authProvider, and returns the corresponding authProvider object, and an error if there is any.
func (c *authProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.AuthProvider, err error) {
	result = &v1alpha1.AuthProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("authproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AuthProviders that match those selectors.
func (c *authProviders) List(opts v1.ListOptions) (result *v1alpha1.AuthProviderList, err error) {
	result = &v1alpha1.AuthProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("authproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested authProviders.
func (c *authProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("authproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched authProvider.
func (c *authProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AuthProvider, err error) {
	result = &v1alpha1.AuthProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("authproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
