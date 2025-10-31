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

// TlsConfigurationResponseAttributesAllOf struct for TlsConfigurationResponseAttributesAllOf
type TlsConfigurationResponseAttributesAllOf struct {
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

type _TlsConfigurationResponseAttributesAllOf TlsConfigurationResponseAttributesAllOf

// NewTlsConfigurationResponseAttributesAllOf instantiates a new TlsConfigurationResponseAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsConfigurationResponseAttributesAllOf() *TlsConfigurationResponseAttributesAllOf {
	this := TlsConfigurationResponseAttributesAllOf{}
	return &this
}

// NewTlsConfigurationResponseAttributesAllOfWithDefaults instantiates a new TlsConfigurationResponseAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsConfigurationResponseAttributesAllOfWithDefaults() *TlsConfigurationResponseAttributesAllOf {
	this := TlsConfigurationResponseAttributesAllOf{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *TlsConfigurationResponseAttributesAllOf) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationResponseAttributesAllOf) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributesAllOf) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *TlsConfigurationResponseAttributesAllOf) SetDefault(v bool) {
	o.Default = &v
}

// GetHttpProtocols returns the HttpProtocols field value if set, zero value otherwise.
func (o *TlsConfigurationResponseAttributesAllOf) GetHttpProtocols() []string {
	if o == nil || o.HttpProtocols == nil {
		var ret []string
		return ret
	}
	return o.HttpProtocols
}

// GetHttpProtocolsOk returns a tuple with the HttpProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationResponseAttributesAllOf) GetHttpProtocolsOk() ([]string, bool) {
	if o == nil || o.HttpProtocols == nil {
		return nil, false
	}
	return o.HttpProtocols, true
}

// HasHttpProtocols returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributesAllOf) HasHttpProtocols() bool {
	if o != nil && o.HttpProtocols != nil {
		return true
	}

	return false
}

// SetHttpProtocols gets a reference to the given []string and assigns it to the HttpProtocols field.
func (o *TlsConfigurationResponseAttributesAllOf) SetHttpProtocols(v []string) {
	o.HttpProtocols = v
}

// GetTlsProtocols returns the TlsProtocols field value if set, zero value otherwise.
func (o *TlsConfigurationResponseAttributesAllOf) GetTlsProtocols() []string {
	if o == nil || o.TlsProtocols == nil {
		var ret []string
		return ret
	}
	return o.TlsProtocols
}

// GetTlsProtocolsOk returns a tuple with the TlsProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationResponseAttributesAllOf) GetTlsProtocolsOk() ([]string, bool) {
	if o == nil || o.TlsProtocols == nil {
		return nil, false
	}
	return o.TlsProtocols, true
}

// HasTlsProtocols returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributesAllOf) HasTlsProtocols() bool {
	if o != nil && o.TlsProtocols != nil {
		return true
	}

	return false
}

// SetTlsProtocols gets a reference to the given []string and assigns it to the TlsProtocols field.
func (o *TlsConfigurationResponseAttributesAllOf) SetTlsProtocols(v []string) {
	o.TlsProtocols = v
}

// GetBulk returns the Bulk field value if set, zero value otherwise.
func (o *TlsConfigurationResponseAttributesAllOf) GetBulk() bool {
	if o == nil || o.Bulk == nil {
		var ret bool
		return ret
	}
	return *o.Bulk
}

// GetBulkOk returns a tuple with the Bulk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationResponseAttributesAllOf) GetBulkOk() (*bool, bool) {
	if o == nil || o.Bulk == nil {
		return nil, false
	}
	return o.Bulk, true
}

// HasBulk returns a boolean if a field has been set.
func (o *TlsConfigurationResponseAttributesAllOf) HasBulk() bool {
	if o != nil && o.Bulk != nil {
		return true
	}

	return false
}

// SetBulk gets a reference to the given bool and assigns it to the Bulk field.
func (o *TlsConfigurationResponseAttributesAllOf) SetBulk(v bool) {
	o.Bulk = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsConfigurationResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *TlsConfigurationResponseAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTlsConfigurationResponseAttributesAllOf := _TlsConfigurationResponseAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varTlsConfigurationResponseAttributesAllOf); err == nil {
		*o = TlsConfigurationResponseAttributesAllOf(varTlsConfigurationResponseAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "default")
		delete(additionalProperties, "http_protocols")
		delete(additionalProperties, "tls_protocols")
		delete(additionalProperties, "bulk")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsConfigurationResponseAttributesAllOf is a helper abstraction for handling nullable tlsconfigurationresponseattributesallof types.
type NullableTlsConfigurationResponseAttributesAllOf struct {
	value *TlsConfigurationResponseAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTlsConfigurationResponseAttributesAllOf) Get() *TlsConfigurationResponseAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsConfigurationResponseAttributesAllOf) Set(val *TlsConfigurationResponseAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsConfigurationResponseAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsConfigurationResponseAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsConfigurationResponseAttributesAllOf returns a pointer to a new instance of NullableTlsConfigurationResponseAttributesAllOf.
func NewNullableTlsConfigurationResponseAttributesAllOf(val *TlsConfigurationResponseAttributesAllOf) *NullableTlsConfigurationResponseAttributesAllOf {
	return &NullableTlsConfigurationResponseAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsConfigurationResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsConfigurationResponseAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
