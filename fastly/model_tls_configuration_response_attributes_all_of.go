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

// TLSConfigurationResponseAttributesAllOf struct for TLSConfigurationResponseAttributesAllOf
type TLSConfigurationResponseAttributesAllOf struct {
	// Signifies whether or not Fastly will use this configuration as a default when creating a new [TLS Activation](https://www.fastly.com/documentation/reference/api/tls/custom-certs/activations/).
	Default *bool `json:"default,omitempty"`
	// HTTP protocols available on your configuration.
	HTTPProtocols []string `json:"http_protocols,omitempty"`
	// TLS protocols available on your configuration.
	TLSProtocols []string `json:"tls_protocols,omitempty"`
	// Signifies whether the configuration is used for Platform TLS or not.
	Bulk                 *bool `json:"bulk,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSConfigurationResponseAttributesAllOf TLSConfigurationResponseAttributesAllOf

// NewTLSConfigurationResponseAttributesAllOf instantiates a new TLSConfigurationResponseAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSConfigurationResponseAttributesAllOf() *TLSConfigurationResponseAttributesAllOf {
	this := TLSConfigurationResponseAttributesAllOf{}
	return &this
}

// NewTLSConfigurationResponseAttributesAllOfWithDefaults instantiates a new TLSConfigurationResponseAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSConfigurationResponseAttributesAllOfWithDefaults() *TLSConfigurationResponseAttributesAllOf {
	this := TLSConfigurationResponseAttributesAllOf{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *TLSConfigurationResponseAttributesAllOf) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationResponseAttributesAllOf) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *TLSConfigurationResponseAttributesAllOf) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *TLSConfigurationResponseAttributesAllOf) SetDefault(v bool) {
	o.Default = &v
}

// GetHTTPProtocols returns the HTTPProtocols field value if set, zero value otherwise.
func (o *TLSConfigurationResponseAttributesAllOf) GetHTTPProtocols() []string {
	if o == nil || o.HTTPProtocols == nil {
		var ret []string
		return ret
	}
	return o.HTTPProtocols
}

// GetHTTPProtocolsOk returns a tuple with the HTTPProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationResponseAttributesAllOf) GetHTTPProtocolsOk() ([]string, bool) {
	if o == nil || o.HTTPProtocols == nil {
		return nil, false
	}
	return o.HTTPProtocols, true
}

// HasHTTPProtocols returns a boolean if a field has been set.
func (o *TLSConfigurationResponseAttributesAllOf) HasHTTPProtocols() bool {
	if o != nil && o.HTTPProtocols != nil {
		return true
	}

	return false
}

// SetHTTPProtocols gets a reference to the given []string and assigns it to the HTTPProtocols field.
func (o *TLSConfigurationResponseAttributesAllOf) SetHTTPProtocols(v []string) {
	o.HTTPProtocols = v
}

// GetTLSProtocols returns the TLSProtocols field value if set, zero value otherwise.
func (o *TLSConfigurationResponseAttributesAllOf) GetTLSProtocols() []string {
	if o == nil || o.TLSProtocols == nil {
		var ret []string
		return ret
	}
	return o.TLSProtocols
}

// GetTLSProtocolsOk returns a tuple with the TLSProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationResponseAttributesAllOf) GetTLSProtocolsOk() ([]string, bool) {
	if o == nil || o.TLSProtocols == nil {
		return nil, false
	}
	return o.TLSProtocols, true
}

// HasTLSProtocols returns a boolean if a field has been set.
func (o *TLSConfigurationResponseAttributesAllOf) HasTLSProtocols() bool {
	if o != nil && o.TLSProtocols != nil {
		return true
	}

	return false
}

// SetTLSProtocols gets a reference to the given []string and assigns it to the TLSProtocols field.
func (o *TLSConfigurationResponseAttributesAllOf) SetTLSProtocols(v []string) {
	o.TLSProtocols = v
}

// GetBulk returns the Bulk field value if set, zero value otherwise.
func (o *TLSConfigurationResponseAttributesAllOf) GetBulk() bool {
	if o == nil || o.Bulk == nil {
		var ret bool
		return ret
	}
	return *o.Bulk
}

// GetBulkOk returns a tuple with the Bulk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationResponseAttributesAllOf) GetBulkOk() (*bool, bool) {
	if o == nil || o.Bulk == nil {
		return nil, false
	}
	return o.Bulk, true
}

// HasBulk returns a boolean if a field has been set.
func (o *TLSConfigurationResponseAttributesAllOf) HasBulk() bool {
	if o != nil && o.Bulk != nil {
		return true
	}

	return false
}

// SetBulk gets a reference to the given bool and assigns it to the Bulk field.
func (o *TLSConfigurationResponseAttributesAllOf) SetBulk(v bool) {
	o.Bulk = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSConfigurationResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.HTTPProtocols != nil {
		toSerialize["http_protocols"] = o.HTTPProtocols
	}
	if o.TLSProtocols != nil {
		toSerialize["tls_protocols"] = o.TLSProtocols
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
func (o *TLSConfigurationResponseAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTLSConfigurationResponseAttributesAllOf := _TLSConfigurationResponseAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varTLSConfigurationResponseAttributesAllOf); err == nil {
		*o = TLSConfigurationResponseAttributesAllOf(varTLSConfigurationResponseAttributesAllOf)
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

// NullableTLSConfigurationResponseAttributesAllOf is a helper abstraction for handling nullable tlsconfigurationresponseattributesallof types.
type NullableTLSConfigurationResponseAttributesAllOf struct {
	value *TLSConfigurationResponseAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTLSConfigurationResponseAttributesAllOf) Get() *TLSConfigurationResponseAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSConfigurationResponseAttributesAllOf) Set(val *TLSConfigurationResponseAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSConfigurationResponseAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSConfigurationResponseAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSConfigurationResponseAttributesAllOf returns a pointer to a new instance of NullableTLSConfigurationResponseAttributesAllOf.
func NewNullableTLSConfigurationResponseAttributesAllOf(val *TLSConfigurationResponseAttributesAllOf) *NullableTLSConfigurationResponseAttributesAllOf {
	return &NullableTLSConfigurationResponseAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSConfigurationResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSConfigurationResponseAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
