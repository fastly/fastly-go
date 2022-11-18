# UserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | The real life name of the user. | [optional] 
**LimitServices** | Pointer to **bool** | Indicates that the user has limited access to the customer&#39;s services. | [optional] 
**Locked** | Pointer to **bool** | Indicates whether the is account is locked for editing or not. | [optional] 
**RequireNewPassword** | Pointer to **bool** | Indicates if a new password is required at next login. | [optional] 
**Role** | Pointer to [**RoleUser**](RoleUser.md) |  | [optional] 
**TwoFactorAuthEnabled** | Pointer to **bool** | Indicates if 2FA is enabled on the user. | [optional] 
**TwoFactorSetupRequired** | Pointer to **bool** | Indicates if 2FA is required by the user&#39;s customer account. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ID** | Pointer to **string** |  | [optional] [readonly] 
**EmailHash** | Pointer to **string** | The alphanumeric string identifying a email login. | [optional] [readonly] 
**CustomerID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewUserResponse

`func NewUserResponse() *UserResponse`

NewUserResponse instantiates a new UserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseWithDefaults

`func NewUserResponseWithDefaults() *UserResponse`

NewUserResponseWithDefaults instantiates a new UserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *UserResponse) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserResponse) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserResponse) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *UserResponse) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetName

`func (o *UserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLimitServices

`func (o *UserResponse) GetLimitServices() bool`

GetLimitServices returns the LimitServices field if non-nil, zero value otherwise.

### GetLimitServicesOk

`func (o *UserResponse) GetLimitServicesOk() (*bool, bool)`

GetLimitServicesOk returns a tuple with the LimitServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitServices

`func (o *UserResponse) SetLimitServices(v bool)`

SetLimitServices sets LimitServices field to given value.

### HasLimitServices

`func (o *UserResponse) HasLimitServices() bool`

HasLimitServices returns a boolean if a field has been set.

### GetLocked

`func (o *UserResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UserResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UserResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UserResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetRequireNewPassword

`func (o *UserResponse) GetRequireNewPassword() bool`

GetRequireNewPassword returns the RequireNewPassword field if non-nil, zero value otherwise.

### GetRequireNewPasswordOk

`func (o *UserResponse) GetRequireNewPasswordOk() (*bool, bool)`

GetRequireNewPasswordOk returns a tuple with the RequireNewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireNewPassword

`func (o *UserResponse) SetRequireNewPassword(v bool)`

SetRequireNewPassword sets RequireNewPassword field to given value.

### HasRequireNewPassword

`func (o *UserResponse) HasRequireNewPassword() bool`

HasRequireNewPassword returns a boolean if a field has been set.

### GetRole

`func (o *UserResponse) GetRole() RoleUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserResponse) GetRoleOk() (*RoleUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserResponse) SetRole(v RoleUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTwoFactorAuthEnabled

`func (o *UserResponse) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *UserResponse) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *UserResponse) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.

### HasTwoFactorAuthEnabled

`func (o *UserResponse) HasTwoFactorAuthEnabled() bool`

HasTwoFactorAuthEnabled returns a boolean if a field has been set.

### GetTwoFactorSetupRequired

`func (o *UserResponse) GetTwoFactorSetupRequired() bool`

GetTwoFactorSetupRequired returns the TwoFactorSetupRequired field if non-nil, zero value otherwise.

### GetTwoFactorSetupRequiredOk

`func (o *UserResponse) GetTwoFactorSetupRequiredOk() (*bool, bool)`

GetTwoFactorSetupRequiredOk returns a tuple with the TwoFactorSetupRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorSetupRequired

`func (o *UserResponse) SetTwoFactorSetupRequired(v bool)`

SetTwoFactorSetupRequired sets TwoFactorSetupRequired field to given value.

### HasTwoFactorSetupRequired

`func (o *UserResponse) HasTwoFactorSetupRequired() bool`

HasTwoFactorSetupRequired returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *UserResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *UserResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *UserResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UserResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UserResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *UserResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *UserResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *UserResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *UserResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *UserResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *UserResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetID

`func (o *UserResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *UserResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *UserResponse) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *UserResponse) HasID() bool`

HasID returns a boolean if a field has been set.

### GetEmailHash

`func (o *UserResponse) GetEmailHash() string`

GetEmailHash returns the EmailHash field if non-nil, zero value otherwise.

### GetEmailHashOk

`func (o *UserResponse) GetEmailHashOk() (*string, bool)`

GetEmailHashOk returns a tuple with the EmailHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHash

`func (o *UserResponse) SetEmailHash(v string)`

SetEmailHash sets EmailHash field to given value.

### HasEmailHash

`func (o *UserResponse) HasEmailHash() bool`

HasEmailHash returns a boolean if a field has been set.

### GetCustomerID

`func (o *UserResponse) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *UserResponse) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *UserResponse) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *UserResponse) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
