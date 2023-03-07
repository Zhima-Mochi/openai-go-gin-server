/*
 * OpenAI API
 *
 * APIs for sampling from and fine-tuning language models
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateChatCompletionResponse struct {

	Id string `json:"id"`

	Object string `json:"object"`

	Created int32 `json:"created"`

	Model string `json:"model"`

	Choices []CreateChatCompletionResponseChoicesInner `json:"choices"`

	Usage CreateCompletionResponseUsage `json:"usage,omitempty"`
}