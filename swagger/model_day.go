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

// Day
type Day struct {
	// An ISO8601 date, excluding time and timezone/offset.  The year, month and day formatted according to ISO8601. The date represents the local date at the timezone given in the `tz` parameter.
	Date string `json:"date"`
	DocumentCount int64 `json:"document_count"`
	Hours []Hour `json:"hours"`
}