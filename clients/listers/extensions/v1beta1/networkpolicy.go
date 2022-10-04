//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

package v1beta1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	extensionsv1beta1listers "k8s.io/client-go/listers/extensions/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// NetworkPolicyClusterLister can list NetworkPolicies across all workspaces, or scope down to a NetworkPolicyLister for one workspace.
type NetworkPolicyClusterLister interface {
	List(selector labels.Selector) (ret []*extensionsv1beta1.NetworkPolicy, err error)
	Cluster(cluster logicalcluster.Name) extensionsv1beta1listers.NetworkPolicyLister
}

type networkPolicyClusterLister struct {
	indexer cache.Indexer
}

// NewNetworkPolicyClusterLister returns a new NetworkPolicyClusterLister.
func NewNetworkPolicyClusterLister(indexer cache.Indexer) *networkPolicyClusterLister {
	return &networkPolicyClusterLister{indexer: indexer}
}

// List lists all NetworkPolicies in the indexer across all workspaces.
func (s *networkPolicyClusterLister) List(selector labels.Selector) (ret []*extensionsv1beta1.NetworkPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*extensionsv1beta1.NetworkPolicy))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get NetworkPolicies.
func (s *networkPolicyClusterLister) Cluster(cluster logicalcluster.Name) extensionsv1beta1listers.NetworkPolicyLister {
	return &networkPolicyLister{indexer: s.indexer, cluster: cluster}
}

// networkPolicyLister implements the extensionsv1beta1listers.NetworkPolicyLister interface.
type networkPolicyLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all NetworkPolicies in the indexer for a workspace.
func (s *networkPolicyLister) List(selector labels.Selector) (ret []*extensionsv1beta1.NetworkPolicy, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*extensionsv1beta1.NetworkPolicy)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

// NetworkPolicies returns an object that can list and get NetworkPolicies in one namespace.
func (s *networkPolicyLister) NetworkPolicies(namespace string) extensionsv1beta1listers.NetworkPolicyNamespaceLister {
	return &networkPolicyNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// networkPolicyNamespaceLister implements the extensionsv1beta1listers.NetworkPolicyNamespaceLister interface.
type networkPolicyNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all NetworkPolicies in the indexer for a given workspace and namespace.
func (s *networkPolicyNamespaceLister) List(selector labels.Selector) (ret []*extensionsv1beta1.NetworkPolicy, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*extensionsv1beta1.NetworkPolicy)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}
	return ret, err
}

// Get retrieves the NetworkPolicy from the indexer for a given workspace, namespace and name.
func (s *networkPolicyNamespaceLister) Get(name string) (*extensionsv1beta1.NetworkPolicy, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(extensionsv1beta1.Resource("NetworkPolicy"), name)
	}
	return obj.(*extensionsv1beta1.NetworkPolicy), nil
}
