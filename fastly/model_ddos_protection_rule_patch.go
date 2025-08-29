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

// DdosProtectionRulePatch struct for DdosProtectionRulePatch
type DdosProtectionRulePatch struct {
	// Action types for a rule. Supported action values are default, block, log, off. The default action value follows the current protection mode of the associated service.
	Action               *string `json:"action,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionRulePatch DdosProtectionRulePatch

// NewDdosProtectionRulePatch instantiates a new DdosProtectionRulePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionRulePatch() *DdosProtectionRulePatch {
	this := DdosProtectionRulePatch{}
	var action string = "default"
	this.Action = &action
	return &this
}

// NewDdosProtectionRulePatchWithDefaults instantiates a new DdosProtectionRulePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionRulePatchWithDefaults() *DdosProtectionRulePatch {
	this := DdosProtectionRulePatch{}
	var action string = "default"
	this.Action = &action
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *DdosProtectionRulePatch) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRulePatch) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *DdosProtectionRulePatch) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *DdosProtectionRulePatch) SetAction(v string) {
	o.Action = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionRulePatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionRulePatch) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionRulePatch := _DdosProtectionRulePatch{}

	if err = json.Unmarshal(bytes, &varDdosProtectionRulePatch); err == nil {
		*o = DdosProtectionRulePatch(varDdosProtectionRulePatch)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionRulePatch is a helper abstraction for handling nullable ddosprotectionrulepatch types.
type NullableDdosProtectionRulePatch struct {
	value *DdosProtectionRulePatch
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionRulePatch) Get() *DdosProtectionRulePatch {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionRulePatch) Set(val *DdosProtectionRulePatch) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionRulePatch) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionRulePatch) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionRulePatch returns a pointer to a new instance of NullableDdosProtectionRulePatch.
func NewNullableDdosProtectionRulePatch(val *DdosProtectionRulePatch) *NullableDdosProtectionRulePatch {
	return &NullableDdosProtectionRulePatch{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionRulePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionRulePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
