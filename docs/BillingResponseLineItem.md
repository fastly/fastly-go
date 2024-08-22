# BillingResponseLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
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

### NewBillingResponseLineItem

`func NewBillingResponseLineItem() *BillingResponseLineItem`

NewBillingResponseLineItem instantiates a new BillingResponseLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingResponseLineItemWithDefaults

`func NewBillingResponseLineItemWithDefaults() *BillingResponseLineItem`

NewBillingResponseLineItemWithDefaults instantiates a new BillingResponseLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BillingResponseLineItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingResponseLineItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingResponseLineItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingResponseLineItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BillingResponseLineItem) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BillingResponseLineItem) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *BillingResponseLineItem) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BillingResponseLineItem) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BillingResponseLineItem) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BillingResponseLineItem) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BillingResponseLineItem) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BillingResponseLineItem) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BillingResponseLineItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingResponseLineItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingResponseLineItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BillingResponseLineItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BillingResponseLineItem) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BillingResponseLineItem) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetAmount

`func (o *BillingResponseLineItem) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingResponseLineItem) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingResponseLineItem) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingResponseLineItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAriaInvoiceID

`func (o *BillingResponseLineItem) GetAriaInvoiceID() string`

GetAriaInvoiceID returns the AriaInvoiceID field if non-nil, zero value otherwise.

### GetAriaInvoiceIDOk

`func (o *BillingResponseLineItem) GetAriaInvoiceIDOk() (*string, bool)`

GetAriaInvoiceIDOk returns a tuple with the AriaInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAriaInvoiceID

`func (o *BillingResponseLineItem) SetAriaInvoiceID(v string)`

SetAriaInvoiceID sets AriaInvoiceID field to given value.

### HasAriaInvoiceID

`func (o *BillingResponseLineItem) HasAriaInvoiceID() bool`

HasAriaInvoiceID returns a boolean if a field has been set.

### GetClientServiceID

`func (o *BillingResponseLineItem) GetClientServiceID() string`

GetClientServiceID returns the ClientServiceID field if non-nil, zero value otherwise.

### GetClientServiceIDOk

`func (o *BillingResponseLineItem) GetClientServiceIDOk() (*string, bool)`

GetClientServiceIDOk returns a tuple with the ClientServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientServiceID

`func (o *BillingResponseLineItem) SetClientServiceID(v string)`

SetClientServiceID sets ClientServiceID field to given value.

### HasClientServiceID

`func (o *BillingResponseLineItem) HasClientServiceID() bool`

HasClientServiceID returns a boolean if a field has been set.

### GetCreditCouponCode

`func (o *BillingResponseLineItem) GetCreditCouponCode() string`

GetCreditCouponCode returns the CreditCouponCode field if non-nil, zero value otherwise.

### GetCreditCouponCodeOk

`func (o *BillingResponseLineItem) GetCreditCouponCodeOk() (*string, bool)`

GetCreditCouponCodeOk returns a tuple with the CreditCouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCouponCode

`func (o *BillingResponseLineItem) SetCreditCouponCode(v string)`

SetCreditCouponCode sets CreditCouponCode field to given value.

### HasCreditCouponCode

`func (o *BillingResponseLineItem) HasCreditCouponCode() bool`

HasCreditCouponCode returns a boolean if a field has been set.

### SetCreditCouponCodeNil

`func (o *BillingResponseLineItem) SetCreditCouponCodeNil(b bool)`

 SetCreditCouponCodeNil sets the value for CreditCouponCode to be an explicit nil

### UnsetCreditCouponCode
`func (o *BillingResponseLineItem) UnsetCreditCouponCode()`

UnsetCreditCouponCode ensures that no value is present for CreditCouponCode, not even an explicit nil
### GetDescription

`func (o *BillingResponseLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingResponseLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingResponseLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingResponseLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetID

`func (o *BillingResponseLineItem) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *BillingResponseLineItem) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *BillingResponseLineItem) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *BillingResponseLineItem) HasID() bool`

HasID returns a boolean if a field has been set.

### GetLineNumber

`func (o *BillingResponseLineItem) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *BillingResponseLineItem) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *BillingResponseLineItem) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *BillingResponseLineItem) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetPlanName

`func (o *BillingResponseLineItem) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *BillingResponseLineItem) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *BillingResponseLineItem) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *BillingResponseLineItem) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanNo

`func (o *BillingResponseLineItem) GetPlanNo() float32`

GetPlanNo returns the PlanNo field if non-nil, zero value otherwise.

### GetPlanNoOk

`func (o *BillingResponseLineItem) GetPlanNoOk() (*float32, bool)`

GetPlanNoOk returns a tuple with the PlanNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanNo

`func (o *BillingResponseLineItem) SetPlanNo(v float32)`

SetPlanNo sets PlanNo field to given value.

### HasPlanNo

`func (o *BillingResponseLineItem) HasPlanNo() bool`

HasPlanNo returns a boolean if a field has been set.

### GetRatePerUnit

`func (o *BillingResponseLineItem) GetRatePerUnit() float32`

GetRatePerUnit returns the RatePerUnit field if non-nil, zero value otherwise.

### GetRatePerUnitOk

`func (o *BillingResponseLineItem) GetRatePerUnitOk() (*float32, bool)`

GetRatePerUnitOk returns a tuple with the RatePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerUnit

`func (o *BillingResponseLineItem) SetRatePerUnit(v float32)`

SetRatePerUnit sets RatePerUnit field to given value.

### HasRatePerUnit

`func (o *BillingResponseLineItem) HasRatePerUnit() bool`

HasRatePerUnit returns a boolean if a field has been set.

### SetRatePerUnitNil

`func (o *BillingResponseLineItem) SetRatePerUnitNil(b bool)`

 SetRatePerUnitNil sets the value for RatePerUnit to be an explicit nil

### UnsetRatePerUnit
`func (o *BillingResponseLineItem) UnsetRatePerUnit()`

UnsetRatePerUnit ensures that no value is present for RatePerUnit, not even an explicit nil
### GetRateScheduleNo

`func (o *BillingResponseLineItem) GetRateScheduleNo() float32`

GetRateScheduleNo returns the RateScheduleNo field if non-nil, zero value otherwise.

### GetRateScheduleNoOk

`func (o *BillingResponseLineItem) GetRateScheduleNoOk() (*float32, bool)`

GetRateScheduleNoOk returns a tuple with the RateScheduleNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateScheduleNo

`func (o *BillingResponseLineItem) SetRateScheduleNo(v float32)`

SetRateScheduleNo sets RateScheduleNo field to given value.

### HasRateScheduleNo

`func (o *BillingResponseLineItem) HasRateScheduleNo() bool`

HasRateScheduleNo returns a boolean if a field has been set.

### SetRateScheduleNoNil

`func (o *BillingResponseLineItem) SetRateScheduleNoNil(b bool)`

 SetRateScheduleNoNil sets the value for RateScheduleNo to be an explicit nil

### UnsetRateScheduleNo
`func (o *BillingResponseLineItem) UnsetRateScheduleNo()`

UnsetRateScheduleNo ensures that no value is present for RateScheduleNo, not even an explicit nil
### GetRateScheduleTierNo

`func (o *BillingResponseLineItem) GetRateScheduleTierNo() float32`

GetRateScheduleTierNo returns the RateScheduleTierNo field if non-nil, zero value otherwise.

### GetRateScheduleTierNoOk

`func (o *BillingResponseLineItem) GetRateScheduleTierNoOk() (*float32, bool)`

GetRateScheduleTierNoOk returns a tuple with the RateScheduleTierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateScheduleTierNo

`func (o *BillingResponseLineItem) SetRateScheduleTierNo(v float32)`

SetRateScheduleTierNo sets RateScheduleTierNo field to given value.

### HasRateScheduleTierNo

`func (o *BillingResponseLineItem) HasRateScheduleTierNo() bool`

HasRateScheduleTierNo returns a boolean if a field has been set.

### SetRateScheduleTierNoNil

`func (o *BillingResponseLineItem) SetRateScheduleTierNoNil(b bool)`

 SetRateScheduleTierNoNil sets the value for RateScheduleTierNo to be an explicit nil

### UnsetRateScheduleTierNo
`func (o *BillingResponseLineItem) UnsetRateScheduleTierNo()`

UnsetRateScheduleTierNo ensures that no value is present for RateScheduleTierNo, not even an explicit nil
### GetServiceName

`func (o *BillingResponseLineItem) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *BillingResponseLineItem) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *BillingResponseLineItem) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *BillingResponseLineItem) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceNo

`func (o *BillingResponseLineItem) GetServiceNo() float32`

GetServiceNo returns the ServiceNo field if non-nil, zero value otherwise.

### GetServiceNoOk

`func (o *BillingResponseLineItem) GetServiceNoOk() (*float32, bool)`

GetServiceNoOk returns a tuple with the ServiceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNo

`func (o *BillingResponseLineItem) SetServiceNo(v float32)`

SetServiceNo sets ServiceNo field to given value.

### HasServiceNo

`func (o *BillingResponseLineItem) HasServiceNo() bool`

HasServiceNo returns a boolean if a field has been set.

### GetUnits

`func (o *BillingResponseLineItem) GetUnits() float32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *BillingResponseLineItem) GetUnitsOk() (*float32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *BillingResponseLineItem) SetUnits(v float32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *BillingResponseLineItem) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetUsageTypeCd

`func (o *BillingResponseLineItem) GetUsageTypeCd() string`

GetUsageTypeCd returns the UsageTypeCd field if non-nil, zero value otherwise.

### GetUsageTypeCdOk

`func (o *BillingResponseLineItem) GetUsageTypeCdOk() (*string, bool)`

GetUsageTypeCdOk returns a tuple with the UsageTypeCd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTypeCd

`func (o *BillingResponseLineItem) SetUsageTypeCd(v string)`

SetUsageTypeCd sets UsageTypeCd field to given value.

### HasUsageTypeCd

`func (o *BillingResponseLineItem) HasUsageTypeCd() bool`

HasUsageTypeCd returns a boolean if a field has been set.

### SetUsageTypeCdNil

`func (o *BillingResponseLineItem) SetUsageTypeCdNil(b bool)`

 SetUsageTypeCdNil sets the value for UsageTypeCd to be an explicit nil

### UnsetUsageTypeCd
`func (o *BillingResponseLineItem) UnsetUsageTypeCd()`

UnsetUsageTypeCd ensures that no value is present for UsageTypeCd, not even an explicit nil
### GetUsageTypeNo

`func (o *BillingResponseLineItem) GetUsageTypeNo() float32`

GetUsageTypeNo returns the UsageTypeNo field if non-nil, zero value otherwise.

### GetUsageTypeNoOk

`func (o *BillingResponseLineItem) GetUsageTypeNoOk() (*float32, bool)`

GetUsageTypeNoOk returns a tuple with the UsageTypeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTypeNo

`func (o *BillingResponseLineItem) SetUsageTypeNo(v float32)`

SetUsageTypeNo sets UsageTypeNo field to given value.

### HasUsageTypeNo

`func (o *BillingResponseLineItem) HasUsageTypeNo() bool`

HasUsageTypeNo returns a boolean if a field has been set.

### SetUsageTypeNoNil

`func (o *BillingResponseLineItem) SetUsageTypeNoNil(b bool)`

 SetUsageTypeNoNil sets the value for UsageTypeNo to be an explicit nil

### UnsetUsageTypeNo
`func (o *BillingResponseLineItem) UnsetUsageTypeNo()`

UnsetUsageTypeNo ensures that no value is present for UsageTypeNo, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
