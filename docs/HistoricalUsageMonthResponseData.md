# HistoricalUsageMonthResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** |  | [optional] [readonly] 
**Services** | Pointer to [**map[string]HistoricalUsageService**](HistoricalUsageService.md) | Organized by *service id*. | [optional] 
**Total** | Pointer to [**map[string]HistoricalUsageData**](HistoricalUsageData.md) | Organized by *region*. | [optional] 

## Methods

### NewHistoricalUsageMonthResponseData

`func NewHistoricalUsageMonthResponseData() *HistoricalUsageMonthResponseData`

NewHistoricalUsageMonthResponseData instantiates a new HistoricalUsageMonthResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalUsageMonthResponseDataWithDefaults

`func NewHistoricalUsageMonthResponseDataWithDefaults() *HistoricalUsageMonthResponseData`

NewHistoricalUsageMonthResponseDataWithDefaults instantiates a new HistoricalUsageMonthResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *HistoricalUsageMonthResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *HistoricalUsageMonthResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *HistoricalUsageMonthResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *HistoricalUsageMonthResponseData) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetServices

`func (o *HistoricalUsageMonthResponseData) GetServices() map[string]HistoricalUsageService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *HistoricalUsageMonthResponseData) GetServicesOk() (*map[string]HistoricalUsageService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *HistoricalUsageMonthResponseData) SetServices(v map[string]HistoricalUsageService)`

SetServices sets Services field to given value.

### HasServices

`func (o *HistoricalUsageMonthResponseData) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTotal

`func (o *HistoricalUsageMonthResponseData) GetTotal() map[string]HistoricalUsageData`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HistoricalUsageMonthResponseData) GetTotalOk() (*map[string]HistoricalUsageData, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HistoricalUsageMonthResponseData) SetTotal(v map[string]HistoricalUsageData)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HistoricalUsageMonthResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


