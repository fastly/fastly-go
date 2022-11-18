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
	"time"
)

// WafFirewallVersionResponseDataAttributes struct for WafFirewallVersionResponseDataAttributes
type WafFirewallVersionResponseDataAttributes struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
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
	Error *string `json:"error,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallVersionResponseDataAttributes WafFirewallVersionResponseDataAttributes

// NewWafFirewallVersionResponseDataAttributes instantiates a new WafFirewallVersionResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallVersionResponseDataAttributes() *WafFirewallVersionResponseDataAttributes {
	this := WafFirewallVersionResponseDataAttributes{}
	return &this
}

// NewWafFirewallVersionResponseDataAttributesWithDefaults instantiates a new WafFirewallVersionResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallVersionResponseDataAttributesWithDefaults() *WafFirewallVersionResponseDataAttributes {
	this := WafFirewallVersionResponseDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafFirewallVersionResponseDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafFirewallVersionResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *WafFirewallVersionResponseDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *WafFirewallVersionResponseDataAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *WafFirewallVersionResponseDataAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafFirewallVersionResponseDataAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafFirewallVersionResponseDataAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *WafFirewallVersionResponseDataAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *WafFirewallVersionResponseDataAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *WafFirewallVersionResponseDataAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafFirewallVersionResponseDataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafFirewallVersionResponseDataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *WafFirewallVersionResponseDataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *WafFirewallVersionResponseDataAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *WafFirewallVersionResponseDataAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WafFirewallVersionResponseDataAttributes) SetActive(v bool) {
	o.Active = &v
}

// GetActiveRulesFastlyBlockCount returns the ActiveRulesFastlyBlockCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyBlockCount() int32 {
	if o == nil || o.ActiveRulesFastlyBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesFastlyBlockCount
}

// GetActiveRulesFastlyBlockCountOk returns a tuple with the ActiveRulesFastlyBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyBlockCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesFastlyBlockCount == nil {
		return nil, false
	}
	return o.ActiveRulesFastlyBlockCount, true
}

// HasActiveRulesFastlyBlockCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesFastlyBlockCount() bool {
	if o != nil && o.ActiveRulesFastlyBlockCount != nil {
		return true
	}

	return false
}

// SetActiveRulesFastlyBlockCount gets a reference to the given int32 and assigns it to the ActiveRulesFastlyBlockCount field.
func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesFastlyBlockCount(v int32) {
	o.ActiveRulesFastlyBlockCount = &v
}

// GetActiveRulesFastlyLogCount returns the ActiveRulesFastlyLogCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyLogCount() int32 {
	if o == nil || o.ActiveRulesFastlyLogCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesFastlyLogCount
}

// GetActiveRulesFastlyLogCountOk returns a tuple with the ActiveRulesFastlyLogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyLogCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesFastlyLogCount == nil {
		return nil, false
	}
	return o.ActiveRulesFastlyLogCount, true
}

// HasActiveRulesFastlyLogCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesFastlyLogCount() bool {
	if o != nil && o.ActiveRulesFastlyLogCount != nil {
		return true
	}

	return false
}

// SetActiveRulesFastlyLogCount gets a reference to the given int32 and assigns it to the ActiveRulesFastlyLogCount field.
func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesFastlyLogCount(v int32) {
	o.ActiveRulesFastlyLogCount = &v
}

// GetActiveRulesFastlyScoreCount returns the ActiveRulesFastlyScoreCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyScoreCount() int32 {
	if o == nil || o.ActiveRulesFastlyScoreCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesFastlyScoreCount
}

// GetActiveRulesFastlyScoreCountOk returns a tuple with the ActiveRulesFastlyScoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyScoreCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesFastlyScoreCount == nil {
		return nil, false
	}
	return o.ActiveRulesFastlyScoreCount, true
}

// HasActiveRulesFastlyScoreCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesFastlyScoreCount() bool {
	if o != nil && o.ActiveRulesFastlyScoreCount != nil {
		return true
	}

	return false
}

// SetActiveRulesFastlyScoreCount gets a reference to the given int32 and assigns it to the ActiveRulesFastlyScoreCount field.
func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesFastlyScoreCount(v int32) {
	o.ActiveRulesFastlyScoreCount = &v
}

// GetActiveRulesOwaspBlockCount returns the ActiveRulesOwaspBlockCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspBlockCount() int32 {
	if o == nil || o.ActiveRulesOwaspBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesOwaspBlockCount
}

// GetActiveRulesOwaspBlockCountOk returns a tuple with the ActiveRulesOwaspBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspBlockCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesOwaspBlockCount == nil {
		return nil, false
	}
	return o.ActiveRulesOwaspBlockCount, true
}

// HasActiveRulesOwaspBlockCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesOwaspBlockCount() bool {
	if o != nil && o.ActiveRulesOwaspBlockCount != nil {
		return true
	}

	return false
}

// SetActiveRulesOwaspBlockCount gets a reference to the given int32 and assigns it to the ActiveRulesOwaspBlockCount field.
func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesOwaspBlockCount(v int32) {
	o.ActiveRulesOwaspBlockCount = &v
}

// GetActiveRulesOwaspLogCount returns the ActiveRulesOwaspLogCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspLogCount() int32 {
	if o == nil || o.ActiveRulesOwaspLogCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesOwaspLogCount
}

// GetActiveRulesOwaspLogCountOk returns a tuple with the ActiveRulesOwaspLogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspLogCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesOwaspLogCount == nil {
		return nil, false
	}
	return o.ActiveRulesOwaspLogCount, true
}

// HasActiveRulesOwaspLogCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesOwaspLogCount() bool {
	if o != nil && o.ActiveRulesOwaspLogCount != nil {
		return true
	}

	return false
}

// SetActiveRulesOwaspLogCount gets a reference to the given int32 and assigns it to the ActiveRulesOwaspLogCount field.
func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesOwaspLogCount(v int32) {
	o.ActiveRulesOwaspLogCount = &v
}

// GetActiveRulesOwaspScoreCount returns the ActiveRulesOwaspScoreCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspScoreCount() int32 {
	if o == nil || o.ActiveRulesOwaspScoreCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesOwaspScoreCount
}

// GetActiveRulesOwaspScoreCountOk returns a tuple with the ActiveRulesOwaspScoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspScoreCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesOwaspScoreCount == nil {
		return nil, false
	}
	return o.ActiveRulesOwaspScoreCount, true
}

// HasActiveRulesOwaspScoreCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesOwaspScoreCount() bool {
	if o != nil && o.ActiveRulesOwaspScoreCount != nil {
		return true
	}

	return false
}

// SetActiveRulesOwaspScoreCount gets a reference to the given int32 and assigns it to the ActiveRulesOwaspScoreCount field.
func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesOwaspScoreCount(v int32) {
	o.ActiveRulesOwaspScoreCount = &v
}

// GetActiveRulesTrustwaveBlockCount returns the ActiveRulesTrustwaveBlockCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesTrustwaveBlockCount() int32 {
	if o == nil || o.ActiveRulesTrustwaveBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesTrustwaveBlockCount
}

// GetActiveRulesTrustwaveBlockCountOk returns a tuple with the ActiveRulesTrustwaveBlockCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesTrustwaveBlockCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesTrustwaveBlockCount == nil {
		return nil, false
	}
	return o.ActiveRulesTrustwaveBlockCount, true
}

// HasActiveRulesTrustwaveBlockCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesTrustwaveBlockCount() bool {
	if o != nil && o.ActiveRulesTrustwaveBlockCount != nil {
		return true
	}

	return false
}

// SetActiveRulesTrustwaveBlockCount gets a reference to the given int32 and assigns it to the ActiveRulesTrustwaveBlockCount field.
func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesTrustwaveBlockCount(v int32) {
	o.ActiveRulesTrustwaveBlockCount = &v
}

// GetActiveRulesTrustwaveLogCount returns the ActiveRulesTrustwaveLogCount field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesTrustwaveLogCount() int32 {
	if o == nil || o.ActiveRulesTrustwaveLogCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesTrustwaveLogCount
}

// GetActiveRulesTrustwaveLogCountOk returns a tuple with the ActiveRulesTrustwaveLogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesTrustwaveLogCountOk() (*int32, bool) {
	if o == nil || o.ActiveRulesTrustwaveLogCount == nil {
		return nil, false
	}
	return o.ActiveRulesTrustwaveLogCount, true
}

// HasActiveRulesTrustwaveLogCount returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasActiveRulesTrustwaveLogCount() bool {
	if o != nil && o.ActiveRulesTrustwaveLogCount != nil {
		return true
	}

	return false
}

// SetActiveRulesTrustwaveLogCount gets a reference to the given int32 and assigns it to the ActiveRulesTrustwaveLogCount field.
func (o *WafFirewallVersionResponseDataAttributes) SetActiveRulesTrustwaveLogCount(v int32) {
	o.ActiveRulesTrustwaveLogCount = &v
}

// GetLastDeploymentStatus returns the LastDeploymentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafFirewallVersionResponseDataAttributes) GetLastDeploymentStatus() string {
	if o == nil || o.LastDeploymentStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastDeploymentStatus.Get()
}

// GetLastDeploymentStatusOk returns a tuple with the LastDeploymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafFirewallVersionResponseDataAttributes) GetLastDeploymentStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastDeploymentStatus.Get(), o.LastDeploymentStatus.IsSet()
}

// HasLastDeploymentStatus returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasLastDeploymentStatus() bool {
	if o != nil && o.LastDeploymentStatus.IsSet() {
		return true
	}

	return false
}

// SetLastDeploymentStatus gets a reference to the given NullableString and assigns it to the LastDeploymentStatus field.
func (o *WafFirewallVersionResponseDataAttributes) SetLastDeploymentStatus(v string) {
	o.LastDeploymentStatus.Set(&v)
}
// SetLastDeploymentStatusNil sets the value for LastDeploymentStatus to be an explicit nil
func (o *WafFirewallVersionResponseDataAttributes) SetLastDeploymentStatusNil() {
	o.LastDeploymentStatus.Set(nil)
}

// UnsetLastDeploymentStatus ensures that no value is present for LastDeploymentStatus, not even an explicit nil
func (o *WafFirewallVersionResponseDataAttributes) UnsetLastDeploymentStatus() {
	o.LastDeploymentStatus.Unset()
}

// GetDeployedAt returns the DeployedAt field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetDeployedAt() string {
	if o == nil || o.DeployedAt == nil {
		var ret string
		return ret
	}
	return *o.DeployedAt
}

// GetDeployedAtOk returns a tuple with the DeployedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetDeployedAtOk() (*string, bool) {
	if o == nil || o.DeployedAt == nil {
		return nil, false
	}
	return o.DeployedAt, true
}

// HasDeployedAt returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasDeployedAt() bool {
	if o != nil && o.DeployedAt != nil {
		return true
	}

	return false
}

// SetDeployedAt gets a reference to the given string and assigns it to the DeployedAt field.
func (o *WafFirewallVersionResponseDataAttributes) SetDeployedAt(v string) {
	o.DeployedAt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAttributes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAttributes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAttributes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *WafFirewallVersionResponseDataAttributes) SetError(v string) {
	o.Error = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallVersionResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
func (o *WafFirewallVersionResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallVersionResponseDataAttributes := _WafFirewallVersionResponseDataAttributes{}

	if err = json.Unmarshal(bytes, &varWafFirewallVersionResponseDataAttributes); err == nil {
		*o = WafFirewallVersionResponseDataAttributes(varWafFirewallVersionResponseDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
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

// NullableWafFirewallVersionResponseDataAttributes is a helper abstraction for handling nullable waffirewallversionresponsedataattributes types. 
type NullableWafFirewallVersionResponseDataAttributes struct {
	value *WafFirewallVersionResponseDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallVersionResponseDataAttributes) Get() *WafFirewallVersionResponseDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallVersionResponseDataAttributes) Set(val *WafFirewallVersionResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallVersionResponseDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallVersionResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallVersionResponseDataAttributes returns a pointer to a new instance of NullableWafFirewallVersionResponseDataAttributes.
func NewNullableWafFirewallVersionResponseDataAttributes(val *WafFirewallVersionResponseDataAttributes) *NullableWafFirewallVersionResponseDataAttributes {
	return &NullableWafFirewallVersionResponseDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallVersionResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallVersionResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
