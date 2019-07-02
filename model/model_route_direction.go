/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 
type RouteDirection struct {
	// The distance in the route at which the action applies
	Distance float32 `json:"distance,omitempty"`
	// The action of this direction
	Action int32 `json:"action,omitempty"`
	Name string `json:"name,omitempty"`
}
