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

type InstancegroupNamedPort struct {
	/* The name which the port will be mapped to. */
	Name string `json:"name"`

	/* The port number to map the name to. */
	Port int `json:"port"`
}

type ComputeInstanceGroupSpec struct {
	/* Immutable. An optional textual description of the instance group. */
	// +optional
	Description *string `json:"description,omitempty"`

	/*  */
	// +optional
	Instances []v1alpha1.ResourceRef `json:"instances,omitempty"`

	/* The named port configuration. */
	// +optional
	NamedPort []InstancegroupNamedPort `json:"namedPort,omitempty"`

	/*  */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The zone that this instance group should be created in. */
	Zone string `json:"zone"`
}

type ComputeInstanceGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeInstanceGroup's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* The URI of the created resource. */
	SelfLink string `json:"selfLink,omitempty"`
	/* The number of instances in the group. */
	Size int `json:"size,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInstanceGroup is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeInstanceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInstanceGroupSpec   `json:"spec,omitempty"`
	Status ComputeInstanceGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInstanceGroupList contains a list of ComputeInstanceGroup
type ComputeInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstanceGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstanceGroup{}, &ComputeInstanceGroupList{})
}