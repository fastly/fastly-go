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

// TLSActivationResponse struct for TLSActivationResponse
type TLSActivationResponse struct {
	Data                 *TLSActivationResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSActivationResponse TLSActivationResponse

// NewTLSActivationResponse instantiates a new TLSActivationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSActivationResponse() *TLSActivationResponse {
	this := TLSActivationResponse{}
	return &this
}

// NewTLSActivationResponseWithDefaults instantiates a new TLSActivationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSActivationResponseWithDefaults() *TLSActivationResponse {
	this := TLSActivationResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TLSActivationResponse) GetData() TLSActivationResponseData {
	if o == nil || o.Data == nil {
		var ret TLSActivationResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSActivationResponse) GetDataOk() (*TLSActivationResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TLSActivationResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given TLSActivationResponseData and assigns it to the Data field.
func (o *TLSActivationResponse) SetData(v TLSActivationResponseData) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSActivationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TLSActivationResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTLSActivationResponse := _TLSActivationResponse{}

	if err = json.Unmarshal(bytes, &varTLSActivationResponse); err == nil {
		*o = TLSActivationResponse(varTLSActivationResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSActivationResponse is a helper abstraction for handling nullable tlsactivationresponse types.
type NullableTLSActivationResponse struct {
	value *TLSActivationResponse
	isSet bool
}

// Get returns the value.
func (v NullableTLSActivationResponse) Get() *TLSActivationResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSActivationResponse) Set(val *TLSActivationResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSActivationResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSActivationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSActivationResponse returns a pointer to a new instance of NullableTLSActivationResponse.
func NewNullableTLSActivationResponse(val *TLSActivationResponse) *NullableTLSActivationResponse {
	return &NullableTLSActivationResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSActivationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSActivationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
