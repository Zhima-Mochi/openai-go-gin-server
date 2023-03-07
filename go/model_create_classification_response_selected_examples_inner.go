/*
 * OpenAI API
 *
 * APIs for sampling from and fine-tuning language models
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateClassificationResponseSelectedExamplesInner struct {

	Document int32 `json:"document,omitempty"`

	Text string `json:"text,omitempty"`

	Label string `json:"label,omitempty"`
}