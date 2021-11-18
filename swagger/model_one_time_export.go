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

// A one-time export
type OneTimeExport struct {
	// List of tag names
	Tags []Tag `json:"tags,omitempty"`
	// Start date (UTC) in ISO 8601 standard. E.g. \"2020-03-25T00:00:00Z\"
	StartDate string `json:"start_date"`
	// List of search ids
	SearchIds []int32 `json:"search_ids"`
	// End date (UTC) in ISO 8601 standard. E.g. \"2020-03-29T23:59:59Z\"
	EndDate string `json:"end_date"`
}