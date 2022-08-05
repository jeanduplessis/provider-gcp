//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternativeNameServerConfigObservation) DeepCopyInto(out *AlternativeNameServerConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternativeNameServerConfigObservation.
func (in *AlternativeNameServerConfigObservation) DeepCopy() *AlternativeNameServerConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AlternativeNameServerConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternativeNameServerConfigParameters) DeepCopyInto(out *AlternativeNameServerConfigParameters) {
	*out = *in
	if in.TargetNameServers != nil {
		in, out := &in.TargetNameServers, &out.TargetNameServers
		*out = make([]AlternativeNameServerConfigTargetNameServersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternativeNameServerConfigParameters.
func (in *AlternativeNameServerConfigParameters) DeepCopy() *AlternativeNameServerConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AlternativeNameServerConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternativeNameServerConfigTargetNameServersObservation) DeepCopyInto(out *AlternativeNameServerConfigTargetNameServersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternativeNameServerConfigTargetNameServersObservation.
func (in *AlternativeNameServerConfigTargetNameServersObservation) DeepCopy() *AlternativeNameServerConfigTargetNameServersObservation {
	if in == nil {
		return nil
	}
	out := new(AlternativeNameServerConfigTargetNameServersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternativeNameServerConfigTargetNameServersParameters) DeepCopyInto(out *AlternativeNameServerConfigTargetNameServersParameters) {
	*out = *in
	if in.ForwardingPath != nil {
		in, out := &in.ForwardingPath, &out.ForwardingPath
		*out = new(string)
		**out = **in
	}
	if in.IPv4Address != nil {
		in, out := &in.IPv4Address, &out.IPv4Address
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternativeNameServerConfigTargetNameServersParameters.
func (in *AlternativeNameServerConfigTargetNameServersParameters) DeepCopy() *AlternativeNameServerConfigTargetNameServersParameters {
	if in == nil {
		return nil
	}
	out := new(AlternativeNameServerConfigTargetNameServersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSECConfigObservation) DeepCopyInto(out *DNSSECConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSECConfigObservation.
func (in *DNSSECConfigObservation) DeepCopy() *DNSSECConfigObservation {
	if in == nil {
		return nil
	}
	out := new(DNSSECConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSECConfigParameters) DeepCopyInto(out *DNSSECConfigParameters) {
	*out = *in
	if in.DefaultKeySpecs != nil {
		in, out := &in.DefaultKeySpecs, &out.DefaultKeySpecs
		*out = make([]DefaultKeySpecsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.NonExistence != nil {
		in, out := &in.NonExistence, &out.NonExistence
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSECConfigParameters.
func (in *DNSSECConfigParameters) DeepCopy() *DNSSECConfigParameters {
	if in == nil {
		return nil
	}
	out := new(DNSSECConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultKeySpecsObservation) DeepCopyInto(out *DefaultKeySpecsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultKeySpecsObservation.
func (in *DefaultKeySpecsObservation) DeepCopy() *DefaultKeySpecsObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultKeySpecsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultKeySpecsParameters) DeepCopyInto(out *DefaultKeySpecsParameters) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.KeyLength != nil {
		in, out := &in.KeyLength, &out.KeyLength
		*out = new(float64)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultKeySpecsParameters.
func (in *DefaultKeySpecsParameters) DeepCopy() *DefaultKeySpecsParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultKeySpecsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingConfigObservation) DeepCopyInto(out *ForwardingConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingConfigObservation.
func (in *ForwardingConfigObservation) DeepCopy() *ForwardingConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ForwardingConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingConfigParameters) DeepCopyInto(out *ForwardingConfigParameters) {
	*out = *in
	if in.TargetNameServers != nil {
		in, out := &in.TargetNameServers, &out.TargetNameServers
		*out = make([]TargetNameServersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingConfigParameters.
func (in *ForwardingConfigParameters) DeepCopy() *ForwardingConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ForwardingConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeoObservation) DeepCopyInto(out *GeoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeoObservation.
func (in *GeoObservation) DeepCopy() *GeoObservation {
	if in == nil {
		return nil
	}
	out := new(GeoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeoParameters) DeepCopyInto(out *GeoParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeoParameters.
func (in *GeoParameters) DeepCopy() *GeoParameters {
	if in == nil {
		return nil
	}
	out := new(GeoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZone) DeepCopyInto(out *ManagedZone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZone.
func (in *ManagedZone) DeepCopy() *ManagedZone {
	if in == nil {
		return nil
	}
	out := new(ManagedZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedZone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneList) DeepCopyInto(out *ManagedZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedZone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneList.
func (in *ManagedZoneList) DeepCopy() *ManagedZoneList {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneObservation) DeepCopyInto(out *ManagedZoneObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NameServers != nil {
		in, out := &in.NameServers, &out.NameServers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneObservation.
func (in *ManagedZoneObservation) DeepCopy() *ManagedZoneObservation {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneParameters) DeepCopyInto(out *ManagedZoneParameters) {
	*out = *in
	if in.DNSName != nil {
		in, out := &in.DNSName, &out.DNSName
		*out = new(string)
		**out = **in
	}
	if in.DNSSECConfig != nil {
		in, out := &in.DNSSECConfig, &out.DNSSECConfig
		*out = make([]DNSSECConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.ForwardingConfig != nil {
		in, out := &in.ForwardingConfig, &out.ForwardingConfig
		*out = make([]ForwardingConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PeeringConfig != nil {
		in, out := &in.PeeringConfig, &out.PeeringConfig
		*out = make([]PeeringConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateVisibilityConfig != nil {
		in, out := &in.PrivateVisibilityConfig, &out.PrivateVisibilityConfig
		*out = make([]PrivateVisibilityConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneParameters.
func (in *ManagedZoneParameters) DeepCopy() *ManagedZoneParameters {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpec) DeepCopyInto(out *ManagedZoneSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpec.
func (in *ManagedZoneSpec) DeepCopy() *ManagedZoneSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneStatus) DeepCopyInto(out *ManagedZoneStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneStatus.
func (in *ManagedZoneStatus) DeepCopy() *ManagedZoneStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworksObservation) DeepCopyInto(out *NetworksObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworksObservation.
func (in *NetworksObservation) DeepCopy() *NetworksObservation {
	if in == nil {
		return nil
	}
	out := new(NetworksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworksParameters) DeepCopyInto(out *NetworksParameters) {
	*out = *in
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	if in.NetworkURLRef != nil {
		in, out := &in.NetworkURLRef, &out.NetworkURLRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkURLSelector != nil {
		in, out := &in.NetworkURLSelector, &out.NetworkURLSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworksParameters.
func (in *NetworksParameters) DeepCopy() *NetworksParameters {
	if in == nil {
		return nil
	}
	out := new(NetworksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeeringConfigObservation) DeepCopyInto(out *PeeringConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeeringConfigObservation.
func (in *PeeringConfigObservation) DeepCopy() *PeeringConfigObservation {
	if in == nil {
		return nil
	}
	out := new(PeeringConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeeringConfigParameters) DeepCopyInto(out *PeeringConfigParameters) {
	*out = *in
	if in.TargetNetwork != nil {
		in, out := &in.TargetNetwork, &out.TargetNetwork
		*out = make([]TargetNetworkParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeeringConfigParameters.
func (in *PeeringConfigParameters) DeepCopy() *PeeringConfigParameters {
	if in == nil {
		return nil
	}
	out := new(PeeringConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyNetworksObservation) DeepCopyInto(out *PolicyNetworksObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyNetworksObservation.
func (in *PolicyNetworksObservation) DeepCopy() *PolicyNetworksObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyNetworksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyNetworksParameters) DeepCopyInto(out *PolicyNetworksParameters) {
	*out = *in
	if in.NetworkRef != nil {
		in, out := &in.NetworkRef, &out.NetworkRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkSelector != nil {
		in, out := &in.NetworkSelector, &out.NetworkSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyNetworksParameters.
func (in *PolicyNetworksParameters) DeepCopy() *PolicyNetworksParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyNetworksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyObservation) DeepCopyInto(out *PolicyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyObservation.
func (in *PolicyObservation) DeepCopy() *PolicyObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyParameters) DeepCopyInto(out *PolicyParameters) {
	*out = *in
	if in.AlternativeNameServerConfig != nil {
		in, out := &in.AlternativeNameServerConfig, &out.AlternativeNameServerConfig
		*out = make([]AlternativeNameServerConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnableInboundForwarding != nil {
		in, out := &in.EnableInboundForwarding, &out.EnableInboundForwarding
		*out = new(bool)
		**out = **in
	}
	if in.EnableLogging != nil {
		in, out := &in.EnableLogging, &out.EnableLogging
		*out = new(bool)
		**out = **in
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]PolicyNetworksParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyParameters.
func (in *PolicyParameters) DeepCopy() *PolicyParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateVisibilityConfigObservation) DeepCopyInto(out *PrivateVisibilityConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateVisibilityConfigObservation.
func (in *PrivateVisibilityConfigObservation) DeepCopy() *PrivateVisibilityConfigObservation {
	if in == nil {
		return nil
	}
	out := new(PrivateVisibilityConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateVisibilityConfigParameters) DeepCopyInto(out *PrivateVisibilityConfigParameters) {
	*out = *in
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworksParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateVisibilityConfigParameters.
func (in *PrivateVisibilityConfigParameters) DeepCopy() *PrivateVisibilityConfigParameters {
	if in == nil {
		return nil
	}
	out := new(PrivateVisibilityConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSet) DeepCopyInto(out *RecordSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSet.
func (in *RecordSet) DeepCopy() *RecordSet {
	if in == nil {
		return nil
	}
	out := new(RecordSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecordSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetList) DeepCopyInto(out *RecordSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RecordSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetList.
func (in *RecordSetList) DeepCopy() *RecordSetList {
	if in == nil {
		return nil
	}
	out := new(RecordSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecordSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetObservation) DeepCopyInto(out *RecordSetObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetObservation.
func (in *RecordSetObservation) DeepCopy() *RecordSetObservation {
	if in == nil {
		return nil
	}
	out := new(RecordSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetParameters) DeepCopyInto(out *RecordSetParameters) {
	*out = *in
	if in.ManagedZone != nil {
		in, out := &in.ManagedZone, &out.ManagedZone
		*out = new(string)
		**out = **in
	}
	if in.ManagedZoneRef != nil {
		in, out := &in.ManagedZoneRef, &out.ManagedZoneRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedZoneSelector != nil {
		in, out := &in.ManagedZoneSelector, &out.ManagedZoneSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameRef != nil {
		in, out := &in.NameRef, &out.NameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.RoutingPolicy != nil {
		in, out := &in.RoutingPolicy, &out.RoutingPolicy
		*out = make([]RoutingPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetParameters.
func (in *RecordSetParameters) DeepCopy() *RecordSetParameters {
	if in == nil {
		return nil
	}
	out := new(RecordSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetSpec) DeepCopyInto(out *RecordSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetSpec.
func (in *RecordSetSpec) DeepCopy() *RecordSetSpec {
	if in == nil {
		return nil
	}
	out := new(RecordSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetStatus) DeepCopyInto(out *RecordSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetStatus.
func (in *RecordSetStatus) DeepCopy() *RecordSetStatus {
	if in == nil {
		return nil
	}
	out := new(RecordSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyObservation) DeepCopyInto(out *RoutingPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyObservation.
func (in *RoutingPolicyObservation) DeepCopy() *RoutingPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyParameters) DeepCopyInto(out *RoutingPolicyParameters) {
	*out = *in
	if in.Geo != nil {
		in, out := &in.Geo, &out.Geo
		*out = make([]GeoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Wrr != nil {
		in, out := &in.Wrr, &out.Wrr
		*out = make([]WrrParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyParameters.
func (in *RoutingPolicyParameters) DeepCopy() *RoutingPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetNameServersObservation) DeepCopyInto(out *TargetNameServersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetNameServersObservation.
func (in *TargetNameServersObservation) DeepCopy() *TargetNameServersObservation {
	if in == nil {
		return nil
	}
	out := new(TargetNameServersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetNameServersParameters) DeepCopyInto(out *TargetNameServersParameters) {
	*out = *in
	if in.ForwardingPath != nil {
		in, out := &in.ForwardingPath, &out.ForwardingPath
		*out = new(string)
		**out = **in
	}
	if in.IPv4Address != nil {
		in, out := &in.IPv4Address, &out.IPv4Address
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetNameServersParameters.
func (in *TargetNameServersParameters) DeepCopy() *TargetNameServersParameters {
	if in == nil {
		return nil
	}
	out := new(TargetNameServersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetNetworkObservation) DeepCopyInto(out *TargetNetworkObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetNetworkObservation.
func (in *TargetNetworkObservation) DeepCopy() *TargetNetworkObservation {
	if in == nil {
		return nil
	}
	out := new(TargetNetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetNetworkParameters) DeepCopyInto(out *TargetNetworkParameters) {
	*out = *in
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	if in.NetworkURLRef != nil {
		in, out := &in.NetworkURLRef, &out.NetworkURLRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkURLSelector != nil {
		in, out := &in.NetworkURLSelector, &out.NetworkURLSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetNetworkParameters.
func (in *TargetNetworkParameters) DeepCopy() *TargetNetworkParameters {
	if in == nil {
		return nil
	}
	out := new(TargetNetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WrrObservation) DeepCopyInto(out *WrrObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WrrObservation.
func (in *WrrObservation) DeepCopy() *WrrObservation {
	if in == nil {
		return nil
	}
	out := new(WrrObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WrrParameters) DeepCopyInto(out *WrrParameters) {
	*out = *in
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WrrParameters.
func (in *WrrParameters) DeepCopy() *WrrParameters {
	if in == nil {
		return nil
	}
	out := new(WrrParameters)
	in.DeepCopyInto(out)
	return out
}
