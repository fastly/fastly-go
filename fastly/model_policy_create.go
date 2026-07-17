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

// PolicyCreate struct for PolicyCreate
type PolicyCreate struct {
	Name                 string      `json:"name"`
	Description          *string     `json:"description,omitempty"`
	Mode                 string      `json:"mode"`
	Directives           []Directive `json:"directives,omitempty"`
	AdditionalProperties map[string]any
}

type _PolicyCreate PolicyCreate

// NewPolicyCreate instantiates a new PolicyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyCreate(name string, mode string) *PolicyCreate {
	this := PolicyCreate{}
	this.Name = name
	this.Mode = mode
	return &this
}

// NewPolicyCreateWithDefaults instantiates a new PolicyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyCreateWithDefaults() *PolicyCreate {
	this := PolicyCreate{}
	return &this
}

// GetName returns the Name field value
func (o *PolicyCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicyCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicyCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PolicyCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PolicyCreate) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value
func (o *PolicyCreate) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *PolicyCreate) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *PolicyCreate) SetMode(v string) {
	o.Mode = v
}

// GetDirectives returns the Directives field value if set, zero value otherwise.
func (o *PolicyCreate) GetDirectives() []Directive {
	if o == nil || o.Directives == nil {
		var ret []Directive
		return ret
	}
	return o.Directives
}

// GetDirectivesOk returns a tuple with the Directives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyCreate) GetDirectivesOk() ([]Directive, bool) {
	if o == nil || o.Directives == nil {
		return nil, false
	}
	return o.Directives, true
}

// HasDirectives returns a boolean if a field has been set.
func (o *PolicyCreate) HasDirectives() bool {
	if o != nil && o.Directives != nil {
		return true
	}

	return false
}

// SetDirectives gets a reference to the given []Directive and assigns it to the Directives field.
func (o *PolicyCreate) SetDirectives(v []Directive) {
	o.Directives = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PolicyCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.Directives != nil {
		toSerialize["directives"] = o.Directives
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PolicyCreate) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyCreate := _PolicyCreate{}

	if err = json.Unmarshal(bytes, &varPolicyCreate); err == nil {
		*o = PolicyCreate(varPolicyCreate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "directives")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePolicyCreate is a helper abstraction for handling nullable policycreate types.
type NullablePolicyCreate struct {
	value *PolicyCreate
	isSet bool
}

// Get returns the value.
func (v NullablePolicyCreate) Get() *PolicyCreate {
	return v.value
}

// Set modifies the value.
func (v *NullablePolicyCreate) Set(val *PolicyCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePolicyCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePolicyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePolicyCreate returns a pointer to a new instance of NullablePolicyCreate.
func NewNullablePolicyCreate(val *PolicyCreate) *NullablePolicyCreate {
	return &NullablePolicyCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePolicyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePolicyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
