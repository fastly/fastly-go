# HistoricalDdos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Whether or not we were able to successfully execute the query. | [optional] 
**Meta** | Pointer to [**HistoricalDdosMeta**](HistoricalDdosMeta.md) |  | [optional] 
**Msg** | Pointer to **NullableString** | If the query was not successful, this will provide a string that explains why. | [optional] 
**Data** | Pointer to [**[]PlatformDdosDataItems**](PlatformDdosDataItems.md) | A list of [entries](#entry-data-model). | [optional] 

## Methods

### NewHistoricalDdos

`func NewHistoricalDdos() *HistoricalDdos`

NewHistoricalDdos instantiates a new HistoricalDdos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalDdosWithDefaults

`func NewHistoricalDdosWithDefaults() *HistoricalDdos`

NewHistoricalDdosWithDefaults instantiates a new HistoricalDdos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HistoricalDdos) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HistoricalDdos) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HistoricalDdos) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HistoricalDdos) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *HistoricalDdos) GetMeta() HistoricalDdosMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HistoricalDdos) GetMetaOk() (*HistoricalDdosMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HistoricalDdos) SetMeta(v HistoricalDdosMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HistoricalDdos) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMsg

`func (o *HistoricalDdos) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *HistoricalDdos) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *HistoricalDdos) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *HistoricalDdos) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *HistoricalDdos) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *HistoricalDdos) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetData

`func (o *HistoricalDdos) GetData() []PlatformDdosDataItems`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalDdos) GetDataOk() (*[]PlatformDdosDataItems, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalDdos) SetData(v []PlatformDdosDataItems)`

SetData sets Data field to given value.

### HasData

`func (o *HistoricalDdos) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
