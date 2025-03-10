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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CertificateAuthorizationAttemptInfo struct {
	/* Human readable explanation for reaching the state. Provided to help
	address the configuration issues.
	Not guaranteed to be stable. For programmatic access use 'failure_reason' field. */
	// +optional
	Details *string `json:"details,omitempty"`

	/* Domain name of the authorization attempt. */
	// +optional
	Domain *string `json:"domain,omitempty"`

	/* Reason for failure of the authorization attempt for the domain. */
	// +optional
	FailureReason *string `json:"failureReason,omitempty"`

	/* State of the domain for managed certificate issuance. */
	// +optional
	State *string `json:"state,omitempty"`
}

type CertificateCertificatePem struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *CertificateValueFrom `json:"valueFrom,omitempty"`
}

type CertificateManaged struct {
	/* Detailed state of the latest authorization attempt for each domain
	specified for this Managed Certificate. */
	// +optional
	AuthorizationAttemptInfo []CertificateAuthorizationAttemptInfo `json:"authorizationAttemptInfo,omitempty"`

	/* Immutable. Authorizations that will be used for performing domain authorization. */
	// +optional
	DnsAuthorizations []string `json:"dnsAuthorizations,omitempty"`

	/* Immutable. The domains for which a managed SSL certificate will be generated.
	Wildcard domains are only supported with DNS challenge resolution. */
	// +optional
	Domains []string `json:"domains,omitempty"`

	/* Information about issues with provisioning this Managed Certificate. */
	// +optional
	ProvisioningIssue []CertificateProvisioningIssue `json:"provisioningIssue,omitempty"`

	/* A state of this Managed Certificate. */
	// +optional
	State *string `json:"state,omitempty"`
}

type CertificatePemPrivateKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *CertificateValueFrom `json:"valueFrom,omitempty"`
}

type CertificatePrivateKeyPem struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *CertificateValueFrom `json:"valueFrom,omitempty"`
}

type CertificateProvisioningIssue struct {
	/* Human readable explanation about the issue. Provided to help address
	the configuration issues.
	Not guaranteed to be stable. For programmatic access use 'reason' field. */
	// +optional
	Details *string `json:"details,omitempty"`

	/* Reason for provisioning failures. */
	// +optional
	Reason *string `json:"reason,omitempty"`
}

type CertificateSelfManaged struct {
	/* DEPRECATED. Deprecated in favor of `pem_certificate`. **Deprecated** The certificate chain in PEM-encoded form.

	Leaf certificate comes first, followed by intermediate ones if any. */
	// +optional
	CertificatePem *CertificateCertificatePem `json:"certificatePem,omitempty"`

	/* The certificate chain in PEM-encoded form.

	Leaf certificate comes first, followed by intermediate ones if any. */
	// +optional
	PemCertificate *string `json:"pemCertificate,omitempty"`

	/* The private key of the leaf certificate in PEM-encoded form. */
	// +optional
	PemPrivateKey *CertificatePemPrivateKey `json:"pemPrivateKey,omitempty"`

	/* DEPRECATED. Deprecated in favor of `pem_private_key`. **Deprecated** The private key of the leaf certificate in PEM-encoded form. */
	// +optional
	PrivateKeyPem *CertificatePrivateKeyPem `json:"privateKeyPem,omitempty"`
}

type CertificateValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type CertificateManagerCertificateSpec struct {
	/* A human-readable description of the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Configuration and state of a Managed Certificate.
	Certificate Manager provisions and renews Managed Certificates
	automatically, for as long as it's authorized to do so. */
	// +optional
	Managed *CertificateManaged `json:"managed,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The scope of the certificate.

	DEFAULT: Certificates with default scope are served from core Google data centers.
	If unsure, choose this option.

	EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
	served from non-core Google data centers.
	Currently allowed only for managed certificates. */
	// +optional
	Scope *string `json:"scope,omitempty"`

	/* Immutable. Certificate data for a SelfManaged Certificate.
	SelfManaged Certificates are uploaded by the user. Updating such
	certificates before they expire remains the user's responsibility. */
	// +optional
	SelfManaged *CertificateSelfManaged `json:"selfManaged,omitempty"`
}

type CertificateManagerCertificateStatus struct {
	/* Conditions represent the latest available observations of the
	   CertificateManagerCertificate's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateManagerCertificate is the Schema for the certificatemanager API
// +k8s:openapi-gen=true
type CertificateManagerCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateManagerCertificateSpec   `json:"spec,omitempty"`
	Status CertificateManagerCertificateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateManagerCertificateList contains a list of CertificateManagerCertificate
type CertificateManagerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateManagerCertificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateManagerCertificate{}, &CertificateManagerCertificateList{})
}
