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

// DefaultSettingsError struct for DefaultSettingsError
type DefaultSettingsError struct {
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	Detail *string `json:"detail,omitempty"`
	AdditionalProperties map[string]any
}

type _DefaultSettingsError DefaultSettingsError

// NewDefaultSettingsError instantiates a new DefaultSettingsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultSettingsError() *DefaultSettingsError {
	this := DefaultSettingsError{}
	return &this
}

// NewDefaultSettingsErrorWithDefaults instantiates a new DefaultSettingsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultSettingsErrorWithDefaults() *DefaultSettingsError {
	this := DefaultSettingsError{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DefaultSettingsError) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultSettingsError) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DefaultSettingsError) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DefaultSettingsError) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DefaultSettingsError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultSettingsError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DefaultSettingsError) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DefaultSettingsError) SetType(v string) {
	o.Type = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *DefaultSettingsError) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultSettingsError) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *DefaultSettingsError) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *DefaultSettingsError) SetDetail(v string) {
	o.Detail = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DefaultSettingsError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *DefaultSettingsError) UnmarshalJSON(bytes []byte) (err error) {
	varDefaultSettingsError := _DefaultSettingsError{}

	if err = json.Unmarshal(bytes, &varDefaultSettingsError); err == nil {
		*o = DefaultSettingsError(varDefaultSettingsError)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDefaultSettingsError is a helper abstraction for handling nullable defaultsettingserror types. 
type NullableDefaultSettingsError struct {
	value *DefaultSettingsError
	isSet bool
}

// Get returns the value.
func (v NullableDefaultSettingsError) Get() *DefaultSettingsError {
	return v.value
}

// Set modifies the value.
func (v *NullableDefaultSettingsError) Set(val *DefaultSettingsError) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDefaultSettingsError) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDefaultSettingsError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDefaultSettingsError returns a pointer to a new instance of NullableDefaultSettingsError.
func NewNullableDefaultSettingsError(val *DefaultSettingsError) *NullableDefaultSettingsError {
	return &NullableDefaultSettingsError{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDefaultSettingsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableDefaultSettingsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
