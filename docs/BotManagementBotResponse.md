# BotManagementBotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**BotManagementBotAction**](BotManagementBotAction.md) |  | 
**BotID** | **string** | The unique identifier of the bot. | [readonly] 
**Name** | **string** | The display name of the bot. | [readonly] 

## Methods

### NewBotManagementBotResponse

`func NewBotManagementBotResponse(action BotManagementBotAction, botID string, name string, ) *BotManagementBotResponse`

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


### GetBotID

`func (o *BotManagementBotResponse) GetBotID() string`

GetBotID returns the BotID field if non-nil, zero value otherwise.

### GetBotIDOk

`func (o *BotManagementBotResponse) GetBotIDOk() (*string, bool)`

GetBotIDOk returns a tuple with the BotID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotID

`func (o *BotManagementBotResponse) SetBotID(v string)`

SetBotID sets BotID field to given value.


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
