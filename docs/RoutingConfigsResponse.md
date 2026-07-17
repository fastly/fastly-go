# RoutingConfigsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RoutingConfigResponse**](RoutingConfigResponse.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewRoutingConfigsResponse

`func NewRoutingConfigsResponse() *RoutingConfigsResponse`

NewRoutingConfigsResponse instantiates a new RoutingConfigsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingConfigsResponseWithDefaults

`func NewRoutingConfigsResponseWithDefaults() *RoutingConfigsResponse`

NewRoutingConfigsResponseWithDefaults instantiates a new RoutingConfigsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RoutingConfigsResponse) GetData() []RoutingConfigResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RoutingConfigsResponse) GetDataOk() (*[]RoutingConfigResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RoutingConfigsResponse) SetData(v []RoutingConfigResponse)`

SetData sets Data field to given value.

### HasData

`func (o *RoutingConfigsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *RoutingConfigsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RoutingConfigsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RoutingConfigsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RoutingConfigsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


