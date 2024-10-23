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

// Batch struct for Batch
type Batch struct {
	// A descriptor for the response of the entire batch
	Title *string `json:"title,omitempty"`
	// If an error is present in any of the requests, this field will describe that error
	Type *string `json:"type,omitempty"`
	// Per-key errors which failed to parse, validate, or otherwise transmit
	Errors               []BatchErrors `json:"errors,omitempty"`
	AdditionalProperties map[string]any
}

type _Batch Batch

// NewBatch instantiates a new Batch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatch() *Batch {
	this := Batch{}
	return &this
}

// NewBatchWithDefaults instantiates a new Batch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchWithDefaults() *Batch {
	this := Batch{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Batch) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Batch) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Batch) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Batch) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Batch) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Batch) SetType(v string) {
	o.Type = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Batch) GetErrors() []BatchErrors {
	if o == nil || o.Errors == nil {
		var ret []BatchErrors
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetErrorsOk() ([]BatchErrors, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Batch) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []BatchErrors and assigns it to the Errors field.
func (o *Batch) SetErrors(v []BatchErrors) {
	o.Errors = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Batch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Batch) UnmarshalJSON(bytes []byte) (err error) {
	varBatch := _Batch{}

	if err = json.Unmarshal(bytes, &varBatch); err == nil {
		*o = Batch(varBatch)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBatch is a helper abstraction for handling nullable batch types.
type NullableBatch struct {
	value *Batch
	isSet bool
}

// Get returns the value.
func (v NullableBatch) Get() *Batch {
	return v.value
}

// Set modifies the value.
func (v *NullableBatch) Set(val *Batch) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBatch) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBatch) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBatch returns a pointer to a new instance of NullableBatch.
func NewNullableBatch(val *Batch) *NullableBatch {
	return &NullableBatch{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
