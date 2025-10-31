# TlsDomainsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]TlsDomainData**](TlsDomainData.md) |  | [optional] 

## Methods

### NewTlsDomainsResponse

`func NewTlsDomainsResponse() *TlsDomainsResponse`

NewTlsDomainsResponse instantiates a new TlsDomainsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsDomainsResponseWithDefaults

`func NewTlsDomainsResponseWithDefaults() *TlsDomainsResponse`

NewTlsDomainsResponseWithDefaults instantiates a new TlsDomainsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TlsDomainsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TlsDomainsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TlsDomainsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TlsDomainsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *TlsDomainsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TlsDomainsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TlsDomainsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TlsDomainsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *TlsDomainsResponse) GetData() []TlsDomainData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TlsDomainsResponse) GetDataOk() (*[]TlsDomainData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TlsDomainsResponse) SetData(v []TlsDomainData)`

SetData sets Data field to given value.

### HasData

`func (o *TlsDomainsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


