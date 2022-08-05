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

type ImageIAMPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ImageIAMPolicyParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Image
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// +kubebuilder:validation:Optional
	ImageRef *v1.Reference `json:"imageRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ImageSelector *v1.Selector `json:"imageSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Image
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("project",false)
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`
}

// ImageIAMPolicySpec defines the desired state of ImageIAMPolicy
type ImageIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageIAMPolicyParameters `json:"forProvider"`
}

// ImageIAMPolicyStatus defines the observed state of ImageIAMPolicy.
type ImageIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ImageIAMPolicy is the Schema for the ImageIAMPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ImageIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageIAMPolicySpec   `json:"spec"`
	Status            ImageIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageIAMPolicyList contains a list of ImageIAMPolicys
type ImageIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	ImageIAMPolicy_Kind             = "ImageIAMPolicy"
	ImageIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageIAMPolicy_Kind}.String()
	ImageIAMPolicy_KindAPIVersion   = ImageIAMPolicy_Kind + "." + CRDGroupVersion.String()
	ImageIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ImageIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageIAMPolicy{}, &ImageIAMPolicyList{})
}
