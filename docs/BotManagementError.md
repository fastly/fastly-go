# BotManagementError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | A short summary of the problem. | 
**Status** | **int32** | The HTTP status code of the response. | 
**Detail** | Pointer to **string** | An explanation of the problem, specific to this occurrence. | [optional] 

## Methods

### NewBotManagementError

`func NewBotManagementError(title string, status int32, ) *BotManagementError`

NewBotManagementError instantiates a new BotManagementError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementErrorWithDefaults

`func NewBotManagementErrorWithDefaults() *BotManagementError`

NewBotManagementErrorWithDefaults instantiates a new BotManagementError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *BotManagementError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BotManagementError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BotManagementError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *BotManagementError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BotManagementError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BotManagementError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *BotManagementError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BotManagementError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BotManagementError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *BotManagementError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
