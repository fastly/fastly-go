# BotManagementBotsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**BotManagementMeta**](BotManagementMeta.md) |  | 
**Data** | [**[]BotManagementBotResponse**](BotManagementBotResponse.md) | The bots returned by the request. | 

## Methods

### NewBotManagementBotsResponse

`func NewBotManagementBotsResponse(meta BotManagementMeta, data []BotManagementBotResponse, ) *BotManagementBotsResponse`

NewBotManagementBotsResponse instantiates a new BotManagementBotsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementBotsResponseWithDefaults

`func NewBotManagementBotsResponseWithDefaults() *BotManagementBotsResponse`

NewBotManagementBotsResponseWithDefaults instantiates a new BotManagementBotsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *BotManagementBotsResponse) GetMeta() BotManagementMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BotManagementBotsResponse) GetMetaOk() (*BotManagementMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BotManagementBotsResponse) SetMeta(v BotManagementMeta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *BotManagementBotsResponse) GetData() []BotManagementBotResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BotManagementBotsResponse) GetDataOk() (*[]BotManagementBotResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BotManagementBotsResponse) SetData(v []BotManagementBotResponse)`

SetData sets Data field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


