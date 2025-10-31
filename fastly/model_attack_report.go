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

// AttackReport struct for AttackReport
type AttackReport struct {
	// ID of the workspace.
	Id string `json:"id"`
	// Name of the workspace.
	Name string `json:"name"`
	// Total request count
	TotalCount int32 `json:"total_count"`
	// Blocked request count
	BlockedCount int32 `json:"blocked_count"`
	// Flagged request count
	FlaggedCount int32 `json:"flagged_count"`
	// Attack request count
	AttackCount int32 `json:"attack_count"`
	// Count of IPs that have been flagged
	AllFlaggedIpCount int32 `json:"all_flagged_ip_count"`
	// Count of currently flagged IPs
	FlaggedIpCount       int32          `json:"flagged_ip_count"`
	TopAttackSignals     []AttackSignal `json:"top_attack_signals"`
	TopAttackSources     []AttackSource `json:"top_attack_sources"`
	AdditionalProperties map[string]any
}

type _AttackReport AttackReport

// NewAttackReport instantiates a new AttackReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReport(id string, name string, totalCount int32, blockedCount int32, flaggedCount int32, attackCount int32, allFlaggedIpCount int32, flaggedIpCount int32, topAttackSignals []AttackSignal, topAttackSources []AttackSource) *AttackReport {
	this := AttackReport{}
	this.Id = id
	this.Name = name
	this.TotalCount = totalCount
	this.BlockedCount = blockedCount
	this.FlaggedCount = flaggedCount
	this.AttackCount = attackCount
	this.AllFlaggedIpCount = allFlaggedIpCount
	this.FlaggedIpCount = flaggedIpCount
	this.TopAttackSignals = topAttackSignals
	this.TopAttackSources = topAttackSources
	return &this
}

// NewAttackReportWithDefaults instantiates a new AttackReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportWithDefaults() *AttackReport {
	this := AttackReport{}
	return &this
}

// GetId returns the Id field value
func (o *AttackReport) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttackReport) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttackReport) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AttackReport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AttackReport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AttackReport) SetName(v string) {
	o.Name = v
}

// GetTotalCount returns the TotalCount field value
func (o *AttackReport) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *AttackReport) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *AttackReport) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetBlockedCount returns the BlockedCount field value
func (o *AttackReport) GetBlockedCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockedCount
}

// GetBlockedCountOk returns a tuple with the BlockedCount field value
// and a boolean to check if the value has been set.
func (o *AttackReport) GetBlockedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockedCount, true
}

// SetBlockedCount sets field value
func (o *AttackReport) SetBlockedCount(v int32) {
	o.BlockedCount = v
}

// GetFlaggedCount returns the FlaggedCount field value
func (o *AttackReport) GetFlaggedCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FlaggedCount
}

// GetFlaggedCountOk returns a tuple with the FlaggedCount field value
// and a boolean to check if the value has been set.
func (o *AttackReport) GetFlaggedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlaggedCount, true
}

// SetFlaggedCount sets field value
func (o *AttackReport) SetFlaggedCount(v int32) {
	o.FlaggedCount = v
}

// GetAttackCount returns the AttackCount field value
func (o *AttackReport) GetAttackCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttackCount
}

// GetAttackCountOk returns a tuple with the AttackCount field value
// and a boolean to check if the value has been set.
func (o *AttackReport) GetAttackCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttackCount, true
}

// SetAttackCount sets field value
func (o *AttackReport) SetAttackCount(v int32) {
	o.AttackCount = v
}

// GetAllFlaggedIpCount returns the AllFlaggedIpCount field value
func (o *AttackReport) GetAllFlaggedIpCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AllFlaggedIpCount
}

// GetAllFlaggedIpCountOk returns a tuple with the AllFlaggedIpCount field value
// and a boolean to check if the value has been set.
func (o *AttackReport) GetAllFlaggedIpCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllFlaggedIpCount, true
}

// SetAllFlaggedIpCount sets field value
func (o *AttackReport) SetAllFlaggedIpCount(v int32) {
	o.AllFlaggedIpCount = v
}

// GetFlaggedIpCount returns the FlaggedIpCount field value
func (o *AttackReport) GetFlaggedIpCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FlaggedIpCount
}

// GetFlaggedIpCountOk returns a tuple with the FlaggedIpCount field value
// and a boolean to check if the value has been set.
func (o *AttackReport) GetFlaggedIpCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlaggedIpCount, true
}

// SetFlaggedIpCount sets field value
func (o *AttackReport) SetFlaggedIpCount(v int32) {
	o.FlaggedIpCount = v
}

// GetTopAttackSignals returns the TopAttackSignals field value
func (o *AttackReport) GetTopAttackSignals() []AttackSignal {
	if o == nil {
		var ret []AttackSignal
		return ret
	}

	return o.TopAttackSignals
}

// GetTopAttackSignalsOk returns a tuple with the TopAttackSignals field value
// and a boolean to check if the value has been set.
func (o *AttackReport) GetTopAttackSignalsOk() ([]AttackSignal, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopAttackSignals, true
}

// SetTopAttackSignals sets field value
func (o *AttackReport) SetTopAttackSignals(v []AttackSignal) {
	o.TopAttackSignals = v
}

// GetTopAttackSources returns the TopAttackSources field value
func (o *AttackReport) GetTopAttackSources() []AttackSource {
	if o == nil {
		var ret []AttackSource
		return ret
	}

	return o.TopAttackSources
}

// GetTopAttackSourcesOk returns a tuple with the TopAttackSources field value
// and a boolean to check if the value has been set.
func (o *AttackReport) GetTopAttackSourcesOk() ([]AttackSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopAttackSources, true
}

// SetTopAttackSources sets field value
func (o *AttackReport) SetTopAttackSources(v []AttackSource) {
	o.TopAttackSources = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AttackReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}
	if true {
		toSerialize["blocked_count"] = o.BlockedCount
	}
	if true {
		toSerialize["flagged_count"] = o.FlaggedCount
	}
	if true {
		toSerialize["attack_count"] = o.AttackCount
	}
	if true {
		toSerialize["all_flagged_ip_count"] = o.AllFlaggedIpCount
	}
	if true {
		toSerialize["flagged_ip_count"] = o.FlaggedIpCount
	}
	if true {
		toSerialize["top_attack_signals"] = o.TopAttackSignals
	}
	if true {
		toSerialize["top_attack_sources"] = o.TopAttackSources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AttackReport) UnmarshalJSON(bytes []byte) (err error) {
	varAttackReport := _AttackReport{}

	if err = json.Unmarshal(bytes, &varAttackReport); err == nil {
		*o = AttackReport(varAttackReport)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "total_count")
		delete(additionalProperties, "blocked_count")
		delete(additionalProperties, "flagged_count")
		delete(additionalProperties, "attack_count")
		delete(additionalProperties, "all_flagged_ip_count")
		delete(additionalProperties, "flagged_ip_count")
		delete(additionalProperties, "top_attack_signals")
		delete(additionalProperties, "top_attack_sources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAttackReport is a helper abstraction for handling nullable attackreport types.
type NullableAttackReport struct {
	value *AttackReport
	isSet bool
}

// Get returns the value.
func (v NullableAttackReport) Get() *AttackReport {
	return v.value
}

// Set modifies the value.
func (v *NullableAttackReport) Set(val *AttackReport) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAttackReport) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAttackReport) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAttackReport returns a pointer to a new instance of NullableAttackReport.
func NewNullableAttackReport(val *AttackReport) *NullableAttackReport {
	return &NullableAttackReport{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAttackReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAttackReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
