# BotManagementWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The display name of the workspace. | [optional] 
**Description** | Pointer to **string** | A description of the workspace. | [optional] 
**ProtectionMode** | Pointer to [**BotManagementProtectionMode**](BotManagementProtectionMode.md) |  | [optional] 

## Methods

### NewBotManagementWorkspace

`func NewBotManagementWorkspace() *BotManagementWorkspace`

NewBotManagementWorkspace instantiates a new BotManagementWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementWorkspaceWithDefaults

`func NewBotManagementWorkspaceWithDefaults() *BotManagementWorkspace`

NewBotManagementWorkspaceWithDefaults instantiates a new BotManagementWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BotManagementWorkspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BotManagementWorkspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BotManagementWorkspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BotManagementWorkspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BotManagementWorkspace) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BotManagementWorkspace) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BotManagementWorkspace) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BotManagementWorkspace) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProtectionMode

`func (o *BotManagementWorkspace) GetProtectionMode() BotManagementProtectionMode`

GetProtectionMode returns the ProtectionMode field if non-nil, zero value otherwise.

### GetProtectionModeOk

`func (o *BotManagementWorkspace) GetProtectionModeOk() (*BotManagementProtectionMode, bool)`

GetProtectionModeOk returns a tuple with the ProtectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionMode

`func (o *BotManagementWorkspace) SetProtectionMode(v BotManagementProtectionMode)`

SetProtectionMode sets ProtectionMode field to given value.

### HasProtectionMode

`func (o *BotManagementWorkspace) HasProtectionMode() bool`

HasProtectionMode returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
