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

type TunnelIAMMemberConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TunnelIAMMemberConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TunnelIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type TunnelIAMMemberInitParameters struct {
	Condition []TunnelIAMMemberConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type TunnelIAMMemberObservation struct {
	Condition []TunnelIAMMemberConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type TunnelIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []TunnelIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// TunnelIAMMemberSpec defines the desired state of TunnelIAMMember
type TunnelIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TunnelIAMMemberParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TunnelIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// TunnelIAMMemberStatus defines the observed state of TunnelIAMMember.
type TunnelIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TunnelIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TunnelIAMMember is the Schema for the TunnelIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TunnelIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TunnelIAMMemberSpec   `json:"spec"`
	Status            TunnelIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TunnelIAMMemberList contains a list of TunnelIAMMembers
type TunnelIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TunnelIAMMember `json:"items"`
}

// Repository type metadata.
var (
	TunnelIAMMember_Kind             = "TunnelIAMMember"
	TunnelIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TunnelIAMMember_Kind}.String()
	TunnelIAMMember_KindAPIVersion   = TunnelIAMMember_Kind + "." + CRDGroupVersion.String()
	TunnelIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(TunnelIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&TunnelIAMMember{}, &TunnelIAMMemberList{})
}
