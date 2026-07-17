# VersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RoutingConfigVersionResponse**](RoutingConfigVersionResponse.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewVersionsResponse

`func NewVersionsResponse() *VersionsResponse`

NewVersionsResponse instantiates a new VersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionsResponseWithDefaults

`func NewVersionsResponseWithDefaults() *VersionsResponse`

NewVersionsResponseWithDefaults instantiates a new VersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VersionsResponse) GetData() []RoutingConfigVersionResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VersionsResponse) GetDataOk() (*[]RoutingConfigVersionResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VersionsResponse) SetData(v []RoutingConfigVersionResponse)`

SetData sets Data field to given value.

### HasData

`func (o *VersionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *VersionsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VersionsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VersionsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VersionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


