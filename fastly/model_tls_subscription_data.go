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

// TlsSubscriptionData struct for TlsSubscriptionData
type TlsSubscriptionData struct {
	Type                 *TypeTlsSubscription             `json:"type,omitempty"`
	Attributes           *TlsSubscriptionDataAttributes   `json:"attributes,omitempty"`
	Relationships        *RelationshipsForTlsSubscription `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsSubscriptionData TlsSubscriptionData

// NewTlsSubscriptionData instantiates a new TlsSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsSubscriptionData() *TlsSubscriptionData {
	this := TlsSubscriptionData{}
	var type_ TypeTlsSubscription = TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION
	this.Type = &type_
	return &this
}

// NewTlsSubscriptionDataWithDefaults instantiates a new TlsSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsSubscriptionDataWithDefaults() *TlsSubscriptionData {
	this := TlsSubscriptionData{}
	var type_ TypeTlsSubscription = TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsSubscriptionData) GetType() TypeTlsSubscription {
	if o == nil || o.Type == nil {
		var ret TypeTlsSubscription
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsSubscriptionData) GetTypeOk() (*TypeTlsSubscription, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsSubscriptionData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsSubscription and assigns it to the Type field.
func (o *TlsSubscriptionData) SetType(v TypeTlsSubscription) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsSubscriptionData) GetAttributes() TlsSubscriptionDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TlsSubscriptionDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsSubscriptionData) GetAttributesOk() (*TlsSubscriptionDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsSubscriptionData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TlsSubscriptionDataAttributes and assigns it to the Attributes field.
func (o *TlsSubscriptionData) SetAttributes(v TlsSubscriptionDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TlsSubscriptionData) GetRelationships() RelationshipsForTlsSubscription {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTlsSubscription
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsSubscriptionData) GetRelationshipsOk() (*RelationshipsForTlsSubscription, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TlsSubscriptionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTlsSubscription and assigns it to the Relationships field.
func (o *TlsSubscriptionData) SetRelationships(v RelationshipsForTlsSubscription) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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
func (o *TlsSubscriptionData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsSubscriptionData := _TlsSubscriptionData{}

	if err = json.Unmarshal(bytes, &varTlsSubscriptionData); err == nil {
		*o = TlsSubscriptionData(varTlsSubscriptionData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsSubscriptionData is a helper abstraction for handling nullable tlssubscriptiondata types.
type NullableTlsSubscriptionData struct {
	value *TlsSubscriptionData
	isSet bool
}

// Get returns the value.
func (v NullableTlsSubscriptionData) Get() *TlsSubscriptionData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsSubscriptionData) Set(val *TlsSubscriptionData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsSubscriptionData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsSubscriptionData returns a pointer to a new instance of NullableTlsSubscriptionData.
func NewNullableTlsSubscriptionData(val *TlsSubscriptionData) *NullableTlsSubscriptionData {
	return &NullableTlsSubscriptionData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
