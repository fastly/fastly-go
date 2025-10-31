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

// RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate struct for RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate
type RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate struct {
	Data                 *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificateData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate

// NewRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate instantiates a new RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate() *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate {
	this := RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate{}
	return &this
}

// NewRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificateWithDefaults instantiates a new RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificateWithDefaults() *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate {
	this := RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) GetData() RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificateData {
	if o == nil || o.Data == nil {
		var ret RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificateData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) GetDataOk() (*RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificateData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificateData and assigns it to the Data field.
func (o *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) SetData(v RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificateData) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate := _RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate{}

	if err = json.Unmarshal(bytes, &varRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate); err == nil {
		*o = RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate(varRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate is a helper abstraction for handling nullable relationshipdefaultecdsatlscertificatedefaultecdsacertificate types.
type NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate struct {
	value *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) Get() *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) Set(val *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate returns a pointer to a new instance of NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate.
func NewNullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate(val *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) *NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate {
	return &NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
