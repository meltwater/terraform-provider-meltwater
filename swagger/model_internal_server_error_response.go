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

// Internal Server error
type InternalServerErrorResponse struct {
	// List of problems with the request
	Errors []ErrorObject `json:"errors,omitempty"`
}
