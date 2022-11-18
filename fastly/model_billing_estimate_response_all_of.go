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

// BillingEstimateResponseAllOf struct for BillingEstimateResponseAllOf
type BillingEstimateResponseAllOf struct {
	Lines []BillingEstimateResponseAllOfLines `json:"lines,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingEstimateResponseAllOf BillingEstimateResponseAllOf

// NewBillingEstimateResponseAllOf instantiates a new BillingEstimateResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingEstimateResponseAllOf() *BillingEstimateResponseAllOf {
	this := BillingEstimateResponseAllOf{}
	return &this
}

// NewBillingEstimateResponseAllOfWithDefaults instantiates a new BillingEstimateResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingEstimateResponseAllOfWithDefaults() *BillingEstimateResponseAllOf {
	this := BillingEstimateResponseAllOf{}
	return &this
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOf) GetLines() []BillingEstimateResponseAllOfLines {
	if o == nil || o.Lines == nil {
		var ret []BillingEstimateResponseAllOfLines
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOf) GetLinesOk() ([]BillingEstimateResponseAllOfLines, bool) {
	if o == nil || o.Lines == nil {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOf) HasLines() bool {
	if o != nil && o.Lines != nil {
		return true
	}

	return false
}

// SetLines gets a reference to the given []BillingEstimateResponseAllOfLines and assigns it to the Lines field.
func (o *BillingEstimateResponseAllOf) SetLines(v []BillingEstimateResponseAllOfLines) {
	o.Lines = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingEstimateResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Lines != nil {
		toSerialize["lines"] = o.Lines
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingEstimateResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBillingEstimateResponseAllOf := _BillingEstimateResponseAllOf{}

	if err = json.Unmarshal(bytes, &varBillingEstimateResponseAllOf); err == nil {
		*o = BillingEstimateResponseAllOf(varBillingEstimateResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "lines")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingEstimateResponseAllOf is a helper abstraction for handling nullable billingestimateresponseallof types. 
type NullableBillingEstimateResponseAllOf struct {
	value *BillingEstimateResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableBillingEstimateResponseAllOf) Get() *BillingEstimateResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingEstimateResponseAllOf) Set(val *BillingEstimateResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingEstimateResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingEstimateResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingEstimateResponseAllOf returns a pointer to a new instance of NullableBillingEstimateResponseAllOf.
func NewNullableBillingEstimateResponseAllOf(val *BillingEstimateResponseAllOf) *NullableBillingEstimateResponseAllOf {
	return &NullableBillingEstimateResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingEstimateResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingEstimateResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
