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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	auth "kope.io/auth/pkg/apis/auth"
)

// FakeUsers implements UserInterface
type FakeUsers struct {
	Fake *FakeAuth
	ns   string
}

var usersResource = schema.GroupVersionResource{Group: "auth.kope.io", Version: "", Resource: "users"}

func (c *FakeUsers) Create(user *auth.User) (result *auth.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(usersResource, c.ns, user), &auth.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*auth.User), err
}

func (c *FakeUsers) Update(user *auth.User) (result *auth.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(usersResource, c.ns, user), &auth.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*auth.User), err
}

func (c *FakeUsers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(usersResource, c.ns, name), &auth.User{})

	return err
}

func (c *FakeUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(usersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &auth.UserList{})
	return err
}

func (c *FakeUsers) Get(name string, options v1.GetOptions) (result *auth.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(usersResource, c.ns, name), &auth.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*auth.User), err
}

func (c *FakeUsers) List(opts v1.ListOptions) (result *auth.UserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(usersResource, c.ns, opts), &auth.UserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &auth.UserList{}
	for _, item := range obj.(*auth.UserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested users.
func (c *FakeUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(usersResource, c.ns, opts))

}

// Patch applies the patch and returns the patched user.
func (c *FakeUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(usersResource, c.ns, name, data, subresources...), &auth.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*auth.User), err
}
