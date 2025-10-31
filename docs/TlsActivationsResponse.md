# TlsActivationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]TlsActivationResponseData**](TlsActivationResponseData.md) |  | [optional] 

## Methods

### NewTlsActivationsResponse

`func NewTlsActivationsResponse() *TlsActivationsResponse`

NewTlsActivationsResponse instantiates a new TlsActivationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsActivationsResponseWithDefaults

`func NewTlsActivationsResponseWithDefaults() *TlsActivationsResponse`

NewTlsActivationsResponseWithDefaults instantiates a new TlsActivationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TlsActivationsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TlsActivationsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TlsActivationsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TlsActivationsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *TlsActivationsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TlsActivationsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TlsActivationsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TlsActivationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *TlsActivationsResponse) GetData() []TlsActivationResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TlsActivationsResponse) GetDataOk() (*[]TlsActivationResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TlsActivationsResponse) SetData(v []TlsActivationResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *TlsActivationsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


