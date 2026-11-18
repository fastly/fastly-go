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

// TLSSubscriptionData struct for TLSSubscriptionData
type TLSSubscriptionData struct {
	Type                 *TypeTLSSubscription             `json:"type,omitempty"`
	Attributes           *TLSSubscriptionDataAttributes   `json:"attributes,omitempty"`
	Relationships        *RelationshipsForTLSSubscription `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSSubscriptionData TLSSubscriptionData

// NewTLSSubscriptionData instantiates a new TLSSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSSubscriptionData() *TLSSubscriptionData {
	this := TLSSubscriptionData{}
	var resourceType TypeTLSSubscription = TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION
	this.Type = &resourceType
	return &this
}

// NewTLSSubscriptionDataWithDefaults instantiates a new TLSSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSSubscriptionDataWithDefaults() *TLSSubscriptionData {
	this := TLSSubscriptionData{}
	var resourceType TypeTLSSubscription = TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSSubscriptionData) GetType() TypeTLSSubscription {
	if o == nil || o.Type == nil {
		var ret TypeTLSSubscription
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSSubscriptionData) GetTypeOk() (*TypeTLSSubscription, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSSubscriptionData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSSubscription and assigns it to the Type field.
func (o *TLSSubscriptionData) SetType(v TypeTLSSubscription) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSSubscriptionData) GetAttributes() TLSSubscriptionDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSSubscriptionDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSSubscriptionData) GetAttributesOk() (*TLSSubscriptionDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSSubscriptionData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSSubscriptionDataAttributes and assigns it to the Attributes field.
func (o *TLSSubscriptionData) SetAttributes(v TLSSubscriptionDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSSubscriptionData) GetRelationships() RelationshipsForTLSSubscription {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTLSSubscription
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSSubscriptionData) GetRelationshipsOk() (*RelationshipsForTLSSubscription, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSSubscriptionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTLSSubscription and assigns it to the Relationships field.
func (o *TLSSubscriptionData) SetRelationships(v RelationshipsForTLSSubscription) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSSubscriptionData) MarshalJSON() ([]byte, error) {
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
func (o *TLSSubscriptionData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSSubscriptionData := _TLSSubscriptionData{}

	if err = json.Unmarshal(bytes, &varTLSSubscriptionData); err == nil {
		*o = TLSSubscriptionData(varTLSSubscriptionData)
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

// NullableTLSSubscriptionData is a helper abstraction for handling nullable tlssubscriptiondata types.
type NullableTLSSubscriptionData struct {
	value *TLSSubscriptionData
	isSet bool
}

// Get returns the value.
func (v NullableTLSSubscriptionData) Get() *TLSSubscriptionData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSSubscriptionData) Set(val *TLSSubscriptionData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSSubscriptionData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSSubscriptionData returns a pointer to a new instance of NullableTLSSubscriptionData.
func NewNullableTLSSubscriptionData(val *TLSSubscriptionData) *NullableTLSSubscriptionData {
	return &NullableTLSSubscriptionData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
