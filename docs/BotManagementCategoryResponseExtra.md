# BotManagementCategoryResponseExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryID** | **string** | The unique identifier of the category. | [readonly] 
**Name** | **string** | The display name of the category. | [readonly] 
**Description** | Pointer to **string** | A description of the bots the category contains. | [optional] [readonly] 

## Methods

### NewBotManagementCategoryResponseExtra

`func NewBotManagementCategoryResponseExtra(categoryID string, name string, ) *BotManagementCategoryResponseExtra`

NewBotManagementCategoryResponseExtra instantiates a new BotManagementCategoryResponseExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementCategoryResponseExtraWithDefaults

`func NewBotManagementCategoryResponseExtraWithDefaults() *BotManagementCategoryResponseExtra`

NewBotManagementCategoryResponseExtraWithDefaults instantiates a new BotManagementCategoryResponseExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryID

`func (o *BotManagementCategoryResponseExtra) GetCategoryID() string`

GetCategoryID returns the CategoryID field if non-nil, zero value otherwise.

### GetCategoryIDOk

`func (o *BotManagementCategoryResponseExtra) GetCategoryIDOk() (*string, bool)`

GetCategoryIDOk returns a tuple with the CategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryID

`func (o *BotManagementCategoryResponseExtra) SetCategoryID(v string)`

SetCategoryID sets CategoryID field to given value.


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
