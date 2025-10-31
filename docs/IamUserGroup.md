# IamUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Id** | Pointer to **string** | Alphanumeric string identifying the user group. | [optional] 
**Name** | Pointer to **string** | Name of the user group. | [optional] 
**Description** | Pointer to **string** | Description of the user group. | [optional] 
**InvitationsCount** | Pointer to **int32** | Number of invitations added to the user group. | [optional] 
**UsersCount** | Pointer to **int32** | Number of users added to the user group. | [optional] 
**RolesCount** | Pointer to **int32** | Number of roles added to the user group. | [optional] 

## Methods

### NewIamUserGroup

`func NewIamUserGroup() *IamUserGroup`

NewIamUserGroup instantiates a new IamUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserGroupWithDefaults

`func NewIamUserGroupWithDefaults() *IamUserGroup`

NewIamUserGroupWithDefaults instantiates a new IamUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IamUserGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamUserGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamUserGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamUserGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IamUserGroup) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IamUserGroup) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *IamUserGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamUserGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamUserGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamUserGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IamUserGroup) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IamUserGroup) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *IamUserGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamUserGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamUserGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamUserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IamUserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamUserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamUserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamUserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IamUserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamUserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamUserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamUserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInvitationsCount

`func (o *IamUserGroup) GetInvitationsCount() int32`

GetInvitationsCount returns the InvitationsCount field if non-nil, zero value otherwise.

### GetInvitationsCountOk

`func (o *IamUserGroup) GetInvitationsCountOk() (*int32, bool)`

GetInvitationsCountOk returns a tuple with the InvitationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationsCount

`func (o *IamUserGroup) SetInvitationsCount(v int32)`

SetInvitationsCount sets InvitationsCount field to given value.

### HasInvitationsCount

`func (o *IamUserGroup) HasInvitationsCount() bool`

HasInvitationsCount returns a boolean if a field has been set.

### GetUsersCount

`func (o *IamUserGroup) GetUsersCount() int32`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *IamUserGroup) GetUsersCountOk() (*int32, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *IamUserGroup) SetUsersCount(v int32)`

SetUsersCount sets UsersCount field to given value.

### HasUsersCount

`func (o *IamUserGroup) HasUsersCount() bool`

HasUsersCount returns a boolean if a field has been set.

### GetRolesCount

`func (o *IamUserGroup) GetRolesCount() int32`

GetRolesCount returns the RolesCount field if non-nil, zero value otherwise.

### GetRolesCountOk

`func (o *IamUserGroup) GetRolesCountOk() (*int32, bool)`

GetRolesCountOk returns a tuple with the RolesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesCount

`func (o *IamUserGroup) SetRolesCount(v int32)`

SetRolesCount sets RolesCount field to given value.

### HasRolesCount

`func (o *IamUserGroup) HasRolesCount() bool`

HasRolesCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


