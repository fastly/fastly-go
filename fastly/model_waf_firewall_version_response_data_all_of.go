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

// WafFirewallVersionResponseDataAllOf struct for WafFirewallVersionResponseDataAllOf
type WafFirewallVersionResponseDataAllOf struct {
	// Alphanumeric string identifying a Firewall version.
	ID *string `json:"id,omitempty"`
	Attributes *WafFirewallVersionResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipsForWafFirewallVersion `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallVersionResponseDataAllOf WafFirewallVersionResponseDataAllOf

// NewWafFirewallVersionResponseDataAllOf instantiates a new WafFirewallVersionResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallVersionResponseDataAllOf() *WafFirewallVersionResponseDataAllOf {
	this := WafFirewallVersionResponseDataAllOf{}
	return &this
}

// NewWafFirewallVersionResponseDataAllOfWithDefaults instantiates a new WafFirewallVersionResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallVersionResponseDataAllOfWithDefaults() *WafFirewallVersionResponseDataAllOf {
	this := WafFirewallVersionResponseDataAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafFirewallVersionResponseDataAllOf) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAllOf) GetAttributes() WafFirewallVersionResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafFirewallVersionResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAllOf) GetAttributesOk() (*WafFirewallVersionResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafFirewallVersionResponseDataAttributes and assigns it to the Attributes field.
func (o *WafFirewallVersionResponseDataAllOf) SetAttributes(v WafFirewallVersionResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseDataAllOf) GetRelationships() RelationshipsForWafFirewallVersion {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForWafFirewallVersion
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseDataAllOf) GetRelationshipsOk() (*RelationshipsForWafFirewallVersion, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseDataAllOf) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForWafFirewallVersion and assigns it to the Relationships field.
func (o *WafFirewallVersionResponseDataAllOf) SetRelationships(v RelationshipsForWafFirewallVersion) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallVersionResponseDataAllOf) MarshalJSON() ([]byte, error) {
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
func (o *WafFirewallVersionResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallVersionResponseDataAllOf := _WafFirewallVersionResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varWafFirewallVersionResponseDataAllOf); err == nil {
		*o = WafFirewallVersionResponseDataAllOf(varWafFirewallVersionResponseDataAllOf)
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

// NullableWafFirewallVersionResponseDataAllOf is a helper abstraction for handling nullable waffirewallversionresponsedataallof types. 
type NullableWafFirewallVersionResponseDataAllOf struct {
	value *WafFirewallVersionResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallVersionResponseDataAllOf) Get() *WafFirewallVersionResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallVersionResponseDataAllOf) Set(val *WafFirewallVersionResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallVersionResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallVersionResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallVersionResponseDataAllOf returns a pointer to a new instance of NullableWafFirewallVersionResponseDataAllOf.
func NewNullableWafFirewallVersionResponseDataAllOf(val *WafFirewallVersionResponseDataAllOf) *NullableWafFirewallVersionResponseDataAllOf {
	return &NullableWafFirewallVersionResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallVersionResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallVersionResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
