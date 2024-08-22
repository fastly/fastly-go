// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://www.fastly.com/documentation/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
	"time"
)

// BillingResponseLineItem struct for BillingResponseLineItem
type BillingResponseLineItem struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	Amount *float32 `json:"amount,omitempty"`
	// An alphanumeric string identifying the invoice.
	AriaInvoiceID *string `json:"aria_invoice_id,omitempty"`
	ClientServiceID *string `json:"client_service_id,omitempty"`
	CreditCouponCode NullableString `json:"credit_coupon_code,omitempty"`
	Description *string `json:"description,omitempty"`
	ID *string `json:"id,omitempty"`
	LineNumber *int32 `json:"line_number,omitempty"`
	PlanName *string `json:"plan_name,omitempty"`
	PlanNo *float32 `json:"plan_no,omitempty"`
	RatePerUnit NullableFloat32 `json:"rate_per_unit,omitempty"`
	RateScheduleNo NullableFloat32 `json:"rate_schedule_no,omitempty"`
	RateScheduleTierNo NullableFloat32 `json:"rate_schedule_tier_no,omitempty"`
	ServiceName *string `json:"service_name,omitempty"`
	ServiceNo *float32 `json:"service_no,omitempty"`
	Units *float32 `json:"units,omitempty"`
	UsageTypeCd NullableString `json:"usage_type_cd,omitempty"`
	UsageTypeNo NullableFloat32 `json:"usage_type_no,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingResponseLineItem BillingResponseLineItem

// NewBillingResponseLineItem instantiates a new BillingResponseLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingResponseLineItem() *BillingResponseLineItem {
	this := BillingResponseLineItem{}
	return &this
}

// NewBillingResponseLineItemWithDefaults instantiates a new BillingResponseLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingResponseLineItemWithDefaults() *BillingResponseLineItem {
	this := BillingResponseLineItem{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *BillingResponseLineItem) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *BillingResponseLineItem) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *BillingResponseLineItem) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItem) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItem) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *BillingResponseLineItem) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *BillingResponseLineItem) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *BillingResponseLineItem) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItem) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *BillingResponseLineItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *BillingResponseLineItem) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *BillingResponseLineItem) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *BillingResponseLineItem) SetAmount(v float32) {
	o.Amount = &v
}

// GetAriaInvoiceID returns the AriaInvoiceID field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetAriaInvoiceID() string {
	if o == nil || o.AriaInvoiceID == nil {
		var ret string
		return ret
	}
	return *o.AriaInvoiceID
}

// GetAriaInvoiceIDOk returns a tuple with the AriaInvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetAriaInvoiceIDOk() (*string, bool) {
	if o == nil || o.AriaInvoiceID == nil {
		return nil, false
	}
	return o.AriaInvoiceID, true
}

// HasAriaInvoiceID returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasAriaInvoiceID() bool {
	if o != nil && o.AriaInvoiceID != nil {
		return true
	}

	return false
}

// SetAriaInvoiceID gets a reference to the given string and assigns it to the AriaInvoiceID field.
func (o *BillingResponseLineItem) SetAriaInvoiceID(v string) {
	o.AriaInvoiceID = &v
}

// GetClientServiceID returns the ClientServiceID field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetClientServiceID() string {
	if o == nil || o.ClientServiceID == nil {
		var ret string
		return ret
	}
	return *o.ClientServiceID
}

// GetClientServiceIDOk returns a tuple with the ClientServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetClientServiceIDOk() (*string, bool) {
	if o == nil || o.ClientServiceID == nil {
		return nil, false
	}
	return o.ClientServiceID, true
}

// HasClientServiceID returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasClientServiceID() bool {
	if o != nil && o.ClientServiceID != nil {
		return true
	}

	return false
}

// SetClientServiceID gets a reference to the given string and assigns it to the ClientServiceID field.
func (o *BillingResponseLineItem) SetClientServiceID(v string) {
	o.ClientServiceID = &v
}

// GetCreditCouponCode returns the CreditCouponCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItem) GetCreditCouponCode() string {
	if o == nil || o.CreditCouponCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreditCouponCode.Get()
}

// GetCreditCouponCodeOk returns a tuple with the CreditCouponCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItem) GetCreditCouponCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreditCouponCode.Get(), o.CreditCouponCode.IsSet()
}

// HasCreditCouponCode returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasCreditCouponCode() bool {
	if o != nil && o.CreditCouponCode.IsSet() {
		return true
	}

	return false
}

// SetCreditCouponCode gets a reference to the given NullableString and assigns it to the CreditCouponCode field.
func (o *BillingResponseLineItem) SetCreditCouponCode(v string) {
	o.CreditCouponCode.Set(&v)
}
// SetCreditCouponCodeNil sets the value for CreditCouponCode to be an explicit nil
func (o *BillingResponseLineItem) SetCreditCouponCodeNil() {
	o.CreditCouponCode.Set(nil)
}

// UnsetCreditCouponCode ensures that no value is present for CreditCouponCode, not even an explicit nil
func (o *BillingResponseLineItem) UnsetCreditCouponCode() {
	o.CreditCouponCode.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BillingResponseLineItem) SetDescription(v string) {
	o.Description = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *BillingResponseLineItem) SetID(v string) {
	o.ID = &v
}

// GetLineNumber returns the LineNumber field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetLineNumber() int32 {
	if o == nil || o.LineNumber == nil {
		var ret int32
		return ret
	}
	return *o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetLineNumberOk() (*int32, bool) {
	if o == nil || o.LineNumber == nil {
		return nil, false
	}
	return o.LineNumber, true
}

// HasLineNumber returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasLineNumber() bool {
	if o != nil && o.LineNumber != nil {
		return true
	}

	return false
}

// SetLineNumber gets a reference to the given int32 and assigns it to the LineNumber field.
func (o *BillingResponseLineItem) SetLineNumber(v int32) {
	o.LineNumber = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetPlanName() string {
	if o == nil || o.PlanName == nil {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetPlanNameOk() (*string, bool) {
	if o == nil || o.PlanName == nil {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasPlanName() bool {
	if o != nil && o.PlanName != nil {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *BillingResponseLineItem) SetPlanName(v string) {
	o.PlanName = &v
}

// GetPlanNo returns the PlanNo field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetPlanNo() float32 {
	if o == nil || o.PlanNo == nil {
		var ret float32
		return ret
	}
	return *o.PlanNo
}

// GetPlanNoOk returns a tuple with the PlanNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetPlanNoOk() (*float32, bool) {
	if o == nil || o.PlanNo == nil {
		return nil, false
	}
	return o.PlanNo, true
}

// HasPlanNo returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasPlanNo() bool {
	if o != nil && o.PlanNo != nil {
		return true
	}

	return false
}

// SetPlanNo gets a reference to the given float32 and assigns it to the PlanNo field.
func (o *BillingResponseLineItem) SetPlanNo(v float32) {
	o.PlanNo = &v
}

// GetRatePerUnit returns the RatePerUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItem) GetRatePerUnit() float32 {
	if o == nil || o.RatePerUnit.Get() == nil {
		var ret float32
		return ret
	}
	return *o.RatePerUnit.Get()
}

// GetRatePerUnitOk returns a tuple with the RatePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItem) GetRatePerUnitOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RatePerUnit.Get(), o.RatePerUnit.IsSet()
}

// HasRatePerUnit returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasRatePerUnit() bool {
	if o != nil && o.RatePerUnit.IsSet() {
		return true
	}

	return false
}

// SetRatePerUnit gets a reference to the given NullableFloat32 and assigns it to the RatePerUnit field.
func (o *BillingResponseLineItem) SetRatePerUnit(v float32) {
	o.RatePerUnit.Set(&v)
}
// SetRatePerUnitNil sets the value for RatePerUnit to be an explicit nil
func (o *BillingResponseLineItem) SetRatePerUnitNil() {
	o.RatePerUnit.Set(nil)
}

// UnsetRatePerUnit ensures that no value is present for RatePerUnit, not even an explicit nil
func (o *BillingResponseLineItem) UnsetRatePerUnit() {
	o.RatePerUnit.Unset()
}

// GetRateScheduleNo returns the RateScheduleNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItem) GetRateScheduleNo() float32 {
	if o == nil || o.RateScheduleNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.RateScheduleNo.Get()
}

// GetRateScheduleNoOk returns a tuple with the RateScheduleNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItem) GetRateScheduleNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RateScheduleNo.Get(), o.RateScheduleNo.IsSet()
}

// HasRateScheduleNo returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasRateScheduleNo() bool {
	if o != nil && o.RateScheduleNo.IsSet() {
		return true
	}

	return false
}

// SetRateScheduleNo gets a reference to the given NullableFloat32 and assigns it to the RateScheduleNo field.
func (o *BillingResponseLineItem) SetRateScheduleNo(v float32) {
	o.RateScheduleNo.Set(&v)
}
// SetRateScheduleNoNil sets the value for RateScheduleNo to be an explicit nil
func (o *BillingResponseLineItem) SetRateScheduleNoNil() {
	o.RateScheduleNo.Set(nil)
}

// UnsetRateScheduleNo ensures that no value is present for RateScheduleNo, not even an explicit nil
func (o *BillingResponseLineItem) UnsetRateScheduleNo() {
	o.RateScheduleNo.Unset()
}

// GetRateScheduleTierNo returns the RateScheduleTierNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItem) GetRateScheduleTierNo() float32 {
	if o == nil || o.RateScheduleTierNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.RateScheduleTierNo.Get()
}

// GetRateScheduleTierNoOk returns a tuple with the RateScheduleTierNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItem) GetRateScheduleTierNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RateScheduleTierNo.Get(), o.RateScheduleTierNo.IsSet()
}

// HasRateScheduleTierNo returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasRateScheduleTierNo() bool {
	if o != nil && o.RateScheduleTierNo.IsSet() {
		return true
	}

	return false
}

// SetRateScheduleTierNo gets a reference to the given NullableFloat32 and assigns it to the RateScheduleTierNo field.
func (o *BillingResponseLineItem) SetRateScheduleTierNo(v float32) {
	o.RateScheduleTierNo.Set(&v)
}
// SetRateScheduleTierNoNil sets the value for RateScheduleTierNo to be an explicit nil
func (o *BillingResponseLineItem) SetRateScheduleTierNoNil() {
	o.RateScheduleTierNo.Set(nil)
}

// UnsetRateScheduleTierNo ensures that no value is present for RateScheduleTierNo, not even an explicit nil
func (o *BillingResponseLineItem) UnsetRateScheduleTierNo() {
	o.RateScheduleTierNo.Unset()
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *BillingResponseLineItem) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServiceNo returns the ServiceNo field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetServiceNo() float32 {
	if o == nil || o.ServiceNo == nil {
		var ret float32
		return ret
	}
	return *o.ServiceNo
}

// GetServiceNoOk returns a tuple with the ServiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetServiceNoOk() (*float32, bool) {
	if o == nil || o.ServiceNo == nil {
		return nil, false
	}
	return o.ServiceNo, true
}

// HasServiceNo returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasServiceNo() bool {
	if o != nil && o.ServiceNo != nil {
		return true
	}

	return false
}

// SetServiceNo gets a reference to the given float32 and assigns it to the ServiceNo field.
func (o *BillingResponseLineItem) SetServiceNo(v float32) {
	o.ServiceNo = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BillingResponseLineItem) GetUnits() float32 {
	if o == nil || o.Units == nil {
		var ret float32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItem) GetUnitsOk() (*float32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given float32 and assigns it to the Units field.
func (o *BillingResponseLineItem) SetUnits(v float32) {
	o.Units = &v
}

// GetUsageTypeCd returns the UsageTypeCd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItem) GetUsageTypeCd() string {
	if o == nil || o.UsageTypeCd.Get() == nil {
		var ret string
		return ret
	}
	return *o.UsageTypeCd.Get()
}

// GetUsageTypeCdOk returns a tuple with the UsageTypeCd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItem) GetUsageTypeCdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsageTypeCd.Get(), o.UsageTypeCd.IsSet()
}

// HasUsageTypeCd returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasUsageTypeCd() bool {
	if o != nil && o.UsageTypeCd.IsSet() {
		return true
	}

	return false
}

// SetUsageTypeCd gets a reference to the given NullableString and assigns it to the UsageTypeCd field.
func (o *BillingResponseLineItem) SetUsageTypeCd(v string) {
	o.UsageTypeCd.Set(&v)
}
// SetUsageTypeCdNil sets the value for UsageTypeCd to be an explicit nil
func (o *BillingResponseLineItem) SetUsageTypeCdNil() {
	o.UsageTypeCd.Set(nil)
}

// UnsetUsageTypeCd ensures that no value is present for UsageTypeCd, not even an explicit nil
func (o *BillingResponseLineItem) UnsetUsageTypeCd() {
	o.UsageTypeCd.Unset()
}

// GetUsageTypeNo returns the UsageTypeNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItem) GetUsageTypeNo() float32 {
	if o == nil || o.UsageTypeNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.UsageTypeNo.Get()
}

// GetUsageTypeNoOk returns a tuple with the UsageTypeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItem) GetUsageTypeNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsageTypeNo.Get(), o.UsageTypeNo.IsSet()
}

// HasUsageTypeNo returns a boolean if a field has been set.
func (o *BillingResponseLineItem) HasUsageTypeNo() bool {
	if o != nil && o.UsageTypeNo.IsSet() {
		return true
	}

	return false
}

// SetUsageTypeNo gets a reference to the given NullableFloat32 and assigns it to the UsageTypeNo field.
func (o *BillingResponseLineItem) SetUsageTypeNo(v float32) {
	o.UsageTypeNo.Set(&v)
}
// SetUsageTypeNoNil sets the value for UsageTypeNo to be an explicit nil
func (o *BillingResponseLineItem) SetUsageTypeNoNil() {
	o.UsageTypeNo.Set(nil)
}

// UnsetUsageTypeNo ensures that no value is present for UsageTypeNo, not even an explicit nil
func (o *BillingResponseLineItem) UnsetUsageTypeNo() {
	o.UsageTypeNo.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingResponseLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.AriaInvoiceID != nil {
		toSerialize["aria_invoice_id"] = o.AriaInvoiceID
	}
	if o.ClientServiceID != nil {
		toSerialize["client_service_id"] = o.ClientServiceID
	}
	if o.CreditCouponCode.IsSet() {
		toSerialize["credit_coupon_code"] = o.CreditCouponCode.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.LineNumber != nil {
		toSerialize["line_number"] = o.LineNumber
	}
	if o.PlanName != nil {
		toSerialize["plan_name"] = o.PlanName
	}
	if o.PlanNo != nil {
		toSerialize["plan_no"] = o.PlanNo
	}
	if o.RatePerUnit.IsSet() {
		toSerialize["rate_per_unit"] = o.RatePerUnit.Get()
	}
	if o.RateScheduleNo.IsSet() {
		toSerialize["rate_schedule_no"] = o.RateScheduleNo.Get()
	}
	if o.RateScheduleTierNo.IsSet() {
		toSerialize["rate_schedule_tier_no"] = o.RateScheduleTierNo.Get()
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.ServiceNo != nil {
		toSerialize["service_no"] = o.ServiceNo
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.UsageTypeCd.IsSet() {
		toSerialize["usage_type_cd"] = o.UsageTypeCd.Get()
	}
	if o.UsageTypeNo.IsSet() {
		toSerialize["usage_type_no"] = o.UsageTypeNo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingResponseLineItem) UnmarshalJSON(bytes []byte) (err error) {
	varBillingResponseLineItem := _BillingResponseLineItem{}

	if err = json.Unmarshal(bytes, &varBillingResponseLineItem); err == nil {
		*o = BillingResponseLineItem(varBillingResponseLineItem)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "aria_invoice_id")
		delete(additionalProperties, "client_service_id")
		delete(additionalProperties, "credit_coupon_code")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "line_number")
		delete(additionalProperties, "plan_name")
		delete(additionalProperties, "plan_no")
		delete(additionalProperties, "rate_per_unit")
		delete(additionalProperties, "rate_schedule_no")
		delete(additionalProperties, "rate_schedule_tier_no")
		delete(additionalProperties, "service_name")
		delete(additionalProperties, "service_no")
		delete(additionalProperties, "units")
		delete(additionalProperties, "usage_type_cd")
		delete(additionalProperties, "usage_type_no")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingResponseLineItem is a helper abstraction for handling nullable billingresponselineitem types. 
type NullableBillingResponseLineItem struct {
	value *BillingResponseLineItem
	isSet bool
}

// Get returns the value.
func (v NullableBillingResponseLineItem) Get() *BillingResponseLineItem {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingResponseLineItem) Set(val *BillingResponseLineItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingResponseLineItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingResponseLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingResponseLineItem returns a pointer to a new instance of NullableBillingResponseLineItem.
func NewNullableBillingResponseLineItem(val *BillingResponseLineItem) *NullableBillingResponseLineItem {
	return &NullableBillingResponseLineItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingResponseLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingResponseLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
