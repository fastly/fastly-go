# BotManagementWorkspaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The display name of the workspace. | 
**Description** | Pointer to **string** | A description of the workspace. | [optional] 
**ProtectionMode** | [**BotManagementProtectionMode**](BotManagementProtectionMode.md) |  | 
**ID** | **string** | A UUID used to uniquely identify the workspace. | [readonly] 
**Services** | Pointer to **[]string** | The IDs of the services attached to this workspace. | [optional] [readonly] 
**CreatedAt** | **time.Time** | The date and time the workspace was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the workspace was last updated, in RFC 3339 format. | [readonly] 

## Methods

### NewBotManagementWorkspaceResponse

`func NewBotManagementWorkspaceResponse(name string, protectionMode BotManagementProtectionMode, id string, createdAt time.Time, updatedAt time.Time, ) *BotManagementWorkspaceResponse`

NewBotManagementWorkspaceResponse instantiates a new BotManagementWorkspaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementWorkspaceResponseWithDefaults

`func NewBotManagementWorkspaceResponseWithDefaults() *BotManagementWorkspaceResponse`

NewBotManagementWorkspaceResponseWithDefaults instantiates a new BotManagementWorkspaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BotManagementWorkspaceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BotManagementWorkspaceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BotManagementWorkspaceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BotManagementWorkspaceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BotManagementWorkspaceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BotManagementWorkspaceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BotManagementWorkspaceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProtectionMode

`func (o *BotManagementWorkspaceResponse) GetProtectionMode() BotManagementProtectionMode`

GetProtectionMode returns the ProtectionMode field if non-nil, zero value otherwise.

### GetProtectionModeOk

`func (o *BotManagementWorkspaceResponse) GetProtectionModeOk() (*BotManagementProtectionMode, bool)`

GetProtectionModeOk returns a tuple with the ProtectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionMode

`func (o *BotManagementWorkspaceResponse) SetProtectionMode(v BotManagementProtectionMode)`

SetProtectionMode sets ProtectionMode field to given value.


### GetID

`func (o *BotManagementWorkspaceResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *BotManagementWorkspaceResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *BotManagementWorkspaceResponse) SetID(v string)`

SetID sets ID field to given value.


### GetServices

`func (o *BotManagementWorkspaceResponse) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BotManagementWorkspaceResponse) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BotManagementWorkspaceResponse) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *BotManagementWorkspaceResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BotManagementWorkspaceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BotManagementWorkspaceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BotManagementWorkspaceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BotManagementWorkspaceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BotManagementWorkspaceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BotManagementWorkspaceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
