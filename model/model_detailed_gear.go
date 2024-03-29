/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type DetailedGear struct {
	// The gear's unique identifier.
	Id string `json:"id,omitempty"`
	// Resource state, indicates level of detail. Possible values: 2 -> \"summary\", 3 -> \"detail\"
	ResourceState int32 `json:"resource_state,omitempty"`
	// Whether this gear's is the owner's default one.
	Primary bool `json:"primary,omitempty"`
	// The gear's name.
	Name string `json:"name,omitempty"`
	// The distance logged with this gear.
	Distance float32 `json:"distance,omitempty"`
	// The gear's brand name.
	BrandName string `json:"brand_name,omitempty"`
	// The gear's model name.
	ModelName string `json:"model_name,omitempty"`
	// The gear's frame type (bike only).
	FrameType int32 `json:"frame_type,omitempty"`
	// The gear's description.
	Description string `json:"description,omitempty"`
}
