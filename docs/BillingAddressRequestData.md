# BillingAddressRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeBillingAddress**](TypeBillingAddress.md) |  | [optional] [default to TYPEBILLINGADDRESS_BILLING_ADDRESS]
**Attributes** | Pointer to [**BillingAddressAttributes**](BillingAddressAttributes.md) |  | [optional] 

## Methods

### NewBillingAddressRequestData

`func NewBillingAddressRequestData() *BillingAddressRequestData`

NewBillingAddressRequestData instantiates a new BillingAddressRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAddressRequestDataWithDefaults

`func NewBillingAddressRequestDataWithDefaults() *BillingAddressRequestData`

NewBillingAddressRequestDataWithDefaults instantiates a new BillingAddressRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BillingAddressRequestData) GetType() TypeBillingAddress`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingAddressRequestData) GetTypeOk() (*TypeBillingAddress, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingAddressRequestData) SetType(v TypeBillingAddress)`

SetType sets Type field to given value.

### HasType

`func (o *BillingAddressRequestData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *BillingAddressRequestData) GetAttributes() BillingAddressAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BillingAddressRequestData) GetAttributesOk() (*BillingAddressAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BillingAddressRequestData) SetAttributes(v BillingAddressAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BillingAddressRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
