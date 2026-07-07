// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package egressgateway

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cilium/cilium/pkg/option"
)

func TestSupportedIdentityAllocationMode(t *testing.T) {
	tests := []struct {
		name string
		mode string
		want bool
	}{
		{name: "crd", mode: option.IdentityAllocationModeCRD, want: true},
		{name: "kvstore", mode: option.IdentityAllocationModeKVstore, want: true},
		{name: "doublewrite read kvstore", mode: option.IdentityAllocationModeDoubleWriteReadKVstore, want: true},
		{name: "doublewrite read crd", mode: option.IdentityAllocationModeDoubleWriteReadCRD, want: true},
		{name: "unknown", mode: "unknown", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, isSupportedIdentityAllocationMode(tt.mode))
		})
	}
}
