# BotManagementRateLimitError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Status** | **int32** |  | 
**Detail** | Pointer to **string** | An explanation of the problem, specific to this occurrence. | [optional] 

## Methods

### NewBotManagementRateLimitError

`func NewBotManagementRateLimitError(title string, status int32, ) *BotManagementRateLimitError`

NewBotManagementRateLimitError instantiates a new BotManagementRateLimitError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementRateLimitErrorWithDefaults

`func NewBotManagementRateLimitErrorWithDefaults() *BotManagementRateLimitError`

NewBotManagementRateLimitErrorWithDefaults instantiates a new BotManagementRateLimitError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *BotManagementRateLimitError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BotManagementRateLimitError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BotManagementRateLimitError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *BotManagementRateLimitError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BotManagementRateLimitError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BotManagementRateLimitError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *BotManagementRateLimitError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BotManagementRateLimitError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BotManagementRateLimitError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *BotManagementRateLimitError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


