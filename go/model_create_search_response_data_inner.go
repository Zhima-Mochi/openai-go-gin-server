/*
 * OpenAI API
 *
 * APIs for sampling from and fine-tuning language models
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateSearchResponseDataInner struct {

	Object string `json:"object,omitempty"`

	Document int32 `json:"document,omitempty"`

	Score float32 `json:"score,omitempty"`
}
