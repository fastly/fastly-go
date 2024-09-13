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
)

// LineItemData struct for LineItemData
type LineItemData struct {
	Amount NullableFloat32 `json:"amount,omitempty"`
	// An alphanumeric string identifying the invoice.
	AriaInvoiceID NullableString `json:"aria_invoice_id,omitempty"`
	ClientServiceID NullableString `json:"client_service_id,omitempty"`
	CreditCouponCode NullableString `json:"credit_coupon_code,omitempty"`
	Description NullableString `json:"description,omitempty"`
	ID *string `json:"id,omitempty"`
	LineNumber NullableInt32 `json:"line_number,omitempty"`
	PlanName NullableString `json:"plan_name,omitempty"`
	PlanNo NullableFloat32 `json:"plan_no,omitempty"`
	RatePerUnit NullableFloat32 `json:"rate_per_unit,omitempty"`
	RateScheduleNo NullableFloat32 `json:"rate_schedule_no,omitempty"`
	RateScheduleTierNo NullableFloat32 `json:"rate_schedule_tier_no,omitempty"`
	ServiceName NullableString `json:"service_name,omitempty"`
	ServiceNo NullableFloat32 `json:"service_no,omitempty"`
	Units NullableFloat32 `json:"units,omitempty"`
	UsageTypeCd NullableString `json:"usage_type_cd,omitempty"`
	UsageTypeNo NullableFloat32 `json:"usage_type_no,omitempty"`
	AdditionalProperties map[string]any
}

type _LineItemData LineItemData

// NewLineItemData instantiates a new LineItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemData() *LineItemData {
	this := LineItemData{}
	return &this
}

// NewLineItemDataWithDefaults instantiates a new LineItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataWithDefaults() *LineItemData {
	this := LineItemData{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetAmount() float32 {
	if o == nil || o.Amount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *LineItemData) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *LineItemData) SetAmount(v float32) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *LineItemData) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *LineItemData) UnsetAmount() {
	o.Amount.Unset()
}

// GetAriaInvoiceID returns the AriaInvoiceID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetAriaInvoiceID() string {
	if o == nil || o.AriaInvoiceID.Get() == nil {
		var ret string
		return ret
	}
	return *o.AriaInvoiceID.Get()
}

// GetAriaInvoiceIDOk returns a tuple with the AriaInvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetAriaInvoiceIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AriaInvoiceID.Get(), o.AriaInvoiceID.IsSet()
}

// HasAriaInvoiceID returns a boolean if a field has been set.
func (o *LineItemData) HasAriaInvoiceID() bool {
	if o != nil && o.AriaInvoiceID.IsSet() {
		return true
	}

	return false
}

// SetAriaInvoiceID gets a reference to the given NullableString and assigns it to the AriaInvoiceID field.
func (o *LineItemData) SetAriaInvoiceID(v string) {
	o.AriaInvoiceID.Set(&v)
}
// SetAriaInvoiceIDNil sets the value for AriaInvoiceID to be an explicit nil
func (o *LineItemData) SetAriaInvoiceIDNil() {
	o.AriaInvoiceID.Set(nil)
}

// UnsetAriaInvoiceID ensures that no value is present for AriaInvoiceID, not even an explicit nil
func (o *LineItemData) UnsetAriaInvoiceID() {
	o.AriaInvoiceID.Unset()
}

// GetClientServiceID returns the ClientServiceID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetClientServiceID() string {
	if o == nil || o.ClientServiceID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientServiceID.Get()
}

// GetClientServiceIDOk returns a tuple with the ClientServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetClientServiceIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientServiceID.Get(), o.ClientServiceID.IsSet()
}

// HasClientServiceID returns a boolean if a field has been set.
func (o *LineItemData) HasClientServiceID() bool {
	if o != nil && o.ClientServiceID.IsSet() {
		return true
	}

	return false
}

// SetClientServiceID gets a reference to the given NullableString and assigns it to the ClientServiceID field.
func (o *LineItemData) SetClientServiceID(v string) {
	o.ClientServiceID.Set(&v)
}
// SetClientServiceIDNil sets the value for ClientServiceID to be an explicit nil
func (o *LineItemData) SetClientServiceIDNil() {
	o.ClientServiceID.Set(nil)
}

// UnsetClientServiceID ensures that no value is present for ClientServiceID, not even an explicit nil
func (o *LineItemData) UnsetClientServiceID() {
	o.ClientServiceID.Unset()
}

// GetCreditCouponCode returns the CreditCouponCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetCreditCouponCode() string {
	if o == nil || o.CreditCouponCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreditCouponCode.Get()
}

// GetCreditCouponCodeOk returns a tuple with the CreditCouponCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetCreditCouponCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreditCouponCode.Get(), o.CreditCouponCode.IsSet()
}

// HasCreditCouponCode returns a boolean if a field has been set.
func (o *LineItemData) HasCreditCouponCode() bool {
	if o != nil && o.CreditCouponCode.IsSet() {
		return true
	}

	return false
}

// SetCreditCouponCode gets a reference to the given NullableString and assigns it to the CreditCouponCode field.
func (o *LineItemData) SetCreditCouponCode(v string) {
	o.CreditCouponCode.Set(&v)
}
// SetCreditCouponCodeNil sets the value for CreditCouponCode to be an explicit nil
func (o *LineItemData) SetCreditCouponCodeNil() {
	o.CreditCouponCode.Set(nil)
}

// UnsetCreditCouponCode ensures that no value is present for CreditCouponCode, not even an explicit nil
func (o *LineItemData) UnsetCreditCouponCode() {
	o.CreditCouponCode.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *LineItemData) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *LineItemData) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *LineItemData) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *LineItemData) UnsetDescription() {
	o.Description.Unset()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *LineItemData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *LineItemData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *LineItemData) SetID(v string) {
	o.ID = &v
}

// GetLineNumber returns the LineNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetLineNumber() int32 {
	if o == nil || o.LineNumber.Get() == nil {
		var ret int32
		return ret
	}
	return *o.LineNumber.Get()
}

// GetLineNumberOk returns a tuple with the LineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetLineNumberOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LineNumber.Get(), o.LineNumber.IsSet()
}

// HasLineNumber returns a boolean if a field has been set.
func (o *LineItemData) HasLineNumber() bool {
	if o != nil && o.LineNumber.IsSet() {
		return true
	}

	return false
}

// SetLineNumber gets a reference to the given NullableInt32 and assigns it to the LineNumber field.
func (o *LineItemData) SetLineNumber(v int32) {
	o.LineNumber.Set(&v)
}
// SetLineNumberNil sets the value for LineNumber to be an explicit nil
func (o *LineItemData) SetLineNumberNil() {
	o.LineNumber.Set(nil)
}

// UnsetLineNumber ensures that no value is present for LineNumber, not even an explicit nil
func (o *LineItemData) UnsetLineNumber() {
	o.LineNumber.Unset()
}

// GetPlanName returns the PlanName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetPlanName() string {
	if o == nil || o.PlanName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PlanName.Get()
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetPlanNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PlanName.Get(), o.PlanName.IsSet()
}

// HasPlanName returns a boolean if a field has been set.
func (o *LineItemData) HasPlanName() bool {
	if o != nil && o.PlanName.IsSet() {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given NullableString and assigns it to the PlanName field.
func (o *LineItemData) SetPlanName(v string) {
	o.PlanName.Set(&v)
}
// SetPlanNameNil sets the value for PlanName to be an explicit nil
func (o *LineItemData) SetPlanNameNil() {
	o.PlanName.Set(nil)
}

// UnsetPlanName ensures that no value is present for PlanName, not even an explicit nil
func (o *LineItemData) UnsetPlanName() {
	o.PlanName.Unset()
}

// GetPlanNo returns the PlanNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetPlanNo() float32 {
	if o == nil || o.PlanNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.PlanNo.Get()
}

// GetPlanNoOk returns a tuple with the PlanNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetPlanNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PlanNo.Get(), o.PlanNo.IsSet()
}

// HasPlanNo returns a boolean if a field has been set.
func (o *LineItemData) HasPlanNo() bool {
	if o != nil && o.PlanNo.IsSet() {
		return true
	}

	return false
}

// SetPlanNo gets a reference to the given NullableFloat32 and assigns it to the PlanNo field.
func (o *LineItemData) SetPlanNo(v float32) {
	o.PlanNo.Set(&v)
}
// SetPlanNoNil sets the value for PlanNo to be an explicit nil
func (o *LineItemData) SetPlanNoNil() {
	o.PlanNo.Set(nil)
}

// UnsetPlanNo ensures that no value is present for PlanNo, not even an explicit nil
func (o *LineItemData) UnsetPlanNo() {
	o.PlanNo.Unset()
}

// GetRatePerUnit returns the RatePerUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetRatePerUnit() float32 {
	if o == nil || o.RatePerUnit.Get() == nil {
		var ret float32
		return ret
	}
	return *o.RatePerUnit.Get()
}

// GetRatePerUnitOk returns a tuple with the RatePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetRatePerUnitOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RatePerUnit.Get(), o.RatePerUnit.IsSet()
}

// HasRatePerUnit returns a boolean if a field has been set.
func (o *LineItemData) HasRatePerUnit() bool {
	if o != nil && o.RatePerUnit.IsSet() {
		return true
	}

	return false
}

// SetRatePerUnit gets a reference to the given NullableFloat32 and assigns it to the RatePerUnit field.
func (o *LineItemData) SetRatePerUnit(v float32) {
	o.RatePerUnit.Set(&v)
}
// SetRatePerUnitNil sets the value for RatePerUnit to be an explicit nil
func (o *LineItemData) SetRatePerUnitNil() {
	o.RatePerUnit.Set(nil)
}

// UnsetRatePerUnit ensures that no value is present for RatePerUnit, not even an explicit nil
func (o *LineItemData) UnsetRatePerUnit() {
	o.RatePerUnit.Unset()
}

// GetRateScheduleNo returns the RateScheduleNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetRateScheduleNo() float32 {
	if o == nil || o.RateScheduleNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.RateScheduleNo.Get()
}

// GetRateScheduleNoOk returns a tuple with the RateScheduleNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetRateScheduleNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RateScheduleNo.Get(), o.RateScheduleNo.IsSet()
}

// HasRateScheduleNo returns a boolean if a field has been set.
func (o *LineItemData) HasRateScheduleNo() bool {
	if o != nil && o.RateScheduleNo.IsSet() {
		return true
	}

	return false
}

// SetRateScheduleNo gets a reference to the given NullableFloat32 and assigns it to the RateScheduleNo field.
func (o *LineItemData) SetRateScheduleNo(v float32) {
	o.RateScheduleNo.Set(&v)
}
// SetRateScheduleNoNil sets the value for RateScheduleNo to be an explicit nil
func (o *LineItemData) SetRateScheduleNoNil() {
	o.RateScheduleNo.Set(nil)
}

// UnsetRateScheduleNo ensures that no value is present for RateScheduleNo, not even an explicit nil
func (o *LineItemData) UnsetRateScheduleNo() {
	o.RateScheduleNo.Unset()
}

// GetRateScheduleTierNo returns the RateScheduleTierNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetRateScheduleTierNo() float32 {
	if o == nil || o.RateScheduleTierNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.RateScheduleTierNo.Get()
}

// GetRateScheduleTierNoOk returns a tuple with the RateScheduleTierNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetRateScheduleTierNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RateScheduleTierNo.Get(), o.RateScheduleTierNo.IsSet()
}

// HasRateScheduleTierNo returns a boolean if a field has been set.
func (o *LineItemData) HasRateScheduleTierNo() bool {
	if o != nil && o.RateScheduleTierNo.IsSet() {
		return true
	}

	return false
}

// SetRateScheduleTierNo gets a reference to the given NullableFloat32 and assigns it to the RateScheduleTierNo field.
func (o *LineItemData) SetRateScheduleTierNo(v float32) {
	o.RateScheduleTierNo.Set(&v)
}
// SetRateScheduleTierNoNil sets the value for RateScheduleTierNo to be an explicit nil
func (o *LineItemData) SetRateScheduleTierNoNil() {
	o.RateScheduleTierNo.Set(nil)
}

// UnsetRateScheduleTierNo ensures that no value is present for RateScheduleTierNo, not even an explicit nil
func (o *LineItemData) UnsetRateScheduleTierNo() {
	o.RateScheduleTierNo.Unset()
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetServiceName() string {
	if o == nil || o.ServiceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceName.Get()
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetServiceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceName.Get(), o.ServiceName.IsSet()
}

// HasServiceName returns a boolean if a field has been set.
func (o *LineItemData) HasServiceName() bool {
	if o != nil && o.ServiceName.IsSet() {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given NullableString and assigns it to the ServiceName field.
func (o *LineItemData) SetServiceName(v string) {
	o.ServiceName.Set(&v)
}
// SetServiceNameNil sets the value for ServiceName to be an explicit nil
func (o *LineItemData) SetServiceNameNil() {
	o.ServiceName.Set(nil)
}

// UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
func (o *LineItemData) UnsetServiceName() {
	o.ServiceName.Unset()
}

// GetServiceNo returns the ServiceNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetServiceNo() float32 {
	if o == nil || o.ServiceNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.ServiceNo.Get()
}

// GetServiceNoOk returns a tuple with the ServiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetServiceNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceNo.Get(), o.ServiceNo.IsSet()
}

// HasServiceNo returns a boolean if a field has been set.
func (o *LineItemData) HasServiceNo() bool {
	if o != nil && o.ServiceNo.IsSet() {
		return true
	}

	return false
}

// SetServiceNo gets a reference to the given NullableFloat32 and assigns it to the ServiceNo field.
func (o *LineItemData) SetServiceNo(v float32) {
	o.ServiceNo.Set(&v)
}
// SetServiceNoNil sets the value for ServiceNo to be an explicit nil
func (o *LineItemData) SetServiceNoNil() {
	o.ServiceNo.Set(nil)
}

// UnsetServiceNo ensures that no value is present for ServiceNo, not even an explicit nil
func (o *LineItemData) UnsetServiceNo() {
	o.ServiceNo.Unset()
}

// GetUnits returns the Units field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetUnits() float32 {
	if o == nil || o.Units.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Units.Get()
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetUnitsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Units.Get(), o.Units.IsSet()
}

// HasUnits returns a boolean if a field has been set.
func (o *LineItemData) HasUnits() bool {
	if o != nil && o.Units.IsSet() {
		return true
	}

	return false
}

// SetUnits gets a reference to the given NullableFloat32 and assigns it to the Units field.
func (o *LineItemData) SetUnits(v float32) {
	o.Units.Set(&v)
}
// SetUnitsNil sets the value for Units to be an explicit nil
func (o *LineItemData) SetUnitsNil() {
	o.Units.Set(nil)
}

// UnsetUnits ensures that no value is present for Units, not even an explicit nil
func (o *LineItemData) UnsetUnits() {
	o.Units.Unset()
}

// GetUsageTypeCd returns the UsageTypeCd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetUsageTypeCd() string {
	if o == nil || o.UsageTypeCd.Get() == nil {
		var ret string
		return ret
	}
	return *o.UsageTypeCd.Get()
}

// GetUsageTypeCdOk returns a tuple with the UsageTypeCd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetUsageTypeCdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsageTypeCd.Get(), o.UsageTypeCd.IsSet()
}

// HasUsageTypeCd returns a boolean if a field has been set.
func (o *LineItemData) HasUsageTypeCd() bool {
	if o != nil && o.UsageTypeCd.IsSet() {
		return true
	}

	return false
}

// SetUsageTypeCd gets a reference to the given NullableString and assigns it to the UsageTypeCd field.
func (o *LineItemData) SetUsageTypeCd(v string) {
	o.UsageTypeCd.Set(&v)
}
// SetUsageTypeCdNil sets the value for UsageTypeCd to be an explicit nil
func (o *LineItemData) SetUsageTypeCdNil() {
	o.UsageTypeCd.Set(nil)
}

// UnsetUsageTypeCd ensures that no value is present for UsageTypeCd, not even an explicit nil
func (o *LineItemData) UnsetUsageTypeCd() {
	o.UsageTypeCd.Unset()
}

// GetUsageTypeNo returns the UsageTypeNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemData) GetUsageTypeNo() float32 {
	if o == nil || o.UsageTypeNo.Get() == nil {
		var ret float32
		return ret
	}
	return *o.UsageTypeNo.Get()
}

// GetUsageTypeNoOk returns a tuple with the UsageTypeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemData) GetUsageTypeNoOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsageTypeNo.Get(), o.UsageTypeNo.IsSet()
}

// HasUsageTypeNo returns a boolean if a field has been set.
func (o *LineItemData) HasUsageTypeNo() bool {
	if o != nil && o.UsageTypeNo.IsSet() {
		return true
	}

	return false
}

// SetUsageTypeNo gets a reference to the given NullableFloat32 and assigns it to the UsageTypeNo field.
func (o *LineItemData) SetUsageTypeNo(v float32) {
	o.UsageTypeNo.Set(&v)
}
// SetUsageTypeNoNil sets the value for UsageTypeNo to be an explicit nil
func (o *LineItemData) SetUsageTypeNoNil() {
	o.UsageTypeNo.Set(nil)
}

// UnsetUsageTypeNo ensures that no value is present for UsageTypeNo, not even an explicit nil
func (o *LineItemData) UnsetUsageTypeNo() {
	o.UsageTypeNo.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LineItemData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.AriaInvoiceID.IsSet() {
		toSerialize["aria_invoice_id"] = o.AriaInvoiceID.Get()
	}
	if o.ClientServiceID.IsSet() {
		toSerialize["client_service_id"] = o.ClientServiceID.Get()
	}
	if o.CreditCouponCode.IsSet() {
		toSerialize["credit_coupon_code"] = o.CreditCouponCode.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.LineNumber.IsSet() {
		toSerialize["line_number"] = o.LineNumber.Get()
	}
	if o.PlanName.IsSet() {
		toSerialize["plan_name"] = o.PlanName.Get()
	}
	if o.PlanNo.IsSet() {
		toSerialize["plan_no"] = o.PlanNo.Get()
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
	if o.ServiceName.IsSet() {
		toSerialize["service_name"] = o.ServiceName.Get()
	}
	if o.ServiceNo.IsSet() {
		toSerialize["service_no"] = o.ServiceNo.Get()
	}
	if o.Units.IsSet() {
		toSerialize["units"] = o.Units.Get()
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
func (o *LineItemData) UnmarshalJSON(bytes []byte) (err error) {
	varLineItemData := _LineItemData{}

	if err = json.Unmarshal(bytes, &varLineItemData); err == nil {
		*o = LineItemData(varLineItemData)
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

// NullableLineItemData is a helper abstraction for handling nullable lineitemdata types. 
type NullableLineItemData struct {
	value *LineItemData
	isSet bool
}

// Get returns the value.
func (v NullableLineItemData) Get() *LineItemData {
	return v.value
}

// Set modifies the value.
func (v *NullableLineItemData) Set(val *LineItemData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLineItemData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLineItemData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLineItemData returns a pointer to a new instance of NullableLineItemData.
func NewNullableLineItemData(val *LineItemData) *NullableLineItemData {
	return &NullableLineItemData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLineItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLineItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
