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

// RoutingConfigCondition A condition that must be met for a rule to match.
type RoutingConfigCondition struct {
	Type     ConditionType     `json:"type"`
	Operator ConditionOperator `json:"operator"`
	// The key to evaluate. For `header` conditions this is the header name. Required for `header` conditions.
	Key *string `json:"key,omitempty"`
	// The value to compare against using the operator.
	Value                string `json:"value"`
	AdditionalProperties map[string]any
}

type _RoutingConfigCondition RoutingConfigCondition

// NewRoutingConfigCondition instantiates a new RoutingConfigCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingConfigCondition(type_ ConditionType, operator ConditionOperator, value string) *RoutingConfigCondition {
	this := RoutingConfigCondition{}
	this.Type = type_
	this.Operator = operator
	this.Value = value
	return &this
}

// NewRoutingConfigConditionWithDefaults instantiates a new RoutingConfigCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingConfigConditionWithDefaults() *RoutingConfigCondition {
	this := RoutingConfigCondition{}
	return &this
}

// GetType returns the Type field value
func (o *RoutingConfigCondition) GetType() ConditionType {
	if o == nil {
		var ret ConditionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingConfigCondition) GetTypeOk() (*ConditionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoutingConfigCondition) SetType(v ConditionType) {
	o.Type = v
}

// GetOperator returns the Operator field value
func (o *RoutingConfigCondition) GetOperator() ConditionOperator {
	if o == nil {
		var ret ConditionOperator
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *RoutingConfigCondition) GetOperatorOk() (*ConditionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *RoutingConfigCondition) SetOperator(v ConditionOperator) {
	o.Operator = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RoutingConfigCondition) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfigCondition) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RoutingConfigCondition) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RoutingConfigCondition) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value
func (o *RoutingConfigCondition) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RoutingConfigCondition) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RoutingConfigCondition) SetValue(v string) {
	o.Value = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RoutingConfigCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RoutingConfigCondition) UnmarshalJSON(bytes []byte) (err error) {
	varRoutingConfigCondition := _RoutingConfigCondition{}

	if err = json.Unmarshal(bytes, &varRoutingConfigCondition); err == nil {
		*o = RoutingConfigCondition(varRoutingConfigCondition)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRoutingConfigCondition is a helper abstraction for handling nullable routingconfigcondition types.
type NullableRoutingConfigCondition struct {
	value *RoutingConfigCondition
	isSet bool
}

// Get returns the value.
func (v NullableRoutingConfigCondition) Get() *RoutingConfigCondition {
	return v.value
}

// Set modifies the value.
func (v *NullableRoutingConfigCondition) Set(val *RoutingConfigCondition) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRoutingConfigCondition) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRoutingConfigCondition) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRoutingConfigCondition returns a pointer to a new instance of NullableRoutingConfigCondition.
func NewNullableRoutingConfigCondition(val *RoutingConfigCondition) *NullableRoutingConfigCondition {
	return &NullableRoutingConfigCondition{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRoutingConfigCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRoutingConfigCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
