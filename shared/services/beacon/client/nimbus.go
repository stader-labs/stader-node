package client

import "github.com/stader-labs/stader-node/shared/services/beacon"

type NimbusClient struct {
	StandardHttpClient
}

// Create a new client instance
func NewNimbusClient(providerAddress string) *NimbusClient {
	return &NimbusClient{
		StandardHttpClient: *NewStandardHttpClient(providerAddress),
	}
}

func (n *NimbusClient) GetClientType() (beacon.BeaconClientType, error) {
	return beacon.SingleProcess, nil
}
