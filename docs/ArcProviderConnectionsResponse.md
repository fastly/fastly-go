# ArcProviderConnectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ArcProviderConnectionListItemResponse**](ArcProviderConnectionListItemResponse.md) | The provider connections returned by the request. | 
**Meta** | [**ArcProviderConnectionsMeta**](ArcProviderConnectionsMeta.md) |  | 

## Methods

### NewArcProviderConnectionsResponse

`func NewArcProviderConnectionsResponse(data []ArcProviderConnectionListItemResponse, meta ArcProviderConnectionsMeta, ) *ArcProviderConnectionsResponse`

NewArcProviderConnectionsResponse instantiates a new ArcProviderConnectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProviderConnectionsResponseWithDefaults

`func NewArcProviderConnectionsResponseWithDefaults() *ArcProviderConnectionsResponse`

NewArcProviderConnectionsResponseWithDefaults instantiates a new ArcProviderConnectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ArcProviderConnectionsResponse) GetData() []ArcProviderConnectionListItemResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ArcProviderConnectionsResponse) GetDataOk() (*[]ArcProviderConnectionListItemResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ArcProviderConnectionsResponse) SetData(v []ArcProviderConnectionListItemResponse)`

SetData sets Data field to given value.


### GetMeta

`func (o *ArcProviderConnectionsResponse) GetMeta() ArcProviderConnectionsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ArcProviderConnectionsResponse) GetMetaOk() (*ArcProviderConnectionsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ArcProviderConnectionsResponse) SetMeta(v ArcProviderConnectionsMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
