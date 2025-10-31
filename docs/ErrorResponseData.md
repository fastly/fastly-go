# ErrorResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorResponseData

`func NewErrorResponseData() *ErrorResponseData`

NewErrorResponseData instantiates a new ErrorResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseDataWithDefaults

`func NewErrorResponseDataWithDefaults() *ErrorResponseData`

NewErrorResponseDataWithDefaults instantiates a new ErrorResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ErrorResponseData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorResponseData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorResponseData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ErrorResponseData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *ErrorResponseData) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorResponseData) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorResponseData) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorResponseData) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


