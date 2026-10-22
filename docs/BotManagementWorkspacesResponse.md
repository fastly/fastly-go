# BotManagementWorkspacesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**BotManagementMeta**](BotManagementMeta.md) |  | 
**Data** | [**[]BotManagementWorkspaceResponse**](BotManagementWorkspaceResponse.md) | The workspaces returned by the request. | 

## Methods

### NewBotManagementWorkspacesResponse

`func NewBotManagementWorkspacesResponse(meta BotManagementMeta, data []BotManagementWorkspaceResponse, ) *BotManagementWorkspacesResponse`

NewBotManagementWorkspacesResponse instantiates a new BotManagementWorkspacesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementWorkspacesResponseWithDefaults

`func NewBotManagementWorkspacesResponseWithDefaults() *BotManagementWorkspacesResponse`

NewBotManagementWorkspacesResponseWithDefaults instantiates a new BotManagementWorkspacesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *BotManagementWorkspacesResponse) GetMeta() BotManagementMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BotManagementWorkspacesResponse) GetMetaOk() (*BotManagementMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BotManagementWorkspacesResponse) SetMeta(v BotManagementMeta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *BotManagementWorkspacesResponse) GetData() []BotManagementWorkspaceResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BotManagementWorkspacesResponse) GetDataOk() (*[]BotManagementWorkspaceResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BotManagementWorkspacesResponse) SetData(v []BotManagementWorkspaceResponse)`

SetData sets Data field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


