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

// LegacyWafRuleset struct for LegacyWafRuleset
type LegacyWafRuleset struct {
	// Date and time WAF ruleset VCL was last deployed.
	LastPush *string `json:"last_push,omitempty"`
	// The WAF ruleset VCL currently deployed.
	Vcl *string `json:"vcl,omitempty"`
	AdditionalProperties map[string]any
}

type _LegacyWafRuleset LegacyWafRuleset

// NewLegacyWafRuleset instantiates a new LegacyWafRuleset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyWafRuleset() *LegacyWafRuleset {
	this := LegacyWafRuleset{}
	return &this
}

// NewLegacyWafRulesetWithDefaults instantiates a new LegacyWafRuleset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyWafRulesetWithDefaults() *LegacyWafRuleset {
	this := LegacyWafRuleset{}
	return &this
}

// GetLastPush returns the LastPush field value if set, zero value otherwise.
func (o *LegacyWafRuleset) GetLastPush() string {
	if o == nil || o.LastPush == nil {
		var ret string
		return ret
	}
	return *o.LastPush
}

// GetLastPushOk returns a tuple with the LastPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafRuleset) GetLastPushOk() (*string, bool) {
	if o == nil || o.LastPush == nil {
		return nil, false
	}
	return o.LastPush, true
}

// HasLastPush returns a boolean if a field has been set.
func (o *LegacyWafRuleset) HasLastPush() bool {
	if o != nil && o.LastPush != nil {
		return true
	}

	return false
}

// SetLastPush gets a reference to the given string and assigns it to the LastPush field.
func (o *LegacyWafRuleset) SetLastPush(v string) {
	o.LastPush = &v
}

// GetVcl returns the Vcl field value if set, zero value otherwise.
func (o *LegacyWafRuleset) GetVcl() string {
	if o == nil || o.Vcl == nil {
		var ret string
		return ret
	}
	return *o.Vcl
}

// GetVclOk returns a tuple with the Vcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafRuleset) GetVclOk() (*string, bool) {
	if o == nil || o.Vcl == nil {
		return nil, false
	}
	return o.Vcl, true
}

// HasVcl returns a boolean if a field has been set.
func (o *LegacyWafRuleset) HasVcl() bool {
	if o != nil && o.Vcl != nil {
		return true
	}

	return false
}

// SetVcl gets a reference to the given string and assigns it to the Vcl field.
func (o *LegacyWafRuleset) SetVcl(v string) {
	o.Vcl = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LegacyWafRuleset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.LastPush != nil {
		toSerialize["last_push"] = o.LastPush
	}
	if o.Vcl != nil {
		toSerialize["vcl"] = o.Vcl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LegacyWafRuleset) UnmarshalJSON(bytes []byte) (err error) {
	varLegacyWafRuleset := _LegacyWafRuleset{}

	if err = json.Unmarshal(bytes, &varLegacyWafRuleset); err == nil {
		*o = LegacyWafRuleset(varLegacyWafRuleset)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "last_push")
		delete(additionalProperties, "vcl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLegacyWafRuleset is a helper abstraction for handling nullable legacywafruleset types. 
type NullableLegacyWafRuleset struct {
	value *LegacyWafRuleset
	isSet bool
}

// Get returns the value.
func (v NullableLegacyWafRuleset) Get() *LegacyWafRuleset {
	return v.value
}

// Set modifies the value.
func (v *NullableLegacyWafRuleset) Set(val *LegacyWafRuleset) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLegacyWafRuleset) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLegacyWafRuleset) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLegacyWafRuleset returns a pointer to a new instance of NullableLegacyWafRuleset.
func NewNullableLegacyWafRuleset(val *LegacyWafRuleset) *NullableLegacyWafRuleset {
	return &NullableLegacyWafRuleset{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLegacyWafRuleset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLegacyWafRuleset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
