# IamPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Alphanumeric string identifying the permission. | [optional] 
**Object** | Pointer to **string** | The type of the object. | [optional] 
**Name** | Pointer to **string** | Name of the permission. | [optional] 
**Description** | Pointer to **string** | The description of the permission. | [optional] 
**ResourceName** | Pointer to **string** | The name of the resource the operation will be performed on. | [optional] 
**ResourceDescription** | Pointer to **string** | The description of the resource. | [optional] 
**Scope** | Pointer to **string** | Permissions are either \&quot;service\&quot; level or \&quot;account\&quot; level. | [optional] 

## Methods

### NewIamPermission

`func NewIamPermission() *IamPermission`

NewIamPermission instantiates a new IamPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPermissionWithDefaults

`func NewIamPermissionWithDefaults() *IamPermission`

NewIamPermissionWithDefaults instantiates a new IamPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IamPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamPermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamPermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *IamPermission) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IamPermission) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IamPermission) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *IamPermission) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetName

`func (o *IamPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamPermission) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamPermission) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IamPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceName

`func (o *IamPermission) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *IamPermission) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *IamPermission) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *IamPermission) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceDescription

`func (o *IamPermission) GetResourceDescription() string`

GetResourceDescription returns the ResourceDescription field if non-nil, zero value otherwise.

### GetResourceDescriptionOk

`func (o *IamPermission) GetResourceDescriptionOk() (*string, bool)`

GetResourceDescriptionOk returns a tuple with the ResourceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDescription

`func (o *IamPermission) SetResourceDescription(v string)`

SetResourceDescription sets ResourceDescription field to given value.

### HasResourceDescription

`func (o *IamPermission) HasResourceDescription() bool`

HasResourceDescription returns a boolean if a field has been set.

### GetScope

`func (o *IamPermission) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IamPermission) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IamPermission) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IamPermission) HasScope() bool`

HasScope returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


