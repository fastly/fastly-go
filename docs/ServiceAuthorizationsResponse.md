# ServiceAuthorizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]ServiceAuthorizationResponseData**](ServiceAuthorizationResponseData.md) |  | [optional] 

## Methods

### NewServiceAuthorizationsResponse

`func NewServiceAuthorizationsResponse() *ServiceAuthorizationsResponse`

NewServiceAuthorizationsResponse instantiates a new ServiceAuthorizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAuthorizationsResponseWithDefaults

`func NewServiceAuthorizationsResponseWithDefaults() *ServiceAuthorizationsResponse`

NewServiceAuthorizationsResponseWithDefaults instantiates a new ServiceAuthorizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ServiceAuthorizationsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceAuthorizationsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceAuthorizationsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceAuthorizationsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *ServiceAuthorizationsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServiceAuthorizationsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServiceAuthorizationsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ServiceAuthorizationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *ServiceAuthorizationsResponse) GetData() []ServiceAuthorizationResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceAuthorizationsResponse) GetDataOk() (*[]ServiceAuthorizationResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceAuthorizationsResponse) SetData(v []ServiceAuthorizationResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceAuthorizationsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
