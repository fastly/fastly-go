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

// Listinvoices struct for Listinvoices
type Listinvoices struct {
	Data []Invoice `json:"data,omitempty"`
	Meta *Metadata `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _Listinvoices Listinvoices

// NewListinvoices instantiates a new Listinvoices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListinvoices() *Listinvoices {
	this := Listinvoices{}
	return &this
}

// NewListinvoicesWithDefaults instantiates a new Listinvoices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListinvoicesWithDefaults() *Listinvoices {
	this := Listinvoices{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Listinvoices) GetData() []Invoice {
	if o == nil || o.Data == nil {
		var ret []Invoice
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listinvoices) GetDataOk() ([]Invoice, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Listinvoices) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Invoice and assigns it to the Data field.
func (o *Listinvoices) SetData(v []Invoice) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Listinvoices) GetMeta() Metadata {
	if o == nil || o.Meta == nil {
		var ret Metadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listinvoices) GetMetaOk() (*Metadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Listinvoices) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Metadata and assigns it to the Meta field.
func (o *Listinvoices) SetMeta(v Metadata) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Listinvoices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Listinvoices) UnmarshalJSON(bytes []byte) (err error) {
	varListinvoices := _Listinvoices{}

	if err = json.Unmarshal(bytes, &varListinvoices); err == nil {
		*o = Listinvoices(varListinvoices)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableListinvoices is a helper abstraction for handling nullable listinvoices types. 
type NullableListinvoices struct {
	value *Listinvoices
	isSet bool
}

// Get returns the value.
func (v NullableListinvoices) Get() *Listinvoices {
	return v.value
}

// Set modifies the value.
func (v *NullableListinvoices) Set(val *Listinvoices) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableListinvoices) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableListinvoices) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableListinvoices returns a pointer to a new instance of NullableListinvoices.
func NewNullableListinvoices(val *Listinvoices) *NullableListinvoices {
	return &NullableListinvoices{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableListinvoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableListinvoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
