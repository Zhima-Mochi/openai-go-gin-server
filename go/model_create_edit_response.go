/*
 * OpenAI API
 *
 * APIs for sampling from and fine-tuning language models
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateEditResponse struct {

	Object string `json:"object"`

	Created int32 `json:"created"`

	Choices []CreateCompletionResponseChoicesInner `json:"choices"`

	Usage CreateCompletionResponseUsage `json:"usage"`
}