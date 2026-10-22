# ArcVirtualKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A display name for the virtual key. | 
**ExpiresAt** | Pointer to **NullableTime** | The date and time the key expires, in RFC 3339 format. Omit for a key that doesn&#39;t expire. | [optional] 
**SecurityEnabled** | Pointer to **bool** | Whether AI Firewall security is enabled for requests made with this key. | [optional] [default to false]
**SecurityAction** | Pointer to **string** | The action AI Firewall takes when it detects a violation on a request made with this key. Only meaningful when `security_enabled` is `true`. | [optional] 
**RpmLimit** | Pointer to **NullableInt32** | The requests-per-minute limit for this key. Omit for no limit. | [optional] 
**TpmLimit** | Pointer to **NullableInt32** | The tokens-per-minute limit for this key. Omit for no limit. | [optional] 

## Methods

### NewArcVirtualKey

`func NewArcVirtualKey(name string, ) *ArcVirtualKey`

NewArcVirtualKey instantiates a new ArcVirtualKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcVirtualKeyWithDefaults

`func NewArcVirtualKeyWithDefaults() *ArcVirtualKey`

NewArcVirtualKeyWithDefaults instantiates a new ArcVirtualKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArcVirtualKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcVirtualKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcVirtualKey) SetName(v string)`

SetName sets Name field to given value.


### GetExpiresAt

`func (o *ArcVirtualKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ArcVirtualKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ArcVirtualKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ArcVirtualKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ArcVirtualKey) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ArcVirtualKey) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetSecurityEnabled

`func (o *ArcVirtualKey) GetSecurityEnabled() bool`

GetSecurityEnabled returns the SecurityEnabled field if non-nil, zero value otherwise.

### GetSecurityEnabledOk

`func (o *ArcVirtualKey) GetSecurityEnabledOk() (*bool, bool)`

GetSecurityEnabledOk returns a tuple with the SecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEnabled

`func (o *ArcVirtualKey) SetSecurityEnabled(v bool)`

SetSecurityEnabled sets SecurityEnabled field to given value.

### HasSecurityEnabled

`func (o *ArcVirtualKey) HasSecurityEnabled() bool`

HasSecurityEnabled returns a boolean if a field has been set.

### GetSecurityAction

`func (o *ArcVirtualKey) GetSecurityAction() string`

GetSecurityAction returns the SecurityAction field if non-nil, zero value otherwise.

### GetSecurityActionOk

`func (o *ArcVirtualKey) GetSecurityActionOk() (*string, bool)`

GetSecurityActionOk returns a tuple with the SecurityAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAction

`func (o *ArcVirtualKey) SetSecurityAction(v string)`

SetSecurityAction sets SecurityAction field to given value.

### HasSecurityAction

`func (o *ArcVirtualKey) HasSecurityAction() bool`

HasSecurityAction returns a boolean if a field has been set.

### GetRpmLimit

`func (o *ArcVirtualKey) GetRpmLimit() int32`

GetRpmLimit returns the RpmLimit field if non-nil, zero value otherwise.

### GetRpmLimitOk

`func (o *ArcVirtualKey) GetRpmLimitOk() (*int32, bool)`

GetRpmLimitOk returns a tuple with the RpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmLimit

`func (o *ArcVirtualKey) SetRpmLimit(v int32)`

SetRpmLimit sets RpmLimit field to given value.

### HasRpmLimit

`func (o *ArcVirtualKey) HasRpmLimit() bool`

HasRpmLimit returns a boolean if a field has been set.

### SetRpmLimitNil

`func (o *ArcVirtualKey) SetRpmLimitNil(b bool)`

 SetRpmLimitNil sets the value for RpmLimit to be an explicit nil

### UnsetRpmLimit
`func (o *ArcVirtualKey) UnsetRpmLimit()`

UnsetRpmLimit ensures that no value is present for RpmLimit, not even an explicit nil
### GetTpmLimit

`func (o *ArcVirtualKey) GetTpmLimit() int32`

GetTpmLimit returns the TpmLimit field if non-nil, zero value otherwise.

### GetTpmLimitOk

`func (o *ArcVirtualKey) GetTpmLimitOk() (*int32, bool)`

GetTpmLimitOk returns a tuple with the TpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmLimit

`func (o *ArcVirtualKey) SetTpmLimit(v int32)`

SetTpmLimit sets TpmLimit field to given value.

### HasTpmLimit

`func (o *ArcVirtualKey) HasTpmLimit() bool`

HasTpmLimit returns a boolean if a field has been set.

### SetTpmLimitNil

`func (o *ArcVirtualKey) SetTpmLimitNil(b bool)`

 SetTpmLimitNil sets the value for TpmLimit to be an explicit nil

### UnsetTpmLimit
`func (o *ArcVirtualKey) UnsetTpmLimit()`

UnsetTpmLimit ensures that no value is present for TpmLimit, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


