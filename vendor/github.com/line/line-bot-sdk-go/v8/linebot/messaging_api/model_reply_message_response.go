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

// ReplyMessageResponse
// ReplyMessageResponse
// https://developers.line.biz/en/reference/messaging-api/#send-reply-message-response
type ReplyMessageResponse struct {

	/**
	 * Array of sent messages. (Required)
	 */
	SentMessages []SentMessage `json:"sentMessages"`
}