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

// SchemasUserResponseReadOnly struct for SchemasUserResponseReadOnly
type SchemasUserResponseReadOnly struct {
	Id *string `json:"id,omitempty"`
	// The alphanumeric string identifying a email login.
	EmailHash            *string `json:"email_hash,omitempty"`
	CustomerId           *string `json:"customer_id,omitempty"`
	AdditionalProperties map[string]any
}

type _SchemasUserResponseReadOnly SchemasUserResponseReadOnly

// NewSchemasUserResponseReadOnly instantiates a new SchemasUserResponseReadOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUserResponseReadOnly() *SchemasUserResponseReadOnly {
	this := SchemasUserResponseReadOnly{}
	return &this
}

// NewSchemasUserResponseReadOnlyWithDefaults instantiates a new SchemasUserResponseReadOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUserResponseReadOnlyWithDefaults() *SchemasUserResponseReadOnly {
	this := SchemasUserResponseReadOnly{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SchemasUserResponseReadOnly) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponseReadOnly) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SchemasUserResponseReadOnly) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SchemasUserResponseReadOnly) SetId(v string) {
	o.Id = &v
}

// GetEmailHash returns the EmailHash field value if set, zero value otherwise.
func (o *SchemasUserResponseReadOnly) GetEmailHash() string {
	if o == nil || o.EmailHash == nil {
		var ret string
		return ret
	}
	return *o.EmailHash
}

// GetEmailHashOk returns a tuple with the EmailHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponseReadOnly) GetEmailHashOk() (*string, bool) {
	if o == nil || o.EmailHash == nil {
		return nil, false
	}
	return o.EmailHash, true
}

// HasEmailHash returns a boolean if a field has been set.
func (o *SchemasUserResponseReadOnly) HasEmailHash() bool {
	if o != nil && o.EmailHash != nil {
		return true
	}

	return false
}

// SetEmailHash gets a reference to the given string and assigns it to the EmailHash field.
func (o *SchemasUserResponseReadOnly) SetEmailHash(v string) {
	o.EmailHash = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *SchemasUserResponseReadOnly) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponseReadOnly) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *SchemasUserResponseReadOnly) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *SchemasUserResponseReadOnly) SetCustomerId(v string) {
	o.CustomerId = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SchemasUserResponseReadOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.EmailHash != nil {
		toSerialize["email_hash"] = o.EmailHash
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *SchemasUserResponseReadOnly) UnmarshalJSON(bytes []byte) (err error) {
	varSchemasUserResponseReadOnly := _SchemasUserResponseReadOnly{}

	if err = json.Unmarshal(bytes, &varSchemasUserResponseReadOnly); err == nil {
		*o = SchemasUserResponseReadOnly(varSchemasUserResponseReadOnly)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "email_hash")
		delete(additionalProperties, "customer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSchemasUserResponseReadOnly is a helper abstraction for handling nullable schemasuserresponsereadonly types.
type NullableSchemasUserResponseReadOnly struct {
	value *SchemasUserResponseReadOnly
	isSet bool
}

// Get returns the value.
func (v NullableSchemasUserResponseReadOnly) Get() *SchemasUserResponseReadOnly {
	return v.value
}

// Set modifies the value.
func (v *NullableSchemasUserResponseReadOnly) Set(val *SchemasUserResponseReadOnly) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSchemasUserResponseReadOnly) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSchemasUserResponseReadOnly) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSchemasUserResponseReadOnly returns a pointer to a new instance of NullableSchemasUserResponseReadOnly.
func NewNullableSchemasUserResponseReadOnly(val *SchemasUserResponseReadOnly) *NullableSchemasUserResponseReadOnly {
	return &NullableSchemasUserResponseReadOnly{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSchemasUserResponseReadOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableSchemasUserResponseReadOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
