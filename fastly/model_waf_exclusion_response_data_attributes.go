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
	"time"
)

// WafExclusionResponseDataAttributes struct for WafExclusionResponseDataAttributes
type WafExclusionResponseDataAttributes struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
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
	Variable NullableString `json:"variable,omitempty"`
	AdditionalProperties map[string]any
}

type _WafExclusionResponseDataAttributes WafExclusionResponseDataAttributes

// NewWafExclusionResponseDataAttributes instantiates a new WafExclusionResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafExclusionResponseDataAttributes() *WafExclusionResponseDataAttributes {
	this := WafExclusionResponseDataAttributes{}
	var logging bool = true
	this.Logging = &logging
	return &this
}

// NewWafExclusionResponseDataAttributesWithDefaults instantiates a new WafExclusionResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafExclusionResponseDataAttributesWithDefaults() *WafExclusionResponseDataAttributes {
	this := WafExclusionResponseDataAttributes{}
	var logging bool = true
	this.Logging = &logging
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafExclusionResponseDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafExclusionResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *WafExclusionResponseDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *WafExclusionResponseDataAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *WafExclusionResponseDataAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafExclusionResponseDataAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafExclusionResponseDataAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributes) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *WafExclusionResponseDataAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *WafExclusionResponseDataAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *WafExclusionResponseDataAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafExclusionResponseDataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafExclusionResponseDataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *WafExclusionResponseDataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *WafExclusionResponseDataAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *WafExclusionResponseDataAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAttributes) GetCondition() string {
	if o == nil || o.Condition == nil {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAttributes) GetConditionOk() (*string, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributes) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *WafExclusionResponseDataAttributes) SetCondition(v string) {
	o.Condition = &v
}

// GetExclusionType returns the ExclusionType field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAttributes) GetExclusionType() string {
	if o == nil || o.ExclusionType == nil {
		var ret string
		return ret
	}
	return *o.ExclusionType
}

// GetExclusionTypeOk returns a tuple with the ExclusionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAttributes) GetExclusionTypeOk() (*string, bool) {
	if o == nil || o.ExclusionType == nil {
		return nil, false
	}
	return o.ExclusionType, true
}

// HasExclusionType returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributes) HasExclusionType() bool {
	if o != nil && o.ExclusionType != nil {
		return true
	}

	return false
}

// SetExclusionType gets a reference to the given string and assigns it to the ExclusionType field.
func (o *WafExclusionResponseDataAttributes) SetExclusionType(v string) {
	o.ExclusionType = &v
}

// GetLogging returns the Logging field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAttributes) GetLogging() bool {
	if o == nil || o.Logging == nil {
		var ret bool
		return ret
	}
	return *o.Logging
}

// GetLoggingOk returns a tuple with the Logging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAttributes) GetLoggingOk() (*bool, bool) {
	if o == nil || o.Logging == nil {
		return nil, false
	}
	return o.Logging, true
}

// HasLogging returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributes) HasLogging() bool {
	if o != nil && o.Logging != nil {
		return true
	}

	return false
}

// SetLogging gets a reference to the given bool and assigns it to the Logging field.
func (o *WafExclusionResponseDataAttributes) SetLogging(v bool) {
	o.Logging = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafExclusionResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAttributes) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAttributes) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributes) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *WafExclusionResponseDataAttributes) SetNumber(v int32) {
	o.Number = &v
}

// GetVariable returns the Variable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafExclusionResponseDataAttributes) GetVariable() string {
	if o == nil || o.Variable.Get() == nil {
		var ret string
		return ret
	}
	return *o.Variable.Get()
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafExclusionResponseDataAttributes) GetVariableOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Variable.Get(), o.Variable.IsSet()
}

// HasVariable returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAttributes) HasVariable() bool {
	if o != nil && o.Variable.IsSet() {
		return true
	}

	return false
}

// SetVariable gets a reference to the given NullableString and assigns it to the Variable field.
func (o *WafExclusionResponseDataAttributes) SetVariable(v string) {
	o.Variable.Set(&v)
}
// SetVariableNil sets the value for Variable to be an explicit nil
func (o *WafExclusionResponseDataAttributes) SetVariableNil() {
	o.Variable.Set(nil)
}

// UnsetVariable ensures that no value is present for Variable, not even an explicit nil
func (o *WafExclusionResponseDataAttributes) UnsetVariable() {
	o.Variable.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafExclusionResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
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
func (o *WafExclusionResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varWafExclusionResponseDataAttributes := _WafExclusionResponseDataAttributes{}

	if err = json.Unmarshal(bytes, &varWafExclusionResponseDataAttributes); err == nil {
		*o = WafExclusionResponseDataAttributes(varWafExclusionResponseDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
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

// NullableWafExclusionResponseDataAttributes is a helper abstraction for handling nullable wafexclusionresponsedataattributes types. 
type NullableWafExclusionResponseDataAttributes struct {
	value *WafExclusionResponseDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableWafExclusionResponseDataAttributes) Get() *WafExclusionResponseDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableWafExclusionResponseDataAttributes) Set(val *WafExclusionResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafExclusionResponseDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafExclusionResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafExclusionResponseDataAttributes returns a pointer to a new instance of NullableWafExclusionResponseDataAttributes.
func NewNullableWafExclusionResponseDataAttributes(val *WafExclusionResponseDataAttributes) *NullableWafExclusionResponseDataAttributes {
	return &NullableWafExclusionResponseDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafExclusionResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafExclusionResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
