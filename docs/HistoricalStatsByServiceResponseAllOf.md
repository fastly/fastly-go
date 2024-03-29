# HistoricalStatsByServiceResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**map[string][]Results**](array.md) | Contains the results of the query, organized by *service ID*, into arrays where each element describes one service over a *time span*. | [optional] 

## Methods

### NewHistoricalStatsByServiceResponseAllOf

`func NewHistoricalStatsByServiceResponseAllOf() *HistoricalStatsByServiceResponseAllOf`

NewHistoricalStatsByServiceResponseAllOf instantiates a new HistoricalStatsByServiceResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalStatsByServiceResponseAllOfWithDefaults

`func NewHistoricalStatsByServiceResponseAllOfWithDefaults() *HistoricalStatsByServiceResponseAllOf`

NewHistoricalStatsByServiceResponseAllOfWithDefaults instantiates a new HistoricalStatsByServiceResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HistoricalStatsByServiceResponseAllOf) GetData() map[string][]Results`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalStatsByServiceResponseAllOf) GetDataOk() (*map[string][]Results, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalStatsByServiceResponseAllOf) SetData(v map[string][]Results)`

SetData sets Data field to given value.

### HasData

`func (o *HistoricalStatsByServiceResponseAllOf) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
