# LineItemData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**AriaInvoiceID** | Pointer to **string** | An alphanumeric string identifying the invoice. | [optional] [readonly] 
**ClientServiceID** | Pointer to **string** |  | [optional] 
**CreditCouponCode** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**LineNumber** | Pointer to **int32** |  | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 
**PlanNo** | Pointer to **float32** |  | [optional] 
**RatePerUnit** | Pointer to **NullableFloat32** |  | [optional] 
**RateScheduleNo** | Pointer to **NullableFloat32** |  | [optional] 
**RateScheduleTierNo** | Pointer to **NullableFloat32** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**ServiceNo** | Pointer to **float32** |  | [optional] 
**Units** | Pointer to **float32** |  | [optional] 
**UsageTypeCd** | Pointer to **NullableString** |  | [optional] 
**UsageTypeNo** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewLineItemData

`func NewLineItemData() *LineItemData`

NewLineItemData instantiates a new LineItemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemDataWithDefaults

`func NewLineItemDataWithDefaults() *LineItemData`

NewLineItemDataWithDefaults instantiates a new LineItemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *LineItemData) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LineItemData) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LineItemData) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *LineItemData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAriaInvoiceID

`func (o *LineItemData) GetAriaInvoiceID() string`

GetAriaInvoiceID returns the AriaInvoiceID field if non-nil, zero value otherwise.

### GetAriaInvoiceIDOk

`func (o *LineItemData) GetAriaInvoiceIDOk() (*string, bool)`

GetAriaInvoiceIDOk returns a tuple with the AriaInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAriaInvoiceID

`func (o *LineItemData) SetAriaInvoiceID(v string)`

SetAriaInvoiceID sets AriaInvoiceID field to given value.

### HasAriaInvoiceID

`func (o *LineItemData) HasAriaInvoiceID() bool`

HasAriaInvoiceID returns a boolean if a field has been set.

### GetClientServiceID

`func (o *LineItemData) GetClientServiceID() string`

GetClientServiceID returns the ClientServiceID field if non-nil, zero value otherwise.

### GetClientServiceIDOk

`func (o *LineItemData) GetClientServiceIDOk() (*string, bool)`

GetClientServiceIDOk returns a tuple with the ClientServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientServiceID

`func (o *LineItemData) SetClientServiceID(v string)`

SetClientServiceID sets ClientServiceID field to given value.

### HasClientServiceID

`func (o *LineItemData) HasClientServiceID() bool`

HasClientServiceID returns a boolean if a field has been set.

### GetCreditCouponCode

`func (o *LineItemData) GetCreditCouponCode() string`

GetCreditCouponCode returns the CreditCouponCode field if non-nil, zero value otherwise.

### GetCreditCouponCodeOk

`func (o *LineItemData) GetCreditCouponCodeOk() (*string, bool)`

GetCreditCouponCodeOk returns a tuple with the CreditCouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCouponCode

`func (o *LineItemData) SetCreditCouponCode(v string)`

SetCreditCouponCode sets CreditCouponCode field to given value.

### HasCreditCouponCode

`func (o *LineItemData) HasCreditCouponCode() bool`

HasCreditCouponCode returns a boolean if a field has been set.

### SetCreditCouponCodeNil

`func (o *LineItemData) SetCreditCouponCodeNil(b bool)`

 SetCreditCouponCodeNil sets the value for CreditCouponCode to be an explicit nil

### UnsetCreditCouponCode
`func (o *LineItemData) UnsetCreditCouponCode()`

UnsetCreditCouponCode ensures that no value is present for CreditCouponCode, not even an explicit nil
### GetDescription

`func (o *LineItemData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LineItemData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LineItemData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LineItemData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetID

`func (o *LineItemData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *LineItemData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *LineItemData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *LineItemData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetLineNumber

`func (o *LineItemData) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *LineItemData) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *LineItemData) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *LineItemData) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetPlanName

`func (o *LineItemData) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *LineItemData) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *LineItemData) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *LineItemData) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanNo

`func (o *LineItemData) GetPlanNo() float32`

GetPlanNo returns the PlanNo field if non-nil, zero value otherwise.

### GetPlanNoOk

`func (o *LineItemData) GetPlanNoOk() (*float32, bool)`

GetPlanNoOk returns a tuple with the PlanNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanNo

`func (o *LineItemData) SetPlanNo(v float32)`

SetPlanNo sets PlanNo field to given value.

### HasPlanNo

`func (o *LineItemData) HasPlanNo() bool`

HasPlanNo returns a boolean if a field has been set.

### GetRatePerUnit

`func (o *LineItemData) GetRatePerUnit() float32`

GetRatePerUnit returns the RatePerUnit field if non-nil, zero value otherwise.

### GetRatePerUnitOk

`func (o *LineItemData) GetRatePerUnitOk() (*float32, bool)`

GetRatePerUnitOk returns a tuple with the RatePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerUnit

`func (o *LineItemData) SetRatePerUnit(v float32)`

SetRatePerUnit sets RatePerUnit field to given value.

### HasRatePerUnit

`func (o *LineItemData) HasRatePerUnit() bool`

HasRatePerUnit returns a boolean if a field has been set.

### SetRatePerUnitNil

`func (o *LineItemData) SetRatePerUnitNil(b bool)`

 SetRatePerUnitNil sets the value for RatePerUnit to be an explicit nil

### UnsetRatePerUnit
`func (o *LineItemData) UnsetRatePerUnit()`

UnsetRatePerUnit ensures that no value is present for RatePerUnit, not even an explicit nil
### GetRateScheduleNo

`func (o *LineItemData) GetRateScheduleNo() float32`

GetRateScheduleNo returns the RateScheduleNo field if non-nil, zero value otherwise.

### GetRateScheduleNoOk

`func (o *LineItemData) GetRateScheduleNoOk() (*float32, bool)`

GetRateScheduleNoOk returns a tuple with the RateScheduleNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateScheduleNo

`func (o *LineItemData) SetRateScheduleNo(v float32)`

SetRateScheduleNo sets RateScheduleNo field to given value.

### HasRateScheduleNo

`func (o *LineItemData) HasRateScheduleNo() bool`

HasRateScheduleNo returns a boolean if a field has been set.

### SetRateScheduleNoNil

`func (o *LineItemData) SetRateScheduleNoNil(b bool)`

 SetRateScheduleNoNil sets the value for RateScheduleNo to be an explicit nil

### UnsetRateScheduleNo
`func (o *LineItemData) UnsetRateScheduleNo()`

UnsetRateScheduleNo ensures that no value is present for RateScheduleNo, not even an explicit nil
### GetRateScheduleTierNo

`func (o *LineItemData) GetRateScheduleTierNo() float32`

GetRateScheduleTierNo returns the RateScheduleTierNo field if non-nil, zero value otherwise.

### GetRateScheduleTierNoOk

`func (o *LineItemData) GetRateScheduleTierNoOk() (*float32, bool)`

GetRateScheduleTierNoOk returns a tuple with the RateScheduleTierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateScheduleTierNo

`func (o *LineItemData) SetRateScheduleTierNo(v float32)`

SetRateScheduleTierNo sets RateScheduleTierNo field to given value.

### HasRateScheduleTierNo

`func (o *LineItemData) HasRateScheduleTierNo() bool`

HasRateScheduleTierNo returns a boolean if a field has been set.

### SetRateScheduleTierNoNil

`func (o *LineItemData) SetRateScheduleTierNoNil(b bool)`

 SetRateScheduleTierNoNil sets the value for RateScheduleTierNo to be an explicit nil

### UnsetRateScheduleTierNo
`func (o *LineItemData) UnsetRateScheduleTierNo()`

UnsetRateScheduleTierNo ensures that no value is present for RateScheduleTierNo, not even an explicit nil
### GetServiceName

`func (o *LineItemData) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *LineItemData) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *LineItemData) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *LineItemData) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceNo

`func (o *LineItemData) GetServiceNo() float32`

GetServiceNo returns the ServiceNo field if non-nil, zero value otherwise.

### GetServiceNoOk

`func (o *LineItemData) GetServiceNoOk() (*float32, bool)`

GetServiceNoOk returns a tuple with the ServiceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNo

`func (o *LineItemData) SetServiceNo(v float32)`

SetServiceNo sets ServiceNo field to given value.

### HasServiceNo

`func (o *LineItemData) HasServiceNo() bool`

HasServiceNo returns a boolean if a field has been set.

### GetUnits

`func (o *LineItemData) GetUnits() float32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *LineItemData) GetUnitsOk() (*float32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *LineItemData) SetUnits(v float32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *LineItemData) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetUsageTypeCd

`func (o *LineItemData) GetUsageTypeCd() string`

GetUsageTypeCd returns the UsageTypeCd field if non-nil, zero value otherwise.

### GetUsageTypeCdOk

`func (o *LineItemData) GetUsageTypeCdOk() (*string, bool)`

GetUsageTypeCdOk returns a tuple with the UsageTypeCd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTypeCd

`func (o *LineItemData) SetUsageTypeCd(v string)`

SetUsageTypeCd sets UsageTypeCd field to given value.

### HasUsageTypeCd

`func (o *LineItemData) HasUsageTypeCd() bool`

HasUsageTypeCd returns a boolean if a field has been set.

### SetUsageTypeCdNil

`func (o *LineItemData) SetUsageTypeCdNil(b bool)`

 SetUsageTypeCdNil sets the value for UsageTypeCd to be an explicit nil

### UnsetUsageTypeCd
`func (o *LineItemData) UnsetUsageTypeCd()`

UnsetUsageTypeCd ensures that no value is present for UsageTypeCd, not even an explicit nil
### GetUsageTypeNo

`func (o *LineItemData) GetUsageTypeNo() float32`

GetUsageTypeNo returns the UsageTypeNo field if non-nil, zero value otherwise.

### GetUsageTypeNoOk

`func (o *LineItemData) GetUsageTypeNoOk() (*float32, bool)`

GetUsageTypeNoOk returns a tuple with the UsageTypeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTypeNo

`func (o *LineItemData) SetUsageTypeNo(v float32)`

SetUsageTypeNo sets UsageTypeNo field to given value.

### HasUsageTypeNo

`func (o *LineItemData) HasUsageTypeNo() bool`

HasUsageTypeNo returns a boolean if a field has been set.

### SetUsageTypeNoNil

`func (o *LineItemData) SetUsageTypeNoNil(b bool)`

 SetUsageTypeNoNil sets the value for UsageTypeNo to be an explicit nil

### UnsetUsageTypeNo
`func (o *LineItemData) UnsetUsageTypeNo()`

UnsetUsageTypeNo ensures that no value is present for UsageTypeNo, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
