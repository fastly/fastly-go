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

// BillingEstimateLinesLineItems struct for BillingEstimateLinesLineItems
type BillingEstimateLinesLineItems struct {
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

type _BillingEstimateLinesLineItems BillingEstimateLinesLineItems

// NewBillingEstimateLinesLineItems instantiates a new BillingEstimateLinesLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingEstimateLinesLineItems() *BillingEstimateLinesLineItems {
	this := BillingEstimateLinesLineItems{}
	return &this
}

// NewBillingEstimateLinesLineItemsWithDefaults instantiates a new BillingEstimateLinesLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingEstimateLinesLineItemsWithDefaults() *BillingEstimateLinesLineItems {
	this := BillingEstimateLinesLineItems{}
	return &this
}

// GetPlanNo returns the PlanNo field value if set, zero value otherwise.
func (o *BillingEstimateLinesLineItems) GetPlanNo() int32 {
	if o == nil || o.PlanNo == nil {
		var ret int32
		return ret
	}
	return *o.PlanNo
}

// GetPlanNoOk returns a tuple with the PlanNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateLinesLineItems) GetPlanNoOk() (*int32, bool) {
	if o == nil || o.PlanNo == nil {
		return nil, false
	}
	return o.PlanNo, true
}

// HasPlanNo returns a boolean if a field has been set.
func (o *BillingEstimateLinesLineItems) HasPlanNo() bool {
	if o != nil && o.PlanNo != nil {
		return true
	}

	return false
}

// SetPlanNo gets a reference to the given int32 and assigns it to the PlanNo field.
func (o *BillingEstimateLinesLineItems) SetPlanNo(v int32) {
	o.PlanNo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BillingEstimateLinesLineItems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateLinesLineItems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BillingEstimateLinesLineItems) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BillingEstimateLinesLineItems) SetDescription(v string) {
	o.Description = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BillingEstimateLinesLineItems) GetUnits() float32 {
	if o == nil || o.Units == nil {
		var ret float32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateLinesLineItems) GetUnitsOk() (*float32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BillingEstimateLinesLineItems) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given float32 and assigns it to the Units field.
func (o *BillingEstimateLinesLineItems) SetUnits(v float32) {
	o.Units = &v
}

// GetPerUnitCost returns the PerUnitCost field value if set, zero value otherwise.
func (o *BillingEstimateLinesLineItems) GetPerUnitCost() float32 {
	if o == nil || o.PerUnitCost == nil {
		var ret float32
		return ret
	}
	return *o.PerUnitCost
}

// GetPerUnitCostOk returns a tuple with the PerUnitCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateLinesLineItems) GetPerUnitCostOk() (*float32, bool) {
	if o == nil || o.PerUnitCost == nil {
		return nil, false
	}
	return o.PerUnitCost, true
}

// HasPerUnitCost returns a boolean if a field has been set.
func (o *BillingEstimateLinesLineItems) HasPerUnitCost() bool {
	if o != nil && o.PerUnitCost != nil {
		return true
	}

	return false
}

// SetPerUnitCost gets a reference to the given float32 and assigns it to the PerUnitCost field.
func (o *BillingEstimateLinesLineItems) SetPerUnitCost(v float32) {
	o.PerUnitCost = &v
}

// GetServiceNo returns the ServiceNo field value if set, zero value otherwise.
func (o *BillingEstimateLinesLineItems) GetServiceNo() float32 {
	if o == nil || o.ServiceNo == nil {
		var ret float32
		return ret
	}
	return *o.ServiceNo
}

// GetServiceNoOk returns a tuple with the ServiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateLinesLineItems) GetServiceNoOk() (*float32, bool) {
	if o == nil || o.ServiceNo == nil {
		return nil, false
	}
	return o.ServiceNo, true
}

// HasServiceNo returns a boolean if a field has been set.
func (o *BillingEstimateLinesLineItems) HasServiceNo() bool {
	if o != nil && o.ServiceNo != nil {
		return true
	}

	return false
}

// SetServiceNo gets a reference to the given float32 and assigns it to the ServiceNo field.
func (o *BillingEstimateLinesLineItems) SetServiceNo(v float32) {
	o.ServiceNo = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *BillingEstimateLinesLineItems) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateLinesLineItems) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *BillingEstimateLinesLineItems) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *BillingEstimateLinesLineItems) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BillingEstimateLinesLineItems) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateLinesLineItems) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BillingEstimateLinesLineItems) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *BillingEstimateLinesLineItems) SetAmount(v float32) {
	o.Amount = &v
}

// GetClientServiceID returns the ClientServiceID field value if set, zero value otherwise.
func (o *BillingEstimateLinesLineItems) GetClientServiceID() string {
	if o == nil || o.ClientServiceID == nil {
		var ret string
		return ret
	}
	return *o.ClientServiceID
}

// GetClientServiceIDOk returns a tuple with the ClientServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateLinesLineItems) GetClientServiceIDOk() (*string, bool) {
	if o == nil || o.ClientServiceID == nil {
		return nil, false
	}
	return o.ClientServiceID, true
}

// HasClientServiceID returns a boolean if a field has been set.
func (o *BillingEstimateLinesLineItems) HasClientServiceID() bool {
	if o != nil && o.ClientServiceID != nil {
		return true
	}

	return false
}

// SetClientServiceID gets a reference to the given string and assigns it to the ClientServiceID field.
func (o *BillingEstimateLinesLineItems) SetClientServiceID(v string) {
	o.ClientServiceID = &v
}

// GetClientPlanID returns the ClientPlanID field value if set, zero value otherwise.
func (o *BillingEstimateLinesLineItems) GetClientPlanID() string {
	if o == nil || o.ClientPlanID == nil {
		var ret string
		return ret
	}
	return *o.ClientPlanID
}

// GetClientPlanIDOk returns a tuple with the ClientPlanID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateLinesLineItems) GetClientPlanIDOk() (*string, bool) {
	if o == nil || o.ClientPlanID == nil {
		return nil, false
	}
	return o.ClientPlanID, true
}

// HasClientPlanID returns a boolean if a field has been set.
func (o *BillingEstimateLinesLineItems) HasClientPlanID() bool {
	if o != nil && o.ClientPlanID != nil {
		return true
	}

	return false
}

// SetClientPlanID gets a reference to the given string and assigns it to the ClientPlanID field.
func (o *BillingEstimateLinesLineItems) SetClientPlanID(v string) {
	o.ClientPlanID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingEstimateLinesLineItems) MarshalJSON() ([]byte, error) {
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
func (o *BillingEstimateLinesLineItems) UnmarshalJSON(bytes []byte) (err error) {
	varBillingEstimateLinesLineItems := _BillingEstimateLinesLineItems{}

	if err = json.Unmarshal(bytes, &varBillingEstimateLinesLineItems); err == nil {
		*o = BillingEstimateLinesLineItems(varBillingEstimateLinesLineItems)
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

// NullableBillingEstimateLinesLineItems is a helper abstraction for handling nullable billingestimatelineslineitems types. 
type NullableBillingEstimateLinesLineItems struct {
	value *BillingEstimateLinesLineItems
	isSet bool
}

// Get returns the value.
func (v NullableBillingEstimateLinesLineItems) Get() *BillingEstimateLinesLineItems {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingEstimateLinesLineItems) Set(val *BillingEstimateLinesLineItems) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingEstimateLinesLineItems) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingEstimateLinesLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingEstimateLinesLineItems returns a pointer to a new instance of NullableBillingEstimateLinesLineItems.
func NewNullableBillingEstimateLinesLineItems(val *BillingEstimateLinesLineItems) *NullableBillingEstimateLinesLineItems {
	return &NullableBillingEstimateLinesLineItems{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingEstimateLinesLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingEstimateLinesLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
