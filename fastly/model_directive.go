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

// Directive struct for Directive
type Directive struct {
	// CSP directive name (e.g., script-src, style-src)
	Name *string `json:"name,omitempty"`
	// Directive values
	Values               []string `json:"values,omitempty"`
	AdditionalProperties map[string]any
}

type _Directive Directive

// NewDirective instantiates a new Directive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirective() *Directive {
	this := Directive{}
	return &this
}

// NewDirectiveWithDefaults instantiates a new Directive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectiveWithDefaults() *Directive {
	this := Directive{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Directive) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Directive) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Directive) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Directive) SetName(v string) {
	o.Name = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Directive) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Directive) GetValuesOk() ([]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Directive) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *Directive) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Directive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
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
func (o *Directive) UnmarshalJSON(bytes []byte) (err error) {
	varDirective := _Directive{}

	if err = json.Unmarshal(bytes, &varDirective); err == nil {
		*o = Directive(varDirective)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDirective is a helper abstraction for handling nullable directive types.
type NullableDirective struct {
	value *Directive
	isSet bool
}

// Get returns the value.
func (v NullableDirective) Get() *Directive {
	return v.value
}

// Set modifies the value.
func (v *NullableDirective) Set(val *Directive) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDirective) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDirective) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDirective returns a pointer to a new instance of NullableDirective.
func NewNullableDirective(val *Directive) *NullableDirective {
	return &NullableDirective{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDirective) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDirective) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
