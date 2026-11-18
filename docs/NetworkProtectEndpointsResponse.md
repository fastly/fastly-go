# NetworkProtectEndpointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]NetworkProtectEndpointResponse**](NetworkProtectEndpointResponse.md) | Endpoints returned on the current page. | 
**Meta** | [**NetworkProtectListMeta**](NetworkProtectListMeta.md) |  | 

## Methods

### NewNetworkProtectEndpointsResponse

`func NewNetworkProtectEndpointsResponse(data []NetworkProtectEndpointResponse, meta NetworkProtectListMeta, ) *NetworkProtectEndpointsResponse`

NewNetworkProtectEndpointsResponse instantiates a new NetworkProtectEndpointsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectEndpointsResponseWithDefaults

`func NewNetworkProtectEndpointsResponseWithDefaults() *NetworkProtectEndpointsResponse`

NewNetworkProtectEndpointsResponseWithDefaults instantiates a new NetworkProtectEndpointsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NetworkProtectEndpointsResponse) GetData() []NetworkProtectEndpointResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkProtectEndpointsResponse) GetDataOk() (*[]NetworkProtectEndpointResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkProtectEndpointsResponse) SetData(v []NetworkProtectEndpointResponse)`

SetData sets Data field to given value.


### GetMeta

`func (o *NetworkProtectEndpointsResponse) GetMeta() NetworkProtectListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkProtectEndpointsResponse) GetMetaOk() (*NetworkProtectListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkProtectEndpointsResponse) SetMeta(v NetworkProtectListMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
