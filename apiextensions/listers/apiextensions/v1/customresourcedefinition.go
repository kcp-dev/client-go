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

package v1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsv1listers "k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CustomResourceDefinitionClusterLister can list CustomResourceDefinitions across all workspaces, or scope down to a CustomResourceDefinitionLister for one workspace.
// All objects returned here must be treated as read-only.
type CustomResourceDefinitionClusterLister interface {
	// List lists all CustomResourceDefinitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apiextensionsv1.CustomResourceDefinition, err error)
	// Cluster returns a lister that can list and get CustomResourceDefinitions in one workspace.
	Cluster(clusterName logicalcluster.Name) apiextensionsv1listers.CustomResourceDefinitionLister
	CustomResourceDefinitionClusterListerExpansion
}

type customResourceDefinitionClusterLister struct {
	indexer cache.Indexer
}

// NewCustomResourceDefinitionClusterLister returns a new CustomResourceDefinitionClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewCustomResourceDefinitionClusterLister(indexer cache.Indexer) *customResourceDefinitionClusterLister {
	return &customResourceDefinitionClusterLister{indexer: indexer}
}

// List lists all CustomResourceDefinitions in the indexer across all workspaces.
func (s *customResourceDefinitionClusterLister) List(selector labels.Selector) (ret []*apiextensionsv1.CustomResourceDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*apiextensionsv1.CustomResourceDefinition))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get CustomResourceDefinitions.
func (s *customResourceDefinitionClusterLister) Cluster(clusterName logicalcluster.Name) apiextensionsv1listers.CustomResourceDefinitionLister {
	return &customResourceDefinitionLister{indexer: s.indexer, clusterName: clusterName}
}

// customResourceDefinitionLister implements the apiextensionsv1listers.CustomResourceDefinitionLister interface.
type customResourceDefinitionLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all CustomResourceDefinitions in the indexer for a workspace.
func (s *customResourceDefinitionLister) List(selector labels.Selector) (ret []*apiextensionsv1.CustomResourceDefinition, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*apiextensionsv1.CustomResourceDefinition))
	})
	return ret, err
}

// Get retrieves the CustomResourceDefinition from the indexer for a given workspace and name.
func (s *customResourceDefinitionLister) Get(name string) (*apiextensionsv1.CustomResourceDefinition, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(apiextensionsv1.Resource("customresourcedefinitions"), name)
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), nil
}
