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

// ErrorResponseData struct for ErrorResponseData
type ErrorResponseData struct {
	Title *string `json:"title,omitempty"`
	Detail *string `json:"detail,omitempty"`
	AdditionalProperties map[string]any
}

type _ErrorResponseData ErrorResponseData

// NewErrorResponseData instantiates a new ErrorResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseData() *ErrorResponseData {
	this := ErrorResponseData{}
	return &this
}

// NewErrorResponseDataWithDefaults instantiates a new ErrorResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseDataWithDefaults() *ErrorResponseData {
	this := ErrorResponseData{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ErrorResponseData) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseData) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ErrorResponseData) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ErrorResponseData) SetTitle(v string) {
	o.Title = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ErrorResponseData) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseData) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorResponseData) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ErrorResponseData) SetDetail(v string) {
	o.Detail = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ErrorResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
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
func (o *ErrorResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varErrorResponseData := _ErrorResponseData{}

	if err = json.Unmarshal(bytes, &varErrorResponseData); err == nil {
		*o = ErrorResponseData(varErrorResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableErrorResponseData is a helper abstraction for handling nullable errorresponsedata types. 
type NullableErrorResponseData struct {
	value *ErrorResponseData
	isSet bool
}

// Get returns the value.
func (v NullableErrorResponseData) Get() *ErrorResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableErrorResponseData) Set(val *ErrorResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableErrorResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableErrorResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableErrorResponseData returns a pointer to a new instance of NullableErrorResponseData.
func NewNullableErrorResponseData(val *ErrorResponseData) *NullableErrorResponseData {
	return &NullableErrorResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableErrorResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableErrorResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
