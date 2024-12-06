# LogInsights

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | Pointer to [**LogInsightsDimensions**](LogInsightsDimensions.md) |  | [optional] 
**DimensionAttributes** | Pointer to [**LogInsightsDimensionAttributes**](LogInsightsDimensionAttributes.md) |  | [optional] 
**Values** | Pointer to [**LogInsightsValues**](LogInsightsValues.md) |  | [optional] 

## Methods

### NewLogInsights

`func NewLogInsights() *LogInsights`

NewLogInsights instantiates a new LogInsights object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogInsightsWithDefaults

`func NewLogInsightsWithDefaults() *LogInsights`

NewLogInsightsWithDefaults instantiates a new LogInsights object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *LogInsights) GetDimensions() LogInsightsDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *LogInsights) GetDimensionsOk() (*LogInsightsDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *LogInsights) SetDimensions(v LogInsightsDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *LogInsights) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetDimensionAttributes

`func (o *LogInsights) GetDimensionAttributes() LogInsightsDimensionAttributes`

GetDimensionAttributes returns the DimensionAttributes field if non-nil, zero value otherwise.

### GetDimensionAttributesOk

`func (o *LogInsights) GetDimensionAttributesOk() (*LogInsightsDimensionAttributes, bool)`

GetDimensionAttributesOk returns a tuple with the DimensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionAttributes

`func (o *LogInsights) SetDimensionAttributes(v LogInsightsDimensionAttributes)`

SetDimensionAttributes sets DimensionAttributes field to given value.

### HasDimensionAttributes

`func (o *LogInsights) HasDimensionAttributes() bool`

HasDimensionAttributes returns a boolean if a field has been set.

### GetValues

`func (o *LogInsights) GetValues() LogInsightsValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LogInsights) GetValuesOk() (*LogInsightsValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LogInsights) SetValues(v LogInsightsValues)`

SetValues sets Values field to given value.

### HasValues

`func (o *LogInsights) HasValues() bool`

HasValues returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
