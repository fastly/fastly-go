# BotManagementCategoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**BotManagementMeta**](BotManagementMeta.md) |  | 
**Data** | [**[]BotManagementCategoryResponse**](BotManagementCategoryResponse.md) | The categories returned by the request. | 

## Methods

### NewBotManagementCategoriesResponse

`func NewBotManagementCategoriesResponse(meta BotManagementMeta, data []BotManagementCategoryResponse, ) *BotManagementCategoriesResponse`

NewBotManagementCategoriesResponse instantiates a new BotManagementCategoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementCategoriesResponseWithDefaults

`func NewBotManagementCategoriesResponseWithDefaults() *BotManagementCategoriesResponse`

NewBotManagementCategoriesResponseWithDefaults instantiates a new BotManagementCategoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *BotManagementCategoriesResponse) GetMeta() BotManagementMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BotManagementCategoriesResponse) GetMetaOk() (*BotManagementMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BotManagementCategoriesResponse) SetMeta(v BotManagementMeta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *BotManagementCategoriesResponse) GetData() []BotManagementCategoryResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BotManagementCategoriesResponse) GetDataOk() (*[]BotManagementCategoryResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BotManagementCategoriesResponse) SetData(v []BotManagementCategoryResponse)`

SetData sets Data field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


