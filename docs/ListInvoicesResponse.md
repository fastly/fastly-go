# ListInvoicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Invoice**](Invoice.md) |  | [optional] 
**Meta** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewListInvoicesResponse

`func NewListInvoicesResponse() *ListInvoicesResponse`

NewListInvoicesResponse instantiates a new ListInvoicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoicesResponseWithDefaults

`func NewListInvoicesResponseWithDefaults() *ListInvoicesResponse`

NewListInvoicesResponseWithDefaults instantiates a new ListInvoicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListInvoicesResponse) GetData() []Invoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListInvoicesResponse) GetDataOk() (*[]Invoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListInvoicesResponse) SetData(v []Invoice)`

SetData sets Data field to given value.

### HasData

`func (o *ListInvoicesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListInvoicesResponse) GetMeta() Metadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInvoicesResponse) GetMetaOk() (*Metadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInvoicesResponse) SetMeta(v Metadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInvoicesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
