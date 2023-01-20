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

// AppRoleWriteSecretIDResponse struct for AppRoleWriteSecretIDResponse
type AppRoleWriteSecretIDResponse struct {
	// Secret ID attached to the role.
	SecretId string `json:"secret_id"`
	// Accessor of the secret ID
	SecretIdAccessor string `json:"secret_id_accessor"`
	// Number of times a secret ID can access the role, after which the secret ID will expire.
	SecretIdNumUses int32 `json:"secret_id_num_uses"`
	// Duration in seconds after which the issued secret ID expires.
	SecretIdTtl int32 `json:"secret_id_ttl"`
}

// NewAppRoleWriteSecretIDResponseWithDefaults instantiates a new AppRoleWriteSecretIDResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteSecretIDResponseWithDefaults() *AppRoleWriteSecretIDResponse {
	var this AppRoleWriteSecretIDResponse

	return &this
}

func (o AppRoleWriteSecretIDResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id"] = o.SecretId
	toSerialize["secret_id_accessor"] = o.SecretIdAccessor
	toSerialize["secret_id_num_uses"] = o.SecretIdNumUses
	toSerialize["secret_id_ttl"] = o.SecretIdTtl

	return json.Marshal(toSerialize)
}