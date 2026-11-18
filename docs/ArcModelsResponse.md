# ArcModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ArcModel**](ArcModel.md) | The models returned by the request. | 
**Meta** | [**ArcModelsMeta**](ArcModelsMeta.md) |  | 

## Methods

### NewArcModelsResponse

`func NewArcModelsResponse(data []ArcModel, meta ArcModelsMeta, ) *ArcModelsResponse`

NewArcModelsResponse instantiates a new ArcModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcModelsResponseWithDefaults

`func NewArcModelsResponseWithDefaults() *ArcModelsResponse`

NewArcModelsResponseWithDefaults instantiates a new ArcModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ArcModelsResponse) GetData() []ArcModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ArcModelsResponse) GetDataOk() (*[]ArcModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ArcModelsResponse) SetData(v []ArcModel)`

SetData sets Data field to given value.


### GetMeta

`func (o *ArcModelsResponse) GetMeta() ArcModelsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ArcModelsResponse) GetMetaOk() (*ArcModelsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ArcModelsResponse) SetMeta(v ArcModelsMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
