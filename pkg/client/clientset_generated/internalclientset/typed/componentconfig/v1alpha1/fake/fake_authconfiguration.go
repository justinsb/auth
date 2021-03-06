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
	v1alpha1 "kope.io/auth/pkg/apis/componentconfig/v1alpha1"
)

// FakeAuthConfigurations implements AuthConfigurationInterface
type FakeAuthConfigurations struct {
	Fake *FakeComponentconfigV1alpha1
	ns   string
}

var authconfigurationsResource = schema.GroupVersionResource{Group: "config.auth.kope.io", Version: "v1alpha1", Resource: "authconfigurations"}

func (c *FakeAuthConfigurations) Create(authConfiguration *v1alpha1.AuthConfiguration) (result *v1alpha1.AuthConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(authconfigurationsResource, c.ns, authConfiguration), &v1alpha1.AuthConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthConfiguration), err
}

func (c *FakeAuthConfigurations) Update(authConfiguration *v1alpha1.AuthConfiguration) (result *v1alpha1.AuthConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(authconfigurationsResource, c.ns, authConfiguration), &v1alpha1.AuthConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthConfiguration), err
}

func (c *FakeAuthConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(authconfigurationsResource, c.ns, name), &v1alpha1.AuthConfiguration{})

	return err
}

func (c *FakeAuthConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(authconfigurationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AuthConfigurationList{})
	return err
}

func (c *FakeAuthConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.AuthConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(authconfigurationsResource, c.ns, name), &v1alpha1.AuthConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthConfiguration), err
}

func (c *FakeAuthConfigurations) List(opts v1.ListOptions) (result *v1alpha1.AuthConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(authconfigurationsResource, c.ns, opts), &v1alpha1.AuthConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AuthConfigurationList{}
	for _, item := range obj.(*v1alpha1.AuthConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested authConfigurations.
func (c *FakeAuthConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(authconfigurationsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched authConfiguration.
func (c *FakeAuthConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AuthConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(authconfigurationsResource, c.ns, name, data, subresources...), &v1alpha1.AuthConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthConfiguration), err
}
