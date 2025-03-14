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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudasset/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudAssetOrganizationFeeds implements CloudAssetOrganizationFeedInterface
type FakeCloudAssetOrganizationFeeds struct {
	Fake *FakeCloudassetV1alpha1
	ns   string
}

var cloudassetorganizationfeedsResource = schema.GroupVersionResource{Group: "cloudasset.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "cloudassetorganizationfeeds"}

var cloudassetorganizationfeedsKind = schema.GroupVersionKind{Group: "cloudasset.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudAssetOrganizationFeed"}

// Get takes name of the cloudAssetOrganizationFeed, and returns the corresponding cloudAssetOrganizationFeed object, and an error if there is any.
func (c *FakeCloudAssetOrganizationFeeds) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudAssetOrganizationFeed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudassetorganizationfeedsResource, c.ns, name), &v1alpha1.CloudAssetOrganizationFeed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudAssetOrganizationFeed), err
}

// List takes label and field selectors, and returns the list of CloudAssetOrganizationFeeds that match those selectors.
func (c *FakeCloudAssetOrganizationFeeds) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudAssetOrganizationFeedList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudassetorganizationfeedsResource, cloudassetorganizationfeedsKind, c.ns, opts), &v1alpha1.CloudAssetOrganizationFeedList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudAssetOrganizationFeedList{ListMeta: obj.(*v1alpha1.CloudAssetOrganizationFeedList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudAssetOrganizationFeedList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudAssetOrganizationFeeds.
func (c *FakeCloudAssetOrganizationFeeds) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudassetorganizationfeedsResource, c.ns, opts))

}

// Create takes the representation of a cloudAssetOrganizationFeed and creates it.  Returns the server's representation of the cloudAssetOrganizationFeed, and an error, if there is any.
func (c *FakeCloudAssetOrganizationFeeds) Create(ctx context.Context, cloudAssetOrganizationFeed *v1alpha1.CloudAssetOrganizationFeed, opts v1.CreateOptions) (result *v1alpha1.CloudAssetOrganizationFeed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudassetorganizationfeedsResource, c.ns, cloudAssetOrganizationFeed), &v1alpha1.CloudAssetOrganizationFeed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudAssetOrganizationFeed), err
}

// Update takes the representation of a cloudAssetOrganizationFeed and updates it. Returns the server's representation of the cloudAssetOrganizationFeed, and an error, if there is any.
func (c *FakeCloudAssetOrganizationFeeds) Update(ctx context.Context, cloudAssetOrganizationFeed *v1alpha1.CloudAssetOrganizationFeed, opts v1.UpdateOptions) (result *v1alpha1.CloudAssetOrganizationFeed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudassetorganizationfeedsResource, c.ns, cloudAssetOrganizationFeed), &v1alpha1.CloudAssetOrganizationFeed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudAssetOrganizationFeed), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudAssetOrganizationFeeds) UpdateStatus(ctx context.Context, cloudAssetOrganizationFeed *v1alpha1.CloudAssetOrganizationFeed, opts v1.UpdateOptions) (*v1alpha1.CloudAssetOrganizationFeed, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudassetorganizationfeedsResource, "status", c.ns, cloudAssetOrganizationFeed), &v1alpha1.CloudAssetOrganizationFeed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudAssetOrganizationFeed), err
}

// Delete takes name of the cloudAssetOrganizationFeed and deletes it. Returns an error if one occurs.
func (c *FakeCloudAssetOrganizationFeeds) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cloudassetorganizationfeedsResource, c.ns, name, opts), &v1alpha1.CloudAssetOrganizationFeed{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudAssetOrganizationFeeds) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudassetorganizationfeedsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudAssetOrganizationFeedList{})
	return err
}

// Patch applies the patch and returns the patched cloudAssetOrganizationFeed.
func (c *FakeCloudAssetOrganizationFeeds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudAssetOrganizationFeed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudassetorganizationfeedsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudAssetOrganizationFeed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudAssetOrganizationFeed), err
}
