// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeExternalVPNGateways implements ComputeExternalVPNGatewayInterface
type FakeComputeExternalVPNGateways struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computeexternalvpngatewaysResource = schema.GroupVersionResource{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Resource: "computeexternalvpngateways"}

var computeexternalvpngatewaysKind = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeExternalVPNGateway"}

// Get takes name of the computeExternalVPNGateway, and returns the corresponding computeExternalVPNGateway object, and an error if there is any.
func (c *FakeComputeExternalVPNGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeExternalVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeexternalvpngatewaysResource, c.ns, name), &v1beta1.ComputeExternalVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeExternalVPNGateway), err
}

// List takes label and field selectors, and returns the list of ComputeExternalVPNGateways that match those selectors.
func (c *FakeComputeExternalVPNGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeExternalVPNGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeexternalvpngatewaysResource, computeexternalvpngatewaysKind, c.ns, opts), &v1beta1.ComputeExternalVPNGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeExternalVPNGatewayList{ListMeta: obj.(*v1beta1.ComputeExternalVPNGatewayList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeExternalVPNGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeExternalVPNGateways.
func (c *FakeComputeExternalVPNGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeexternalvpngatewaysResource, c.ns, opts))

}

// Create takes the representation of a computeExternalVPNGateway and creates it.  Returns the server's representation of the computeExternalVPNGateway, and an error, if there is any.
func (c *FakeComputeExternalVPNGateways) Create(ctx context.Context, computeExternalVPNGateway *v1beta1.ComputeExternalVPNGateway, opts v1.CreateOptions) (result *v1beta1.ComputeExternalVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeexternalvpngatewaysResource, c.ns, computeExternalVPNGateway), &v1beta1.ComputeExternalVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeExternalVPNGateway), err
}

// Update takes the representation of a computeExternalVPNGateway and updates it. Returns the server's representation of the computeExternalVPNGateway, and an error, if there is any.
func (c *FakeComputeExternalVPNGateways) Update(ctx context.Context, computeExternalVPNGateway *v1beta1.ComputeExternalVPNGateway, opts v1.UpdateOptions) (result *v1beta1.ComputeExternalVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeexternalvpngatewaysResource, c.ns, computeExternalVPNGateway), &v1beta1.ComputeExternalVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeExternalVPNGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeExternalVPNGateways) UpdateStatus(ctx context.Context, computeExternalVPNGateway *v1beta1.ComputeExternalVPNGateway, opts v1.UpdateOptions) (*v1beta1.ComputeExternalVPNGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeexternalvpngatewaysResource, "status", c.ns, computeExternalVPNGateway), &v1beta1.ComputeExternalVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeExternalVPNGateway), err
}

// Delete takes name of the computeExternalVPNGateway and deletes it. Returns an error if one occurs.
func (c *FakeComputeExternalVPNGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computeexternalvpngatewaysResource, c.ns, name, opts), &v1beta1.ComputeExternalVPNGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeExternalVPNGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeexternalvpngatewaysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeExternalVPNGatewayList{})
	return err
}

// Patch applies the patch and returns the patched computeExternalVPNGateway.
func (c *FakeComputeExternalVPNGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeExternalVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeexternalvpngatewaysResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeExternalVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeExternalVPNGateway), err
}
