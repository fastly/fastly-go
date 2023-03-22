# CacheSettingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.  | [optional] 
**CacheCondition** | Pointer to **NullableString** | Name of the cache condition controlling when this configuration applies. | [optional] 
**Name** | Pointer to **string** | Name for the cache settings object. | [optional] 
**StaleTTL** | Pointer to **int32** | Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as &#39;stale if error&#39;). | [optional] 
**TTL** | Pointer to **int32** | Maximum time to consider the object fresh in the cache (the cache &#39;time to live&#39;). | [optional] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **int32** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 

## Methods

### NewCacheSettingResponse

`func NewCacheSettingResponse() *CacheSettingResponse`

NewCacheSettingResponse instantiates a new CacheSettingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingResponseWithDefaults

`func NewCacheSettingResponseWithDefaults() *CacheSettingResponse`

NewCacheSettingResponseWithDefaults instantiates a new CacheSettingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CacheSettingResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CacheSettingResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CacheSettingResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CacheSettingResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *CacheSettingResponse) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CacheSettingResponse) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetCacheCondition

`func (o *CacheSettingResponse) GetCacheCondition() string`

GetCacheCondition returns the CacheCondition field if non-nil, zero value otherwise.

### GetCacheConditionOk

`func (o *CacheSettingResponse) GetCacheConditionOk() (*string, bool)`

GetCacheConditionOk returns a tuple with the CacheCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCondition

`func (o *CacheSettingResponse) SetCacheCondition(v string)`

SetCacheCondition sets CacheCondition field to given value.

### HasCacheCondition

`func (o *CacheSettingResponse) HasCacheCondition() bool`

HasCacheCondition returns a boolean if a field has been set.

### SetCacheConditionNil

`func (o *CacheSettingResponse) SetCacheConditionNil(b bool)`

 SetCacheConditionNil sets the value for CacheCondition to be an explicit nil

### UnsetCacheCondition
`func (o *CacheSettingResponse) UnsetCacheCondition()`

UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
### GetName

`func (o *CacheSettingResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CacheSettingResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CacheSettingResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CacheSettingResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStaleTTL

`func (o *CacheSettingResponse) GetStaleTTL() int32`

GetStaleTTL returns the StaleTTL field if non-nil, zero value otherwise.

### GetStaleTTLOk

`func (o *CacheSettingResponse) GetStaleTTLOk() (*int32, bool)`

GetStaleTTLOk returns a tuple with the StaleTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleTTL

`func (o *CacheSettingResponse) SetStaleTTL(v int32)`

SetStaleTTL sets StaleTTL field to given value.

### HasStaleTTL

`func (o *CacheSettingResponse) HasStaleTTL() bool`

HasStaleTTL returns a boolean if a field has been set.

### GetTTL

`func (o *CacheSettingResponse) GetTTL() int32`

GetTTL returns the TTL field if non-nil, zero value otherwise.

### GetTTLOk

`func (o *CacheSettingResponse) GetTTLOk() (*int32, bool)`

GetTTLOk returns a tuple with the TTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTL

`func (o *CacheSettingResponse) SetTTL(v int32)`

SetTTL sets TTL field to given value.

### HasTTL

`func (o *CacheSettingResponse) HasTTL() bool`

HasTTL returns a boolean if a field has been set.

### GetServiceID

`func (o *CacheSettingResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *CacheSettingResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *CacheSettingResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *CacheSettingResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *CacheSettingResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CacheSettingResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CacheSettingResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CacheSettingResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CacheSettingResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CacheSettingResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CacheSettingResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CacheSettingResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CacheSettingResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CacheSettingResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *CacheSettingResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CacheSettingResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CacheSettingResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CacheSettingResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *CacheSettingResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *CacheSettingResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *CacheSettingResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CacheSettingResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CacheSettingResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CacheSettingResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CacheSettingResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CacheSettingResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
