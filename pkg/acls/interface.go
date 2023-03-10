/*
Copyright 2017 The Kubernetes Authors.

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

package acls

import (
	"context"

	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/util/pkg/vfs"
)

// ACLStrategy is the interface implemented by ACL strategy providers
type ACLStrategy interface {
	// GetACL returns the ACL if this strategy handles the vfs.Path, when writing for the specified cluster
	GetACL(ctx context.Context, p vfs.Path, cluster *kops.Cluster) (vfs.ACL, error)
}
