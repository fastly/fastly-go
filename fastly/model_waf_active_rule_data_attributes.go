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

// WafActiveRuleDataAttributes struct for WafActiveRuleDataAttributes
type WafActiveRuleDataAttributes struct {
	// The ModSecurity rule ID of the associated rule revision.
	ModsecRuleID *int32 `json:"modsec_rule_id,omitempty"`
	Revision *WafRuleRevisionOrLatest `json:"revision,omitempty"`
	// Describes the behavior for the particular rule revision within this firewall version.
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]any
}

type _WafActiveRuleDataAttributes WafActiveRuleDataAttributes

// NewWafActiveRuleDataAttributes instantiates a new WafActiveRuleDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafActiveRuleDataAttributes() *WafActiveRuleDataAttributes {
	this := WafActiveRuleDataAttributes{}
	return &this
}

// NewWafActiveRuleDataAttributesWithDefaults instantiates a new WafActiveRuleDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafActiveRuleDataAttributesWithDefaults() *WafActiveRuleDataAttributes {
	this := WafActiveRuleDataAttributes{}
	return &this
}

// GetModsecRuleID returns the ModsecRuleID field value if set, zero value otherwise.
func (o *WafActiveRuleDataAttributes) GetModsecRuleID() int32 {
	if o == nil || o.ModsecRuleID == nil {
		var ret int32
		return ret
	}
	return *o.ModsecRuleID
}

// GetModsecRuleIDOk returns a tuple with the ModsecRuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleDataAttributes) GetModsecRuleIDOk() (*int32, bool) {
	if o == nil || o.ModsecRuleID == nil {
		return nil, false
	}
	return o.ModsecRuleID, true
}

// HasModsecRuleID returns a boolean if a field has been set.
func (o *WafActiveRuleDataAttributes) HasModsecRuleID() bool {
	if o != nil && o.ModsecRuleID != nil {
		return true
	}

	return false
}

// SetModsecRuleID gets a reference to the given int32 and assigns it to the ModsecRuleID field.
func (o *WafActiveRuleDataAttributes) SetModsecRuleID(v int32) {
	o.ModsecRuleID = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *WafActiveRuleDataAttributes) GetRevision() WafRuleRevisionOrLatest {
	if o == nil || o.Revision == nil {
		var ret WafRuleRevisionOrLatest
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleDataAttributes) GetRevisionOk() (*WafRuleRevisionOrLatest, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *WafActiveRuleDataAttributes) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given WafRuleRevisionOrLatest and assigns it to the Revision field.
func (o *WafActiveRuleDataAttributes) SetRevision(v WafRuleRevisionOrLatest) {
	o.Revision = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WafActiveRuleDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WafActiveRuleDataAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WafActiveRuleDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafActiveRuleDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ModsecRuleID != nil {
		toSerialize["modsec_rule_id"] = o.ModsecRuleID
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WafActiveRuleDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varWafActiveRuleDataAttributes := _WafActiveRuleDataAttributes{}

	if err = json.Unmarshal(bytes, &varWafActiveRuleDataAttributes); err == nil {
		*o = WafActiveRuleDataAttributes(varWafActiveRuleDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "modsec_rule_id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafActiveRuleDataAttributes is a helper abstraction for handling nullable wafactiveruledataattributes types. 
type NullableWafActiveRuleDataAttributes struct {
	value *WafActiveRuleDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableWafActiveRuleDataAttributes) Get() *WafActiveRuleDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableWafActiveRuleDataAttributes) Set(val *WafActiveRuleDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafActiveRuleDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafActiveRuleDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafActiveRuleDataAttributes returns a pointer to a new instance of NullableWafActiveRuleDataAttributes.
func NewNullableWafActiveRuleDataAttributes(val *WafActiveRuleDataAttributes) *NullableWafActiveRuleDataAttributes {
	return &NullableWafActiveRuleDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafActiveRuleDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafActiveRuleDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
