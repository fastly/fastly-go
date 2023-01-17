# RequestSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Allows you to terminate request handling and immediately perform an action. | [optional] 
**BypassBusyWait** | Pointer to **int32** | Disable collapsed forwarding, so you don&#39;t wait for other objects to origin. | [optional] 
**DefaultHost** | Pointer to **NullableString** | Sets the host header. | [optional] 
**ForceMiss** | Pointer to **int32** | Allows you to force a cache miss for the request. Replaces the item in the cache if the content is cacheable. | [optional] 
**ForceSsl** | Pointer to **int32** | Forces the request use SSL (redirects a non-SSL to SSL). | [optional] 
**GeoHeaders** | Pointer to **int32** | Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers. | [optional] 
**HashKeys** | Pointer to **NullableString** | Comma separated list of varnish request object fields that should be in the hash key. | [optional] 
**MaxStaleAge** | Pointer to **int32** | How old an object is allowed to be to serve stale-if-error or stale-while-revalidate. | [optional] 
**Name** | Pointer to **string** | Name for the request settings. | [optional] 
**RequestCondition** | Pointer to **NullableString** | Condition which, if met, will select this configuration during a request. Optional. | [optional] 
**TimerSupport** | Pointer to **int32** | Injects the X-Timer info into the request for viewing origin fetch durations. | [optional] 
**Xff** | Pointer to **string** | Short for X-Forwarded-For. | [optional] 

## Methods

### NewRequestSettings

`func NewRequestSettings() *RequestSettings`

NewRequestSettings instantiates a new RequestSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSettingsWithDefaults

`func NewRequestSettingsWithDefaults() *RequestSettings`

NewRequestSettingsWithDefaults instantiates a new RequestSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RequestSettings) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RequestSettings) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RequestSettings) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RequestSettings) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *RequestSettings) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *RequestSettings) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetBypassBusyWait

`func (o *RequestSettings) GetBypassBusyWait() int32`

GetBypassBusyWait returns the BypassBusyWait field if non-nil, zero value otherwise.

### GetBypassBusyWaitOk

`func (o *RequestSettings) GetBypassBusyWaitOk() (*int32, bool)`

GetBypassBusyWaitOk returns a tuple with the BypassBusyWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassBusyWait

`func (o *RequestSettings) SetBypassBusyWait(v int32)`

SetBypassBusyWait sets BypassBusyWait field to given value.

### HasBypassBusyWait

`func (o *RequestSettings) HasBypassBusyWait() bool`

HasBypassBusyWait returns a boolean if a field has been set.

### GetDefaultHost

`func (o *RequestSettings) GetDefaultHost() string`

GetDefaultHost returns the DefaultHost field if non-nil, zero value otherwise.

### GetDefaultHostOk

`func (o *RequestSettings) GetDefaultHostOk() (*string, bool)`

GetDefaultHostOk returns a tuple with the DefaultHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHost

`func (o *RequestSettings) SetDefaultHost(v string)`

SetDefaultHost sets DefaultHost field to given value.

### HasDefaultHost

`func (o *RequestSettings) HasDefaultHost() bool`

HasDefaultHost returns a boolean if a field has been set.

### SetDefaultHostNil

`func (o *RequestSettings) SetDefaultHostNil(b bool)`

 SetDefaultHostNil sets the value for DefaultHost to be an explicit nil

### UnsetDefaultHost
`func (o *RequestSettings) UnsetDefaultHost()`

UnsetDefaultHost ensures that no value is present for DefaultHost, not even an explicit nil
### GetForceMiss

`func (o *RequestSettings) GetForceMiss() int32`

GetForceMiss returns the ForceMiss field if non-nil, zero value otherwise.

### GetForceMissOk

`func (o *RequestSettings) GetForceMissOk() (*int32, bool)`

GetForceMissOk returns a tuple with the ForceMiss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceMiss

`func (o *RequestSettings) SetForceMiss(v int32)`

SetForceMiss sets ForceMiss field to given value.

### HasForceMiss

`func (o *RequestSettings) HasForceMiss() bool`

HasForceMiss returns a boolean if a field has been set.

### GetForceSsl

`func (o *RequestSettings) GetForceSsl() int32`

GetForceSsl returns the ForceSsl field if non-nil, zero value otherwise.

### GetForceSslOk

`func (o *RequestSettings) GetForceSslOk() (*int32, bool)`

GetForceSslOk returns a tuple with the ForceSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSsl

`func (o *RequestSettings) SetForceSsl(v int32)`

SetForceSsl sets ForceSsl field to given value.

### HasForceSsl

`func (o *RequestSettings) HasForceSsl() bool`

HasForceSsl returns a boolean if a field has been set.

### GetGeoHeaders

`func (o *RequestSettings) GetGeoHeaders() int32`

GetGeoHeaders returns the GeoHeaders field if non-nil, zero value otherwise.

### GetGeoHeadersOk

`func (o *RequestSettings) GetGeoHeadersOk() (*int32, bool)`

GetGeoHeadersOk returns a tuple with the GeoHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoHeaders

`func (o *RequestSettings) SetGeoHeaders(v int32)`

SetGeoHeaders sets GeoHeaders field to given value.

### HasGeoHeaders

`func (o *RequestSettings) HasGeoHeaders() bool`

HasGeoHeaders returns a boolean if a field has been set.

### GetHashKeys

`func (o *RequestSettings) GetHashKeys() string`

GetHashKeys returns the HashKeys field if non-nil, zero value otherwise.

### GetHashKeysOk

`func (o *RequestSettings) GetHashKeysOk() (*string, bool)`

GetHashKeysOk returns a tuple with the HashKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashKeys

`func (o *RequestSettings) SetHashKeys(v string)`

SetHashKeys sets HashKeys field to given value.

### HasHashKeys

`func (o *RequestSettings) HasHashKeys() bool`

HasHashKeys returns a boolean if a field has been set.

### SetHashKeysNil

`func (o *RequestSettings) SetHashKeysNil(b bool)`

 SetHashKeysNil sets the value for HashKeys to be an explicit nil

### UnsetHashKeys
`func (o *RequestSettings) UnsetHashKeys()`

UnsetHashKeys ensures that no value is present for HashKeys, not even an explicit nil
### GetMaxStaleAge

`func (o *RequestSettings) GetMaxStaleAge() int32`

GetMaxStaleAge returns the MaxStaleAge field if non-nil, zero value otherwise.

### GetMaxStaleAgeOk

`func (o *RequestSettings) GetMaxStaleAgeOk() (*int32, bool)`

GetMaxStaleAgeOk returns a tuple with the MaxStaleAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaleAge

`func (o *RequestSettings) SetMaxStaleAge(v int32)`

SetMaxStaleAge sets MaxStaleAge field to given value.

### HasMaxStaleAge

`func (o *RequestSettings) HasMaxStaleAge() bool`

HasMaxStaleAge returns a boolean if a field has been set.

### GetName

`func (o *RequestSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequestCondition

`func (o *RequestSettings) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *RequestSettings) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *RequestSettings) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *RequestSettings) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *RequestSettings) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *RequestSettings) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetTimerSupport

`func (o *RequestSettings) GetTimerSupport() int32`

GetTimerSupport returns the TimerSupport field if non-nil, zero value otherwise.

### GetTimerSupportOk

`func (o *RequestSettings) GetTimerSupportOk() (*int32, bool)`

GetTimerSupportOk returns a tuple with the TimerSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerSupport

`func (o *RequestSettings) SetTimerSupport(v int32)`

SetTimerSupport sets TimerSupport field to given value.

### HasTimerSupport

`func (o *RequestSettings) HasTimerSupport() bool`

HasTimerSupport returns a boolean if a field has been set.

### GetXff

`func (o *RequestSettings) GetXff() string`

GetXff returns the Xff field if non-nil, zero value otherwise.

### GetXffOk

`func (o *RequestSettings) GetXffOk() (*string, bool)`

GetXffOk returns a tuple with the Xff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXff

`func (o *RequestSettings) SetXff(v string)`

SetXff sets Xff field to given value.

### HasXff

`func (o *RequestSettings) HasXff() bool`

HasXff returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
