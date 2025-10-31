# OriginInspectorHistorical

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Whether or not we were able to successfully execute the query. | [optional] 
**Meta** | Pointer to [**OriginInspectorHistoricalMeta**](OriginInspectorHistoricalMeta.md) |  | [optional] 
**Msg** | Pointer to **NullableString** | If the query was not successful, this will provide a string that explains why. | [optional] 
**Data** | Pointer to [**[]OriginInspectorHistoricalData**](OriginInspectorHistoricalData.md) | A list of [entries](#entry-data-model), each representing one unique combination of dimensions, such as origin host, region or POP. | [optional] 

## Methods

### NewOriginInspectorHistorical

`func NewOriginInspectorHistorical() *OriginInspectorHistorical`

NewOriginInspectorHistorical instantiates a new OriginInspectorHistorical object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginInspectorHistoricalWithDefaults

`func NewOriginInspectorHistoricalWithDefaults() *OriginInspectorHistorical`

NewOriginInspectorHistoricalWithDefaults instantiates a new OriginInspectorHistorical object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *OriginInspectorHistorical) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OriginInspectorHistorical) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OriginInspectorHistorical) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OriginInspectorHistorical) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *OriginInspectorHistorical) GetMeta() OriginInspectorHistoricalMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OriginInspectorHistorical) GetMetaOk() (*OriginInspectorHistoricalMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OriginInspectorHistorical) SetMeta(v OriginInspectorHistoricalMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OriginInspectorHistorical) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMsg

`func (o *OriginInspectorHistorical) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OriginInspectorHistorical) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OriginInspectorHistorical) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OriginInspectorHistorical) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *OriginInspectorHistorical) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *OriginInspectorHistorical) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetData

`func (o *OriginInspectorHistorical) GetData() []OriginInspectorHistoricalData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OriginInspectorHistorical) GetDataOk() (*[]OriginInspectorHistoricalData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OriginInspectorHistorical) SetData(v []OriginInspectorHistoricalData)`

SetData sets Data field to given value.

### HasData

`func (o *OriginInspectorHistorical) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


