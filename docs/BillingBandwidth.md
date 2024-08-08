# BillingBandwidth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** |  | [optional] 
**Tiers** | Pointer to [**[]BillingBandwidthTiers**](BillingBandwidthTiers.md) |  | [optional] 

## Methods

### NewBillingBandwidth

`func NewBillingBandwidth() *BillingBandwidth`

NewBillingBandwidth instantiates a new BillingBandwidth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingBandwidthWithDefaults

`func NewBillingBandwidthWithDefaults() *BillingBandwidth`

NewBillingBandwidthWithDefaults instantiates a new BillingBandwidth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *BillingBandwidth) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BillingBandwidth) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BillingBandwidth) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BillingBandwidth) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTiers

`func (o *BillingBandwidth) GetTiers() []BillingBandwidthTiers`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *BillingBandwidth) GetTiersOk() (*[]BillingBandwidthTiers, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *BillingBandwidth) SetTiers(v []BillingBandwidthTiers)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *BillingBandwidth) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
