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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "kope.io/auth/pkg/client/clientset_generated/internalclientset"
	authinternalversion "kope.io/auth/pkg/client/clientset_generated/internalclientset/typed/auth/internalversion"
	fakeauthinternalversion "kope.io/auth/pkg/client/clientset_generated/internalclientset/typed/auth/internalversion/fake"
	authv1alpha1 "kope.io/auth/pkg/client/clientset_generated/internalclientset/typed/auth/v1alpha1"
	fakeauthv1alpha1 "kope.io/auth/pkg/client/clientset_generated/internalclientset/typed/auth/v1alpha1/fake"
	componentconfiginternalversion "kope.io/auth/pkg/client/clientset_generated/internalclientset/typed/componentconfig/internalversion"
	fakecomponentconfiginternalversion "kope.io/auth/pkg/client/clientset_generated/internalclientset/typed/componentconfig/internalversion/fake"
	componentconfigv1alpha1 "kope.io/auth/pkg/client/clientset_generated/internalclientset/typed/componentconfig/v1alpha1"
	fakecomponentconfigv1alpha1 "kope.io/auth/pkg/client/clientset_generated/internalclientset/typed/componentconfig/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(registry, scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o, registry.RESTMapper()))

	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return &fakediscovery.FakeDiscovery{Fake: &c.Fake}
}

var _ clientset.Interface = &Clientset{}

// Auth retrieves the AuthClient
func (c *Clientset) Auth() authinternalversion.AuthInterface {
	return &fakeauthinternalversion.FakeAuth{Fake: &c.Fake}
}

// AuthV1alpha1 retrieves the AuthV1alpha1Client
func (c *Clientset) AuthV1alpha1() authv1alpha1.AuthV1alpha1Interface {
	return &fakeauthv1alpha1.FakeAuthV1alpha1{Fake: &c.Fake}
}

// Componentconfig retrieves the ComponentconfigClient
func (c *Clientset) Componentconfig() componentconfiginternalversion.ComponentconfigInterface {
	return &fakecomponentconfiginternalversion.FakeComponentconfig{Fake: &c.Fake}
}

// ComponentconfigV1alpha1 retrieves the ComponentconfigV1alpha1Client
func (c *Clientset) ComponentconfigV1alpha1() componentconfigv1alpha1.ComponentconfigV1alpha1Interface {
	return &fakecomponentconfigv1alpha1.FakeComponentconfigV1alpha1{Fake: &c.Fake}
}
