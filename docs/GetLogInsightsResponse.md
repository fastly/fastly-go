# GetLogInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]LogInsights**](LogInsights.md) |  | [optional] 
**Meta** | Pointer to [**LogInsightsMeta**](LogInsightsMeta.md) |  | [optional] 

## Methods

### NewGetLogInsightsResponse

`func NewGetLogInsightsResponse() *GetLogInsightsResponse`

NewGetLogInsightsResponse instantiates a new GetLogInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogInsightsResponseWithDefaults

`func NewGetLogInsightsResponseWithDefaults() *GetLogInsightsResponse`

NewGetLogInsightsResponseWithDefaults instantiates a new GetLogInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetLogInsightsResponse) GetData() []LogInsights`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLogInsightsResponse) GetDataOk() (*[]LogInsights, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLogInsightsResponse) SetData(v []LogInsights)`

SetData sets Data field to given value.

### HasData

`func (o *GetLogInsightsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetLogInsightsResponse) GetMeta() LogInsightsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetLogInsightsResponse) GetMetaOk() (*LogInsightsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetLogInsightsResponse) SetMeta(v LogInsightsMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetLogInsightsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
