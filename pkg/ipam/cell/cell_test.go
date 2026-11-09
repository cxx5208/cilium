// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium
// This file was not reviewed by a human before submission

package ipamcell

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ipamOption "github.com/cilium/cilium/pkg/ipam/option"
	k8sResources "github.com/cilium/cilium/pkg/k8s"
	"github.com/cilium/cilium/pkg/option"
)

// The CiliumPodIPPool resource may only be created for multi-pool IPAM.
// Creating it also registers the CiliumPodIPPool CRD with the agent's CRD
// sync barrier, so creating it for other modes would block agent startup
// on a CRD those modes never use.
func TestNewCiliumPodIPPoolResourceSkipsNonMultiPoolModes(t *testing.T) {
	modes := []string{
		ipamOption.IPAMKubernetes,
		ipamOption.IPAMClusterPool,
		ipamOption.IPAMCRD,
		ipamOption.IPAMENI,
		ipamOption.IPAMAzure,
	}
	for _, mode := range modes {
		res, err := newCiliumPodIPPoolResource(
			k8sResources.CiliumResourceParams{},
			&option.DaemonConfig{IPAM: mode},
		)
		require.NoError(t, err)
		assert.Nil(t, res, "IPAM mode %q must not create the CiliumPodIPPool resource", mode)
	}
}
