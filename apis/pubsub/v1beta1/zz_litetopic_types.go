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

type CapacityInitParameters struct {

	// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	PublishMibPerSec *float64 `json:"publishMibPerSec,omitempty" tf:"publish_mib_per_sec,omitempty"`

	// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	SubscribeMibPerSec *float64 `json:"subscribeMibPerSec,omitempty" tf:"subscribe_mib_per_sec,omitempty"`
}

type CapacityObservation struct {

	// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	PublishMibPerSec *float64 `json:"publishMibPerSec,omitempty" tf:"publish_mib_per_sec,omitempty"`

	// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	SubscribeMibPerSec *float64 `json:"subscribeMibPerSec,omitempty" tf:"subscribe_mib_per_sec,omitempty"`
}

type CapacityParameters struct {

	// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	// +kubebuilder:validation:Optional
	PublishMibPerSec *float64 `json:"publishMibPerSec" tf:"publish_mib_per_sec,omitempty"`

	// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	// +kubebuilder:validation:Optional
	SubscribeMibPerSec *float64 `json:"subscribeMibPerSec" tf:"subscribe_mib_per_sec,omitempty"`
}

type LiteTopicInitParameters struct {

	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig []PartitionConfigInitParameters `json:"partitionConfig,omitempty" tf:"partition_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the pubsub lite topic.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The settings for this topic's Reservation usage.
	// Structure is documented below.
	ReservationConfig []ReservationConfigInitParameters `json:"reservationConfig,omitempty" tf:"reservation_config,omitempty"`

	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig []RetentionConfigInitParameters `json:"retentionConfig,omitempty" tf:"retention_config,omitempty"`
}

type LiteTopicObservation struct {

	// an identifier for the resource with format projects/{{project}}/locations/{{zone}}/topics/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig []PartitionConfigObservation `json:"partitionConfig,omitempty" tf:"partition_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the pubsub lite topic.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The settings for this topic's Reservation usage.
	// Structure is documented below.
	ReservationConfig []ReservationConfigObservation `json:"reservationConfig,omitempty" tf:"reservation_config,omitempty"`

	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig []RetentionConfigObservation `json:"retentionConfig,omitempty" tf:"retention_config,omitempty"`

	// The zone of the pubsub lite topic.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type LiteTopicParameters struct {

	// The settings for this topic's partitions.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PartitionConfig []PartitionConfigParameters `json:"partitionConfig,omitempty" tf:"partition_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the pubsub lite topic.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The settings for this topic's Reservation usage.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ReservationConfig []ReservationConfigParameters `json:"reservationConfig,omitempty" tf:"reservation_config,omitempty"`

	// The settings for a topic's message retention.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RetentionConfig []RetentionConfigParameters `json:"retentionConfig,omitempty" tf:"retention_config,omitempty"`

	// The zone of the pubsub lite topic.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type PartitionConfigInitParameters struct {

	// The capacity configuration.
	// Structure is documented below.
	Capacity []CapacityInitParameters `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The number of partitions in the topic. Must be at least 1.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type PartitionConfigObservation struct {

	// The capacity configuration.
	// Structure is documented below.
	Capacity []CapacityObservation `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The number of partitions in the topic. Must be at least 1.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type PartitionConfigParameters struct {

	// The capacity configuration.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Capacity []CapacityParameters `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The number of partitions in the topic. Must be at least 1.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`
}

type ReservationConfigInitParameters struct {
}

type ReservationConfigObservation struct {

	// The Reservation to use for this topic's throughput capacity.
	ThroughputReservation *string `json:"throughputReservation,omitempty" tf:"throughput_reservation,omitempty"`
}

type ReservationConfigParameters struct {

	// The Reservation to use for this topic's throughput capacity.
	// +crossplane:generate:reference:type=LiteReservation
	// +kubebuilder:validation:Optional
	ThroughputReservation *string `json:"throughputReservation,omitempty" tf:"throughput_reservation,omitempty"`

	// Reference to a LiteReservation to populate throughputReservation.
	// +kubebuilder:validation:Optional
	ThroughputReservationRef *v1.Reference `json:"throughputReservationRef,omitempty" tf:"-"`

	// Selector for a LiteReservation to populate throughputReservation.
	// +kubebuilder:validation:Optional
	ThroughputReservationSelector *v1.Selector `json:"throughputReservationSelector,omitempty" tf:"-"`
}

type RetentionConfigInitParameters struct {

	// The provisioned storage, in bytes, per partition. If the number of bytes stored
	// in any of the topic's partitions grows beyond this value, older messages will be
	// dropped to make room for newer ones, regardless of the value of period.
	PerPartitionBytes *string `json:"perPartitionBytes,omitempty" tf:"per_partition_bytes,omitempty"`

	// How long a published message is retained. If unset, messages will be retained as
	// long as the bytes retained for each partition is below perPartitionBytes. A
	// duration in seconds with up to nine fractional digits, terminated by 's'.
	// Example: "3.5s".
	Period *string `json:"period,omitempty" tf:"period,omitempty"`
}

type RetentionConfigObservation struct {

	// The provisioned storage, in bytes, per partition. If the number of bytes stored
	// in any of the topic's partitions grows beyond this value, older messages will be
	// dropped to make room for newer ones, regardless of the value of period.
	PerPartitionBytes *string `json:"perPartitionBytes,omitempty" tf:"per_partition_bytes,omitempty"`

	// How long a published message is retained. If unset, messages will be retained as
	// long as the bytes retained for each partition is below perPartitionBytes. A
	// duration in seconds with up to nine fractional digits, terminated by 's'.
	// Example: "3.5s".
	Period *string `json:"period,omitempty" tf:"period,omitempty"`
}

type RetentionConfigParameters struct {

	// The provisioned storage, in bytes, per partition. If the number of bytes stored
	// in any of the topic's partitions grows beyond this value, older messages will be
	// dropped to make room for newer ones, regardless of the value of period.
	// +kubebuilder:validation:Optional
	PerPartitionBytes *string `json:"perPartitionBytes" tf:"per_partition_bytes,omitempty"`

	// How long a published message is retained. If unset, messages will be retained as
	// long as the bytes retained for each partition is below perPartitionBytes. A
	// duration in seconds with up to nine fractional digits, terminated by 's'.
	// Example: "3.5s".
	// +kubebuilder:validation:Optional
	Period *string `json:"period,omitempty" tf:"period,omitempty"`
}

// LiteTopicSpec defines the desired state of LiteTopic
type LiteTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LiteTopicParameters `json:"forProvider"`
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
	InitProvider LiteTopicInitParameters `json:"initProvider,omitempty"`
}

// LiteTopicStatus defines the observed state of LiteTopic.
type LiteTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LiteTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LiteTopic is the Schema for the LiteTopics API. A named resource to which messages are sent by publishers.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type LiteTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.partitionConfig) || has(self.initProvider.partitionConfig)",message="partitionConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.retentionConfig) || has(self.initProvider.retentionConfig)",message="retentionConfig is a required parameter"
	Spec   LiteTopicSpec   `json:"spec"`
	Status LiteTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LiteTopicList contains a list of LiteTopics
type LiteTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiteTopic `json:"items"`
}

// Repository type metadata.
var (
	LiteTopic_Kind             = "LiteTopic"
	LiteTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LiteTopic_Kind}.String()
	LiteTopic_KindAPIVersion   = LiteTopic_Kind + "." + CRDGroupVersion.String()
	LiteTopic_GroupVersionKind = CRDGroupVersion.WithKind(LiteTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&LiteTopic{}, &LiteTopicList{})
}
