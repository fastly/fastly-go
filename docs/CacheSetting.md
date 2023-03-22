# CacheSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.  | [optional] 
**CacheCondition** | Pointer to **NullableString** | Name of the cache condition controlling when this configuration applies. | [optional] 
**Name** | Pointer to **string** | Name for the cache settings object. | [optional] 
**StaleTTL** | Pointer to **int32** | Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as &#39;stale if error&#39;). | [optional] 
**TTL** | Pointer to **int32** | Maximum time to consider the object fresh in the cache (the cache &#39;time to live&#39;). | [optional] 

## Methods

### NewCacheSetting

`func NewCacheSetting() *CacheSetting`

NewCacheSetting instantiates a new CacheSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingWithDefaults

`func NewCacheSettingWithDefaults() *CacheSetting`

NewCacheSettingWithDefaults instantiates a new CacheSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CacheSetting) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CacheSetting) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CacheSetting) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CacheSetting) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *CacheSetting) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CacheSetting) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetCacheCondition

`func (o *CacheSetting) GetCacheCondition() string`

GetCacheCondition returns the CacheCondition field if non-nil, zero value otherwise.

### GetCacheConditionOk

`func (o *CacheSetting) GetCacheConditionOk() (*string, bool)`

GetCacheConditionOk returns a tuple with the CacheCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCondition

`func (o *CacheSetting) SetCacheCondition(v string)`

SetCacheCondition sets CacheCondition field to given value.

### HasCacheCondition

`func (o *CacheSetting) HasCacheCondition() bool`

HasCacheCondition returns a boolean if a field has been set.

### SetCacheConditionNil

`func (o *CacheSetting) SetCacheConditionNil(b bool)`

 SetCacheConditionNil sets the value for CacheCondition to be an explicit nil

### UnsetCacheCondition
`func (o *CacheSetting) UnsetCacheCondition()`

UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
### GetName

`func (o *CacheSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CacheSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CacheSetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CacheSetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStaleTTL

`func (o *CacheSetting) GetStaleTTL() int32`

GetStaleTTL returns the StaleTTL field if non-nil, zero value otherwise.

### GetStaleTTLOk

`func (o *CacheSetting) GetStaleTTLOk() (*int32, bool)`

GetStaleTTLOk returns a tuple with the StaleTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleTTL

`func (o *CacheSetting) SetStaleTTL(v int32)`

SetStaleTTL sets StaleTTL field to given value.

### HasStaleTTL

`func (o *CacheSetting) HasStaleTTL() bool`

HasStaleTTL returns a boolean if a field has been set.

### GetTTL

`func (o *CacheSetting) GetTTL() int32`

GetTTL returns the TTL field if non-nil, zero value otherwise.

### GetTTLOk

`func (o *CacheSetting) GetTTLOk() (*int32, bool)`

GetTTLOk returns a tuple with the TTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTL

`func (o *CacheSetting) SetTTL(v int32)`

SetTTL sets TTL field to given value.

### HasTTL

`func (o *CacheSetting) HasTTL() bool`

HasTTL returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
