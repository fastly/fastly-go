# BatchErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key that the error corresponds to. This field will be empty if the object or one of its fields was unable to be parsed. | [optional] 
**Index** | Pointer to **int32** | The line number of the payload on which the error occurred (starting from 0 for the first line). | [optional] 
**Code** | Pointer to **string** | The HTTP response code for the request, or a 400 if the request was not able to be completed. | [optional] 
**Reason** | Pointer to **string** | A descriptor of this particular item&#39;s error. | [optional] 

## Methods

### NewBatchErrors

`func NewBatchErrors() *BatchErrors`

NewBatchErrors instantiates a new BatchErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchErrorsWithDefaults

`func NewBatchErrorsWithDefaults() *BatchErrors`

NewBatchErrorsWithDefaults instantiates a new BatchErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *BatchErrors) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BatchErrors) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BatchErrors) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BatchErrors) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetIndex

`func (o *BatchErrors) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BatchErrors) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BatchErrors) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *BatchErrors) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetCode

`func (o *BatchErrors) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BatchErrors) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BatchErrors) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BatchErrors) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *BatchErrors) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BatchErrors) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BatchErrors) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BatchErrors) HasReason() bool`

HasReason returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
