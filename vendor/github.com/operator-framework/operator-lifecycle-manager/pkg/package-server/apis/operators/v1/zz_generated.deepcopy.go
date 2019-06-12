// +build !ignore_autogenerated

/*
Copyright 2019 Red Hat, Inc.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	v1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppLink) DeepCopyInto(out *AppLink) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppLink.
func (in *AppLink) DeepCopy() *AppLink {
	if in == nil {
		return nil
	}
	out := new(AppLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSVDescription) DeepCopyInto(out *CSVDescription) {
	*out = *in
	if in.Icon != nil {
		in, out := &in.Icon, &out.Icon
		*out = make([]Icon, len(*in))
		copy(*out, *in)
	}
	in.Version.DeepCopyInto(&out.Version)
	out.Provider = in.Provider
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.InstallModes != nil {
		in, out := &in.InstallModes, &out.InstallModes
		*out = make([]v1alpha1.InstallMode, len(*in))
		copy(*out, *in)
	}
	in.CustomResourceDefinitions.DeepCopyInto(&out.CustomResourceDefinitions)
	in.APIServiceDefinitions.DeepCopyInto(&out.APIServiceDefinitions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSVDescription.
func (in *CSVDescription) DeepCopy() *CSVDescription {
	if in == nil {
		return nil
	}
	out := new(CSVDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Icon) DeepCopyInto(out *Icon) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Icon.
func (in *Icon) DeepCopy() *Icon {
	if in == nil {
		return nil
	}
	out := new(Icon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageChannel) DeepCopyInto(out *PackageChannel) {
	*out = *in
	in.CurrentCSVDesc.DeepCopyInto(&out.CurrentCSVDesc)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageChannel.
func (in *PackageChannel) DeepCopy() *PackageChannel {
	if in == nil {
		return nil
	}
	out := new(PackageChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifest) DeepCopyInto(out *PackageManifest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifest.
func (in *PackageManifest) DeepCopy() *PackageManifest {
	if in == nil {
		return nil
	}
	out := new(PackageManifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageManifest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestList) DeepCopyInto(out *PackageManifestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PackageManifest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestList.
func (in *PackageManifestList) DeepCopy() *PackageManifestList {
	if in == nil {
		return nil
	}
	out := new(PackageManifestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageManifestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestSpec) DeepCopyInto(out *PackageManifestSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestSpec.
func (in *PackageManifestSpec) DeepCopy() *PackageManifestSpec {
	if in == nil {
		return nil
	}
	out := new(PackageManifestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestStatus) DeepCopyInto(out *PackageManifestStatus) {
	*out = *in
	out.Provider = in.Provider
	if in.Channels != nil {
		in, out := &in.Channels, &out.Channels
		*out = make([]PackageChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestStatus.
func (in *PackageManifestStatus) DeepCopy() *PackageManifestStatus {
	if in == nil {
		return nil
	}
	out := new(PackageManifestStatus)
	in.DeepCopyInto(out)
	return out
}
