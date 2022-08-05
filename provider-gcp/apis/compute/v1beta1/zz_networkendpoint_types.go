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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkEndpointObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NetworkEndpointParameters struct {

	// IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// The name for a specific VM instance that the IP address belongs to.
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone of network endpoint group.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Instance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// The network endpoint group this endpoint is part of.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.NetworkEndpointGroup
	// +kubebuilder:validation:Optional
	NetworkEndpointGroup *string `json:"networkEndpointGroup,omitempty" tf:"network_endpoint_group,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkEndpointGroupRef *v1.Reference `json:"networkEndpointGroupRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkEndpointGroupSelector *v1.Selector `json:"networkEndpointGroupSelector,omitempty" tf:"-"`

	// Port number of network endpoint.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Zone where the containing network endpoint group is located.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// NetworkEndpointSpec defines the desired state of NetworkEndpoint
type NetworkEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkEndpointParameters `json:"forProvider"`
}

// NetworkEndpointStatus defines the observed state of NetworkEndpoint.
type NetworkEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkEndpoint is the Schema for the NetworkEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NetworkEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkEndpointSpec   `json:"spec"`
	Status            NetworkEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkEndpointList contains a list of NetworkEndpoints
type NetworkEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkEndpoint `json:"items"`
}

// Repository type metadata.
var (
	NetworkEndpoint_Kind             = "NetworkEndpoint"
	NetworkEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkEndpoint_Kind}.String()
	NetworkEndpoint_KindAPIVersion   = NetworkEndpoint_Kind + "." + CRDGroupVersion.String()
	NetworkEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(NetworkEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkEndpoint{}, &NetworkEndpointList{})
}
