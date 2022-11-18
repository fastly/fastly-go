# HistoricalUsageMonthResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerID** | Pointer to **string** |  | [optional] [readonly] 
**Services** | Pointer to [**map[string]map[string]HistoricalUsageResults**](map.md) |  | [optional] 
**Total** | Pointer to [**HistoricalUsageResults**](HistoricalUsageResults.md) |  | [optional] 

## Methods

### NewHistoricalUsageMonthResponseAllOfData

`func NewHistoricalUsageMonthResponseAllOfData() *HistoricalUsageMonthResponseAllOfData`

NewHistoricalUsageMonthResponseAllOfData instantiates a new HistoricalUsageMonthResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalUsageMonthResponseAllOfDataWithDefaults

`func NewHistoricalUsageMonthResponseAllOfDataWithDefaults() *HistoricalUsageMonthResponseAllOfData`

NewHistoricalUsageMonthResponseAllOfDataWithDefaults instantiates a new HistoricalUsageMonthResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerID

`func (o *HistoricalUsageMonthResponseAllOfData) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *HistoricalUsageMonthResponseAllOfData) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *HistoricalUsageMonthResponseAllOfData) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *HistoricalUsageMonthResponseAllOfData) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetServices

`func (o *HistoricalUsageMonthResponseAllOfData) GetServices() map[string]map[string]HistoricalUsageResults`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *HistoricalUsageMonthResponseAllOfData) GetServicesOk() (*map[string]map[string]HistoricalUsageResults, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *HistoricalUsageMonthResponseAllOfData) SetServices(v map[string]map[string]HistoricalUsageResults)`

SetServices sets Services field to given value.

### HasServices

`func (o *HistoricalUsageMonthResponseAllOfData) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTotal

`func (o *HistoricalUsageMonthResponseAllOfData) GetTotal() HistoricalUsageResults`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HistoricalUsageMonthResponseAllOfData) GetTotalOk() (*HistoricalUsageResults, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HistoricalUsageMonthResponseAllOfData) SetTotal(v HistoricalUsageResults)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HistoricalUsageMonthResponseAllOfData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
