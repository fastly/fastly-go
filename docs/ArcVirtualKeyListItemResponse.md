# ArcVirtualKeyListItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A display name for the virtual key. | 
**ExpiresAt** | Pointer to **NullableTime** | The date and time the key expires, in RFC 3339 format. Omit for a key that doesn&#39;t expire. | [optional] 
**SecurityEnabled** | Pointer to **bool** | Whether AI Firewall security is enabled for requests made with this key. | [optional] [default to false]
**SecurityAction** | Pointer to **string** | The action AI Firewall takes when it detects a violation on a request made with this key. Only meaningful when `security_enabled` is `true`. | [optional] 
**RpmLimit** | Pointer to **NullableInt32** | The requests-per-minute limit for this key. Omit for no limit. | [optional] 
**TpmLimit** | Pointer to **NullableInt32** | The tokens-per-minute limit for this key. Omit for no limit. | [optional] 
**Id** | **string** | The unique identifier of the virtual key. | [readonly] 
**CustomerId** | **string** | The ID of the customer account that owns the key. | [readonly] 
**Type** | **string** | The type of the resource. | [readonly] 
**UserId** | **string** | The ID of the user who created the key. | [readonly] 
**UserName** | Pointer to **string** | The display name of the user who created the key. | [optional] [readonly] 
**LastUsedAt** | Pointer to **NullableTime** | The date and time the key was last used to make a request, in RFC 3339 format. `null` if the key has never been used. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | The date and time the key was deleted, in RFC 3339 format. `null` unless the key was deleted and `include_deleted` was used to retrieve it. | [optional] [readonly] 
**CreatedAt** | **time.Time** | The date and time the key was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the key was last updated, in RFC 3339 format. | [readonly] 
**SpendLimitEnabled** | Pointer to **bool** | Whether a spend limit is enabled for this key. | [optional] [readonly] 
**SpendLimitEnabledOn** | Pointer to **NullableTime** | The date and time `spend_limit_enabled` was last set to `true`, in RFC 3339 format. `null` if the spend limit has never been enabled. This timestamp is not cleared when `spend_limit_enabled` is later set back to `false`, so it can be non-null even when the spend limit is currently disabled.  | [optional] [readonly] 

## Methods

### NewArcVirtualKeyListItemResponse

`func NewArcVirtualKeyListItemResponse(name string, id string, customerId string, type_ string, userId string, createdAt time.Time, updatedAt time.Time, ) *ArcVirtualKeyListItemResponse`

NewArcVirtualKeyListItemResponse instantiates a new ArcVirtualKeyListItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcVirtualKeyListItemResponseWithDefaults

`func NewArcVirtualKeyListItemResponseWithDefaults() *ArcVirtualKeyListItemResponse`

NewArcVirtualKeyListItemResponseWithDefaults instantiates a new ArcVirtualKeyListItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArcVirtualKeyListItemResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcVirtualKeyListItemResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcVirtualKeyListItemResponse) SetName(v string)`

SetName sets Name field to given value.


### GetExpiresAt

`func (o *ArcVirtualKeyListItemResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ArcVirtualKeyListItemResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ArcVirtualKeyListItemResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ArcVirtualKeyListItemResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ArcVirtualKeyListItemResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ArcVirtualKeyListItemResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetSecurityEnabled

`func (o *ArcVirtualKeyListItemResponse) GetSecurityEnabled() bool`

GetSecurityEnabled returns the SecurityEnabled field if non-nil, zero value otherwise.

### GetSecurityEnabledOk

`func (o *ArcVirtualKeyListItemResponse) GetSecurityEnabledOk() (*bool, bool)`

GetSecurityEnabledOk returns a tuple with the SecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEnabled

`func (o *ArcVirtualKeyListItemResponse) SetSecurityEnabled(v bool)`

SetSecurityEnabled sets SecurityEnabled field to given value.

### HasSecurityEnabled

`func (o *ArcVirtualKeyListItemResponse) HasSecurityEnabled() bool`

HasSecurityEnabled returns a boolean if a field has been set.

### GetSecurityAction

`func (o *ArcVirtualKeyListItemResponse) GetSecurityAction() string`

GetSecurityAction returns the SecurityAction field if non-nil, zero value otherwise.

### GetSecurityActionOk

`func (o *ArcVirtualKeyListItemResponse) GetSecurityActionOk() (*string, bool)`

GetSecurityActionOk returns a tuple with the SecurityAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAction

`func (o *ArcVirtualKeyListItemResponse) SetSecurityAction(v string)`

SetSecurityAction sets SecurityAction field to given value.

### HasSecurityAction

`func (o *ArcVirtualKeyListItemResponse) HasSecurityAction() bool`

HasSecurityAction returns a boolean if a field has been set.

### GetRpmLimit

`func (o *ArcVirtualKeyListItemResponse) GetRpmLimit() int32`

GetRpmLimit returns the RpmLimit field if non-nil, zero value otherwise.

### GetRpmLimitOk

`func (o *ArcVirtualKeyListItemResponse) GetRpmLimitOk() (*int32, bool)`

GetRpmLimitOk returns a tuple with the RpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmLimit

`func (o *ArcVirtualKeyListItemResponse) SetRpmLimit(v int32)`

SetRpmLimit sets RpmLimit field to given value.

### HasRpmLimit

`func (o *ArcVirtualKeyListItemResponse) HasRpmLimit() bool`

HasRpmLimit returns a boolean if a field has been set.

### SetRpmLimitNil

`func (o *ArcVirtualKeyListItemResponse) SetRpmLimitNil(b bool)`

 SetRpmLimitNil sets the value for RpmLimit to be an explicit nil

### UnsetRpmLimit
`func (o *ArcVirtualKeyListItemResponse) UnsetRpmLimit()`

UnsetRpmLimit ensures that no value is present for RpmLimit, not even an explicit nil
### GetTpmLimit

`func (o *ArcVirtualKeyListItemResponse) GetTpmLimit() int32`

GetTpmLimit returns the TpmLimit field if non-nil, zero value otherwise.

### GetTpmLimitOk

`func (o *ArcVirtualKeyListItemResponse) GetTpmLimitOk() (*int32, bool)`

GetTpmLimitOk returns a tuple with the TpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmLimit

`func (o *ArcVirtualKeyListItemResponse) SetTpmLimit(v int32)`

SetTpmLimit sets TpmLimit field to given value.

### HasTpmLimit

`func (o *ArcVirtualKeyListItemResponse) HasTpmLimit() bool`

HasTpmLimit returns a boolean if a field has been set.

### SetTpmLimitNil

`func (o *ArcVirtualKeyListItemResponse) SetTpmLimitNil(b bool)`

 SetTpmLimitNil sets the value for TpmLimit to be an explicit nil

### UnsetTpmLimit
`func (o *ArcVirtualKeyListItemResponse) UnsetTpmLimit()`

UnsetTpmLimit ensures that no value is present for TpmLimit, not even an explicit nil
### GetId

`func (o *ArcVirtualKeyListItemResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArcVirtualKeyListItemResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArcVirtualKeyListItemResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *ArcVirtualKeyListItemResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ArcVirtualKeyListItemResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ArcVirtualKeyListItemResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetType

`func (o *ArcVirtualKeyListItemResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArcVirtualKeyListItemResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArcVirtualKeyListItemResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUserId

`func (o *ArcVirtualKeyListItemResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ArcVirtualKeyListItemResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ArcVirtualKeyListItemResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserName

`func (o *ArcVirtualKeyListItemResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ArcVirtualKeyListItemResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ArcVirtualKeyListItemResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ArcVirtualKeyListItemResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ArcVirtualKeyListItemResponse) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ArcVirtualKeyListItemResponse) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ArcVirtualKeyListItemResponse) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ArcVirtualKeyListItemResponse) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *ArcVirtualKeyListItemResponse) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *ArcVirtualKeyListItemResponse) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetDeletedAt

`func (o *ArcVirtualKeyListItemResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArcVirtualKeyListItemResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArcVirtualKeyListItemResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArcVirtualKeyListItemResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ArcVirtualKeyListItemResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ArcVirtualKeyListItemResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetCreatedAt

`func (o *ArcVirtualKeyListItemResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArcVirtualKeyListItemResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArcVirtualKeyListItemResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArcVirtualKeyListItemResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArcVirtualKeyListItemResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArcVirtualKeyListItemResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSpendLimitEnabled

`func (o *ArcVirtualKeyListItemResponse) GetSpendLimitEnabled() bool`

GetSpendLimitEnabled returns the SpendLimitEnabled field if non-nil, zero value otherwise.

### GetSpendLimitEnabledOk

`func (o *ArcVirtualKeyListItemResponse) GetSpendLimitEnabledOk() (*bool, bool)`

GetSpendLimitEnabledOk returns a tuple with the SpendLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitEnabled

`func (o *ArcVirtualKeyListItemResponse) SetSpendLimitEnabled(v bool)`

SetSpendLimitEnabled sets SpendLimitEnabled field to given value.

### HasSpendLimitEnabled

`func (o *ArcVirtualKeyListItemResponse) HasSpendLimitEnabled() bool`

HasSpendLimitEnabled returns a boolean if a field has been set.

### GetSpendLimitEnabledOn

`func (o *ArcVirtualKeyListItemResponse) GetSpendLimitEnabledOn() time.Time`

GetSpendLimitEnabledOn returns the SpendLimitEnabledOn field if non-nil, zero value otherwise.

### GetSpendLimitEnabledOnOk

`func (o *ArcVirtualKeyListItemResponse) GetSpendLimitEnabledOnOk() (*time.Time, bool)`

GetSpendLimitEnabledOnOk returns a tuple with the SpendLimitEnabledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitEnabledOn

`func (o *ArcVirtualKeyListItemResponse) SetSpendLimitEnabledOn(v time.Time)`

SetSpendLimitEnabledOn sets SpendLimitEnabledOn field to given value.

### HasSpendLimitEnabledOn

`func (o *ArcVirtualKeyListItemResponse) HasSpendLimitEnabledOn() bool`

HasSpendLimitEnabledOn returns a boolean if a field has been set.

### SetSpendLimitEnabledOnNil

`func (o *ArcVirtualKeyListItemResponse) SetSpendLimitEnabledOnNil(b bool)`

 SetSpendLimitEnabledOnNil sets the value for SpendLimitEnabledOn to be an explicit nil

### UnsetSpendLimitEnabledOn
`func (o *ArcVirtualKeyListItemResponse) UnsetSpendLimitEnabledOn()`

UnsetSpendLimitEnabledOn ensures that no value is present for SpendLimitEnabledOn, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


