# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]any** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorResponse

`func NewErrorResponse() *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *ErrorResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorResponse) GetErrors() []any`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorResponse) GetErrorsOk() (*[]any, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorResponse) SetErrors(v []any)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *ErrorResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *ErrorResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ErrorResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
