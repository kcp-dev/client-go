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

package v1alpha1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	resourcev1alpha1 "k8s.io/api/resource/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	resourcev1alpha1listers "k8s.io/client-go/listers/resource/v1alpha2"
	"k8s.io/client-go/tools/cache"
)

// ResourceClassClusterLister can list ResourceClasses across all workspaces, or scope down to a ResourceClassLister for one workspace.
// All objects returned here must be treated as read-only.
type ResourceClassClusterLister interface {
	// List lists all ResourceClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*resourcev1alpha1.ResourceClass, err error)
	// Cluster returns a lister that can list and get ResourceClasses in one workspace.
	Cluster(clusterName logicalcluster.Name) resourcev1alpha1listers.ResourceClassLister
	ResourceClassClusterListerExpansion
}

type resourceClassClusterLister struct {
	indexer cache.Indexer
}

// NewResourceClassClusterLister returns a new ResourceClassClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewResourceClassClusterLister(indexer cache.Indexer) *resourceClassClusterLister {
	return &resourceClassClusterLister{indexer: indexer}
}

// List lists all ResourceClasses in the indexer across all workspaces.
func (s *resourceClassClusterLister) List(selector labels.Selector) (ret []*resourcev1alpha1.ResourceClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*resourcev1alpha1.ResourceClass))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ResourceClasses.
func (s *resourceClassClusterLister) Cluster(clusterName logicalcluster.Name) resourcev1alpha1listers.ResourceClassLister {
	return &resourceClassLister{indexer: s.indexer, clusterName: clusterName}
}

// resourceClassLister implements the resourcev1alpha1listers.ResourceClassLister interface.
type resourceClassLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all ResourceClasses in the indexer for a workspace.
func (s *resourceClassLister) List(selector labels.Selector) (ret []*resourcev1alpha1.ResourceClass, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*resourcev1alpha1.ResourceClass))
	})
	return ret, err
}

// Get retrieves the ResourceClass from the indexer for a given workspace and name.
func (s *resourceClassLister) Get(name string) (*resourcev1alpha1.ResourceClass, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(resourcev1alpha1.Resource("resourceclasses"), name)
	}
	return obj.(*resourcev1alpha1.ResourceClass), nil
}
