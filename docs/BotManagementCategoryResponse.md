# BotManagementCategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**BotManagementAction**](BotManagementAction.md) |  | 
**CategoryId** | **string** | The unique identifier of the category. | [readonly] 
**Name** | **string** | The display name of the category. | [readonly] 
**Description** | Pointer to **string** | A description of the bots the category contains. | [optional] [readonly] 

## Methods

### NewBotManagementCategoryResponse

`func NewBotManagementCategoryResponse(action BotManagementAction, categoryId string, name string, ) *BotManagementCategoryResponse`

NewBotManagementCategoryResponse instantiates a new BotManagementCategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementCategoryResponseWithDefaults

`func NewBotManagementCategoryResponseWithDefaults() *BotManagementCategoryResponse`

NewBotManagementCategoryResponseWithDefaults instantiates a new BotManagementCategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BotManagementCategoryResponse) GetAction() BotManagementAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BotManagementCategoryResponse) GetActionOk() (*BotManagementAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BotManagementCategoryResponse) SetAction(v BotManagementAction)`

SetAction sets Action field to given value.


### GetCategoryId

`func (o *BotManagementCategoryResponse) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *BotManagementCategoryResponse) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *BotManagementCategoryResponse) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetName

`func (o *BotManagementCategoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BotManagementCategoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BotManagementCategoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BotManagementCategoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BotManagementCategoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BotManagementCategoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BotManagementCategoryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


