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

// InlineResponse2006 struct for InlineResponse2006
type InlineResponse2006 struct {
	Results              []Suggestion `json:"results,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse2006 InlineResponse2006

// NewInlineResponse2006 instantiates a new InlineResponse2006 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2006() *InlineResponse2006 {
	this := InlineResponse2006{}
	return &this
}

// NewInlineResponse2006WithDefaults instantiates a new InlineResponse2006 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2006WithDefaults() *InlineResponse2006 {
	this := InlineResponse2006{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InlineResponse2006) GetResults() []Suggestion {
	if o == nil || o.Results == nil {
		var ret []Suggestion
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetResultsOk() ([]Suggestion, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InlineResponse2006) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Suggestion and assigns it to the Results field.
func (o *InlineResponse2006) SetResults(v []Suggestion) {
	o.Results = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse2006) MarshalJSON() ([]byte, error) {
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
func (o *InlineResponse2006) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2006 := _InlineResponse2006{}

	if err = json.Unmarshal(bytes, &varInlineResponse2006); err == nil {
		*o = InlineResponse2006(varInlineResponse2006)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse2006 is a helper abstraction for handling nullable inlineresponse2006 types.
type NullableInlineResponse2006 struct {
	value *InlineResponse2006
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse2006) Get() *InlineResponse2006 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse2006) Set(val *InlineResponse2006) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse2006) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse2006) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse2006 returns a pointer to a new instance of NullableInlineResponse2006.
func NewNullableInlineResponse2006(val *InlineResponse2006) *NullableInlineResponse2006 {
	return &NullableInlineResponse2006{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse2006) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse2006) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
