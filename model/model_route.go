/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type Route struct {
	Athlete *SummaryAthlete `json:"athlete,omitempty"`
	// The description of the route
	Description string `json:"description,omitempty"`
	// The route's distance, in meters
	Distance float32 `json:"distance,omitempty"`
	// The route's elevation gain.
	ElevationGain float32 `json:"elevation_gain,omitempty"`
	// The unique identifier of this route
	Id int32 `json:"id,omitempty"`
	Map_ *PolylineMap `json:"map,omitempty"`
	// The name of this route
	Name string `json:"name,omitempty"`
	// Whether this route is private
	Private bool `json:"private,omitempty"`
	// Whether this route is starred by the logged-in athlete
	Starred bool `json:"starred,omitempty"`
	Timestamp int32 `json:"timestamp,omitempty"`
	// This route's type (1 for ride, 2 for runs)
	Type_ int32 `json:"type,omitempty"`
	// This route's sub-type (1 for road, 2 for mountain bike, 3 for cross, 4 for trail, 5 for mixed)
	SubType int32 `json:"sub_type,omitempty"`
	// The segments traversed by this route
	Segments []SummarySegment `json:"segments,omitempty"`
	// The directions of this route
	Directions []RouteDirection `json:"directions,omitempty"`
}
