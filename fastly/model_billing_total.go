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

// BillingTotal Complete summary of the billing information.
type BillingTotal struct {
	// The total amount of bandwidth used this month (See bandwidth_units for measurement).
	Bandwidth *float32 `json:"bandwidth,omitempty"`
	// The cost of the bandwidth used this month in USD.
	BandwidthCost *float32 `json:"bandwidth_cost,omitempty"`
	// Bandwidth measurement units based on billing plan.
	BandwidthUnits NullableString `json:"bandwidth_units,omitempty"`
	// The final amount to be paid.
	Cost *float32 `json:"cost,omitempty"`
	// Total incurred cost plus extras cost.
	CostBeforeDiscount *float32 `json:"cost_before_discount,omitempty"`
	// Calculated discount rate.
	Discount *float32 `json:"discount,omitempty"`
	// A list of any extras for this invoice.
	Extras []BillingTotalExtras `json:"extras,omitempty"`
	// Total cost of all extras.
	ExtrasCost *float32 `json:"extras_cost,omitempty"`
	// The total cost of bandwidth and requests used this month.
	IncurredCost *float32 `json:"incurred_cost,omitempty"`
	// How much over the plan minimum has been incurred.
	Overage *float32 `json:"overage,omitempty"`
	// The short code the plan this invoice was generated under.
	PlanCode *string `json:"plan_code,omitempty"`
	// The minimum cost of this plan.
	PlanMinimum *float32 `json:"plan_minimum,omitempty"`
	// The name of the plan this invoice was generated under.
	PlanName *string `json:"plan_name,omitempty"`
	// The total number of requests used this month.
	Requests *float32 `json:"requests,omitempty"`
	// The cost of the requests used this month.
	RequestsCost *float32 `json:"requests_cost,omitempty"`
	// Payment terms. Almost always Net15.
	Terms                *string `json:"terms,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingTotal BillingTotal

// NewBillingTotal instantiates a new BillingTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingTotal() *BillingTotal {
	this := BillingTotal{}
	return &this
}

// NewBillingTotalWithDefaults instantiates a new BillingTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingTotalWithDefaults() *BillingTotal {
	this := BillingTotal{}
	return &this
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *BillingTotal) GetBandwidth() float32 {
	if o == nil || o.Bandwidth == nil {
		var ret float32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetBandwidthOk() (*float32, bool) {
	if o == nil || o.Bandwidth == nil {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *BillingTotal) HasBandwidth() bool {
	if o != nil && o.Bandwidth != nil {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given float32 and assigns it to the Bandwidth field.
func (o *BillingTotal) SetBandwidth(v float32) {
	o.Bandwidth = &v
}

// GetBandwidthCost returns the BandwidthCost field value if set, zero value otherwise.
func (o *BillingTotal) GetBandwidthCost() float32 {
	if o == nil || o.BandwidthCost == nil {
		var ret float32
		return ret
	}
	return *o.BandwidthCost
}

// GetBandwidthCostOk returns a tuple with the BandwidthCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetBandwidthCostOk() (*float32, bool) {
	if o == nil || o.BandwidthCost == nil {
		return nil, false
	}
	return o.BandwidthCost, true
}

// HasBandwidthCost returns a boolean if a field has been set.
func (o *BillingTotal) HasBandwidthCost() bool {
	if o != nil && o.BandwidthCost != nil {
		return true
	}

	return false
}

// SetBandwidthCost gets a reference to the given float32 and assigns it to the BandwidthCost field.
func (o *BillingTotal) SetBandwidthCost(v float32) {
	o.BandwidthCost = &v
}

// GetBandwidthUnits returns the BandwidthUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingTotal) GetBandwidthUnits() string {
	if o == nil || o.BandwidthUnits.Get() == nil {
		var ret string
		return ret
	}
	return *o.BandwidthUnits.Get()
}

// GetBandwidthUnitsOk returns a tuple with the BandwidthUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingTotal) GetBandwidthUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BandwidthUnits.Get(), o.BandwidthUnits.IsSet()
}

// HasBandwidthUnits returns a boolean if a field has been set.
func (o *BillingTotal) HasBandwidthUnits() bool {
	if o != nil && o.BandwidthUnits.IsSet() {
		return true
	}

	return false
}

// SetBandwidthUnits gets a reference to the given NullableString and assigns it to the BandwidthUnits field.
func (o *BillingTotal) SetBandwidthUnits(v string) {
	o.BandwidthUnits.Set(&v)
}

// SetBandwidthUnitsNil sets the value for BandwidthUnits to be an explicit nil
func (o *BillingTotal) SetBandwidthUnitsNil() {
	o.BandwidthUnits.Set(nil)
}

// UnsetBandwidthUnits ensures that no value is present for BandwidthUnits, not even an explicit nil
func (o *BillingTotal) UnsetBandwidthUnits() {
	o.BandwidthUnits.Unset()
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingTotal) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingTotal) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingTotal) SetCost(v float32) {
	o.Cost = &v
}

// GetCostBeforeDiscount returns the CostBeforeDiscount field value if set, zero value otherwise.
func (o *BillingTotal) GetCostBeforeDiscount() float32 {
	if o == nil || o.CostBeforeDiscount == nil {
		var ret float32
		return ret
	}
	return *o.CostBeforeDiscount
}

// GetCostBeforeDiscountOk returns a tuple with the CostBeforeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetCostBeforeDiscountOk() (*float32, bool) {
	if o == nil || o.CostBeforeDiscount == nil {
		return nil, false
	}
	return o.CostBeforeDiscount, true
}

// HasCostBeforeDiscount returns a boolean if a field has been set.
func (o *BillingTotal) HasCostBeforeDiscount() bool {
	if o != nil && o.CostBeforeDiscount != nil {
		return true
	}

	return false
}

// SetCostBeforeDiscount gets a reference to the given float32 and assigns it to the CostBeforeDiscount field.
func (o *BillingTotal) SetCostBeforeDiscount(v float32) {
	o.CostBeforeDiscount = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *BillingTotal) GetDiscount() float32 {
	if o == nil || o.Discount == nil {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetDiscountOk() (*float32, bool) {
	if o == nil || o.Discount == nil {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *BillingTotal) HasDiscount() bool {
	if o != nil && o.Discount != nil {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *BillingTotal) SetDiscount(v float32) {
	o.Discount = &v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *BillingTotal) GetExtras() []BillingTotalExtras {
	if o == nil || o.Extras == nil {
		var ret []BillingTotalExtras
		return ret
	}
	return o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetExtrasOk() ([]BillingTotalExtras, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *BillingTotal) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given []BillingTotalExtras and assigns it to the Extras field.
func (o *BillingTotal) SetExtras(v []BillingTotalExtras) {
	o.Extras = v
}

// GetExtrasCost returns the ExtrasCost field value if set, zero value otherwise.
func (o *BillingTotal) GetExtrasCost() float32 {
	if o == nil || o.ExtrasCost == nil {
		var ret float32
		return ret
	}
	return *o.ExtrasCost
}

// GetExtrasCostOk returns a tuple with the ExtrasCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetExtrasCostOk() (*float32, bool) {
	if o == nil || o.ExtrasCost == nil {
		return nil, false
	}
	return o.ExtrasCost, true
}

// HasExtrasCost returns a boolean if a field has been set.
func (o *BillingTotal) HasExtrasCost() bool {
	if o != nil && o.ExtrasCost != nil {
		return true
	}

	return false
}

// SetExtrasCost gets a reference to the given float32 and assigns it to the ExtrasCost field.
func (o *BillingTotal) SetExtrasCost(v float32) {
	o.ExtrasCost = &v
}

// GetIncurredCost returns the IncurredCost field value if set, zero value otherwise.
func (o *BillingTotal) GetIncurredCost() float32 {
	if o == nil || o.IncurredCost == nil {
		var ret float32
		return ret
	}
	return *o.IncurredCost
}

// GetIncurredCostOk returns a tuple with the IncurredCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetIncurredCostOk() (*float32, bool) {
	if o == nil || o.IncurredCost == nil {
		return nil, false
	}
	return o.IncurredCost, true
}

// HasIncurredCost returns a boolean if a field has been set.
func (o *BillingTotal) HasIncurredCost() bool {
	if o != nil && o.IncurredCost != nil {
		return true
	}

	return false
}

// SetIncurredCost gets a reference to the given float32 and assigns it to the IncurredCost field.
func (o *BillingTotal) SetIncurredCost(v float32) {
	o.IncurredCost = &v
}

// GetOverage returns the Overage field value if set, zero value otherwise.
func (o *BillingTotal) GetOverage() float32 {
	if o == nil || o.Overage == nil {
		var ret float32
		return ret
	}
	return *o.Overage
}

// GetOverageOk returns a tuple with the Overage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetOverageOk() (*float32, bool) {
	if o == nil || o.Overage == nil {
		return nil, false
	}
	return o.Overage, true
}

// HasOverage returns a boolean if a field has been set.
func (o *BillingTotal) HasOverage() bool {
	if o != nil && o.Overage != nil {
		return true
	}

	return false
}

// SetOverage gets a reference to the given float32 and assigns it to the Overage field.
func (o *BillingTotal) SetOverage(v float32) {
	o.Overage = &v
}

// GetPlanCode returns the PlanCode field value if set, zero value otherwise.
func (o *BillingTotal) GetPlanCode() string {
	if o == nil || o.PlanCode == nil {
		var ret string
		return ret
	}
	return *o.PlanCode
}

// GetPlanCodeOk returns a tuple with the PlanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetPlanCodeOk() (*string, bool) {
	if o == nil || o.PlanCode == nil {
		return nil, false
	}
	return o.PlanCode, true
}

// HasPlanCode returns a boolean if a field has been set.
func (o *BillingTotal) HasPlanCode() bool {
	if o != nil && o.PlanCode != nil {
		return true
	}

	return false
}

// SetPlanCode gets a reference to the given string and assigns it to the PlanCode field.
func (o *BillingTotal) SetPlanCode(v string) {
	o.PlanCode = &v
}

// GetPlanMinimum returns the PlanMinimum field value if set, zero value otherwise.
func (o *BillingTotal) GetPlanMinimum() float32 {
	if o == nil || o.PlanMinimum == nil {
		var ret float32
		return ret
	}
	return *o.PlanMinimum
}

// GetPlanMinimumOk returns a tuple with the PlanMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetPlanMinimumOk() (*float32, bool) {
	if o == nil || o.PlanMinimum == nil {
		return nil, false
	}
	return o.PlanMinimum, true
}

// HasPlanMinimum returns a boolean if a field has been set.
func (o *BillingTotal) HasPlanMinimum() bool {
	if o != nil && o.PlanMinimum != nil {
		return true
	}

	return false
}

// SetPlanMinimum gets a reference to the given float32 and assigns it to the PlanMinimum field.
func (o *BillingTotal) SetPlanMinimum(v float32) {
	o.PlanMinimum = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *BillingTotal) GetPlanName() string {
	if o == nil || o.PlanName == nil {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetPlanNameOk() (*string, bool) {
	if o == nil || o.PlanName == nil {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *BillingTotal) HasPlanName() bool {
	if o != nil && o.PlanName != nil {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *BillingTotal) SetPlanName(v string) {
	o.PlanName = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *BillingTotal) GetRequests() float32 {
	if o == nil || o.Requests == nil {
		var ret float32
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetRequestsOk() (*float32, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *BillingTotal) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given float32 and assigns it to the Requests field.
func (o *BillingTotal) SetRequests(v float32) {
	o.Requests = &v
}

// GetRequestsCost returns the RequestsCost field value if set, zero value otherwise.
func (o *BillingTotal) GetRequestsCost() float32 {
	if o == nil || o.RequestsCost == nil {
		var ret float32
		return ret
	}
	return *o.RequestsCost
}

// GetRequestsCostOk returns a tuple with the RequestsCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetRequestsCostOk() (*float32, bool) {
	if o == nil || o.RequestsCost == nil {
		return nil, false
	}
	return o.RequestsCost, true
}

// HasRequestsCost returns a boolean if a field has been set.
func (o *BillingTotal) HasRequestsCost() bool {
	if o != nil && o.RequestsCost != nil {
		return true
	}

	return false
}

// SetRequestsCost gets a reference to the given float32 and assigns it to the RequestsCost field.
func (o *BillingTotal) SetRequestsCost(v float32) {
	o.RequestsCost = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *BillingTotal) GetTerms() string {
	if o == nil || o.Terms == nil {
		var ret string
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotal) GetTermsOk() (*string, bool) {
	if o == nil || o.Terms == nil {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *BillingTotal) HasTerms() bool {
	if o != nil && o.Terms != nil {
		return true
	}

	return false
}

// SetTerms gets a reference to the given string and assigns it to the Terms field.
func (o *BillingTotal) SetTerms(v string) {
	o.Terms = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingTotal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Bandwidth != nil {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if o.BandwidthCost != nil {
		toSerialize["bandwidth_cost"] = o.BandwidthCost
	}
	if o.BandwidthUnits.IsSet() {
		toSerialize["bandwidth_units"] = o.BandwidthUnits.Get()
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.CostBeforeDiscount != nil {
		toSerialize["cost_before_discount"] = o.CostBeforeDiscount
	}
	if o.Discount != nil {
		toSerialize["discount"] = o.Discount
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.ExtrasCost != nil {
		toSerialize["extras_cost"] = o.ExtrasCost
	}
	if o.IncurredCost != nil {
		toSerialize["incurred_cost"] = o.IncurredCost
	}
	if o.Overage != nil {
		toSerialize["overage"] = o.Overage
	}
	if o.PlanCode != nil {
		toSerialize["plan_code"] = o.PlanCode
	}
	if o.PlanMinimum != nil {
		toSerialize["plan_minimum"] = o.PlanMinimum
	}
	if o.PlanName != nil {
		toSerialize["plan_name"] = o.PlanName
	}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}
	if o.RequestsCost != nil {
		toSerialize["requests_cost"] = o.RequestsCost
	}
	if o.Terms != nil {
		toSerialize["terms"] = o.Terms
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BillingTotal) UnmarshalJSON(bytes []byte) (err error) {
	varBillingTotal := _BillingTotal{}

	if err = json.Unmarshal(bytes, &varBillingTotal); err == nil {
		*o = BillingTotal(varBillingTotal)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "bandwidth_cost")
		delete(additionalProperties, "bandwidth_units")
		delete(additionalProperties, "cost")
		delete(additionalProperties, "cost_before_discount")
		delete(additionalProperties, "discount")
		delete(additionalProperties, "extras")
		delete(additionalProperties, "extras_cost")
		delete(additionalProperties, "incurred_cost")
		delete(additionalProperties, "overage")
		delete(additionalProperties, "plan_code")
		delete(additionalProperties, "plan_minimum")
		delete(additionalProperties, "plan_name")
		delete(additionalProperties, "requests")
		delete(additionalProperties, "requests_cost")
		delete(additionalProperties, "terms")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingTotal is a helper abstraction for handling nullable billingtotal types.
type NullableBillingTotal struct {
	value *BillingTotal
	isSet bool
}

// Get returns the value.
func (v NullableBillingTotal) Get() *BillingTotal {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingTotal) Set(val *BillingTotal) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingTotal) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingTotal) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingTotal returns a pointer to a new instance of NullableBillingTotal.
func NewNullableBillingTotal(val *BillingTotal) *NullableBillingTotal {
	return &NullableBillingTotal{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBillingTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
