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

// ArcError An error response.
type ArcError struct {
	// A URL identifying the error type.
	Type *string `json:"type,omitempty"`
	// A short summary of the problem.
	Title *string `json:"title,omitempty"`
	// A machine-readable error code.
	Code                 *string `json:"code,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcError ArcError

// NewArcError instantiates a new ArcError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcError() *ArcError {
	this := ArcError{}
	return &this
}

// NewArcErrorWithDefaults instantiates a new ArcError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcErrorWithDefaults() *ArcError {
	this := ArcError{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ArcError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ArcError) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ArcError) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ArcError) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcError) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ArcError) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ArcError) SetTitle(v string) {
	o.Title = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ArcError) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcError) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ArcError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ArcError) SetCode(v string) {
	o.Code = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcError) UnmarshalJSON(bytes []byte) (err error) {
	varArcError := _ArcError{}

	if err = json.Unmarshal(bytes, &varArcError); err == nil {
		*o = ArcError(varArcError)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "title")
		delete(additionalProperties, "code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcError is a helper abstraction for handling nullable arcerror types.
type NullableArcError struct {
	value *ArcError
	isSet bool
}

// Get returns the value.
func (v NullableArcError) Get() *ArcError {
	return v.value
}

// Set modifies the value.
func (v *NullableArcError) Set(val *ArcError) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcError) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcError returns a pointer to a new instance of NullableArcError.
func NewNullableArcError(val *ArcError) *NullableArcError {
	return &NullableArcError{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
