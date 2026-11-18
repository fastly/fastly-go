# BotManagementNotAuthenticatedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Status** | **int32** |  | 
**Detail** | Pointer to **string** | An explanation of the problem, specific to this occurrence. | [optional] 

## Methods

### NewBotManagementNotAuthenticatedError

`func NewBotManagementNotAuthenticatedError(title string, status int32, ) *BotManagementNotAuthenticatedError`

NewBotManagementNotAuthenticatedError instantiates a new BotManagementNotAuthenticatedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementNotAuthenticatedErrorWithDefaults

`func NewBotManagementNotAuthenticatedErrorWithDefaults() *BotManagementNotAuthenticatedError`

NewBotManagementNotAuthenticatedErrorWithDefaults instantiates a new BotManagementNotAuthenticatedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *BotManagementNotAuthenticatedError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BotManagementNotAuthenticatedError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BotManagementNotAuthenticatedError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *BotManagementNotAuthenticatedError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BotManagementNotAuthenticatedError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BotManagementNotAuthenticatedError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *BotManagementNotAuthenticatedError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BotManagementNotAuthenticatedError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BotManagementNotAuthenticatedError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *BotManagementNotAuthenticatedError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
