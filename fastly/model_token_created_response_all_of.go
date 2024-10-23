// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://www.fastly.com/documentation/reference/api/)

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.

import (
	"encoding/json"
)

// TokenCreatedResponseAllOf struct for TokenCreatedResponseAllOf
type TokenCreatedResponseAllOf struct {
	// The alphanumeric string for accessing the API (only available on token creation).
	AccessToken          *string `json:"access_token,omitempty"`
	AdditionalProperties map[string]any
}

type _TokenCreatedResponseAllOf TokenCreatedResponseAllOf

// NewTokenCreatedResponseAllOf instantiates a new TokenCreatedResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenCreatedResponseAllOf() *TokenCreatedResponseAllOf {
	this := TokenCreatedResponseAllOf{}
	return &this
}

// NewTokenCreatedResponseAllOfWithDefaults instantiates a new TokenCreatedResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCreatedResponseAllOfWithDefaults() *TokenCreatedResponseAllOf {
	this := TokenCreatedResponseAllOf{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TokenCreatedResponseAllOf) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCreatedResponseAllOf) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TokenCreatedResponseAllOf) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TokenCreatedResponseAllOf) SetAccessToken(v string) {
	o.AccessToken = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TokenCreatedResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TokenCreatedResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTokenCreatedResponseAllOf := _TokenCreatedResponseAllOf{}

	if err = json.Unmarshal(bytes, &varTokenCreatedResponseAllOf); err == nil {
		*o = TokenCreatedResponseAllOf(varTokenCreatedResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTokenCreatedResponseAllOf is a helper abstraction for handling nullable tokencreatedresponseallof types.
type NullableTokenCreatedResponseAllOf struct {
	value *TokenCreatedResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTokenCreatedResponseAllOf) Get() *TokenCreatedResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTokenCreatedResponseAllOf) Set(val *TokenCreatedResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTokenCreatedResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTokenCreatedResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTokenCreatedResponseAllOf returns a pointer to a new instance of NullableTokenCreatedResponseAllOf.
func NewNullableTokenCreatedResponseAllOf(val *TokenCreatedResponseAllOf) *NullableTokenCreatedResponseAllOf {
	return &NullableTokenCreatedResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTokenCreatedResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTokenCreatedResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
