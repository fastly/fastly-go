# BotManagementWorkspaceResponseExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | A UUID used to uniquely identify the workspace. | [readonly] 
**Services** | Pointer to **[]string** | The IDs of the services attached to this workspace. | [optional] [readonly] 
**CreatedAt** | **time.Time** | The date and time the workspace was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the workspace was last updated, in RFC 3339 format. | [readonly] 

## Methods

### NewBotManagementWorkspaceResponseExtra

`func NewBotManagementWorkspaceResponseExtra(id string, createdAt time.Time, updatedAt time.Time, ) *BotManagementWorkspaceResponseExtra`

NewBotManagementWorkspaceResponseExtra instantiates a new BotManagementWorkspaceResponseExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementWorkspaceResponseExtraWithDefaults

`func NewBotManagementWorkspaceResponseExtraWithDefaults() *BotManagementWorkspaceResponseExtra`

NewBotManagementWorkspaceResponseExtraWithDefaults instantiates a new BotManagementWorkspaceResponseExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *BotManagementWorkspaceResponseExtra) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *BotManagementWorkspaceResponseExtra) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *BotManagementWorkspaceResponseExtra) SetID(v string)`

SetID sets ID field to given value.


### GetServices

`func (o *BotManagementWorkspaceResponseExtra) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BotManagementWorkspaceResponseExtra) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BotManagementWorkspaceResponseExtra) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *BotManagementWorkspaceResponseExtra) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BotManagementWorkspaceResponseExtra) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BotManagementWorkspaceResponseExtra) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BotManagementWorkspaceResponseExtra) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BotManagementWorkspaceResponseExtra) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BotManagementWorkspaceResponseExtra) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BotManagementWorkspaceResponseExtra) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
