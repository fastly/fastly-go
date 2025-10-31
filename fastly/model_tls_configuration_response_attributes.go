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
	"time"
)

// TlsConfigurationResponseAttributes struct for TlsConfigurationResponseAttributes
type TlsConfigurationResponseAttributes struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// Signifies whether or not Fastly will use this configuration as a default when creating a new [TLS Activation](https://www.fastly.com/documentation/reference/api/tls/custom-certs/activations/).
	Default *bool `json:"default,omitempty"`
	// HTTP protocols available on your configuration.
	HttpProtocols []string `json:"http_protocols,omitempty"`
	// TLS protocols available on your configuration.
	TlsProtocols []string `json:"tls_protocols,omitempty"`
	// Signifies whether the configuration is used for Platform TLS or not.
	Bulk                 *bool `json:"bulk,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsConfigurationResponseAttributes TlsConfigurationResponseAttributes

// NewTlsConfigurationResponseAttributes instantiates a new TlsConfigurationResponseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsConfigurationResponseAttributes() *TlsConfigurationResponseAttributes {
	this := TlsConfigurationResponseAttributes{}
	return &this
}

// NewTlsConfigurationResponseAttributesWithDefaults instantiates a new TlsConfigurationResponseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsConfigurationResponseAttributesWithDefaults() *TlsConfigurationResponseAttributes {
	this := TlsConfigurationResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsConfigurationResponseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsConfigurationResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *TlsConfigurationResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *TlsConfigurationResponseAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *TlsConfigurationResponseAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsConfigurationResponseAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsConfigurationResponseAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributes) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *TlsConfigurationResponseAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *TlsConfigurationResponseAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *TlsConfigurationResponseAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsConfigurationResponseAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsConfigurationResponseAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *TlsConfigurationResponseAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *TlsConfigurationResponseAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *TlsConfigurationResponseAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *TlsConfigurationResponseAttributes) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationResponseAttributes) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributes) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *TlsConfigurationResponseAttributes) SetDefault(v bool) {
	o.Default = &v
}

// GetHttpProtocols returns the HttpProtocols field value if set, zero value otherwise.
func (o *TlsConfigurationResponseAttributes) GetHttpProtocols() []string {
	if o == nil || o.HttpProtocols == nil {
		var ret []string
		return ret
	}
	return o.HttpProtocols
}

// GetHttpProtocolsOk returns a tuple with the HttpProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationResponseAttributes) GetHttpProtocolsOk() ([]string, bool) {
	if o == nil || o.HttpProtocols == nil {
		return nil, false
	}
	return o.HttpProtocols, true
}

// HasHttpProtocols returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributes) HasHttpProtocols() bool {
	if o != nil && o.HttpProtocols != nil {
		return true
	}

	return false
}

// SetHttpProtocols gets a reference to the given []string and assigns it to the HttpProtocols field.
func (o *TlsConfigurationResponseAttributes) SetHttpProtocols(v []string) {
	o.HttpProtocols = v
}

// GetTlsProtocols returns the TlsProtocols field value if set, zero value otherwise.
func (o *TlsConfigurationResponseAttributes) GetTlsProtocols() []string {
	if o == nil || o.TlsProtocols == nil {
		var ret []string
		return ret
	}
	return o.TlsProtocols
}

// GetTlsProtocolsOk returns a tuple with the TlsProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationResponseAttributes) GetTlsProtocolsOk() ([]string, bool) {
	if o == nil || o.TlsProtocols == nil {
		return nil, false
	}
	return o.TlsProtocols, true
}

// HasTlsProtocols returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributes) HasTlsProtocols() bool {
	if o != nil && o.TlsProtocols != nil {
		return true
	}

	return false
}

// SetTlsProtocols gets a reference to the given []string and assigns it to the TlsProtocols field.
func (o *TlsConfigurationResponseAttributes) SetTlsProtocols(v []string) {
	o.TlsProtocols = v
}

// GetBulk returns the Bulk field value if set, zero value otherwise.
func (o *TlsConfigurationResponseAttributes) GetBulk() bool {
	if o == nil || o.Bulk == nil {
		var ret bool
		return ret
	}
	return *o.Bulk
}

// GetBulkOk returns a tuple with the Bulk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationResponseAttributes) GetBulkOk() (*bool, bool) {
	if o == nil || o.Bulk == nil {
		return nil, false
	}
	return o.Bulk, true
}

// HasBulk returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributes) HasBulk() bool {
	if o != nil && o.Bulk != nil {
		return true
	}

	return false
}

// SetBulk gets a reference to the given bool and assigns it to the Bulk field.
func (o *TlsConfigurationResponseAttributes) SetBulk(v bool) {
	o.Bulk = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsConfigurationResponseAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.HttpProtocols != nil {
		toSerialize["http_protocols"] = o.HttpProtocols
	}
	if o.TlsProtocols != nil {
		toSerialize["tls_protocols"] = o.TlsProtocols
	}
	if o.Bulk != nil {
		toSerialize["bulk"] = o.Bulk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsConfigurationResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTlsConfigurationResponseAttributes := _TlsConfigurationResponseAttributes{}

	if err = json.Unmarshal(bytes, &varTlsConfigurationResponseAttributes); err == nil {
		*o = TlsConfigurationResponseAttributes(varTlsConfigurationResponseAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "default")
		delete(additionalProperties, "http_protocols")
		delete(additionalProperties, "tls_protocols")
		delete(additionalProperties, "bulk")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsConfigurationResponseAttributes is a helper abstraction for handling nullable tlsconfigurationresponseattributes types.
type NullableTlsConfigurationResponseAttributes struct {
	value *TlsConfigurationResponseAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTlsConfigurationResponseAttributes) Get() *TlsConfigurationResponseAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsConfigurationResponseAttributes) Set(val *TlsConfigurationResponseAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsConfigurationResponseAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsConfigurationResponseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsConfigurationResponseAttributes returns a pointer to a new instance of NullableTlsConfigurationResponseAttributes.
func NewNullableTlsConfigurationResponseAttributes(val *TlsConfigurationResponseAttributes) *NullableTlsConfigurationResponseAttributes {
	return &NullableTlsConfigurationResponseAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsConfigurationResponseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsConfigurationResponseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
