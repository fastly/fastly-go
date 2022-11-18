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

// StarData struct for StarData
type StarData struct {
	Type *TypeStar `json:"type,omitempty"`
	Relationships *RelationshipsForStar `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _StarData StarData

// NewStarData instantiates a new StarData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStarData() *StarData {
	this := StarData{}
	var resourceType TypeStar = TYPESTAR_STAR
	this.Type = &resourceType
	return &this
}

// NewStarDataWithDefaults instantiates a new StarData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStarDataWithDefaults() *StarData {
	this := StarData{}
	var resourceType TypeStar = TYPESTAR_STAR
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StarData) GetType() TypeStar {
	if o == nil || o.Type == nil {
		var ret TypeStar
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StarData) GetTypeOk() (*TypeStar, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StarData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeStar and assigns it to the Type field.
func (o *StarData) SetType(v TypeStar) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StarData) GetRelationships() RelationshipsForStar {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForStar
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StarData) GetRelationshipsOk() (*RelationshipsForStar, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StarData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForStar and assigns it to the Relationships field.
func (o *StarData) SetRelationships(v RelationshipsForStar) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o StarData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *StarData) UnmarshalJSON(bytes []byte) (err error) {
	varStarData := _StarData{}

	if err = json.Unmarshal(bytes, &varStarData); err == nil {
		*o = StarData(varStarData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableStarData is a helper abstraction for handling nullable stardata types. 
type NullableStarData struct {
	value *StarData
	isSet bool
}

// Get returns the value.
func (v NullableStarData) Get() *StarData {
	return v.value
}

// Set modifies the value.
func (v *NullableStarData) Set(val *StarData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableStarData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableStarData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableStarData returns a pointer to a new instance of NullableStarData.
func NewNullableStarData(val *StarData) *NullableStarData {
	return &NullableStarData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableStarData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableStarData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
