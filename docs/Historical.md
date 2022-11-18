# Historical

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Whether or not we were able to successfully execute the query. | [optional] 
**Meta** | Pointer to [**HistoricalMeta**](HistoricalMeta.md) |  | [optional] 
**Msg** | Pointer to **NullableString** | If the query was not successful, this will provide a string that explains why. | [optional] 

## Methods

### NewHistorical

`func NewHistorical() *Historical`

NewHistorical instantiates a new Historical object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalWithDefaults

`func NewHistoricalWithDefaults() *Historical`

NewHistoricalWithDefaults instantiates a new Historical object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Historical) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Historical) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Historical) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Historical) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *Historical) GetMeta() HistoricalMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Historical) GetMetaOk() (*HistoricalMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Historical) SetMeta(v HistoricalMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Historical) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMsg

`func (o *Historical) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *Historical) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *Historical) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *Historical) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *Historical) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *Historical) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
