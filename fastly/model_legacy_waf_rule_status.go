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

// LegacyWafRuleStatus struct for LegacyWafRuleStatus
type LegacyWafRuleStatus struct {
	// Describes the behavior for the particular rule within this firewall. Permitted values: `log`, `block`, and `disabled`.
	Status *string `json:"status,omitempty"`
	// The ModSecurity rule ID.
	ModsecRuleID *string `json:"modsec_rule_id,omitempty"`
	// The Rule ID.
	UniqueRuleID         *string `json:"unique_rule_id,omitempty"`
	AdditionalProperties map[string]any
}

type _LegacyWafRuleStatus LegacyWafRuleStatus

// NewLegacyWafRuleStatus instantiates a new LegacyWafRuleStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyWafRuleStatus() *LegacyWafRuleStatus {
	this := LegacyWafRuleStatus{}
	return &this
}

// NewLegacyWafRuleStatusWithDefaults instantiates a new LegacyWafRuleStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyWafRuleStatusWithDefaults() *LegacyWafRuleStatus {
	this := LegacyWafRuleStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LegacyWafRuleStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafRuleStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LegacyWafRuleStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LegacyWafRuleStatus) SetStatus(v string) {
	o.Status = &v
}

// GetModsecRuleID returns the ModsecRuleID field value if set, zero value otherwise.
func (o *LegacyWafRuleStatus) GetModsecRuleID() string {
	if o == nil || o.ModsecRuleID == nil {
		var ret string
		return ret
	}
	return *o.ModsecRuleID
}

// GetModsecRuleIDOk returns a tuple with the ModsecRuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafRuleStatus) GetModsecRuleIDOk() (*string, bool) {
	if o == nil || o.ModsecRuleID == nil {
		return nil, false
	}
	return o.ModsecRuleID, true
}

// HasModsecRuleID returns a boolean if a field has been set.
func (o *LegacyWafRuleStatus) HasModsecRuleID() bool {
	if o != nil && o.ModsecRuleID != nil {
		return true
	}

	return false
}

// SetModsecRuleID gets a reference to the given string and assigns it to the ModsecRuleID field.
func (o *LegacyWafRuleStatus) SetModsecRuleID(v string) {
	o.ModsecRuleID = &v
}

// GetUniqueRuleID returns the UniqueRuleID field value if set, zero value otherwise.
func (o *LegacyWafRuleStatus) GetUniqueRuleID() string {
	if o == nil || o.UniqueRuleID == nil {
		var ret string
		return ret
	}
	return *o.UniqueRuleID
}

// GetUniqueRuleIDOk returns a tuple with the UniqueRuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafRuleStatus) GetUniqueRuleIDOk() (*string, bool) {
	if o == nil || o.UniqueRuleID == nil {
		return nil, false
	}
	return o.UniqueRuleID, true
}

// HasUniqueRuleID returns a boolean if a field has been set.
func (o *LegacyWafRuleStatus) HasUniqueRuleID() bool {
	if o != nil && o.UniqueRuleID != nil {
		return true
	}

	return false
}

// SetUniqueRuleID gets a reference to the given string and assigns it to the UniqueRuleID field.
func (o *LegacyWafRuleStatus) SetUniqueRuleID(v string) {
	o.UniqueRuleID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LegacyWafRuleStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ModsecRuleID != nil {
		toSerialize["modsec_rule_id"] = o.ModsecRuleID
	}
	if o.UniqueRuleID != nil {
		toSerialize["unique_rule_id"] = o.UniqueRuleID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LegacyWafRuleStatus) UnmarshalJSON(bytes []byte) (err error) {
	varLegacyWafRuleStatus := _LegacyWafRuleStatus{}

	if err = json.Unmarshal(bytes, &varLegacyWafRuleStatus); err == nil {
		*o = LegacyWafRuleStatus(varLegacyWafRuleStatus)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "modsec_rule_id")
		delete(additionalProperties, "unique_rule_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLegacyWafRuleStatus is a helper abstraction for handling nullable legacywafrulestatus types.
type NullableLegacyWafRuleStatus struct {
	value *LegacyWafRuleStatus
	isSet bool
}

// Get returns the value.
func (v NullableLegacyWafRuleStatus) Get() *LegacyWafRuleStatus {
	return v.value
}

// Set modifies the value.
func (v *NullableLegacyWafRuleStatus) Set(val *LegacyWafRuleStatus) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLegacyWafRuleStatus) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLegacyWafRuleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLegacyWafRuleStatus returns a pointer to a new instance of NullableLegacyWafRuleStatus.
func NewNullableLegacyWafRuleStatus(val *LegacyWafRuleStatus) *NullableLegacyWafRuleStatus {
	return &NullableLegacyWafRuleStatus{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLegacyWafRuleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLegacyWafRuleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
