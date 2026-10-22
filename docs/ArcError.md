# ArcError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A URL identifying the error type. | [optional] 
**Title** | Pointer to **string** | A short summary of the problem. | [optional] 
**Code** | Pointer to **string** | A machine-readable error code. | [optional] 

## Methods

### NewArcError

`func NewArcError() *ArcError`

NewArcError instantiates a new ArcError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcErrorWithDefaults

`func NewArcErrorWithDefaults() *ArcError`

NewArcErrorWithDefaults instantiates a new ArcError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArcError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArcError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArcError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ArcError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *ArcError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArcError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArcError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ArcError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCode

`func (o *ArcError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ArcError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ArcError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ArcError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


