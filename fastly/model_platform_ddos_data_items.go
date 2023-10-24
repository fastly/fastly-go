// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// PlatformDdosDataItems struct for PlatformDdosDataItems
type PlatformDdosDataItems struct {
	Values *Values `json:"values,omitempty"`
	AdditionalProperties map[string]any
}

type _PlatformDdosDataItems PlatformDdosDataItems

// NewPlatformDdosDataItems instantiates a new PlatformDdosDataItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformDdosDataItems() *PlatformDdosDataItems {
	this := PlatformDdosDataItems{}
	return &this
}

// NewPlatformDdosDataItemsWithDefaults instantiates a new PlatformDdosDataItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformDdosDataItemsWithDefaults() *PlatformDdosDataItems {
	this := PlatformDdosDataItems{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *PlatformDdosDataItems) GetValues() Values {
	if o == nil || o.Values == nil {
		var ret Values
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformDdosDataItems) GetValuesOk() (*Values, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *PlatformDdosDataItems) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given Values and assigns it to the Values field.
func (o *PlatformDdosDataItems) SetValues(v Values) {
	o.Values = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PlatformDdosDataItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *PlatformDdosDataItems) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformDdosDataItems := _PlatformDdosDataItems{}

	if err = json.Unmarshal(bytes, &varPlatformDdosDataItems); err == nil {
		*o = PlatformDdosDataItems(varPlatformDdosDataItems)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePlatformDdosDataItems is a helper abstraction for handling nullable platformddosdataitems types. 
type NullablePlatformDdosDataItems struct {
	value *PlatformDdosDataItems
	isSet bool
}

// Get returns the value.
func (v NullablePlatformDdosDataItems) Get() *PlatformDdosDataItems {
	return v.value
}

// Set modifies the value.
func (v *NullablePlatformDdosDataItems) Set(val *PlatformDdosDataItems) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePlatformDdosDataItems) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePlatformDdosDataItems) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePlatformDdosDataItems returns a pointer to a new instance of NullablePlatformDdosDataItems.
func NewNullablePlatformDdosDataItems(val *PlatformDdosDataItems) *NullablePlatformDdosDataItems {
	return &NullablePlatformDdosDataItems{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePlatformDdosDataItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePlatformDdosDataItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
