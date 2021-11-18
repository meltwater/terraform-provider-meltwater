# {{classname}}

All URIs are relative to *https://api.meltwater.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSourceSelections**](SourceSelectionsApi.md#ListSourceSelections) | **Get** /v3/source_selections | Get a list of all your source selections

# **ListSourceSelections**
> []SourceSelection ListSourceSelections(ctx, optional)
Get a list of all your source selections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SourceSelectionsApiListSourceSelectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceSelectionsApiListSourceSelectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **optional.String**| Company which owns the given source selections. If not specified your default company is used. | 
 **type_** | **optional.String**| Only return source selections with the specified type. If this parameter is omitted source selections of all types will be returned. | 

### Return type

[**[]SourceSelection**](array.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

