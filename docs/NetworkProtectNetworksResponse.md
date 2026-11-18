# NetworkProtectNetworksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]NetworkProtectNetworkResponse**](NetworkProtectNetworkResponse.md) | Networks returned on the current page. | 
**Meta** | [**NetworkProtectListMeta**](NetworkProtectListMeta.md) |  | 

## Methods

### NewNetworkProtectNetworksResponse

`func NewNetworkProtectNetworksResponse(data []NetworkProtectNetworkResponse, meta NetworkProtectListMeta, ) *NetworkProtectNetworksResponse`

NewNetworkProtectNetworksResponse instantiates a new NetworkProtectNetworksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectNetworksResponseWithDefaults

`func NewNetworkProtectNetworksResponseWithDefaults() *NetworkProtectNetworksResponse`

NewNetworkProtectNetworksResponseWithDefaults instantiates a new NetworkProtectNetworksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NetworkProtectNetworksResponse) GetData() []NetworkProtectNetworkResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkProtectNetworksResponse) GetDataOk() (*[]NetworkProtectNetworkResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkProtectNetworksResponse) SetData(v []NetworkProtectNetworkResponse)`

SetData sets Data field to given value.


### GetMeta

`func (o *NetworkProtectNetworksResponse) GetMeta() NetworkProtectListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkProtectNetworksResponse) GetMetaOk() (*NetworkProtectListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkProtectNetworksResponse) SetMeta(v NetworkProtectListMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
