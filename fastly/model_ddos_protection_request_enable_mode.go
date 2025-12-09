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

// DdosProtectionRequestEnableMode struct for DdosProtectionRequestEnableMode
type DdosProtectionRequestEnableMode struct {
	// Operation mode
	Mode                 *string `json:"mode,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionRequestEnableMode DdosProtectionRequestEnableMode

// NewDdosProtectionRequestEnableMode instantiates a new DdosProtectionRequestEnableMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionRequestEnableMode() *DdosProtectionRequestEnableMode {
	this := DdosProtectionRequestEnableMode{}
	return &this
}

// NewDdosProtectionRequestEnableModeWithDefaults instantiates a new DdosProtectionRequestEnableMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionRequestEnableModeWithDefaults() *DdosProtectionRequestEnableMode {
	this := DdosProtectionRequestEnableMode{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DdosProtectionRequestEnableMode) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionRequestEnableMode) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DdosProtectionRequestEnableMode) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *DdosProtectionRequestEnableMode) SetMode(v string) {
	o.Mode = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionRequestEnableMode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionRequestEnableMode) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionRequestEnableMode := _DdosProtectionRequestEnableMode{}

	if err = json.Unmarshal(bytes, &varDdosProtectionRequestEnableMode); err == nil {
		*o = DdosProtectionRequestEnableMode(varDdosProtectionRequestEnableMode)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionRequestEnableMode is a helper abstraction for handling nullable ddosprotectionrequestenablemode types.
type NullableDdosProtectionRequestEnableMode struct {
	value *DdosProtectionRequestEnableMode
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionRequestEnableMode) Get() *DdosProtectionRequestEnableMode {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionRequestEnableMode) Set(val *DdosProtectionRequestEnableMode) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionRequestEnableMode) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionRequestEnableMode) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionRequestEnableMode returns a pointer to a new instance of NullableDdosProtectionRequestEnableMode.
func NewNullableDdosProtectionRequestEnableMode(val *DdosProtectionRequestEnableMode) *NullableDdosProtectionRequestEnableMode {
	return &NullableDdosProtectionRequestEnableMode{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionRequestEnableMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionRequestEnableMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
