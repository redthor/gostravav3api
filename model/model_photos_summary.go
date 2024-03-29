/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type PhotosSummary struct {
	// The number of photos
	Count int32 `json:"count,omitempty"`
	Primary *PhotosSummaryPrimary `json:"primary,omitempty"`
}
