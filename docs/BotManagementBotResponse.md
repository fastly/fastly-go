# BotManagementBotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**BotManagementBotAction**](BotManagementBotAction.md) |  | 
**BotId** | **string** | The unique identifier of the bot. | [readonly] 
**Name** | **string** | The display name of the bot. | [readonly] 

## Methods

### NewBotManagementBotResponse

`func NewBotManagementBotResponse(action BotManagementBotAction, botId string, name string, ) *BotManagementBotResponse`

NewBotManagementBotResponse instantiates a new BotManagementBotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementBotResponseWithDefaults

`func NewBotManagementBotResponseWithDefaults() *BotManagementBotResponse`

NewBotManagementBotResponseWithDefaults instantiates a new BotManagementBotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BotManagementBotResponse) GetAction() BotManagementBotAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BotManagementBotResponse) GetActionOk() (*BotManagementBotAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BotManagementBotResponse) SetAction(v BotManagementBotAction)`

SetAction sets Action field to given value.


### GetBotId

`func (o *BotManagementBotResponse) GetBotId() string`

GetBotId returns the BotId field if non-nil, zero value otherwise.

### GetBotIdOk

`func (o *BotManagementBotResponse) GetBotIdOk() (*string, bool)`

GetBotIdOk returns a tuple with the BotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotId

`func (o *BotManagementBotResponse) SetBotId(v string)`

SetBotId sets BotId field to given value.


### GetName

`func (o *BotManagementBotResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BotManagementBotResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BotManagementBotResponse) SetName(v string)`

SetName sets Name field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


