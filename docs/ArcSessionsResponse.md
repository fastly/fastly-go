# ArcSessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ArcSession**](ArcSession.md) | The sessions returned by the request. | 
**Meta** | [**ArcSessionsMeta**](ArcSessionsMeta.md) |  | 

## Methods

### NewArcSessionsResponse

`func NewArcSessionsResponse(data []ArcSession, meta ArcSessionsMeta, ) *ArcSessionsResponse`

NewArcSessionsResponse instantiates a new ArcSessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcSessionsResponseWithDefaults

`func NewArcSessionsResponseWithDefaults() *ArcSessionsResponse`

NewArcSessionsResponseWithDefaults instantiates a new ArcSessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ArcSessionsResponse) GetData() []ArcSession`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ArcSessionsResponse) GetDataOk() (*[]ArcSession, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ArcSessionsResponse) SetData(v []ArcSession)`

SetData sets Data field to given value.


### GetMeta

`func (o *ArcSessionsResponse) GetMeta() ArcSessionsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ArcSessionsResponse) GetMetaOk() (*ArcSessionsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ArcSessionsResponse) SetMeta(v ArcSessionsMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


