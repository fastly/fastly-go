# TimeseriesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | Pointer to **map[string]map[string]any** | An object containing each requested dimension and time as properties. | [optional] 
**Values** | Pointer to **map[string]map[string]any** | An object containing each requested series as a property. | [optional] 

## Methods

### NewTimeseriesResult

`func NewTimeseriesResult() *TimeseriesResult`

NewTimeseriesResult instantiates a new TimeseriesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesResultWithDefaults

`func NewTimeseriesResultWithDefaults() *TimeseriesResult`

NewTimeseriesResultWithDefaults instantiates a new TimeseriesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *TimeseriesResult) GetDimensions() map[string]map[string]any`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *TimeseriesResult) GetDimensionsOk() (*map[string]map[string]any, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *TimeseriesResult) SetDimensions(v map[string]map[string]any)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *TimeseriesResult) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetValues

`func (o *TimeseriesResult) GetValues() map[string]map[string]any`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TimeseriesResult) GetValuesOk() (*map[string]map[string]any, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TimeseriesResult) SetValues(v map[string]map[string]any)`

SetValues sets Values field to given value.

### HasValues

`func (o *TimeseriesResult) HasValues() bool`

HasValues returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
