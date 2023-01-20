# GetStoresResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]StoreResponse**](StoreResponse.md) |  | [optional] 
**Meta** | Pointer to [**GetStoresResponseMeta**](GetStoresResponseMeta.md) |  | [optional] 

## Methods

### NewGetStoresResponse

`func NewGetStoresResponse() *GetStoresResponse`

NewGetStoresResponse instantiates a new GetStoresResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoresResponseWithDefaults

`func NewGetStoresResponseWithDefaults() *GetStoresResponse`

NewGetStoresResponseWithDefaults instantiates a new GetStoresResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetStoresResponse) GetData() []StoreResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetStoresResponse) GetDataOk() (*[]StoreResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetStoresResponse) SetData(v []StoreResponse)`

SetData sets Data field to given value.

### HasData

`func (o *GetStoresResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetStoresResponse) GetMeta() GetStoresResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetStoresResponse) GetMetaOk() (*GetStoresResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetStoresResponse) SetMeta(v GetStoresResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetStoresResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
