/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type PhotosSummaryPrimary struct {
	Id int64 `json:"id,omitempty"`
	Source int32 `json:"source,omitempty"`
	UniqueId string `json:"unique_id,omitempty"`
	Urls map[string]string `json:"urls,omitempty"`
}
