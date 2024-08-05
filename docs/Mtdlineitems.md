# Mtdlineitems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Invoice line item transaction name. | [optional] 
**Amount** | Pointer to **float32** | Billed amount for line item. | [optional] 
**Rate** | Pointer to **float32** | Price per unit. | [optional] 
**Units** | Pointer to **float32** | Total number of units of usage. | [optional] 
**ProductName** | Pointer to **string** | The name of the product. | [optional] 
**ProductGroup** | Pointer to **string** | The broader classification of the product (e.g., `Compute` or `Full-Site Delivery`). | [optional] 
**ProductLine** | Pointer to **string** | The broader classification of the product (e.g., `Network Services` or `Security`). | [optional] 
**Region** | Pointer to **string** | The geographical area applicable for regionally based products. | [optional] 
**UsageType** | Pointer to **string** | The unit of measure (e.g., `requests` or `bandwidth`). | [optional] 

## Methods

### NewMtdlineitems

`func NewMtdlineitems() *Mtdlineitems`

NewMtdlineitems instantiates a new Mtdlineitems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMtdlineitemsWithDefaults

`func NewMtdlineitemsWithDefaults() *Mtdlineitems`

NewMtdlineitemsWithDefaults instantiates a new Mtdlineitems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Mtdlineitems) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Mtdlineitems) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Mtdlineitems) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Mtdlineitems) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *Mtdlineitems) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Mtdlineitems) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Mtdlineitems) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Mtdlineitems) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetRate

`func (o *Mtdlineitems) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *Mtdlineitems) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *Mtdlineitems) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *Mtdlineitems) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetUnits

`func (o *Mtdlineitems) GetUnits() float32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Mtdlineitems) GetUnitsOk() (*float32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Mtdlineitems) SetUnits(v float32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Mtdlineitems) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetProductName

`func (o *Mtdlineitems) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *Mtdlineitems) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *Mtdlineitems) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *Mtdlineitems) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductGroup

`func (o *Mtdlineitems) GetProductGroup() string`

GetProductGroup returns the ProductGroup field if non-nil, zero value otherwise.

### GetProductGroupOk

`func (o *Mtdlineitems) GetProductGroupOk() (*string, bool)`

GetProductGroupOk returns a tuple with the ProductGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroup

`func (o *Mtdlineitems) SetProductGroup(v string)`

SetProductGroup sets ProductGroup field to given value.

### HasProductGroup

`func (o *Mtdlineitems) HasProductGroup() bool`

HasProductGroup returns a boolean if a field has been set.

### GetProductLine

`func (o *Mtdlineitems) GetProductLine() string`

GetProductLine returns the ProductLine field if non-nil, zero value otherwise.

### GetProductLineOk

`func (o *Mtdlineitems) GetProductLineOk() (*string, bool)`

GetProductLineOk returns a tuple with the ProductLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLine

`func (o *Mtdlineitems) SetProductLine(v string)`

SetProductLine sets ProductLine field to given value.

### HasProductLine

`func (o *Mtdlineitems) HasProductLine() bool`

HasProductLine returns a boolean if a field has been set.

### GetRegion

`func (o *Mtdlineitems) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Mtdlineitems) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Mtdlineitems) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Mtdlineitems) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUsageType

`func (o *Mtdlineitems) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *Mtdlineitems) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *Mtdlineitems) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *Mtdlineitems) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
