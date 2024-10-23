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

// WafFirewallVersionResponseData struct for WafFirewallVersionResponseData
type WafFirewallVersionResponseData struct {
	Type       *TypeWafFirewallVersion                   `json:"type,omitempty"`
	Attributes *WafFirewallVersionResponseDataAttributes `json:"attributes,omitempty"`
	// Alphanumeric string identifying a Firewall version.
	ID                   *string                             `json:"id,omitempty"`
	Relationships        *RelationshipsForWafFirewallVersion `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallVersionResponseData WafFirewallVersionResponseData

// NewWafFirewallVersionResponseData instantiates a new WafFirewallVersionResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallVersionResponseData() *WafFirewallVersionResponseData {
	this := WafFirewallVersionResponseData{}
	var resourceType TypeWafFirewallVersion = TYPEWAFFIREWALLVERSION_WAF_FIREWALL_VERSION
	this.Type = &resourceType
	return &this
}

// NewWafFirewallVersionResponseDataWithDefaults instantiates a new WafFirewallVersionResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallVersionResponseDataWithDefaults() *WafFirewallVersionResponseData {
	this := WafFirewallVersionResponseData{}
	var resourceType TypeWafFirewallVersion = TYPEWAFFIREWALLVERSION_WAF_FIREWALL_VERSION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseData) GetType() TypeWafFirewallVersion {
	if o == nil || o.Type == nil {
		var ret TypeWafFirewallVersion
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseData) GetTypeOk() (*TypeWafFirewallVersion, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafFirewallVersion and assigns it to the Type field.
func (o *WafFirewallVersionResponseData) SetType(v TypeWafFirewallVersion) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseData) GetAttributes() WafFirewallVersionResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafFirewallVersionResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseData) GetAttributesOk() (*WafFirewallVersionResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafFirewallVersionResponseDataAttributes and assigns it to the Attributes field.
func (o *WafFirewallVersionResponseData) SetAttributes(v WafFirewallVersionResponseDataAttributes) {
	o.Attributes = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafFirewallVersionResponseData) SetID(v string) {
	o.ID = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafFirewallVersionResponseData) GetRelationships() RelationshipsForWafFirewallVersion {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForWafFirewallVersion
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionResponseData) GetRelationshipsOk() (*RelationshipsForWafFirewallVersion, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafFirewallVersionResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForWafFirewallVersion and assigns it to the Relationships field.
func (o *WafFirewallVersionResponseData) SetRelationships(v RelationshipsForWafFirewallVersion) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallVersionResponseData) MarshalJSON() ([]byte, error) {
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
func (o *WafFirewallVersionResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallVersionResponseData := _WafFirewallVersionResponseData{}

	if err = json.Unmarshal(bytes, &varWafFirewallVersionResponseData); err == nil {
		*o = WafFirewallVersionResponseData(varWafFirewallVersionResponseData)
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

// NullableWafFirewallVersionResponseData is a helper abstraction for handling nullable waffirewallversionresponsedata types.
type NullableWafFirewallVersionResponseData struct {
	value *WafFirewallVersionResponseData
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallVersionResponseData) Get() *WafFirewallVersionResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallVersionResponseData) Set(val *WafFirewallVersionResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallVersionResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallVersionResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallVersionResponseData returns a pointer to a new instance of NullableWafFirewallVersionResponseData.
func NewNullableWafFirewallVersionResponseData(val *WafFirewallVersionResponseData) *NullableWafFirewallVersionResponseData {
	return &NullableWafFirewallVersionResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallVersionResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafFirewallVersionResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
