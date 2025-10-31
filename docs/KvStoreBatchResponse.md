# KvStoreBatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A descriptor for the response of the entire batch | [optional] 
**Type** | Pointer to **string** | If an error is present in any of the requests, this field will describe that error | [optional] 
**Errors** | Pointer to [**[]KvStoreBatchResponseErrors**](KvStoreBatchResponseErrors.md) | Errors which occurred while handling the items in the request | [optional] 

## Methods

### NewKvStoreBatchResponse

`func NewKvStoreBatchResponse() *KvStoreBatchResponse`

NewKvStoreBatchResponse instantiates a new KvStoreBatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvStoreBatchResponseWithDefaults

`func NewKvStoreBatchResponseWithDefaults() *KvStoreBatchResponse`

NewKvStoreBatchResponseWithDefaults instantiates a new KvStoreBatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *KvStoreBatchResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *KvStoreBatchResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *KvStoreBatchResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *KvStoreBatchResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *KvStoreBatchResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KvStoreBatchResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KvStoreBatchResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KvStoreBatchResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetErrors

`func (o *KvStoreBatchResponse) GetErrors() []KvStoreBatchResponseErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *KvStoreBatchResponse) GetErrorsOk() (*[]KvStoreBatchResponseErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *KvStoreBatchResponse) SetErrors(v []KvStoreBatchResponseErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *KvStoreBatchResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


