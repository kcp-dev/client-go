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

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var namespacesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "namespaces"}
var namespacesKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}

type namespacesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *namespacesClusterClient) Cluster(clusterPath logicalcluster.Path) corev1client.NamespaceInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &namespacesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Namespaces that match those selectors across all clusters.
func (c *namespacesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.NamespaceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(namespacesResource, namespacesKind, logicalcluster.Wildcard, opts), &corev1.NamespaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.NamespaceList{ListMeta: obj.(*corev1.NamespaceList).ListMeta}
	for _, item := range obj.(*corev1.NamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Namespaces across all clusters.
func (c *namespacesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(namespacesResource, logicalcluster.Wildcard, opts))
}

type namespacesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *namespacesClient) Create(ctx context.Context, namespace *corev1.Namespace, opts metav1.CreateOptions) (*corev1.Namespace, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(namespacesResource, c.ClusterPath, namespace), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

func (c *namespacesClient) Update(ctx context.Context, namespace *corev1.Namespace, opts metav1.UpdateOptions) (*corev1.Namespace, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(namespacesResource, c.ClusterPath, namespace), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

func (c *namespacesClient) UpdateStatus(ctx context.Context, namespace *corev1.Namespace, opts metav1.UpdateOptions) (*corev1.Namespace, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(namespacesResource, c.ClusterPath, "status", namespace), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

func (c *namespacesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(namespacesResource, c.ClusterPath, name, opts), &corev1.Namespace{})
	return err
}

func (c *namespacesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*corev1.Namespace, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(namespacesResource, c.ClusterPath, name), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

// List takes label and field selectors, and returns the list of Namespaces that match those selectors.
func (c *namespacesClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.NamespaceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(namespacesResource, namespacesKind, c.ClusterPath, opts), &corev1.NamespaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.NamespaceList{ListMeta: obj.(*corev1.NamespaceList).ListMeta}
	for _, item := range obj.(*corev1.NamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *namespacesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(namespacesResource, c.ClusterPath, opts))
}

func (c *namespacesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.Namespace, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(namespacesResource, c.ClusterPath, name, pt, data, subresources...), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

func (c *namespacesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscorev1.NamespaceApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Namespace, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(namespacesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

func (c *namespacesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscorev1.NamespaceApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Namespace, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(namespacesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}
