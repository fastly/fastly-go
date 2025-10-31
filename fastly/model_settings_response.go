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

// SettingsResponse struct for SettingsResponse
type SettingsResponse struct {
	// The default host name for the version.
	GeneralDefaultHost *string `json:"general.default_host,omitempty"`
	// The default time-to-live (TTL) for the version.
	GeneralDefaultTtl *int32 `json:"general.default_ttl,omitempty"`
	// Enables serving a stale object if there is an error.
	GeneralStaleIfError *bool `json:"general.stale_if_error,omitempty"`
	// The default time-to-live (TTL) for serving the stale object for the version.
	GeneralStaleIfErrorTtl *int32  `json:"general.stale_if_error_ttl,omitempty"`
	ServiceId              *string `json:"service_id,omitempty"`
	Version                *int32  `json:"version,omitempty"`
	AdditionalProperties   map[string]any
}

type _SettingsResponse SettingsResponse

// NewSettingsResponse instantiates a new SettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsResponse() *SettingsResponse {
	this := SettingsResponse{}
	var generalStaleIfError bool = false
	this.GeneralStaleIfError = &generalStaleIfError
	var generalStaleIfErrorTtl int32 = 43200
	this.GeneralStaleIfErrorTtl = &generalStaleIfErrorTtl
	return &this
}

// NewSettingsResponseWithDefaults instantiates a new SettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsResponseWithDefaults() *SettingsResponse {
	this := SettingsResponse{}
	var generalStaleIfError bool = false
	this.GeneralStaleIfError = &generalStaleIfError
	var generalStaleIfErrorTtl int32 = 43200
	this.GeneralStaleIfErrorTtl = &generalStaleIfErrorTtl
	return &this
}

// GetGeneralDefaultHost returns the GeneralDefaultHost field value if set, zero value otherwise.
func (o *SettingsResponse) GetGeneralDefaultHost() string {
	if o == nil || o.GeneralDefaultHost == nil {
		var ret string
		return ret
	}
	return *o.GeneralDefaultHost
}

// GetGeneralDefaultHostOk returns a tuple with the GeneralDefaultHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetGeneralDefaultHostOk() (*string, bool) {
	if o == nil || o.GeneralDefaultHost == nil {
		return nil, false
	}
	return o.GeneralDefaultHost, true
}

// HasGeneralDefaultHost returns a boolean if a field has been set.
func (o *SettingsResponse) HasGeneralDefaultHost() bool {
	if o != nil && o.GeneralDefaultHost != nil {
		return true
	}

	return false
}

// SetGeneralDefaultHost gets a reference to the given string and assigns it to the GeneralDefaultHost field.
func (o *SettingsResponse) SetGeneralDefaultHost(v string) {
	o.GeneralDefaultHost = &v
}

// GetGeneralDefaultTtl returns the GeneralDefaultTtl field value if set, zero value otherwise.
func (o *SettingsResponse) GetGeneralDefaultTtl() int32 {
	if o == nil || o.GeneralDefaultTtl == nil {
		var ret int32
		return ret
	}
	return *o.GeneralDefaultTtl
}

// GetGeneralDefaultTtlOk returns a tuple with the GeneralDefaultTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetGeneralDefaultTtlOk() (*int32, bool) {
	if o == nil || o.GeneralDefaultTtl == nil {
		return nil, false
	}
	return o.GeneralDefaultTtl, true
}

// HasGeneralDefaultTtl returns a boolean if a field has been set.
func (o *SettingsResponse) HasGeneralDefaultTtl() bool {
	if o != nil && o.GeneralDefaultTtl != nil {
		return true
	}

	return false
}

// SetGeneralDefaultTtl gets a reference to the given int32 and assigns it to the GeneralDefaultTtl field.
func (o *SettingsResponse) SetGeneralDefaultTtl(v int32) {
	o.GeneralDefaultTtl = &v
}

// GetGeneralStaleIfError returns the GeneralStaleIfError field value if set, zero value otherwise.
func (o *SettingsResponse) GetGeneralStaleIfError() bool {
	if o == nil || o.GeneralStaleIfError == nil {
		var ret bool
		return ret
	}
	return *o.GeneralStaleIfError
}

// GetGeneralStaleIfErrorOk returns a tuple with the GeneralStaleIfError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetGeneralStaleIfErrorOk() (*bool, bool) {
	if o == nil || o.GeneralStaleIfError == nil {
		return nil, false
	}
	return o.GeneralStaleIfError, true
}

// HasGeneralStaleIfError returns a boolean if a field has been set.
func (o *SettingsResponse) HasGeneralStaleIfError() bool {
	if o != nil && o.GeneralStaleIfError != nil {
		return true
	}

	return false
}

// SetGeneralStaleIfError gets a reference to the given bool and assigns it to the GeneralStaleIfError field.
func (o *SettingsResponse) SetGeneralStaleIfError(v bool) {
	o.GeneralStaleIfError = &v
}

// GetGeneralStaleIfErrorTtl returns the GeneralStaleIfErrorTtl field value if set, zero value otherwise.
func (o *SettingsResponse) GetGeneralStaleIfErrorTtl() int32 {
	if o == nil || o.GeneralStaleIfErrorTtl == nil {
		var ret int32
		return ret
	}
	return *o.GeneralStaleIfErrorTtl
}

// GetGeneralStaleIfErrorTtlOk returns a tuple with the GeneralStaleIfErrorTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetGeneralStaleIfErrorTtlOk() (*int32, bool) {
	if o == nil || o.GeneralStaleIfErrorTtl == nil {
		return nil, false
	}
	return o.GeneralStaleIfErrorTtl, true
}

// HasGeneralStaleIfErrorTtl returns a boolean if a field has been set.
func (o *SettingsResponse) HasGeneralStaleIfErrorTtl() bool {
	if o != nil && o.GeneralStaleIfErrorTtl != nil {
		return true
	}

	return false
}

// SetGeneralStaleIfErrorTtl gets a reference to the given int32 and assigns it to the GeneralStaleIfErrorTtl field.
func (o *SettingsResponse) SetGeneralStaleIfErrorTtl(v int32) {
	o.GeneralStaleIfErrorTtl = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *SettingsResponse) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *SettingsResponse) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *SettingsResponse) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SettingsResponse) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsResponse) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SettingsResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SettingsResponse) SetVersion(v int32) {
	o.Version = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.GeneralDefaultHost != nil {
		toSerialize["general.default_host"] = o.GeneralDefaultHost
	}
	if o.GeneralDefaultTtl != nil {
		toSerialize["general.default_ttl"] = o.GeneralDefaultTtl
	}
	if o.GeneralStaleIfError != nil {
		toSerialize["general.stale_if_error"] = o.GeneralStaleIfError
	}
	if o.GeneralStaleIfErrorTtl != nil {
		toSerialize["general.stale_if_error_ttl"] = o.GeneralStaleIfErrorTtl
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *SettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSettingsResponse := _SettingsResponse{}

	if err = json.Unmarshal(bytes, &varSettingsResponse); err == nil {
		*o = SettingsResponse(varSettingsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "general.default_host")
		delete(additionalProperties, "general.default_ttl")
		delete(additionalProperties, "general.stale_if_error")
		delete(additionalProperties, "general.stale_if_error_ttl")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSettingsResponse is a helper abstraction for handling nullable settingsresponse types.
type NullableSettingsResponse struct {
	value *SettingsResponse
	isSet bool
}

// Get returns the value.
func (v NullableSettingsResponse) Get() *SettingsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableSettingsResponse) Set(val *SettingsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSettingsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSettingsResponse returns a pointer to a new instance of NullableSettingsResponse.
func NewNullableSettingsResponse(val *SettingsResponse) *NullableSettingsResponse {
	return &NullableSettingsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
