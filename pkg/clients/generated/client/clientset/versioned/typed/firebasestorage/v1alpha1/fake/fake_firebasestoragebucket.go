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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/firebasestorage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFirebaseStorageBuckets implements FirebaseStorageBucketInterface
type FakeFirebaseStorageBuckets struct {
	Fake *FakeFirebasestorageV1alpha1
	ns   string
}

var firebasestoragebucketsResource = schema.GroupVersionResource{Group: "firebasestorage.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "firebasestoragebuckets"}

var firebasestoragebucketsKind = schema.GroupVersionKind{Group: "firebasestorage.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "FirebaseStorageBucket"}

// Get takes name of the firebaseStorageBucket, and returns the corresponding firebaseStorageBucket object, and an error if there is any.
func (c *FakeFirebaseStorageBuckets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FirebaseStorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(firebasestoragebucketsResource, c.ns, name), &v1alpha1.FirebaseStorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirebaseStorageBucket), err
}

// List takes label and field selectors, and returns the list of FirebaseStorageBuckets that match those selectors.
func (c *FakeFirebaseStorageBuckets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FirebaseStorageBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(firebasestoragebucketsResource, firebasestoragebucketsKind, c.ns, opts), &v1alpha1.FirebaseStorageBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FirebaseStorageBucketList{ListMeta: obj.(*v1alpha1.FirebaseStorageBucketList).ListMeta}
	for _, item := range obj.(*v1alpha1.FirebaseStorageBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested firebaseStorageBuckets.
func (c *FakeFirebaseStorageBuckets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(firebasestoragebucketsResource, c.ns, opts))

}

// Create takes the representation of a firebaseStorageBucket and creates it.  Returns the server's representation of the firebaseStorageBucket, and an error, if there is any.
func (c *FakeFirebaseStorageBuckets) Create(ctx context.Context, firebaseStorageBucket *v1alpha1.FirebaseStorageBucket, opts v1.CreateOptions) (result *v1alpha1.FirebaseStorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(firebasestoragebucketsResource, c.ns, firebaseStorageBucket), &v1alpha1.FirebaseStorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirebaseStorageBucket), err
}

// Update takes the representation of a firebaseStorageBucket and updates it. Returns the server's representation of the firebaseStorageBucket, and an error, if there is any.
func (c *FakeFirebaseStorageBuckets) Update(ctx context.Context, firebaseStorageBucket *v1alpha1.FirebaseStorageBucket, opts v1.UpdateOptions) (result *v1alpha1.FirebaseStorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(firebasestoragebucketsResource, c.ns, firebaseStorageBucket), &v1alpha1.FirebaseStorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirebaseStorageBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFirebaseStorageBuckets) UpdateStatus(ctx context.Context, firebaseStorageBucket *v1alpha1.FirebaseStorageBucket, opts v1.UpdateOptions) (*v1alpha1.FirebaseStorageBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(firebasestoragebucketsResource, "status", c.ns, firebaseStorageBucket), &v1alpha1.FirebaseStorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirebaseStorageBucket), err
}

// Delete takes name of the firebaseStorageBucket and deletes it. Returns an error if one occurs.
func (c *FakeFirebaseStorageBuckets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(firebasestoragebucketsResource, c.ns, name, opts), &v1alpha1.FirebaseStorageBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFirebaseStorageBuckets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(firebasestoragebucketsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FirebaseStorageBucketList{})
	return err
}

// Patch applies the patch and returns the patched firebaseStorageBucket.
func (c *FakeFirebaseStorageBuckets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FirebaseStorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(firebasestoragebucketsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FirebaseStorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirebaseStorageBucket), err
}
