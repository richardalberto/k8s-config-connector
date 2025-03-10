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

type ImageGuestOsFeatures struct {
	/* Immutable. The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options. Possible values: ["MULTI_IP_SUBNET", "SECURE_BOOT", "SEV_CAPABLE", "UEFI_COMPATIBLE", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "GVNIC", "SEV_LIVE_MIGRATABLE"]. */
	Type string `json:"type"`
}

type ImageImageEncryptionKey struct {
	/* The self link of the encryption key that is stored in Google Cloud
	KMS. */
	// +optional
	KmsKeySelfLinkRef *v1alpha1.ResourceRef `json:"kmsKeySelfLinkRef,omitempty"`

	/* The service account being used for the encryption request for the
	given KMS key. If absent, the Compute Engine default service account
	is used. */
	// +optional
	KmsKeyServiceAccountRef *v1alpha1.ResourceRef `json:"kmsKeyServiceAccountRef,omitempty"`
}

type ImageRawDisk struct {
	/* Immutable. The format used to encode and transmit the block device, which
	should be TAR. This is just a container and transmission format
	and not a runtime format. Provided by the client when the disk
	image is created. Default value: "TAR" Possible values: ["TAR"]. */
	// +optional
	ContainerType *string `json:"containerType,omitempty"`

	/* Immutable. An optional SHA1 checksum of the disk image before unpackaging.
	This is provided by the client when the disk image is created. */
	// +optional
	Sha1 *string `json:"sha1,omitempty"`

	/* Immutable. The full Google Cloud Storage URL where disk storage is stored
	You must provide either this property or the sourceDisk property
	but not both. */
	Source string `json:"source"`
}

type ComputeImageSpec struct {
	/* Immutable. An optional description of this resource. Provide this property when
	you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The source disk to create this image based on.
	You must provide either this property or the
	rawDisk.source property but not both to create an image. */
	// +optional
	DiskRef *v1alpha1.ResourceRef `json:"diskRef,omitempty"`

	/* Immutable. Size of the image when restored onto a persistent disk (in GB). */
	// +optional
	DiskSizeGb *int `json:"diskSizeGb,omitempty"`

	/* Immutable. The name of the image family to which this image belongs. You can
	create disks by specifying an image family instead of a specific
	image name. The image family always returns its latest image that is
	not deprecated. The name of the image family must comply with
	RFC1035. */
	// +optional
	Family *string `json:"family,omitempty"`

	/* Immutable. A list of features to enable on the guest operating system.
	Applicable only for bootable images. */
	// +optional
	GuestOsFeatures []ImageGuestOsFeatures `json:"guestOsFeatures,omitempty"`

	/* Immutable. Encrypts the image using a customer-supplied encryption key.

	After you encrypt an image with a customer-supplied key, you must
	provide the same key if you use the image later (e.g. to create a
	disk from the image). */
	// +optional
	ImageEncryptionKey *ImageImageEncryptionKey `json:"imageEncryptionKey,omitempty"`

	/* Immutable. Any applicable license URI. */
	// +optional
	Licenses []string `json:"licenses,omitempty"`

	/* Immutable. The parameters of the raw disk image. */
	// +optional
	RawDisk *ImageRawDisk `json:"rawDisk,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The source image used to create this image. */
	// +optional
	SourceImageRef *v1alpha1.ResourceRef `json:"sourceImageRef,omitempty"`

	/* The source snapshot used to create this image. */
	// +optional
	SourceSnapshotRef *v1alpha1.ResourceRef `json:"sourceSnapshotRef,omitempty"`
}

type ComputeImageStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeImage's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Size of the image tar.gz archive stored in Google Cloud Storage (in
	bytes). */
	// +optional
	ArchiveSizeBytes *int `json:"archiveSizeBytes,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* The fingerprint used for optimistic locking of this resource. Used
	internally during updates. */
	// +optional
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeImage is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeImageSpec   `json:"spec,omitempty"`
	Status ComputeImageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeImageList contains a list of ComputeImage
type ComputeImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeImage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeImage{}, &ComputeImageList{})
}
