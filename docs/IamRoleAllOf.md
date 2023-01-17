# IamRoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Alphanumeric string identifying the role. | [optional] 
**Object** | Pointer to **string** | The type of the object. | [optional] 
**Name** | Pointer to **string** | Name of the role. | [optional] 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**Custom** | Pointer to **bool** | This attribute is set to `true` if the role is managed by the customer. It is set to `false` if the role was created by Fastly. | [optional] 
**PermissionsCount** | Pointer to **int32** | Number of permissions assigned to the role. | [optional] 

## Methods

### NewIamRoleAllOf

`func NewIamRoleAllOf() *IamRoleAllOf`

NewIamRoleAllOf instantiates a new IamRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRoleAllOfWithDefaults

`func NewIamRoleAllOfWithDefaults() *IamRoleAllOf`

NewIamRoleAllOfWithDefaults instantiates a new IamRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *IamRoleAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *IamRoleAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *IamRoleAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *IamRoleAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetObject

`func (o *IamRoleAllOf) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IamRoleAllOf) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IamRoleAllOf) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *IamRoleAllOf) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetName

`func (o *IamRoleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamRoleAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamRoleAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamRoleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IamRoleAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamRoleAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamRoleAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamRoleAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustom

`func (o *IamRoleAllOf) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *IamRoleAllOf) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *IamRoleAllOf) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *IamRoleAllOf) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetPermissionsCount

`func (o *IamRoleAllOf) GetPermissionsCount() int32`

GetPermissionsCount returns the PermissionsCount field if non-nil, zero value otherwise.

### GetPermissionsCountOk

`func (o *IamRoleAllOf) GetPermissionsCountOk() (*int32, bool)`

GetPermissionsCountOk returns a tuple with the PermissionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsCount

`func (o *IamRoleAllOf) SetPermissionsCount(v int32)`

SetPermissionsCount sets PermissionsCount field to given value.

### HasPermissionsCount

`func (o *IamRoleAllOf) HasPermissionsCount() bool`

HasPermissionsCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
