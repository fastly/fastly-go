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

// OperationBulkDelete struct for OperationBulkDelete
type OperationBulkDelete struct {
	// List of operation IDs to delete.
	OperationIds         []string `json:"operation_ids"`
	AdditionalProperties map[string]any
}

type _OperationBulkDelete OperationBulkDelete

// NewOperationBulkDelete instantiates a new OperationBulkDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationBulkDelete(operationIds []string) *OperationBulkDelete {
	this := OperationBulkDelete{}
	this.OperationIds = operationIds
	return &this
}

// NewOperationBulkDeleteWithDefaults instantiates a new OperationBulkDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationBulkDeleteWithDefaults() *OperationBulkDelete {
	this := OperationBulkDelete{}
	return &this
}

// GetOperationIds returns the OperationIds field value
func (o *OperationBulkDelete) GetOperationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OperationIds
}

// GetOperationIdsOk returns a tuple with the OperationIds field value
// and a boolean to check if the value has been set.
func (o *OperationBulkDelete) GetOperationIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationIds, true
}

// SetOperationIds sets field value
func (o *OperationBulkDelete) SetOperationIds(v []string) {
	o.OperationIds = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OperationBulkDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["operation_ids"] = o.OperationIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OperationBulkDelete) UnmarshalJSON(bytes []byte) (err error) {
	varOperationBulkDelete := _OperationBulkDelete{}

	if err = json.Unmarshal(bytes, &varOperationBulkDelete); err == nil {
		*o = OperationBulkDelete(varOperationBulkDelete)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operation_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOperationBulkDelete is a helper abstraction for handling nullable operationbulkdelete types.
type NullableOperationBulkDelete struct {
	value *OperationBulkDelete
	isSet bool
}

// Get returns the value.
func (v NullableOperationBulkDelete) Get() *OperationBulkDelete {
	return v.value
}

// Set modifies the value.
func (v *NullableOperationBulkDelete) Set(val *OperationBulkDelete) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOperationBulkDelete) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOperationBulkDelete) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOperationBulkDelete returns a pointer to a new instance of NullableOperationBulkDelete.
func NewNullableOperationBulkDelete(val *OperationBulkDelete) *NullableOperationBulkDelete {
	return &NullableOperationBulkDelete{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOperationBulkDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOperationBulkDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
