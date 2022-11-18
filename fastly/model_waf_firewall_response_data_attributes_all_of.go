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

// WafFirewallResponseDataAttributesAllOf struct for WafFirewallResponseDataAttributesAllOf
type WafFirewallResponseDataAttributesAllOf struct {
	ServiceID *string `json:"service_id,omitempty"`
	// The number of active Fastly rules set to block on the active or latest firewall version.
	ActiveRulesFastlyBlockCount *int32 `json:"active_rules_fastly_block_count,omitempty"`
	// The number of active Fastly rules set to log on the active or latest firewall version.
	ActiveRulesFastlyLogCount *int32 `json:"active_rules_fastly_log_count,omitempty"`
	// The number of active Fastly rules set to score on the active or latest firewall version.
	ActiveRulesFastlyScoreCount *int32 `json:"active_rules_fastly_score_count,omitempty"`
	// The number of active OWASP rules set to block on the active or latest firewall version.
	ActiveRulesOwaspBlockCount *int32 `json:"active_rules_owasp_block_count,omitempty"`
	// The number of active OWASP rules set to log on the active or latest firewall version.
	ActiveRulesOwaspLogCount *int32 `json:"active_rules_owasp_log_count,omitempty"`
	// The number of active OWASP rules set to score on the active or latest firewall version.
	ActiveRulesOwaspScoreCount *int32 `json:"active_rules_owasp_score_count,omitempty"`
	// The number of active Trustwave rules set to block on the active or latest firewall version.
	ActiveRulesTrustwaveBlockCount *int32 `json:"active_rules_trustwave_block_count,omitempty"`
	// The number of active Trustwave rules set to log on the active or latest firewall version.
	ActiveRulesTrustwaveLogCount *int32 `json:"active_rules_trustwave_log_count,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallResponseDataAttributesAllOf WafFirewallResponseDataAttributesAllOf

// NewWafFirewallResponseDataAttributesAllOf instantiates a new WafFirewallResponseDataAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallResponseDataAttributesAllOf() *WafFirewallResponseDataAttributesAllOf {
	this := WafFirewallResponseDataAttributesAllOf{}
	return &this
}

// NewWafFirewallResponseDataAttributesAllOfWithDefaults instantiates a new WafFirewallResponseDataAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallResponseDataAttributesAllOfWithDefaults() *WafFirewallResponseDataAttributesAllOf {
	this := WafFirewallResponseDataAttributesAllOf{}
	return &this
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAttributesAllOf) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAttributesAllOf) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAttributesAllOf) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *WafFirewallResponseDataAttributesAllOf) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetActiveRulesFastlyBlockCount returns the ActiveRulesFastlyBlockCount field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyBlockCount() int32 {
	if o == nil || o.ActiveRulesFastlyBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesFastlyBlockCount
}

// GetActiveRulesFastlyBlockCountOk returns a tuple with the ActiveRulesFastlyBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyBlockCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesFastlyBlockCount == nil {
		return nil, false
	}
	return o.ActiveRulesFastlyBlockCount, true
}

// HasActiveRulesFastlyBlockCount returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesFastlyBlockCount() bool {
	if o != nil && o.ActiveRulesFastlyBlockCount != nil {
		return true
	}

	return false
}

// SetActiveRulesFastlyBlockCount gets a reference to the given int32 and assigns it to the ActiveRulesFastlyBlockCount field.
func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesFastlyBlockCount(v int32) {
	o.ActiveRulesFastlyBlockCount = &v
}

// GetActiveRulesFastlyLogCount returns the ActiveRulesFastlyLogCount field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyLogCount() int32 {
	if o == nil || o.ActiveRulesFastlyLogCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesFastlyLogCount
}

// GetActiveRulesFastlyLogCountOk returns a tuple with the ActiveRulesFastlyLogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyLogCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesFastlyLogCount == nil {
		return nil, false
	}
	return o.ActiveRulesFastlyLogCount, true
}

// HasActiveRulesFastlyLogCount returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesFastlyLogCount() bool {
	if o != nil && o.ActiveRulesFastlyLogCount != nil {
		return true
	}

	return false
}

// SetActiveRulesFastlyLogCount gets a reference to the given int32 and assigns it to the ActiveRulesFastlyLogCount field.
func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesFastlyLogCount(v int32) {
	o.ActiveRulesFastlyLogCount = &v
}

// GetActiveRulesFastlyScoreCount returns the ActiveRulesFastlyScoreCount field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyScoreCount() int32 {
	if o == nil || o.ActiveRulesFastlyScoreCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesFastlyScoreCount
}

// GetActiveRulesFastlyScoreCountOk returns a tuple with the ActiveRulesFastlyScoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesFastlyScoreCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesFastlyScoreCount == nil {
		return nil, false
	}
	return o.ActiveRulesFastlyScoreCount, true
}

// HasActiveRulesFastlyScoreCount returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesFastlyScoreCount() bool {
	if o != nil && o.ActiveRulesFastlyScoreCount != nil {
		return true
	}

	return false
}

// SetActiveRulesFastlyScoreCount gets a reference to the given int32 and assigns it to the ActiveRulesFastlyScoreCount field.
func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesFastlyScoreCount(v int32) {
	o.ActiveRulesFastlyScoreCount = &v
}

// GetActiveRulesOwaspBlockCount returns the ActiveRulesOwaspBlockCount field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspBlockCount() int32 {
	if o == nil || o.ActiveRulesOwaspBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesOwaspBlockCount
}

// GetActiveRulesOwaspBlockCountOk returns a tuple with the ActiveRulesOwaspBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspBlockCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesOwaspBlockCount == nil {
		return nil, false
	}
	return o.ActiveRulesOwaspBlockCount, true
}

// HasActiveRulesOwaspBlockCount returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesOwaspBlockCount() bool {
	if o != nil && o.ActiveRulesOwaspBlockCount != nil {
		return true
	}

	return false
}

// SetActiveRulesOwaspBlockCount gets a reference to the given int32 and assigns it to the ActiveRulesOwaspBlockCount field.
func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesOwaspBlockCount(v int32) {
	o.ActiveRulesOwaspBlockCount = &v
}

// GetActiveRulesOwaspLogCount returns the ActiveRulesOwaspLogCount field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspLogCount() int32 {
	if o == nil || o.ActiveRulesOwaspLogCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesOwaspLogCount
}

// GetActiveRulesOwaspLogCountOk returns a tuple with the ActiveRulesOwaspLogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspLogCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesOwaspLogCount == nil {
		return nil, false
	}
	return o.ActiveRulesOwaspLogCount, true
}

// HasActiveRulesOwaspLogCount returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesOwaspLogCount() bool {
	if o != nil && o.ActiveRulesOwaspLogCount != nil {
		return true
	}

	return false
}

// SetActiveRulesOwaspLogCount gets a reference to the given int32 and assigns it to the ActiveRulesOwaspLogCount field.
func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesOwaspLogCount(v int32) {
	o.ActiveRulesOwaspLogCount = &v
}

// GetActiveRulesOwaspScoreCount returns the ActiveRulesOwaspScoreCount field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspScoreCount() int32 {
	if o == nil || o.ActiveRulesOwaspScoreCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesOwaspScoreCount
}

// GetActiveRulesOwaspScoreCountOk returns a tuple with the ActiveRulesOwaspScoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesOwaspScoreCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesOwaspScoreCount == nil {
		return nil, false
	}
	return o.ActiveRulesOwaspScoreCount, true
}

// HasActiveRulesOwaspScoreCount returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesOwaspScoreCount() bool {
	if o != nil && o.ActiveRulesOwaspScoreCount != nil {
		return true
	}

	return false
}

// SetActiveRulesOwaspScoreCount gets a reference to the given int32 and assigns it to the ActiveRulesOwaspScoreCount field.
func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesOwaspScoreCount(v int32) {
	o.ActiveRulesOwaspScoreCount = &v
}

// GetActiveRulesTrustwaveBlockCount returns the ActiveRulesTrustwaveBlockCount field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesTrustwaveBlockCount() int32 {
	if o == nil || o.ActiveRulesTrustwaveBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesTrustwaveBlockCount
}

// GetActiveRulesTrustwaveBlockCountOk returns a tuple with the ActiveRulesTrustwaveBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesTrustwaveBlockCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesTrustwaveBlockCount == nil {
		return nil, false
	}
	return o.ActiveRulesTrustwaveBlockCount, true
}

// HasActiveRulesTrustwaveBlockCount returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesTrustwaveBlockCount() bool {
	if o != nil && o.ActiveRulesTrustwaveBlockCount != nil {
		return true
	}

	return false
}

// SetActiveRulesTrustwaveBlockCount gets a reference to the given int32 and assigns it to the ActiveRulesTrustwaveBlockCount field.
func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesTrustwaveBlockCount(v int32) {
	o.ActiveRulesTrustwaveBlockCount = &v
}

// GetActiveRulesTrustwaveLogCount returns the ActiveRulesTrustwaveLogCount field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesTrustwaveLogCount() int32 {
	if o == nil || o.ActiveRulesTrustwaveLogCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesTrustwaveLogCount
}

// GetActiveRulesTrustwaveLogCountOk returns a tuple with the ActiveRulesTrustwaveLogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAttributesAllOf) GetActiveRulesTrustwaveLogCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesTrustwaveLogCount == nil {
		return nil, false
	}
	return o.ActiveRulesTrustwaveLogCount, true
}

// HasActiveRulesTrustwaveLogCount returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAttributesAllOf) HasActiveRulesTrustwaveLogCount() bool {
	if o != nil && o.ActiveRulesTrustwaveLogCount != nil {
		return true
	}

	return false
}

// SetActiveRulesTrustwaveLogCount gets a reference to the given int32 and assigns it to the ActiveRulesTrustwaveLogCount field.
func (o *WafFirewallResponseDataAttributesAllOf) SetActiveRulesTrustwaveLogCount(v int32) {
	o.ActiveRulesTrustwaveLogCount = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallResponseDataAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.ActiveRulesFastlyBlockCount != nil {
		toSerialize["active_rules_fastly_block_count"] = o.ActiveRulesFastlyBlockCount
	}
	if o.ActiveRulesFastlyLogCount != nil {
		toSerialize["active_rules_fastly_log_count"] = o.ActiveRulesFastlyLogCount
	}
	if o.ActiveRulesFastlyScoreCount != nil {
		toSerialize["active_rules_fastly_score_count"] = o.ActiveRulesFastlyScoreCount
	}
	if o.ActiveRulesOwaspBlockCount != nil {
		toSerialize["active_rules_owasp_block_count"] = o.ActiveRulesOwaspBlockCount
	}
	if o.ActiveRulesOwaspLogCount != nil {
		toSerialize["active_rules_owasp_log_count"] = o.ActiveRulesOwaspLogCount
	}
	if o.ActiveRulesOwaspScoreCount != nil {
		toSerialize["active_rules_owasp_score_count"] = o.ActiveRulesOwaspScoreCount
	}
	if o.ActiveRulesTrustwaveBlockCount != nil {
		toSerialize["active_rules_trustwave_block_count"] = o.ActiveRulesTrustwaveBlockCount
	}
	if o.ActiveRulesTrustwaveLogCount != nil {
		toSerialize["active_rules_trustwave_log_count"] = o.ActiveRulesTrustwaveLogCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WafFirewallResponseDataAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallResponseDataAttributesAllOf := _WafFirewallResponseDataAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varWafFirewallResponseDataAttributesAllOf); err == nil {
		*o = WafFirewallResponseDataAttributesAllOf(varWafFirewallResponseDataAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "active_rules_fastly_block_count")
		delete(additionalProperties, "active_rules_fastly_log_count")
		delete(additionalProperties, "active_rules_fastly_score_count")
		delete(additionalProperties, "active_rules_owasp_block_count")
		delete(additionalProperties, "active_rules_owasp_log_count")
		delete(additionalProperties, "active_rules_owasp_score_count")
		delete(additionalProperties, "active_rules_trustwave_block_count")
		delete(additionalProperties, "active_rules_trustwave_log_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallResponseDataAttributesAllOf is a helper abstraction for handling nullable waffirewallresponsedataattributesallof types. 
type NullableWafFirewallResponseDataAttributesAllOf struct {
	value *WafFirewallResponseDataAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallResponseDataAttributesAllOf) Get() *WafFirewallResponseDataAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallResponseDataAttributesAllOf) Set(val *WafFirewallResponseDataAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallResponseDataAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallResponseDataAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallResponseDataAttributesAllOf returns a pointer to a new instance of NullableWafFirewallResponseDataAttributesAllOf.
func NewNullableWafFirewallResponseDataAttributesAllOf(val *WafFirewallResponseDataAttributesAllOf) *NullableWafFirewallResponseDataAttributesAllOf {
	return &NullableWafFirewallResponseDataAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallResponseDataAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallResponseDataAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
