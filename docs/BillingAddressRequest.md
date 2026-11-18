# BillingAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipVerification** | Pointer to **bool** | When set to true, the address will be saved without verification | [optional] 
**Data** | Pointer to [**BillingAddressRequestData**](BillingAddressRequestData.md) |  | [optional] 

## Methods

### NewBillingAddressRequest

`func NewBillingAddressRequest() *BillingAddressRequest`

NewBillingAddressRequest instantiates a new BillingAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAddressRequestWithDefaults

`func NewBillingAddressRequestWithDefaults() *BillingAddressRequest`

NewBillingAddressRequestWithDefaults instantiates a new BillingAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipVerification

`func (o *BillingAddressRequest) GetSkipVerification() bool`

GetSkipVerification returns the SkipVerification field if non-nil, zero value otherwise.

### GetSkipVerificationOk

`func (o *BillingAddressRequest) GetSkipVerificationOk() (*bool, bool)`

GetSkipVerificationOk returns a tuple with the SkipVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipVerification

`func (o *BillingAddressRequest) SetSkipVerification(v bool)`

SetSkipVerification sets SkipVerification field to given value.

### HasSkipVerification

`func (o *BillingAddressRequest) HasSkipVerification() bool`

HasSkipVerification returns a boolean if a field has been set.

### GetData

`func (o *BillingAddressRequest) GetData() BillingAddressRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillingAddressRequest) GetDataOk() (*BillingAddressRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillingAddressRequest) SetData(v BillingAddressRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *BillingAddressRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
