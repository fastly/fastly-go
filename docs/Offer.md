# Offer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | Pointer to **string** | The aftermarket vendor. | [optional] 
**Price** | Pointer to **string** | The price for the domain from the aftermarket vendor. | [optional] 
**Currency** | Pointer to **string** | The currency for the aftermarket offer. | [optional] 

## Methods

### NewOffer

`func NewOffer() *Offer`

NewOffer instantiates a new Offer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferWithDefaults

`func NewOfferWithDefaults() *Offer`

NewOfferWithDefaults instantiates a new Offer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *Offer) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Offer) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Offer) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Offer) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetPrice

`func (o *Offer) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Offer) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Offer) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Offer) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *Offer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Offer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Offer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Offer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


