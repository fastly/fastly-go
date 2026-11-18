# BotRequestMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BotRequestMetricsDataItem**](BotRequestMetricsDataItem.md) | An array of time series data items containing metric counts. Items are grouped by `timestamp` by default, or by a combination of `timestamp` and `bot` or `datacenter` or `region` when the `group_by` parameter is set. This dataset is sparse: only keys with non-zero values are included.  | [optional] 
**Meta** | Pointer to [**BotRequestMetricsMetadata**](BotRequestMetricsMetadata.md) |  | [optional] 

## Methods

### NewBotRequestMetricsResponse

`func NewBotRequestMetricsResponse() *BotRequestMetricsResponse`

NewBotRequestMetricsResponse instantiates a new BotRequestMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotRequestMetricsResponseWithDefaults

`func NewBotRequestMetricsResponseWithDefaults() *BotRequestMetricsResponse`

NewBotRequestMetricsResponseWithDefaults instantiates a new BotRequestMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BotRequestMetricsResponse) GetData() []BotRequestMetricsDataItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BotRequestMetricsResponse) GetDataOk() (*[]BotRequestMetricsDataItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BotRequestMetricsResponse) SetData(v []BotRequestMetricsDataItem)`

SetData sets Data field to given value.

### HasData

`func (o *BotRequestMetricsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *BotRequestMetricsResponse) GetMeta() BotRequestMetricsMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BotRequestMetricsResponse) GetMetaOk() (*BotRequestMetricsMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BotRequestMetricsResponse) SetMeta(v BotRequestMetricsMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BotRequestMetricsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
