# OneTimeExportResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | **string** | The datetime the export was last updated | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | List of tag names used for the export | [optional] [default to null]
**StatusReason** | **string** | Optional explanation for current status | [optional] [default to null]
**Status** | **string** | Current status of the one-time export | [optional] [default to null]
**StartDate** | **string** | Start date of the export (UTC) in ISO 8601 format | [default to null]
**Searches** | [**[]SearchDetails**](SearchDetails.md) | List of searches used for the export | [default to null]
**InsertedAt** | **string** | The datetime the export was created | [optional] [default to null]
**Id** | **int32** | ID of the export | [default to null]
**EndDate** | **string** | End date of the export (UTC) in ISO 8601 format | [default to null]
**DataUrl** | **string** | The URL at which the export data can be retrieved | [optional] [default to null]
**CompanyName** | **string** | Company the export belongs to | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

