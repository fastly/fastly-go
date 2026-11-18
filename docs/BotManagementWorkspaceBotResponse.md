# BotManagementWorkspaceBotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**BotManagementBotAction**](BotManagementBotAction.md) |  | 
**BotID** | **string** | The unique identifier of the bot. | [readonly] 
**Name** | **string** | The display name of the bot. | [readonly] 
**CategoryID** | **string** | The ID of the category the bot belongs to. | [readonly] 

## Methods

### NewBotManagementWorkspaceBotResponse

`func NewBotManagementWorkspaceBotResponse(action BotManagementBotAction, botID string, name string, categoryID string, ) *BotManagementWorkspaceBotResponse`

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


### GetBotID

`func (o *BotManagementWorkspaceBotResponse) GetBotID() string`

GetBotID returns the BotID field if non-nil, zero value otherwise.

### GetBotIDOk

`func (o *BotManagementWorkspaceBotResponse) GetBotIDOk() (*string, bool)`

GetBotIDOk returns a tuple with the BotID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotID

`func (o *BotManagementWorkspaceBotResponse) SetBotID(v string)`

SetBotID sets BotID field to given value.


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


### GetCategoryID

`func (o *BotManagementWorkspaceBotResponse) GetCategoryID() string`

GetCategoryID returns the CategoryID field if non-nil, zero value otherwise.

### GetCategoryIDOk

`func (o *BotManagementWorkspaceBotResponse) GetCategoryIDOk() (*string, bool)`

GetCategoryIDOk returns a tuple with the CategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryID

`func (o *BotManagementWorkspaceBotResponse) SetCategoryID(v string)`

SetCategoryID sets CategoryID field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
