# DashboardItemPropertyDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The source of the data to display. | 
**Config** | [**DashboardItemPropertyDataSourcePropertyConfig**](DashboardItemPropertyDataSourcePropertyConfig.md) |  | 

## Methods

### NewDashboardItemPropertyDataSource

`func NewDashboardItemPropertyDataSource(type_ string, config DashboardItemPropertyDataSourcePropertyConfig, ) *DashboardItemPropertyDataSource`

NewDashboardItemPropertyDataSource instantiates a new DashboardItemPropertyDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardItemPropertyDataSourceWithDefaults

`func NewDashboardItemPropertyDataSourceWithDefaults() *DashboardItemPropertyDataSource`

NewDashboardItemPropertyDataSourceWithDefaults instantiates a new DashboardItemPropertyDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DashboardItemPropertyDataSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardItemPropertyDataSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardItemPropertyDataSource) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *DashboardItemPropertyDataSource) GetConfig() DashboardItemPropertyDataSourcePropertyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DashboardItemPropertyDataSource) GetConfigOk() (*DashboardItemPropertyDataSourcePropertyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DashboardItemPropertyDataSource) SetConfig(v DashboardItemPropertyDataSourcePropertyConfig)`

SetConfig sets Config field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


