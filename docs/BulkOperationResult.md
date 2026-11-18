# BulkOperationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | The operation ID. | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code for this operation. | [optional] 
**Reason** | Pointer to **string** | Error reason if the operation failed. | [optional] 

## Methods

### NewBulkOperationResult

`func NewBulkOperationResult() *BulkOperationResult`

NewBulkOperationResult instantiates a new BulkOperationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkOperationResultWithDefaults

`func NewBulkOperationResultWithDefaults() *BulkOperationResult`

NewBulkOperationResultWithDefaults instantiates a new BulkOperationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *BulkOperationResult) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *BulkOperationResult) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *BulkOperationResult) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *BulkOperationResult) HasID() bool`

HasID returns a boolean if a field has been set.

### GetStatusCode

`func (o *BulkOperationResult) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BulkOperationResult) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BulkOperationResult) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BulkOperationResult) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetReason

`func (o *BulkOperationResult) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BulkOperationResult) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BulkOperationResult) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BulkOperationResult) HasReason() bool`

HasReason returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
