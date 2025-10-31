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

// ValidatorResultData struct for ValidatorResultData
type ValidatorResultData struct {
	Id                   *string                        `json:"id,omitempty"`
	Type                 *string                        `json:"type,omitempty"`
	Attributes           *ValidatorResultDataAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _ValidatorResultData ValidatorResultData

// NewValidatorResultData instantiates a new ValidatorResultData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatorResultData() *ValidatorResultData {
	this := ValidatorResultData{}
	return &this
}

// NewValidatorResultDataWithDefaults instantiates a new ValidatorResultData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatorResultDataWithDefaults() *ValidatorResultData {
	this := ValidatorResultData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ValidatorResultData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatorResultData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ValidatorResultData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ValidatorResultData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ValidatorResultData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatorResultData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ValidatorResultData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ValidatorResultData) SetType(v string) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ValidatorResultData) GetAttributes() ValidatorResultDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ValidatorResultDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatorResultData) GetAttributesOk() (*ValidatorResultDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ValidatorResultData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ValidatorResultDataAttributes and assigns it to the Attributes field.
func (o *ValidatorResultData) SetAttributes(v ValidatorResultDataAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValidatorResultData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ValidatorResultData) UnmarshalJSON(bytes []byte) (err error) {
	varValidatorResultData := _ValidatorResultData{}

	if err = json.Unmarshal(bytes, &varValidatorResultData); err == nil {
		*o = ValidatorResultData(varValidatorResultData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValidatorResultData is a helper abstraction for handling nullable validatorresultdata types.
type NullableValidatorResultData struct {
	value *ValidatorResultData
	isSet bool
}

// Get returns the value.
func (v NullableValidatorResultData) Get() *ValidatorResultData {
	return v.value
}

// Set modifies the value.
func (v *NullableValidatorResultData) Set(val *ValidatorResultData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValidatorResultData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValidatorResultData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValidatorResultData returns a pointer to a new instance of NullableValidatorResultData.
func NewNullableValidatorResultData(val *ValidatorResultData) *NullableValidatorResultData {
	return &NullableValidatorResultData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValidatorResultData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValidatorResultData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
