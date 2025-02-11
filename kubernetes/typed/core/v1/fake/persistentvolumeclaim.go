/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/testing"

	kcpcorev1 "github.com/kcp-dev/client-go/kubernetes/typed/core/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var persistentVolumeClaimsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "persistentvolumeclaims"}
var persistentVolumeClaimsKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "PersistentVolumeClaim"}

type persistentVolumeClaimsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *persistentVolumeClaimsClusterClient) Cluster(clusterPath logicalcluster.Path) kcpcorev1.PersistentVolumeClaimsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &persistentVolumeClaimsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of PersistentVolumeClaims that match those selectors across all clusters.
func (c *persistentVolumeClaimsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.PersistentVolumeClaimList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(persistentVolumeClaimsResource, persistentVolumeClaimsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &corev1.PersistentVolumeClaimList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.PersistentVolumeClaimList{ListMeta: obj.(*corev1.PersistentVolumeClaimList).ListMeta}
	for _, item := range obj.(*corev1.PersistentVolumeClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested PersistentVolumeClaims across all clusters.
func (c *persistentVolumeClaimsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(persistentVolumeClaimsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type persistentVolumeClaimsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *persistentVolumeClaimsNamespacer) Namespace(namespace string) corev1client.PersistentVolumeClaimInterface {
	return &persistentVolumeClaimsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type persistentVolumeClaimsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *persistentVolumeClaimsClient) Create(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaim, opts metav1.CreateOptions) (*corev1.PersistentVolumeClaim, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(persistentVolumeClaimsResource, c.ClusterPath, c.Namespace, persistentVolumeClaim), &corev1.PersistentVolumeClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolumeClaim), err
}

func (c *persistentVolumeClaimsClient) Update(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaim, opts metav1.UpdateOptions) (*corev1.PersistentVolumeClaim, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(persistentVolumeClaimsResource, c.ClusterPath, c.Namespace, persistentVolumeClaim), &corev1.PersistentVolumeClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolumeClaim), err
}

func (c *persistentVolumeClaimsClient) UpdateStatus(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaim, opts metav1.UpdateOptions) (*corev1.PersistentVolumeClaim, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(persistentVolumeClaimsResource, c.ClusterPath, "status", c.Namespace, persistentVolumeClaim), &corev1.PersistentVolumeClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolumeClaim), err
}

func (c *persistentVolumeClaimsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(persistentVolumeClaimsResource, c.ClusterPath, c.Namespace, name, opts), &corev1.PersistentVolumeClaim{})
	return err
}

func (c *persistentVolumeClaimsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(persistentVolumeClaimsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.PersistentVolumeClaimList{})
	return err
}

func (c *persistentVolumeClaimsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*corev1.PersistentVolumeClaim, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(persistentVolumeClaimsResource, c.ClusterPath, c.Namespace, name), &corev1.PersistentVolumeClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolumeClaim), err
}

// List takes label and field selectors, and returns the list of PersistentVolumeClaims that match those selectors.
func (c *persistentVolumeClaimsClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.PersistentVolumeClaimList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(persistentVolumeClaimsResource, persistentVolumeClaimsKind, c.ClusterPath, c.Namespace, opts), &corev1.PersistentVolumeClaimList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.PersistentVolumeClaimList{ListMeta: obj.(*corev1.PersistentVolumeClaimList).ListMeta}
	for _, item := range obj.(*corev1.PersistentVolumeClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *persistentVolumeClaimsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(persistentVolumeClaimsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *persistentVolumeClaimsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.PersistentVolumeClaim, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(persistentVolumeClaimsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &corev1.PersistentVolumeClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolumeClaim), err
}

func (c *persistentVolumeClaimsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscorev1.PersistentVolumeClaimApplyConfiguration, opts metav1.ApplyOptions) (*corev1.PersistentVolumeClaim, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(persistentVolumeClaimsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &corev1.PersistentVolumeClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolumeClaim), err
}

func (c *persistentVolumeClaimsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscorev1.PersistentVolumeClaimApplyConfiguration, opts metav1.ApplyOptions) (*corev1.PersistentVolumeClaim, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(persistentVolumeClaimsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &corev1.PersistentVolumeClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolumeClaim), err
}
