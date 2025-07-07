/*
Copyright 2014 The Kubernetes Authors.
Modifications Copyright 2022 The KCP Authors.

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

package fake

import (
	"context"

	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	core "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var deploymentsResource = schema.GroupVersionResource{Group: "extensions", Version: "v1beta1", Resource: "deployments"}
var deploymentsKind = schema.GroupVersionKind{Group: "extensions", Version: "v1beta1", Kind: "Deployment"}

func (c *deploymentScopedClient) Rollback(ctx context.Context, deploymentRollback *v1beta1.DeploymentRollback, opts metav1.CreateOptions) error {
	action := core.CreateActionImpl{}
	action.Verb = "create"
	action.Resource = deploymentsResource
	action.Subresource = "rollback"
	action.Object = deploymentRollback
	action.ClusterPath = c.ClusterPath

	_, err := c.Fake.Invokes(action, deploymentRollback)
	return err
}
