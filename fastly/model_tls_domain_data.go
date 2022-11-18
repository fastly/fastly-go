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

// TLSDomainData struct for TLSDomainData
type TLSDomainData struct {
	// The domain name.
	ID *string `json:"id,omitempty"`
	Type *TypeTLSDomain `json:"type,omitempty"`
	Relationships *RelationshipsForTLSDomain `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSDomainData TLSDomainData

// NewTLSDomainData instantiates a new TLSDomainData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSDomainData() *TLSDomainData {
	this := TLSDomainData{}
	var resourceType TypeTLSDomain = TYPETLSDOMAIN_TLS_DOMAIN
	this.Type = &resourceType
	return &this
}

// NewTLSDomainDataWithDefaults instantiates a new TLSDomainData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSDomainDataWithDefaults() *TLSDomainData {
	this := TLSDomainData{}
	var resourceType TypeTLSDomain = TYPETLSDOMAIN_TLS_DOMAIN
	this.Type = &resourceType
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TLSDomainData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSDomainData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TLSDomainData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TLSDomainData) SetID(v string) {
	o.ID = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSDomainData) GetType() TypeTLSDomain {
	if o == nil || o.Type == nil {
		var ret TypeTLSDomain
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSDomainData) GetTypeOk() (*TypeTLSDomain, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSDomainData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSDomain and assigns it to the Type field.
func (o *TLSDomainData) SetType(v TypeTLSDomain) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSDomainData) GetRelationships() RelationshipsForTLSDomain {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTLSDomain
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSDomainData) GetRelationshipsOk() (*RelationshipsForTLSDomain, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSDomainData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTLSDomain and assigns it to the Relationships field.
func (o *TLSDomainData) SetRelationships(v RelationshipsForTLSDomain) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSDomainData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
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
func (o *TLSDomainData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSDomainData := _TLSDomainData{}

	if err = json.Unmarshal(bytes, &varTLSDomainData); err == nil {
		*o = TLSDomainData(varTLSDomainData)
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

// NullableTLSDomainData is a helper abstraction for handling nullable tlsdomaindata types. 
type NullableTLSDomainData struct {
	value *TLSDomainData
	isSet bool
}

// Get returns the value.
func (v NullableTLSDomainData) Get() *TLSDomainData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSDomainData) Set(val *TLSDomainData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSDomainData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSDomainData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSDomainData returns a pointer to a new instance of NullableTLSDomainData.
func NewNullableTLSDomainData(val *TLSDomainData) *NullableTLSDomainData {
	return &NullableTLSDomainData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSDomainData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSDomainData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
