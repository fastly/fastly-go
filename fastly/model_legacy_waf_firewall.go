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

// LegacyWafFirewall struct for LegacyWafFirewall
type LegacyWafFirewall struct {
	// Date and time that VCL was last pushed to cache nodes.
	LastPush *string `json:"last_push,omitempty"`
	// Name of the corresponding condition object.
	PrefetchCondition *string `json:"prefetch_condition,omitempty"`
	// Name of the corresponding response object.
	Response  *string            `json:"response,omitempty"`
	Version   *ReadOnlyVersion   `json:"version,omitempty"`
	ServiceID *ReadOnlyServiceID `json:"service_id,omitempty"`
	// The status of the firewall.
	Disabled *bool `json:"disabled,omitempty"`
	// The number of rule statuses set to log.
	RuleStatusesLogCount *int32 `json:"rule_statuses_log_count,omitempty"`
	// The number of rule statuses set to block.
	RuleStatusesBlockCount *int32 `json:"rule_statuses_block_count,omitempty"`
	// The number of rule statuses set to disabled.
	RuleStatusesDisabledCount *int32 `json:"rule_statuses_disabled_count,omitempty"`
	AdditionalProperties      map[string]any
}

type _LegacyWafFirewall LegacyWafFirewall

// NewLegacyWafFirewall instantiates a new LegacyWafFirewall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyWafFirewall() *LegacyWafFirewall {
	this := LegacyWafFirewall{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// NewLegacyWafFirewallWithDefaults instantiates a new LegacyWafFirewall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyWafFirewallWithDefaults() *LegacyWafFirewall {
	this := LegacyWafFirewall{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetLastPush returns the LastPush field value if set, zero value otherwise.
func (o *LegacyWafFirewall) GetLastPush() string {
	if o == nil || o.LastPush == nil {
		var ret string
		return ret
	}
	return *o.LastPush
}

// GetLastPushOk returns a tuple with the LastPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafFirewall) GetLastPushOk() (*string, bool) {
	if o == nil || o.LastPush == nil {
		return nil, false
	}
	return o.LastPush, true
}

// HasLastPush returns a boolean if a field has been set.
func (o *LegacyWafFirewall) HasLastPush() bool {
	if o != nil && o.LastPush != nil {
		return true
	}

	return false
}

// SetLastPush gets a reference to the given string and assigns it to the LastPush field.
func (o *LegacyWafFirewall) SetLastPush(v string) {
	o.LastPush = &v
}

// GetPrefetchCondition returns the PrefetchCondition field value if set, zero value otherwise.
func (o *LegacyWafFirewall) GetPrefetchCondition() string {
	if o == nil || o.PrefetchCondition == nil {
		var ret string
		return ret
	}
	return *o.PrefetchCondition
}

// GetPrefetchConditionOk returns a tuple with the PrefetchCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafFirewall) GetPrefetchConditionOk() (*string, bool) {
	if o == nil || o.PrefetchCondition == nil {
		return nil, false
	}
	return o.PrefetchCondition, true
}

// HasPrefetchCondition returns a boolean if a field has been set.
func (o *LegacyWafFirewall) HasPrefetchCondition() bool {
	if o != nil && o.PrefetchCondition != nil {
		return true
	}

	return false
}

// SetPrefetchCondition gets a reference to the given string and assigns it to the PrefetchCondition field.
func (o *LegacyWafFirewall) SetPrefetchCondition(v string) {
	o.PrefetchCondition = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *LegacyWafFirewall) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafFirewall) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *LegacyWafFirewall) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *LegacyWafFirewall) SetResponse(v string) {
	o.Response = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LegacyWafFirewall) GetVersion() ReadOnlyVersion {
	if o == nil || o.Version == nil {
		var ret ReadOnlyVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafFirewall) GetVersionOk() (*ReadOnlyVersion, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LegacyWafFirewall) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given ReadOnlyVersion and assigns it to the Version field.
func (o *LegacyWafFirewall) SetVersion(v ReadOnlyVersion) {
	o.Version = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *LegacyWafFirewall) GetServiceID() ReadOnlyServiceID {
	if o == nil || o.ServiceID == nil {
		var ret ReadOnlyServiceID
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafFirewall) GetServiceIDOk() (*ReadOnlyServiceID, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *LegacyWafFirewall) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given ReadOnlyServiceID and assigns it to the ServiceID field.
func (o *LegacyWafFirewall) SetServiceID(v ReadOnlyServiceID) {
	o.ServiceID = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *LegacyWafFirewall) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafFirewall) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *LegacyWafFirewall) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *LegacyWafFirewall) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetRuleStatusesLogCount returns the RuleStatusesLogCount field value if set, zero value otherwise.
func (o *LegacyWafFirewall) GetRuleStatusesLogCount() int32 {
	if o == nil || o.RuleStatusesLogCount == nil {
		var ret int32
		return ret
	}
	return *o.RuleStatusesLogCount
}

// GetRuleStatusesLogCountOk returns a tuple with the RuleStatusesLogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafFirewall) GetRuleStatusesLogCountOk() (*int32, bool) {
	if o == nil || o.RuleStatusesLogCount == nil {
		return nil, false
	}
	return o.RuleStatusesLogCount, true
}

// HasRuleStatusesLogCount returns a boolean if a field has been set.
func (o *LegacyWafFirewall) HasRuleStatusesLogCount() bool {
	if o != nil && o.RuleStatusesLogCount != nil {
		return true
	}

	return false
}

// SetRuleStatusesLogCount gets a reference to the given int32 and assigns it to the RuleStatusesLogCount field.
func (o *LegacyWafFirewall) SetRuleStatusesLogCount(v int32) {
	o.RuleStatusesLogCount = &v
}

// GetRuleStatusesBlockCount returns the RuleStatusesBlockCount field value if set, zero value otherwise.
func (o *LegacyWafFirewall) GetRuleStatusesBlockCount() int32 {
	if o == nil || o.RuleStatusesBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.RuleStatusesBlockCount
}

// GetRuleStatusesBlockCountOk returns a tuple with the RuleStatusesBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafFirewall) GetRuleStatusesBlockCountOk() (*int32, bool) {
	if o == nil || o.RuleStatusesBlockCount == nil {
		return nil, false
	}
	return o.RuleStatusesBlockCount, true
}

// HasRuleStatusesBlockCount returns a boolean if a field has been set.
func (o *LegacyWafFirewall) HasRuleStatusesBlockCount() bool {
	if o != nil && o.RuleStatusesBlockCount != nil {
		return true
	}

	return false
}

// SetRuleStatusesBlockCount gets a reference to the given int32 and assigns it to the RuleStatusesBlockCount field.
func (o *LegacyWafFirewall) SetRuleStatusesBlockCount(v int32) {
	o.RuleStatusesBlockCount = &v
}

// GetRuleStatusesDisabledCount returns the RuleStatusesDisabledCount field value if set, zero value otherwise.
func (o *LegacyWafFirewall) GetRuleStatusesDisabledCount() int32 {
	if o == nil || o.RuleStatusesDisabledCount == nil {
		var ret int32
		return ret
	}
	return *o.RuleStatusesDisabledCount
}

// GetRuleStatusesDisabledCountOk returns a tuple with the RuleStatusesDisabledCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafFirewall) GetRuleStatusesDisabledCountOk() (*int32, bool) {
	if o == nil || o.RuleStatusesDisabledCount == nil {
		return nil, false
	}
	return o.RuleStatusesDisabledCount, true
}

// HasRuleStatusesDisabledCount returns a boolean if a field has been set.
func (o *LegacyWafFirewall) HasRuleStatusesDisabledCount() bool {
	if o != nil && o.RuleStatusesDisabledCount != nil {
		return true
	}

	return false
}

// SetRuleStatusesDisabledCount gets a reference to the given int32 and assigns it to the RuleStatusesDisabledCount field.
func (o *LegacyWafFirewall) SetRuleStatusesDisabledCount(v int32) {
	o.RuleStatusesDisabledCount = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LegacyWafFirewall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.LastPush != nil {
		toSerialize["last_push"] = o.LastPush
	}
	if o.PrefetchCondition != nil {
		toSerialize["prefetch_condition"] = o.PrefetchCondition
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.RuleStatusesLogCount != nil {
		toSerialize["rule_statuses_log_count"] = o.RuleStatusesLogCount
	}
	if o.RuleStatusesBlockCount != nil {
		toSerialize["rule_statuses_block_count"] = o.RuleStatusesBlockCount
	}
	if o.RuleStatusesDisabledCount != nil {
		toSerialize["rule_statuses_disabled_count"] = o.RuleStatusesDisabledCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LegacyWafFirewall) UnmarshalJSON(bytes []byte) (err error) {
	varLegacyWafFirewall := _LegacyWafFirewall{}

	if err = json.Unmarshal(bytes, &varLegacyWafFirewall); err == nil {
		*o = LegacyWafFirewall(varLegacyWafFirewall)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "last_push")
		delete(additionalProperties, "prefetch_condition")
		delete(additionalProperties, "response")
		delete(additionalProperties, "version")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "rule_statuses_log_count")
		delete(additionalProperties, "rule_statuses_block_count")
		delete(additionalProperties, "rule_statuses_disabled_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLegacyWafFirewall is a helper abstraction for handling nullable legacywaffirewall types.
type NullableLegacyWafFirewall struct {
	value *LegacyWafFirewall
	isSet bool
}

// Get returns the value.
func (v NullableLegacyWafFirewall) Get() *LegacyWafFirewall {
	return v.value
}

// Set modifies the value.
func (v *NullableLegacyWafFirewall) Set(val *LegacyWafFirewall) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLegacyWafFirewall) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLegacyWafFirewall) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLegacyWafFirewall returns a pointer to a new instance of NullableLegacyWafFirewall.
func NewNullableLegacyWafFirewall(val *LegacyWafFirewall) *NullableLegacyWafFirewall {
	return &NullableLegacyWafFirewall{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLegacyWafFirewall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLegacyWafFirewall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
