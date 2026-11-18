# Invoicelineitems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Invoice line item transaction name. | [optional] 
**Amount** | Pointer to **float32** | Billed amount for line item. | [optional] 
**CreditCouponCode** | Pointer to **string** | Discount coupon associated with the invoice for any account or service credits. | [optional] 
**Rate** | Pointer to **float32** | Price per unit. | [optional] 
**Units** | Pointer to **float32** | Total number of units of usage. | [optional] 
**ProductName** | Pointer to **string** | The name of the product. | [optional] 
**ProductGroup** | Pointer to **string** | The broader classification of the product (e.g., `Compute` or `Full-Site Delivery`). | [optional] 
**ProductLine** | Pointer to **string** | The broader classification of the product (e.g., `Network Services` or `Security`). | [optional] 
**Region** | Pointer to **string** | The geographical area applicable for regionally based products. | [optional] 
**UsageType** | Pointer to **string** | The unit of measure (e.g., `requests` or `bandwidth`). | [optional] 

## Methods

### NewInvoicelineitems

`func NewInvoicelineitems() *Invoicelineitems`

NewInvoicelineitems instantiates a new Invoicelineitems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicelineitemsWithDefaults

`func NewInvoicelineitemsWithDefaults() *Invoicelineitems`

NewInvoicelineitemsWithDefaults instantiates a new Invoicelineitems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Invoicelineitems) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Invoicelineitems) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Invoicelineitems) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Invoicelineitems) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *Invoicelineitems) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Invoicelineitems) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Invoicelineitems) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Invoicelineitems) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreditCouponCode

`func (o *Invoicelineitems) GetCreditCouponCode() string`

GetCreditCouponCode returns the CreditCouponCode field if non-nil, zero value otherwise.

### GetCreditCouponCodeOk

`func (o *Invoicelineitems) GetCreditCouponCodeOk() (*string, bool)`

GetCreditCouponCodeOk returns a tuple with the CreditCouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCouponCode

`func (o *Invoicelineitems) SetCreditCouponCode(v string)`

SetCreditCouponCode sets CreditCouponCode field to given value.

### HasCreditCouponCode

`func (o *Invoicelineitems) HasCreditCouponCode() bool`

HasCreditCouponCode returns a boolean if a field has been set.

### GetRate

`func (o *Invoicelineitems) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *Invoicelineitems) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *Invoicelineitems) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *Invoicelineitems) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetUnits

`func (o *Invoicelineitems) GetUnits() float32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Invoicelineitems) GetUnitsOk() (*float32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Invoicelineitems) SetUnits(v float32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Invoicelineitems) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetProductName

`func (o *Invoicelineitems) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *Invoicelineitems) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *Invoicelineitems) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *Invoicelineitems) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductGroup

`func (o *Invoicelineitems) GetProductGroup() string`

GetProductGroup returns the ProductGroup field if non-nil, zero value otherwise.

### GetProductGroupOk

`func (o *Invoicelineitems) GetProductGroupOk() (*string, bool)`

GetProductGroupOk returns a tuple with the ProductGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroup

`func (o *Invoicelineitems) SetProductGroup(v string)`

SetProductGroup sets ProductGroup field to given value.

### HasProductGroup

`func (o *Invoicelineitems) HasProductGroup() bool`

HasProductGroup returns a boolean if a field has been set.

### GetProductLine

`func (o *Invoicelineitems) GetProductLine() string`

GetProductLine returns the ProductLine field if non-nil, zero value otherwise.

### GetProductLineOk

`func (o *Invoicelineitems) GetProductLineOk() (*string, bool)`

GetProductLineOk returns a tuple with the ProductLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLine

`func (o *Invoicelineitems) SetProductLine(v string)`

SetProductLine sets ProductLine field to given value.

### HasProductLine

`func (o *Invoicelineitems) HasProductLine() bool`

HasProductLine returns a boolean if a field has been set.

### GetRegion

`func (o *Invoicelineitems) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Invoicelineitems) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Invoicelineitems) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Invoicelineitems) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUsageType

`func (o *Invoicelineitems) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *Invoicelineitems) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *Invoicelineitems) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *Invoicelineitems) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
