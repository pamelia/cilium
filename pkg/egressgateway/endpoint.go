// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package egressgateway

import (
	"fmt"
	"net/netip"

	endpointTables "github.com/cilium/cilium/pkg/endpoint/tables"
	"github.com/cilium/cilium/pkg/labels"
)

// endpointMetadata stores relevant metadata associated with an endpoint that's updated during endpoint
// add/update events
type endpointMetadata struct {
	// Endpoint labels
	labels map[string]string
	// Endpoint ID
	id endpointID
	// ips are endpoint's unique IPs
	ips []netip.Addr
	// nodeIP is the IP of the node the endpoint is on
	nodeIP string
}

type endpointID string

func getEndpointMetadata(endpoint *endpointTables.Endpoint, identityLabels labels.Labels) (*endpointMetadata, error) {
	if endpoint.Key.IsZero() {
		return nil, fmt.Errorf("endpoint has empty key")
	}

	if len(endpoint.IPs) == 0 {
		return nil, fmt.Errorf("failed to get valid endpoint IPs")
	}

	nodeIP := ""
	if endpoint.HostIP.IsValid() {
		nodeIP = endpoint.HostIP.String()
	}

	data := &endpointMetadata{
		ips:    append([]netip.Addr{}, endpoint.IPs...),
		labels: identityLabels.K8sStringMap(),
		id:     endpointID(endpoint.Key.String()),
		nodeIP: nodeIP,
	}

	return data, nil
}
