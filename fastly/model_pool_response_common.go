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

// PoolResponseCommon struct for PoolResponseCommon
type PoolResponseCommon struct {
	// Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using `bereq.between_bytes_timeout`.
	BetweenBytesTimeout *string `json:"between_bytes_timeout,omitempty"`
	// How long to wait for a timeout in milliseconds.
	ConnectTimeout *string `json:"connect_timeout,omitempty"`
	// How long to wait for the first byte in milliseconds.
	FirstByteTimeout *string `json:"first_byte_timeout,omitempty"`
	// Maximum number of connections.
	MaxConnDefault *string `json:"max_conn_default,omitempty"`
	// Be strict on checking TLS certs.
	TLSCheckCert         NullableString `json:"tls_check_cert,omitempty"`
	ID                   *string        `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _PoolResponseCommon PoolResponseCommon

// NewPoolResponseCommon instantiates a new PoolResponseCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolResponseCommon() *PoolResponseCommon {
	this := PoolResponseCommon{}
	var maxConnDefault string = "200"
	this.MaxConnDefault = &maxConnDefault
	return &this
}

// NewPoolResponseCommonWithDefaults instantiates a new PoolResponseCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolResponseCommonWithDefaults() *PoolResponseCommon {
	this := PoolResponseCommon{}
	var maxConnDefault string = "200"
	this.MaxConnDefault = &maxConnDefault
	return &this
}

// GetBetweenBytesTimeout returns the BetweenBytesTimeout field value if set, zero value otherwise.
func (o *PoolResponseCommon) GetBetweenBytesTimeout() string {
	if o == nil || o.BetweenBytesTimeout == nil {
		var ret string
		return ret
	}
	return *o.BetweenBytesTimeout
}

// GetBetweenBytesTimeoutOk returns a tuple with the BetweenBytesTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponseCommon) GetBetweenBytesTimeoutOk() (*string, bool) {
	if o == nil || o.BetweenBytesTimeout == nil {
		return nil, false
	}
	return o.BetweenBytesTimeout, true
}

// HasBetweenBytesTimeout returns a boolean if a field has been set.
func (o *PoolResponseCommon) HasBetweenBytesTimeout() bool {
	if o != nil && o.BetweenBytesTimeout != nil {
		return true
	}

	return false
}

// SetBetweenBytesTimeout gets a reference to the given string and assigns it to the BetweenBytesTimeout field.
func (o *PoolResponseCommon) SetBetweenBytesTimeout(v string) {
	o.BetweenBytesTimeout = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *PoolResponseCommon) GetConnectTimeout() string {
	if o == nil || o.ConnectTimeout == nil {
		var ret string
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponseCommon) GetConnectTimeoutOk() (*string, bool) {
	if o == nil || o.ConnectTimeout == nil {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *PoolResponseCommon) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout != nil {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given string and assigns it to the ConnectTimeout field.
func (o *PoolResponseCommon) SetConnectTimeout(v string) {
	o.ConnectTimeout = &v
}

// GetFirstByteTimeout returns the FirstByteTimeout field value if set, zero value otherwise.
func (o *PoolResponseCommon) GetFirstByteTimeout() string {
	if o == nil || o.FirstByteTimeout == nil {
		var ret string
		return ret
	}
	return *o.FirstByteTimeout
}

// GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponseCommon) GetFirstByteTimeoutOk() (*string, bool) {
	if o == nil || o.FirstByteTimeout == nil {
		return nil, false
	}
	return o.FirstByteTimeout, true
}

// HasFirstByteTimeout returns a boolean if a field has been set.
func (o *PoolResponseCommon) HasFirstByteTimeout() bool {
	if o != nil && o.FirstByteTimeout != nil {
		return true
	}

	return false
}

// SetFirstByteTimeout gets a reference to the given string and assigns it to the FirstByteTimeout field.
func (o *PoolResponseCommon) SetFirstByteTimeout(v string) {
	o.FirstByteTimeout = &v
}

// GetMaxConnDefault returns the MaxConnDefault field value if set, zero value otherwise.
func (o *PoolResponseCommon) GetMaxConnDefault() string {
	if o == nil || o.MaxConnDefault == nil {
		var ret string
		return ret
	}
	return *o.MaxConnDefault
}

// GetMaxConnDefaultOk returns a tuple with the MaxConnDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponseCommon) GetMaxConnDefaultOk() (*string, bool) {
	if o == nil || o.MaxConnDefault == nil {
		return nil, false
	}
	return o.MaxConnDefault, true
}

// HasMaxConnDefault returns a boolean if a field has been set.
func (o *PoolResponseCommon) HasMaxConnDefault() bool {
	if o != nil && o.MaxConnDefault != nil {
		return true
	}

	return false
}

// SetMaxConnDefault gets a reference to the given string and assigns it to the MaxConnDefault field.
func (o *PoolResponseCommon) SetMaxConnDefault(v string) {
	o.MaxConnDefault = &v
}

// GetTLSCheckCert returns the TLSCheckCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponseCommon) GetTLSCheckCert() string {
	if o == nil || o.TLSCheckCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCheckCert.Get()
}

// GetTLSCheckCertOk returns a tuple with the TLSCheckCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponseCommon) GetTLSCheckCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TLSCheckCert.Get(), o.TLSCheckCert.IsSet()
}

// HasTLSCheckCert returns a boolean if a field has been set.
func (o *PoolResponseCommon) HasTLSCheckCert() bool {
	if o != nil && o.TLSCheckCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCheckCert gets a reference to the given NullableString and assigns it to the TLSCheckCert field.
func (o *PoolResponseCommon) SetTLSCheckCert(v string) {
	o.TLSCheckCert.Set(&v)
}

// SetTLSCheckCertNil sets the value for TLSCheckCert to be an explicit nil
func (o *PoolResponseCommon) SetTLSCheckCertNil() {
	o.TLSCheckCert.Set(nil)
}

// UnsetTLSCheckCert ensures that no value is present for TLSCheckCert, not even an explicit nil
func (o *PoolResponseCommon) UnsetTLSCheckCert() {
	o.TLSCheckCert.Unset()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *PoolResponseCommon) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponseCommon) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *PoolResponseCommon) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *PoolResponseCommon) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PoolResponseCommon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.BetweenBytesTimeout != nil {
		toSerialize["between_bytes_timeout"] = o.BetweenBytesTimeout
	}
	if o.ConnectTimeout != nil {
		toSerialize["connect_timeout"] = o.ConnectTimeout
	}
	if o.FirstByteTimeout != nil {
		toSerialize["first_byte_timeout"] = o.FirstByteTimeout
	}
	if o.MaxConnDefault != nil {
		toSerialize["max_conn_default"] = o.MaxConnDefault
	}
	if o.TLSCheckCert.IsSet() {
		toSerialize["tls_check_cert"] = o.TLSCheckCert.Get()
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PoolResponseCommon) UnmarshalJSON(bytes []byte) (err error) {
	varPoolResponseCommon := _PoolResponseCommon{}

	if err = json.Unmarshal(bytes, &varPoolResponseCommon); err == nil {
		*o = PoolResponseCommon(varPoolResponseCommon)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "between_bytes_timeout")
		delete(additionalProperties, "connect_timeout")
		delete(additionalProperties, "first_byte_timeout")
		delete(additionalProperties, "max_conn_default")
		delete(additionalProperties, "tls_check_cert")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePoolResponseCommon is a helper abstraction for handling nullable poolresponsecommon types.
type NullablePoolResponseCommon struct {
	value *PoolResponseCommon
	isSet bool
}

// Get returns the value.
func (v NullablePoolResponseCommon) Get() *PoolResponseCommon {
	return v.value
}

// Set modifies the value.
func (v *NullablePoolResponseCommon) Set(val *PoolResponseCommon) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePoolResponseCommon) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePoolResponseCommon) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePoolResponseCommon returns a pointer to a new instance of NullablePoolResponseCommon.
func NewNullablePoolResponseCommon(val *PoolResponseCommon) *NullablePoolResponseCommon {
	return &NullablePoolResponseCommon{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePoolResponseCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePoolResponseCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
