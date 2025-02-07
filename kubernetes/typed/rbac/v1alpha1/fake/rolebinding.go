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

	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsrbacv1alpha1 "k8s.io/client-go/applyconfigurations/rbac/v1alpha1"
	rbacv1alpha1client "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	"k8s.io/client-go/testing"

	kcprbacv1alpha1 "github.com/kcp-dev/client-go/kubernetes/typed/rbac/v1alpha1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var roleBindingsResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "v1alpha1", Resource: "rolebindings"}
var roleBindingsKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1alpha1", Kind: "RoleBinding"}

type roleBindingsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *roleBindingsClusterClient) Cluster(clusterPath logicalcluster.Path) kcprbacv1alpha1.RoleBindingsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &roleBindingsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of RoleBindings that match those selectors across all clusters.
func (c *roleBindingsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1alpha1.RoleBindingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(roleBindingsResource, roleBindingsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &rbacv1alpha1.RoleBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1alpha1.RoleBindingList{ListMeta: obj.(*rbacv1alpha1.RoleBindingList).ListMeta}
	for _, item := range obj.(*rbacv1alpha1.RoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested RoleBindings across all clusters.
func (c *roleBindingsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(roleBindingsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type roleBindingsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *roleBindingsNamespacer) Namespace(namespace string) rbacv1alpha1client.RoleBindingInterface {
	return &roleBindingsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type roleBindingsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *roleBindingsClient) Create(ctx context.Context, roleBinding *rbacv1alpha1.RoleBinding, opts metav1.CreateOptions) (*rbacv1alpha1.RoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(roleBindingsResource, c.ClusterPath, c.Namespace, roleBinding), &rbacv1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.RoleBinding), err
}

func (c *roleBindingsClient) Update(ctx context.Context, roleBinding *rbacv1alpha1.RoleBinding, opts metav1.UpdateOptions) (*rbacv1alpha1.RoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(roleBindingsResource, c.ClusterPath, c.Namespace, roleBinding), &rbacv1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.RoleBinding), err
}

func (c *roleBindingsClient) UpdateStatus(ctx context.Context, roleBinding *rbacv1alpha1.RoleBinding, opts metav1.UpdateOptions) (*rbacv1alpha1.RoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(roleBindingsResource, c.ClusterPath, "status", c.Namespace, roleBinding), &rbacv1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.RoleBinding), err
}

func (c *roleBindingsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(roleBindingsResource, c.ClusterPath, c.Namespace, name, opts), &rbacv1alpha1.RoleBinding{})
	return err
}

func (c *roleBindingsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(roleBindingsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &rbacv1alpha1.RoleBindingList{})
	return err
}

func (c *roleBindingsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*rbacv1alpha1.RoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(roleBindingsResource, c.ClusterPath, c.Namespace, name), &rbacv1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.RoleBinding), err
}

// List takes label and field selectors, and returns the list of RoleBindings that match those selectors.
func (c *roleBindingsClient) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1alpha1.RoleBindingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(roleBindingsResource, roleBindingsKind, c.ClusterPath, c.Namespace, opts), &rbacv1alpha1.RoleBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1alpha1.RoleBindingList{ListMeta: obj.(*rbacv1alpha1.RoleBindingList).ListMeta}
	for _, item := range obj.(*rbacv1alpha1.RoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *roleBindingsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(roleBindingsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *roleBindingsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*rbacv1alpha1.RoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(roleBindingsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &rbacv1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.RoleBinding), err
}

func (c *roleBindingsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsrbacv1alpha1.RoleBindingApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1alpha1.RoleBinding, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(roleBindingsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &rbacv1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.RoleBinding), err
}

func (c *roleBindingsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsrbacv1alpha1.RoleBindingApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1alpha1.RoleBinding, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(roleBindingsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &rbacv1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.RoleBinding), err
}
