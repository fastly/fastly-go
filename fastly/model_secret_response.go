// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
	"time"
)

// SecretResponse struct for SecretResponse
type SecretResponse struct {
	// Name of the secret.
	Name *string `json:"name,omitempty"`
	// An opaque identifier of the plaintext secret value. This can be used to determine if a secret value has changed.
	Digest *string `json:"digest,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// True if the secret replaced a secret with the same name.
	Recreated NullableBool `json:"recreated,omitempty"`
	AdditionalProperties map[string]any
}

type _SecretResponse SecretResponse

// NewSecretResponse instantiates a new SecretResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretResponse() *SecretResponse {
	this := SecretResponse{}
	return &this
}

// NewSecretResponseWithDefaults instantiates a new SecretResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretResponseWithDefaults() *SecretResponse {
	this := SecretResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecretResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecretResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecretResponse) SetName(v string) {
	o.Name = &v
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *SecretResponse) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretResponse) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *SecretResponse) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *SecretResponse) SetDigest(v string) {
	o.Digest = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecretResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecretResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SecretResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *SecretResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *SecretResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *SecretResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetRecreated returns the Recreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecretResponse) GetRecreated() bool {
	if o == nil || o.Recreated.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Recreated.Get()
}

// GetRecreatedOk returns a tuple with the Recreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecretResponse) GetRecreatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Recreated.Get(), o.Recreated.IsSet()
}

// HasRecreated returns a boolean if a field has been set.
func (o *SecretResponse) HasRecreated() bool {
	if o != nil && o.Recreated.IsSet() {
		return true
	}

	return false
}

// SetRecreated gets a reference to the given NullableBool and assigns it to the Recreated field.
func (o *SecretResponse) SetRecreated(v bool) {
	o.Recreated.Set(&v)
}
// SetRecreatedNil sets the value for Recreated to be an explicit nil
func (o *SecretResponse) SetRecreatedNil() {
	o.Recreated.Set(nil)
}

// UnsetRecreated ensures that no value is present for Recreated, not even an explicit nil
func (o *SecretResponse) UnsetRecreated() {
	o.Recreated.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SecretResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Recreated.IsSet() {
		toSerialize["recreated"] = o.Recreated.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *SecretResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSecretResponse := _SecretResponse{}

	if err = json.Unmarshal(bytes, &varSecretResponse); err == nil {
		*o = SecretResponse(varSecretResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "digest")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "recreated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSecretResponse is a helper abstraction for handling nullable secretresponse types. 
type NullableSecretResponse struct {
	value *SecretResponse
	isSet bool
}

// Get returns the value.
func (v NullableSecretResponse) Get() *SecretResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableSecretResponse) Set(val *SecretResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSecretResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSecretResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSecretResponse returns a pointer to a new instance of NullableSecretResponse.
func NewNullableSecretResponse(val *SecretResponse) *NullableSecretResponse {
	return &NullableSecretResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSecretResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSecretResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
