/*
Copyright 2020 The Kubernetes Authors.

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

package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2022-05-01/network"
	"github.com/Azure/go-autorest/autorest"
)

// NetworkInterfacesClient is a client for managing Network Interfaces.
type NetworkInterfacesClient interface {
	ListScaleSetsNetworkInterfaces(ctx context.Context, resourceGroupName, vmssName string) ([]network.Interface, error)
}

type networkInterfacesClientImpl struct {
	c *network.InterfacesClient
}

var _ NetworkInterfacesClient = &networkInterfacesClientImpl{}

func (c *networkInterfacesClientImpl) ListScaleSetsNetworkInterfaces(ctx context.Context, resourceGroupName, vmssName string) ([]network.Interface, error) {
	var l []network.Interface
	for iter, err := c.c.ListVirtualMachineScaleSetNetworkInterfacesComplete(ctx, resourceGroupName, vmssName); iter.NotDone(); err = iter.Next() {
		if err != nil {
			return nil, err
		}
		l = append(l, iter.Value())
	}
	return l, nil
}

func newNetworkInterfacesClientImpl(subscriptionID string, authorizer autorest.Authorizer) *networkInterfacesClientImpl {
	c := network.NewInterfacesClient(subscriptionID)
	c.Authorizer = authorizer
	return &networkInterfacesClientImpl{
		c: &c,
	}
}
