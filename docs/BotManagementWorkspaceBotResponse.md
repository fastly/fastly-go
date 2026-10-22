# BotManagementWorkspaceBotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**BotManagementBotAction**](BotManagementBotAction.md) |  | 
**BotId** | **string** | The unique identifier of the bot. | [readonly] 
**Name** | **string** | The display name of the bot. | [readonly] 
**CategoryId** | **string** | The ID of the category the bot belongs to. | [readonly] 

## Methods

### NewBotManagementWorkspaceBotResponse

`func NewBotManagementWorkspaceBotResponse(action BotManagementBotAction, botId string, name string, categoryId string, ) *BotManagementWorkspaceBotResponse`

NewBotManagementWorkspaceBotResponse instantiates a new BotManagementWorkspaceBotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementWorkspaceBotResponseWithDefaults

`func NewBotManagementWorkspaceBotResponseWithDefaults() *BotManagementWorkspaceBotResponse`

NewBotManagementWorkspaceBotResponseWithDefaults instantiates a new BotManagementWorkspaceBotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BotManagementWorkspaceBotResponse) GetAction() BotManagementBotAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BotManagementWorkspaceBotResponse) GetActionOk() (*BotManagementBotAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BotManagementWorkspaceBotResponse) SetAction(v BotManagementBotAction)`

SetAction sets Action field to given value.


### GetBotId

`func (o *BotManagementWorkspaceBotResponse) GetBotId() string`

GetBotId returns the BotId field if non-nil, zero value otherwise.

### GetBotIdOk

`func (o *BotManagementWorkspaceBotResponse) GetBotIdOk() (*string, bool)`

GetBotIdOk returns a tuple with the BotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotId

`func (o *BotManagementWorkspaceBotResponse) SetBotId(v string)`

SetBotId sets BotId field to given value.


### GetName

`func (o *BotManagementWorkspaceBotResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BotManagementWorkspaceBotResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BotManagementWorkspaceBotResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCategoryId

`func (o *BotManagementWorkspaceBotResponse) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *BotManagementWorkspaceBotResponse) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *BotManagementWorkspaceBotResponse) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


