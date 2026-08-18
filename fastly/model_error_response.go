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

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	// A URI reference that identifies the problem type.
	Type *string `json:"type,omitempty"`
	// A short, human-readable summary of the problem type.
	Title *string `json:"title,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail               *string               `json:"detail,omitempty"`
	Errors               []ErrorResponseErrors `json:"errors,omitempty"`
	AdditionalProperties map[string]any
}

type _ErrorResponse ErrorResponse

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrorResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrorResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErrorResponse) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ErrorResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ErrorResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ErrorResponse) SetTitle(v string) {
	o.Title = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ErrorResponse) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorResponse) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ErrorResponse) SetDetail(v string) {
	o.Detail = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ErrorResponse) GetErrors() []ErrorResponseErrors {
	if o == nil || o.Errors == nil {
		var ret []ErrorResponseErrors
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorsOk() ([]ErrorResponseErrors, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ErrorResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorResponseErrors and assigns it to the Errors field.
func (o *ErrorResponse) SetErrors(v []ErrorResponseErrors) {
	o.Errors = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ErrorResponse) UnmarshalJSON(bytes []byte) (err error) {
	varErrorResponse := _ErrorResponse{}

	if err = json.Unmarshal(bytes, &varErrorResponse); err == nil {
		*o = ErrorResponse(varErrorResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "title")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableErrorResponse is a helper abstraction for handling nullable errorresponse types.
type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

// Get returns the value.
func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableErrorResponse returns a pointer to a new instance of NullableErrorResponse.
func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
