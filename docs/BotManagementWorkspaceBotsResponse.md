# BotManagementWorkspaceBotsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**BotManagementMeta**](BotManagementMeta.md) |  | 
**Data** | [**[]BotManagementWorkspaceBotResponse**](BotManagementWorkspaceBotResponse.md) | The bots returned by the request. | 

## Methods

### NewBotManagementWorkspaceBotsResponse

`func NewBotManagementWorkspaceBotsResponse(meta BotManagementMeta, data []BotManagementWorkspaceBotResponse, ) *BotManagementWorkspaceBotsResponse`

NewBotManagementWorkspaceBotsResponse instantiates a new BotManagementWorkspaceBotsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementWorkspaceBotsResponseWithDefaults

`func NewBotManagementWorkspaceBotsResponseWithDefaults() *BotManagementWorkspaceBotsResponse`

NewBotManagementWorkspaceBotsResponseWithDefaults instantiates a new BotManagementWorkspaceBotsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *BotManagementWorkspaceBotsResponse) GetMeta() BotManagementMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BotManagementWorkspaceBotsResponse) GetMetaOk() (*BotManagementMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BotManagementWorkspaceBotsResponse) SetMeta(v BotManagementMeta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *BotManagementWorkspaceBotsResponse) GetData() []BotManagementWorkspaceBotResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BotManagementWorkspaceBotsResponse) GetDataOk() (*[]BotManagementWorkspaceBotResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BotManagementWorkspaceBotsResponse) SetData(v []BotManagementWorkspaceBotResponse)`

SetData sets Data field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
