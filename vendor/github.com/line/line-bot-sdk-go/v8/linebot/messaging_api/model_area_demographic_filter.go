/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package messaging_api

import (
	"encoding/json"
)

// AreaDemographicFilter
// AreaDemographicFilter

type AreaDemographicFilter struct {
	DemographicFilter

	/**
	 * Get OneOf
	 */
	OneOf []AreaDemographic `json:"oneOf,omitempty"`
}

// MarshalJSON customizes the JSON serialization of the AreaDemographicFilter struct.
func (r *AreaDemographicFilter) MarshalJSON() ([]byte, error) {

	type Alias AreaDemographicFilter
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type,omitempty"`
	}{
		Alias: (*Alias)(r),

		Type: "area",
	})
}