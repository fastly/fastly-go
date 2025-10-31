# DashboardItemPropertyDataSourcePropertyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | **[]string** | The metrics to visualize. Valid options are defined by the selected [data source](#field_data_source). | 

## Methods

### NewDashboardItemPropertyDataSourcePropertyConfig

`func NewDashboardItemPropertyDataSourcePropertyConfig(metrics []string, ) *DashboardItemPropertyDataSourcePropertyConfig`

NewDashboardItemPropertyDataSourcePropertyConfig instantiates a new DashboardItemPropertyDataSourcePropertyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardItemPropertyDataSourcePropertyConfigWithDefaults

`func NewDashboardItemPropertyDataSourcePropertyConfigWithDefaults() *DashboardItemPropertyDataSourcePropertyConfig`

NewDashboardItemPropertyDataSourcePropertyConfigWithDefaults instantiates a new DashboardItemPropertyDataSourcePropertyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *DashboardItemPropertyDataSourcePropertyConfig) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DashboardItemPropertyDataSourcePropertyConfig) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DashboardItemPropertyDataSourcePropertyConfig) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


