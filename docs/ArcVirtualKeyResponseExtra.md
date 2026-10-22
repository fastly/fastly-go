# ArcVirtualKeyResponseExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the virtual key. | [readonly] 
**Type** | **string** | The type of the resource. | [readonly] 
**UserId** | **string** | The ID of the user who created the key. | [readonly] 
**CreatedBy** | Pointer to **string** | The display name of the user who created the key. | [optional] [readonly] 
**LastUsedAt** | Pointer to **NullableTime** | The date and time the key was last used to make a request, in RFC 3339 format. `null` if the key has never been used. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | The date and time the key was deleted, in RFC 3339 format. `null` unless the key was deleted and `include_deleted` was used to retrieve it. | [optional] [readonly] 
**CreatedAt** | **time.Time** | The date and time the key was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the key was last updated, in RFC 3339 format. | [readonly] 
**SpendLimitEnabled** | Pointer to **bool** | Whether a spend limit is enabled for this key. | [optional] [readonly] 
**SpendLimitEnabledOn** | Pointer to **NullableTime** | The date and time `spend_limit_enabled` was last set to `true`, in RFC 3339 format. `null` if the spend limit has never been enabled. This timestamp is not cleared when `spend_limit_enabled` is later set back to `false`, so it can be non-null even when the spend limit is currently disabled.  | [optional] [readonly] 

## Methods

### NewArcVirtualKeyResponseExtra

`func NewArcVirtualKeyResponseExtra(id string, type_ string, userId string, createdAt time.Time, updatedAt time.Time, ) *ArcVirtualKeyResponseExtra`

NewArcVirtualKeyResponseExtra instantiates a new ArcVirtualKeyResponseExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcVirtualKeyResponseExtraWithDefaults

`func NewArcVirtualKeyResponseExtraWithDefaults() *ArcVirtualKeyResponseExtra`

NewArcVirtualKeyResponseExtraWithDefaults instantiates a new ArcVirtualKeyResponseExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArcVirtualKeyResponseExtra) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArcVirtualKeyResponseExtra) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArcVirtualKeyResponseExtra) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ArcVirtualKeyResponseExtra) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArcVirtualKeyResponseExtra) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArcVirtualKeyResponseExtra) SetType(v string)`

SetType sets Type field to given value.


### GetUserId

`func (o *ArcVirtualKeyResponseExtra) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ArcVirtualKeyResponseExtra) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ArcVirtualKeyResponseExtra) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedBy

`func (o *ArcVirtualKeyResponseExtra) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArcVirtualKeyResponseExtra) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArcVirtualKeyResponseExtra) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArcVirtualKeyResponseExtra) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ArcVirtualKeyResponseExtra) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ArcVirtualKeyResponseExtra) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ArcVirtualKeyResponseExtra) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ArcVirtualKeyResponseExtra) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *ArcVirtualKeyResponseExtra) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *ArcVirtualKeyResponseExtra) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetDeletedAt

`func (o *ArcVirtualKeyResponseExtra) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArcVirtualKeyResponseExtra) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArcVirtualKeyResponseExtra) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArcVirtualKeyResponseExtra) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ArcVirtualKeyResponseExtra) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ArcVirtualKeyResponseExtra) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetCreatedAt

`func (o *ArcVirtualKeyResponseExtra) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArcVirtualKeyResponseExtra) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArcVirtualKeyResponseExtra) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArcVirtualKeyResponseExtra) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArcVirtualKeyResponseExtra) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArcVirtualKeyResponseExtra) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSpendLimitEnabled

`func (o *ArcVirtualKeyResponseExtra) GetSpendLimitEnabled() bool`

GetSpendLimitEnabled returns the SpendLimitEnabled field if non-nil, zero value otherwise.

### GetSpendLimitEnabledOk

`func (o *ArcVirtualKeyResponseExtra) GetSpendLimitEnabledOk() (*bool, bool)`

GetSpendLimitEnabledOk returns a tuple with the SpendLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitEnabled

`func (o *ArcVirtualKeyResponseExtra) SetSpendLimitEnabled(v bool)`

SetSpendLimitEnabled sets SpendLimitEnabled field to given value.

### HasSpendLimitEnabled

`func (o *ArcVirtualKeyResponseExtra) HasSpendLimitEnabled() bool`

HasSpendLimitEnabled returns a boolean if a field has been set.

### GetSpendLimitEnabledOn

`func (o *ArcVirtualKeyResponseExtra) GetSpendLimitEnabledOn() time.Time`

GetSpendLimitEnabledOn returns the SpendLimitEnabledOn field if non-nil, zero value otherwise.

### GetSpendLimitEnabledOnOk

`func (o *ArcVirtualKeyResponseExtra) GetSpendLimitEnabledOnOk() (*time.Time, bool)`

GetSpendLimitEnabledOnOk returns a tuple with the SpendLimitEnabledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitEnabledOn

`func (o *ArcVirtualKeyResponseExtra) SetSpendLimitEnabledOn(v time.Time)`

SetSpendLimitEnabledOn sets SpendLimitEnabledOn field to given value.

### HasSpendLimitEnabledOn

`func (o *ArcVirtualKeyResponseExtra) HasSpendLimitEnabledOn() bool`

HasSpendLimitEnabledOn returns a boolean if a field has been set.

### SetSpendLimitEnabledOnNil

`func (o *ArcVirtualKeyResponseExtra) SetSpendLimitEnabledOnNil(b bool)`

 SetSpendLimitEnabledOnNil sets the value for SpendLimitEnabledOn to be an explicit nil

### UnsetSpendLimitEnabledOn
`func (o *ArcVirtualKeyResponseExtra) UnsetSpendLimitEnabledOn()`

UnsetSpendLimitEnabledOn ensures that no value is present for SpendLimitEnabledOn, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


