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

// WafRuleRevisionAttributes struct for WafRuleRevisionAttributes
type WafRuleRevisionAttributes struct {
	// Message metadata for the rule.
	Message *string `json:"message,omitempty"`
	// Corresponding ModSecurity rule ID.
	ModsecRuleID *int32 `json:"modsec_rule_id,omitempty"`
	// Paranoia level for the rule.
	ParanoiaLevel *int32 `json:"paranoia_level,omitempty"`
	// Revision number.
	Revision *int32 `json:"revision,omitempty"`
	// Severity metadata for the rule.
	Severity *int32 `json:"severity,omitempty"`
	// The ModSecurity rule logic.
	Source *string `json:"source,omitempty"`
	// The state, indicating if the revision is the most recent version of the rule.
	State *string `json:"state,omitempty"`
	// The VCL representation of the rule logic.
	Vcl *string `json:"vcl,omitempty"`
	AdditionalProperties map[string]any
}

type _WafRuleRevisionAttributes WafRuleRevisionAttributes

// NewWafRuleRevisionAttributes instantiates a new WafRuleRevisionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleRevisionAttributes() *WafRuleRevisionAttributes {
	this := WafRuleRevisionAttributes{}
	return &this
}

// NewWafRuleRevisionAttributesWithDefaults instantiates a new WafRuleRevisionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleRevisionAttributesWithDefaults() *WafRuleRevisionAttributes {
	this := WafRuleRevisionAttributes{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *WafRuleRevisionAttributes) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionAttributes) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *WafRuleRevisionAttributes) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *WafRuleRevisionAttributes) SetMessage(v string) {
	o.Message = &v
}

// GetModsecRuleID returns the ModsecRuleID field value if set, zero value otherwise.
func (o *WafRuleRevisionAttributes) GetModsecRuleID() int32 {
	if o == nil || o.ModsecRuleID == nil {
		var ret int32
		return ret
	}
	return *o.ModsecRuleID
}

// GetModsecRuleIDOk returns a tuple with the ModsecRuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionAttributes) GetModsecRuleIDOk() (*int32, bool) {
	if o == nil || o.ModsecRuleID == nil {
		return nil, false
	}
	return o.ModsecRuleID, true
}

// HasModsecRuleID returns a boolean if a field has been set.
func (o *WafRuleRevisionAttributes) HasModsecRuleID() bool {
	if o != nil && o.ModsecRuleID != nil {
		return true
	}

	return false
}

// SetModsecRuleID gets a reference to the given int32 and assigns it to the ModsecRuleID field.
func (o *WafRuleRevisionAttributes) SetModsecRuleID(v int32) {
	o.ModsecRuleID = &v
}

// GetParanoiaLevel returns the ParanoiaLevel field value if set, zero value otherwise.
func (o *WafRuleRevisionAttributes) GetParanoiaLevel() int32 {
	if o == nil || o.ParanoiaLevel == nil {
		var ret int32
		return ret
	}
	return *o.ParanoiaLevel
}

// GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionAttributes) GetParanoiaLevelOk() (*int32, bool) {
	if o == nil || o.ParanoiaLevel == nil {
		return nil, false
	}
	return o.ParanoiaLevel, true
}

// HasParanoiaLevel returns a boolean if a field has been set.
func (o *WafRuleRevisionAttributes) HasParanoiaLevel() bool {
	if o != nil && o.ParanoiaLevel != nil {
		return true
	}

	return false
}

// SetParanoiaLevel gets a reference to the given int32 and assigns it to the ParanoiaLevel field.
func (o *WafRuleRevisionAttributes) SetParanoiaLevel(v int32) {
	o.ParanoiaLevel = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *WafRuleRevisionAttributes) GetRevision() int32 {
	if o == nil || o.Revision == nil {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionAttributes) GetRevisionOk() (*int32, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *WafRuleRevisionAttributes) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *WafRuleRevisionAttributes) SetRevision(v int32) {
	o.Revision = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *WafRuleRevisionAttributes) GetSeverity() int32 {
	if o == nil || o.Severity == nil {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionAttributes) GetSeverityOk() (*int32, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *WafRuleRevisionAttributes) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *WafRuleRevisionAttributes) SetSeverity(v int32) {
	o.Severity = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *WafRuleRevisionAttributes) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionAttributes) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *WafRuleRevisionAttributes) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *WafRuleRevisionAttributes) SetSource(v string) {
	o.Source = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *WafRuleRevisionAttributes) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionAttributes) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *WafRuleRevisionAttributes) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *WafRuleRevisionAttributes) SetState(v string) {
	o.State = &v
}

// GetVcl returns the Vcl field value if set, zero value otherwise.
func (o *WafRuleRevisionAttributes) GetVcl() string {
	if o == nil || o.Vcl == nil {
		var ret string
		return ret
	}
	return *o.Vcl
}

// GetVclOk returns a tuple with the Vcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionAttributes) GetVclOk() (*string, bool) {
	if o == nil || o.Vcl == nil {
		return nil, false
	}
	return o.Vcl, true
}

// HasVcl returns a boolean if a field has been set.
func (o *WafRuleRevisionAttributes) HasVcl() bool {
	if o != nil && o.Vcl != nil {
		return true
	}

	return false
}

// SetVcl gets a reference to the given string and assigns it to the Vcl field.
func (o *WafRuleRevisionAttributes) SetVcl(v string) {
	o.Vcl = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRuleRevisionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ModsecRuleID != nil {
		toSerialize["modsec_rule_id"] = o.ModsecRuleID
	}
	if o.ParanoiaLevel != nil {
		toSerialize["paranoia_level"] = o.ParanoiaLevel
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.State != nil {
		toSerialize["state"] = o.State
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
func (o *WafRuleRevisionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varWafRuleRevisionAttributes := _WafRuleRevisionAttributes{}

	if err = json.Unmarshal(bytes, &varWafRuleRevisionAttributes); err == nil {
		*o = WafRuleRevisionAttributes(varWafRuleRevisionAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "modsec_rule_id")
		delete(additionalProperties, "paranoia_level")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "source")
		delete(additionalProperties, "state")
		delete(additionalProperties, "vcl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafRuleRevisionAttributes is a helper abstraction for handling nullable wafrulerevisionattributes types. 
type NullableWafRuleRevisionAttributes struct {
	value *WafRuleRevisionAttributes
	isSet bool
}

// Get returns the value.
func (v NullableWafRuleRevisionAttributes) Get() *WafRuleRevisionAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRuleRevisionAttributes) Set(val *WafRuleRevisionAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRuleRevisionAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRuleRevisionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRuleRevisionAttributes returns a pointer to a new instance of NullableWafRuleRevisionAttributes.
func NewNullableWafRuleRevisionAttributes(val *WafRuleRevisionAttributes) *NullableWafRuleRevisionAttributes {
	return &NullableWafRuleRevisionAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRuleRevisionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafRuleRevisionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
