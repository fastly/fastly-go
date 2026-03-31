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

// OperationCreateExtra struct for OperationCreateExtra
type OperationCreateExtra struct {
	// The status to assign to the operation. Defaults to SAVED if omitted.
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]any
}

type _OperationCreateExtra OperationCreateExtra

// NewOperationCreateExtra instantiates a new OperationCreateExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationCreateExtra() *OperationCreateExtra {
	this := OperationCreateExtra{}
	var status string = "SAVED"
	this.Status = &status
	return &this
}

// NewOperationCreateExtraWithDefaults instantiates a new OperationCreateExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationCreateExtraWithDefaults() *OperationCreateExtra {
	this := OperationCreateExtra{}
	var status string = "SAVED"
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OperationCreateExtra) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationCreateExtra) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OperationCreateExtra) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OperationCreateExtra) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OperationCreateExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OperationCreateExtra) UnmarshalJSON(bytes []byte) (err error) {
	varOperationCreateExtra := _OperationCreateExtra{}

	if err = json.Unmarshal(bytes, &varOperationCreateExtra); err == nil {
		*o = OperationCreateExtra(varOperationCreateExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOperationCreateExtra is a helper abstraction for handling nullable operationcreateextra types.
type NullableOperationCreateExtra struct {
	value *OperationCreateExtra
	isSet bool
}

// Get returns the value.
func (v NullableOperationCreateExtra) Get() *OperationCreateExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableOperationCreateExtra) Set(val *OperationCreateExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOperationCreateExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOperationCreateExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOperationCreateExtra returns a pointer to a new instance of NullableOperationCreateExtra.
func NewNullableOperationCreateExtra(val *OperationCreateExtra) *NullableOperationCreateExtra {
	return &NullableOperationCreateExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOperationCreateExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOperationCreateExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
