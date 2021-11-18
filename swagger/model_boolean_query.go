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

type BooleanQuery struct {
	// The optional source selection id for the search query (mi only)
	SourceSelectionId int64 `json:"source_selection_id,omitempty"`
	FilterSet *FilterSet `json:"filter_set,omitempty"`
	// The case sensitivity for this query
	CaseSensitivity string `json:"case_sensitivity"`
	// The boolean query string
	Boolean string `json:"boolean"`
}