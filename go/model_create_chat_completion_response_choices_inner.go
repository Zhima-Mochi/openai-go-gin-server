/*
 * OpenAI API
 *
 * APIs for sampling from and fine-tuning language models
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateChatCompletionResponseChoicesInner struct {

	Index int32 `json:"index,omitempty"`

	Message ChatCompletionResponseMessage `json:"message,omitempty"`

	FinishReason string `json:"finish_reason,omitempty"`
}
