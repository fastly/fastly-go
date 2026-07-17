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

// Action The action to take when a rule matches.
type Action struct {
	Type ActionType `json:"type"`
	// The value for the action. For `service` actions, this is the target Fastly service ID.
	Value                string `json:"value"`
	AdditionalProperties map[string]any
}

type _Action Action

// NewAction instantiates a new Action object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAction(type_ ActionType, value string) *Action {
	this := Action{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewActionWithDefaults instantiates a new Action object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionWithDefaults() *Action {
	this := Action{}
	return &this
}

// GetType returns the Type field value
func (o *Action) GetType() ActionType {
	if o == nil {
		var ret ActionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Action) GetTypeOk() (*ActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Action) SetType(v ActionType) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *Action) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Action) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Action) SetValue(v string) {
	o.Value = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Action) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Action) UnmarshalJSON(bytes []byte) (err error) {
	varAction := _Action{}

	if err = json.Unmarshal(bytes, &varAction); err == nil {
		*o = Action(varAction)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAction is a helper abstraction for handling nullable action types.
type NullableAction struct {
	value *Action
	isSet bool
}

// Get returns the value.
func (v NullableAction) Get() *Action {
	return v.value
}

// Set modifies the value.
func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAction) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAction returns a pointer to a new instance of NullableAction.
func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
