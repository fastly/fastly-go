# ListEomInvoicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Invoice**](Invoice.md) |  | [optional] 
**Meta** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewListEomInvoicesResponse

`func NewListEomInvoicesResponse() *ListEomInvoicesResponse`

NewListEomInvoicesResponse instantiates a new ListEomInvoicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEomInvoicesResponseWithDefaults

`func NewListEomInvoicesResponseWithDefaults() *ListEomInvoicesResponse`

NewListEomInvoicesResponseWithDefaults instantiates a new ListEomInvoicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListEomInvoicesResponse) GetData() []Invoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEomInvoicesResponse) GetDataOk() (*[]Invoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEomInvoicesResponse) SetData(v []Invoice)`

SetData sets Data field to given value.

### HasData

`func (o *ListEomInvoicesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListEomInvoicesResponse) GetMeta() Metadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEomInvoicesResponse) GetMetaOk() (*Metadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEomInvoicesResponse) SetMeta(v Metadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEomInvoicesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
