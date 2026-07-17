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

// Position The position at which to place a new rule. At most one of `before` or `after` may be set. Cannot be used when creating a default (empty-conditions) rule.
type Position struct {
	// The id of the rule that this rule should be placed before.
	Before *string `json:"before,omitempty"`
	// The id of the rule that this rule should be placed after.
	After                *string `json:"after,omitempty"`
	AdditionalProperties map[string]any
}

type _Position Position

// NewPosition instantiates a new Position object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPosition() *Position {
	this := Position{}
	return &this
}

// NewPositionWithDefaults instantiates a new Position object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionWithDefaults() *Position {
	this := Position{}
	return &this
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *Position) GetBefore() string {
	if o == nil || o.Before == nil {
		var ret string
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Position) GetBeforeOk() (*string, bool) {
	if o == nil || o.Before == nil {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *Position) HasBefore() bool {
	if o != nil && o.Before != nil {
		return true
	}

	return false
}

// SetBefore gets a reference to the given string and assigns it to the Before field.
func (o *Position) SetBefore(v string) {
	o.Before = &v
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *Position) GetAfter() string {
	if o == nil || o.After == nil {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Position) GetAfterOk() (*string, bool) {
	if o == nil || o.After == nil {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *Position) HasAfter() bool {
	if o != nil && o.After != nil {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *Position) SetAfter(v string) {
	o.After = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Position) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Before != nil {
		toSerialize["before"] = o.Before
	}
	if o.After != nil {
		toSerialize["after"] = o.After
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Position) UnmarshalJSON(bytes []byte) (err error) {
	varPosition := _Position{}

	if err = json.Unmarshal(bytes, &varPosition); err == nil {
		*o = Position(varPosition)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "before")
		delete(additionalProperties, "after")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePosition is a helper abstraction for handling nullable position types.
type NullablePosition struct {
	value *Position
	isSet bool
}

// Get returns the value.
func (v NullablePosition) Get() *Position {
	return v.value
}

// Set modifies the value.
func (v *NullablePosition) Set(val *Position) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePosition) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePosition) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePosition returns a pointer to a new instance of NullablePosition.
func NewNullablePosition(val *Position) *NullablePosition {
	return &NullablePosition{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
