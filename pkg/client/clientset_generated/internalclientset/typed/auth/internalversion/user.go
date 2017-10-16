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

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	auth "kope.io/auth/pkg/apis/auth"
	scheme "kope.io/auth/pkg/client/clientset_generated/internalclientset/scheme"
)

// UsersGetter has a method to return a UserInterface.
// A group's client should implement this interface.
type UsersGetter interface {
	Users(namespace string) UserInterface
}

// UserInterface has methods to work with User resources.
type UserInterface interface {
	Create(*auth.User) (*auth.User, error)
	Update(*auth.User) (*auth.User, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*auth.User, error)
	List(opts v1.ListOptions) (*auth.UserList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.User, err error)
	UserExpansion
}

// users implements UserInterface
type users struct {
	client rest.Interface
	ns     string
}

// newUsers returns a Users
func newUsers(c *AuthClient, namespace string) *users {
	return &users{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the user, and returns the corresponding user object, and an error if there is any.
func (c *users) Get(name string, options v1.GetOptions) (result *auth.User, err error) {
	result = &auth.User{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("users").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Users that match those selectors.
func (c *users) List(opts v1.ListOptions) (result *auth.UserList, err error) {
	result = &auth.UserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("users").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested users.
func (c *users) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("users").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a user and creates it.  Returns the server's representation of the user, and an error, if there is any.
func (c *users) Create(user *auth.User) (result *auth.User, err error) {
	result = &auth.User{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("users").
		Body(user).
		Do().
		Into(result)
	return
}

// Update takes the representation of a user and updates it. Returns the server's representation of the user, and an error, if there is any.
func (c *users) Update(user *auth.User) (result *auth.User, err error) {
	result = &auth.User{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("users").
		Name(user.Name).
		Body(user).
		Do().
		Into(result)
	return
}

// Delete takes name of the user and deletes it. Returns an error if one occurs.
func (c *users) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("users").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *users) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("users").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched user.
func (c *users) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.User, err error) {
	result = &auth.User{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("users").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
