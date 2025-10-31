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
	// Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, for Delivery services, the response received so far will be considered complete and the fetch will end. For Compute services, timeout expiration is treated as a failure of the backend connection, and an error is generated. May be set at runtime using `bereq.between_bytes_timeout`.
	BetweenBytesTimeout *string `json:"between_bytes_timeout,omitempty"`
	// How long to wait for a timeout in milliseconds.
	ConnectTimeout *string `json:"connect_timeout,omitempty"`
	// How long to wait for the first byte in milliseconds.
	FirstByteTimeout *string `json:"first_byte_timeout,omitempty"`
	// Maximum number of connections.
	MaxConnDefault *string `json:"max_conn_default,omitempty"`
	// Be strict on checking TLS certs.
	TlsCheckCert         NullableString `json:"tls_check_cert,omitempty"`
	Id                   *string        `json:"id,omitempty"`
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

// GetTlsCheckCert returns the TlsCheckCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolResponseCommon) GetTlsCheckCert() string {
	if o == nil || o.TlsCheckCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCheckCert.Get()
}

// GetTlsCheckCertOk returns a tuple with the TlsCheckCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolResponseCommon) GetTlsCheckCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCheckCert.Get(), o.TlsCheckCert.IsSet()
}

// HasTlsCheckCert returns a boolean if a field has been set.
func (o *PoolResponseCommon) HasTlsCheckCert() bool {
	if o != nil && o.TlsCheckCert.IsSet() {
		return true
	}

	return false
}

// SetTlsCheckCert gets a reference to the given NullableString and assigns it to the TlsCheckCert field.
func (o *PoolResponseCommon) SetTlsCheckCert(v string) {
	o.TlsCheckCert.Set(&v)
}

// SetTlsCheckCertNil sets the value for TlsCheckCert to be an explicit nil
func (o *PoolResponseCommon) SetTlsCheckCertNil() {
	o.TlsCheckCert.Set(nil)
}

// UnsetTlsCheckCert ensures that no value is present for TlsCheckCert, not even an explicit nil
func (o *PoolResponseCommon) UnsetTlsCheckCert() {
	o.TlsCheckCert.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoolResponseCommon) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolResponseCommon) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoolResponseCommon) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PoolResponseCommon) SetId(v string) {
	o.Id = &v
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
	if o.TlsCheckCert.IsSet() {
		toSerialize["tls_check_cert"] = o.TlsCheckCert.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
