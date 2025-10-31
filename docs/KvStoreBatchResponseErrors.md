# KvStoreBatchResponseErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key that the error corresponds to. This field will be empty if the object or one of its fields was not parseable. | [optional] 
**Index** | Pointer to **int32** | The line number of the batch request body on which the error occurred (starting from 0 for the first line). | [optional] 
**Code** | Pointer to **string** | The HTTP response code for the item, or a 400 if the item&#39;s operation was not completed. | [optional] 
**Reason** | Pointer to **string** | A descriptor of this particular item&#39;s error. | [optional] 

## Methods

### NewKvStoreBatchResponseErrors

`func NewKvStoreBatchResponseErrors() *KvStoreBatchResponseErrors`

NewKvStoreBatchResponseErrors instantiates a new KvStoreBatchResponseErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvStoreBatchResponseErrorsWithDefaults

`func NewKvStoreBatchResponseErrorsWithDefaults() *KvStoreBatchResponseErrors`

NewKvStoreBatchResponseErrorsWithDefaults instantiates a new KvStoreBatchResponseErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KvStoreBatchResponseErrors) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KvStoreBatchResponseErrors) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KvStoreBatchResponseErrors) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KvStoreBatchResponseErrors) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetIndex

`func (o *KvStoreBatchResponseErrors) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *KvStoreBatchResponseErrors) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *KvStoreBatchResponseErrors) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *KvStoreBatchResponseErrors) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetCode

`func (o *KvStoreBatchResponseErrors) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *KvStoreBatchResponseErrors) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *KvStoreBatchResponseErrors) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *KvStoreBatchResponseErrors) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *KvStoreBatchResponseErrors) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *KvStoreBatchResponseErrors) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *KvStoreBatchResponseErrors) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *KvStoreBatchResponseErrors) HasReason() bool`

HasReason returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


