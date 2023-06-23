# HistoricalUsageMonthResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerID** | Pointer to **string** |  | [optional] [readonly] 
**Services** | Pointer to [**map[string]HistoricalService**](HistoricalService.md) |  | [optional] 
**Total** | Pointer to [**HistoricalUsageResults**](HistoricalUsageResults.md) |  | [optional] 

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

### GetCustomerID

`func (o *HistoricalUsageMonthResponseData) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *HistoricalUsageMonthResponseData) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *HistoricalUsageMonthResponseData) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *HistoricalUsageMonthResponseData) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetServices

`func (o *HistoricalUsageMonthResponseData) GetServices() map[string]HistoricalService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *HistoricalUsageMonthResponseData) GetServicesOk() (*map[string]HistoricalService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *HistoricalUsageMonthResponseData) SetServices(v map[string]HistoricalService)`

SetServices sets Services field to given value.

### HasServices

`func (o *HistoricalUsageMonthResponseData) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTotal

`func (o *HistoricalUsageMonthResponseData) GetTotal() HistoricalUsageResults`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HistoricalUsageMonthResponseData) GetTotalOk() (*HistoricalUsageResults, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HistoricalUsageMonthResponseData) SetTotal(v HistoricalUsageResults)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HistoricalUsageMonthResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
