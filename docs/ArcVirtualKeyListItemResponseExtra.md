# ArcVirtualKeyListItemResponseExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | The unique identifier of the virtual key. | [readonly] 
**CustomerID** | **string** | The ID of the customer account that owns the key. | [readonly] 
**Type** | **string** | The type of the resource. | [readonly] 
**UserID** | **string** | The ID of the user who created the key. | [readonly] 
**UserName** | Pointer to **string** | The display name of the user who created the key. | [optional] [readonly] 
**LastUsedAt** | Pointer to **NullableTime** | The date and time the key was last used to make a request, in RFC 3339 format. `null` if the key has never been used. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | The date and time the key was deleted, in RFC 3339 format. `null` unless the key was deleted and `include_deleted` was used to retrieve it. | [optional] [readonly] 
**CreatedAt** | **time.Time** | The date and time the key was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the key was last updated, in RFC 3339 format. | [readonly] 
**SpendLimitEnabled** | Pointer to **bool** | Whether a spend limit is enabled for this key. | [optional] [readonly] 
**SpendLimitEnabledOn** | Pointer to **NullableTime** | The date and time `spend_limit_enabled` was last set to `true`, in RFC 3339 format. `null` if the spend limit has never been enabled. This timestamp is not cleared when `spend_limit_enabled` is later set back to `false`, so it can be non-null even when the spend limit is currently disabled.  | [optional] [readonly] 

## Methods

### NewArcVirtualKeyListItemResponseExtra

`func NewArcVirtualKeyListItemResponseExtra(id string, customerID string, resourceType string, userID string, createdAt time.Time, updatedAt time.Time, ) *ArcVirtualKeyListItemResponseExtra`

NewArcVirtualKeyListItemResponseExtra instantiates a new ArcVirtualKeyListItemResponseExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcVirtualKeyListItemResponseExtraWithDefaults

`func NewArcVirtualKeyListItemResponseExtraWithDefaults() *ArcVirtualKeyListItemResponseExtra`

NewArcVirtualKeyListItemResponseExtraWithDefaults instantiates a new ArcVirtualKeyListItemResponseExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ArcVirtualKeyListItemResponseExtra) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ArcVirtualKeyListItemResponseExtra) SetID(v string)`

SetID sets ID field to given value.


### GetCustomerID

`func (o *ArcVirtualKeyListItemResponseExtra) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *ArcVirtualKeyListItemResponseExtra) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.


### GetType

`func (o *ArcVirtualKeyListItemResponseExtra) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArcVirtualKeyListItemResponseExtra) SetType(v string)`

SetType sets Type field to given value.


### GetUserID

`func (o *ArcVirtualKeyListItemResponseExtra) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *ArcVirtualKeyListItemResponseExtra) SetUserID(v string)`

SetUserID sets UserID field to given value.


### GetUserName

`func (o *ArcVirtualKeyListItemResponseExtra) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ArcVirtualKeyListItemResponseExtra) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ArcVirtualKeyListItemResponseExtra) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ArcVirtualKeyListItemResponseExtra) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ArcVirtualKeyListItemResponseExtra) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ArcVirtualKeyListItemResponseExtra) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *ArcVirtualKeyListItemResponseExtra) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *ArcVirtualKeyListItemResponseExtra) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetDeletedAt

`func (o *ArcVirtualKeyListItemResponseExtra) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArcVirtualKeyListItemResponseExtra) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArcVirtualKeyListItemResponseExtra) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ArcVirtualKeyListItemResponseExtra) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ArcVirtualKeyListItemResponseExtra) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetCreatedAt

`func (o *ArcVirtualKeyListItemResponseExtra) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArcVirtualKeyListItemResponseExtra) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArcVirtualKeyListItemResponseExtra) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArcVirtualKeyListItemResponseExtra) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSpendLimitEnabled

`func (o *ArcVirtualKeyListItemResponseExtra) GetSpendLimitEnabled() bool`

GetSpendLimitEnabled returns the SpendLimitEnabled field if non-nil, zero value otherwise.

### GetSpendLimitEnabledOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetSpendLimitEnabledOk() (*bool, bool)`

GetSpendLimitEnabledOk returns a tuple with the SpendLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitEnabled

`func (o *ArcVirtualKeyListItemResponseExtra) SetSpendLimitEnabled(v bool)`

SetSpendLimitEnabled sets SpendLimitEnabled field to given value.

### HasSpendLimitEnabled

`func (o *ArcVirtualKeyListItemResponseExtra) HasSpendLimitEnabled() bool`

HasSpendLimitEnabled returns a boolean if a field has been set.

### GetSpendLimitEnabledOn

`func (o *ArcVirtualKeyListItemResponseExtra) GetSpendLimitEnabledOn() time.Time`

GetSpendLimitEnabledOn returns the SpendLimitEnabledOn field if non-nil, zero value otherwise.

### GetSpendLimitEnabledOnOk

`func (o *ArcVirtualKeyListItemResponseExtra) GetSpendLimitEnabledOnOk() (*time.Time, bool)`

GetSpendLimitEnabledOnOk returns a tuple with the SpendLimitEnabledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitEnabledOn

`func (o *ArcVirtualKeyListItemResponseExtra) SetSpendLimitEnabledOn(v time.Time)`

SetSpendLimitEnabledOn sets SpendLimitEnabledOn field to given value.

### HasSpendLimitEnabledOn

`func (o *ArcVirtualKeyListItemResponseExtra) HasSpendLimitEnabledOn() bool`

HasSpendLimitEnabledOn returns a boolean if a field has been set.

### SetSpendLimitEnabledOnNil

`func (o *ArcVirtualKeyListItemResponseExtra) SetSpendLimitEnabledOnNil(b bool)`

 SetSpendLimitEnabledOnNil sets the value for SpendLimitEnabledOn to be an explicit nil

### UnsetSpendLimitEnabledOn
`func (o *ArcVirtualKeyListItemResponseExtra) UnsetSpendLimitEnabledOn()`

UnsetSpendLimitEnabledOn ensures that no value is present for SpendLimitEnabledOn, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
