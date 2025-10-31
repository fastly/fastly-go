# CacheSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.  | [optional] 
**CacheCondition** | Pointer to **NullableString** | Name of the cache condition controlling when this configuration applies. | [optional] 
**Name** | Pointer to **string** | Name for the cache settings object. | [optional] 
**StaleTtl** | Pointer to **string** | Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as &#39;stale if error&#39;). | [optional] 
**Ttl** | Pointer to **string** | Maximum time to consider the object fresh in the cache (the cache &#39;time to live&#39;). | [optional] 

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

### GetStaleTtl

`func (o *CacheSetting) GetStaleTtl() string`

GetStaleTtl returns the StaleTtl field if non-nil, zero value otherwise.

### GetStaleTtlOk

`func (o *CacheSetting) GetStaleTtlOk() (*string, bool)`

GetStaleTtlOk returns a tuple with the StaleTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleTtl

`func (o *CacheSetting) SetStaleTtl(v string)`

SetStaleTtl sets StaleTtl field to given value.

### HasStaleTtl

`func (o *CacheSetting) HasStaleTtl() bool`

HasStaleTtl returns a boolean if a field has been set.

### GetTtl

`func (o *CacheSetting) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CacheSetting) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CacheSetting) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CacheSetting) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


