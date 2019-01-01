// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new networks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for networks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AllocateOrReleaseFromIPPool allocate or release from Ip pool API
*/
func (a *Client) AllocateOrReleaseFromIPPool(params *AllocateOrReleaseFromIPPoolParams) (*AllocateOrReleaseFromIPPoolCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocateOrReleaseFromIPPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "allocateOrReleaseFromIpPool",
		Method:             "POST",
		PathPattern:        "/networks/pools/ipPools/{poolId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocateOrReleaseFromIPPoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllocateOrReleaseFromIPPoolCreated), nil

}

/*
CreateIPPool create Ip pool API
*/
func (a *Client) CreateIPPool(params *CreateIPPoolParams) (*CreateIPPoolCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIPPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createIpPool",
		Method:             "POST",
		PathPattern:        "/networks/pools/ipPools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateIPPoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateIPPoolCreated), nil

}

/*
CreateNetwork create network API
*/
func (a *Client) CreateNetwork(params *CreateNetworkParams) (*CreateNetworkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNetwork",
		Method:             "POST",
		PathPattern:        "/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateNetworkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateNetworkCreated), nil

}

/*
GetNetworks get networks API
*/
func (a *Client) GetNetworks(params *GetNetworksParams) (*GetNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworks",
		Method:             "GET",
		PathPattern:        "/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNetworksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNetworksOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
