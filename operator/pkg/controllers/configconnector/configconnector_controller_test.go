// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configconnector

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"
	testmocks "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/mocks"

	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func TestHandleReconcileFailed(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	c := mgr.GetClient()
	mockEventRecorder := testmocks.NewMockEventRecorder(t, mgr.GetScheme())
	r := ConfigConnectorReconciler{
		client:   c,
		recorder: mockEventRecorder,
	}

	apiVersion, kind := corev1beta1.ConfigConnectorGroupVersionKind.ToAPIVersionAndKind()
	nn := types.NamespacedName{
		Name: "kcc-for-reconcile-failed-test",
	}

	tc := testCaseStruct{
		cc: &corev1beta1.ConfigConnector{
			TypeMeta: metav1.TypeMeta{
				Kind:       kind,
				APIVersion: apiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: nn.Name,
			},
			Spec: corev1beta1.ConfigConnectorSpec{
				Mode: "namespaced",
			},
		},
	}

	if err := c.Create(ctx, tc.cc); err != nil {
		t.Fatalf("failed to create ConfigConnector: %v", err)
	}
	reconcileErr := fmt.Errorf("reconciliation error")
	if err := r.handleReconcileFailed(ctx, nn, reconcileErr); err != nil {
		t.Errorf("error handling failed reconciliation: %v", err)
	}

	expectedErrMsg := fmt.Sprintf(k8s.ReconcileErrMsgTmpl, reconcileErr)
	mockEventRecorder.AssertEventRecorded(kind, nn, v1.EventTypeWarning, k8s.UpdateFailed, expectedErrMsg)

	newCC := &corev1beta1.ConfigConnector{}
	if err := c.Get(ctx, nn, newCC); err != nil {
		t.Errorf("failed to get ConfigConnector after attempt to handle failed reconciliation: %v", err)
	}
	status := newCC.GetCommonStatus()
	if status.Healthy {
		t.Errorf("unexpected value for status.healthy: got 'true', want 'false'")
	}
	if len(status.Errors) != 1 {
		t.Errorf("unexpected number of errors in status.errors: got %v errors, want 1 error. Got the errors: %v", len(status.Errors), status.Errors)
	} else if errMsg := status.Errors[0]; errMsg != expectedErrMsg {
		t.Errorf("unexpected error in status.errors: got '%v', want '%v'", errMsg, expectedErrMsg)
	}
}

func TestHandleReconcileSucceeded(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	c := mgr.GetClient()
	mockEventRecorder := testmocks.NewMockEventRecorder(t, mgr.GetScheme())
	r := ConfigConnectorReconciler{
		client:   c,
		recorder: mockEventRecorder,
	}
	apiVersion, kind := corev1beta1.ConfigConnectorGroupVersionKind.ToAPIVersionAndKind()
	nn := types.NamespacedName{
		Name: "kcc-for-reconcile-succeeded-test",
	}
	tc := testCaseStruct{
		cc: &corev1beta1.ConfigConnector{
			TypeMeta: metav1.TypeMeta{
				Kind:       kind,
				APIVersion: apiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: nn.Name,
			},
			Spec: corev1beta1.ConfigConnectorSpec{
				Mode: "namespaced",
			},
		},
	}

	if err := c.Create(ctx, tc.cc); err != nil {
		t.Fatalf("failed to create ConfigConnector: %v", err)
	}
	if err := r.handleReconcileSucceeded(ctx, nn); err != nil {
		t.Errorf("error handling successful reconciliation: %v", err)
	}
	mockEventRecorder.AssertEventRecorded(kind, nn, v1.EventTypeNormal, k8s.UpToDate, k8s.UpToDateMessage)

	newCC := &corev1beta1.ConfigConnector{}
	if err := c.Get(ctx, nn, newCC); err != nil {
		t.Errorf("failed to get ConfigConnector after attempt to handle successful reconciliation: %v", err)
	}
	status := newCC.GetCommonStatus()
	if !status.Healthy {
		t.Errorf("unexpected value for status.healthy: got 'false', want 'true'")
	}
	if len(status.Errors) != 0 {
		t.Errorf("unexpected number of errors in status.errors: got %v errors, want 0 errors. Got the errors: %v", len(status.Errors), status.Errors)
	}
}

func TestHandleConfigConnectorCreate(t *testing.T) {
	t.Parallel()
	tests := []testCaseStruct{
		{
			name: "1 CC and 1 CCContext, namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-1",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      k8s.ConfigConnectorContextAllowedName,
						Namespace: "foo-ns",
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			loadedManifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return []string{testcontroller.FooCRD, testcontroller.SystemNs}
			},
		},
		{
			name: "1 CC and no CCContext, namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-2",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			loadedManifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return []string{testcontroller.FooCRD, testcontroller.SystemNs}
			},
		},
		{
			name: "1 CC in cluster-mode with workload identity and no CCContext",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-1",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			loadedManifest: testcontroller.GetClusterModeWorkloadIdentityManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
		},
		{
			name: "1 CC in cluster-mode with gcp identity and no CCContext",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-1",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			loadedManifest: testcontroller.GetClusterModeGCPManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key")
			},
		},
		{
			name: "1 CC with 1 CCC (ignored), cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-1",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			cccs: []corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      k8s.ConfigConnectorContextAllowedName,
						Namespace: "foo-ns",
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			loadedManifest: testcontroller.GetClusterModeWorkloadIdentityManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			testcontroller.EnsureNamespaceExists(c, k8s.CNRMSystemNamespace)
			testcontroller.EnsureNamespaceExists(c, k8s.OperatorSystemNamespace)
			m := testcontroller.ParseObjects(t, ctx, tc.loadedManifest)
			r := newConfigConnectorReconciler(c)

			if err := c.Create(ctx, tc.cc); err != nil {
				t.Fatalf("error creating %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}

			for _, ccc := range tc.cccs {
				testcontroller.EnsureNamespaceExists(c, ccc.Namespace)
				if err := c.Create(ctx, &ccc); err != nil {
					t.Fatalf("error creating %v %v: %v", ccc.Kind, ccc.Name, err)
				}
			}

			if err := handleLifecycles(t, ctx, r, tc.cc, m); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedObjs := tc.resultsFunc(t, c)
			expectedManifest := testcontroller.ParseObjects(t, ctx, expectedObjs)
			expectedJson, err := expectedManifest.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			resJson, err := m.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(resJson, expectedJson) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(resJson, expectedJson))
			}

			// Verify that CCC objects are NOT attached finalizers by the CC controller.
			for _, ccc := range tc.cccs {
				o := &corev1beta1.ConfigConnectorContext{}
				contextKey := types.NamespacedName{
					Name:      ccc.Name,
					Namespace: ccc.Namespace,
				}
				if err := c.Get(ctx, contextKey, o); err != nil {
					t.Fatalf("error getting ConfigConnector %v: %v", contextKey, err)
				}
				if testcontroller.HasOperatorFinalizer(o) {
					t.Fatalf("%v finalizer was found in %v", k8s.OperatorFinalizer, ccc)
				}
			}
			// Verify that CC contains the operator finalizer.
			ccKey := types.NamespacedName{
				Name:      tc.cc.Name,
				Namespace: tc.cc.Namespace,
			}
			if err := c.Get(ctx, ccKey, tc.cc); err != nil {
				t.Fatalf("error getting ConfigConnector %v: %v", ccKey, err)
			}
			if !testcontroller.HasOperatorFinalizer(tc.cc) {
				t.Fatalf("no %v finalizer was found in %v", k8s.OperatorFinalizer, tc.cc)
			}
		})
	}
}

func TestHandleConfigConnectorDelete(t *testing.T) {
	tests := []struct {
		name                 string
		cc                   *corev1beta1.ConfigConnector
		cccs                 []corev1beta1.ConfigConnectorContext
		installedObjectsFunc func(t *testing.T, c client.Client) []string
		resultsFunc          func(t *testing.T, c client.Client) []string
	}{
		{
			name: "cluster mode workload identity uninstall",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:       "test-kcc",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
		},
		{
			name: "cluster mode gcp identity uninstall",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:       "test-kcc",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key")
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
		},
		{
			name: "namespaced mode CC, 1 CCContext, delete CC",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:       "test-kcc",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       k8s.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				res := make([]string, 0)
				res = append(res, testcontroller.GetSharedComponentsManifest()...)
				namespacedManifest := testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.NamespacedComponentsTemplate, "foo-ns", "foo-ns@bar.iam.gserviceaccount.com", false, "", c)
				res = append(res, namespacedManifest...)
				res = append(res, testcontroller.PerNamespaceControllerManagerPod)
				return res
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			testcontroller.EnsureNamespaceExists(c, k8s.OperatorSystemNamespace)
			testcontroller.EnsureNamespaceExists(c, k8s.CNRMSystemNamespace)
			m := testcontroller.ParseObjects(t, ctx, tc.installedObjectsFunc(t, c))
			r := newConfigConnectorReconciler(c)
			if err := c.Create(ctx, tc.cc); err != nil {
				t.Fatalf("error creating %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			key := types.NamespacedName{
				Name: tc.cc.Name,
			}
			if err := c.Get(ctx, key, tc.cc); err != nil {
				t.Fatalf("error getting %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			for _, ccc := range tc.cccs {
				testcontroller.EnsureNamespaceExists(c, ccc.Namespace)
				if err := c.Create(ctx, &ccc); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", ccc.Kind, ccc.Namespace, ccc.Name, err)
				}
			}

			for _, item := range m.Items {
				if err := c.Create(ctx, item.UnstructuredObject()); err != nil && !apierrors.IsAlreadyExists(err) {
					t.Fatalf("error creating %v %v: %v", item.GroupKind(), item.GetName(), err)
				}
			}

			// issue the delete request for the configconnector object
			if err := c.Delete(ctx, tc.cc); err != nil {
				t.Fatalf("error deleting %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			if err := c.Get(ctx, key, tc.cc); err != nil {
				t.Fatalf("error getting %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			if len(tc.cccs) > 0 {
				// Expect that the first attempt returns an error.
				if err := handleLifecycles(t, ctx, r, tc.cc, m); err == nil {
					t.Fatalf("expect to have an error because the controller manager pod per namespace is not deleted, but got nil")
				}
				// Simulate that CCC controller kicks in and deletes the controller manager pod.
				for _, item := range m.Items {
					if item.Kind == "Pod" && strings.Contains(item.GetName(), "cnrm-controller-manager") {
						if err := c.Delete(ctx, item.UnstructuredObject()); err != nil {
							t.Fatalf("error deleting %v %v: %v", item.GroupKind(), item.GetName(), err)
						}
					}
				}
			}
			if err := handleLifecycles(t, ctx, r, tc.cc, m); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedObjs := tc.resultsFunc(t, c)
			expectedManifest := testcontroller.ParseObjects(t, ctx, expectedObjs)
			expectedJson, err := expectedManifest.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			resJson, err := m.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(resJson, expectedJson) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(resJson, expectedJson))
			}

			// Assert that the ConfigConnector object is deleted.
			if err := c.Get(ctx, key, tc.cc); err == nil || !apierrors.IsNotFound(err) {
				t.Fatalf("expect to get %v error, but got error: %v", metav1.StatusReasonNotFound, err)
			}
		})
	}
}

func TestConfigConnectorUpdate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                 string
		cc                   *corev1beta1.ConfigConnector
		updatedCc            *corev1beta1.ConfigConnector
		cccs                 []*corev1beta1.ConfigConnectorContext
		installedObjectsFunc func(t *testing.T, c client.Client) []string
		manifest             []string
		toDeleteObjectsFunc  func(t *testing.T, c client.Client) []string
		resultsFunc          func(t *testing.T, c client.Client) []string
	}{
		{
			name: "workload identity cluster mode to namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       k8s.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.ClusterModeOnlyWorkloadIdentityComponents, "foo@bar.iam.gserviceaccount.com")
			},
			manifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				return res
			},
		},
		{
			name: "gcp identity cluster mode to namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       k8s.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key")
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.ClusterModeOnlyGCPComponents, "my-key")
			},
			manifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				return res
			},
		},
		{
			name: "namespaced mode to workload identity cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       k8s.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				res = append(res, testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.NamespacedComponentsTemplate, "foo-ns", "foo-ns@bar.iam.gserviceaccount.com", false, "", c)...)
				res = append(res, testcontroller.PerNamespaceControllerManagerPod)
				return res
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return []string{}
			},
			manifest: testcontroller.GetClusterModeWorkloadIdentityManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
		},
		{
			name: "namespaced mode to gcp identity cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       k8s.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				res = append(res, testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.NamespacedComponentsTemplate, "foo-ns", "foo-ns@bar.iam.gserviceaccount.com", false, "", c)...)
				res = append(res, testcontroller.PerNamespaceControllerManagerPod)
				return res
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return []string{}
			},
			manifest: testcontroller.GetClusterModeGCPManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key ")
			},
		},
		{
			name: "workload identity cluster mode to gcp identity cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			manifest: testcontroller.GetClusterModeGCPManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key ")
			},
		},
		{
			name: "gcp identity cluster mode to workload identity cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key ")
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			manifest: testcontroller.GetClusterModeWorkloadIdentityManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			r := newConfigConnectorReconciler(c)

			testcontroller.EnsureNamespaceExists(c, k8s.OperatorSystemNamespace)
			if err := c.Create(ctx, tc.cc); err != nil {
				t.Fatalf("error creating %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			for _, ccc := range tc.cccs {
				testcontroller.EnsureNamespaceExists(c, ccc.Namespace)
				if err := c.Create(ctx, ccc); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", ccc.Kind, ccc.Namespace, ccc.Name, err)
				}
			}
			installedComponents := tc.installedObjectsFunc(t, c)
			installedManifest := testcontroller.ParseObjects(t, ctx, installedComponents)
			for _, item := range installedManifest.Items {
				if err := c.Create(ctx, item.UnstructuredObject()); err != nil && !apierrors.IsAlreadyExists(err) {
					t.Fatalf("error creating %v %v: %v", item.GroupKind(), item.GetName(), err)
				}
			}

			// update ConfigConnector
			tc.updatedCc.ResourceVersion = tc.cc.ResourceVersion
			if err := c.Update(ctx, tc.updatedCc); err != nil {
				t.Fatalf("error updating %v %v: %v", tc.updatedCc.Kind, tc.updatedCc.Name, err)
			}

			m := testcontroller.ParseObjects(t, ctx, tc.manifest)
			if tc.cc.GetMode() == "namespaced" {
				if err := handleLifecycles(t, ctx, r, tc.updatedCc, m); err == nil {
					t.Fatalf("got nil, but want to have an error because the controller manager pod per namespace is not deleted")
				}
				// Simulate that CCC controller kicks in and deletes the controller manager pod.
				for _, item := range installedManifest.Items {
					if item.Kind == "Pod" && strings.Contains(item.GetName(), "cnrm-controller-manager") {
						if err := c.Delete(ctx, item.UnstructuredObject()); err != nil {
							t.Fatalf("error deleting %v %v: %v", item.GroupKind(), item.GetName(), err)
						}
					}
				}
			}
			if err := handleLifecycles(t, ctx, r, tc.updatedCc, m); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedObjs := tc.resultsFunc(t, c)
			expectedManifest := testcontroller.ParseObjects(t, ctx, expectedObjs)
			expectedJson, err := expectedManifest.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			resJson, err := m.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(resJson, expectedJson) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(resJson, expectedJson))
			}
			// assert unneeded components are deleted
			unneededComponents := tc.toDeleteObjectsFunc(t, c)
			for _, str := range unneededComponents {
				obj := testcontroller.ToUnstructured(t, str)
				key := types.NamespacedName{Name: obj.GetName(), Namespace: obj.GetNamespace()}
				if err := c.Get(ctx, key, obj); err == nil || !apierrors.IsNotFound(err) {
					t.Fatalf("expect to get %v error for %v %v, but got error: %v", metav1.StatusReasonNotFound, obj.GetKind(), key, err)
				}
			}
		})
	}
}

type testCaseStruct struct {
	name           string
	cc             *corev1beta1.ConfigConnector
	cccs           []corev1beta1.ConfigConnectorContext
	loadedManifest []string
	resultsFunc    func(t *testing.T, c client.Client) []string
}

func handleLifecycles(t *testing.T, ctx context.Context,
	r *ConfigConnectorReconciler, cc *corev1beta1.ConfigConnector, m *manifest.Objects) error {
	t.Helper()

	fn := r.transformForClusterMode()
	if err := fn(ctx, cc, m); err != nil {
		return err
	}
	fn = r.handleConfigConnectorLifecycle()
	if err := fn(ctx, cc, m); err != nil {
		return err
	}
	return nil
}

func newConfigConnectorReconciler(c client.Client) *ConfigConnectorReconciler {
	return &ConfigConnectorReconciler{
		client: c,
		log:    logr.Discard(),
	}
}

func TestSelectingCRDsByVersion(t *testing.T) {
	tests := []struct {
		name              string
		manifests         []string
		version           string
		expectedManifests []string
		hasError          bool
	}{
		{
			name:              "select v1alpha1 CRD from v1alpha1 and v1beta1 CRDs",
			manifests:         testcontroller.GetManifestsWithAlphaAndBetaCRDs(),
			version:           "v1alpha1",
			expectedManifests: testcontroller.GetManifestsWithAlphaCRD(),
		},
		{
			name:              "select v1alpha1 CRD from v1beta1 CRDs",
			manifests:         testcontroller.GetManifestsWithBetaCRD(),
			version:           "v1alpha1",
			expectedManifests: testcontroller.GetManifestsWithNoCRD(),
		},
		{
			name:      "select v1alpha1 CRD from non-KCC CRD",
			manifests: testcontroller.GetManifestsWithNonKCCCRD(),
			version:   "v1alpha1",
			hasError:  true,
		},
		{
			name:      "select v1alpha1 CRD from defective CRD",
			manifests: testcontroller.GetManifestsWithDefectiveCRD(),
			version:   "v1alpha1",
			hasError:  true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			manifests := testcontroller.ParseObjects(t, ctx, tc.manifests)
			r := newConfigConnectorReconciler(c)

			err := r.selectCRDsByVersion(manifests, tc.version)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, want an error")
				}
				return
			} else if err != nil {
				t.Fatalf("error selecting CRDs by version: %v", err)
			}

			processedJSON, err := manifests.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedManifests := testcontroller.ParseObjects(t, ctx, tc.expectedManifests)
			expectedJSON, err := expectedManifests.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(processedJSON, expectedJSON) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(processedJSON, expectedJSON))
			}
		})
	}
}
