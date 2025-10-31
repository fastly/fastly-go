# HistoricalOriginsResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]OriginInspectorEntry**](OriginInspectorEntry.md) | A list of timeseries. Each individual timeseries represents a unique combination of dimensions, such as origin host, region or POP. | [optional] 

## Methods

### NewHistoricalOriginsResponseAllOf

`func NewHistoricalOriginsResponseAllOf() *HistoricalOriginsResponseAllOf`

NewHistoricalOriginsResponseAllOf instantiates a new HistoricalOriginsResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalOriginsResponseAllOfWithDefaults

`func NewHistoricalOriginsResponseAllOfWithDefaults() *HistoricalOriginsResponseAllOf`

NewHistoricalOriginsResponseAllOfWithDefaults instantiates a new HistoricalOriginsResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HistoricalOriginsResponseAllOf) GetData() []OriginInspectorEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalOriginsResponseAllOf) GetDataOk() (*[]OriginInspectorEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalOriginsResponseAllOf) SetData(v []OriginInspectorEntry)`

SetData sets Data field to given value.

### HasData

`func (o *HistoricalOriginsResponseAllOf) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


