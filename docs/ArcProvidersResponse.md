# ArcProvidersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ArcProvider**](ArcProvider.md) | The providers returned by the request. | 
**Meta** | [**ArcProvidersMeta**](ArcProvidersMeta.md) |  | 

## Methods

### NewArcProvidersResponse

`func NewArcProvidersResponse(data []ArcProvider, meta ArcProvidersMeta, ) *ArcProvidersResponse`

NewArcProvidersResponse instantiates a new ArcProvidersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProvidersResponseWithDefaults

`func NewArcProvidersResponseWithDefaults() *ArcProvidersResponse`

NewArcProvidersResponseWithDefaults instantiates a new ArcProvidersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ArcProvidersResponse) GetData() []ArcProvider`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ArcProvidersResponse) GetDataOk() (*[]ArcProvider, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ArcProvidersResponse) SetData(v []ArcProvider)`

SetData sets Data field to given value.


### GetMeta

`func (o *ArcProvidersResponse) GetMeta() ArcProvidersMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ArcProvidersResponse) GetMetaOk() (*ArcProvidersMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ArcProvidersResponse) SetMeta(v ArcProvidersMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


