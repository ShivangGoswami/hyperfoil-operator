//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthSpec) DeepCopyInto(out *AuthSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthSpec.
func (in *AuthSpec) DeepCopy() *AuthSpec {
	if in == nil {
		return nil
	}
	out := new(AuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hyperfoil) DeepCopyInto(out *Hyperfoil) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hyperfoil.
func (in *Hyperfoil) DeepCopy() *Hyperfoil {
	if in == nil {
		return nil
	}
	out := new(Hyperfoil)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hyperfoil) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperfoilList) DeepCopyInto(out *HyperfoilList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hyperfoil, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperfoilList.
func (in *HyperfoilList) DeepCopy() *HyperfoilList {
	if in == nil {
		return nil
	}
	out := new(HyperfoilList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperfoilList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperfoilSpec) DeepCopyInto(out *HyperfoilSpec) {
	*out = *in
	if in.AdditionalArgs != nil {
		in, out := &in.AdditionalArgs, &out.AdditionalArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Route = in.Route
	out.Auth = in.Auth
	if in.PreHooks != nil {
		in, out := &in.PreHooks, &out.PreHooks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostHooks != nil {
		in, out := &in.PostHooks, &out.PostHooks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecretEnvVars != nil {
		in, out := &in.SecretEnvVars, &out.SecretEnvVars
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperfoilSpec.
func (in *HyperfoilSpec) DeepCopy() *HyperfoilSpec {
	if in == nil {
		return nil
	}
	out := new(HyperfoilSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperfoilStatus) DeepCopyInto(out *HyperfoilStatus) {
	*out = *in
	in.LastUpdate.DeepCopyInto(&out.LastUpdate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperfoilStatus.
func (in *HyperfoilStatus) DeepCopy() *HyperfoilStatus {
	if in == nil {
		return nil
	}
	out := new(HyperfoilStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}
