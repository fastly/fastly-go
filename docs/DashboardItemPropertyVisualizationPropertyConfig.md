# DashboardItemPropertyVisualizationPropertyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlotType** | **string** | The type of chart to display.  | 
**Format** | Pointer to **string** | (Optional) The units to use to format the data.  | [optional] [default to "number"]
**CalculationMethod** | Pointer to **string** | (Optional) The aggregation function to apply to the dataset.  | [optional] 

## Methods

### NewDashboardItemPropertyVisualizationPropertyConfig

`func NewDashboardItemPropertyVisualizationPropertyConfig(plotType string, ) *DashboardItemPropertyVisualizationPropertyConfig`

NewDashboardItemPropertyVisualizationPropertyConfig instantiates a new DashboardItemPropertyVisualizationPropertyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardItemPropertyVisualizationPropertyConfigWithDefaults

`func NewDashboardItemPropertyVisualizationPropertyConfigWithDefaults() *DashboardItemPropertyVisualizationPropertyConfig`

NewDashboardItemPropertyVisualizationPropertyConfigWithDefaults instantiates a new DashboardItemPropertyVisualizationPropertyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlotType

`func (o *DashboardItemPropertyVisualizationPropertyConfig) GetPlotType() string`

GetPlotType returns the PlotType field if non-nil, zero value otherwise.

### GetPlotTypeOk

`func (o *DashboardItemPropertyVisualizationPropertyConfig) GetPlotTypeOk() (*string, bool)`

GetPlotTypeOk returns a tuple with the PlotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotType

`func (o *DashboardItemPropertyVisualizationPropertyConfig) SetPlotType(v string)`

SetPlotType sets PlotType field to given value.


### GetFormat

`func (o *DashboardItemPropertyVisualizationPropertyConfig) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DashboardItemPropertyVisualizationPropertyConfig) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DashboardItemPropertyVisualizationPropertyConfig) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DashboardItemPropertyVisualizationPropertyConfig) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetCalculationMethod

`func (o *DashboardItemPropertyVisualizationPropertyConfig) GetCalculationMethod() string`

GetCalculationMethod returns the CalculationMethod field if non-nil, zero value otherwise.

### GetCalculationMethodOk

`func (o *DashboardItemPropertyVisualizationPropertyConfig) GetCalculationMethodOk() (*string, bool)`

GetCalculationMethodOk returns a tuple with the CalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationMethod

`func (o *DashboardItemPropertyVisualizationPropertyConfig) SetCalculationMethod(v string)`

SetCalculationMethod sets CalculationMethod field to given value.

### HasCalculationMethod

`func (o *DashboardItemPropertyVisualizationPropertyConfig) HasCalculationMethod() bool`

HasCalculationMethod returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


