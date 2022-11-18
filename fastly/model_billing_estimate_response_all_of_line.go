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

// BillingEstimateResponseAllOfLine struct for BillingEstimateResponseAllOfLine
type BillingEstimateResponseAllOfLine struct {
	PlanNo *int32 `json:"plan_no,omitempty"`
	Description *string `json:"description,omitempty"`
	Units *float32 `json:"units,omitempty"`
	PerUnitCost *float32 `json:"per_unit_cost,omitempty"`
	ServiceNo *float32 `json:"service_no,omitempty"`
	ServiceType *string `json:"service_type,omitempty"`
	Amount *float32 `json:"amount,omitempty"`
	ClientServiceID *string `json:"client_service_id,omitempty"`
	ClientPlanID *string `json:"client_plan_id,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingEstimateResponseAllOfLine BillingEstimateResponseAllOfLine

// NewBillingEstimateResponseAllOfLine instantiates a new BillingEstimateResponseAllOfLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingEstimateResponseAllOfLine() *BillingEstimateResponseAllOfLine {
	this := BillingEstimateResponseAllOfLine{}
	return &this
}

// NewBillingEstimateResponseAllOfLineWithDefaults instantiates a new BillingEstimateResponseAllOfLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingEstimateResponseAllOfLineWithDefaults() *BillingEstimateResponseAllOfLine {
	this := BillingEstimateResponseAllOfLine{}
	return &this
}

// GetPlanNo returns the PlanNo field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOfLine) GetPlanNo() int32 {
	if o == nil || o.PlanNo == nil {
		var ret int32
		return ret
	}
	return *o.PlanNo
}

// GetPlanNoOk returns a tuple with the PlanNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOfLine) GetPlanNoOk() (*int32, bool) {
	if o == nil || o.PlanNo == nil {
		return nil, false
	}
	return o.PlanNo, true
}

// HasPlanNo returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOfLine) HasPlanNo() bool {
	if o != nil && o.PlanNo != nil {
		return true
	}

	return false
}

// SetPlanNo gets a reference to the given int32 and assigns it to the PlanNo field.
func (o *BillingEstimateResponseAllOfLine) SetPlanNo(v int32) {
	o.PlanNo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOfLine) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOfLine) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOfLine) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BillingEstimateResponseAllOfLine) SetDescription(v string) {
	o.Description = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOfLine) GetUnits() float32 {
	if o == nil || o.Units == nil {
		var ret float32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOfLine) GetUnitsOk() (*float32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOfLine) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given float32 and assigns it to the Units field.
func (o *BillingEstimateResponseAllOfLine) SetUnits(v float32) {
	o.Units = &v
}

// GetPerUnitCost returns the PerUnitCost field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOfLine) GetPerUnitCost() float32 {
	if o == nil || o.PerUnitCost == nil {
		var ret float32
		return ret
	}
	return *o.PerUnitCost
}

// GetPerUnitCostOk returns a tuple with the PerUnitCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOfLine) GetPerUnitCostOk() (*float32, bool) {
	if o == nil || o.PerUnitCost == nil {
		return nil, false
	}
	return o.PerUnitCost, true
}

// HasPerUnitCost returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOfLine) HasPerUnitCost() bool {
	if o != nil && o.PerUnitCost != nil {
		return true
	}

	return false
}

// SetPerUnitCost gets a reference to the given float32 and assigns it to the PerUnitCost field.
func (o *BillingEstimateResponseAllOfLine) SetPerUnitCost(v float32) {
	o.PerUnitCost = &v
}

// GetServiceNo returns the ServiceNo field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOfLine) GetServiceNo() float32 {
	if o == nil || o.ServiceNo == nil {
		var ret float32
		return ret
	}
	return *o.ServiceNo
}

// GetServiceNoOk returns a tuple with the ServiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOfLine) GetServiceNoOk() (*float32, bool) {
	if o == nil || o.ServiceNo == nil {
		return nil, false
	}
	return o.ServiceNo, true
}

// HasServiceNo returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOfLine) HasServiceNo() bool {
	if o != nil && o.ServiceNo != nil {
		return true
	}

	return false
}

// SetServiceNo gets a reference to the given float32 and assigns it to the ServiceNo field.
func (o *BillingEstimateResponseAllOfLine) SetServiceNo(v float32) {
	o.ServiceNo = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOfLine) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOfLine) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOfLine) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *BillingEstimateResponseAllOfLine) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOfLine) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOfLine) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOfLine) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *BillingEstimateResponseAllOfLine) SetAmount(v float32) {
	o.Amount = &v
}

// GetClientServiceID returns the ClientServiceID field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOfLine) GetClientServiceID() string {
	if o == nil || o.ClientServiceID == nil {
		var ret string
		return ret
	}
	return *o.ClientServiceID
}

// GetClientServiceIDOk returns a tuple with the ClientServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOfLine) GetClientServiceIDOk() (*string, bool) {
	if o == nil || o.ClientServiceID == nil {
		return nil, false
	}
	return o.ClientServiceID, true
}

// HasClientServiceID returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOfLine) HasClientServiceID() bool {
	if o != nil && o.ClientServiceID != nil {
		return true
	}

	return false
}

// SetClientServiceID gets a reference to the given string and assigns it to the ClientServiceID field.
func (o *BillingEstimateResponseAllOfLine) SetClientServiceID(v string) {
	o.ClientServiceID = &v
}

// GetClientPlanID returns the ClientPlanID field value if set, zero value otherwise.
func (o *BillingEstimateResponseAllOfLine) GetClientPlanID() string {
	if o == nil || o.ClientPlanID == nil {
		var ret string
		return ret
	}
	return *o.ClientPlanID
}

// GetClientPlanIDOk returns a tuple with the ClientPlanID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponseAllOfLine) GetClientPlanIDOk() (*string, bool) {
	if o == nil || o.ClientPlanID == nil {
		return nil, false
	}
	return o.ClientPlanID, true
}

// HasClientPlanID returns a boolean if a field has been set.
func (o *BillingEstimateResponseAllOfLine) HasClientPlanID() bool {
	if o != nil && o.ClientPlanID != nil {
		return true
	}

	return false
}

// SetClientPlanID gets a reference to the given string and assigns it to the ClientPlanID field.
func (o *BillingEstimateResponseAllOfLine) SetClientPlanID(v string) {
	o.ClientPlanID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingEstimateResponseAllOfLine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.PlanNo != nil {
		toSerialize["plan_no"] = o.PlanNo
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.PerUnitCost != nil {
		toSerialize["per_unit_cost"] = o.PerUnitCost
	}
	if o.ServiceNo != nil {
		toSerialize["service_no"] = o.ServiceNo
	}
	if o.ServiceType != nil {
		toSerialize["service_type"] = o.ServiceType
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.ClientServiceID != nil {
		toSerialize["client_service_id"] = o.ClientServiceID
	}
	if o.ClientPlanID != nil {
		toSerialize["client_plan_id"] = o.ClientPlanID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingEstimateResponseAllOfLine) UnmarshalJSON(bytes []byte) (err error) {
	varBillingEstimateResponseAllOfLine := _BillingEstimateResponseAllOfLine{}

	if err = json.Unmarshal(bytes, &varBillingEstimateResponseAllOfLine); err == nil {
		*o = BillingEstimateResponseAllOfLine(varBillingEstimateResponseAllOfLine)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "plan_no")
		delete(additionalProperties, "description")
		delete(additionalProperties, "units")
		delete(additionalProperties, "per_unit_cost")
		delete(additionalProperties, "service_no")
		delete(additionalProperties, "service_type")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "client_service_id")
		delete(additionalProperties, "client_plan_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingEstimateResponseAllOfLine is a helper abstraction for handling nullable billingestimateresponseallofline types. 
type NullableBillingEstimateResponseAllOfLine struct {
	value *BillingEstimateResponseAllOfLine
	isSet bool
}

// Get returns the value.
func (v NullableBillingEstimateResponseAllOfLine) Get() *BillingEstimateResponseAllOfLine {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingEstimateResponseAllOfLine) Set(val *BillingEstimateResponseAllOfLine) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingEstimateResponseAllOfLine) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingEstimateResponseAllOfLine) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingEstimateResponseAllOfLine returns a pointer to a new instance of NullableBillingEstimateResponseAllOfLine.
func NewNullableBillingEstimateResponseAllOfLine(val *BillingEstimateResponseAllOfLine) *NullableBillingEstimateResponseAllOfLine {
	return &NullableBillingEstimateResponseAllOfLine{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingEstimateResponseAllOfLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingEstimateResponseAllOfLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
