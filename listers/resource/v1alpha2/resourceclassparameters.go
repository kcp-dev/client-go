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

package v1alpha2

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	resourcev1alpha2 "k8s.io/api/resource/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	resourcev1alpha2listers "k8s.io/client-go/listers/resource/v1alpha2"
	"k8s.io/client-go/tools/cache"
)

// ResourceClassParametersClusterLister can list ResourceClassParameters across all workspaces, or scope down to a ResourceClassParametersLister for one workspace.
// All objects returned here must be treated as read-only.
type ResourceClassParametersClusterLister interface {
	// List lists all ResourceClassParameters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*resourcev1alpha2.ResourceClassParameters, err error)
	// Cluster returns a lister that can list and get ResourceClassParameters in one workspace.
	Cluster(clusterName logicalcluster.Name) resourcev1alpha2listers.ResourceClassParametersLister
	ResourceClassParametersClusterListerExpansion
}

type resourceClassParametersClusterLister struct {
	indexer cache.Indexer
}

// NewResourceClassParametersClusterLister returns a new ResourceClassParametersClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewResourceClassParametersClusterLister(indexer cache.Indexer) *resourceClassParametersClusterLister {
	return &resourceClassParametersClusterLister{indexer: indexer}
}

// List lists all ResourceClassParameters in the indexer across all workspaces.
func (s *resourceClassParametersClusterLister) List(selector labels.Selector) (ret []*resourcev1alpha2.ResourceClassParameters, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*resourcev1alpha2.ResourceClassParameters))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ResourceClassParameters.
func (s *resourceClassParametersClusterLister) Cluster(clusterName logicalcluster.Name) resourcev1alpha2listers.ResourceClassParametersLister {
	return &resourceClassParametersLister{indexer: s.indexer, clusterName: clusterName}
}

// resourceClassParametersLister implements the resourcev1alpha2listers.ResourceClassParametersLister interface.
type resourceClassParametersLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all ResourceClassParameters in the indexer for a workspace.
func (s *resourceClassParametersLister) List(selector labels.Selector) (ret []*resourcev1alpha2.ResourceClassParameters, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*resourcev1alpha2.ResourceClassParameters))
	})
	return ret, err
}

// ResourceClassParameters returns an object that can list and get ResourceClassParameters in one namespace.
func (s *resourceClassParametersLister) ResourceClassParameters(namespace string) resourcev1alpha2listers.ResourceClassParametersNamespaceLister {
	return &resourceClassParametersNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// resourceClassParametersNamespaceLister implements the resourcev1alpha2listers.ResourceClassParametersNamespaceLister interface.
type resourceClassParametersNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all ResourceClassParameters in the indexer for a given workspace and namespace.
func (s *resourceClassParametersNamespaceLister) List(selector labels.Selector) (ret []*resourcev1alpha2.ResourceClassParameters, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*resourcev1alpha2.ResourceClassParameters))
	})
	return ret, err
}

// Get retrieves the ResourceClassParameters from the indexer for a given workspace, namespace and name.
func (s *resourceClassParametersNamespaceLister) Get(name string) (*resourcev1alpha2.ResourceClassParameters, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(resourcev1alpha2.Resource("resourceclassparameters"), name)
	}
	return obj.(*resourcev1alpha2.ResourceClassParameters), nil
}
