# BillingRegions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **float32** |  | [optional] 
**Bandwidth** | Pointer to [**BillingBandwidth**](BillingBandwidth.md) |  | [optional] 
**Percentile** | Pointer to [**BillingBandwidth**](BillingBandwidth.md) |  | [optional] 
**Requests** | Pointer to [**BillingBandwidth**](BillingBandwidth.md) |  | [optional] 

## Methods

### NewBillingRegions

`func NewBillingRegions() *BillingRegions`

NewBillingRegions instantiates a new BillingRegions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingRegionsWithDefaults

`func NewBillingRegionsWithDefaults() *BillingRegions`

NewBillingRegionsWithDefaults instantiates a new BillingRegions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *BillingRegions) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingRegions) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingRegions) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingRegions) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetBandwidth

`func (o *BillingRegions) GetBandwidth() BillingBandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *BillingRegions) GetBandwidthOk() (*BillingBandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *BillingRegions) SetBandwidth(v BillingBandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *BillingRegions) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetPercentile

`func (o *BillingRegions) GetPercentile() BillingBandwidth`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *BillingRegions) GetPercentileOk() (*BillingBandwidth, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *BillingRegions) SetPercentile(v BillingBandwidth)`

SetPercentile sets Percentile field to given value.

### HasPercentile

`func (o *BillingRegions) HasPercentile() bool`

HasPercentile returns a boolean if a field has been set.

### GetRequests

`func (o *BillingRegions) GetRequests() BillingBandwidth`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *BillingRegions) GetRequestsOk() (*BillingBandwidth, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *BillingRegions) SetRequests(v BillingBandwidth)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *BillingRegions) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
