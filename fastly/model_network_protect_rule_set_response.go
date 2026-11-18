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

// NetworkProtectRuleSetResponse A Network Protect rule-set response.
type NetworkProtectRuleSetResponse struct {
	DefaultAction NetworkProtectRuleSetDefaultAction `json:"default_action"`
	// Rules in evaluation order with read-only timestamps.
	Rules []NetworkProtectRuleResponse `json:"rules"`
	// A read-only unique identifier for a Network Protect resource.
	ID string `json:"id"`
	// An ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at"`
	// An ISO 8601 timestamp in UTC.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleSetResponse NetworkProtectRuleSetResponse

// NewNetworkProtectRuleSetResponse instantiates a new NetworkProtectRuleSetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleSetResponse(defaultAction NetworkProtectRuleSetDefaultAction, rules []NetworkProtectRuleResponse, id string, createdAt time.Time, updatedAt time.Time) *NetworkProtectRuleSetResponse {
	this := NetworkProtectRuleSetResponse{}
	this.DefaultAction = defaultAction
	this.Rules = rules
	this.ID = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewNetworkProtectRuleSetResponseWithDefaults instantiates a new NetworkProtectRuleSetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleSetResponseWithDefaults() *NetworkProtectRuleSetResponse {
	this := NetworkProtectRuleSetResponse{}
	return &this
}

// GetDefaultAction returns the DefaultAction field value
func (o *NetworkProtectRuleSetResponse) GetDefaultAction() NetworkProtectRuleSetDefaultAction {
	if o == nil {
		var ret NetworkProtectRuleSetDefaultAction
		return ret
	}

	return o.DefaultAction
}

// GetDefaultActionOk returns a tuple with the DefaultAction field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetResponse) GetDefaultActionOk() (*NetworkProtectRuleSetDefaultAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultAction, true
}

// SetDefaultAction sets field value
func (o *NetworkProtectRuleSetResponse) SetDefaultAction(v NetworkProtectRuleSetDefaultAction) {
	o.DefaultAction = v
}

// GetRules returns the Rules field value
func (o *NetworkProtectRuleSetResponse) GetRules() []NetworkProtectRuleResponse {
	if o == nil {
		var ret []NetworkProtectRuleResponse
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetResponse) GetRulesOk() ([]NetworkProtectRuleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *NetworkProtectRuleSetResponse) SetRules(v []NetworkProtectRuleResponse) {
	o.Rules = v
}

// GetID returns the ID field value
func (o *NetworkProtectRuleSetResponse) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetResponse) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *NetworkProtectRuleSetResponse) SetID(v string) {
	o.ID = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NetworkProtectRuleSetResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NetworkProtectRuleSetResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NetworkProtectRuleSetResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NetworkProtectRuleSetResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleSetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["default_action"] = o.DefaultAction
	}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if true {
		toSerialize["id"] = o.ID
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleSetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleSetResponse := _NetworkProtectRuleSetResponse{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleSetResponse); err == nil {
		*o = NetworkProtectRuleSetResponse(varNetworkProtectRuleSetResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "default_action")
		delete(additionalProperties, "rules")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleSetResponse is a helper abstraction for handling nullable networkprotectrulesetresponse types.
type NullableNetworkProtectRuleSetResponse struct {
	value *NetworkProtectRuleSetResponse
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleSetResponse) Get() *NetworkProtectRuleSetResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleSetResponse) Set(val *NetworkProtectRuleSetResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleSetResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleSetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleSetResponse returns a pointer to a new instance of NullableNetworkProtectRuleSetResponse.
func NewNullableNetworkProtectRuleSetResponse(val *NetworkProtectRuleSetResponse) *NullableNetworkProtectRuleSetResponse {
	return &NullableNetworkProtectRuleSetResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleSetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleSetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
