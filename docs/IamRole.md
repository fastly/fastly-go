# IamRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ID** | Pointer to **string** | Alphanumeric string identifying the role. | [optional] 
**Object** | Pointer to **string** | The type of the object. | [optional] 
**Name** | Pointer to **string** | Name of the role. | [optional] 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**Custom** | Pointer to **bool** | This attribute is set to `true` if the role is managed by the customer. It is set to `false` if the role was created by Fastly. | [optional] 
**PermissionsCount** | Pointer to **int32** | Number of permissions assigned to the role. | [optional] 

## Methods

### NewIamRole

`func NewIamRole() *IamRole`

NewIamRole instantiates a new IamRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRoleWithDefaults

`func NewIamRoleWithDefaults() *IamRole`

NewIamRoleWithDefaults instantiates a new IamRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IamRole) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamRole) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamRole) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IamRole) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IamRole) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *IamRole) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamRole) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamRole) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamRole) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IamRole) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IamRole) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetID

`func (o *IamRole) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *IamRole) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *IamRole) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *IamRole) HasID() bool`

HasID returns a boolean if a field has been set.

### GetObject

`func (o *IamRole) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IamRole) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IamRole) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *IamRole) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetName

`func (o *IamRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IamRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustom

`func (o *IamRole) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *IamRole) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *IamRole) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *IamRole) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetPermissionsCount

`func (o *IamRole) GetPermissionsCount() int32`

GetPermissionsCount returns the PermissionsCount field if non-nil, zero value otherwise.

### GetPermissionsCountOk

`func (o *IamRole) GetPermissionsCountOk() (*int32, bool)`

GetPermissionsCountOk returns a tuple with the PermissionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsCount

`func (o *IamRole) SetPermissionsCount(v int32)`

SetPermissionsCount sets PermissionsCount field to given value.

### HasPermissionsCount

`func (o *IamRole) HasPermissionsCount() bool`

HasPermissionsCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
