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

// BillingEstimateResponseAllOfLines struct for BillingEstimateResponseAllOfLines
type BillingEstimateResponseAllOfLines struct {
	Line *BillingEstimateResponseAllOfLine `json:"line,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingEstimateResponseAllOfLines BillingEstimateResponseAllOfLines

// NewBillingEstimateResponseAllOfLines instantiates a new BillingEstimateResponseAllOfLines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingEstimateResponseAllOfLines() *BillingEstimateResponseAllOfLines {
	this := BillingEstimateResponseAllOfLines{}
	return &this
}

// NewBillingEstimateResponseAllOfLinesWithDefaults instantiates a new BillingEstimateResponseAllOfLines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingEstimateResponseAllOfLinesWithDefaults() *BillingEstimateResponseAllOfLines {
	this := BillingEstimateResponseAllOfLines{}
	return &this
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOfLines) GetLine() BillingEstimateResponseAllOfLine {
	if o == nil || o.Line == nil {
		var ret BillingEstimateResponseAllOfLine
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOfLines) GetLineOk() (*BillingEstimateResponseAllOfLine, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOfLines) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given BillingEstimateResponseAllOfLine and assigns it to the Line field.
func (o *BillingEstimateResponseAllOfLines) SetLine(v BillingEstimateResponseAllOfLine) {
	o.Line = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingEstimateResponseAllOfLines) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingEstimateResponseAllOfLines) UnmarshalJSON(bytes []byte) (err error) {
	varBillingEstimateResponseAllOfLines := _BillingEstimateResponseAllOfLines{}

	if err = json.Unmarshal(bytes, &varBillingEstimateResponseAllOfLines); err == nil {
		*o = BillingEstimateResponseAllOfLines(varBillingEstimateResponseAllOfLines)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "line")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingEstimateResponseAllOfLines is a helper abstraction for handling nullable billingestimateresponsealloflines types. 
type NullableBillingEstimateResponseAllOfLines struct {
	value *BillingEstimateResponseAllOfLines
	isSet bool
}

// Get returns the value.
func (v NullableBillingEstimateResponseAllOfLines) Get() *BillingEstimateResponseAllOfLines {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingEstimateResponseAllOfLines) Set(val *BillingEstimateResponseAllOfLines) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingEstimateResponseAllOfLines) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingEstimateResponseAllOfLines) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingEstimateResponseAllOfLines returns a pointer to a new instance of NullableBillingEstimateResponseAllOfLines.
func NewNullableBillingEstimateResponseAllOfLines(val *BillingEstimateResponseAllOfLines) *NullableBillingEstimateResponseAllOfLines {
	return &NullableBillingEstimateResponseAllOfLines{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingEstimateResponseAllOfLines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingEstimateResponseAllOfLines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
