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

// LegacyWafRule struct for LegacyWafRule
type LegacyWafRule struct {
	// Message metadata for the rule.
	Message *string `json:"message,omitempty"`
	// Corresponding ModSecurity rule ID.
	RuleID *string `json:"rule_id,omitempty"`
	// Severity metadata for the rule.
	Severity *int32 `json:"severity,omitempty"`
	// The ModSecurity rule logic.
	Source *string `json:"source,omitempty"`
	// The VCL representation of the rule logic.
	Vcl                  *string `json:"vcl,omitempty"`
	AdditionalProperties map[string]any
}

type _LegacyWafRule LegacyWafRule

// NewLegacyWafRule instantiates a new LegacyWafRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyWafRule() *LegacyWafRule {
	this := LegacyWafRule{}
	return &this
}

// NewLegacyWafRuleWithDefaults instantiates a new LegacyWafRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyWafRuleWithDefaults() *LegacyWafRule {
	this := LegacyWafRule{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LegacyWafRule) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafRule) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LegacyWafRule) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LegacyWafRule) SetMessage(v string) {
	o.Message = &v
}

// GetRuleID returns the RuleID field value if set, zero value otherwise.
func (o *LegacyWafRule) GetRuleID() string {
	if o == nil || o.RuleID == nil {
		var ret string
		return ret
	}
	return *o.RuleID
}

// GetRuleIDOk returns a tuple with the RuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafRule) GetRuleIDOk() (*string, bool) {
	if o == nil || o.RuleID == nil {
		return nil, false
	}
	return o.RuleID, true
}

// HasRuleID returns a boolean if a field has been set.
func (o *LegacyWafRule) HasRuleID() bool {
	if o != nil && o.RuleID != nil {
		return true
	}

	return false
}

// SetRuleID gets a reference to the given string and assigns it to the RuleID field.
func (o *LegacyWafRule) SetRuleID(v string) {
	o.RuleID = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *LegacyWafRule) GetSeverity() int32 {
	if o == nil || o.Severity == nil {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafRule) GetSeverityOk() (*int32, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *LegacyWafRule) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *LegacyWafRule) SetSeverity(v int32) {
	o.Severity = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *LegacyWafRule) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafRule) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *LegacyWafRule) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *LegacyWafRule) SetSource(v string) {
	o.Source = &v
}

// GetVcl returns the Vcl field value if set, zero value otherwise.
func (o *LegacyWafRule) GetVcl() string {
	if o == nil || o.Vcl == nil {
		var ret string
		return ret
	}
	return *o.Vcl
}

// GetVclOk returns a tuple with the Vcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafRule) GetVclOk() (*string, bool) {
	if o == nil || o.Vcl == nil {
		return nil, false
	}
	return o.Vcl, true
}

// HasVcl returns a boolean if a field has been set.
func (o *LegacyWafRule) HasVcl() bool {
	if o != nil && o.Vcl != nil {
		return true
	}

	return false
}

// SetVcl gets a reference to the given string and assigns it to the Vcl field.
func (o *LegacyWafRule) SetVcl(v string) {
	o.Vcl = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LegacyWafRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.RuleID != nil {
		toSerialize["rule_id"] = o.RuleID
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
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
func (o *LegacyWafRule) UnmarshalJSON(bytes []byte) (err error) {
	varLegacyWafRule := _LegacyWafRule{}

	if err = json.Unmarshal(bytes, &varLegacyWafRule); err == nil {
		*o = LegacyWafRule(varLegacyWafRule)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "rule_id")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "source")
		delete(additionalProperties, "vcl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLegacyWafRule is a helper abstraction for handling nullable legacywafrule types.
type NullableLegacyWafRule struct {
	value *LegacyWafRule
	isSet bool
}

// Get returns the value.
func (v NullableLegacyWafRule) Get() *LegacyWafRule {
	return v.value
}

// Set modifies the value.
func (v *NullableLegacyWafRule) Set(val *LegacyWafRule) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLegacyWafRule) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLegacyWafRule) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLegacyWafRule returns a pointer to a new instance of NullableLegacyWafRule.
func NewNullableLegacyWafRule(val *LegacyWafRule) *NullableLegacyWafRule {
	return &NullableLegacyWafRule{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLegacyWafRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLegacyWafRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
