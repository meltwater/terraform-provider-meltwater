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

type ErrorObject struct {
	// Category of error
	Type_ string `json:"type"`
	// Title of error
	Title string `json:"title"`
	// Details about the error
	Details string `json:"details"`
}
