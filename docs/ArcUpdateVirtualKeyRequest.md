# ArcUpdateVirtualKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A display name for the virtual key. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | The date and time the key expires, in RFC 3339 format. Set explicitly to `null` to clear the expiration.  | [optional] 
**SecurityEnabled** | Pointer to **bool** | Whether AI Firewall security is enabled for requests made with this key. | [optional] 
**SecurityAction** | Pointer to **string** | The action AI Firewall takes when it detects a violation on a request made with this key. | [optional] 
**RpmLimit** | Pointer to **NullableInt32** | The requests-per-minute limit for this key. Omit this field to leave the limit unchanged, set it to a positive integer to set a limit, or set it explicitly to `null` to remove the limit.  | [optional] 
**TpmLimit** | Pointer to **NullableInt32** | The tokens-per-minute limit for this key. Omit this field to leave the limit unchanged, set it to a positive integer to set a limit, or set it explicitly to `null` to remove the limit.  | [optional] 
**SpendLimitEnabled** | Pointer to **bool** | Whether a spend limit is enabled for this key. This field cannot be set when the key is created; it can only be changed with this endpoint. | [optional] 

## Methods

### NewArcUpdateVirtualKeyRequest

`func NewArcUpdateVirtualKeyRequest() *ArcUpdateVirtualKeyRequest`

NewArcUpdateVirtualKeyRequest instantiates a new ArcUpdateVirtualKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcUpdateVirtualKeyRequestWithDefaults

`func NewArcUpdateVirtualKeyRequestWithDefaults() *ArcUpdateVirtualKeyRequest`

NewArcUpdateVirtualKeyRequestWithDefaults instantiates a new ArcUpdateVirtualKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArcUpdateVirtualKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcUpdateVirtualKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcUpdateVirtualKeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArcUpdateVirtualKeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ArcUpdateVirtualKeyRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ArcUpdateVirtualKeyRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ArcUpdateVirtualKeyRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ArcUpdateVirtualKeyRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ArcUpdateVirtualKeyRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ArcUpdateVirtualKeyRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetSecurityEnabled

`func (o *ArcUpdateVirtualKeyRequest) GetSecurityEnabled() bool`

GetSecurityEnabled returns the SecurityEnabled field if non-nil, zero value otherwise.

### GetSecurityEnabledOk

`func (o *ArcUpdateVirtualKeyRequest) GetSecurityEnabledOk() (*bool, bool)`

GetSecurityEnabledOk returns a tuple with the SecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEnabled

`func (o *ArcUpdateVirtualKeyRequest) SetSecurityEnabled(v bool)`

SetSecurityEnabled sets SecurityEnabled field to given value.

### HasSecurityEnabled

`func (o *ArcUpdateVirtualKeyRequest) HasSecurityEnabled() bool`

HasSecurityEnabled returns a boolean if a field has been set.

### GetSecurityAction

`func (o *ArcUpdateVirtualKeyRequest) GetSecurityAction() string`

GetSecurityAction returns the SecurityAction field if non-nil, zero value otherwise.

### GetSecurityActionOk

`func (o *ArcUpdateVirtualKeyRequest) GetSecurityActionOk() (*string, bool)`

GetSecurityActionOk returns a tuple with the SecurityAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAction

`func (o *ArcUpdateVirtualKeyRequest) SetSecurityAction(v string)`

SetSecurityAction sets SecurityAction field to given value.

### HasSecurityAction

`func (o *ArcUpdateVirtualKeyRequest) HasSecurityAction() bool`

HasSecurityAction returns a boolean if a field has been set.

### GetRpmLimit

`func (o *ArcUpdateVirtualKeyRequest) GetRpmLimit() int32`

GetRpmLimit returns the RpmLimit field if non-nil, zero value otherwise.

### GetRpmLimitOk

`func (o *ArcUpdateVirtualKeyRequest) GetRpmLimitOk() (*int32, bool)`

GetRpmLimitOk returns a tuple with the RpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmLimit

`func (o *ArcUpdateVirtualKeyRequest) SetRpmLimit(v int32)`

SetRpmLimit sets RpmLimit field to given value.

### HasRpmLimit

`func (o *ArcUpdateVirtualKeyRequest) HasRpmLimit() bool`

HasRpmLimit returns a boolean if a field has been set.

### SetRpmLimitNil

`func (o *ArcUpdateVirtualKeyRequest) SetRpmLimitNil(b bool)`

 SetRpmLimitNil sets the value for RpmLimit to be an explicit nil

### UnsetRpmLimit
`func (o *ArcUpdateVirtualKeyRequest) UnsetRpmLimit()`

UnsetRpmLimit ensures that no value is present for RpmLimit, not even an explicit nil
### GetTpmLimit

`func (o *ArcUpdateVirtualKeyRequest) GetTpmLimit() int32`

GetTpmLimit returns the TpmLimit field if non-nil, zero value otherwise.

### GetTpmLimitOk

`func (o *ArcUpdateVirtualKeyRequest) GetTpmLimitOk() (*int32, bool)`

GetTpmLimitOk returns a tuple with the TpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmLimit

`func (o *ArcUpdateVirtualKeyRequest) SetTpmLimit(v int32)`

SetTpmLimit sets TpmLimit field to given value.

### HasTpmLimit

`func (o *ArcUpdateVirtualKeyRequest) HasTpmLimit() bool`

HasTpmLimit returns a boolean if a field has been set.

### SetTpmLimitNil

`func (o *ArcUpdateVirtualKeyRequest) SetTpmLimitNil(b bool)`

 SetTpmLimitNil sets the value for TpmLimit to be an explicit nil

### UnsetTpmLimit
`func (o *ArcUpdateVirtualKeyRequest) UnsetTpmLimit()`

UnsetTpmLimit ensures that no value is present for TpmLimit, not even an explicit nil
### GetSpendLimitEnabled

`func (o *ArcUpdateVirtualKeyRequest) GetSpendLimitEnabled() bool`

GetSpendLimitEnabled returns the SpendLimitEnabled field if non-nil, zero value otherwise.

### GetSpendLimitEnabledOk

`func (o *ArcUpdateVirtualKeyRequest) GetSpendLimitEnabledOk() (*bool, bool)`

GetSpendLimitEnabledOk returns a tuple with the SpendLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitEnabled

`func (o *ArcUpdateVirtualKeyRequest) SetSpendLimitEnabled(v bool)`

SetSpendLimitEnabled sets SpendLimitEnabled field to given value.

### HasSpendLimitEnabled

`func (o *ArcUpdateVirtualKeyRequest) HasSpendLimitEnabled() bool`

HasSpendLimitEnabled returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
