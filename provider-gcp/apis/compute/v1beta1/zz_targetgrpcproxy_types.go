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

type TargetGRPCProxyObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in
	// this object. This field is used in optimistic locking. This field
	// will be ignored when inserting a TargetGrpcProxy. An up-to-date
	// fingerprint must be provided in order to patch/update the
	// TargetGrpcProxy; otherwise, the request will fail with error
	// 412 conditionNotMet. To see the latest fingerprint, make a get()
	// request to retrieve the TargetGrpcProxy. A base64-encoded string.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/targetGrpcProxies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Server-defined URL with id for the resource.
	SelfLinkWithID *string `json:"selfLinkWithId,omitempty" tf:"self_link_with_id,omitempty"`
}

type TargetGRPCProxyParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.URLMap
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	URLMap *string `json:"urlMap,omitempty" tf:"url_map,omitempty"`

	// Reference to a URLMap in compute to populate urlMap.
	// +kubebuilder:validation:Optional
	URLMapRef *v1.Reference `json:"urlMapRef,omitempty" tf:"-"`

	// Selector for a URLMap in compute to populate urlMap.
	// +kubebuilder:validation:Optional
	URLMapSelector *v1.Selector `json:"urlMapSelector,omitempty" tf:"-"`

	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	// +kubebuilder:validation:Optional
	ValidateForProxyless *bool `json:"validateForProxyless,omitempty" tf:"validate_for_proxyless,omitempty"`
}

// TargetGRPCProxySpec defines the desired state of TargetGRPCProxy
type TargetGRPCProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetGRPCProxyParameters `json:"forProvider"`
}

// TargetGRPCProxyStatus defines the observed state of TargetGRPCProxy.
type TargetGRPCProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetGRPCProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TargetGRPCProxy is the Schema for the TargetGRPCProxys API. Represents a Target gRPC Proxy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TargetGRPCProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetGRPCProxySpec   `json:"spec"`
	Status            TargetGRPCProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetGRPCProxyList contains a list of TargetGRPCProxys
type TargetGRPCProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetGRPCProxy `json:"items"`
}

// Repository type metadata.
var (
	TargetGRPCProxy_Kind             = "TargetGRPCProxy"
	TargetGRPCProxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetGRPCProxy_Kind}.String()
	TargetGRPCProxy_KindAPIVersion   = TargetGRPCProxy_Kind + "." + CRDGroupVersion.String()
	TargetGRPCProxy_GroupVersionKind = CRDGroupVersion.WithKind(TargetGRPCProxy_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetGRPCProxy{}, &TargetGRPCProxyList{})
}
