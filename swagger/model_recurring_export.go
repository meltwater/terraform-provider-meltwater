/*
 * Meltwater API
 *
 * The Meltwater Public API
 *
 * API version: 1.0
 * Contact: support@api.meltwater.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A recurring export
type RecurringExport struct {
	// An integer representation of what day of the week the export window starts. Integer values are: Monday=1, Tuesday=2, Wednesday=3, Thursday=4, Friday=5, Saturday=6, Sunday=7 Defaults to 1
	WindowWeekday int32 `json:"window_weekday,omitempty"`
	// Window Time Unit. The time unit the export window represents. Defaults to \"DAY\"
	WindowTimeUnit string `json:"window_time_unit,omitempty"`
	// Window Time. The time of day the export window begins. Defaults to \"00:00:00\"
	WindowTime string `json:"window_time,omitempty"`
	// Window size. The number of 'window_time_unit's the export window covers Defaults to 1
	WindowSize int32 `json:"window_size,omitempty"`
	// An integer representation of what day of the month the export window starts. Integer values 1-28 represent the day of the month e.g. 1 represents 1st, 2 represents the 2nd, etc. Integer value 0 represents the last day of the month. Defaults to 1
	WindowMonthday int32 `json:"window_monthday,omitempty"`
	// Timezone for window_time field. Must be a valid timezone in the IANA database. Defaults to \"Etc/UTC\"
	Timezone string `json:"timezone,omitempty"`
	// List of tag names
	Tags []Tag `json:"tags,omitempty"`
	// List of search ids
	SearchIds []int32 `json:"search_ids"`
}
