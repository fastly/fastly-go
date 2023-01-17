# HistoricalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Whether or not we were able to successfully execute the query. | [optional] 
**Meta** | Pointer to [**HistoricalMeta**](HistoricalMeta.md) |  | [optional] 
**Msg** | Pointer to **NullableString** | If the query was not successful, this will provide a string that explains why. | [optional] 
**Data** | Pointer to [**map[string][]Results**](array.md) | Contains the results of the query, organized by *service ID*, into arrays where each element describes one service over a *time span*. | [optional] 

## Methods

### NewHistoricalResponse

`func NewHistoricalResponse() *HistoricalResponse`

NewHistoricalResponse instantiates a new HistoricalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalResponseWithDefaults

`func NewHistoricalResponseWithDefaults() *HistoricalResponse`

NewHistoricalResponseWithDefaults instantiates a new HistoricalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HistoricalResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HistoricalResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HistoricalResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HistoricalResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *HistoricalResponse) GetMeta() HistoricalMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HistoricalResponse) GetMetaOk() (*HistoricalMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HistoricalResponse) SetMeta(v HistoricalMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HistoricalResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMsg

`func (o *HistoricalResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *HistoricalResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *HistoricalResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *HistoricalResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *HistoricalResponse) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *HistoricalResponse) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetData

`func (o *HistoricalResponse) GetData() map[string][]Results`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalResponse) GetDataOk() (*map[string][]Results, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalResponse) SetData(v map[string][]Results)`

SetData sets Data field to given value.

### HasData

`func (o *HistoricalResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
