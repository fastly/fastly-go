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

// VersionDetailSettings List of default settings for this service.
type VersionDetailSettings struct {
	// The default host name for the version.
	GeneralDefaultHost *string `json:"general.default_host,omitempty"`
	// The default time-to-live (TTL) for the version.
	GeneralDefaultTTL *int32 `json:"general.default_ttl,omitempty"`
	// Enables serving a stale object if there is an error.
	GeneralStaleIfError *bool `json:"general.stale_if_error,omitempty"`
	// The default time-to-live (TTL) for serving the stale object for the version.
	GeneralStaleIfErrorTTL *int32 `json:"general.stale_if_error_ttl,omitempty"`
	AdditionalProperties   map[string]any
}

type _VersionDetailSettings VersionDetailSettings

// NewVersionDetailSettings instantiates a new VersionDetailSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionDetailSettings() *VersionDetailSettings {
	this := VersionDetailSettings{}
	var generalStaleIfError bool = false
	this.GeneralStaleIfError = &generalStaleIfError
	var generalStaleIfErrorTTL int32 = 43200
	this.GeneralStaleIfErrorTTL = &generalStaleIfErrorTTL
	return &this
}

// NewVersionDetailSettingsWithDefaults instantiates a new VersionDetailSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionDetailSettingsWithDefaults() *VersionDetailSettings {
	this := VersionDetailSettings{}
	var generalStaleIfError bool = false
	this.GeneralStaleIfError = &generalStaleIfError
	var generalStaleIfErrorTTL int32 = 43200
	this.GeneralStaleIfErrorTTL = &generalStaleIfErrorTTL
	return &this
}

// GetGeneralDefaultHost returns the GeneralDefaultHost field value if set, zero value otherwise.
func (o *VersionDetailSettings) GetGeneralDefaultHost() string {
	if o == nil || o.GeneralDefaultHost == nil {
		var ret string
		return ret
	}
	return *o.GeneralDefaultHost
}

// GetGeneralDefaultHostOk returns a tuple with the GeneralDefaultHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetailSettings) GetGeneralDefaultHostOk() (*string, bool) {
	if o == nil || o.GeneralDefaultHost == nil {
		return nil, false
	}
	return o.GeneralDefaultHost, true
}

// HasGeneralDefaultHost returns a boolean if a field has been set.
func (o *VersionDetailSettings) HasGeneralDefaultHost() bool {
	if o != nil && o.GeneralDefaultHost != nil {
		return true
	}

	return false
}

// SetGeneralDefaultHost gets a reference to the given string and assigns it to the GeneralDefaultHost field.
func (o *VersionDetailSettings) SetGeneralDefaultHost(v string) {
	o.GeneralDefaultHost = &v
}

// GetGeneralDefaultTTL returns the GeneralDefaultTTL field value if set, zero value otherwise.
func (o *VersionDetailSettings) GetGeneralDefaultTTL() int32 {
	if o == nil || o.GeneralDefaultTTL == nil {
		var ret int32
		return ret
	}
	return *o.GeneralDefaultTTL
}

// GetGeneralDefaultTTLOk returns a tuple with the GeneralDefaultTTL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetailSettings) GetGeneralDefaultTTLOk() (*int32, bool) {
	if o == nil || o.GeneralDefaultTTL == nil {
		return nil, false
	}
	return o.GeneralDefaultTTL, true
}

// HasGeneralDefaultTTL returns a boolean if a field has been set.
func (o *VersionDetailSettings) HasGeneralDefaultTTL() bool {
	if o != nil && o.GeneralDefaultTTL != nil {
		return true
	}

	return false
}

// SetGeneralDefaultTTL gets a reference to the given int32 and assigns it to the GeneralDefaultTTL field.
func (o *VersionDetailSettings) SetGeneralDefaultTTL(v int32) {
	o.GeneralDefaultTTL = &v
}

// GetGeneralStaleIfError returns the GeneralStaleIfError field value if set, zero value otherwise.
func (o *VersionDetailSettings) GetGeneralStaleIfError() bool {
	if o == nil || o.GeneralStaleIfError == nil {
		var ret bool
		return ret
	}
	return *o.GeneralStaleIfError
}

// GetGeneralStaleIfErrorOk returns a tuple with the GeneralStaleIfError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetailSettings) GetGeneralStaleIfErrorOk() (*bool, bool) {
	if o == nil || o.GeneralStaleIfError == nil {
		return nil, false
	}
	return o.GeneralStaleIfError, true
}

// HasGeneralStaleIfError returns a boolean if a field has been set.
func (o *VersionDetailSettings) HasGeneralStaleIfError() bool {
	if o != nil && o.GeneralStaleIfError != nil {
		return true
	}

	return false
}

// SetGeneralStaleIfError gets a reference to the given bool and assigns it to the GeneralStaleIfError field.
func (o *VersionDetailSettings) SetGeneralStaleIfError(v bool) {
	o.GeneralStaleIfError = &v
}

// GetGeneralStaleIfErrorTTL returns the GeneralStaleIfErrorTTL field value if set, zero value otherwise.
func (o *VersionDetailSettings) GetGeneralStaleIfErrorTTL() int32 {
	if o == nil || o.GeneralStaleIfErrorTTL == nil {
		var ret int32
		return ret
	}
	return *o.GeneralStaleIfErrorTTL
}

// GetGeneralStaleIfErrorTTLOk returns a tuple with the GeneralStaleIfErrorTTL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetailSettings) GetGeneralStaleIfErrorTTLOk() (*int32, bool) {
	if o == nil || o.GeneralStaleIfErrorTTL == nil {
		return nil, false
	}
	return o.GeneralStaleIfErrorTTL, true
}

// HasGeneralStaleIfErrorTTL returns a boolean if a field has been set.
func (o *VersionDetailSettings) HasGeneralStaleIfErrorTTL() bool {
	if o != nil && o.GeneralStaleIfErrorTTL != nil {
		return true
	}

	return false
}

// SetGeneralStaleIfErrorTTL gets a reference to the given int32 and assigns it to the GeneralStaleIfErrorTTL field.
func (o *VersionDetailSettings) SetGeneralStaleIfErrorTTL(v int32) {
	o.GeneralStaleIfErrorTTL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o VersionDetailSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.GeneralDefaultHost != nil {
		toSerialize["general.default_host"] = o.GeneralDefaultHost
	}
	if o.GeneralDefaultTTL != nil {
		toSerialize["general.default_ttl"] = o.GeneralDefaultTTL
	}
	if o.GeneralStaleIfError != nil {
		toSerialize["general.stale_if_error"] = o.GeneralStaleIfError
	}
	if o.GeneralStaleIfErrorTTL != nil {
		toSerialize["general.stale_if_error_ttl"] = o.GeneralStaleIfErrorTTL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *VersionDetailSettings) UnmarshalJSON(bytes []byte) (err error) {
	varVersionDetailSettings := _VersionDetailSettings{}

	if err = json.Unmarshal(bytes, &varVersionDetailSettings); err == nil {
		*o = VersionDetailSettings(varVersionDetailSettings)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "general.default_host")
		delete(additionalProperties, "general.default_ttl")
		delete(additionalProperties, "general.stale_if_error")
		delete(additionalProperties, "general.stale_if_error_ttl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableVersionDetailSettings is a helper abstraction for handling nullable versiondetailsettings types.
type NullableVersionDetailSettings struct {
	value *VersionDetailSettings
	isSet bool
}

// Get returns the value.
func (v NullableVersionDetailSettings) Get() *VersionDetailSettings {
	return v.value
}

// Set modifies the value.
func (v *NullableVersionDetailSettings) Set(val *VersionDetailSettings) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableVersionDetailSettings) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableVersionDetailSettings) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableVersionDetailSettings returns a pointer to a new instance of NullableVersionDetailSettings.
func NewNullableVersionDetailSettings(val *VersionDetailSettings) *NullableVersionDetailSettings {
	return &NullableVersionDetailSettings{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableVersionDetailSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableVersionDetailSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
