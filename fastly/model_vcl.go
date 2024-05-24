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

// Vcl struct for Vcl
type Vcl struct {
	// The VCL code to be included.
	Content *string `json:"content,omitempty"`
	// Set to `true` when this is the main VCL, otherwise `false`.
	Main *bool `json:"main,omitempty"`
	// The name of this VCL.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _Vcl Vcl

// NewVcl instantiates a new Vcl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcl() *Vcl {
	this := Vcl{}
	return &this
}

// NewVclWithDefaults instantiates a new Vcl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVclWithDefaults() *Vcl {
	this := Vcl{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Vcl) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vcl) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Vcl) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Vcl) SetContent(v string) {
	o.Content = &v
}

// GetMain returns the Main field value if set, zero value otherwise.
func (o *Vcl) GetMain() bool {
	if o == nil || o.Main == nil {
		var ret bool
		return ret
	}
	return *o.Main
}

// GetMainOk returns a tuple with the Main field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vcl) GetMainOk() (*bool, bool) {
	if o == nil || o.Main == nil {
		return nil, false
	}
	return o.Main, true
}

// HasMain returns a boolean if a field has been set.
func (o *Vcl) HasMain() bool {
	if o != nil && o.Main != nil {
		return true
	}

	return false
}

// SetMain gets a reference to the given bool and assigns it to the Main field.
func (o *Vcl) SetMain(v bool) {
	o.Main = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Vcl) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vcl) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Vcl) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Vcl) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Vcl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Main != nil {
		toSerialize["main"] = o.Main
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Vcl) UnmarshalJSON(bytes []byte) (err error) {
	varVcl := _Vcl{}

	if err = json.Unmarshal(bytes, &varVcl); err == nil {
		*o = Vcl(varVcl)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "main")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableVcl is a helper abstraction for handling nullable vcl types. 
type NullableVcl struct {
	value *Vcl
	isSet bool
}

// Get returns the value.
func (v NullableVcl) Get() *Vcl {
	return v.value
}

// Set modifies the value.
func (v *NullableVcl) Set(val *Vcl) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableVcl) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableVcl) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableVcl returns a pointer to a new instance of NullableVcl.
func NewNullableVcl(val *Vcl) *NullableVcl {
	return &NullableVcl{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableVcl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableVcl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
