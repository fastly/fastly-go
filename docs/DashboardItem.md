# DashboardItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Dashboard item identifier (UUID) | [optional] [readonly] 
**Title** | **string** | A human-readable title for the dashboard item | 
**Subtitle** | **string** | A human-readable subtitle for the dashboard item. Often a description of the visualization. | 
**DataSource** | [**DashboardItemPropertyDataSource**](DashboardItemPropertyDataSource.md) |  | 
**Visualization** | [**DashboardItemPropertyVisualization**](DashboardItemPropertyVisualization.md) |  | 
**Span** | Pointer to **int32** | The number of columns for the dashboard item to span. Dashboards are rendered on a 12-column grid on \&quot;desktop\&quot; screen sizes. | [optional] [default to 4]

## Methods

### NewDashboardItem

`func NewDashboardItem(title string, subtitle string, dataSource DashboardItemPropertyDataSource, visualization DashboardItemPropertyVisualization, ) *DashboardItem`

NewDashboardItem instantiates a new DashboardItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardItemWithDefaults

`func NewDashboardItemWithDefaults() *DashboardItem`

NewDashboardItemWithDefaults instantiates a new DashboardItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DashboardItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *DashboardItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubtitle

`func (o *DashboardItem) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *DashboardItem) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *DashboardItem) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.


### GetDataSource

`func (o *DashboardItem) GetDataSource() DashboardItemPropertyDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DashboardItem) GetDataSourceOk() (*DashboardItemPropertyDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DashboardItem) SetDataSource(v DashboardItemPropertyDataSource)`

SetDataSource sets DataSource field to given value.


### GetVisualization

`func (o *DashboardItem) GetVisualization() DashboardItemPropertyVisualization`

GetVisualization returns the Visualization field if non-nil, zero value otherwise.

### GetVisualizationOk

`func (o *DashboardItem) GetVisualizationOk() (*DashboardItemPropertyVisualization, bool)`

GetVisualizationOk returns a tuple with the Visualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualization

`func (o *DashboardItem) SetVisualization(v DashboardItemPropertyVisualization)`

SetVisualization sets Visualization field to given value.


### GetSpan

`func (o *DashboardItem) GetSpan() int32`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *DashboardItem) GetSpanOk() (*int32, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *DashboardItem) SetSpan(v int32)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *DashboardItem) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


