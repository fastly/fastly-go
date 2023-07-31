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
)

// RequestSettingsResponseAllOf struct for RequestSettingsResponseAllOf
type RequestSettingsResponseAllOf struct {
	// Disable collapsed forwarding, so you don't wait for other objects to origin.
	BypassBusyWait *string `json:"bypass_busy_wait,omitempty"`
	// Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable.
	ForceMiss *string `json:"force_miss,omitempty"`
	// Forces the request use SSL (redirects a non-SSL to SSL).
	ForceSsl *string `json:"force_ssl,omitempty"`
	// Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers.
	GeoHeaders *string `json:"geo_headers,omitempty"`
	// How old an object is allowed to be to serve stale-if-error or stale-while-revalidate.
	MaxStaleAge *string `json:"max_stale_age,omitempty"`
	// Injects the X-Timer info into the request for viewing origin fetch durations.
	TimerSupport *string `json:"timer_support,omitempty"`
	AdditionalProperties map[string]any
}

type _RequestSettingsResponseAllOf RequestSettingsResponseAllOf

// NewRequestSettingsResponseAllOf instantiates a new RequestSettingsResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSettingsResponseAllOf() *RequestSettingsResponseAllOf {
	this := RequestSettingsResponseAllOf{}
	return &this
}

// NewRequestSettingsResponseAllOfWithDefaults instantiates a new RequestSettingsResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSettingsResponseAllOfWithDefaults() *RequestSettingsResponseAllOf {
	this := RequestSettingsResponseAllOf{}
	return &this
}

// GetBypassBusyWait returns the BypassBusyWait field value if set, zero value otherwise.
func (o *RequestSettingsResponseAllOf) GetBypassBusyWait() string {
	if o == nil || o.BypassBusyWait == nil {
		var ret string
		return ret
	}
	return *o.BypassBusyWait
}

// GetBypassBusyWaitOk returns a tuple with the BypassBusyWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsResponseAllOf) GetBypassBusyWaitOk() (*string, bool) {
	if o == nil || o.BypassBusyWait == nil {
		return nil, false
	}
	return o.BypassBusyWait, true
}

// HasBypassBusyWait returns a boolean if a field has been set.
func (o *RequestSettingsResponseAllOf) HasBypassBusyWait() bool {
	if o != nil && o.BypassBusyWait != nil {
		return true
	}

	return false
}

// SetBypassBusyWait gets a reference to the given string and assigns it to the BypassBusyWait field.
func (o *RequestSettingsResponseAllOf) SetBypassBusyWait(v string) {
	o.BypassBusyWait = &v
}

// GetForceMiss returns the ForceMiss field value if set, zero value otherwise.
func (o *RequestSettingsResponseAllOf) GetForceMiss() string {
	if o == nil || o.ForceMiss == nil {
		var ret string
		return ret
	}
	return *o.ForceMiss
}

// GetForceMissOk returns a tuple with the ForceMiss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsResponseAllOf) GetForceMissOk() (*string, bool) {
	if o == nil || o.ForceMiss == nil {
		return nil, false
	}
	return o.ForceMiss, true
}

// HasForceMiss returns a boolean if a field has been set.
func (o *RequestSettingsResponseAllOf) HasForceMiss() bool {
	if o != nil && o.ForceMiss != nil {
		return true
	}

	return false
}

// SetForceMiss gets a reference to the given string and assigns it to the ForceMiss field.
func (o *RequestSettingsResponseAllOf) SetForceMiss(v string) {
	o.ForceMiss = &v
}

// GetForceSsl returns the ForceSsl field value if set, zero value otherwise.
func (o *RequestSettingsResponseAllOf) GetForceSsl() string {
	if o == nil || o.ForceSsl == nil {
		var ret string
		return ret
	}
	return *o.ForceSsl
}

// GetForceSslOk returns a tuple with the ForceSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsResponseAllOf) GetForceSslOk() (*string, bool) {
	if o == nil || o.ForceSsl == nil {
		return nil, false
	}
	return o.ForceSsl, true
}

// HasForceSsl returns a boolean if a field has been set.
func (o *RequestSettingsResponseAllOf) HasForceSsl() bool {
	if o != nil && o.ForceSsl != nil {
		return true
	}

	return false
}

// SetForceSsl gets a reference to the given string and assigns it to the ForceSsl field.
func (o *RequestSettingsResponseAllOf) SetForceSsl(v string) {
	o.ForceSsl = &v
}

// GetGeoHeaders returns the GeoHeaders field value if set, zero value otherwise.
func (o *RequestSettingsResponseAllOf) GetGeoHeaders() string {
	if o == nil || o.GeoHeaders == nil {
		var ret string
		return ret
	}
	return *o.GeoHeaders
}

// GetGeoHeadersOk returns a tuple with the GeoHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsResponseAllOf) GetGeoHeadersOk() (*string, bool) {
	if o == nil || o.GeoHeaders == nil {
		return nil, false
	}
	return o.GeoHeaders, true
}

// HasGeoHeaders returns a boolean if a field has been set.
func (o *RequestSettingsResponseAllOf) HasGeoHeaders() bool {
	if o != nil && o.GeoHeaders != nil {
		return true
	}

	return false
}

// SetGeoHeaders gets a reference to the given string and assigns it to the GeoHeaders field.
func (o *RequestSettingsResponseAllOf) SetGeoHeaders(v string) {
	o.GeoHeaders = &v
}

// GetMaxStaleAge returns the MaxStaleAge field value if set, zero value otherwise.
func (o *RequestSettingsResponseAllOf) GetMaxStaleAge() string {
	if o == nil || o.MaxStaleAge == nil {
		var ret string
		return ret
	}
	return *o.MaxStaleAge
}

// GetMaxStaleAgeOk returns a tuple with the MaxStaleAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsResponseAllOf) GetMaxStaleAgeOk() (*string, bool) {
	if o == nil || o.MaxStaleAge == nil {
		return nil, false
	}
	return o.MaxStaleAge, true
}

// HasMaxStaleAge returns a boolean if a field has been set.
func (o *RequestSettingsResponseAllOf) HasMaxStaleAge() bool {
	if o != nil && o.MaxStaleAge != nil {
		return true
	}

	return false
}

// SetMaxStaleAge gets a reference to the given string and assigns it to the MaxStaleAge field.
func (o *RequestSettingsResponseAllOf) SetMaxStaleAge(v string) {
	o.MaxStaleAge = &v
}

// GetTimerSupport returns the TimerSupport field value if set, zero value otherwise.
func (o *RequestSettingsResponseAllOf) GetTimerSupport() string {
	if o == nil || o.TimerSupport == nil {
		var ret string
		return ret
	}
	return *o.TimerSupport
}

// GetTimerSupportOk returns a tuple with the TimerSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsResponseAllOf) GetTimerSupportOk() (*string, bool) {
	if o == nil || o.TimerSupport == nil {
		return nil, false
	}
	return o.TimerSupport, true
}

// HasTimerSupport returns a boolean if a field has been set.
func (o *RequestSettingsResponseAllOf) HasTimerSupport() bool {
	if o != nil && o.TimerSupport != nil {
		return true
	}

	return false
}

// SetTimerSupport gets a reference to the given string and assigns it to the TimerSupport field.
func (o *RequestSettingsResponseAllOf) SetTimerSupport(v string) {
	o.TimerSupport = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RequestSettingsResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.BypassBusyWait != nil {
		toSerialize["bypass_busy_wait"] = o.BypassBusyWait
	}
	if o.ForceMiss != nil {
		toSerialize["force_miss"] = o.ForceMiss
	}
	if o.ForceSsl != nil {
		toSerialize["force_ssl"] = o.ForceSsl
	}
	if o.GeoHeaders != nil {
		toSerialize["geo_headers"] = o.GeoHeaders
	}
	if o.MaxStaleAge != nil {
		toSerialize["max_stale_age"] = o.MaxStaleAge
	}
	if o.TimerSupport != nil {
		toSerialize["timer_support"] = o.TimerSupport
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RequestSettingsResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRequestSettingsResponseAllOf := _RequestSettingsResponseAllOf{}

	if err = json.Unmarshal(bytes, &varRequestSettingsResponseAllOf); err == nil {
		*o = RequestSettingsResponseAllOf(varRequestSettingsResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bypass_busy_wait")
		delete(additionalProperties, "force_miss")
		delete(additionalProperties, "force_ssl")
		delete(additionalProperties, "geo_headers")
		delete(additionalProperties, "max_stale_age")
		delete(additionalProperties, "timer_support")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRequestSettingsResponseAllOf is a helper abstraction for handling nullable requestsettingsresponseallof types. 
type NullableRequestSettingsResponseAllOf struct {
	value *RequestSettingsResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableRequestSettingsResponseAllOf) Get() *RequestSettingsResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableRequestSettingsResponseAllOf) Set(val *RequestSettingsResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRequestSettingsResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRequestSettingsResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRequestSettingsResponseAllOf returns a pointer to a new instance of NullableRequestSettingsResponseAllOf.
func NewNullableRequestSettingsResponseAllOf(val *RequestSettingsResponseAllOf) *NullableRequestSettingsResponseAllOf {
	return &NullableRequestSettingsResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRequestSettingsResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRequestSettingsResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
