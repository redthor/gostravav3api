/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

import (
	"time"
)

type RunningRace struct {
	// The unique identifier of this race.
	Id int32 `json:"id,omitempty"`
	// The name of this race.
	Name string `json:"name,omitempty"`
	// The type of this race.
	RunningRaceType int32 `json:"running_race_type,omitempty"`
	// The race's distance, in meters.
	Distance float32 `json:"distance,omitempty"`
	// The time at which the race begins started in the local timezone.
	StartDateLocal time.Time `json:"start_date_local,omitempty"`
	// The name of the city in which the race is taking place.
	City string `json:"city,omitempty"`
	// The name of the state or geographical region in which the race is taking place.
	State string `json:"state,omitempty"`
	// The name of the country in which the race is taking place.
	Country string `json:"country,omitempty"`
	// The set of routes that cover this race's course.
	RouteIds []int32 `json:"route_ids,omitempty"`
	// The unit system in which the race should be displayed.
	MeasurementPreference string `json:"measurement_preference,omitempty"`
	// The vanity URL of this race on Strava.
	Url string `json:"url,omitempty"`
	// The URL of this race's website.
	WebsiteUrl string `json:"website_url,omitempty"`
}
