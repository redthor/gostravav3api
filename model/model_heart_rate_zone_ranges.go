/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type HeartRateZoneRanges struct {
	// Whether the athlete has set their own custom heart rate zones
	CustomZones bool `json:"custom_zones,omitempty"`
	Zones *ZoneRanges `json:"zones,omitempty"`
}
