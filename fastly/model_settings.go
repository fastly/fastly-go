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

// Settings struct for Settings
type Settings struct {
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

type _Settings Settings

// NewSettings instantiates a new Settings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettings() *Settings {
	this := Settings{}
	var generalStaleIfError bool = false
	this.GeneralStaleIfError = &generalStaleIfError
	var generalStaleIfErrorTTL int32 = 43200
	this.GeneralStaleIfErrorTTL = &generalStaleIfErrorTTL
	return &this
}

// NewSettingsWithDefaults instantiates a new Settings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsWithDefaults() *Settings {
	this := Settings{}
	var generalStaleIfError bool = false
	this.GeneralStaleIfError = &generalStaleIfError
	var generalStaleIfErrorTTL int32 = 43200
	this.GeneralStaleIfErrorTTL = &generalStaleIfErrorTTL
	return &this
}

// GetGeneralDefaultHost returns the GeneralDefaultHost field value if set, zero value otherwise.
func (o *Settings) GetGeneralDefaultHost() string {
	if o == nil || o.GeneralDefaultHost == nil {
		var ret string
		return ret
	}
	return *o.GeneralDefaultHost
}

// GetGeneralDefaultHostOk returns a tuple with the GeneralDefaultHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetGeneralDefaultHostOk() (*string, bool) {
	if o == nil || o.GeneralDefaultHost == nil {
		return nil, false
	}
	return o.GeneralDefaultHost, true
}

// HasGeneralDefaultHost returns a boolean if a field has been set.
func (o *Settings) HasGeneralDefaultHost() bool {
	if o != nil && o.GeneralDefaultHost != nil {
		return true
	}

	return false
}

// SetGeneralDefaultHost gets a reference to the given string and assigns it to the GeneralDefaultHost field.
func (o *Settings) SetGeneralDefaultHost(v string) {
	o.GeneralDefaultHost = &v
}

// GetGeneralDefaultTTL returns the GeneralDefaultTTL field value if set, zero value otherwise.
func (o *Settings) GetGeneralDefaultTTL() int32 {
	if o == nil || o.GeneralDefaultTTL == nil {
		var ret int32
		return ret
	}
	return *o.GeneralDefaultTTL
}

// GetGeneralDefaultTTLOk returns a tuple with the GeneralDefaultTTL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetGeneralDefaultTTLOk() (*int32, bool) {
	if o == nil || o.GeneralDefaultTTL == nil {
		return nil, false
	}
	return o.GeneralDefaultTTL, true
}

// HasGeneralDefaultTTL returns a boolean if a field has been set.
func (o *Settings) HasGeneralDefaultTTL() bool {
	if o != nil && o.GeneralDefaultTTL != nil {
		return true
	}

	return false
}

// SetGeneralDefaultTTL gets a reference to the given int32 and assigns it to the GeneralDefaultTTL field.
func (o *Settings) SetGeneralDefaultTTL(v int32) {
	o.GeneralDefaultTTL = &v
}

// GetGeneralStaleIfError returns the GeneralStaleIfError field value if set, zero value otherwise.
func (o *Settings) GetGeneralStaleIfError() bool {
	if o == nil || o.GeneralStaleIfError == nil {
		var ret bool
		return ret
	}
	return *o.GeneralStaleIfError
}

// GetGeneralStaleIfErrorOk returns a tuple with the GeneralStaleIfError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetGeneralStaleIfErrorOk() (*bool, bool) {
	if o == nil || o.GeneralStaleIfError == nil {
		return nil, false
	}
	return o.GeneralStaleIfError, true
}

// HasGeneralStaleIfError returns a boolean if a field has been set.
func (o *Settings) HasGeneralStaleIfError() bool {
	if o != nil && o.GeneralStaleIfError != nil {
		return true
	}

	return false
}

// SetGeneralStaleIfError gets a reference to the given bool and assigns it to the GeneralStaleIfError field.
func (o *Settings) SetGeneralStaleIfError(v bool) {
	o.GeneralStaleIfError = &v
}

// GetGeneralStaleIfErrorTTL returns the GeneralStaleIfErrorTTL field value if set, zero value otherwise.
func (o *Settings) GetGeneralStaleIfErrorTTL() int32 {
	if o == nil || o.GeneralStaleIfErrorTTL == nil {
		var ret int32
		return ret
	}
	return *o.GeneralStaleIfErrorTTL
}

// GetGeneralStaleIfErrorTTLOk returns a tuple with the GeneralStaleIfErrorTTL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetGeneralStaleIfErrorTTLOk() (*int32, bool) {
	if o == nil || o.GeneralStaleIfErrorTTL == nil {
		return nil, false
	}
	return o.GeneralStaleIfErrorTTL, true
}

// HasGeneralStaleIfErrorTTL returns a boolean if a field has been set.
func (o *Settings) HasGeneralStaleIfErrorTTL() bool {
	if o != nil && o.GeneralStaleIfErrorTTL != nil {
		return true
	}

	return false
}

// SetGeneralStaleIfErrorTTL gets a reference to the given int32 and assigns it to the GeneralStaleIfErrorTTL field.
func (o *Settings) SetGeneralStaleIfErrorTTL(v int32) {
	o.GeneralStaleIfErrorTTL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Settings) MarshalJSON() ([]byte, error) {
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
func (o *Settings) UnmarshalJSON(bytes []byte) (err error) {
	varSettings := _Settings{}

	if err = json.Unmarshal(bytes, &varSettings); err == nil {
		*o = Settings(varSettings)
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

// NullableSettings is a helper abstraction for handling nullable settings types.
type NullableSettings struct {
	value *Settings
	isSet bool
}

// Get returns the value.
func (v NullableSettings) Get() *Settings {
	return v.value
}

// Set modifies the value.
func (v *NullableSettings) Set(val *Settings) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSettings) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSettings) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSettings returns a pointer to a new instance of NullableSettings.
func NewNullableSettings(val *Settings) *NullableSettings {
	return &NullableSettings{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
