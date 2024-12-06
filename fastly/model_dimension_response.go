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

// DimensionResponse struct for DimensionResponse
type DimensionResponse struct {
	// The HTTP reason phrase for this dimension.
	Response             *string `json:"response,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionResponse DimensionResponse

// NewDimensionResponse instantiates a new DimensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionResponse() *DimensionResponse {
	this := DimensionResponse{}
	return &this
}

// NewDimensionResponseWithDefaults instantiates a new DimensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionResponseWithDefaults() *DimensionResponse {
	this := DimensionResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DimensionResponse) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionResponse) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DimensionResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *DimensionResponse) SetResponse(v string) {
	o.Response = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionResponse := _DimensionResponse{}

	if err = json.Unmarshal(bytes, &varDimensionResponse); err == nil {
		*o = DimensionResponse(varDimensionResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "response")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionResponse is a helper abstraction for handling nullable dimensionresponse types.
type NullableDimensionResponse struct {
	value *DimensionResponse
	isSet bool
}

// Get returns the value.
func (v NullableDimensionResponse) Get() *DimensionResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionResponse) Set(val *DimensionResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionResponse returns a pointer to a new instance of NullableDimensionResponse.
func NewNullableDimensionResponse(val *DimensionResponse) *NullableDimensionResponse {
	return &NullableDimensionResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
