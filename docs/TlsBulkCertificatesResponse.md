# TLSBulkCertificatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]TLSBulkCertificateResponseData**](TlsBulkCertificateResponseData.md) |  | [optional] 

## Methods

### NewTLSBulkCertificatesResponse

`func NewTLSBulkCertificatesResponse() *TLSBulkCertificatesResponse`

NewTLSBulkCertificatesResponse instantiates a new TLSBulkCertificatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSBulkCertificatesResponseWithDefaults

`func NewTLSBulkCertificatesResponseWithDefaults() *TLSBulkCertificatesResponse`

NewTLSBulkCertificatesResponseWithDefaults instantiates a new TLSBulkCertificatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TLSBulkCertificatesResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TLSBulkCertificatesResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TLSBulkCertificatesResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TLSBulkCertificatesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *TLSBulkCertificatesResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TLSBulkCertificatesResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TLSBulkCertificatesResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TLSBulkCertificatesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *TLSBulkCertificatesResponse) GetData() []TLSBulkCertificateResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TLSBulkCertificatesResponse) GetDataOk() (*[]TLSBulkCertificateResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TLSBulkCertificatesResponse) SetData(v []TLSBulkCertificateResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *TLSBulkCertificatesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
