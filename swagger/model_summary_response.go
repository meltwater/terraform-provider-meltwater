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

// SummaryResponse
type SummaryResponse struct {
	UniqueAuthors int64 `json:"unique_authors"`
	TopCountries []CountryStats `json:"top_countries"`
	Sentiment *Sentiments `json:"sentiment"`
	TopLanguages []LanguageStats `json:"top_languages"`
	TimeSeries []Day `json:"time_series"`
	Volume *Volume `json:"volume"`
}