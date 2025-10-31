# SchemasUserResponse

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
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**EmailHash** | Pointer to **string** | The alphanumeric string identifying a email login. | [optional] [readonly] 
**CustomerId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSchemasUserResponse

`func NewSchemasUserResponse() *SchemasUserResponse`

NewSchemasUserResponse instantiates a new SchemasUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemasUserResponseWithDefaults

`func NewSchemasUserResponseWithDefaults() *SchemasUserResponse`

NewSchemasUserResponseWithDefaults instantiates a new SchemasUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *SchemasUserResponse) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *SchemasUserResponse) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *SchemasUserResponse) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *SchemasUserResponse) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetName

`func (o *SchemasUserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemasUserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemasUserResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchemasUserResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLimitServices

`func (o *SchemasUserResponse) GetLimitServices() bool`

GetLimitServices returns the LimitServices field if non-nil, zero value otherwise.

### GetLimitServicesOk

`func (o *SchemasUserResponse) GetLimitServicesOk() (*bool, bool)`

GetLimitServicesOk returns a tuple with the LimitServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitServices

`func (o *SchemasUserResponse) SetLimitServices(v bool)`

SetLimitServices sets LimitServices field to given value.

### HasLimitServices

`func (o *SchemasUserResponse) HasLimitServices() bool`

HasLimitServices returns a boolean if a field has been set.

### GetLocked

`func (o *SchemasUserResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SchemasUserResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SchemasUserResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SchemasUserResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *SchemasUserResponse) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *SchemasUserResponse) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetRequireNewPassword

`func (o *SchemasUserResponse) GetRequireNewPassword() bool`

GetRequireNewPassword returns the RequireNewPassword field if non-nil, zero value otherwise.

### GetRequireNewPasswordOk

`func (o *SchemasUserResponse) GetRequireNewPasswordOk() (*bool, bool)`

GetRequireNewPasswordOk returns a tuple with the RequireNewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireNewPassword

`func (o *SchemasUserResponse) SetRequireNewPassword(v bool)`

SetRequireNewPassword sets RequireNewPassword field to given value.

### HasRequireNewPassword

`func (o *SchemasUserResponse) HasRequireNewPassword() bool`

HasRequireNewPassword returns a boolean if a field has been set.

### SetRequireNewPasswordNil

`func (o *SchemasUserResponse) SetRequireNewPasswordNil(b bool)`

 SetRequireNewPasswordNil sets the value for RequireNewPassword to be an explicit nil

### UnsetRequireNewPassword
`func (o *SchemasUserResponse) UnsetRequireNewPassword()`

UnsetRequireNewPassword ensures that no value is present for RequireNewPassword, not even an explicit nil
### GetRole

`func (o *SchemasUserResponse) GetRole() RoleUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SchemasUserResponse) GetRoleOk() (*RoleUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SchemasUserResponse) SetRole(v RoleUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *SchemasUserResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoles

`func (o *SchemasUserResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SchemasUserResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SchemasUserResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *SchemasUserResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTwoFactorAuthEnabled

`func (o *SchemasUserResponse) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *SchemasUserResponse) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *SchemasUserResponse) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.

### HasTwoFactorAuthEnabled

`func (o *SchemasUserResponse) HasTwoFactorAuthEnabled() bool`

HasTwoFactorAuthEnabled returns a boolean if a field has been set.

### SetTwoFactorAuthEnabledNil

`func (o *SchemasUserResponse) SetTwoFactorAuthEnabledNil(b bool)`

 SetTwoFactorAuthEnabledNil sets the value for TwoFactorAuthEnabled to be an explicit nil

### UnsetTwoFactorAuthEnabled
`func (o *SchemasUserResponse) UnsetTwoFactorAuthEnabled()`

UnsetTwoFactorAuthEnabled ensures that no value is present for TwoFactorAuthEnabled, not even an explicit nil
### GetTwoFactorSetupRequired

`func (o *SchemasUserResponse) GetTwoFactorSetupRequired() bool`

GetTwoFactorSetupRequired returns the TwoFactorSetupRequired field if non-nil, zero value otherwise.

### GetTwoFactorSetupRequiredOk

`func (o *SchemasUserResponse) GetTwoFactorSetupRequiredOk() (*bool, bool)`

GetTwoFactorSetupRequiredOk returns a tuple with the TwoFactorSetupRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorSetupRequired

`func (o *SchemasUserResponse) SetTwoFactorSetupRequired(v bool)`

SetTwoFactorSetupRequired sets TwoFactorSetupRequired field to given value.

### HasTwoFactorSetupRequired

`func (o *SchemasUserResponse) HasTwoFactorSetupRequired() bool`

HasTwoFactorSetupRequired returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SchemasUserResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SchemasUserResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SchemasUserResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SchemasUserResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *SchemasUserResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SchemasUserResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *SchemasUserResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SchemasUserResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SchemasUserResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SchemasUserResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *SchemasUserResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *SchemasUserResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *SchemasUserResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SchemasUserResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SchemasUserResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SchemasUserResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SchemasUserResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SchemasUserResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *SchemasUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemasUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemasUserResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemasUserResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmailHash

`func (o *SchemasUserResponse) GetEmailHash() string`

GetEmailHash returns the EmailHash field if non-nil, zero value otherwise.

### GetEmailHashOk

`func (o *SchemasUserResponse) GetEmailHashOk() (*string, bool)`

GetEmailHashOk returns a tuple with the EmailHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHash

`func (o *SchemasUserResponse) SetEmailHash(v string)`

SetEmailHash sets EmailHash field to given value.

### HasEmailHash

`func (o *SchemasUserResponse) HasEmailHash() bool`

HasEmailHash returns a boolean if a field has been set.

### GetCustomerId

`func (o *SchemasUserResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SchemasUserResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SchemasUserResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SchemasUserResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


