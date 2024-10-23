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

// Mtdlineitems struct for Mtdlineitems
type Mtdlineitems struct {
	// Invoice line item transaction name.
	Description *string `json:"description,omitempty"`
	// Billed amount for line item.
	Amount *float32 `json:"amount,omitempty"`
	// Price per unit.
	Rate *float32 `json:"rate,omitempty"`
	// Total number of units of usage.
	Units *float32 `json:"units,omitempty"`
	// The name of the product.
	ProductName *string `json:"product_name,omitempty"`
	// The broader classification of the product (e.g., `Compute` or `Full-Site Delivery`).
	ProductGroup *string `json:"product_group,omitempty"`
	// The broader classification of the product (e.g., `Network Services` or `Security`).
	ProductLine *string `json:"product_line,omitempty"`
	// The geographical area applicable for regionally based products.
	Region *string `json:"region,omitempty"`
	// The unit of measure (e.g., `requests` or `bandwidth`).
	UsageType            *string `json:"usage_type,omitempty"`
	AdditionalProperties map[string]any
}

type _Mtdlineitems Mtdlineitems

// NewMtdlineitems instantiates a new Mtdlineitems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtdlineitems() *Mtdlineitems {
	this := Mtdlineitems{}
	return &this
}

// NewMtdlineitemsWithDefaults instantiates a new Mtdlineitems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtdlineitemsWithDefaults() *Mtdlineitems {
	this := Mtdlineitems{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Mtdlineitems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdlineitems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Mtdlineitems) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Mtdlineitems) SetDescription(v string) {
	o.Description = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Mtdlineitems) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdlineitems) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Mtdlineitems) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *Mtdlineitems) SetAmount(v float32) {
	o.Amount = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *Mtdlineitems) GetRate() float32 {
	if o == nil || o.Rate == nil {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdlineitems) GetRateOk() (*float32, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *Mtdlineitems) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *Mtdlineitems) SetRate(v float32) {
	o.Rate = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *Mtdlineitems) GetUnits() float32 {
	if o == nil || o.Units == nil {
		var ret float32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdlineitems) GetUnitsOk() (*float32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *Mtdlineitems) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given float32 and assigns it to the Units field.
func (o *Mtdlineitems) SetUnits(v float32) {
	o.Units = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *Mtdlineitems) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdlineitems) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *Mtdlineitems) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *Mtdlineitems) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductGroup returns the ProductGroup field value if set, zero value otherwise.
func (o *Mtdlineitems) GetProductGroup() string {
	if o == nil || o.ProductGroup == nil {
		var ret string
		return ret
	}
	return *o.ProductGroup
}

// GetProductGroupOk returns a tuple with the ProductGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdlineitems) GetProductGroupOk() (*string, bool) {
	if o == nil || o.ProductGroup == nil {
		return nil, false
	}
	return o.ProductGroup, true
}

// HasProductGroup returns a boolean if a field has been set.
func (o *Mtdlineitems) HasProductGroup() bool {
	if o != nil && o.ProductGroup != nil {
		return true
	}

	return false
}

// SetProductGroup gets a reference to the given string and assigns it to the ProductGroup field.
func (o *Mtdlineitems) SetProductGroup(v string) {
	o.ProductGroup = &v
}

// GetProductLine returns the ProductLine field value if set, zero value otherwise.
func (o *Mtdlineitems) GetProductLine() string {
	if o == nil || o.ProductLine == nil {
		var ret string
		return ret
	}
	return *o.ProductLine
}

// GetProductLineOk returns a tuple with the ProductLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdlineitems) GetProductLineOk() (*string, bool) {
	if o == nil || o.ProductLine == nil {
		return nil, false
	}
	return o.ProductLine, true
}

// HasProductLine returns a boolean if a field has been set.
func (o *Mtdlineitems) HasProductLine() bool {
	if o != nil && o.ProductLine != nil {
		return true
	}

	return false
}

// SetProductLine gets a reference to the given string and assigns it to the ProductLine field.
func (o *Mtdlineitems) SetProductLine(v string) {
	o.ProductLine = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Mtdlineitems) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdlineitems) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Mtdlineitems) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Mtdlineitems) SetRegion(v string) {
	o.Region = &v
}

// GetUsageType returns the UsageType field value if set, zero value otherwise.
func (o *Mtdlineitems) GetUsageType() string {
	if o == nil || o.UsageType == nil {
		var ret string
		return ret
	}
	return *o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdlineitems) GetUsageTypeOk() (*string, bool) {
	if o == nil || o.UsageType == nil {
		return nil, false
	}
	return o.UsageType, true
}

// HasUsageType returns a boolean if a field has been set.
func (o *Mtdlineitems) HasUsageType() bool {
	if o != nil && o.UsageType != nil {
		return true
	}

	return false
}

// SetUsageType gets a reference to the given string and assigns it to the UsageType field.
func (o *Mtdlineitems) SetUsageType(v string) {
	o.UsageType = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Mtdlineitems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.ProductName != nil {
		toSerialize["product_name"] = o.ProductName
	}
	if o.ProductGroup != nil {
		toSerialize["product_group"] = o.ProductGroup
	}
	if o.ProductLine != nil {
		toSerialize["product_line"] = o.ProductLine
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.UsageType != nil {
		toSerialize["usage_type"] = o.UsageType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Mtdlineitems) UnmarshalJSON(bytes []byte) (err error) {
	varMtdlineitems := _Mtdlineitems{}

	if err = json.Unmarshal(bytes, &varMtdlineitems); err == nil {
		*o = Mtdlineitems(varMtdlineitems)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "rate")
		delete(additionalProperties, "units")
		delete(additionalProperties, "product_name")
		delete(additionalProperties, "product_group")
		delete(additionalProperties, "product_line")
		delete(additionalProperties, "region")
		delete(additionalProperties, "usage_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableMtdlineitems is a helper abstraction for handling nullable mtdlineitems types.
type NullableMtdlineitems struct {
	value *Mtdlineitems
	isSet bool
}

// Get returns the value.
func (v NullableMtdlineitems) Get() *Mtdlineitems {
	return v.value
}

// Set modifies the value.
func (v *NullableMtdlineitems) Set(val *Mtdlineitems) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableMtdlineitems) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableMtdlineitems) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableMtdlineitems returns a pointer to a new instance of NullableMtdlineitems.
func NewNullableMtdlineitems(val *Mtdlineitems) *NullableMtdlineitems {
	return &NullableMtdlineitems{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableMtdlineitems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableMtdlineitems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
