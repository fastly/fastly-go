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

// RequestSettingsResponse struct for RequestSettingsResponse
type RequestSettingsResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ServiceID *string      `json:"service_id,omitempty"`
	Version   *string      `json:"version,omitempty"`
	// Allows you to terminate request handling and immediately perform an action.
	Action NullableString `json:"action,omitempty"`
	// Sets the host header.
	DefaultHost NullableString `json:"default_host,omitempty"`
	// Comma separated list of varnish request object fields that should be in the hash key.
	HashKeys NullableString `json:"hash_keys,omitempty"`
	// Name for the request settings.
	Name *string `json:"name,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition NullableString `json:"request_condition,omitempty"`
	// Short for X-Forwarded-For.
	Xff NullableString `json:"xff,omitempty"`
	// Disable collapsed forwarding, so you don't wait for other objects to origin.
	BypassBusyWait NullableString `json:"bypass_busy_wait,omitempty"`
	// Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable.
	ForceMiss NullableString `json:"force_miss,omitempty"`
	// Forces the request use SSL (redirects a non-SSL to SSL).
	ForceSsl *string `json:"force_ssl,omitempty"`
	// Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers.
	GeoHeaders NullableString `json:"geo_headers,omitempty"`
	// How old an object is allowed to be to serve stale-if-error or stale-while-revalidate.
	MaxStaleAge NullableString `json:"max_stale_age,omitempty"`
	// Injects the X-Timer info into the request for viewing origin fetch durations.
	TimerSupport         NullableString `json:"timer_support,omitempty"`
	AdditionalProperties map[string]any
}

type _RequestSettingsResponse RequestSettingsResponse

// NewRequestSettingsResponse instantiates a new RequestSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSettingsResponse() *RequestSettingsResponse {
	this := RequestSettingsResponse{}
	return &this
}

// NewRequestSettingsResponseWithDefaults instantiates a new RequestSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSettingsResponseWithDefaults() *RequestSettingsResponse {
	this := RequestSettingsResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *RequestSettingsResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *RequestSettingsResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *RequestSettingsResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *RequestSettingsResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *RequestSettingsResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *RequestSettingsResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *RequestSettingsResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *RequestSettingsResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *RequestSettingsResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *RequestSettingsResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *RequestSettingsResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RequestSettingsResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RequestSettingsResponse) SetVersion(v string) {
	o.Version = &v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetAction() string {
	if o == nil || o.Action.Get() == nil {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *RequestSettingsResponse) SetAction(v string) {
	o.Action.Set(&v)
}

// SetActionNil sets the value for Action to be an explicit nil
func (o *RequestSettingsResponse) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *RequestSettingsResponse) UnsetAction() {
	o.Action.Unset()
}

// GetDefaultHost returns the DefaultHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetDefaultHost() string {
	if o == nil || o.DefaultHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultHost.Get()
}

// GetDefaultHostOk returns a tuple with the DefaultHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetDefaultHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultHost.Get(), o.DefaultHost.IsSet()
}

// HasDefaultHost returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasDefaultHost() bool {
	if o != nil && o.DefaultHost.IsSet() {
		return true
	}

	return false
}

// SetDefaultHost gets a reference to the given NullableString and assigns it to the DefaultHost field.
func (o *RequestSettingsResponse) SetDefaultHost(v string) {
	o.DefaultHost.Set(&v)
}

// SetDefaultHostNil sets the value for DefaultHost to be an explicit nil
func (o *RequestSettingsResponse) SetDefaultHostNil() {
	o.DefaultHost.Set(nil)
}

// UnsetDefaultHost ensures that no value is present for DefaultHost, not even an explicit nil
func (o *RequestSettingsResponse) UnsetDefaultHost() {
	o.DefaultHost.Unset()
}

// GetHashKeys returns the HashKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetHashKeys() string {
	if o == nil || o.HashKeys.Get() == nil {
		var ret string
		return ret
	}
	return *o.HashKeys.Get()
}

// GetHashKeysOk returns a tuple with the HashKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetHashKeysOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HashKeys.Get(), o.HashKeys.IsSet()
}

// HasHashKeys returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasHashKeys() bool {
	if o != nil && o.HashKeys.IsSet() {
		return true
	}

	return false
}

// SetHashKeys gets a reference to the given NullableString and assigns it to the HashKeys field.
func (o *RequestSettingsResponse) SetHashKeys(v string) {
	o.HashKeys.Set(&v)
}

// SetHashKeysNil sets the value for HashKeys to be an explicit nil
func (o *RequestSettingsResponse) SetHashKeysNil() {
	o.HashKeys.Set(nil)
}

// UnsetHashKeys ensures that no value is present for HashKeys, not even an explicit nil
func (o *RequestSettingsResponse) UnsetHashKeys() {
	o.HashKeys.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestSettingsResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestSettingsResponse) SetName(v string) {
	o.Name = &v
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetRequestConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *RequestSettingsResponse) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}

// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *RequestSettingsResponse) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *RequestSettingsResponse) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// GetXff returns the Xff field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetXff() string {
	if o == nil || o.Xff.Get() == nil {
		var ret string
		return ret
	}
	return *o.Xff.Get()
}

// GetXffOk returns a tuple with the Xff field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetXffOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Xff.Get(), o.Xff.IsSet()
}

// HasXff returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasXff() bool {
	if o != nil && o.Xff.IsSet() {
		return true
	}

	return false
}

// SetXff gets a reference to the given NullableString and assigns it to the Xff field.
func (o *RequestSettingsResponse) SetXff(v string) {
	o.Xff.Set(&v)
}

// SetXffNil sets the value for Xff to be an explicit nil
func (o *RequestSettingsResponse) SetXffNil() {
	o.Xff.Set(nil)
}

// UnsetXff ensures that no value is present for Xff, not even an explicit nil
func (o *RequestSettingsResponse) UnsetXff() {
	o.Xff.Unset()
}

// GetBypassBusyWait returns the BypassBusyWait field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetBypassBusyWait() string {
	if o == nil || o.BypassBusyWait.Get() == nil {
		var ret string
		return ret
	}
	return *o.BypassBusyWait.Get()
}

// GetBypassBusyWaitOk returns a tuple with the BypassBusyWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetBypassBusyWaitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BypassBusyWait.Get(), o.BypassBusyWait.IsSet()
}

// HasBypassBusyWait returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasBypassBusyWait() bool {
	if o != nil && o.BypassBusyWait.IsSet() {
		return true
	}

	return false
}

// SetBypassBusyWait gets a reference to the given NullableString and assigns it to the BypassBusyWait field.
func (o *RequestSettingsResponse) SetBypassBusyWait(v string) {
	o.BypassBusyWait.Set(&v)
}

// SetBypassBusyWaitNil sets the value for BypassBusyWait to be an explicit nil
func (o *RequestSettingsResponse) SetBypassBusyWaitNil() {
	o.BypassBusyWait.Set(nil)
}

// UnsetBypassBusyWait ensures that no value is present for BypassBusyWait, not even an explicit nil
func (o *RequestSettingsResponse) UnsetBypassBusyWait() {
	o.BypassBusyWait.Unset()
}

// GetForceMiss returns the ForceMiss field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetForceMiss() string {
	if o == nil || o.ForceMiss.Get() == nil {
		var ret string
		return ret
	}
	return *o.ForceMiss.Get()
}

// GetForceMissOk returns a tuple with the ForceMiss field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetForceMissOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForceMiss.Get(), o.ForceMiss.IsSet()
}

// HasForceMiss returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasForceMiss() bool {
	if o != nil && o.ForceMiss.IsSet() {
		return true
	}

	return false
}

// SetForceMiss gets a reference to the given NullableString and assigns it to the ForceMiss field.
func (o *RequestSettingsResponse) SetForceMiss(v string) {
	o.ForceMiss.Set(&v)
}

// SetForceMissNil sets the value for ForceMiss to be an explicit nil
func (o *RequestSettingsResponse) SetForceMissNil() {
	o.ForceMiss.Set(nil)
}

// UnsetForceMiss ensures that no value is present for ForceMiss, not even an explicit nil
func (o *RequestSettingsResponse) UnsetForceMiss() {
	o.ForceMiss.Unset()
}

// GetForceSsl returns the ForceSsl field value if set, zero value otherwise.
func (o *RequestSettingsResponse) GetForceSsl() string {
	if o == nil || o.ForceSsl == nil {
		var ret string
		return ret
	}
	return *o.ForceSsl
}

// GetForceSslOk returns a tuple with the ForceSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsResponse) GetForceSslOk() (*string, bool) {
	if o == nil || o.ForceSsl == nil {
		return nil, false
	}
	return o.ForceSsl, true
}

// HasForceSsl returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasForceSsl() bool {
	if o != nil && o.ForceSsl != nil {
		return true
	}

	return false
}

// SetForceSsl gets a reference to the given string and assigns it to the ForceSsl field.
func (o *RequestSettingsResponse) SetForceSsl(v string) {
	o.ForceSsl = &v
}

// GetGeoHeaders returns the GeoHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetGeoHeaders() string {
	if o == nil || o.GeoHeaders.Get() == nil {
		var ret string
		return ret
	}
	return *o.GeoHeaders.Get()
}

// GetGeoHeadersOk returns a tuple with the GeoHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetGeoHeadersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GeoHeaders.Get(), o.GeoHeaders.IsSet()
}

// HasGeoHeaders returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasGeoHeaders() bool {
	if o != nil && o.GeoHeaders.IsSet() {
		return true
	}

	return false
}

// SetGeoHeaders gets a reference to the given NullableString and assigns it to the GeoHeaders field.
func (o *RequestSettingsResponse) SetGeoHeaders(v string) {
	o.GeoHeaders.Set(&v)
}

// SetGeoHeadersNil sets the value for GeoHeaders to be an explicit nil
func (o *RequestSettingsResponse) SetGeoHeadersNil() {
	o.GeoHeaders.Set(nil)
}

// UnsetGeoHeaders ensures that no value is present for GeoHeaders, not even an explicit nil
func (o *RequestSettingsResponse) UnsetGeoHeaders() {
	o.GeoHeaders.Unset()
}

// GetMaxStaleAge returns the MaxStaleAge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetMaxStaleAge() string {
	if o == nil || o.MaxStaleAge.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaxStaleAge.Get()
}

// GetMaxStaleAgeOk returns a tuple with the MaxStaleAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetMaxStaleAgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStaleAge.Get(), o.MaxStaleAge.IsSet()
}

// HasMaxStaleAge returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasMaxStaleAge() bool {
	if o != nil && o.MaxStaleAge.IsSet() {
		return true
	}

	return false
}

// SetMaxStaleAge gets a reference to the given NullableString and assigns it to the MaxStaleAge field.
func (o *RequestSettingsResponse) SetMaxStaleAge(v string) {
	o.MaxStaleAge.Set(&v)
}

// SetMaxStaleAgeNil sets the value for MaxStaleAge to be an explicit nil
func (o *RequestSettingsResponse) SetMaxStaleAgeNil() {
	o.MaxStaleAge.Set(nil)
}

// UnsetMaxStaleAge ensures that no value is present for MaxStaleAge, not even an explicit nil
func (o *RequestSettingsResponse) UnsetMaxStaleAge() {
	o.MaxStaleAge.Unset()
}

// GetTimerSupport returns the TimerSupport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsResponse) GetTimerSupport() string {
	if o == nil || o.TimerSupport.Get() == nil {
		var ret string
		return ret
	}
	return *o.TimerSupport.Get()
}

// GetTimerSupportOk returns a tuple with the TimerSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsResponse) GetTimerSupportOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimerSupport.Get(), o.TimerSupport.IsSet()
}

// HasTimerSupport returns a boolean if a field has been set.
func (o *RequestSettingsResponse) HasTimerSupport() bool {
	if o != nil && o.TimerSupport.IsSet() {
		return true
	}

	return false
}

// SetTimerSupport gets a reference to the given NullableString and assigns it to the TimerSupport field.
func (o *RequestSettingsResponse) SetTimerSupport(v string) {
	o.TimerSupport.Set(&v)
}

// SetTimerSupportNil sets the value for TimerSupport to be an explicit nil
func (o *RequestSettingsResponse) SetTimerSupportNil() {
	o.TimerSupport.Set(nil)
}

// UnsetTimerSupport ensures that no value is present for TimerSupport, not even an explicit nil
func (o *RequestSettingsResponse) UnsetTimerSupport() {
	o.TimerSupport.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RequestSettingsResponse) MarshalJSON() ([]byte, error) {
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
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.DefaultHost.IsSet() {
		toSerialize["default_host"] = o.DefaultHost.Get()
	}
	if o.HashKeys.IsSet() {
		toSerialize["hash_keys"] = o.HashKeys.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RequestCondition.IsSet() {
		toSerialize["request_condition"] = o.RequestCondition.Get()
	}
	if o.Xff.IsSet() {
		toSerialize["xff"] = o.Xff.Get()
	}
	if o.BypassBusyWait.IsSet() {
		toSerialize["bypass_busy_wait"] = o.BypassBusyWait.Get()
	}
	if o.ForceMiss.IsSet() {
		toSerialize["force_miss"] = o.ForceMiss.Get()
	}
	if o.ForceSsl != nil {
		toSerialize["force_ssl"] = o.ForceSsl
	}
	if o.GeoHeaders.IsSet() {
		toSerialize["geo_headers"] = o.GeoHeaders.Get()
	}
	if o.MaxStaleAge.IsSet() {
		toSerialize["max_stale_age"] = o.MaxStaleAge.Get()
	}
	if o.TimerSupport.IsSet() {
		toSerialize["timer_support"] = o.TimerSupport.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RequestSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRequestSettingsResponse := _RequestSettingsResponse{}

	if err = json.Unmarshal(bytes, &varRequestSettingsResponse); err == nil {
		*o = RequestSettingsResponse(varRequestSettingsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "action")
		delete(additionalProperties, "default_host")
		delete(additionalProperties, "hash_keys")
		delete(additionalProperties, "name")
		delete(additionalProperties, "request_condition")
		delete(additionalProperties, "xff")
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

// NullableRequestSettingsResponse is a helper abstraction for handling nullable requestsettingsresponse types.
type NullableRequestSettingsResponse struct {
	value *RequestSettingsResponse
	isSet bool
}

// Get returns the value.
func (v NullableRequestSettingsResponse) Get() *RequestSettingsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableRequestSettingsResponse) Set(val *RequestSettingsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRequestSettingsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRequestSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRequestSettingsResponse returns a pointer to a new instance of NullableRequestSettingsResponse.
func NewNullableRequestSettingsResponse(val *RequestSettingsResponse) *NullableRequestSettingsResponse {
	return &NullableRequestSettingsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRequestSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRequestSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
