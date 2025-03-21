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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigqueryanalyticshub/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBigQueryAnalyticsHubDataExchanges implements BigQueryAnalyticsHubDataExchangeInterface
type FakeBigQueryAnalyticsHubDataExchanges struct {
	Fake *FakeBigqueryanalyticshubV1alpha1
	ns   string
}

var bigqueryanalyticshubdataexchangesResource = schema.GroupVersionResource{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "bigqueryanalyticshubdataexchanges"}

var bigqueryanalyticshubdataexchangesKind = schema.GroupVersionKind{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "BigQueryAnalyticsHubDataExchange"}

// Get takes name of the bigQueryAnalyticsHubDataExchange, and returns the corresponding bigQueryAnalyticsHubDataExchange object, and an error if there is any.
func (c *FakeBigQueryAnalyticsHubDataExchanges) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BigQueryAnalyticsHubDataExchange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bigqueryanalyticshubdataexchangesResource, c.ns, name), &v1alpha1.BigQueryAnalyticsHubDataExchange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigQueryAnalyticsHubDataExchange), err
}

// List takes label and field selectors, and returns the list of BigQueryAnalyticsHubDataExchanges that match those selectors.
func (c *FakeBigQueryAnalyticsHubDataExchanges) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BigQueryAnalyticsHubDataExchangeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bigqueryanalyticshubdataexchangesResource, bigqueryanalyticshubdataexchangesKind, c.ns, opts), &v1alpha1.BigQueryAnalyticsHubDataExchangeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BigQueryAnalyticsHubDataExchangeList{ListMeta: obj.(*v1alpha1.BigQueryAnalyticsHubDataExchangeList).ListMeta}
	for _, item := range obj.(*v1alpha1.BigQueryAnalyticsHubDataExchangeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bigQueryAnalyticsHubDataExchanges.
func (c *FakeBigQueryAnalyticsHubDataExchanges) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bigqueryanalyticshubdataexchangesResource, c.ns, opts))

}

// Create takes the representation of a bigQueryAnalyticsHubDataExchange and creates it.  Returns the server's representation of the bigQueryAnalyticsHubDataExchange, and an error, if there is any.
func (c *FakeBigQueryAnalyticsHubDataExchanges) Create(ctx context.Context, bigQueryAnalyticsHubDataExchange *v1alpha1.BigQueryAnalyticsHubDataExchange, opts v1.CreateOptions) (result *v1alpha1.BigQueryAnalyticsHubDataExchange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bigqueryanalyticshubdataexchangesResource, c.ns, bigQueryAnalyticsHubDataExchange), &v1alpha1.BigQueryAnalyticsHubDataExchange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigQueryAnalyticsHubDataExchange), err
}

// Update takes the representation of a bigQueryAnalyticsHubDataExchange and updates it. Returns the server's representation of the bigQueryAnalyticsHubDataExchange, and an error, if there is any.
func (c *FakeBigQueryAnalyticsHubDataExchanges) Update(ctx context.Context, bigQueryAnalyticsHubDataExchange *v1alpha1.BigQueryAnalyticsHubDataExchange, opts v1.UpdateOptions) (result *v1alpha1.BigQueryAnalyticsHubDataExchange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bigqueryanalyticshubdataexchangesResource, c.ns, bigQueryAnalyticsHubDataExchange), &v1alpha1.BigQueryAnalyticsHubDataExchange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigQueryAnalyticsHubDataExchange), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBigQueryAnalyticsHubDataExchanges) UpdateStatus(ctx context.Context, bigQueryAnalyticsHubDataExchange *v1alpha1.BigQueryAnalyticsHubDataExchange, opts v1.UpdateOptions) (*v1alpha1.BigQueryAnalyticsHubDataExchange, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bigqueryanalyticshubdataexchangesResource, "status", c.ns, bigQueryAnalyticsHubDataExchange), &v1alpha1.BigQueryAnalyticsHubDataExchange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigQueryAnalyticsHubDataExchange), err
}

// Delete takes name of the bigQueryAnalyticsHubDataExchange and deletes it. Returns an error if one occurs.
func (c *FakeBigQueryAnalyticsHubDataExchanges) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(bigqueryanalyticshubdataexchangesResource, c.ns, name, opts), &v1alpha1.BigQueryAnalyticsHubDataExchange{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBigQueryAnalyticsHubDataExchanges) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bigqueryanalyticshubdataexchangesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BigQueryAnalyticsHubDataExchangeList{})
	return err
}

// Patch applies the patch and returns the patched bigQueryAnalyticsHubDataExchange.
func (c *FakeBigQueryAnalyticsHubDataExchanges) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BigQueryAnalyticsHubDataExchange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bigqueryanalyticshubdataexchangesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BigQueryAnalyticsHubDataExchange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigQueryAnalyticsHubDataExchange), err
}
