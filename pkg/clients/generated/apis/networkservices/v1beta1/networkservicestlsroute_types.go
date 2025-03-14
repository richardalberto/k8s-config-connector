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

type TlsrouteAction struct {
	/* Required. The destination services to which traffic should be forwarded. At least one destination service is required. */
	Destinations []TlsrouteDestinations `json:"destinations"`
}

type TlsrouteDestinations struct {
	ServiceRef v1alpha1.ResourceRef `json:"serviceRef"`

	/* Optional. Specifies the proportion of requests forwareded to the backend referenced by the service_name field. This is computed as: weight/Sum(weights in destinations) Weights in all destinations does not need to sum up to 100. */
	// +optional
	Weight *int `json:"weight,omitempty"`
}

type TlsrouteMatches struct {
	/* Optional. ALPN (Application-Layer Protocol Negotiation) to match against. Examples: "http/1.1", "h2". At least one of sni_host and alpn is required. Up to 5 alpns across all matches can be set. */
	// +optional
	Alpn []string `json:"alpn,omitempty"`

	/* Optional. SNI (server name indicator) to match against. SNI will be matched against all wildcard domains, i.e. www.example.com will be first matched against www.example.com, then *.example.com, then *.com. Partial wildcards are not supported, and values like *w.example.com are invalid. At least one of sni_host and alpn is required. Up to 5 sni hosts across all matches can be set. */
	// +optional
	SniHost []string `json:"sniHost,omitempty"`
}

type TlsrouteRules struct {
	/* Required. The detailed rule defining how to route matched traffic. */
	Action TlsrouteAction `json:"action"`

	/* Required. RouteMatch defines the predicate used to match requests to a given action. Multiple match types are "OR"ed for evaluation. */
	Matches []TlsrouteMatches `json:"matches"`
}

type NetworkServicesTLSRouteSpec struct {
	/* Optional. A free-text description of the resource. Max length 1024 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	// +optional
	Gateways []v1alpha1.ResourceRef `json:"gateways,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	// +optional
	Meshes []v1alpha1.ResourceRef `json:"meshes,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Required. Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match. */
	Rules []TlsrouteRules `json:"rules"`
}

type NetworkServicesTLSRouteStatus struct {
	/* Conditions represent the latest available observations of the
	   NetworkServicesTLSRoute's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The timestamp when the resource was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Output only. Server-defined URL of this resource */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* Output only. The timestamp when the resource was updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesTLSRoute is the Schema for the networkservices API
// +k8s:openapi-gen=true
type NetworkServicesTLSRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesTLSRouteSpec   `json:"spec,omitempty"`
	Status NetworkServicesTLSRouteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesTLSRouteList contains a list of NetworkServicesTLSRoute
type NetworkServicesTLSRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesTLSRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesTLSRoute{}, &NetworkServicesTLSRouteList{})
}
