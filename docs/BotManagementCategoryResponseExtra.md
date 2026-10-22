# BotManagementCategoryResponseExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **string** | The unique identifier of the category. | [readonly] 
**Name** | **string** | The display name of the category. | [readonly] 
**Description** | Pointer to **string** | A description of the bots the category contains. | [optional] [readonly] 

## Methods

### NewBotManagementCategoryResponseExtra

`func NewBotManagementCategoryResponseExtra(categoryId string, name string, ) *BotManagementCategoryResponseExtra`

NewBotManagementCategoryResponseExtra instantiates a new BotManagementCategoryResponseExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementCategoryResponseExtraWithDefaults

`func NewBotManagementCategoryResponseExtraWithDefaults() *BotManagementCategoryResponseExtra`

NewBotManagementCategoryResponseExtraWithDefaults instantiates a new BotManagementCategoryResponseExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *BotManagementCategoryResponseExtra) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *BotManagementCategoryResponseExtra) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *BotManagementCategoryResponseExtra) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetName

`func (o *BotManagementCategoryResponseExtra) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BotManagementCategoryResponseExtra) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BotManagementCategoryResponseExtra) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BotManagementCategoryResponseExtra) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BotManagementCategoryResponseExtra) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BotManagementCategoryResponseExtra) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BotManagementCategoryResponseExtra) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


