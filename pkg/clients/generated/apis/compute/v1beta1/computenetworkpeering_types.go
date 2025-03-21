// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ComputeNetworkPeeringSpec struct {
	/* Whether to export the custom routes to the peer network. Defaults to false. */
	// +optional
	ExportCustomRoutes *bool `json:"exportCustomRoutes,omitempty"`

	/* Immutable. */
	// +optional
	ExportSubnetRoutesWithPublicIp *bool `json:"exportSubnetRoutesWithPublicIp,omitempty"`

	/* Whether to export the custom routes from the peer network. Defaults to false. */
	// +optional
	ImportCustomRoutes *bool `json:"importCustomRoutes,omitempty"`

	/* Immutable. */
	// +optional
	ImportSubnetRoutesWithPublicIp *bool `json:"importSubnetRoutesWithPublicIp,omitempty"`

	NetworkRef v1alpha1.ResourceRef `json:"networkRef"`

	PeerNetworkRef v1alpha1.ResourceRef `json:"peerNetworkRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ComputeNetworkPeeringStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeNetworkPeering's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* State for the peering, either ACTIVE or INACTIVE. The peering is ACTIVE when there's a matching configuration in the peer network. */
	// +optional
	State *string `json:"state,omitempty"`

	/* Details about the current state of the peering. */
	// +optional
	StateDetails *string `json:"stateDetails,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeNetworkPeering is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeNetworkPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkPeeringSpec   `json:"spec,omitempty"`
	Status ComputeNetworkPeeringStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeNetworkPeeringList contains a list of ComputeNetworkPeering
type ComputeNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetworkPeering `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetworkPeering{}, &ComputeNetworkPeeringList{})
}
