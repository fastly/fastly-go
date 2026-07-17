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

// InlineResponse20016 struct for InlineResponse20016
type InlineResponse20016 struct {
	Results              []Suggestion `json:"results,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse20016 InlineResponse20016

// NewInlineResponse20016 instantiates a new InlineResponse20016 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016() *InlineResponse20016 {
	this := InlineResponse20016{}
	return &this
}

// NewInlineResponse20016WithDefaults instantiates a new InlineResponse20016 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016WithDefaults() *InlineResponse20016 {
	this := InlineResponse20016{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InlineResponse20016) GetResults() []Suggestion {
	if o == nil || o.Results == nil {
		var ret []Suggestion
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016) GetResultsOk() ([]Suggestion, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InlineResponse20016) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Suggestion and assigns it to the Results field.
func (o *InlineResponse20016) SetResults(v []Suggestion) {
	o.Results = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse20016) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *InlineResponse20016) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse20016 := _InlineResponse20016{}

	if err = json.Unmarshal(bytes, &varInlineResponse20016); err == nil {
		*o = InlineResponse20016(varInlineResponse20016)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse20016 is a helper abstraction for handling nullable inlineresponse20016 types.
type NullableInlineResponse20016 struct {
	value *InlineResponse20016
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse20016) Get() *InlineResponse20016 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse20016) Set(val *InlineResponse20016) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse20016) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse20016) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse20016 returns a pointer to a new instance of NullableInlineResponse20016.
func NewNullableInlineResponse20016(val *InlineResponse20016) *NullableInlineResponse20016 {
	return &NullableInlineResponse20016{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse20016) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse20016) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
