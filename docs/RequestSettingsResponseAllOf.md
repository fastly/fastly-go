# RequestSettingsResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BypassBusyWait** | Pointer to **NullableString** | Disable collapsed forwarding, so you don&#39;t wait for other objects to origin. | [optional] 
**ForceMiss** | Pointer to **NullableString** | Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable. | [optional] 
**ForceSsl** | Pointer to **string** | Forces the request use SSL (redirects a non-SSL to SSL). | [optional] 
**GeoHeaders** | Pointer to **NullableString** | Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers. | [optional] 
**MaxStaleAge** | Pointer to **NullableString** | How old an object is allowed to be to serve stale-if-error or stale-while-revalidate. | [optional] 
**TimerSupport** | Pointer to **NullableString** | Injects the X-Timer info into the request for viewing origin fetch durations. | [optional] 

## Methods

### NewRequestSettingsResponseAllOf

`func NewRequestSettingsResponseAllOf() *RequestSettingsResponseAllOf`

NewRequestSettingsResponseAllOf instantiates a new RequestSettingsResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSettingsResponseAllOfWithDefaults

`func NewRequestSettingsResponseAllOfWithDefaults() *RequestSettingsResponseAllOf`

NewRequestSettingsResponseAllOfWithDefaults instantiates a new RequestSettingsResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBypassBusyWait

`func (o *RequestSettingsResponseAllOf) GetBypassBusyWait() string`

GetBypassBusyWait returns the BypassBusyWait field if non-nil, zero value otherwise.

### GetBypassBusyWaitOk

`func (o *RequestSettingsResponseAllOf) GetBypassBusyWaitOk() (*string, bool)`

GetBypassBusyWaitOk returns a tuple with the BypassBusyWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassBusyWait

`func (o *RequestSettingsResponseAllOf) SetBypassBusyWait(v string)`

SetBypassBusyWait sets BypassBusyWait field to given value.

### HasBypassBusyWait

`func (o *RequestSettingsResponseAllOf) HasBypassBusyWait() bool`

HasBypassBusyWait returns a boolean if a field has been set.

### SetBypassBusyWaitNil

`func (o *RequestSettingsResponseAllOf) SetBypassBusyWaitNil(b bool)`

 SetBypassBusyWaitNil sets the value for BypassBusyWait to be an explicit nil

### UnsetBypassBusyWait
`func (o *RequestSettingsResponseAllOf) UnsetBypassBusyWait()`

UnsetBypassBusyWait ensures that no value is present for BypassBusyWait, not even an explicit nil
### GetForceMiss

`func (o *RequestSettingsResponseAllOf) GetForceMiss() string`

GetForceMiss returns the ForceMiss field if non-nil, zero value otherwise.

### GetForceMissOk

`func (o *RequestSettingsResponseAllOf) GetForceMissOk() (*string, bool)`

GetForceMissOk returns a tuple with the ForceMiss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceMiss

`func (o *RequestSettingsResponseAllOf) SetForceMiss(v string)`

SetForceMiss sets ForceMiss field to given value.

### HasForceMiss

`func (o *RequestSettingsResponseAllOf) HasForceMiss() bool`

HasForceMiss returns a boolean if a field has been set.

### SetForceMissNil

`func (o *RequestSettingsResponseAllOf) SetForceMissNil(b bool)`

 SetForceMissNil sets the value for ForceMiss to be an explicit nil

### UnsetForceMiss
`func (o *RequestSettingsResponseAllOf) UnsetForceMiss()`

UnsetForceMiss ensures that no value is present for ForceMiss, not even an explicit nil
### GetForceSsl

`func (o *RequestSettingsResponseAllOf) GetForceSsl() string`

GetForceSsl returns the ForceSsl field if non-nil, zero value otherwise.

### GetForceSslOk

`func (o *RequestSettingsResponseAllOf) GetForceSslOk() (*string, bool)`

GetForceSslOk returns a tuple with the ForceSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSsl

`func (o *RequestSettingsResponseAllOf) SetForceSsl(v string)`

SetForceSsl sets ForceSsl field to given value.

### HasForceSsl

`func (o *RequestSettingsResponseAllOf) HasForceSsl() bool`

HasForceSsl returns a boolean if a field has been set.

### GetGeoHeaders

`func (o *RequestSettingsResponseAllOf) GetGeoHeaders() string`

GetGeoHeaders returns the GeoHeaders field if non-nil, zero value otherwise.

### GetGeoHeadersOk

`func (o *RequestSettingsResponseAllOf) GetGeoHeadersOk() (*string, bool)`

GetGeoHeadersOk returns a tuple with the GeoHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoHeaders

`func (o *RequestSettingsResponseAllOf) SetGeoHeaders(v string)`

SetGeoHeaders sets GeoHeaders field to given value.

### HasGeoHeaders

`func (o *RequestSettingsResponseAllOf) HasGeoHeaders() bool`

HasGeoHeaders returns a boolean if a field has been set.

### SetGeoHeadersNil

`func (o *RequestSettingsResponseAllOf) SetGeoHeadersNil(b bool)`

 SetGeoHeadersNil sets the value for GeoHeaders to be an explicit nil

### UnsetGeoHeaders
`func (o *RequestSettingsResponseAllOf) UnsetGeoHeaders()`

UnsetGeoHeaders ensures that no value is present for GeoHeaders, not even an explicit nil
### GetMaxStaleAge

`func (o *RequestSettingsResponseAllOf) GetMaxStaleAge() string`

GetMaxStaleAge returns the MaxStaleAge field if non-nil, zero value otherwise.

### GetMaxStaleAgeOk

`func (o *RequestSettingsResponseAllOf) GetMaxStaleAgeOk() (*string, bool)`

GetMaxStaleAgeOk returns a tuple with the MaxStaleAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaleAge

`func (o *RequestSettingsResponseAllOf) SetMaxStaleAge(v string)`

SetMaxStaleAge sets MaxStaleAge field to given value.

### HasMaxStaleAge

`func (o *RequestSettingsResponseAllOf) HasMaxStaleAge() bool`

HasMaxStaleAge returns a boolean if a field has been set.

### SetMaxStaleAgeNil

`func (o *RequestSettingsResponseAllOf) SetMaxStaleAgeNil(b bool)`

 SetMaxStaleAgeNil sets the value for MaxStaleAge to be an explicit nil

### UnsetMaxStaleAge
`func (o *RequestSettingsResponseAllOf) UnsetMaxStaleAge()`

UnsetMaxStaleAge ensures that no value is present for MaxStaleAge, not even an explicit nil
### GetTimerSupport

`func (o *RequestSettingsResponseAllOf) GetTimerSupport() string`

GetTimerSupport returns the TimerSupport field if non-nil, zero value otherwise.

### GetTimerSupportOk

`func (o *RequestSettingsResponseAllOf) GetTimerSupportOk() (*string, bool)`

GetTimerSupportOk returns a tuple with the TimerSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerSupport

`func (o *RequestSettingsResponseAllOf) SetTimerSupport(v string)`

SetTimerSupport sets TimerSupport field to given value.

### HasTimerSupport

`func (o *RequestSettingsResponseAllOf) HasTimerSupport() bool`

HasTimerSupport returns a boolean if a field has been set.

### SetTimerSupportNil

`func (o *RequestSettingsResponseAllOf) SetTimerSupportNil(b bool)`

 SetTimerSupportNil sets the value for TimerSupport to be an explicit nil

### UnsetTimerSupport
`func (o *RequestSettingsResponseAllOf) UnsetTimerSupport()`

UnsetTimerSupport ensures that no value is present for TimerSupport, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
