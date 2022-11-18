# RequestSettingsResponse

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
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **int32** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 

## Methods

### NewRequestSettingsResponse

`func NewRequestSettingsResponse() *RequestSettingsResponse`

NewRequestSettingsResponse instantiates a new RequestSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSettingsResponseWithDefaults

`func NewRequestSettingsResponseWithDefaults() *RequestSettingsResponse`

NewRequestSettingsResponseWithDefaults instantiates a new RequestSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RequestSettingsResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RequestSettingsResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RequestSettingsResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RequestSettingsResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *RequestSettingsResponse) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *RequestSettingsResponse) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetBypassBusyWait

`func (o *RequestSettingsResponse) GetBypassBusyWait() int32`

GetBypassBusyWait returns the BypassBusyWait field if non-nil, zero value otherwise.

### GetBypassBusyWaitOk

`func (o *RequestSettingsResponse) GetBypassBusyWaitOk() (*int32, bool)`

GetBypassBusyWaitOk returns a tuple with the BypassBusyWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassBusyWait

`func (o *RequestSettingsResponse) SetBypassBusyWait(v int32)`

SetBypassBusyWait sets BypassBusyWait field to given value.

### HasBypassBusyWait

`func (o *RequestSettingsResponse) HasBypassBusyWait() bool`

HasBypassBusyWait returns a boolean if a field has been set.

### GetDefaultHost

`func (o *RequestSettingsResponse) GetDefaultHost() string`

GetDefaultHost returns the DefaultHost field if non-nil, zero value otherwise.

### GetDefaultHostOk

`func (o *RequestSettingsResponse) GetDefaultHostOk() (*string, bool)`

GetDefaultHostOk returns a tuple with the DefaultHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHost

`func (o *RequestSettingsResponse) SetDefaultHost(v string)`

SetDefaultHost sets DefaultHost field to given value.

### HasDefaultHost

`func (o *RequestSettingsResponse) HasDefaultHost() bool`

HasDefaultHost returns a boolean if a field has been set.

### SetDefaultHostNil

`func (o *RequestSettingsResponse) SetDefaultHostNil(b bool)`

 SetDefaultHostNil sets the value for DefaultHost to be an explicit nil

### UnsetDefaultHost
`func (o *RequestSettingsResponse) UnsetDefaultHost()`

UnsetDefaultHost ensures that no value is present for DefaultHost, not even an explicit nil
### GetForceMiss

`func (o *RequestSettingsResponse) GetForceMiss() int32`

GetForceMiss returns the ForceMiss field if non-nil, zero value otherwise.

### GetForceMissOk

`func (o *RequestSettingsResponse) GetForceMissOk() (*int32, bool)`

GetForceMissOk returns a tuple with the ForceMiss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceMiss

`func (o *RequestSettingsResponse) SetForceMiss(v int32)`

SetForceMiss sets ForceMiss field to given value.

### HasForceMiss

`func (o *RequestSettingsResponse) HasForceMiss() bool`

HasForceMiss returns a boolean if a field has been set.

### GetForceSsl

`func (o *RequestSettingsResponse) GetForceSsl() int32`

GetForceSsl returns the ForceSsl field if non-nil, zero value otherwise.

### GetForceSslOk

`func (o *RequestSettingsResponse) GetForceSslOk() (*int32, bool)`

GetForceSslOk returns a tuple with the ForceSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSsl

`func (o *RequestSettingsResponse) SetForceSsl(v int32)`

SetForceSsl sets ForceSsl field to given value.

### HasForceSsl

`func (o *RequestSettingsResponse) HasForceSsl() bool`

HasForceSsl returns a boolean if a field has been set.

### GetGeoHeaders

`func (o *RequestSettingsResponse) GetGeoHeaders() int32`

GetGeoHeaders returns the GeoHeaders field if non-nil, zero value otherwise.

### GetGeoHeadersOk

`func (o *RequestSettingsResponse) GetGeoHeadersOk() (*int32, bool)`

GetGeoHeadersOk returns a tuple with the GeoHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoHeaders

`func (o *RequestSettingsResponse) SetGeoHeaders(v int32)`

SetGeoHeaders sets GeoHeaders field to given value.

### HasGeoHeaders

`func (o *RequestSettingsResponse) HasGeoHeaders() bool`

HasGeoHeaders returns a boolean if a field has been set.

### GetHashKeys

`func (o *RequestSettingsResponse) GetHashKeys() string`

GetHashKeys returns the HashKeys field if non-nil, zero value otherwise.

### GetHashKeysOk

`func (o *RequestSettingsResponse) GetHashKeysOk() (*string, bool)`

GetHashKeysOk returns a tuple with the HashKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashKeys

`func (o *RequestSettingsResponse) SetHashKeys(v string)`

SetHashKeys sets HashKeys field to given value.

### HasHashKeys

`func (o *RequestSettingsResponse) HasHashKeys() bool`

HasHashKeys returns a boolean if a field has been set.

### SetHashKeysNil

`func (o *RequestSettingsResponse) SetHashKeysNil(b bool)`

 SetHashKeysNil sets the value for HashKeys to be an explicit nil

### UnsetHashKeys
`func (o *RequestSettingsResponse) UnsetHashKeys()`

UnsetHashKeys ensures that no value is present for HashKeys, not even an explicit nil
### GetMaxStaleAge

`func (o *RequestSettingsResponse) GetMaxStaleAge() int32`

GetMaxStaleAge returns the MaxStaleAge field if non-nil, zero value otherwise.

### GetMaxStaleAgeOk

`func (o *RequestSettingsResponse) GetMaxStaleAgeOk() (*int32, bool)`

GetMaxStaleAgeOk returns a tuple with the MaxStaleAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaleAge

`func (o *RequestSettingsResponse) SetMaxStaleAge(v int32)`

SetMaxStaleAge sets MaxStaleAge field to given value.

### HasMaxStaleAge

`func (o *RequestSettingsResponse) HasMaxStaleAge() bool`

HasMaxStaleAge returns a boolean if a field has been set.

### GetName

`func (o *RequestSettingsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestSettingsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestSettingsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestSettingsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequestCondition

`func (o *RequestSettingsResponse) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *RequestSettingsResponse) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *RequestSettingsResponse) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *RequestSettingsResponse) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *RequestSettingsResponse) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *RequestSettingsResponse) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetTimerSupport

`func (o *RequestSettingsResponse) GetTimerSupport() int32`

GetTimerSupport returns the TimerSupport field if non-nil, zero value otherwise.

### GetTimerSupportOk

`func (o *RequestSettingsResponse) GetTimerSupportOk() (*int32, bool)`

GetTimerSupportOk returns a tuple with the TimerSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerSupport

`func (o *RequestSettingsResponse) SetTimerSupport(v int32)`

SetTimerSupport sets TimerSupport field to given value.

### HasTimerSupport

`func (o *RequestSettingsResponse) HasTimerSupport() bool`

HasTimerSupport returns a boolean if a field has been set.

### GetXff

`func (o *RequestSettingsResponse) GetXff() string`

GetXff returns the Xff field if non-nil, zero value otherwise.

### GetXffOk

`func (o *RequestSettingsResponse) GetXffOk() (*string, bool)`

GetXffOk returns a tuple with the Xff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXff

`func (o *RequestSettingsResponse) SetXff(v string)`

SetXff sets Xff field to given value.

### HasXff

`func (o *RequestSettingsResponse) HasXff() bool`

HasXff returns a boolean if a field has been set.

### GetServiceID

`func (o *RequestSettingsResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *RequestSettingsResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *RequestSettingsResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *RequestSettingsResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *RequestSettingsResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RequestSettingsResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RequestSettingsResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RequestSettingsResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RequestSettingsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequestSettingsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequestSettingsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RequestSettingsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *RequestSettingsResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *RequestSettingsResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *RequestSettingsResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RequestSettingsResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RequestSettingsResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RequestSettingsResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *RequestSettingsResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *RequestSettingsResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *RequestSettingsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RequestSettingsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RequestSettingsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RequestSettingsResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *RequestSettingsResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *RequestSettingsResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
