// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// BillingResponseLineItemAllOf struct for BillingResponseLineItemAllOf
type BillingResponseLineItemAllOf struct {
	Amount *float32 `json:"amount,omitempty"`
	AriaInvoiceID *string `json:"aria_invoice_id,omitempty"`
	ClientServiceID *string `json:"client_service_id,omitempty"`
	CreditCouponCode NullableString `json:"credit_coupon_code,omitempty"`
	Description *string `json:"description,omitempty"`
	ID *string `json:"id,omitempty"`
	LineNumber *int32 `json:"line_number,omitempty"`
	PlanName *string `json:"plan_name,omitempty"`
	PlanNo *float32 `json:"plan_no,omitempty"`
	RatePerUnit *float32 `json:"rate_per_unit,omitempty"`
	RateScheduleNo NullableFloat32 `json:"rate_schedule_no,omitempty"`
	RateScheduleTierNo NullableFloat32 `json:"rate_schedule_tier_no,omitempty"`
	ServiceName *string `json:"service_name,omitempty"`
	ServiceNo *float32 `json:"service_no,omitempty"`
	Units *float32 `json:"units,omitempty"`
	UsageTypeCd NullableString `json:"usage_type_cd,omitempty"`
	UsageTypeNo NullableFloat32 `json:"usage_type_no,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingResponseLineItemAllOf BillingResponseLineItemAllOf

// NewBillingResponseLineItemAllOf instantiates a new BillingResponseLineItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingResponseLineItemAllOf() *BillingResponseLineItemAllOf {
	this := BillingResponseLineItemAllOf{}
	return &this
}

// NewBillingResponseLineItemAllOfWithDefaults instantiates a new BillingResponseLineItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingResponseLineItemAllOfWithDefaults() *BillingResponseLineItemAllOf {
	this := BillingResponseLineItemAllOf{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *BillingResponseLineItemAllOf) SetAmount(v float32) {
	o.Amount = &v
}

// GetAriaInvoiceID returns the AriaInvoiceID field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetAriaInvoiceID() string {
	if o == nil || o.AriaInvoiceID == nil {
		var ret string
		return ret
	}
	return *o.AriaInvoiceID
}

// GetAriaInvoiceIDOk returns a tuple with the AriaInvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetAriaInvoiceIDOk() (*string, bool) {
	if o == nil || o.AriaInvoiceID == nil {
		return nil, false
	}
	return o.AriaInvoiceID, true
}

// HasAriaInvoiceID returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasAriaInvoiceID() bool {
	if o != nil && o.AriaInvoiceID != nil {
		return true
	}

	return false
}

// SetAriaInvoiceID gets a reference to the given string and assigns it to the AriaInvoiceID field.
func (o *BillingResponseLineItemAllOf) SetAriaInvoiceID(v string) {
	o.AriaInvoiceID = &v
}

// GetClientServiceID returns the ClientServiceID field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetClientServiceID() string {
	if o == nil || o.ClientServiceID == nil {
		var ret string
		return ret
	}
	return *o.ClientServiceID
}

// GetClientServiceIDOk returns a tuple with the ClientServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetClientServiceIDOk() (*string, bool) {
	if o == nil || o.ClientServiceID == nil {
		return nil, false
	}
	return o.ClientServiceID, true
}

// HasClientServiceID returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasClientServiceID() bool {
	if o != nil && o.ClientServiceID != nil {
		return true
	}

	return false
}

// SetClientServiceID gets a reference to the given string and assigns it to the ClientServiceID field.
func (o *BillingResponseLineItemAllOf) SetClientServiceID(v string) {
	o.ClientServiceID = &v
}

// GetCreditCouponCode returns the CreditCouponCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItemAllOf) GetCreditCouponCode() string {
	if o == nil || o.CreditCouponCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreditCouponCode.Get()
}

// GetCreditCouponCodeOk returns a tuple with the CreditCouponCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItemAllOf) GetCreditCouponCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreditCouponCode.Get(), o.CreditCouponCode.IsSet()
}

// HasCreditCouponCode returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasCreditCouponCode() bool {
	if o != nil && o.CreditCouponCode.IsSet() {
		return true
	}

	return false
}

// SetCreditCouponCode gets a reference to the given NullableString and assigns it to the CreditCouponCode field.
func (o *BillingResponseLineItemAllOf) SetCreditCouponCode(v string) {
	o.CreditCouponCode.Set(&v)
}
// SetCreditCouponCodeNil sets the value for CreditCouponCode to be an explicit nil
func (o *BillingResponseLineItemAllOf) SetCreditCouponCodeNil() {
	o.CreditCouponCode.Set(nil)
}

// UnsetCreditCouponCode ensures that no value is present for CreditCouponCode, not even an explicit nil
func (o *BillingResponseLineItemAllOf) UnsetCreditCouponCode() {
	o.CreditCouponCode.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BillingResponseLineItemAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *BillingResponseLineItemAllOf) SetID(v string) {
	o.ID = &v
}

// GetLineNumber returns the LineNumber field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetLineNumber() int32 {
	if o == nil || o.LineNumber == nil {
		var ret int32
		return ret
	}
	return *o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetLineNumberOk() (*int32, bool) {
	if o == nil || o.LineNumber == nil {
		return nil, false
	}
	return o.LineNumber, true
}

// HasLineNumber returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasLineNumber() bool {
	if o != nil && o.LineNumber != nil {
		return true
	}

	return false
}

// SetLineNumber gets a reference to the given int32 and assigns it to the LineNumber field.
func (o *BillingResponseLineItemAllOf) SetLineNumber(v int32) {
	o.LineNumber = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetPlanName() string {
	if o == nil || o.PlanName == nil {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetPlanNameOk() (*string, bool) {
	if o == nil || o.PlanName == nil {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasPlanName() bool {
	if o != nil && o.PlanName != nil {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *BillingResponseLineItemAllOf) SetPlanName(v string) {
	o.PlanName = &v
}

// GetPlanNo returns the PlanNo field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetPlanNo() float32 {
	if o == nil || o.PlanNo == nil {
		var ret float32
		return ret
	}
	return *o.PlanNo
}

// GetPlanNoOk returns a tuple with the PlanNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetPlanNoOk() (*float32, bool) {
	if o == nil || o.PlanNo == nil {
		return nil, false
	}
	return o.PlanNo, true
}

// HasPlanNo returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasPlanNo() bool {
	if o != nil && o.PlanNo != nil {
		return true
	}

	return false
}

// SetPlanNo gets a reference to the given float32 and assigns it to the PlanNo field.
func (o *BillingResponseLineItemAllOf) SetPlanNo(v float32) {
	o.PlanNo = &v
}

// GetRatePerUnit returns the RatePerUnit field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetRatePerUnit() float32 {
	if o == nil || o.RatePerUnit == nil {
		var ret float32
		return ret
	}
	return *o.RatePerUnit
}

// GetRatePerUnitOk returns a tuple with the RatePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetRatePerUnitOk() (*float32, bool) {
	if o == nil || o.RatePerUnit == nil {
		return nil, false
	}
	return o.RatePerUnit, true
}

// HasRatePerUnit returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasRatePerUnit() bool {
	if o != nil && o.RatePerUnit != nil {
		return true
	}

	return false
}

// SetRatePerUnit gets a reference to the given float32 and assigns it to the RatePerUnit field.
func (o *BillingResponseLineItemAllOf) SetRatePerUnit(v float32) {
	o.RatePerUnit = &v
}

// GetRateScheduleNo returns the RateScheduleNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItemAllOf) GetRateScheduleNo() float32 {
	if o == nil || o.RateScheduleNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.RateScheduleNo.Get()
}

// GetRateScheduleNoOk returns a tuple with the RateScheduleNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItemAllOf) GetRateScheduleNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RateScheduleNo.Get(), o.RateScheduleNo.IsSet()
}

// HasRateScheduleNo returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasRateScheduleNo() bool {
	if o != nil && o.RateScheduleNo.IsSet() {
		return true
	}

	return false
}

// SetRateScheduleNo gets a reference to the given NullableFloat32 and assigns it to the RateScheduleNo field.
func (o *BillingResponseLineItemAllOf) SetRateScheduleNo(v float32) {
	o.RateScheduleNo.Set(&v)
}
// SetRateScheduleNoNil sets the value for RateScheduleNo to be an explicit nil
func (o *BillingResponseLineItemAllOf) SetRateScheduleNoNil() {
	o.RateScheduleNo.Set(nil)
}

// UnsetRateScheduleNo ensures that no value is present for RateScheduleNo, not even an explicit nil
func (o *BillingResponseLineItemAllOf) UnsetRateScheduleNo() {
	o.RateScheduleNo.Unset()
}

// GetRateScheduleTierNo returns the RateScheduleTierNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItemAllOf) GetRateScheduleTierNo() float32 {
	if o == nil || o.RateScheduleTierNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.RateScheduleTierNo.Get()
}

// GetRateScheduleTierNoOk returns a tuple with the RateScheduleTierNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItemAllOf) GetRateScheduleTierNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RateScheduleTierNo.Get(), o.RateScheduleTierNo.IsSet()
}

// HasRateScheduleTierNo returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasRateScheduleTierNo() bool {
	if o != nil && o.RateScheduleTierNo.IsSet() {
		return true
	}

	return false
}

// SetRateScheduleTierNo gets a reference to the given NullableFloat32 and assigns it to the RateScheduleTierNo field.
func (o *BillingResponseLineItemAllOf) SetRateScheduleTierNo(v float32) {
	o.RateScheduleTierNo.Set(&v)
}
// SetRateScheduleTierNoNil sets the value for RateScheduleTierNo to be an explicit nil
func (o *BillingResponseLineItemAllOf) SetRateScheduleTierNoNil() {
	o.RateScheduleTierNo.Set(nil)
}

// UnsetRateScheduleTierNo ensures that no value is present for RateScheduleTierNo, not even an explicit nil
func (o *BillingResponseLineItemAllOf) UnsetRateScheduleTierNo() {
	o.RateScheduleTierNo.Unset()
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *BillingResponseLineItemAllOf) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServiceNo returns the ServiceNo field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetServiceNo() float32 {
	if o == nil || o.ServiceNo == nil {
		var ret float32
		return ret
	}
	return *o.ServiceNo
}

// GetServiceNoOk returns a tuple with the ServiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetServiceNoOk() (*float32, bool) {
	if o == nil || o.ServiceNo == nil {
		return nil, false
	}
	return o.ServiceNo, true
}

// HasServiceNo returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasServiceNo() bool {
	if o != nil && o.ServiceNo != nil {
		return true
	}

	return false
}

// SetServiceNo gets a reference to the given float32 and assigns it to the ServiceNo field.
func (o *BillingResponseLineItemAllOf) SetServiceNo(v float32) {
	o.ServiceNo = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BillingResponseLineItemAllOf) GetUnits() float32 {
	if o == nil || o.Units == nil {
		var ret float32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingResponseLineItemAllOf) GetUnitsOk() (*float32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given float32 and assigns it to the Units field.
func (o *BillingResponseLineItemAllOf) SetUnits(v float32) {
	o.Units = &v
}

// GetUsageTypeCd returns the UsageTypeCd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItemAllOf) GetUsageTypeCd() string {
	if o == nil || o.UsageTypeCd.Get() == nil {
		var ret string
		return ret
	}
	return *o.UsageTypeCd.Get()
}

// GetUsageTypeCdOk returns a tuple with the UsageTypeCd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItemAllOf) GetUsageTypeCdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsageTypeCd.Get(), o.UsageTypeCd.IsSet()
}

// HasUsageTypeCd returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasUsageTypeCd() bool {
	if o != nil && o.UsageTypeCd.IsSet() {
		return true
	}

	return false
}

// SetUsageTypeCd gets a reference to the given NullableString and assigns it to the UsageTypeCd field.
func (o *BillingResponseLineItemAllOf) SetUsageTypeCd(v string) {
	o.UsageTypeCd.Set(&v)
}
// SetUsageTypeCdNil sets the value for UsageTypeCd to be an explicit nil
func (o *BillingResponseLineItemAllOf) SetUsageTypeCdNil() {
	o.UsageTypeCd.Set(nil)
}

// UnsetUsageTypeCd ensures that no value is present for UsageTypeCd, not even an explicit nil
func (o *BillingResponseLineItemAllOf) UnsetUsageTypeCd() {
	o.UsageTypeCd.Unset()
}

// GetUsageTypeNo returns the UsageTypeNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingResponseLineItemAllOf) GetUsageTypeNo() float32 {
	if o == nil || o.UsageTypeNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.UsageTypeNo.Get()
}

// GetUsageTypeNoOk returns a tuple with the UsageTypeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingResponseLineItemAllOf) GetUsageTypeNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsageTypeNo.Get(), o.UsageTypeNo.IsSet()
}

// HasUsageTypeNo returns a boolean if a field has been set.
func (o *BillingResponseLineItemAllOf) HasUsageTypeNo() bool {
	if o != nil && o.UsageTypeNo.IsSet() {
		return true
	}

	return false
}

// SetUsageTypeNo gets a reference to the given NullableFloat32 and assigns it to the UsageTypeNo field.
func (o *BillingResponseLineItemAllOf) SetUsageTypeNo(v float32) {
	o.UsageTypeNo.Set(&v)
}
// SetUsageTypeNoNil sets the value for UsageTypeNo to be an explicit nil
func (o *BillingResponseLineItemAllOf) SetUsageTypeNoNil() {
	o.UsageTypeNo.Set(nil)
}

// UnsetUsageTypeNo ensures that no value is present for UsageTypeNo, not even an explicit nil
func (o *BillingResponseLineItemAllOf) UnsetUsageTypeNo() {
	o.UsageTypeNo.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingResponseLineItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
	if o.RatePerUnit != nil {
		toSerialize["rate_per_unit"] = o.RatePerUnit
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
func (o *BillingResponseLineItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBillingResponseLineItemAllOf := _BillingResponseLineItemAllOf{}

	if err = json.Unmarshal(bytes, &varBillingResponseLineItemAllOf); err == nil {
		*o = BillingResponseLineItemAllOf(varBillingResponseLineItemAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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

// NullableBillingResponseLineItemAllOf is a helper abstraction for handling nullable billingresponselineitemallof types. 
type NullableBillingResponseLineItemAllOf struct {
	value *BillingResponseLineItemAllOf
	isSet bool
}

// Get returns the value.
func (v NullableBillingResponseLineItemAllOf) Get() *BillingResponseLineItemAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingResponseLineItemAllOf) Set(val *BillingResponseLineItemAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingResponseLineItemAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingResponseLineItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingResponseLineItemAllOf returns a pointer to a new instance of NullableBillingResponseLineItemAllOf.
func NewNullableBillingResponseLineItemAllOf(val *BillingResponseLineItemAllOf) *NullableBillingResponseLineItemAllOf {
	return &NullableBillingResponseLineItemAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingResponseLineItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingResponseLineItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
