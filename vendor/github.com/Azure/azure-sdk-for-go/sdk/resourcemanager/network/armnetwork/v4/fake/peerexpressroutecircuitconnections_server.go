//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
	"net/http"
	"net/url"
	"regexp"
)

// PeerExpressRouteCircuitConnectionsServer is a fake server for instances of the armnetwork.PeerExpressRouteCircuitConnectionsClient type.
type PeerExpressRouteCircuitConnectionsServer struct {
	// Get is the fake for method PeerExpressRouteCircuitConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *armnetwork.PeerExpressRouteCircuitConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.PeerExpressRouteCircuitConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PeerExpressRouteCircuitConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, circuitName string, peeringName string, options *armnetwork.PeerExpressRouteCircuitConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.PeerExpressRouteCircuitConnectionsClientListResponse])
}

// NewPeerExpressRouteCircuitConnectionsServerTransport creates a new instance of PeerExpressRouteCircuitConnectionsServerTransport with the provided implementation.
// The returned PeerExpressRouteCircuitConnectionsServerTransport instance is connected to an instance of armnetwork.PeerExpressRouteCircuitConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPeerExpressRouteCircuitConnectionsServerTransport(srv *PeerExpressRouteCircuitConnectionsServer) *PeerExpressRouteCircuitConnectionsServerTransport {
	return &PeerExpressRouteCircuitConnectionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.PeerExpressRouteCircuitConnectionsClientListResponse]](),
	}
}

// PeerExpressRouteCircuitConnectionsServerTransport connects instances of armnetwork.PeerExpressRouteCircuitConnectionsClient to instances of PeerExpressRouteCircuitConnectionsServer.
// Don't use this type directly, use NewPeerExpressRouteCircuitConnectionsServerTransport instead.
type PeerExpressRouteCircuitConnectionsServerTransport struct {
	srv          *PeerExpressRouteCircuitConnectionsServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.PeerExpressRouteCircuitConnectionsClientListResponse]]
}

// Do implements the policy.Transporter interface for PeerExpressRouteCircuitConnectionsServerTransport.
func (p *PeerExpressRouteCircuitConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PeerExpressRouteCircuitConnectionsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PeerExpressRouteCircuitConnectionsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PeerExpressRouteCircuitConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	circuitNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
	if err != nil {
		return nil, err
	}
	peeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
	if err != nil {
		return nil, err
	}
	connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, circuitNameParam, peeringNameParam, connectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PeerExpressRouteCircuitConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PeerExpressRouteCircuitConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		peeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, circuitNameParam, peeringNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.PeerExpressRouteCircuitConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}
