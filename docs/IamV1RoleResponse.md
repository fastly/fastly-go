# IamV1RoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]string** | The set of permissions granted to this role. | [optional] 

## Methods

### NewIamV1RoleResponse

`func NewIamV1RoleResponse() *IamV1RoleResponse`

NewIamV1RoleResponse instantiates a new IamV1RoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV1RoleResponseWithDefaults

`func NewIamV1RoleResponseWithDefaults() *IamV1RoleResponse`

NewIamV1RoleResponseWithDefaults instantiates a new IamV1RoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IamV1RoleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV1RoleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV1RoleResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV1RoleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IamV1RoleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamV1RoleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamV1RoleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamV1RoleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV1RoleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV1RoleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV1RoleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV1RoleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *IamV1RoleResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamV1RoleResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamV1RoleResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamV1RoleResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


