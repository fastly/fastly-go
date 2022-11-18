# HistoricalUsageMonthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Whether or not we were able to successfully execute the query. | [optional] 
**Meta** | Pointer to [**HistoricalMeta**](HistoricalMeta.md) |  | [optional] 
**Msg** | Pointer to **NullableString** | If the query was not successful, this will provide a string that explains why. | [optional] 
**Data** | Pointer to [**HistoricalUsageMonthResponseAllOfData**](HistoricalUsageMonthResponseAllOfData.md) |  | [optional] 

## Methods

### NewHistoricalUsageMonthResponse

`func NewHistoricalUsageMonthResponse() *HistoricalUsageMonthResponse`

NewHistoricalUsageMonthResponse instantiates a new HistoricalUsageMonthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalUsageMonthResponseWithDefaults

`func NewHistoricalUsageMonthResponseWithDefaults() *HistoricalUsageMonthResponse`

NewHistoricalUsageMonthResponseWithDefaults instantiates a new HistoricalUsageMonthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HistoricalUsageMonthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HistoricalUsageMonthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HistoricalUsageMonthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HistoricalUsageMonthResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *HistoricalUsageMonthResponse) GetMeta() HistoricalMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HistoricalUsageMonthResponse) GetMetaOk() (*HistoricalMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HistoricalUsageMonthResponse) SetMeta(v HistoricalMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HistoricalUsageMonthResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMsg

`func (o *HistoricalUsageMonthResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *HistoricalUsageMonthResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *HistoricalUsageMonthResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *HistoricalUsageMonthResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *HistoricalUsageMonthResponse) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *HistoricalUsageMonthResponse) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetData

`func (o *HistoricalUsageMonthResponse) GetData() HistoricalUsageMonthResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalUsageMonthResponse) GetDataOk() (*HistoricalUsageMonthResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalUsageMonthResponse) SetData(v HistoricalUsageMonthResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *HistoricalUsageMonthResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
