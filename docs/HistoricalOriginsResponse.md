# HistoricalOriginsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Whether or not we were able to successfully execute the query. | [optional] 
**Meta** | Pointer to [**OriginInspectorHistoricalMeta**](OriginInspectorHistoricalMeta.md) |  | [optional] 
**Msg** | Pointer to **NullableString** | If the query was not successful, this will provide a string that explains why. | [optional] 
**Data** | Pointer to [**[]OriginInspectorEntry**](OriginInspectorEntry.md) | A list of timeseries. Each individual timeseries represents a unique combination of dimensions, such as origin host, region or POP. | [optional] 

## Methods

### NewHistoricalOriginsResponse

`func NewHistoricalOriginsResponse() *HistoricalOriginsResponse`

NewHistoricalOriginsResponse instantiates a new HistoricalOriginsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalOriginsResponseWithDefaults

`func NewHistoricalOriginsResponseWithDefaults() *HistoricalOriginsResponse`

NewHistoricalOriginsResponseWithDefaults instantiates a new HistoricalOriginsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HistoricalOriginsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HistoricalOriginsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HistoricalOriginsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HistoricalOriginsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *HistoricalOriginsResponse) GetMeta() OriginInspectorHistoricalMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HistoricalOriginsResponse) GetMetaOk() (*OriginInspectorHistoricalMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HistoricalOriginsResponse) SetMeta(v OriginInspectorHistoricalMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HistoricalOriginsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMsg

`func (o *HistoricalOriginsResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *HistoricalOriginsResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *HistoricalOriginsResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *HistoricalOriginsResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *HistoricalOriginsResponse) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *HistoricalOriginsResponse) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetData

`func (o *HistoricalOriginsResponse) GetData() []OriginInspectorEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalOriginsResponse) GetDataOk() (*[]OriginInspectorEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalOriginsResponse) SetData(v []OriginInspectorEntry)`

SetData sets Data field to given value.

### HasData

`func (o *HistoricalOriginsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


