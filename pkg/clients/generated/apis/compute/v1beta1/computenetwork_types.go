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

type ComputeNetworkSpec struct {
	/* Immutable. When set to 'true', the network is created in "auto subnet mode" and
	it will create a subnet for each region automatically across the
	'10.128.0.0/9' address range.

	When set to 'false', the network is created in "custom subnet mode" so
	the user can explicitly connect subnetwork resources. */
	// +optional
	AutoCreateSubnetworks *bool `json:"autoCreateSubnetworks,omitempty"`

	/* If set to 'true', default routes ('0.0.0.0/0') will be deleted
	immediately after network creation. Defaults to 'false'. */
	// +optional
	DeleteDefaultRoutesOnCreate *bool `json:"deleteDefaultRoutesOnCreate,omitempty"`

	/* Immutable. An optional description of this resource. The resource must be
	recreated to modify this field. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Enable ULA internal ipv6 on this network. Enabling this feature will assign
	a /48 from google defined ULA prefix fd20::/20. */
	// +optional
	EnableUlaInternalIpv6 *bool `json:"enableUlaInternalIpv6,omitempty"`

	/* Immutable. When enabling ula internal ipv6, caller optionally can specify the /48 range
	they want from the google defined ULA prefix fd20::/20. The input must be a
	valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will
	fail if the speficied /48 is already in used by another resource.
	If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. */
	// +optional
	InternalIpv6Range *string `json:"internalIpv6Range,omitempty"`

	/* Immutable. Maximum Transmission Unit in bytes. The default value is 1460 bytes.
	The minimum value for this field is 1300 and the maximum value is 8896 bytes (jumbo frames).
	Note that packets larger than 1500 bytes (standard Ethernet) can be subject to TCP-MSS clamping or dropped
	with an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or other VPCs
	with varying MTUs. */
	// +optional
	Mtu *int `json:"mtu,omitempty"`

	/* Immutable. Set the order that Firewall Rules and Firewall Policies are evaluated. Needs to be either 'AFTER_CLASSIC_FIREWALL' or 'BEFORE_CLASSIC_FIREWALL' Default 'AFTER_CLASSIC_FIREWALL' Default value: "AFTER_CLASSIC_FIREWALL" Possible values: ["BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL"]. */
	// +optional
	NetworkFirewallPolicyEnforcementOrder *string `json:"networkFirewallPolicyEnforcementOrder,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The network-wide routing mode to use. If set to 'REGIONAL', this
	network's cloud routers will only advertise routes with subnetworks
	of this network in the same region as the router. If set to 'GLOBAL',
	this network's cloud routers will advertise routes with all
	subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"]. */
	// +optional
	RoutingMode *string `json:"routingMode,omitempty"`
}

type ComputeNetworkStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeNetwork's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The gateway address for default routing out of the network. This value
	is selected by GCP. */
	// +optional
	GatewayIpv4 *string `json:"gatewayIpv4,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeNetwork is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkSpec   `json:"spec,omitempty"`
	Status ComputeNetworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeNetworkList contains a list of ComputeNetwork
type ComputeNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetwork{}, &ComputeNetworkList{})
}
