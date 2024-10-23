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

// WafFirewallResponseData struct for WafFirewallResponseData
type WafFirewallResponseData struct {
	Type                 *TypeWafFirewall                   `json:"type,omitempty"`
	Attributes           *WafFirewallResponseDataAttributes `json:"attributes,omitempty"`
	ID                   *string                            `json:"id,omitempty"`
	Relationships        *RelationshipWafFirewallVersions   `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallResponseData WafFirewallResponseData

// NewWafFirewallResponseData instantiates a new WafFirewallResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallResponseData() *WafFirewallResponseData {
	this := WafFirewallResponseData{}
	var resourceType TypeWafFirewall = TYPEWAFFIREWALL_WAF_FIREWALL
	this.Type = &resourceType
	return &this
}

// NewWafFirewallResponseDataWithDefaults instantiates a new WafFirewallResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallResponseDataWithDefaults() *WafFirewallResponseData {
	this := WafFirewallResponseData{}
	var resourceType TypeWafFirewall = TYPEWAFFIREWALL_WAF_FIREWALL
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafFirewallResponseData) GetType() TypeWafFirewall {
	if o == nil || o.Type == nil {
		var ret TypeWafFirewall
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseData) GetTypeOk() (*TypeWafFirewall, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafFirewallResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafFirewall and assigns it to the Type field.
func (o *WafFirewallResponseData) SetType(v TypeWafFirewall) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafFirewallResponseData) GetAttributes() WafFirewallResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafFirewallResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseData) GetAttributesOk() (*WafFirewallResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafFirewallResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafFirewallResponseDataAttributes and assigns it to the Attributes field.
func (o *WafFirewallResponseData) SetAttributes(v WafFirewallResponseDataAttributes) {
	o.Attributes = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafFirewallResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafFirewallResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafFirewallResponseData) SetID(v string) {
	o.ID = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafFirewallResponseData) GetRelationships() RelationshipWafFirewallVersions {
	if o == nil || o.Relationships == nil {
		var ret RelationshipWafFirewallVersions
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseData) GetRelationshipsOk() (*RelationshipWafFirewallVersions, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafFirewallResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipWafFirewallVersions and assigns it to the Relationships field.
func (o *WafFirewallResponseData) SetRelationships(v RelationshipWafFirewallVersions) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafFirewallResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallResponseData := _WafFirewallResponseData{}

	if err = json.Unmarshal(bytes, &varWafFirewallResponseData); err == nil {
		*o = WafFirewallResponseData(varWafFirewallResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "id")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallResponseData is a helper abstraction for handling nullable waffirewallresponsedata types.
type NullableWafFirewallResponseData struct {
	value *WafFirewallResponseData
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallResponseData) Get() *WafFirewallResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallResponseData) Set(val *WafFirewallResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallResponseData returns a pointer to a new instance of NullableWafFirewallResponseData.
func NewNullableWafFirewallResponseData(val *WafFirewallResponseData) *NullableWafFirewallResponseData {
	return &NullableWafFirewallResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafFirewallResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
