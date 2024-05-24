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

// TLSCsrErrorResponse struct for TLSCsrErrorResponse
type TLSCsrErrorResponse struct {
	Errors []ErrorResponseData `json:"errors,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCsrErrorResponse TLSCsrErrorResponse

// NewTLSCsrErrorResponse instantiates a new TLSCsrErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCsrErrorResponse() *TLSCsrErrorResponse {
	this := TLSCsrErrorResponse{}
	return &this
}

// NewTLSCsrErrorResponseWithDefaults instantiates a new TLSCsrErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCsrErrorResponseWithDefaults() *TLSCsrErrorResponse {
	this := TLSCsrErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *TLSCsrErrorResponse) GetErrors() []ErrorResponseData {
	if o == nil || o.Errors == nil {
		var ret []ErrorResponseData
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrErrorResponse) GetErrorsOk() ([]ErrorResponseData, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *TLSCsrErrorResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorResponseData and assigns it to the Errors field.
func (o *TLSCsrErrorResponse) SetErrors(v []ErrorResponseData) {
	o.Errors = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCsrErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *TLSCsrErrorResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCsrErrorResponse := _TLSCsrErrorResponse{}

	if err = json.Unmarshal(bytes, &varTLSCsrErrorResponse); err == nil {
		*o = TLSCsrErrorResponse(varTLSCsrErrorResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSCsrErrorResponse is a helper abstraction for handling nullable tlscsrerrorresponse types. 
type NullableTLSCsrErrorResponse struct {
	value *TLSCsrErrorResponse
	isSet bool
}

// Get returns the value.
func (v NullableTLSCsrErrorResponse) Get() *TLSCsrErrorResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCsrErrorResponse) Set(val *TLSCsrErrorResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCsrErrorResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCsrErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCsrErrorResponse returns a pointer to a new instance of NullableTLSCsrErrorResponse.
func NewNullableTLSCsrErrorResponse(val *TLSCsrErrorResponse) *NullableTLSCsrErrorResponse {
	return &NullableTLSCsrErrorResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCsrErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSCsrErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
