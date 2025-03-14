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

// WafFirewallResponseDataAllOf struct for WafFirewallResponseDataAllOf
type WafFirewallResponseDataAllOf struct {
	ID                   *string                            `json:"id,omitempty"`
	Attributes           *WafFirewallResponseDataAttributes `json:"attributes,omitempty"`
	Relationships        *RelationshipWafFirewallVersions   `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallResponseDataAllOf WafFirewallResponseDataAllOf

// NewWafFirewallResponseDataAllOf instantiates a new WafFirewallResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallResponseDataAllOf() *WafFirewallResponseDataAllOf {
	this := WafFirewallResponseDataAllOf{}
	return &this
}

// NewWafFirewallResponseDataAllOfWithDefaults instantiates a new WafFirewallResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallResponseDataAllOfWithDefaults() *WafFirewallResponseDataAllOf {
	this := WafFirewallResponseDataAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafFirewallResponseDataAllOf) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAllOf) GetAttributes() WafFirewallResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafFirewallResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAllOf) GetAttributesOk() (*WafFirewallResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafFirewallResponseDataAttributes and assigns it to the Attributes field.
func (o *WafFirewallResponseDataAllOf) SetAttributes(v WafFirewallResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafFirewallResponseDataAllOf) GetRelationships() RelationshipWafFirewallVersions {
	if o == nil || o.Relationships == nil {
		var ret RelationshipWafFirewallVersions
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponseDataAllOf) GetRelationshipsOk() (*RelationshipWafFirewallVersions, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafFirewallResponseDataAllOf) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipWafFirewallVersions and assigns it to the Relationships field.
func (o *WafFirewallResponseDataAllOf) SetRelationships(v RelationshipWafFirewallVersions) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallResponseDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
func (o *WafFirewallResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallResponseDataAllOf := _WafFirewallResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varWafFirewallResponseDataAllOf); err == nil {
		*o = WafFirewallResponseDataAllOf(varWafFirewallResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallResponseDataAllOf is a helper abstraction for handling nullable waffirewallresponsedataallof types.
type NullableWafFirewallResponseDataAllOf struct {
	value *WafFirewallResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallResponseDataAllOf) Get() *WafFirewallResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallResponseDataAllOf) Set(val *WafFirewallResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallResponseDataAllOf returns a pointer to a new instance of NullableWafFirewallResponseDataAllOf.
func NewNullableWafFirewallResponseDataAllOf(val *WafFirewallResponseDataAllOf) *NullableWafFirewallResponseDataAllOf {
	return &NullableWafFirewallResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafFirewallResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
