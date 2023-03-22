# HistoricalUsageAggregateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Whether or not we were able to successfully execute the query. | [optional] 
**Meta** | Pointer to [**HistoricalMeta**](HistoricalMeta.md) |  | [optional] 
**Msg** | Pointer to **NullableString** | If the query was not successful, this will provide a string that explains why. | [optional] 
**Data** | Pointer to [**HistoricalUsageResults**](HistoricalUsageResults.md) |  | [optional] 

## Methods

### NewHistoricalUsageAggregateResponse

`func NewHistoricalUsageAggregateResponse() *HistoricalUsageAggregateResponse`

NewHistoricalUsageAggregateResponse instantiates a new HistoricalUsageAggregateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalUsageAggregateResponseWithDefaults

`func NewHistoricalUsageAggregateResponseWithDefaults() *HistoricalUsageAggregateResponse`

NewHistoricalUsageAggregateResponseWithDefaults instantiates a new HistoricalUsageAggregateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HistoricalUsageAggregateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HistoricalUsageAggregateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HistoricalUsageAggregateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HistoricalUsageAggregateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *HistoricalUsageAggregateResponse) GetMeta() HistoricalMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HistoricalUsageAggregateResponse) GetMetaOk() (*HistoricalMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HistoricalUsageAggregateResponse) SetMeta(v HistoricalMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HistoricalUsageAggregateResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMsg

`func (o *HistoricalUsageAggregateResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *HistoricalUsageAggregateResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *HistoricalUsageAggregateResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *HistoricalUsageAggregateResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *HistoricalUsageAggregateResponse) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *HistoricalUsageAggregateResponse) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetData

`func (o *HistoricalUsageAggregateResponse) GetData() HistoricalUsageResults`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalUsageAggregateResponse) GetDataOk() (*HistoricalUsageResults, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalUsageAggregateResponse) SetData(v HistoricalUsageResults)`

SetData sets Data field to given value.

### HasData

`func (o *HistoricalUsageAggregateResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
