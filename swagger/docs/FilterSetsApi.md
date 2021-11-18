# {{classname}}

All URIs are relative to *https://api.meltwater.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFilterSets**](FilterSetsApi.md#ListFilterSets) | **Get** /v3/filter_sets | Get a list of all search filter sets

# **ListFilterSets**
> FilterSetList ListFilterSets(ctx, optional)
Get a list of all search filter sets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilterSetsApiListFilterSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterSetsApiListFilterSetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **optional.String**| Company which owns the given filter sets. If not specified your default company is used. | 
 **includeQuickpicks** | **optional.Bool**| Include meltwater-provided QUICKPICK filter sets, default \&quot;false\&quot; | 

### Return type

[**FilterSetList**](FilterSetList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

