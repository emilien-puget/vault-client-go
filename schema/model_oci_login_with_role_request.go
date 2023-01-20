/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OCILoginWithRoleRequest struct for OCILoginWithRoleRequest
type OCILoginWithRoleRequest struct {
	// The signed headers of the client
	RequestHeaders string `json:"request_headers"`
}

// NewOCILoginWithRoleRequestWithDefaults instantiates a new OCILoginWithRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCILoginWithRoleRequestWithDefaults() *OCILoginWithRoleRequest {
	var this OCILoginWithRoleRequest

	return &this
}

func (o OCILoginWithRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["request_headers"] = o.RequestHeaders

	return json.Marshal(toSerialize)
}