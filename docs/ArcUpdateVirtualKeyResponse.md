# ArcUpdateVirtualKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the virtual key. | [readonly] 
**CustomerId** | **string** | The ID of the customer account that owns the key. | [readonly] 
**Name** | **string** | The display name of the virtual key. | 
**UserId** | **string** | The ID of the user who created the key. | [readonly] 
**ExpiresAt** | Pointer to **NullableTime** | The date and time the key expires, in RFC 3339 format. `null` if the key doesn&#39;t expire. | [optional] 
**RpmLimit** | Pointer to **NullableInt32** | The requests-per-minute limit for this key. `null` if there is no limit. | [optional] 
**TpmLimit** | Pointer to **NullableInt32** | The tokens-per-minute limit for this key. `null` if there is no limit. | [optional] 
**SpendLimitEnabled** | Pointer to **bool** | Whether a spend limit is enabled for this key. | [optional] [readonly] 
**SpendLimitEnabledOn** | Pointer to **NullableTime** | The date and time `spend_limit_enabled` was last set to `true`, in RFC 3339 format. `null` if the spend limit has never been enabled. This timestamp is not cleared when `spend_limit_enabled` is later set back to `false`, so it can be non-null even when the spend limit is currently disabled.  | [optional] [readonly] 
**CreatedAt** | **time.Time** | The date and time the key was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the key was last updated, in RFC 3339 format. | [readonly] 

## Methods

### NewArcUpdateVirtualKeyResponse

`func NewArcUpdateVirtualKeyResponse(id string, customerId string, name string, userId string, createdAt time.Time, updatedAt time.Time, ) *ArcUpdateVirtualKeyResponse`

NewArcUpdateVirtualKeyResponse instantiates a new ArcUpdateVirtualKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcUpdateVirtualKeyResponseWithDefaults

`func NewArcUpdateVirtualKeyResponseWithDefaults() *ArcUpdateVirtualKeyResponse`

NewArcUpdateVirtualKeyResponseWithDefaults instantiates a new ArcUpdateVirtualKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArcUpdateVirtualKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArcUpdateVirtualKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArcUpdateVirtualKeyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *ArcUpdateVirtualKeyResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ArcUpdateVirtualKeyResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ArcUpdateVirtualKeyResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetName

`func (o *ArcUpdateVirtualKeyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcUpdateVirtualKeyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcUpdateVirtualKeyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *ArcUpdateVirtualKeyResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ArcUpdateVirtualKeyResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ArcUpdateVirtualKeyResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetExpiresAt

`func (o *ArcUpdateVirtualKeyResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ArcUpdateVirtualKeyResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ArcUpdateVirtualKeyResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ArcUpdateVirtualKeyResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ArcUpdateVirtualKeyResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ArcUpdateVirtualKeyResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRpmLimit

`func (o *ArcUpdateVirtualKeyResponse) GetRpmLimit() int32`

GetRpmLimit returns the RpmLimit field if non-nil, zero value otherwise.

### GetRpmLimitOk

`func (o *ArcUpdateVirtualKeyResponse) GetRpmLimitOk() (*int32, bool)`

GetRpmLimitOk returns a tuple with the RpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmLimit

`func (o *ArcUpdateVirtualKeyResponse) SetRpmLimit(v int32)`

SetRpmLimit sets RpmLimit field to given value.

### HasRpmLimit

`func (o *ArcUpdateVirtualKeyResponse) HasRpmLimit() bool`

HasRpmLimit returns a boolean if a field has been set.

### SetRpmLimitNil

`func (o *ArcUpdateVirtualKeyResponse) SetRpmLimitNil(b bool)`

 SetRpmLimitNil sets the value for RpmLimit to be an explicit nil

### UnsetRpmLimit
`func (o *ArcUpdateVirtualKeyResponse) UnsetRpmLimit()`

UnsetRpmLimit ensures that no value is present for RpmLimit, not even an explicit nil
### GetTpmLimit

`func (o *ArcUpdateVirtualKeyResponse) GetTpmLimit() int32`

GetTpmLimit returns the TpmLimit field if non-nil, zero value otherwise.

### GetTpmLimitOk

`func (o *ArcUpdateVirtualKeyResponse) GetTpmLimitOk() (*int32, bool)`

GetTpmLimitOk returns a tuple with the TpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmLimit

`func (o *ArcUpdateVirtualKeyResponse) SetTpmLimit(v int32)`

SetTpmLimit sets TpmLimit field to given value.

### HasTpmLimit

`func (o *ArcUpdateVirtualKeyResponse) HasTpmLimit() bool`

HasTpmLimit returns a boolean if a field has been set.

### SetTpmLimitNil

`func (o *ArcUpdateVirtualKeyResponse) SetTpmLimitNil(b bool)`

 SetTpmLimitNil sets the value for TpmLimit to be an explicit nil

### UnsetTpmLimit
`func (o *ArcUpdateVirtualKeyResponse) UnsetTpmLimit()`

UnsetTpmLimit ensures that no value is present for TpmLimit, not even an explicit nil
### GetSpendLimitEnabled

`func (o *ArcUpdateVirtualKeyResponse) GetSpendLimitEnabled() bool`

GetSpendLimitEnabled returns the SpendLimitEnabled field if non-nil, zero value otherwise.

### GetSpendLimitEnabledOk

`func (o *ArcUpdateVirtualKeyResponse) GetSpendLimitEnabledOk() (*bool, bool)`

GetSpendLimitEnabledOk returns a tuple with the SpendLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitEnabled

`func (o *ArcUpdateVirtualKeyResponse) SetSpendLimitEnabled(v bool)`

SetSpendLimitEnabled sets SpendLimitEnabled field to given value.

### HasSpendLimitEnabled

`func (o *ArcUpdateVirtualKeyResponse) HasSpendLimitEnabled() bool`

HasSpendLimitEnabled returns a boolean if a field has been set.

### GetSpendLimitEnabledOn

`func (o *ArcUpdateVirtualKeyResponse) GetSpendLimitEnabledOn() time.Time`

GetSpendLimitEnabledOn returns the SpendLimitEnabledOn field if non-nil, zero value otherwise.

### GetSpendLimitEnabledOnOk

`func (o *ArcUpdateVirtualKeyResponse) GetSpendLimitEnabledOnOk() (*time.Time, bool)`

GetSpendLimitEnabledOnOk returns a tuple with the SpendLimitEnabledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitEnabledOn

`func (o *ArcUpdateVirtualKeyResponse) SetSpendLimitEnabledOn(v time.Time)`

SetSpendLimitEnabledOn sets SpendLimitEnabledOn field to given value.

### HasSpendLimitEnabledOn

`func (o *ArcUpdateVirtualKeyResponse) HasSpendLimitEnabledOn() bool`

HasSpendLimitEnabledOn returns a boolean if a field has been set.

### SetSpendLimitEnabledOnNil

`func (o *ArcUpdateVirtualKeyResponse) SetSpendLimitEnabledOnNil(b bool)`

 SetSpendLimitEnabledOnNil sets the value for SpendLimitEnabledOn to be an explicit nil

### UnsetSpendLimitEnabledOn
`func (o *ArcUpdateVirtualKeyResponse) UnsetSpendLimitEnabledOn()`

UnsetSpendLimitEnabledOn ensures that no value is present for SpendLimitEnabledOn, not even an explicit nil
### GetCreatedAt

`func (o *ArcUpdateVirtualKeyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArcUpdateVirtualKeyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArcUpdateVirtualKeyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArcUpdateVirtualKeyResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArcUpdateVirtualKeyResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArcUpdateVirtualKeyResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


