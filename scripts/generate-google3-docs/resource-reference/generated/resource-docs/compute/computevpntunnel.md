{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}ComputeVPNTunnel{% endblock %}
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
<td>Compute Engine</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/compute/docs/">/compute/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1.vpnTunnels</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/compute/docs/reference/rest/v1/vpnTunnels">/compute/docs/reference/rest/v1/vpnTunnels</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpcomputevpntunnel<br>gcpcomputevpntunnels<br>computevpntunnel</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>compute.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>computevpntunnels.compute.cnrm.cloud.google.com</td>
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
description: string
ikeVersion: integer
localTrafficSelector:
- string
peerExternalGatewayInterface: integer
peerExternalGatewayRef:
  external: string
  name: string
  namespace: string
peerGCPGatewayRef:
  external: string
  name: string
  namespace: string
peerIp: string
region: string
remoteTrafficSelector:
- string
resourceID: string
routerRef:
  external: string
  name: string
  namespace: string
sharedSecret:
  value: string
  valueFrom:
    secretKeyRef:
      key: string
      name: string
targetVPNGatewayRef:
  external: string
  name: string
  namespace: string
vpnGatewayInterface: integer
vpnGatewayRef:
  external: string
  name: string
  namespace: string
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
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. An optional description of this resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>ikeVersion</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. IKE protocol version to use when establishing the VPN tunnel with
peer VPN gateway.
Acceptable IKE versions are 1 or 2. Default version is 2.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>localTrafficSelector</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Immutable. Local traffic selector to use when establishing the VPN tunnel with
peer VPN gateway. The value should be a CIDR formatted string,
for example '192.168.0.0/16'. The ranges should be disjoint.
Only IPv4 is supported.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>localTrafficSelector[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>peerExternalGatewayInterface</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. The interface ID of the external VPN gateway to which this VPN tunnel is connected.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>peerExternalGatewayRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The peer side external VPN gateway to which this VPN tunnel
is connected.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>peerExternalGatewayRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `ComputeExternalVPNGateway` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>peerExternalGatewayRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>peerExternalGatewayRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>peerGCPGatewayRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The peer side HA GCP VPN gateway to which this VPN tunnel is
connected. If provided, the VPN tunnel will automatically use the
same VPN gateway interface ID in the peer GCP VPN gateway.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>peerGCPGatewayRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `ComputeVPNGateway` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>peerGCPGatewayRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>peerGCPGatewayRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>peerIp</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. IP address of the peer VPN gateway. Only IPv4 is supported.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>region</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>remoteTrafficSelector</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Immutable. Remote traffic selector to use when establishing the VPN tunnel with
peer VPN gateway. The value should be a CIDR formatted string,
for example '192.168.0.0/16'. The ranges should be disjoint.
Only IPv4 is supported.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>remoteTrafficSelector[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
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
    <tr>
        <td>
            <p><code>routerRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The router to be used for dynamic routing.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>routerRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `ComputeRouter` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>routerRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>routerRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>sharedSecret</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. Shared secret used to set the secure session between the Cloud VPN
gateway and the peer VPN gateway.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>sharedSecret.value</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Value of the field. Cannot be used if 'valueFrom' is specified.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>sharedSecret.valueFrom</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Source for the field's value. Cannot be used if 'value' is specified.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>sharedSecret.valueFrom.secretKeyRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Reference to a value with the given key in the given Secret in the resource's namespace.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>sharedSecret.valueFrom.secretKeyRef.key</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Key that identifies the value to be extracted.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>sharedSecret.valueFrom.secretKeyRef.name</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the Secret to extract a value from.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetVPNGatewayRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The ComputeTargetVPNGateway with which this VPN tunnel is
associated.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetVPNGatewayRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `ComputeTargetVPNGateway` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetVPNGatewayRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetVPNGatewayRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>vpnGatewayInterface</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. The interface ID of the VPN gateway with which this VPN tunnel is associated.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>vpnGatewayRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The ComputeVPNGateway with which this VPN tunnel is associated.
This must be used if a High Availability VPN gateway resource is
created.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>vpnGatewayRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `ComputeVPNGateway` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>vpnGatewayRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>vpnGatewayRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>


<p>{% verbatim %}* Field is required when parent field is specified{% endverbatim %}</p>


### Status
#### Schema
```yaml
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
creationTimestamp: string
detailedStatus: string
labelFingerprint: string
observedGeneration: integer
selfLink: string
sharedSecretHash: string
tunnelId: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
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
        <td><code>creationTimestamp</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Creation timestamp in RFC3339 text format.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>detailedStatus</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Detailed status message for the VPN tunnel.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>labelFingerprint</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The fingerprint used for optimistic locking of this resource.  Used
internally during updates.{% endverbatim %}</p>
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
        <td><code>selfLink</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>sharedSecretHash</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Hash of the shared secret.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>tunnelId</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The unique identifier for the resource. This identifier is defined by the server.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
```yaml
# Copyright 2020 Google LLC
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

apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeVPNTunnel
metadata:
  name: computevpntunnel-sample
  labels:
    foo: bar
spec:
  peerIp: "15.0.0.120"
  region: us-central1
  sharedSecret:
    valueFrom:
      secretKeyRef:
        name: computevpntunnel-dep
        key: sharedSecret
  targetVPNGatewayRef:
    name: computevpntunnel-dep
  localTrafficSelector:
   - "192.168.0.0/16"
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeAddress
metadata:
  name: computevpntunnel-dep
  labels:
    label-one: "value-one"
spec:
  location: us-central1
  description: "a test regional address"
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeForwardingRule
metadata:
  labels:
    label-one: "value-one"
  name: computevpntunnel-dep1
spec:
  description: "A regional forwarding rule"
  target:
    targetVPNGatewayRef:
      name: computevpntunnel-dep
  ipProtocol: "ESP"
  location: us-central1
  ipAddress:
    addressRef:
      name: computevpntunnel-dep
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeForwardingRule
metadata:
  labels:
    label-one: "value-one"
  name: computevpntunnel-dep2
spec:
  description: "A regional forwarding rule"
  target:
    targetVPNGatewayRef:
      name: computevpntunnel-dep
  ipProtocol: "UDP"
  portRange: "500"
  location: us-central1
  ipAddress:
    addressRef:
      name: computevpntunnel-dep
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeForwardingRule
metadata:
  labels:
    label-one: "value-one"
  name: computevpntunnel-dep3
spec:
  description: "A regional forwarding rule"
  target:
    targetVPNGatewayRef:
      name: computevpntunnel-dep
  ipProtocol: "UDP"
  portRange: "4500"
  location: us-central1
  ipAddress:
    addressRef:
      name: computevpntunnel-dep
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: computevpntunnel-dep
spec:
  routingMode: REGIONAL
  autoCreateSubnetworks: false
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeTargetVPNGateway
metadata:
  name: computevpntunnel-dep
spec:
  description: a test target vpn gateway
  region: us-central1
  networkRef:
    name: computevpntunnel-dep
---
apiVersion: v1
kind: Secret
metadata:
  name: computevpntunnel-dep
stringData:
  sharedSecret: "a secret message"
```


{% endblock %}