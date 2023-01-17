# BillingResponseLineItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**AriaInvoiceID** | Pointer to **string** |  | [optional] [readonly] 
**ClientServiceID** | Pointer to **string** |  | [optional] 
**CreditCouponCode** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**LineNumber** | Pointer to **int32** |  | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 
**PlanNo** | Pointer to **float32** |  | [optional] 
**RatePerUnit** | Pointer to **float32** |  | [optional] 
**RateScheduleNo** | Pointer to **NullableFloat32** |  | [optional] 
**RateScheduleTierNo** | Pointer to **NullableFloat32** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**ServiceNo** | Pointer to **float32** |  | [optional] 
**Units** | Pointer to **float32** |  | [optional] 
**UsageTypeCd** | Pointer to **NullableString** |  | [optional] 
**UsageTypeNo** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewBillingResponseLineItemAllOf

`func NewBillingResponseLineItemAllOf() *BillingResponseLineItemAllOf`

NewBillingResponseLineItemAllOf instantiates a new BillingResponseLineItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingResponseLineItemAllOfWithDefaults

`func NewBillingResponseLineItemAllOfWithDefaults() *BillingResponseLineItemAllOf`

NewBillingResponseLineItemAllOfWithDefaults instantiates a new BillingResponseLineItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingResponseLineItemAllOf) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingResponseLineItemAllOf) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingResponseLineItemAllOf) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingResponseLineItemAllOf) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAriaInvoiceID

`func (o *BillingResponseLineItemAllOf) GetAriaInvoiceID() string`

GetAriaInvoiceID returns the AriaInvoiceID field if non-nil, zero value otherwise.

### GetAriaInvoiceIDOk

`func (o *BillingResponseLineItemAllOf) GetAriaInvoiceIDOk() (*string, bool)`

GetAriaInvoiceIDOk returns a tuple with the AriaInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAriaInvoiceID

`func (o *BillingResponseLineItemAllOf) SetAriaInvoiceID(v string)`

SetAriaInvoiceID sets AriaInvoiceID field to given value.

### HasAriaInvoiceID

`func (o *BillingResponseLineItemAllOf) HasAriaInvoiceID() bool`

HasAriaInvoiceID returns a boolean if a field has been set.

### GetClientServiceID

`func (o *BillingResponseLineItemAllOf) GetClientServiceID() string`

GetClientServiceID returns the ClientServiceID field if non-nil, zero value otherwise.

### GetClientServiceIDOk

`func (o *BillingResponseLineItemAllOf) GetClientServiceIDOk() (*string, bool)`

GetClientServiceIDOk returns a tuple with the ClientServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientServiceID

`func (o *BillingResponseLineItemAllOf) SetClientServiceID(v string)`

SetClientServiceID sets ClientServiceID field to given value.

### HasClientServiceID

`func (o *BillingResponseLineItemAllOf) HasClientServiceID() bool`

HasClientServiceID returns a boolean if a field has been set.

### GetCreditCouponCode

`func (o *BillingResponseLineItemAllOf) GetCreditCouponCode() string`

GetCreditCouponCode returns the CreditCouponCode field if non-nil, zero value otherwise.

### GetCreditCouponCodeOk

`func (o *BillingResponseLineItemAllOf) GetCreditCouponCodeOk() (*string, bool)`

GetCreditCouponCodeOk returns a tuple with the CreditCouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCouponCode

`func (o *BillingResponseLineItemAllOf) SetCreditCouponCode(v string)`

SetCreditCouponCode sets CreditCouponCode field to given value.

### HasCreditCouponCode

`func (o *BillingResponseLineItemAllOf) HasCreditCouponCode() bool`

HasCreditCouponCode returns a boolean if a field has been set.

### SetCreditCouponCodeNil

`func (o *BillingResponseLineItemAllOf) SetCreditCouponCodeNil(b bool)`

 SetCreditCouponCodeNil sets the value for CreditCouponCode to be an explicit nil

### UnsetCreditCouponCode
`func (o *BillingResponseLineItemAllOf) UnsetCreditCouponCode()`

UnsetCreditCouponCode ensures that no value is present for CreditCouponCode, not even an explicit nil
### GetDescription

`func (o *BillingResponseLineItemAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingResponseLineItemAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingResponseLineItemAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingResponseLineItemAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetID

`func (o *BillingResponseLineItemAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *BillingResponseLineItemAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *BillingResponseLineItemAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *BillingResponseLineItemAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetLineNumber

`func (o *BillingResponseLineItemAllOf) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *BillingResponseLineItemAllOf) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *BillingResponseLineItemAllOf) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *BillingResponseLineItemAllOf) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetPlanName

`func (o *BillingResponseLineItemAllOf) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *BillingResponseLineItemAllOf) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *BillingResponseLineItemAllOf) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *BillingResponseLineItemAllOf) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanNo

`func (o *BillingResponseLineItemAllOf) GetPlanNo() float32`

GetPlanNo returns the PlanNo field if non-nil, zero value otherwise.

### GetPlanNoOk

`func (o *BillingResponseLineItemAllOf) GetPlanNoOk() (*float32, bool)`

GetPlanNoOk returns a tuple with the PlanNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanNo

`func (o *BillingResponseLineItemAllOf) SetPlanNo(v float32)`

SetPlanNo sets PlanNo field to given value.

### HasPlanNo

`func (o *BillingResponseLineItemAllOf) HasPlanNo() bool`

HasPlanNo returns a boolean if a field has been set.

### GetRatePerUnit

`func (o *BillingResponseLineItemAllOf) GetRatePerUnit() float32`

GetRatePerUnit returns the RatePerUnit field if non-nil, zero value otherwise.

### GetRatePerUnitOk

`func (o *BillingResponseLineItemAllOf) GetRatePerUnitOk() (*float32, bool)`

GetRatePerUnitOk returns a tuple with the RatePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerUnit

`func (o *BillingResponseLineItemAllOf) SetRatePerUnit(v float32)`

SetRatePerUnit sets RatePerUnit field to given value.

### HasRatePerUnit

`func (o *BillingResponseLineItemAllOf) HasRatePerUnit() bool`

HasRatePerUnit returns a boolean if a field has been set.

### GetRateScheduleNo

`func (o *BillingResponseLineItemAllOf) GetRateScheduleNo() float32`

GetRateScheduleNo returns the RateScheduleNo field if non-nil, zero value otherwise.

### GetRateScheduleNoOk

`func (o *BillingResponseLineItemAllOf) GetRateScheduleNoOk() (*float32, bool)`

GetRateScheduleNoOk returns a tuple with the RateScheduleNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateScheduleNo

`func (o *BillingResponseLineItemAllOf) SetRateScheduleNo(v float32)`

SetRateScheduleNo sets RateScheduleNo field to given value.

### HasRateScheduleNo

`func (o *BillingResponseLineItemAllOf) HasRateScheduleNo() bool`

HasRateScheduleNo returns a boolean if a field has been set.

### SetRateScheduleNoNil

`func (o *BillingResponseLineItemAllOf) SetRateScheduleNoNil(b bool)`

 SetRateScheduleNoNil sets the value for RateScheduleNo to be an explicit nil

### UnsetRateScheduleNo
`func (o *BillingResponseLineItemAllOf) UnsetRateScheduleNo()`

UnsetRateScheduleNo ensures that no value is present for RateScheduleNo, not even an explicit nil
### GetRateScheduleTierNo

`func (o *BillingResponseLineItemAllOf) GetRateScheduleTierNo() float32`

GetRateScheduleTierNo returns the RateScheduleTierNo field if non-nil, zero value otherwise.

### GetRateScheduleTierNoOk

`func (o *BillingResponseLineItemAllOf) GetRateScheduleTierNoOk() (*float32, bool)`

GetRateScheduleTierNoOk returns a tuple with the RateScheduleTierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateScheduleTierNo

`func (o *BillingResponseLineItemAllOf) SetRateScheduleTierNo(v float32)`

SetRateScheduleTierNo sets RateScheduleTierNo field to given value.

### HasRateScheduleTierNo

`func (o *BillingResponseLineItemAllOf) HasRateScheduleTierNo() bool`

HasRateScheduleTierNo returns a boolean if a field has been set.

### SetRateScheduleTierNoNil

`func (o *BillingResponseLineItemAllOf) SetRateScheduleTierNoNil(b bool)`

 SetRateScheduleTierNoNil sets the value for RateScheduleTierNo to be an explicit nil

### UnsetRateScheduleTierNo
`func (o *BillingResponseLineItemAllOf) UnsetRateScheduleTierNo()`

UnsetRateScheduleTierNo ensures that no value is present for RateScheduleTierNo, not even an explicit nil
### GetServiceName

`func (o *BillingResponseLineItemAllOf) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *BillingResponseLineItemAllOf) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *BillingResponseLineItemAllOf) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *BillingResponseLineItemAllOf) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceNo

`func (o *BillingResponseLineItemAllOf) GetServiceNo() float32`

GetServiceNo returns the ServiceNo field if non-nil, zero value otherwise.

### GetServiceNoOk

`func (o *BillingResponseLineItemAllOf) GetServiceNoOk() (*float32, bool)`

GetServiceNoOk returns a tuple with the ServiceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNo

`func (o *BillingResponseLineItemAllOf) SetServiceNo(v float32)`

SetServiceNo sets ServiceNo field to given value.

### HasServiceNo

`func (o *BillingResponseLineItemAllOf) HasServiceNo() bool`

HasServiceNo returns a boolean if a field has been set.

### GetUnits

`func (o *BillingResponseLineItemAllOf) GetUnits() float32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *BillingResponseLineItemAllOf) GetUnitsOk() (*float32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *BillingResponseLineItemAllOf) SetUnits(v float32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *BillingResponseLineItemAllOf) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetUsageTypeCd

`func (o *BillingResponseLineItemAllOf) GetUsageTypeCd() string`

GetUsageTypeCd returns the UsageTypeCd field if non-nil, zero value otherwise.

### GetUsageTypeCdOk

`func (o *BillingResponseLineItemAllOf) GetUsageTypeCdOk() (*string, bool)`

GetUsageTypeCdOk returns a tuple with the UsageTypeCd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTypeCd

`func (o *BillingResponseLineItemAllOf) SetUsageTypeCd(v string)`

SetUsageTypeCd sets UsageTypeCd field to given value.

### HasUsageTypeCd

`func (o *BillingResponseLineItemAllOf) HasUsageTypeCd() bool`

HasUsageTypeCd returns a boolean if a field has been set.

### SetUsageTypeCdNil

`func (o *BillingResponseLineItemAllOf) SetUsageTypeCdNil(b bool)`

 SetUsageTypeCdNil sets the value for UsageTypeCd to be an explicit nil

### UnsetUsageTypeCd
`func (o *BillingResponseLineItemAllOf) UnsetUsageTypeCd()`

UnsetUsageTypeCd ensures that no value is present for UsageTypeCd, not even an explicit nil
### GetUsageTypeNo

`func (o *BillingResponseLineItemAllOf) GetUsageTypeNo() float32`

GetUsageTypeNo returns the UsageTypeNo field if non-nil, zero value otherwise.

### GetUsageTypeNoOk

`func (o *BillingResponseLineItemAllOf) GetUsageTypeNoOk() (*float32, bool)`

GetUsageTypeNoOk returns a tuple with the UsageTypeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTypeNo

`func (o *BillingResponseLineItemAllOf) SetUsageTypeNo(v float32)`

SetUsageTypeNo sets UsageTypeNo field to given value.

### HasUsageTypeNo

`func (o *BillingResponseLineItemAllOf) HasUsageTypeNo() bool`

HasUsageTypeNo returns a boolean if a field has been set.

### SetUsageTypeNoNil

`func (o *BillingResponseLineItemAllOf) SetUsageTypeNoNil(b bool)`

 SetUsageTypeNoNil sets the value for UsageTypeNo to be an explicit nil

### UnsetUsageTypeNo
`func (o *BillingResponseLineItemAllOf) UnsetUsageTypeNo()`

UnsetUsageTypeNo ensures that no value is present for UsageTypeNo, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
