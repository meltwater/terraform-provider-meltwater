# RecurringExport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowWeekday** | **int32** | An integer representation of what day of the week the export window starts. Integer values are: Monday&#x3D;1, Tuesday&#x3D;2, Wednesday&#x3D;3, Thursday&#x3D;4, Friday&#x3D;5, Saturday&#x3D;6, Sunday&#x3D;7 Defaults to 1 | [optional] [default to 1]
**WindowTimeUnit** | **string** | Window Time Unit. The time unit the export window represents. Defaults to \&quot;DAY\&quot; | [optional] [default to WINDOW_TIME_UNIT.DAY]
**WindowTime** | **string** | Window Time. The time of day the export window begins. Defaults to \&quot;00:00:00\&quot; | [optional] [default to 00:00:00]
**WindowSize** | **int32** | Window size. The number of &#x27;window_time_unit&#x27;s the export window covers Defaults to 1 | [optional] [default to 1]
**WindowMonthday** | **int32** | An integer representation of what day of the month the export window starts. Integer values 1-28 represent the day of the month e.g. 1 represents 1st, 2 represents the 2nd, etc. Integer value 0 represents the last day of the month. Defaults to 1 | [optional] [default to 1]
**Timezone** | **string** | Timezone for window_time field. Must be a valid timezone in the IANA database. Defaults to \&quot;Etc/UTC\&quot; | [optional] [default to Etc/UTC]
**Tags** | [**[]Tag**](Tag.md) | List of tag names | [optional] [default to null]
**SearchIds** | **[]int32** | List of search ids | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

