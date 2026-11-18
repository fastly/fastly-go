# NetworkProtectPrefixListsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]NetworkProtectPrefixListResponse**](NetworkProtectPrefixListResponse.md) | Prefix lists returned on the current page. | 
**Meta** | [**NetworkProtectListMeta**](NetworkProtectListMeta.md) |  | 

## Methods

### NewNetworkProtectPrefixListsResponse

`func NewNetworkProtectPrefixListsResponse(data []NetworkProtectPrefixListResponse, meta NetworkProtectListMeta, ) *NetworkProtectPrefixListsResponse`

NewNetworkProtectPrefixListsResponse instantiates a new NetworkProtectPrefixListsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectPrefixListsResponseWithDefaults

`func NewNetworkProtectPrefixListsResponseWithDefaults() *NetworkProtectPrefixListsResponse`

NewNetworkProtectPrefixListsResponseWithDefaults instantiates a new NetworkProtectPrefixListsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NetworkProtectPrefixListsResponse) GetData() []NetworkProtectPrefixListResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkProtectPrefixListsResponse) GetDataOk() (*[]NetworkProtectPrefixListResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkProtectPrefixListsResponse) SetData(v []NetworkProtectPrefixListResponse)`

SetData sets Data field to given value.


### GetMeta

`func (o *NetworkProtectPrefixListsResponse) GetMeta() NetworkProtectListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkProtectPrefixListsResponse) GetMetaOk() (*NetworkProtectListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkProtectPrefixListsResponse) SetMeta(v NetworkProtectListMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
