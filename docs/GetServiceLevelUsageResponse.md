# GetServiceLevelUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceusagemetricsData**](ServiceusagemetricsData.md) |  | [optional] 

## Methods

### NewGetServiceLevelUsageResponse

`func NewGetServiceLevelUsageResponse() *GetServiceLevelUsageResponse`

NewGetServiceLevelUsageResponse instantiates a new GetServiceLevelUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceLevelUsageResponseWithDefaults

`func NewGetServiceLevelUsageResponseWithDefaults() *GetServiceLevelUsageResponse`

NewGetServiceLevelUsageResponseWithDefaults instantiates a new GetServiceLevelUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetServiceLevelUsageResponse) GetData() ServiceusagemetricsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetServiceLevelUsageResponse) GetDataOk() (*ServiceusagemetricsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetServiceLevelUsageResponse) SetData(v ServiceusagemetricsData)`

SetData sets Data field to given value.

### HasData

`func (o *GetServiceLevelUsageResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


