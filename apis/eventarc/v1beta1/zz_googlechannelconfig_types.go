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

type GoogleChannelConfigInitParameters struct {

	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern projects/*/locations/*/keyRings/*/cryptoKeys/*.
	CryptoKeyName *string `json:"cryptoKeyName,omitempty" tf:"crypto_key_name,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type GoogleChannelConfigObservation struct {

	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern projects/*/locations/*/keyRings/*/cryptoKeys/*.
	CryptoKeyName *string `json:"cryptoKeyName,omitempty" tf:"crypto_key_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/googleChannelConfig
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Output only. The last-modified time.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type GoogleChannelConfigParameters struct {

	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern projects/*/locations/*/keyRings/*/cryptoKeys/*.
	// +kubebuilder:validation:Optional
	CryptoKeyName *string `json:"cryptoKeyName,omitempty" tf:"crypto_key_name,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// GoogleChannelConfigSpec defines the desired state of GoogleChannelConfig
type GoogleChannelConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GoogleChannelConfigParameters `json:"forProvider"`
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
	InitProvider GoogleChannelConfigInitParameters `json:"initProvider,omitempty"`
}

// GoogleChannelConfigStatus defines the observed state of GoogleChannelConfig.
type GoogleChannelConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GoogleChannelConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GoogleChannelConfig is the Schema for the GoogleChannelConfigs API. The Eventarc GoogleChannelConfig resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type GoogleChannelConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleChannelConfigSpec   `json:"spec"`
	Status            GoogleChannelConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GoogleChannelConfigList contains a list of GoogleChannelConfigs
type GoogleChannelConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GoogleChannelConfig `json:"items"`
}

// Repository type metadata.
var (
	GoogleChannelConfig_Kind             = "GoogleChannelConfig"
	GoogleChannelConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GoogleChannelConfig_Kind}.String()
	GoogleChannelConfig_KindAPIVersion   = GoogleChannelConfig_Kind + "." + CRDGroupVersion.String()
	GoogleChannelConfig_GroupVersionKind = CRDGroupVersion.WithKind(GoogleChannelConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&GoogleChannelConfig{}, &GoogleChannelConfigList{})
}
