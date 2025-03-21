{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}GKEHubMembership{% endblock %}
{% block body %}



Note: GKE Hub REST documentation is under construction.

<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>GKE Hub</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/anthos/multicluster-management/connect/overview">/anthos/multicluster-management/connect/overview</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1beta1.projects.locations.memberships</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a class="external" href="https://gkehub.googleapis.com/$discovery/rest?version=v1beta1">https://gkehub.googleapis.com/$discovery/rest?version=v1beta1</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpgkehubmembership<br>gcpgkehubmemberships<br>gkehubmembership</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>gkehub.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>gkehubmemberships.gkehub.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


<tr>
<td>{{product_name_short}} Default Average Reconcile Interval In Seconds</td>
<td>600</td>
</tr>
</tbody>
</table>

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/project-id</code></td>
    </tr>
    <tr>
        <td><code>cnrm.cloud.google.com/state-into-spec</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
```yaml
authority:
  issuer: string
description: string
endpoint:
  gkeCluster:
    resourceRef:
      external: string
      name: string
      namespace: string
  kubernetesResource:
    membershipCrManifest: string
    resourceOptions:
      connectVersion: string
      v1beta1Crd: boolean
externalId: string
infrastructureType: string
location: string
resourceID: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>authority</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>authority.issuer</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://` and be a valid URL with length <2000 characters. If set, then Google will allow valid OIDC tokens from this issuer to authenticate within the workload_identity_pool. OIDC discovery will be performed on this URI to validate tokens from the issuer. Clearing `issuer` disables Workload Identity. `issuer` cannot be directly modified; it must be cleared (and Workload Identity disabled) before using a new issuer (and re-enabling Workload Identity).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Description of this membership, limited to 63 characters. Must match the regex: `*` This field is present for legacy purposes.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. Endpoint information to reach this member.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint.gkeCluster</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. GKE-specific information. Only present if this Membership is a GKE cluster.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint.gkeCluster.resourceRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint.gkeCluster.resourceRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Self-link of the GCP resource for the GKE cluster. For example: //container.googleapis.com/projects/my-project/locations/us-west1-a/clusters/my-cluster Zonal clusters are also supported.

Allowed value: The `selfLink` field of a `ContainerCluster` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint.gkeCluster.resourceRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint.gkeCluster.resourceRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint.kubernetesResource</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. The in-cluster Kubernetes Resources that should be applied for a correctly registered cluster, in the steady state. These resources: * Ensure that the cluster is exclusively registered to one and only one Hub Membership. * Propagate Workload Pool Information available in the Membership Authority field. * Ensure proper initial configuration of default Hub Features.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint.kubernetesResource.membershipCrManifest</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Input only. The YAML representation of the Membership CR. This field is ignored for GKE clusters where Hub can read the CR directly. Callers should provide the CR that is currently present in the cluster during CreateMembership or UpdateMembership, or leave this field empty if none exists. The CR manifest is used to validate the cluster has not been registered with another Membership.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint.kubernetesResource.resourceOptions</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. Options for Kubernetes resource generation.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint.kubernetesResource.resourceOptions.connectVersion</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. The Connect agent version to use for connect_resources. Defaults to the latest GKE Connect version. The version must be a currently supported version, obsolete versions will be rejected.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>endpoint.kubernetesResource.resourceOptions.v1beta1Crd</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Optional. Use `apiextensions/v1beta1` instead of `apiextensions/v1` for CustomResourceDefinition resources. This option should be set for clusters with Kubernetes apiserver versions <1.16.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalId</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. The ID must match the regex: `*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>infrastructureType</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. The infrastructure type this Membership is running on. Possible values: INFRASTRUCTURE_TYPE_UNSPECIFIED, ON_PREM, MULTI_CLOUD{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>location</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The location for the resource{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



### Status
#### Schema
```yaml
authority:
  identityProvider: string
  workloadIdentityPool: string
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
createTime: string
deleteTime: string
endpoint:
  kubernetesMetadata:
    kubernetesApiServerVersion: string
    memoryMb: integer
    nodeCount: integer
    nodeProviderId: string
    updateTime: string
    vcpuCount: integer
  kubernetesResource:
    connectResources:
    - clusterScoped: boolean
      manifest: string
    membershipResources:
    - clusterScoped: boolean
      manifest: string
lastConnectionTime: string
observedGeneration: integer
state:
  code: string
uniqueId: string
updateTime: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>authority</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>authority.identityProvider</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. An identity provider that reflects the `issuer` in the workload identity pool.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>authority.workloadIdentityPool</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. The name of the workload identity pool in which `issuer` will be recognized. There is a single Workload Identity Pool per Hub that is shared between all Memberships that belong to that Hub. For a Hub hosted in: {PROJECT_ID}, the workload pool format is `{PROJECT_ID}.hub.id.goog`, although this is subject to change in newer versions of this API.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>createTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. When the Membership was created.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>deleteTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. When the Membership was deleted.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesMetadata</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Output only. Useful Kubernetes-specific metadata.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesMetadata.kubernetesApiServerVersion</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. Kubernetes API server version string as reported by `/version`.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesMetadata.memoryMb</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Output only. The total memory capacity as reported by the sum of all Kubernetes nodes resources, defined in MB.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesMetadata.nodeCount</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Output only. Node count as reported by Kubernetes nodes resources.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesMetadata.nodeProviderId</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. Node providerID as reported by the first node in the list of nodes on the Kubernetes endpoint. On Kubernetes platforms that support zero-node clusters (like GKE-on-GCP), the node_count will be zero and the node_provider_id will be empty.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesMetadata.updateTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. The time at which these details were last updated. This update_time is different from the Membership-level update_time since EndpointDetails are updated internally for API consumers.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesMetadata.vcpuCount</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Output only. vCPU count as reported by Kubernetes nodes resources.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesResource</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesResource.connectResources</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Output only. The Kubernetes resources for installing the GKE Connect agent This field is only populated in the Membership returned from a successful long-running operation from CreateMembership or UpdateMembership. It is not populated during normal GetMembership or ListMemberships requests. To get the resource manifest after the initial registration, the caller should make a UpdateMembership call with an empty field mask.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesResource.connectResources[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesResource.connectResources[].clusterScoped</code></td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Whether the resource provided in the manifest is `cluster_scoped`. If unset, the manifest is assumed to be namespace scoped. This field is used for REST mapping when applying the resource in a cluster.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesResource.connectResources[].manifest</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}YAML manifest of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesResource.membershipResources</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Output only. Additional Kubernetes resources that need to be applied to the cluster after Membership creation, and after every update. This field is only populated in the Membership returned from a successful long-running operation from CreateMembership or UpdateMembership. It is not populated during normal GetMembership or ListMemberships requests. To get the resource manifest after the initial registration, the caller should make a UpdateMembership call with an empty field mask.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesResource.membershipResources[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesResource.membershipResources[].clusterScoped</code></td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Whether the resource provided in the manifest is `cluster_scoped`. If unset, the manifest is assumed to be namespace scoped. This field is used for REST mapping when applying the resource in a cluster.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>endpoint.kubernetesResource.membershipResources[].manifest</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}YAML manifest of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>lastConnectionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. For clusters using Connect, the timestamp of the most recent connection established with Google Cloud. This time is updated every several minutes, not continuously. For clusters that do not use GKE Connect, or that have never connected successfully, this field will be unset.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>state</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Output only. State of the Membership resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>state.code</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. The current state of the Membership resource. Possible values: CODE_UNSPECIFIED, CREATING, READY, DELETING, UPDATING, SERVICE_UPDATING{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>uniqueId</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. Google-generated UUID for this resource. This is unique across all Membership resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>updateTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. When the Membership was last updated.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
```yaml
# Copyright 2021 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: gkehub.cnrm.cloud.google.com/v1beta1
kind: GKEHubMembership
metadata:
  labels:
    label-one: value-one
  name: gkehubmembership-sample
spec:
  location: global
  authority:
    # Issuer must contain a link to a valid JWT issuer. Your ContainerCluster is one. To use it, replace ${PROJECT_ID?} with your project ID.
    issuer: https://container.googleapis.com/v1/projects/${PROJECT_ID?}/locations/us-central1-a/clusters/gkehubmembership-dep
  description: A sample GKE Hub membership
  endpoint:
    gkeCluster:
      resourceRef:
        name: gkehubmembership-dep
---
apiVersion: container.cnrm.cloud.google.com/v1beta1
kind: ContainerCluster
metadata:
  name: gkehubmembership-dep
spec:
  location: us-central1-a
  initialNodeCount: 1
  workloadIdentityConfig:
    # Workload Identity supports only a single namespace based on your project name.
    # Replace ${PROJECT_ID?} below with your project ID.
    workloadPool: ${PROJECT_ID?}.svc.id.goog
```


{% endblock %}
