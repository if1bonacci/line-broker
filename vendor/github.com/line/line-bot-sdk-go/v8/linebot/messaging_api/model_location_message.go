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

// LocationMessage
// LocationMessage
// https://developers.line.biz/en/reference/messaging-api/#location-message
type LocationMessage struct {
	Message

	/**
	 * Get QuickReply
	 */
	QuickReply *QuickReply `json:"quickReply,omitempty"`

	/**
	 * Get Sender
	 */
	Sender *Sender `json:"sender,omitempty"`

	/**
	 * Get Title
	 */
	Title string `json:"title"`

	/**
	 * Get Address
	 */
	Address string `json:"address"`

	/**
	 * Get Latitude
	 */
	Latitude float64 `json:"latitude"`

	/**
	 * Get Longitude
	 */
	Longitude float64 `json:"longitude"`
}

// MarshalJSON customizes the JSON serialization of the LocationMessage struct.
func (r *LocationMessage) MarshalJSON() ([]byte, error) {

	type Alias LocationMessage
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "location",
	})
}
