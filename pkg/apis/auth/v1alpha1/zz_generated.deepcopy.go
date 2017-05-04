// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_IdentitySpec, InType: reflect.TypeOf(&IdentitySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_TokenSpec, InType: reflect.TypeOf(&TokenSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_User, InType: reflect.TypeOf(&User{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_UserList, InType: reflect.TypeOf(&UserList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_UserSpec, InType: reflect.TypeOf(&UserSpec{})},
	)
}

func DeepCopy_v1alpha1_IdentitySpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IdentitySpec)
		out := out.(*IdentitySpec)
		*out = *in
		return nil
	}
}

func DeepCopy_v1alpha1_TokenSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TokenSpec)
		out := out.(*TokenSpec)
		*out = *in
		if in.HashedSecret != nil {
			in, out := &in.HashedSecret, &out.HashedSecret
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		if in.RawSecret != nil {
			in, out := &in.RawSecret, &out.RawSecret
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1alpha1_User(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*User)
		out := out.(*User)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_UserSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1alpha1_UserList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserList)
		out := out.(*UserList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]User, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_User(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_UserSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserSpec)
		out := out.(*UserSpec)
		*out = *in
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Tokens != nil {
			in, out := &in.Tokens, &out.Tokens
			*out = make([]*TokenSpec, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(**TokenSpec)
				}
			}
		}
		if in.Identities != nil {
			in, out := &in.Identities, &out.Identities
			*out = make([]IdentitySpec, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}
