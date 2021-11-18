# {{classname}}

All URIs are relative to *https://api.meltwater.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSearch**](SearchesApi.md#CreateSearch) | **Post** /v3/searches | Create a search
[**DeleteSearch**](SearchesApi.md#DeleteSearch) | **Delete** /v3/searches/{id} | Delete an individual search
[**GetSearch**](SearchesApi.md#GetSearch) | **Get** /v3/searches/{id} | Get an individual search
[**ListSearches**](SearchesApi.md#ListSearches) | **Get** /v3/searches | Get a list of all your searches
[**SearchCount**](SearchesApi.md#SearchCount) | **Get** /v3/searches/{id}/count | Get an approximate count of results for the search over a particular period
[**UpdateSearch**](SearchesApi.md#UpdateSearch) | **Put** /v3/searches/{id} | Update an individual search

# **CreateSearch**
> SingleSearch CreateSearch(ctx, body, optional)
Create a search

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SingleSearchRequest**](SingleSearchRequest.md)| A single search request | 
 **optional** | ***SearchesApiCreateSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchesApiCreateSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **optional.**| Company which owns the given search. If not specified your default company is used. | 
 **dryRun** | **optional.**| If set to \&quot;true\&quot; only performs validation of the search, does not create the resource on success. Default \&quot;false\&quot;. | 

### Return type

[**SingleSearch**](SingleSearch.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSearch**
> DeleteSearch(ctx, id, optional)
Delete an individual search

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Search id | 
 **optional** | ***SearchesApiDeleteSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchesApiDeleteSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **optional.String**| Company which owns the given search. If not specified your default company is used. | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSearch**
> SingleSearch GetSearch(ctx, id, optional)
Get an individual search

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Search id | 
 **optional** | ***SearchesApiGetSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchesApiGetSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **optional.String**| Company which owns the given search. If not specified your default company is used. | 

### Return type

[**SingleSearch**](SingleSearch.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSearches**
> SearchListResponse ListSearches(ctx, optional)
Get a list of all your searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchesApiListSearchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchesApiListSearchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **optional.String**| Company which owns the given search. If not specified your default company is used. | 
 **includeQuery** | **optional.Bool**| Include query object in searches, default \&quot;false\&quot; | 
 **queryType** | **optional.String**| Only return searches of the specified query type, if omitted searches of all query types will be returned | 

### Return type

[**SearchListResponse**](SearchListResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCount**
> SearchCountResponse SearchCount(ctx, id, optional)
Get an approximate count of results for the search over a particular period

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Search id | 
 **optional** | ***SearchesApiSearchCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchesApiSearchCountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **optional.String**| Company which owns the given search. If not specified your default company is used. | 
 **startDate** | **optional.Time**| The start date of the search to count result for. Default: &#x60;two weeks ago&#x60; in ISO8601 format. | 
 **endDate** | **optional.Time**| The end date of the search to count result for. Default: &#x60;now&#x60; in ISO8601 format. | 

### Return type

[**SearchCountResponse**](SearchCountResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSearch**
> SingleSearch UpdateSearch(ctx, body, id, optional)
Update an individual search

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SingleSearchRequest**](SingleSearchRequest.md)| A single search request | 
  **id** | **int64**| Search id | 
 **optional** | ***SearchesApiUpdateSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchesApiUpdateSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **companyId** | **optional.**| Company which owns the given search. If not specified your default company is used. | 
 **dryRun** | **optional.**| If set to \&quot;true\&quot; only performs validation of the search, does not create the resource on success. Default \&quot;false\&quot;. | 

### Return type

[**SingleSearch**](SingleSearch.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

