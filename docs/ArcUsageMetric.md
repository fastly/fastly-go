# ArcUsageMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | The date (or hour) the usage was measured. Formatted as `YYYY-MM-DD` when `granularity` is `daily`, or `YYYY-MM-DD HH:00:00` when `granularity` is `hourly`.  | 
**UsageType** | **string** | The type of usage being measured. | 
**Quantity** | **int32** | The quantity of the usage type for this date. | 
**VirtualKeyId** | Pointer to **string** | The ID of the virtual key this usage is attributed to. | [optional] 
**VirtualKeyName** | Pointer to **string** | The display name of the virtual key this usage is attributed to. | [optional] 
**PrimaryProvider** | Pointer to **string** | The ID of the provider connection originally requested, before any failover occurred. | [optional] 
**Provider** | Pointer to **string** | The name of the provider connection that ultimately served the request. | [optional] 
**ProviderConnectionId** | Pointer to **string** | The ID of the provider connection that ultimately served the request. | [optional] 
**Model** | Pointer to **string** | The AI model identifier used. | [optional] 

## Methods

### NewArcUsageMetric

`func NewArcUsageMetric(date string, usageType string, quantity int32, ) *ArcUsageMetric`

NewArcUsageMetric instantiates a new ArcUsageMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcUsageMetricWithDefaults

`func NewArcUsageMetricWithDefaults() *ArcUsageMetric`

NewArcUsageMetricWithDefaults instantiates a new ArcUsageMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ArcUsageMetric) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ArcUsageMetric) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ArcUsageMetric) SetDate(v string)`

SetDate sets Date field to given value.


### GetUsageType

`func (o *ArcUsageMetric) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *ArcUsageMetric) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *ArcUsageMetric) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.


### GetQuantity

`func (o *ArcUsageMetric) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ArcUsageMetric) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ArcUsageMetric) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetVirtualKeyId

`func (o *ArcUsageMetric) GetVirtualKeyId() string`

GetVirtualKeyId returns the VirtualKeyId field if non-nil, zero value otherwise.

### GetVirtualKeyIdOk

`func (o *ArcUsageMetric) GetVirtualKeyIdOk() (*string, bool)`

GetVirtualKeyIdOk returns a tuple with the VirtualKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualKeyId

`func (o *ArcUsageMetric) SetVirtualKeyId(v string)`

SetVirtualKeyId sets VirtualKeyId field to given value.

### HasVirtualKeyId

`func (o *ArcUsageMetric) HasVirtualKeyId() bool`

HasVirtualKeyId returns a boolean if a field has been set.

### GetVirtualKeyName

`func (o *ArcUsageMetric) GetVirtualKeyName() string`

GetVirtualKeyName returns the VirtualKeyName field if non-nil, zero value otherwise.

### GetVirtualKeyNameOk

`func (o *ArcUsageMetric) GetVirtualKeyNameOk() (*string, bool)`

GetVirtualKeyNameOk returns a tuple with the VirtualKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualKeyName

`func (o *ArcUsageMetric) SetVirtualKeyName(v string)`

SetVirtualKeyName sets VirtualKeyName field to given value.

### HasVirtualKeyName

`func (o *ArcUsageMetric) HasVirtualKeyName() bool`

HasVirtualKeyName returns a boolean if a field has been set.

### GetPrimaryProvider

`func (o *ArcUsageMetric) GetPrimaryProvider() string`

GetPrimaryProvider returns the PrimaryProvider field if non-nil, zero value otherwise.

### GetPrimaryProviderOk

`func (o *ArcUsageMetric) GetPrimaryProviderOk() (*string, bool)`

GetPrimaryProviderOk returns a tuple with the PrimaryProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryProvider

`func (o *ArcUsageMetric) SetPrimaryProvider(v string)`

SetPrimaryProvider sets PrimaryProvider field to given value.

### HasPrimaryProvider

`func (o *ArcUsageMetric) HasPrimaryProvider() bool`

HasPrimaryProvider returns a boolean if a field has been set.

### GetProvider

`func (o *ArcUsageMetric) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ArcUsageMetric) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ArcUsageMetric) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ArcUsageMetric) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderConnectionId

`func (o *ArcUsageMetric) GetProviderConnectionId() string`

GetProviderConnectionId returns the ProviderConnectionId field if non-nil, zero value otherwise.

### GetProviderConnectionIdOk

`func (o *ArcUsageMetric) GetProviderConnectionIdOk() (*string, bool)`

GetProviderConnectionIdOk returns a tuple with the ProviderConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConnectionId

`func (o *ArcUsageMetric) SetProviderConnectionId(v string)`

SetProviderConnectionId sets ProviderConnectionId field to given value.

### HasProviderConnectionId

`func (o *ArcUsageMetric) HasProviderConnectionId() bool`

HasProviderConnectionId returns a boolean if a field has been set.

### GetModel

`func (o *ArcUsageMetric) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ArcUsageMetric) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ArcUsageMetric) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ArcUsageMetric) HasModel() bool`

HasModel returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


