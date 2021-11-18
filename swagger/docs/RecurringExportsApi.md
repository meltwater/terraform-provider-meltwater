# {{classname}}

All URIs are relative to *https://api.meltwater.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecurringExport**](RecurringExportsApi.md#CreateRecurringExport) | **Post** /v3/exports/recurring | Creates a new recurring export
[**DeleteRecurringExport**](RecurringExportsApi.md#DeleteRecurringExport) | **Delete** /v3/exports/recurring/{id} | Removes an existing recurring export
[**ListRecurringExports**](RecurringExportsApi.md#ListRecurringExports) | **Get** /v3/exports/recurring | Get a list of all your recurring exports
[**ShowRecurringExport**](RecurringExportsApi.md#ShowRecurringExport) | **Get** /v3/exports/recurring/{id} | Get details of a recurring export

# **CreateRecurringExport**
> RecurringExportResponse CreateRecurringExport(ctx, optional)
Creates a new recurring export

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecurringExportsApiCreateRecurringExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecurringExportsApiCreateRecurringExportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RecurringExportRequest**](RecurringExportRequest.md)| Export object to create.

The &#x60;timezone&#x60; field specifies what timezone the &#x60;window_time&#x60; field should be interpreted as.

The &#x60;window_time_unit&#x60; field determines the size of the export window and often the export will be refreshed.

Please see the Model Schema for possible values of &#x60;window_time_unit&#x60;.
 | 
 **companyId** | **optional.**| Company which owns the given search_id. If not specified your default company is used. | 

### Return type

[**RecurringExportResponse**](RecurringExportResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRecurringExport**
> DeleteRecurringExport(ctx, id)
Removes an existing recurring export

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Export id | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRecurringExports**
> RecurringExportListResponse ListRecurringExports(ctx, )
Get a list of all your recurring exports

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RecurringExportListResponse**](RecurringExportListResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowRecurringExport**
> RecurringExportResponse ShowRecurringExport(ctx, id)
Get details of a recurring export

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Export id | 

### Return type

[**RecurringExportResponse**](RecurringExportResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

