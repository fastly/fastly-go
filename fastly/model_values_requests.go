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

// ValuesRequests struct for ValuesRequests
type ValuesRequests struct {
	// The percentage of all requests made to the URL in the current dimension.
	RequestPercentage    *float32 `json:"request_percentage,omitempty"`
	AdditionalProperties map[string]any
}

type _ValuesRequests ValuesRequests

// NewValuesRequests instantiates a new ValuesRequests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuesRequests() *ValuesRequests {
	this := ValuesRequests{}
	return &this
}

// NewValuesRequestsWithDefaults instantiates a new ValuesRequests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesRequestsWithDefaults() *ValuesRequests {
	this := ValuesRequests{}
	return &this
}

// GetRequestPercentage returns the RequestPercentage field value if set, zero value otherwise.
func (o *ValuesRequests) GetRequestPercentage() float32 {
	if o == nil || o.RequestPercentage == nil {
		var ret float32
		return ret
	}
	return *o.RequestPercentage
}

// GetRequestPercentageOk returns a tuple with the RequestPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesRequests) GetRequestPercentageOk() (*float32, bool) {
	if o == nil || o.RequestPercentage == nil {
		return nil, false
	}
	return o.RequestPercentage, true
}

// HasRequestPercentage returns a boolean if a field has been set.
func (o *ValuesRequests) HasRequestPercentage() bool {
	if o != nil && o.RequestPercentage != nil {
		return true
	}

	return false
}

// SetRequestPercentage gets a reference to the given float32 and assigns it to the RequestPercentage field.
func (o *ValuesRequests) SetRequestPercentage(v float32) {
	o.RequestPercentage = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValuesRequests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.RequestPercentage != nil {
		toSerialize["request_percentage"] = o.RequestPercentage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ValuesRequests) UnmarshalJSON(bytes []byte) (err error) {
	varValuesRequests := _ValuesRequests{}

	if err = json.Unmarshal(bytes, &varValuesRequests); err == nil {
		*o = ValuesRequests(varValuesRequests)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_percentage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValuesRequests is a helper abstraction for handling nullable valuesrequests types.
type NullableValuesRequests struct {
	value *ValuesRequests
	isSet bool
}

// Get returns the value.
func (v NullableValuesRequests) Get() *ValuesRequests {
	return v.value
}

// Set modifies the value.
func (v *NullableValuesRequests) Set(val *ValuesRequests) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValuesRequests) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValuesRequests) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValuesRequests returns a pointer to a new instance of NullableValuesRequests.
func NewNullableValuesRequests(val *ValuesRequests) *NullableValuesRequests {
	return &NullableValuesRequests{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValuesRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValuesRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
