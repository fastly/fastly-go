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

// ArcSession A session log. Other than `firewall_enabled` (always present), `firewall_*` fields are only present when AI Firewall ran on a request in this session. They're absent (not merely empty) otherwise, so their absence and an empty value mean different things.
type ArcSession struct {
	// The unique identifier of the session.
	ID string `json:"id"`
	// The ID of the virtual key used for this session.
	VirtualKeyID string `json:"virtual_key_id"`
	// The display name of the virtual key used for this session.
	VirtualKeyName *string `json:"virtual_key_name,omitempty"`
	// The AI model identifier used.
	Model string `json:"model"`
	// The name of the provider connection that ultimately served the session.
	Provider string `json:"provider"`
	// The ID of the provider connection that ultimately served the session.
	ProviderConnectionID string `json:"provider_connection_id"`
	// The ID of the provider connection originally requested, before any failover occurred.
	PrimaryProvider *string `json:"primary_provider,omitempty"`
	// The number of requests made in this session.
	Requests int32 `json:"requests"`
	// The total input tokens consumed in this session.
	InputTokens int32 `json:"input_tokens"`
	// The total output tokens generated in this session.
	OutputTokens int32 `json:"output_tokens"`
	// The number of AI Firewall violations detected in this session.
	Violations int32 `json:"violations"`
	// Whether AI Firewall was enabled for this session's virtual key.
	FirewallEnabled bool `json:"firewall_enabled"`
	// The outcome AI Firewall reached for this session. Present only when AI Firewall ran.
	FirewallOutcome *string `json:"firewall_outcome,omitempty"`
	// The AI Firewall stage that produced the outcome (for example, the request or response stage). Present only when AI Firewall ran.
	FirewallStage *string `json:"firewall_stage,omitempty"`
	// The AI Firewall policy that was applied. Present only when AI Firewall ran.
	FirewallPolicy *string `json:"firewall_policy,omitempty"`
	// Why AI Firewall inspection was incomplete, if applicable. Present only when AI Firewall ran and inspection didn't complete.
	FirewallIncompleteReason *string `json:"firewall_incomplete_reason,omitempty"`
	// The date and time the session was created, in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the session was last updated, in RFC 3339 format.
	UpdatedAt time.Time `json:"updated_at"`
	// The request and response pairs recorded for this session.
	Logs                 []ArcSessionLogEntry `json:"logs"`
	AdditionalProperties map[string]any
}

type _ArcSession ArcSession

// NewArcSession instantiates a new ArcSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcSession(id string, virtualKeyID string, model string, provider string, providerConnectionID string, requests int32, inputTokens int32, outputTokens int32, violations int32, firewallEnabled bool, createdAt time.Time, updatedAt time.Time, logs []ArcSessionLogEntry) *ArcSession {
	this := ArcSession{}
	this.ID = id
	this.VirtualKeyID = virtualKeyID
	this.Model = model
	this.Provider = provider
	this.ProviderConnectionID = providerConnectionID
	this.Requests = requests
	this.InputTokens = inputTokens
	this.OutputTokens = outputTokens
	this.Violations = violations
	this.FirewallEnabled = firewallEnabled
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Logs = logs
	return &this
}

// NewArcSessionWithDefaults instantiates a new ArcSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcSessionWithDefaults() *ArcSession {
	this := ArcSession{}
	return &this
}

// GetID returns the ID field value
func (o *ArcSession) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *ArcSession) SetID(v string) {
	o.ID = v
}

// GetVirtualKeyID returns the VirtualKeyID field value
func (o *ArcSession) GetVirtualKeyID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualKeyID
}

// GetVirtualKeyIDOk returns a tuple with the VirtualKeyID field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetVirtualKeyIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualKeyID, true
}

// SetVirtualKeyID sets field value
func (o *ArcSession) SetVirtualKeyID(v string) {
	o.VirtualKeyID = v
}

// GetVirtualKeyName returns the VirtualKeyName field value if set, zero value otherwise.
func (o *ArcSession) GetVirtualKeyName() string {
	if o == nil || o.VirtualKeyName == nil {
		var ret string
		return ret
	}
	return *o.VirtualKeyName
}

// GetVirtualKeyNameOk returns a tuple with the VirtualKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSession) GetVirtualKeyNameOk() (*string, bool) {
	if o == nil || o.VirtualKeyName == nil {
		return nil, false
	}
	return o.VirtualKeyName, true
}

// HasVirtualKeyName returns a boolean if a field has been set.
func (o *ArcSession) HasVirtualKeyName() bool {
	if o != nil && o.VirtualKeyName != nil {
		return true
	}

	return false
}

// SetVirtualKeyName gets a reference to the given string and assigns it to the VirtualKeyName field.
func (o *ArcSession) SetVirtualKeyName(v string) {
	o.VirtualKeyName = &v
}

// GetModel returns the Model field value
func (o *ArcSession) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ArcSession) SetModel(v string) {
	o.Model = v
}

// GetProvider returns the Provider field value
func (o *ArcSession) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *ArcSession) SetProvider(v string) {
	o.Provider = v
}

// GetProviderConnectionID returns the ProviderConnectionID field value
func (o *ArcSession) GetProviderConnectionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderConnectionID
}

// GetProviderConnectionIDOk returns a tuple with the ProviderConnectionID field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetProviderConnectionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderConnectionID, true
}

// SetProviderConnectionID sets field value
func (o *ArcSession) SetProviderConnectionID(v string) {
	o.ProviderConnectionID = v
}

// GetPrimaryProvider returns the PrimaryProvider field value if set, zero value otherwise.
func (o *ArcSession) GetPrimaryProvider() string {
	if o == nil || o.PrimaryProvider == nil {
		var ret string
		return ret
	}
	return *o.PrimaryProvider
}

// GetPrimaryProviderOk returns a tuple with the PrimaryProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSession) GetPrimaryProviderOk() (*string, bool) {
	if o == nil || o.PrimaryProvider == nil {
		return nil, false
	}
	return o.PrimaryProvider, true
}

// HasPrimaryProvider returns a boolean if a field has been set.
func (o *ArcSession) HasPrimaryProvider() bool {
	if o != nil && o.PrimaryProvider != nil {
		return true
	}

	return false
}

// SetPrimaryProvider gets a reference to the given string and assigns it to the PrimaryProvider field.
func (o *ArcSession) SetPrimaryProvider(v string) {
	o.PrimaryProvider = &v
}

// GetRequests returns the Requests field value
func (o *ArcSession) GetRequests() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetRequestsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value
func (o *ArcSession) SetRequests(v int32) {
	o.Requests = v
}

// GetInputTokens returns the InputTokens field value
func (o *ArcSession) GetInputTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InputTokens
}

// GetInputTokensOk returns a tuple with the InputTokens field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetInputTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputTokens, true
}

// SetInputTokens sets field value
func (o *ArcSession) SetInputTokens(v int32) {
	o.InputTokens = v
}

// GetOutputTokens returns the OutputTokens field value
func (o *ArcSession) GetOutputTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OutputTokens
}

// GetOutputTokensOk returns a tuple with the OutputTokens field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetOutputTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputTokens, true
}

// SetOutputTokens sets field value
func (o *ArcSession) SetOutputTokens(v int32) {
	o.OutputTokens = v
}

// GetViolations returns the Violations field value
func (o *ArcSession) GetViolations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Violations
}

// GetViolationsOk returns a tuple with the Violations field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetViolationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Violations, true
}

// SetViolations sets field value
func (o *ArcSession) SetViolations(v int32) {
	o.Violations = v
}

// GetFirewallEnabled returns the FirewallEnabled field value
func (o *ArcSession) GetFirewallEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FirewallEnabled
}

// GetFirewallEnabledOk returns a tuple with the FirewallEnabled field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetFirewallEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirewallEnabled, true
}

// SetFirewallEnabled sets field value
func (o *ArcSession) SetFirewallEnabled(v bool) {
	o.FirewallEnabled = v
}

// GetFirewallOutcome returns the FirewallOutcome field value if set, zero value otherwise.
func (o *ArcSession) GetFirewallOutcome() string {
	if o == nil || o.FirewallOutcome == nil {
		var ret string
		return ret
	}
	return *o.FirewallOutcome
}

// GetFirewallOutcomeOk returns a tuple with the FirewallOutcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSession) GetFirewallOutcomeOk() (*string, bool) {
	if o == nil || o.FirewallOutcome == nil {
		return nil, false
	}
	return o.FirewallOutcome, true
}

// HasFirewallOutcome returns a boolean if a field has been set.
func (o *ArcSession) HasFirewallOutcome() bool {
	if o != nil && o.FirewallOutcome != nil {
		return true
	}

	return false
}

// SetFirewallOutcome gets a reference to the given string and assigns it to the FirewallOutcome field.
func (o *ArcSession) SetFirewallOutcome(v string) {
	o.FirewallOutcome = &v
}

// GetFirewallStage returns the FirewallStage field value if set, zero value otherwise.
func (o *ArcSession) GetFirewallStage() string {
	if o == nil || o.FirewallStage == nil {
		var ret string
		return ret
	}
	return *o.FirewallStage
}

// GetFirewallStageOk returns a tuple with the FirewallStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSession) GetFirewallStageOk() (*string, bool) {
	if o == nil || o.FirewallStage == nil {
		return nil, false
	}
	return o.FirewallStage, true
}

// HasFirewallStage returns a boolean if a field has been set.
func (o *ArcSession) HasFirewallStage() bool {
	if o != nil && o.FirewallStage != nil {
		return true
	}

	return false
}

// SetFirewallStage gets a reference to the given string and assigns it to the FirewallStage field.
func (o *ArcSession) SetFirewallStage(v string) {
	o.FirewallStage = &v
}

// GetFirewallPolicy returns the FirewallPolicy field value if set, zero value otherwise.
func (o *ArcSession) GetFirewallPolicy() string {
	if o == nil || o.FirewallPolicy == nil {
		var ret string
		return ret
	}
	return *o.FirewallPolicy
}

// GetFirewallPolicyOk returns a tuple with the FirewallPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSession) GetFirewallPolicyOk() (*string, bool) {
	if o == nil || o.FirewallPolicy == nil {
		return nil, false
	}
	return o.FirewallPolicy, true
}

// HasFirewallPolicy returns a boolean if a field has been set.
func (o *ArcSession) HasFirewallPolicy() bool {
	if o != nil && o.FirewallPolicy != nil {
		return true
	}

	return false
}

// SetFirewallPolicy gets a reference to the given string and assigns it to the FirewallPolicy field.
func (o *ArcSession) SetFirewallPolicy(v string) {
	o.FirewallPolicy = &v
}

// GetFirewallIncompleteReason returns the FirewallIncompleteReason field value if set, zero value otherwise.
func (o *ArcSession) GetFirewallIncompleteReason() string {
	if o == nil || o.FirewallIncompleteReason == nil {
		var ret string
		return ret
	}
	return *o.FirewallIncompleteReason
}

// GetFirewallIncompleteReasonOk returns a tuple with the FirewallIncompleteReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSession) GetFirewallIncompleteReasonOk() (*string, bool) {
	if o == nil || o.FirewallIncompleteReason == nil {
		return nil, false
	}
	return o.FirewallIncompleteReason, true
}

// HasFirewallIncompleteReason returns a boolean if a field has been set.
func (o *ArcSession) HasFirewallIncompleteReason() bool {
	if o != nil && o.FirewallIncompleteReason != nil {
		return true
	}

	return false
}

// SetFirewallIncompleteReason gets a reference to the given string and assigns it to the FirewallIncompleteReason field.
func (o *ArcSession) SetFirewallIncompleteReason(v string) {
	o.FirewallIncompleteReason = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArcSession) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArcSession) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArcSession) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArcSession) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetLogs returns the Logs field value
func (o *ArcSession) GetLogs() []ArcSessionLogEntry {
	if o == nil {
		var ret []ArcSessionLogEntry
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *ArcSession) GetLogsOk() ([]ArcSessionLogEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs, true
}

// SetLogs sets field value
func (o *ArcSession) SetLogs(v []ArcSessionLogEntry) {
	o.Logs = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.ID
	}
	if true {
		toSerialize["virtual_key_id"] = o.VirtualKeyID
	}
	if o.VirtualKeyName != nil {
		toSerialize["virtual_key_name"] = o.VirtualKeyName
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["provider_connection_id"] = o.ProviderConnectionID
	}
	if o.PrimaryProvider != nil {
		toSerialize["primary_provider"] = o.PrimaryProvider
	}
	if true {
		toSerialize["requests"] = o.Requests
	}
	if true {
		toSerialize["input_tokens"] = o.InputTokens
	}
	if true {
		toSerialize["output_tokens"] = o.OutputTokens
	}
	if true {
		toSerialize["violations"] = o.Violations
	}
	if true {
		toSerialize["firewall_enabled"] = o.FirewallEnabled
	}
	if o.FirewallOutcome != nil {
		toSerialize["firewall_outcome"] = o.FirewallOutcome
	}
	if o.FirewallStage != nil {
		toSerialize["firewall_stage"] = o.FirewallStage
	}
	if o.FirewallPolicy != nil {
		toSerialize["firewall_policy"] = o.FirewallPolicy
	}
	if o.FirewallIncompleteReason != nil {
		toSerialize["firewall_incomplete_reason"] = o.FirewallIncompleteReason
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["logs"] = o.Logs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcSession) UnmarshalJSON(bytes []byte) (err error) {
	varArcSession := _ArcSession{}

	if err = json.Unmarshal(bytes, &varArcSession); err == nil {
		*o = ArcSession(varArcSession)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "virtual_key_id")
		delete(additionalProperties, "virtual_key_name")
		delete(additionalProperties, "model")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "provider_connection_id")
		delete(additionalProperties, "primary_provider")
		delete(additionalProperties, "requests")
		delete(additionalProperties, "input_tokens")
		delete(additionalProperties, "output_tokens")
		delete(additionalProperties, "violations")
		delete(additionalProperties, "firewall_enabled")
		delete(additionalProperties, "firewall_outcome")
		delete(additionalProperties, "firewall_stage")
		delete(additionalProperties, "firewall_policy")
		delete(additionalProperties, "firewall_incomplete_reason")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "logs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcSession is a helper abstraction for handling nullable arcsession types.
type NullableArcSession struct {
	value *ArcSession
	isSet bool
}

// Get returns the value.
func (v NullableArcSession) Get() *ArcSession {
	return v.value
}

// Set modifies the value.
func (v *NullableArcSession) Set(val *ArcSession) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcSession) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcSession) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcSession returns a pointer to a new instance of NullableArcSession.
func NewNullableArcSession(val *ArcSession) *NullableArcSession {
	return &NullableArcSession{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
