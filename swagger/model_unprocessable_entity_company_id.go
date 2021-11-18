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

// Invalid company id
type UnprocessableEntityCompanyId struct {
	// List of problems with the request
	Errors []ValidationError `json:"errors,omitempty"`
}