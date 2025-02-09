//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DenyAssignmentsClient contains the methods for the DenyAssignments group.
// Don't use this type directly, use NewDenyAssignmentsClient() instead.
type DenyAssignmentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDenyAssignmentsClient creates a new instance of DenyAssignmentsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDenyAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DenyAssignmentsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DenyAssignmentsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get the specified deny assignment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01
// scope - The scope of the deny assignment.
// denyAssignmentID - The ID of the deny assignment to get.
// options - DenyAssignmentsClientGetOptions contains the optional parameters for the DenyAssignmentsClient.Get method.
func (client *DenyAssignmentsClient) Get(ctx context.Context, scope string, denyAssignmentID string, options *DenyAssignmentsClientGetOptions) (DenyAssignmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, denyAssignmentID, options)
	if err != nil {
		return DenyAssignmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DenyAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DenyAssignmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DenyAssignmentsClient) getCreateRequest(ctx context.Context, scope string, denyAssignmentID string, options *DenyAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/denyAssignments/{denyAssignmentId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if denyAssignmentID == "" {
		return nil, errors.New("parameter denyAssignmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{denyAssignmentId}", url.PathEscape(denyAssignmentID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DenyAssignmentsClient) getHandleResponse(resp *http.Response) (DenyAssignmentsClientGetResponse, error) {
	result := DenyAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DenyAssignment); err != nil {
		return DenyAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// GetByID - Gets a deny assignment by ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01
// denyAssignmentID - The fully qualified deny assignment ID. For example, use the format, /subscriptions/{guid}/providers/Microsoft.Authorization/denyAssignments/{denyAssignmentId}
// for subscription level deny assignments,
// or /providers/Microsoft.Authorization/denyAssignments/{denyAssignmentId} for tenant level deny assignments.
// options - DenyAssignmentsClientGetByIDOptions contains the optional parameters for the DenyAssignmentsClient.GetByID method.
func (client *DenyAssignmentsClient) GetByID(ctx context.Context, denyAssignmentID string, options *DenyAssignmentsClientGetByIDOptions) (DenyAssignmentsClientGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, denyAssignmentID, options)
	if err != nil {
		return DenyAssignmentsClientGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DenyAssignmentsClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DenyAssignmentsClientGetByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *DenyAssignmentsClient) getByIDCreateRequest(ctx context.Context, denyAssignmentID string, options *DenyAssignmentsClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/{denyAssignmentId}"
	urlPath = strings.ReplaceAll(urlPath, "{denyAssignmentId}", denyAssignmentID)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *DenyAssignmentsClient) getByIDHandleResponse(resp *http.Response) (DenyAssignmentsClientGetByIDResponse, error) {
	result := DenyAssignmentsClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DenyAssignment); err != nil {
		return DenyAssignmentsClientGetByIDResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all deny assignments for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01
// options - DenyAssignmentsClientListOptions contains the optional parameters for the DenyAssignmentsClient.List method.
func (client *DenyAssignmentsClient) NewListPager(options *DenyAssignmentsClientListOptions) *runtime.Pager[DenyAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DenyAssignmentsClientListResponse]{
		More: func(page DenyAssignmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DenyAssignmentsClientListResponse) (DenyAssignmentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DenyAssignmentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DenyAssignmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DenyAssignmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DenyAssignmentsClient) listCreateRequest(ctx context.Context, options *DenyAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/denyAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DenyAssignmentsClient) listHandleResponse(resp *http.Response) (DenyAssignmentsClientListResponse, error) {
	result := DenyAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DenyAssignmentListResult); err != nil {
		return DenyAssignmentsClientListResponse{}, err
	}
	return result, nil
}

// NewListForResourcePager - Gets deny assignments for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceProviderNamespace - The namespace of the resource provider.
// parentResourcePath - The parent resource identity.
// resourceType - The resource type of the resource.
// resourceName - The name of the resource to get deny assignments for.
// options - DenyAssignmentsClientListForResourceOptions contains the optional parameters for the DenyAssignmentsClient.ListForResource
// method.
func (client *DenyAssignmentsClient) NewListForResourcePager(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *DenyAssignmentsClientListForResourceOptions) *runtime.Pager[DenyAssignmentsClientListForResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DenyAssignmentsClientListForResourceResponse]{
		More: func(page DenyAssignmentsClientListForResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DenyAssignmentsClientListForResourceResponse) (DenyAssignmentsClientListForResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForResourceCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DenyAssignmentsClientListForResourceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DenyAssignmentsClientListForResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DenyAssignmentsClientListForResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForResourceHandleResponse(resp)
		},
	})
}

// listForResourceCreateRequest creates the ListForResource request.
func (client *DenyAssignmentsClient) listForResourceCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *DenyAssignmentsClientListForResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/denyAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", resourceProviderNamespace)
	urlPath = strings.ReplaceAll(urlPath, "{parentResourcePath}", parentResourcePath)
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", resourceType)
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForResourceHandleResponse handles the ListForResource response.
func (client *DenyAssignmentsClient) listForResourceHandleResponse(resp *http.Response) (DenyAssignmentsClientListForResourceResponse, error) {
	result := DenyAssignmentsClientListForResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DenyAssignmentListResult); err != nil {
		return DenyAssignmentsClientListForResourceResponse{}, err
	}
	return result, nil
}

// NewListForResourceGroupPager - Gets deny assignments for a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - DenyAssignmentsClientListForResourceGroupOptions contains the optional parameters for the DenyAssignmentsClient.ListForResourceGroup
// method.
func (client *DenyAssignmentsClient) NewListForResourceGroupPager(resourceGroupName string, options *DenyAssignmentsClientListForResourceGroupOptions) *runtime.Pager[DenyAssignmentsClientListForResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DenyAssignmentsClientListForResourceGroupResponse]{
		More: func(page DenyAssignmentsClientListForResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DenyAssignmentsClientListForResourceGroupResponse) (DenyAssignmentsClientListForResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DenyAssignmentsClientListForResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DenyAssignmentsClientListForResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DenyAssignmentsClientListForResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForResourceGroupHandleResponse(resp)
		},
	})
}

// listForResourceGroupCreateRequest creates the ListForResourceGroup request.
func (client *DenyAssignmentsClient) listForResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DenyAssignmentsClientListForResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/denyAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForResourceGroupHandleResponse handles the ListForResourceGroup response.
func (client *DenyAssignmentsClient) listForResourceGroupHandleResponse(resp *http.Response) (DenyAssignmentsClientListForResourceGroupResponse, error) {
	result := DenyAssignmentsClientListForResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DenyAssignmentListResult); err != nil {
		return DenyAssignmentsClientListForResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets deny assignments for a scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01
// scope - The scope of the deny assignments.
// options - DenyAssignmentsClientListForScopeOptions contains the optional parameters for the DenyAssignmentsClient.ListForScope
// method.
func (client *DenyAssignmentsClient) NewListForScopePager(scope string, options *DenyAssignmentsClientListForScopeOptions) *runtime.Pager[DenyAssignmentsClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[DenyAssignmentsClientListForScopeResponse]{
		More: func(page DenyAssignmentsClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DenyAssignmentsClientListForScopeResponse) (DenyAssignmentsClientListForScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForScopeCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DenyAssignmentsClientListForScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DenyAssignmentsClientListForScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DenyAssignmentsClientListForScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForScopeHandleResponse(resp)
		},
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *DenyAssignmentsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *DenyAssignmentsClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/denyAssignments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *DenyAssignmentsClient) listForScopeHandleResponse(resp *http.Response) (DenyAssignmentsClientListForScopeResponse, error) {
	result := DenyAssignmentsClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DenyAssignmentListResult); err != nil {
		return DenyAssignmentsClientListForScopeResponse{}, err
	}
	return result, nil
}
