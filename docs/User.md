# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | The real life name of the user. | [optional] 
**LimitServices** | Pointer to **bool** | Indicates that the user has limited access to the customer&#39;s services. | [optional] 
**Locked** | Pointer to **NullableBool** | Indicates whether the is account is locked for editing or not. | [optional] 
**RequireNewPassword** | Pointer to **NullableBool** | Indicates if a new password is required at next login. | [optional] 
**Role** | Pointer to [**RoleUser**](RoleUser.md) |  | [optional] 
**Roles** | Pointer to **[]string** | A list of role IDs assigned to the user. | [optional] 
**TwoFactorAuthEnabled** | Pointer to **NullableBool** | Indicates if 2FA is enabled on the user. | [optional] 
**TwoFactorSetupRequired** | Pointer to **bool** | Indicates if 2FA is required by the user&#39;s customer account. | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *User) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *User) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *User) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *User) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLimitServices

`func (o *User) GetLimitServices() bool`

GetLimitServices returns the LimitServices field if non-nil, zero value otherwise.

### GetLimitServicesOk

`func (o *User) GetLimitServicesOk() (*bool, bool)`

GetLimitServicesOk returns a tuple with the LimitServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitServices

`func (o *User) SetLimitServices(v bool)`

SetLimitServices sets LimitServices field to given value.

### HasLimitServices

`func (o *User) HasLimitServices() bool`

HasLimitServices returns a boolean if a field has been set.

### GetLocked

`func (o *User) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *User) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *User) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *User) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *User) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *User) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetRequireNewPassword

`func (o *User) GetRequireNewPassword() bool`

GetRequireNewPassword returns the RequireNewPassword field if non-nil, zero value otherwise.

### GetRequireNewPasswordOk

`func (o *User) GetRequireNewPasswordOk() (*bool, bool)`

GetRequireNewPasswordOk returns a tuple with the RequireNewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireNewPassword

`func (o *User) SetRequireNewPassword(v bool)`

SetRequireNewPassword sets RequireNewPassword field to given value.

### HasRequireNewPassword

`func (o *User) HasRequireNewPassword() bool`

HasRequireNewPassword returns a boolean if a field has been set.

### SetRequireNewPasswordNil

`func (o *User) SetRequireNewPasswordNil(b bool)`

 SetRequireNewPasswordNil sets the value for RequireNewPassword to be an explicit nil

### UnsetRequireNewPassword
`func (o *User) UnsetRequireNewPassword()`

UnsetRequireNewPassword ensures that no value is present for RequireNewPassword, not even an explicit nil
### GetRole

`func (o *User) GetRole() RoleUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *User) GetRoleOk() (*RoleUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *User) SetRole(v RoleUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *User) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoles

`func (o *User) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *User) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *User) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTwoFactorAuthEnabled

`func (o *User) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *User) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *User) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.

### HasTwoFactorAuthEnabled

`func (o *User) HasTwoFactorAuthEnabled() bool`

HasTwoFactorAuthEnabled returns a boolean if a field has been set.

### SetTwoFactorAuthEnabledNil

`func (o *User) SetTwoFactorAuthEnabledNil(b bool)`

 SetTwoFactorAuthEnabledNil sets the value for TwoFactorAuthEnabled to be an explicit nil

### UnsetTwoFactorAuthEnabled
`func (o *User) UnsetTwoFactorAuthEnabled()`

UnsetTwoFactorAuthEnabled ensures that no value is present for TwoFactorAuthEnabled, not even an explicit nil
### GetTwoFactorSetupRequired

`func (o *User) GetTwoFactorSetupRequired() bool`

GetTwoFactorSetupRequired returns the TwoFactorSetupRequired field if non-nil, zero value otherwise.

### GetTwoFactorSetupRequiredOk

`func (o *User) GetTwoFactorSetupRequiredOk() (*bool, bool)`

GetTwoFactorSetupRequiredOk returns a tuple with the TwoFactorSetupRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorSetupRequired

`func (o *User) SetTwoFactorSetupRequired(v bool)`

SetTwoFactorSetupRequired sets TwoFactorSetupRequired field to given value.

### HasTwoFactorSetupRequired

`func (o *User) HasTwoFactorSetupRequired() bool`

HasTwoFactorSetupRequired returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


