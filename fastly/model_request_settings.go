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

// RequestSettings struct for RequestSettings
type RequestSettings struct {
	// Allows you to terminate request handling and immediately perform an action.
	Action NullableString `json:"action,omitempty"`
	// Disable collapsed forwarding, so you don't wait for other objects to origin.
	BypassBusyWait *int32 `json:"bypass_busy_wait,omitempty"`
	// Sets the host header.
	DefaultHost NullableString `json:"default_host,omitempty"`
	// Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable.
	ForceMiss *int32 `json:"force_miss,omitempty"`
	// Forces the request use SSL (redirects a non-SSL to SSL).
	ForceSsl *int32 `json:"force_ssl,omitempty"`
	// Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers.
	GeoHeaders *int32 `json:"geo_headers,omitempty"`
	// Comma separated list of varnish request object fields that should be in the hash key.
	HashKeys NullableString `json:"hash_keys,omitempty"`
	// How old an object is allowed to be to serve stale-if-error or stale-while-revalidate.
	MaxStaleAge *int32 `json:"max_stale_age,omitempty"`
	// Name for the request settings.
	Name *string `json:"name,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition NullableString `json:"request_condition,omitempty"`
	// Injects the X-Timer info into the request for viewing origin fetch durations.
	TimerSupport *int32 `json:"timer_support,omitempty"`
	// Short for X-Forwarded-For.
	Xff *string `json:"xff,omitempty"`
	AdditionalProperties map[string]any
}

type _RequestSettings RequestSettings

// NewRequestSettings instantiates a new RequestSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSettings() *RequestSettings {
	this := RequestSettings{}
	return &this
}

// NewRequestSettingsWithDefaults instantiates a new RequestSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSettingsWithDefaults() *RequestSettings {
	this := RequestSettings{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettings) GetAction() string {
	if o == nil || o.Action.Get() == nil {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettings) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *RequestSettings) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *RequestSettings) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *RequestSettings) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *RequestSettings) UnsetAction() {
	o.Action.Unset()
}

// GetBypassBusyWait returns the BypassBusyWait field value if set, zero value otherwise.
func (o *RequestSettings) GetBypassBusyWait() int32 {
	if o == nil || o.BypassBusyWait == nil {
		var ret int32
		return ret
	}
	return *o.BypassBusyWait
}

// GetBypassBusyWaitOk returns a tuple with the BypassBusyWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetBypassBusyWaitOk() (*int32, bool) {
	if o == nil || o.BypassBusyWait == nil {
		return nil, false
	}
	return o.BypassBusyWait, true
}

// HasBypassBusyWait returns a boolean if a field has been set.
func (o *RequestSettings) HasBypassBusyWait() bool {
	if o != nil && o.BypassBusyWait != nil {
		return true
	}

	return false
}

// SetBypassBusyWait gets a reference to the given int32 and assigns it to the BypassBusyWait field.
func (o *RequestSettings) SetBypassBusyWait(v int32) {
	o.BypassBusyWait = &v
}

// GetDefaultHost returns the DefaultHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettings) GetDefaultHost() string {
	if o == nil || o.DefaultHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultHost.Get()
}

// GetDefaultHostOk returns a tuple with the DefaultHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettings) GetDefaultHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultHost.Get(), o.DefaultHost.IsSet()
}

// HasDefaultHost returns a boolean if a field has been set.
func (o *RequestSettings) HasDefaultHost() bool {
	if o != nil && o.DefaultHost.IsSet() {
		return true
	}

	return false
}

// SetDefaultHost gets a reference to the given NullableString and assigns it to the DefaultHost field.
func (o *RequestSettings) SetDefaultHost(v string) {
	o.DefaultHost.Set(&v)
}
// SetDefaultHostNil sets the value for DefaultHost to be an explicit nil
func (o *RequestSettings) SetDefaultHostNil() {
	o.DefaultHost.Set(nil)
}

// UnsetDefaultHost ensures that no value is present for DefaultHost, not even an explicit nil
func (o *RequestSettings) UnsetDefaultHost() {
	o.DefaultHost.Unset()
}

// GetForceMiss returns the ForceMiss field value if set, zero value otherwise.
func (o *RequestSettings) GetForceMiss() int32 {
	if o == nil || o.ForceMiss == nil {
		var ret int32
		return ret
	}
	return *o.ForceMiss
}

// GetForceMissOk returns a tuple with the ForceMiss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetForceMissOk() (*int32, bool) {
	if o == nil || o.ForceMiss == nil {
		return nil, false
	}
	return o.ForceMiss, true
}

// HasForceMiss returns a boolean if a field has been set.
func (o *RequestSettings) HasForceMiss() bool {
	if o != nil && o.ForceMiss != nil {
		return true
	}

	return false
}

// SetForceMiss gets a reference to the given int32 and assigns it to the ForceMiss field.
func (o *RequestSettings) SetForceMiss(v int32) {
	o.ForceMiss = &v
}

// GetForceSsl returns the ForceSsl field value if set, zero value otherwise.
func (o *RequestSettings) GetForceSsl() int32 {
	if o == nil || o.ForceSsl == nil {
		var ret int32
		return ret
	}
	return *o.ForceSsl
}

// GetForceSslOk returns a tuple with the ForceSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetForceSslOk() (*int32, bool) {
	if o == nil || o.ForceSsl == nil {
		return nil, false
	}
	return o.ForceSsl, true
}

// HasForceSsl returns a boolean if a field has been set.
func (o *RequestSettings) HasForceSsl() bool {
	if o != nil && o.ForceSsl != nil {
		return true
	}

	return false
}

// SetForceSsl gets a reference to the given int32 and assigns it to the ForceSsl field.
func (o *RequestSettings) SetForceSsl(v int32) {
	o.ForceSsl = &v
}

// GetGeoHeaders returns the GeoHeaders field value if set, zero value otherwise.
func (o *RequestSettings) GetGeoHeaders() int32 {
	if o == nil || o.GeoHeaders == nil {
		var ret int32
		return ret
	}
	return *o.GeoHeaders
}

// GetGeoHeadersOk returns a tuple with the GeoHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetGeoHeadersOk() (*int32, bool) {
	if o == nil || o.GeoHeaders == nil {
		return nil, false
	}
	return o.GeoHeaders, true
}

// HasGeoHeaders returns a boolean if a field has been set.
func (o *RequestSettings) HasGeoHeaders() bool {
	if o != nil && o.GeoHeaders != nil {
		return true
	}

	return false
}

// SetGeoHeaders gets a reference to the given int32 and assigns it to the GeoHeaders field.
func (o *RequestSettings) SetGeoHeaders(v int32) {
	o.GeoHeaders = &v
}

// GetHashKeys returns the HashKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettings) GetHashKeys() string {
	if o == nil || o.HashKeys.Get() == nil {
		var ret string
		return ret
	}
	return *o.HashKeys.Get()
}

// GetHashKeysOk returns a tuple with the HashKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettings) GetHashKeysOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HashKeys.Get(), o.HashKeys.IsSet()
}

// HasHashKeys returns a boolean if a field has been set.
func (o *RequestSettings) HasHashKeys() bool {
	if o != nil && o.HashKeys.IsSet() {
		return true
	}

	return false
}

// SetHashKeys gets a reference to the given NullableString and assigns it to the HashKeys field.
func (o *RequestSettings) SetHashKeys(v string) {
	o.HashKeys.Set(&v)
}
// SetHashKeysNil sets the value for HashKeys to be an explicit nil
func (o *RequestSettings) SetHashKeysNil() {
	o.HashKeys.Set(nil)
}

// UnsetHashKeys ensures that no value is present for HashKeys, not even an explicit nil
func (o *RequestSettings) UnsetHashKeys() {
	o.HashKeys.Unset()
}

// GetMaxStaleAge returns the MaxStaleAge field value if set, zero value otherwise.
func (o *RequestSettings) GetMaxStaleAge() int32 {
	if o == nil || o.MaxStaleAge == nil {
		var ret int32
		return ret
	}
	return *o.MaxStaleAge
}

// GetMaxStaleAgeOk returns a tuple with the MaxStaleAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetMaxStaleAgeOk() (*int32, bool) {
	if o == nil || o.MaxStaleAge == nil {
		return nil, false
	}
	return o.MaxStaleAge, true
}

// HasMaxStaleAge returns a boolean if a field has been set.
func (o *RequestSettings) HasMaxStaleAge() bool {
	if o != nil && o.MaxStaleAge != nil {
		return true
	}

	return false
}

// SetMaxStaleAge gets a reference to the given int32 and assigns it to the MaxStaleAge field.
func (o *RequestSettings) SetMaxStaleAge(v int32) {
	o.MaxStaleAge = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestSettings) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestSettings) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestSettings) SetName(v string) {
	o.Name = &v
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettings) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettings) GetRequestConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *RequestSettings) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *RequestSettings) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}
// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *RequestSettings) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *RequestSettings) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// GetTimerSupport returns the TimerSupport field value if set, zero value otherwise.
func (o *RequestSettings) GetTimerSupport() int32 {
	if o == nil || o.TimerSupport == nil {
		var ret int32
		return ret
	}
	return *o.TimerSupport
}

// GetTimerSupportOk returns a tuple with the TimerSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetTimerSupportOk() (*int32, bool) {
	if o == nil || o.TimerSupport == nil {
		return nil, false
	}
	return o.TimerSupport, true
}

// HasTimerSupport returns a boolean if a field has been set.
func (o *RequestSettings) HasTimerSupport() bool {
	if o != nil && o.TimerSupport != nil {
		return true
	}

	return false
}

// SetTimerSupport gets a reference to the given int32 and assigns it to the TimerSupport field.
func (o *RequestSettings) SetTimerSupport(v int32) {
	o.TimerSupport = &v
}

// GetXff returns the Xff field value if set, zero value otherwise.
func (o *RequestSettings) GetXff() string {
	if o == nil || o.Xff == nil {
		var ret string
		return ret
	}
	return *o.Xff
}

// GetXffOk returns a tuple with the Xff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettings) GetXffOk() (*string, bool) {
	if o == nil || o.Xff == nil {
		return nil, false
	}
	return o.Xff, true
}

// HasXff returns a boolean if a field has been set.
func (o *RequestSettings) HasXff() bool {
	if o != nil && o.Xff != nil {
		return true
	}

	return false
}

// SetXff gets a reference to the given string and assigns it to the Xff field.
func (o *RequestSettings) SetXff(v string) {
	o.Xff = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RequestSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.BypassBusyWait != nil {
		toSerialize["bypass_busy_wait"] = o.BypassBusyWait
	}
	if o.DefaultHost.IsSet() {
		toSerialize["default_host"] = o.DefaultHost.Get()
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
	if o.HashKeys.IsSet() {
		toSerialize["hash_keys"] = o.HashKeys.Get()
	}
	if o.MaxStaleAge != nil {
		toSerialize["max_stale_age"] = o.MaxStaleAge
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RequestCondition.IsSet() {
		toSerialize["request_condition"] = o.RequestCondition.Get()
	}
	if o.TimerSupport != nil {
		toSerialize["timer_support"] = o.TimerSupport
	}
	if o.Xff != nil {
		toSerialize["xff"] = o.Xff
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RequestSettings) UnmarshalJSON(bytes []byte) (err error) {
	varRequestSettings := _RequestSettings{}

	if err = json.Unmarshal(bytes, &varRequestSettings); err == nil {
		*o = RequestSettings(varRequestSettings)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "bypass_busy_wait")
		delete(additionalProperties, "default_host")
		delete(additionalProperties, "force_miss")
		delete(additionalProperties, "force_ssl")
		delete(additionalProperties, "geo_headers")
		delete(additionalProperties, "hash_keys")
		delete(additionalProperties, "max_stale_age")
		delete(additionalProperties, "name")
		delete(additionalProperties, "request_condition")
		delete(additionalProperties, "timer_support")
		delete(additionalProperties, "xff")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRequestSettings is a helper abstraction for handling nullable requestsettings types. 
type NullableRequestSettings struct {
	value *RequestSettings
	isSet bool
}

// Get returns the value.
func (v NullableRequestSettings) Get() *RequestSettings {
	return v.value
}

// Set modifies the value.
func (v *NullableRequestSettings) Set(val *RequestSettings) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRequestSettings) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRequestSettings) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRequestSettings returns a pointer to a new instance of NullableRequestSettings.
func NewNullableRequestSettings(val *RequestSettings) *NullableRequestSettings {
	return &NullableRequestSettings{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRequestSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRequestSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
