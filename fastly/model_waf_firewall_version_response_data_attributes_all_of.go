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

// WafFirewallVersionResponseDataAttributesAllOf struct for WafFirewallVersionResponseDataAttributesAllOf
type WafFirewallVersionResponseDataAttributesAllOf struct {
	// Whether a specific firewall version is currently deployed.
	Active *bool `json:"active,omitempty"`
	// The number of active Fastly rules set to block.
	ActiveRulesFastlyBlockCount *int32 `json:"active_rules_fastly_block_count,omitempty"`
	// The number of active Fastly rules set to log.
	ActiveRulesFastlyLogCount *int32 `json:"active_rules_fastly_log_count,omitempty"`
	// The number of active Fastly rules set to score.
	ActiveRulesFastlyScoreCount *int32 `json:"active_rules_fastly_score_count,omitempty"`
	// The number of active OWASP rules set to block.
	ActiveRulesOwaspBlockCount *int32 `json:"active_rules_owasp_block_count,omitempty"`
	// The number of active OWASP rules set to log.
	ActiveRulesOwaspLogCount *int32 `json:"active_rules_owasp_log_count,omitempty"`
	// The number of active OWASP rules set to score.
	ActiveRulesOwaspScoreCount *int32 `json:"active_rules_owasp_score_count,omitempty"`
	// The number of active Trustwave rules set to block.
	ActiveRulesTrustwaveBlockCount *int32 `json:"active_rules_trustwave_block_count,omitempty"`
	// The number of active Trustwave rules set to log.
	ActiveRulesTrustwaveLogCount *int32 `json:"active_rules_trustwave_log_count,omitempty"`
	// The status of the last deployment of this firewall version.
	LastDeploymentStatus NullableString `json:"last_deployment_status,omitempty"`
	// Time-stamp (GMT) indicating when the firewall version was last deployed.
	DeployedAt *string `json:"deployed_at,omitempty"`
	// Contains error message if the firewall version fails to deploy.
	Error                *string `json:"error,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallVersionResponseDataAttributesAllOf WafFirewallVersionResponseDataAttributesAllOf

// NewWafFirewallVersionResponseDataAttributesAllOf instantiates a new WafFirewallVersionResponseDataAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallVersionResponseDataAttributesAllOf() *WafFirewallVersionResponseDataAttributesAllOf {
	this := WafFirewallVersionResponseDataAttributesAllOf{}
	return &this
}

// NewWafFirewallVersionResponseDataAttributesAllOfWithDefaults instantiates a new WafFirewallVersionResponseDataAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallVersionResponseDataAttributesAllOfWithDefaults() *WafFirewallVersionResponseDataAttributesAllOf {
	this := WafFirewallVersionResponseDataAttributesAllOf{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActive(v bool) {
	o.Active = &v
}

// GetActiveRulesFastlyBlockCount returns the ActiveRulesFastlyBlockCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyBlockCount() int32 {
	if o == nil || o.ActiveRulesFastlyBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesFastlyBlockCount
}

// GetActiveRulesFastlyBlockCountOk returns a tuple with the ActiveRulesFastlyBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyBlockCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesFastlyBlockCount == nil {
		return nil, false
	}
	return o.ActiveRulesFastlyBlockCount, true
}

// HasActiveRulesFastlyBlockCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesFastlyBlockCount() bool {
	if o != nil && o.ActiveRulesFastlyBlockCount != nil {
		return true
	}

	return false
}

// SetActiveRulesFastlyBlockCount gets a reference to the given int32 and assigns it to the ActiveRulesFastlyBlockCount field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesFastlyBlockCount(v int32) {
	o.ActiveRulesFastlyBlockCount = &v
}

// GetActiveRulesFastlyLogCount returns the ActiveRulesFastlyLogCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyLogCount() int32 {
	if o == nil || o.ActiveRulesFastlyLogCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesFastlyLogCount
}

// GetActiveRulesFastlyLogCountOk returns a tuple with the ActiveRulesFastlyLogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyLogCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesFastlyLogCount == nil {
		return nil, false
	}
	return o.ActiveRulesFastlyLogCount, true
}

// HasActiveRulesFastlyLogCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesFastlyLogCount() bool {
	if o != nil && o.ActiveRulesFastlyLogCount != nil {
		return true
	}

	return false
}

// SetActiveRulesFastlyLogCount gets a reference to the given int32 and assigns it to the ActiveRulesFastlyLogCount field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesFastlyLogCount(v int32) {
	o.ActiveRulesFastlyLogCount = &v
}

// GetActiveRulesFastlyScoreCount returns the ActiveRulesFastlyScoreCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyScoreCount() int32 {
	if o == nil || o.ActiveRulesFastlyScoreCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesFastlyScoreCount
}

// GetActiveRulesFastlyScoreCountOk returns a tuple with the ActiveRulesFastlyScoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesFastlyScoreCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesFastlyScoreCount == nil {
		return nil, false
	}
	return o.ActiveRulesFastlyScoreCount, true
}

// HasActiveRulesFastlyScoreCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesFastlyScoreCount() bool {
	if o != nil && o.ActiveRulesFastlyScoreCount != nil {
		return true
	}

	return false
}

// SetActiveRulesFastlyScoreCount gets a reference to the given int32 and assigns it to the ActiveRulesFastlyScoreCount field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesFastlyScoreCount(v int32) {
	o.ActiveRulesFastlyScoreCount = &v
}

// GetActiveRulesOwaspBlockCount returns the ActiveRulesOwaspBlockCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspBlockCount() int32 {
	if o == nil || o.ActiveRulesOwaspBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesOwaspBlockCount
}

// GetActiveRulesOwaspBlockCountOk returns a tuple with the ActiveRulesOwaspBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspBlockCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesOwaspBlockCount == nil {
		return nil, false
	}
	return o.ActiveRulesOwaspBlockCount, true
}

// HasActiveRulesOwaspBlockCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesOwaspBlockCount() bool {
	if o != nil && o.ActiveRulesOwaspBlockCount != nil {
		return true
	}

	return false
}

// SetActiveRulesOwaspBlockCount gets a reference to the given int32 and assigns it to the ActiveRulesOwaspBlockCount field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesOwaspBlockCount(v int32) {
	o.ActiveRulesOwaspBlockCount = &v
}

// GetActiveRulesOwaspLogCount returns the ActiveRulesOwaspLogCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspLogCount() int32 {
	if o == nil || o.ActiveRulesOwaspLogCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesOwaspLogCount
}

// GetActiveRulesOwaspLogCountOk returns a tuple with the ActiveRulesOwaspLogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspLogCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesOwaspLogCount == nil {
		return nil, false
	}
	return o.ActiveRulesOwaspLogCount, true
}

// HasActiveRulesOwaspLogCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesOwaspLogCount() bool {
	if o != nil && o.ActiveRulesOwaspLogCount != nil {
		return true
	}

	return false
}

// SetActiveRulesOwaspLogCount gets a reference to the given int32 and assigns it to the ActiveRulesOwaspLogCount field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesOwaspLogCount(v int32) {
	o.ActiveRulesOwaspLogCount = &v
}

// GetActiveRulesOwaspScoreCount returns the ActiveRulesOwaspScoreCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspScoreCount() int32 {
	if o == nil || o.ActiveRulesOwaspScoreCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesOwaspScoreCount
}

// GetActiveRulesOwaspScoreCountOk returns a tuple with the ActiveRulesOwaspScoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesOwaspScoreCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesOwaspScoreCount == nil {
		return nil, false
	}
	return o.ActiveRulesOwaspScoreCount, true
}

// HasActiveRulesOwaspScoreCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesOwaspScoreCount() bool {
	if o != nil && o.ActiveRulesOwaspScoreCount != nil {
		return true
	}

	return false
}

// SetActiveRulesOwaspScoreCount gets a reference to the given int32 and assigns it to the ActiveRulesOwaspScoreCount field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesOwaspScoreCount(v int32) {
	o.ActiveRulesOwaspScoreCount = &v
}

// GetActiveRulesTrustwaveBlockCount returns the ActiveRulesTrustwaveBlockCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesTrustwaveBlockCount() int32 {
	if o == nil || o.ActiveRulesTrustwaveBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesTrustwaveBlockCount
}

// GetActiveRulesTrustwaveBlockCountOk returns a tuple with the ActiveRulesTrustwaveBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesTrustwaveBlockCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesTrustwaveBlockCount == nil {
		return nil, false
	}
	return o.ActiveRulesTrustwaveBlockCount, true
}

// HasActiveRulesTrustwaveBlockCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesTrustwaveBlockCount() bool {
	if o != nil && o.ActiveRulesTrustwaveBlockCount != nil {
		return true
	}

	return false
}

// SetActiveRulesTrustwaveBlockCount gets a reference to the given int32 and assigns it to the ActiveRulesTrustwaveBlockCount field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesTrustwaveBlockCount(v int32) {
	o.ActiveRulesTrustwaveBlockCount = &v
}

// GetActiveRulesTrustwaveLogCount returns the ActiveRulesTrustwaveLogCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesTrustwaveLogCount() int32 {
	if o == nil || o.ActiveRulesTrustwaveLogCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesTrustwaveLogCount
}

// GetActiveRulesTrustwaveLogCountOk returns a tuple with the ActiveRulesTrustwaveLogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetActiveRulesTrustwaveLogCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesTrustwaveLogCount == nil {
		return nil, false
	}
	return o.ActiveRulesTrustwaveLogCount, true
}

// HasActiveRulesTrustwaveLogCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasActiveRulesTrustwaveLogCount() bool {
	if o != nil && o.ActiveRulesTrustwaveLogCount != nil {
		return true
	}

	return false
}

// SetActiveRulesTrustwaveLogCount gets a reference to the given int32 and assigns it to the ActiveRulesTrustwaveLogCount field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetActiveRulesTrustwaveLogCount(v int32) {
	o.ActiveRulesTrustwaveLogCount = &v
}

// GetLastDeploymentStatus returns the LastDeploymentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetLastDeploymentStatus() string {
	if o == nil || o.LastDeploymentStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastDeploymentStatus.Get()
}

// GetLastDeploymentStatusOk returns a tuple with the LastDeploymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetLastDeploymentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastDeploymentStatus.Get(), o.LastDeploymentStatus.IsSet()
}

// HasLastDeploymentStatus returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasLastDeploymentStatus() bool {
	if o != nil && o.LastDeploymentStatus.IsSet() {
		return true
	}

	return false
}

// SetLastDeploymentStatus gets a reference to the given NullableString and assigns it to the LastDeploymentStatus field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetLastDeploymentStatus(v string) {
	o.LastDeploymentStatus.Set(&v)
}

// SetLastDeploymentStatusNil sets the value for LastDeploymentStatus to be an explicit nil
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetLastDeploymentStatusNil() {
	o.LastDeploymentStatus.Set(nil)
}

// UnsetLastDeploymentStatus ensures that no value is present for LastDeploymentStatus, not even an explicit nil
func (o *WafFirewallVersionResponseDataAttributesAllOf) UnsetLastDeploymentStatus() {
	o.LastDeploymentStatus.Unset()
}

// GetDeployedAt returns the DeployedAt field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetDeployedAt() string {
	if o == nil || o.DeployedAt == nil {
		var ret string
		return ret
	}
	return *o.DeployedAt
}

// GetDeployedAtOk returns a tuple with the DeployedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetDeployedAtOk() (*string, bool) {
	if o == nil || o.DeployedAt == nil {
		return nil, false
	}
	return o.DeployedAt, true
}

// HasDeployedAt returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasDeployedAt() bool {
	if o != nil && o.DeployedAt != nil {
		return true
	}

	return false
}

// SetDeployedAt gets a reference to the given string and assigns it to the DeployedAt field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetDeployedAt(v string) {
	o.DeployedAt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributesAllOf) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *WafFirewallVersionResponseDataAttributesAllOf) SetError(v string) {
	o.Error = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallVersionResponseDataAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
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
	if o.LastDeploymentStatus.IsSet() {
		toSerialize["last_deployment_status"] = o.LastDeploymentStatus.Get()
	}
	if o.DeployedAt != nil {
		toSerialize["deployed_at"] = o.DeployedAt
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafFirewallVersionResponseDataAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallVersionResponseDataAttributesAllOf := _WafFirewallVersionResponseDataAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varWafFirewallVersionResponseDataAttributesAllOf); err == nil {
		*o = WafFirewallVersionResponseDataAttributesAllOf(varWafFirewallVersionResponseDataAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "active_rules_fastly_block_count")
		delete(additionalProperties, "active_rules_fastly_log_count")
		delete(additionalProperties, "active_rules_fastly_score_count")
		delete(additionalProperties, "active_rules_owasp_block_count")
		delete(additionalProperties, "active_rules_owasp_log_count")
		delete(additionalProperties, "active_rules_owasp_score_count")
		delete(additionalProperties, "active_rules_trustwave_block_count")
		delete(additionalProperties, "active_rules_trustwave_log_count")
		delete(additionalProperties, "last_deployment_status")
		delete(additionalProperties, "deployed_at")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallVersionResponseDataAttributesAllOf is a helper abstraction for handling nullable waffirewallversionresponsedataattributesallof types.
type NullableWafFirewallVersionResponseDataAttributesAllOf struct {
	value *WafFirewallVersionResponseDataAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallVersionResponseDataAttributesAllOf) Get() *WafFirewallVersionResponseDataAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallVersionResponseDataAttributesAllOf) Set(val *WafFirewallVersionResponseDataAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallVersionResponseDataAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallVersionResponseDataAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallVersionResponseDataAttributesAllOf returns a pointer to a new instance of NullableWafFirewallVersionResponseDataAttributesAllOf.
func NewNullableWafFirewallVersionResponseDataAttributesAllOf(val *WafFirewallVersionResponseDataAttributesAllOf) *NullableWafFirewallVersionResponseDataAttributesAllOf {
	return &NullableWafFirewallVersionResponseDataAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallVersionResponseDataAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafFirewallVersionResponseDataAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
