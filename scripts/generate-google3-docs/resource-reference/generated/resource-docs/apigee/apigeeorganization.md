{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}ApigeeOrganization{% endblock %}
{% block body %}



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
<td>Apigee</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/apigee/docs/">/apigee/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>
organizations
</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td>
<a href="/apigee/docs/reference/apis/apigee/rest/v1/organizations">/apigee/docs/reference/apis/apigee/rest/v1/organizations</a>
</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpapigeeorganization<br>gcpapigeeorganizations<br>apigeeorganization</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>apigee.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>apigeeorganizations.apigee.cnrm.cloud.google.com</td>
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
        <td><code>cnrm.cloud.google.com/state-into-spec</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
```yaml
addonsConfig:
  advancedApiOpsConfig:
    enabled: boolean
  monetizationConfig:
    enabled: boolean
analyticsRegion: string
authorizedNetworkRef:
  external: string
  name: string
  namespace: string
description: string
displayName: string
projectRef:
  external: string
  name: string
  namespace: string
properties:
  string: string
resourceID: string
runtimeDatabaseEncryptionKeyRef:
  external: string
  name: string
  namespace: string
runtimeType: string
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
            <p><code>addonsConfig</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Addon configurations of the Apigee organization.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>addonsConfig.advancedApiOpsConfig</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Configuration for the Advanced API Ops add-on.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>addonsConfig.advancedApiOpsConfig.enabled</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Flag that specifies whether the Advanced API Ops add-on is enabled.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>addonsConfig.monetizationConfig</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Configuration for the Monetization add-on.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>addonsConfig.monetizationConfig.enabled</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Flag that specifies whether the Monetization add-on is enabled.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>analyticsRegion</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Required. Primary GCP region for analytics data storage. For valid values, see (https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>authorizedNetworkRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>authorizedNetworkRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See (https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.

Allowed value: The Google Cloud resource name of a `ComputeNetwork` resource (format: `projects/{{project}}/global/networks/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>authorizedNetworkRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>authorizedNetworkRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Description of the Apigee organization.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>displayName</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Display name for the Apigee organization.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. The Project that this resource belongs to.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Required. Name of the GCP project in which to associate the Apigee organization. Pass the information as a query parameter using the following structure in your request: projects/<project> Authorization requires the following IAM permission on the specified resource parent: apigee.organizations.create

Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>properties</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">map (key: string, value: string)</code></p>
            <p>{% verbatim %}Properties defined in the Apigee organization profile.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>runtimeDatabaseEncryptionKeyRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>runtimeDatabaseEncryptionKeyRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. Required when (#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.

Allowed value: The Google Cloud resource name of a `KMSCryptoKey` resource (format: `{{selfLink}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>runtimeDatabaseEncryptionKeyRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>runtimeDatabaseEncryptionKeyRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>runtimeType</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Required. Runtime type of the Apigee organization based on the Apigee subscription purchased. Possible values: RUNTIME_TYPE_UNSPECIFIED, CLOUD, HYBRID{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



### Status
#### Schema
```yaml
billingType: string
caCertificate: string
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
createdAt: integer
environments:
- string
expiresAt: integer
lastModifiedAt: integer
observedGeneration: integer
projectId: string
state: string
subscriptionType: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>billingType</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. Billing type of the Apigee organization. See (https://cloud.google.com/apigee/pricing). Possible values: BILLING_TYPE_UNSPECIFIED, SUBSCRIPTION, EVALUATION{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>caCertificate</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when (#RuntimeType) is `CLOUD`.{% endverbatim %}</p>
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
        <td><code>createdAt</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Output only. Time that the Apigee organization was created in milliseconds since epoch.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>environments</code></td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Output only. List of environments in the Apigee organization.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>environments[]</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>expiresAt</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Output only. Time that the Apigee organization is scheduled for deletion.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>lastModifiedAt</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Output only. Time that the Apigee organization was last modified in milliseconds since epoch.{% endverbatim %}</p>
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
        <td><code>projectId</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. Project ID associated with the Apigee organization.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>state</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. State of the organization. Values other than ACTIVE means the resource is not ready to use. Possible values: SNAPSHOT_STATE_UNSPECIFIED, MISSING, OK_DOCSTORE, OK_SUBMITTED, OK_EXTERNAL, DELETED{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>subscriptionType</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. DEPRECATED: This will eventually be replaced by BillingType. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased). See (https://cloud.google.com/apigee/pricing/). Possible values: SUBSCRIPTION_TYPE_UNSPECIFIED, PAID, TRIAL{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
```yaml
# Copyright 2022 Google LLC
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

apiVersion: apigee.cnrm.cloud.google.com/v1beta1
kind: ApigeeOrganization
metadata:
  name: apigeeorganization-sample
spec:
  projectRef:
    # Replace ${PROJECT_ID?} with your project ID
    external: "projects/${PROJECT_ID?}"
  displayName: "basic-organization"
  description: "A sample organization"
  properties:
    features.mart.connect.enabled: "false"
    features.hybrid.enabled: "true"
  analyticsRegion: "us-west1"
  authorizedNetworkRef:
    name: "apigeeorganization-dep"
  runtimeType: "CLOUD"
  addonsConfig:
    advancedApiOpsConfig:
      enabled: true
    integrationConfig:
      enabled: false
    monetizationConfig:
      enabled: false
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: "apigeeorganization-dep"
spec:
  autoCreateSubnetworks: false
  description: A sample authorized network for an apigee organization
```


{% endblock %}

