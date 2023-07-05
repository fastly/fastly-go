# Batch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A descriptor for the response of the entire batch | [optional] 
**Type** | Pointer to **string** | If an error is present in any of the requests, this field will describe that error | [optional] 
**Errors** | Pointer to [**[]BatchErrors**](BatchErrors.md) | Per-key errors which failed to parse, validate, or otherwise transmit | [optional] 

## Methods

### NewBatch

`func NewBatch() *Batch`

NewBatch instantiates a new Batch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchWithDefaults

`func NewBatchWithDefaults() *Batch`

NewBatchWithDefaults instantiates a new Batch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Batch) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Batch) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Batch) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Batch) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *Batch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Batch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Batch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Batch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetErrors

`func (o *Batch) GetErrors() []BatchErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Batch) GetErrorsOk() (*[]BatchErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Batch) SetErrors(v []BatchErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Batch) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
