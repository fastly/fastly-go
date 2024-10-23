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

// WafExclusionResponseDataAttributesAllOf struct for WafExclusionResponseDataAttributesAllOf
type WafExclusionResponseDataAttributesAllOf struct {
	// A conditional expression in VCL used to determine if the condition is met.
	Condition *string `json:"condition,omitempty"`
	// The type of exclusion.
	ExclusionType *string `json:"exclusion_type,omitempty"`
	// Whether to generate a log upon matching.
	Logging *bool `json:"logging,omitempty"`
	// Name of the exclusion.
	Name *string `json:"name,omitempty"`
	// A numeric ID identifying a WAF exclusion.
	Number *int32 `json:"number,omitempty"`
	// The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type=variable`.
	Variable             NullableString `json:"variable,omitempty"`
	AdditionalProperties map[string]any
}

type _WafExclusionResponseDataAttributesAllOf WafExclusionResponseDataAttributesAllOf

// NewWafExclusionResponseDataAttributesAllOf instantiates a new WafExclusionResponseDataAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafExclusionResponseDataAttributesAllOf() *WafExclusionResponseDataAttributesAllOf {
	this := WafExclusionResponseDataAttributesAllOf{}
	var logging bool = true
	this.Logging = &logging
	return &this
}

// NewWafExclusionResponseDataAttributesAllOfWithDefaults instantiates a new WafExclusionResponseDataAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafExclusionResponseDataAttributesAllOfWithDefaults() *WafExclusionResponseDataAttributesAllOf {
	this := WafExclusionResponseDataAttributesAllOf{}
	var logging bool = true
	this.Logging = &logging
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAttributesAllOf) GetCondition() string {
	if o == nil || o.Condition == nil {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAttributesAllOf) GetConditionOk() (*string, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributesAllOf) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *WafExclusionResponseDataAttributesAllOf) SetCondition(v string) {
	o.Condition = &v
}

// GetExclusionType returns the ExclusionType field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAttributesAllOf) GetExclusionType() string {
	if o == nil || o.ExclusionType == nil {
		var ret string
		return ret
	}
	return *o.ExclusionType
}

// GetExclusionTypeOk returns a tuple with the ExclusionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAttributesAllOf) GetExclusionTypeOk() (*string, bool) {
	if o == nil || o.ExclusionType == nil {
		return nil, false
	}
	return o.ExclusionType, true
}

// HasExclusionType returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributesAllOf) HasExclusionType() bool {
	if o != nil && o.ExclusionType != nil {
		return true
	}

	return false
}

// SetExclusionType gets a reference to the given string and assigns it to the ExclusionType field.
func (o *WafExclusionResponseDataAttributesAllOf) SetExclusionType(v string) {
	o.ExclusionType = &v
}

// GetLogging returns the Logging field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAttributesAllOf) GetLogging() bool {
	if o == nil || o.Logging == nil {
		var ret bool
		return ret
	}
	return *o.Logging
}

// GetLoggingOk returns a tuple with the Logging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAttributesAllOf) GetLoggingOk() (*bool, bool) {
	if o == nil || o.Logging == nil {
		return nil, false
	}
	return o.Logging, true
}

// HasLogging returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributesAllOf) HasLogging() bool {
	if o != nil && o.Logging != nil {
		return true
	}

	return false
}

// SetLogging gets a reference to the given bool and assigns it to the Logging field.
func (o *WafExclusionResponseDataAttributesAllOf) SetLogging(v bool) {
	o.Logging = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAttributesAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAttributesAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributesAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafExclusionResponseDataAttributesAllOf) SetName(v string) {
	o.Name = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAttributesAllOf) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAttributesAllOf) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributesAllOf) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *WafExclusionResponseDataAttributesAllOf) SetNumber(v int32) {
	o.Number = &v
}

// GetVariable returns the Variable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafExclusionResponseDataAttributesAllOf) GetVariable() string {
	if o == nil || o.Variable.Get() == nil {
		var ret string
		return ret
	}
	return *o.Variable.Get()
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafExclusionResponseDataAttributesAllOf) GetVariableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variable.Get(), o.Variable.IsSet()
}

// HasVariable returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributesAllOf) HasVariable() bool {
	if o != nil && o.Variable.IsSet() {
		return true
	}

	return false
}

// SetVariable gets a reference to the given NullableString and assigns it to the Variable field.
func (o *WafExclusionResponseDataAttributesAllOf) SetVariable(v string) {
	o.Variable.Set(&v)
}

// SetVariableNil sets the value for Variable to be an explicit nil
func (o *WafExclusionResponseDataAttributesAllOf) SetVariableNil() {
	o.Variable.Set(nil)
}

// UnsetVariable ensures that no value is present for Variable, not even an explicit nil
func (o *WafExclusionResponseDataAttributesAllOf) UnsetVariable() {
	o.Variable.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafExclusionResponseDataAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.ExclusionType != nil {
		toSerialize["exclusion_type"] = o.ExclusionType
	}
	if o.Logging != nil {
		toSerialize["logging"] = o.Logging
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Variable.IsSet() {
		toSerialize["variable"] = o.Variable.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafExclusionResponseDataAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafExclusionResponseDataAttributesAllOf := _WafExclusionResponseDataAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varWafExclusionResponseDataAttributesAllOf); err == nil {
		*o = WafExclusionResponseDataAttributesAllOf(varWafExclusionResponseDataAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "condition")
		delete(additionalProperties, "exclusion_type")
		delete(additionalProperties, "logging")
		delete(additionalProperties, "name")
		delete(additionalProperties, "number")
		delete(additionalProperties, "variable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafExclusionResponseDataAttributesAllOf is a helper abstraction for handling nullable wafexclusionresponsedataattributesallof types.
type NullableWafExclusionResponseDataAttributesAllOf struct {
	value *WafExclusionResponseDataAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafExclusionResponseDataAttributesAllOf) Get() *WafExclusionResponseDataAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafExclusionResponseDataAttributesAllOf) Set(val *WafExclusionResponseDataAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafExclusionResponseDataAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafExclusionResponseDataAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafExclusionResponseDataAttributesAllOf returns a pointer to a new instance of NullableWafExclusionResponseDataAttributesAllOf.
func NewNullableWafExclusionResponseDataAttributesAllOf(val *WafExclusionResponseDataAttributesAllOf) *NullableWafExclusionResponseDataAttributesAllOf {
	return &NullableWafExclusionResponseDataAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafExclusionResponseDataAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafExclusionResponseDataAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
