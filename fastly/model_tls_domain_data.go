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

// TlsDomainData struct for TlsDomainData
type TlsDomainData struct {
	// The domain name.
	Id                   *string                    `json:"id,omitempty"`
	Type                 *TypeTlsDomain             `json:"type,omitempty"`
	Relationships        *RelationshipsForTlsDomain `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsDomainData TlsDomainData

// NewTlsDomainData instantiates a new TlsDomainData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsDomainData() *TlsDomainData {
	this := TlsDomainData{}
	var type_ TypeTlsDomain = TYPETLSDOMAIN_TLS_DOMAIN
	this.Type = &type_
	return &this
}

// NewTlsDomainDataWithDefaults instantiates a new TlsDomainData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsDomainDataWithDefaults() *TlsDomainData {
	this := TlsDomainData{}
	var type_ TypeTlsDomain = TYPETLSDOMAIN_TLS_DOMAIN
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TlsDomainData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDomainData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TlsDomainData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TlsDomainData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsDomainData) GetType() TypeTlsDomain {
	if o == nil || o.Type == nil {
		var ret TypeTlsDomain
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDomainData) GetTypeOk() (*TypeTlsDomain, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsDomainData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsDomain and assigns it to the Type field.
func (o *TlsDomainData) SetType(v TypeTlsDomain) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TlsDomainData) GetRelationships() RelationshipsForTlsDomain {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTlsDomain
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDomainData) GetRelationshipsOk() (*RelationshipsForTlsDomain, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TlsDomainData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTlsDomain and assigns it to the Relationships field.
func (o *TlsDomainData) SetRelationships(v RelationshipsForTlsDomain) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsDomainData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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
func (o *TlsDomainData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsDomainData := _TlsDomainData{}

	if err = json.Unmarshal(bytes, &varTlsDomainData); err == nil {
		*o = TlsDomainData(varTlsDomainData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsDomainData is a helper abstraction for handling nullable tlsdomaindata types.
type NullableTlsDomainData struct {
	value *TlsDomainData
	isSet bool
}

// Get returns the value.
func (v NullableTlsDomainData) Get() *TlsDomainData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsDomainData) Set(val *TlsDomainData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsDomainData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsDomainData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsDomainData returns a pointer to a new instance of NullableTlsDomainData.
func NewNullableTlsDomainData(val *TlsDomainData) *NullableTlsDomainData {
	return &NullableTlsDomainData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsDomainData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsDomainData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
