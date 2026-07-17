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

// RuleResponse All attributes for a rule response.
type RuleResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// Alphanumeric string identifying the rule. Stable across versions of the routing config.
	Id *string `json:"id,omitempty"`
	// Whether this is the default (catch-all) rule for the path.
	IsDefault *bool   `json:"is_default,omitempty"`
	Action    *Action `json:"action,omitempty"`
	// The conditions a request must satisfy for this rule to match. Empty for the default rule.
	Conditions           []RoutingConfigCondition `json:"conditions,omitempty"`
	AdditionalProperties map[string]any
}

type _RuleResponse RuleResponse

// NewRuleResponse instantiates a new RuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleResponse() *RuleResponse {
	this := RuleResponse{}
	return &this
}

// NewRuleResponseWithDefaults instantiates a new RuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleResponseWithDefaults() *RuleResponse {
	this := RuleResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RuleResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *RuleResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *RuleResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *RuleResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RuleResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *RuleResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *RuleResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *RuleResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RuleResponse) SetId(v string) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *RuleResponse) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleResponse) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *RuleResponse) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *RuleResponse) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RuleResponse) GetAction() Action {
	if o == nil || o.Action == nil {
		var ret Action
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleResponse) GetActionOk() (*Action, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RuleResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given Action and assigns it to the Action field.
func (o *RuleResponse) SetAction(v Action) {
	o.Action = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *RuleResponse) GetConditions() []RoutingConfigCondition {
	if o == nil || o.Conditions == nil {
		var ret []RoutingConfigCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleResponse) GetConditionsOk() ([]RoutingConfigCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *RuleResponse) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []RoutingConfigCondition and assigns it to the Conditions field.
func (o *RuleResponse) SetConditions(v []RoutingConfigCondition) {
	o.Conditions = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRuleResponse := _RuleResponse{}

	if err = json.Unmarshal(bytes, &varRuleResponse); err == nil {
		*o = RuleResponse(varRuleResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_default")
		delete(additionalProperties, "action")
		delete(additionalProperties, "conditions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRuleResponse is a helper abstraction for handling nullable ruleresponse types.
type NullableRuleResponse struct {
	value *RuleResponse
	isSet bool
}

// Get returns the value.
func (v NullableRuleResponse) Get() *RuleResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableRuleResponse) Set(val *RuleResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRuleResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRuleResponse returns a pointer to a new instance of NullableRuleResponse.
func NewNullableRuleResponse(val *RuleResponse) *NullableRuleResponse {
	return &NullableRuleResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
