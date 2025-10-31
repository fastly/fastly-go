# HistoricalDomainsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | Pointer to [**DomainInspectorEntryDimensions**](DomainInspectorEntryDimensions.md) |  | [optional] 
**Values** | Pointer to [**[]Values**](Values.md) | An array of values representing the metric values at each point in time. Note that this dataset is sparse: only the keys with non-zero values will be included in the record.  | [optional] 

## Methods

### NewHistoricalDomainsData

`func NewHistoricalDomainsData() *HistoricalDomainsData`

NewHistoricalDomainsData instantiates a new HistoricalDomainsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalDomainsDataWithDefaults

`func NewHistoricalDomainsDataWithDefaults() *HistoricalDomainsData`

NewHistoricalDomainsDataWithDefaults instantiates a new HistoricalDomainsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *HistoricalDomainsData) GetDimensions() DomainInspectorEntryDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *HistoricalDomainsData) GetDimensionsOk() (*DomainInspectorEntryDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *HistoricalDomainsData) SetDimensions(v DomainInspectorEntryDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *HistoricalDomainsData) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetValues

`func (o *HistoricalDomainsData) GetValues() []Values`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HistoricalDomainsData) GetValuesOk() (*[]Values, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HistoricalDomainsData) SetValues(v []Values)`

SetValues sets Values field to given value.

### HasValues

`func (o *HistoricalDomainsData) HasValues() bool`

HasValues returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


