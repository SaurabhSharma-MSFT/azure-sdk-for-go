//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armelastic

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VMHostClient contains the methods for the VMHost group.
// Don't use this type directly, use NewVMHostClient() instead.
type VMHostClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVMHostClient creates a new instance of VMHostClient with the specified values.
func NewVMHostClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VMHostClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &VMHostClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - List the vm resources currently being monitored by the Elastic monitor resource.
// If the operation fails it returns the *ResourceProviderDefaultErrorResponse error type.
func (client *VMHostClient) List(resourceGroupName string, monitorName string, options *VMHostListOptions) *VMHostListPager {
	return &VMHostListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp VMHostListResponseEnvelope) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VMHostListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VMHostClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *VMHostListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/listVMHost"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VMHostClient) listHandleResponse(resp *http.Response) (VMHostListResponseEnvelope, error) {
	result := VMHostListResponseEnvelope{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMHostListResponse); err != nil {
		return VMHostListResponseEnvelope{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VMHostClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ResourceProviderDefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
