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
	"time"
)

// TLSPrivateKeyResponseAttributes struct for TLSPrivateKeyResponseAttributes
type TLSPrivateKeyResponseAttributes struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// A customizable name for your private key.
	Name *string `json:"name,omitempty"`
	// The key length used to generate the private key.
	KeyLength *int32 `json:"key_length,omitempty"`
	// The algorithm used to generate the private key. Must be `RSA`.
	KeyType *string `json:"key_type,omitempty"`
	// A recommendation from Fastly to replace this private key and all associated certificates.
	Replace *bool `json:"replace,omitempty"`
	// Useful for safely identifying the key.
	PublicKeySha1 *string `json:"public_key_sha1,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSPrivateKeyResponseAttributes TLSPrivateKeyResponseAttributes

// NewTLSPrivateKeyResponseAttributes instantiates a new TLSPrivateKeyResponseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSPrivateKeyResponseAttributes() *TLSPrivateKeyResponseAttributes {
	this := TLSPrivateKeyResponseAttributes{}
	return &this
}

// NewTLSPrivateKeyResponseAttributesWithDefaults instantiates a new TLSPrivateKeyResponseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSPrivateKeyResponseAttributesWithDefaults() *TLSPrivateKeyResponseAttributes {
	this := TLSPrivateKeyResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSPrivateKeyResponseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSPrivateKeyResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *TLSPrivateKeyResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *TLSPrivateKeyResponseAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *TLSPrivateKeyResponseAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSPrivateKeyResponseAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSPrivateKeyResponseAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributes) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *TLSPrivateKeyResponseAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *TLSPrivateKeyResponseAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *TLSPrivateKeyResponseAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSPrivateKeyResponseAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSPrivateKeyResponseAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *TLSPrivateKeyResponseAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *TLSPrivateKeyResponseAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *TLSPrivateKeyResponseAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TLSPrivateKeyResponseAttributes) SetName(v string) {
	o.Name = &v
}

// GetKeyLength returns the KeyLength field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseAttributes) GetKeyLength() int32 {
	if o == nil || o.KeyLength == nil {
		var ret int32
		return ret
	}
	return *o.KeyLength
}

// GetKeyLengthOk returns a tuple with the KeyLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseAttributes) GetKeyLengthOk() (*int32, bool) {
	if o == nil || o.KeyLength == nil {
		return nil, false
	}
	return o.KeyLength, true
}

// HasKeyLength returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributes) HasKeyLength() bool {
	if o != nil && o.KeyLength != nil {
		return true
	}

	return false
}

// SetKeyLength gets a reference to the given int32 and assigns it to the KeyLength field.
func (o *TLSPrivateKeyResponseAttributes) SetKeyLength(v int32) {
	o.KeyLength = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseAttributes) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseAttributes) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributes) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *TLSPrivateKeyResponseAttributes) SetKeyType(v string) {
	o.KeyType = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseAttributes) GetReplace() bool {
	if o == nil || o.Replace == nil {
		var ret bool
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseAttributes) GetReplaceOk() (*bool, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributes) HasReplace() bool {
	if o != nil && o.Replace != nil {
		return true
	}

	return false
}

// SetReplace gets a reference to the given bool and assigns it to the Replace field.
func (o *TLSPrivateKeyResponseAttributes) SetReplace(v bool) {
	o.Replace = &v
}

// GetPublicKeySha1 returns the PublicKeySha1 field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseAttributes) GetPublicKeySha1() string {
	if o == nil || o.PublicKeySha1 == nil {
		var ret string
		return ret
	}
	return *o.PublicKeySha1
}

// GetPublicKeySha1Ok returns a tuple with the PublicKeySha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseAttributes) GetPublicKeySha1Ok() (*string, bool) {
	if o == nil || o.PublicKeySha1 == nil {
		return nil, false
	}
	return o.PublicKeySha1, true
}

// HasPublicKeySha1 returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributes) HasPublicKeySha1() bool {
	if o != nil && o.PublicKeySha1 != nil {
		return true
	}

	return false
}

// SetPublicKeySha1 gets a reference to the given string and assigns it to the PublicKeySha1 field.
func (o *TLSPrivateKeyResponseAttributes) SetPublicKeySha1(v string) {
	o.PublicKeySha1 = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSPrivateKeyResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.KeyLength != nil {
		toSerialize["key_length"] = o.KeyLength
	}
	if o.KeyType != nil {
		toSerialize["key_type"] = o.KeyType
	}
	if o.Replace != nil {
		toSerialize["replace"] = o.Replace
	}
	if o.PublicKeySha1 != nil {
		toSerialize["public_key_sha1"] = o.PublicKeySha1
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *TLSPrivateKeyResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTLSPrivateKeyResponseAttributes := _TLSPrivateKeyResponseAttributes{}

	if err = json.Unmarshal(bytes, &varTLSPrivateKeyResponseAttributes); err == nil {
		*o = TLSPrivateKeyResponseAttributes(varTLSPrivateKeyResponseAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "key_length")
		delete(additionalProperties, "key_type")
		delete(additionalProperties, "replace")
		delete(additionalProperties, "public_key_sha1")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSPrivateKeyResponseAttributes is a helper abstraction for handling nullable tlsprivatekeyresponseattributes types. 
type NullableTLSPrivateKeyResponseAttributes struct {
	value *TLSPrivateKeyResponseAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTLSPrivateKeyResponseAttributes) Get() *TLSPrivateKeyResponseAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSPrivateKeyResponseAttributes) Set(val *TLSPrivateKeyResponseAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSPrivateKeyResponseAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSPrivateKeyResponseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSPrivateKeyResponseAttributes returns a pointer to a new instance of NullableTLSPrivateKeyResponseAttributes.
func NewNullableTLSPrivateKeyResponseAttributes(val *TLSPrivateKeyResponseAttributes) *NullableTLSPrivateKeyResponseAttributes {
	return &NullableTLSPrivateKeyResponseAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSPrivateKeyResponseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSPrivateKeyResponseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
