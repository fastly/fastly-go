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

// OperationBulkCreate struct for OperationBulkCreate
type OperationBulkCreate struct {
	// List of operations to create.
	Operations           []OperationBulkCreateOperations `json:"operations"`
	AdditionalProperties map[string]any
}

type _OperationBulkCreate OperationBulkCreate

// NewOperationBulkCreate instantiates a new OperationBulkCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationBulkCreate(operations []OperationBulkCreateOperations) *OperationBulkCreate {
	this := OperationBulkCreate{}
	this.Operations = operations
	return &this
}

// NewOperationBulkCreateWithDefaults instantiates a new OperationBulkCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationBulkCreateWithDefaults() *OperationBulkCreate {
	this := OperationBulkCreate{}
	return &this
}

// GetOperations returns the Operations field value
func (o *OperationBulkCreate) GetOperations() []OperationBulkCreateOperations {
	if o == nil {
		var ret []OperationBulkCreateOperations
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *OperationBulkCreate) GetOperationsOk() ([]OperationBulkCreateOperations, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *OperationBulkCreate) SetOperations(v []OperationBulkCreateOperations) {
	o.Operations = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OperationBulkCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["operations"] = o.Operations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OperationBulkCreate) UnmarshalJSON(bytes []byte) (err error) {
	varOperationBulkCreate := _OperationBulkCreate{}

	if err = json.Unmarshal(bytes, &varOperationBulkCreate); err == nil {
		*o = OperationBulkCreate(varOperationBulkCreate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOperationBulkCreate is a helper abstraction for handling nullable operationbulkcreate types.
type NullableOperationBulkCreate struct {
	value *OperationBulkCreate
	isSet bool
}

// Get returns the value.
func (v NullableOperationBulkCreate) Get() *OperationBulkCreate {
	return v.value
}

// Set modifies the value.
func (v *NullableOperationBulkCreate) Set(val *OperationBulkCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOperationBulkCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOperationBulkCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOperationBulkCreate returns a pointer to a new instance of NullableOperationBulkCreate.
func NewNullableOperationBulkCreate(val *OperationBulkCreate) *NullableOperationBulkCreate {
	return &NullableOperationBulkCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOperationBulkCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOperationBulkCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
