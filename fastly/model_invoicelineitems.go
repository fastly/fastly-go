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

// Invoicelineitems struct for Invoicelineitems
type Invoicelineitems struct {
	// Invoice line item transaction name.
	Description *string `json:"description,omitempty"`
	// Billed amount for line item.
	Amount *float32 `json:"amount,omitempty"`
	// Discount coupon associated with the invoice for any account or service credits.
	CreditCouponCode *string `json:"credit_coupon_code,omitempty"`
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
	UsageType *string `json:"usage_type,omitempty"`
	AdditionalProperties map[string]any
}

type _Invoicelineitems Invoicelineitems

// NewInvoicelineitems instantiates a new Invoicelineitems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicelineitems() *Invoicelineitems {
	this := Invoicelineitems{}
	return &this
}

// NewInvoicelineitemsWithDefaults instantiates a new Invoicelineitems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicelineitemsWithDefaults() *Invoicelineitems {
	this := Invoicelineitems{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Invoicelineitems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoicelineitems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Invoicelineitems) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Invoicelineitems) SetDescription(v string) {
	o.Description = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Invoicelineitems) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoicelineitems) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Invoicelineitems) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *Invoicelineitems) SetAmount(v float32) {
	o.Amount = &v
}

// GetCreditCouponCode returns the CreditCouponCode field value if set, zero value otherwise.
func (o *Invoicelineitems) GetCreditCouponCode() string {
	if o == nil || o.CreditCouponCode == nil {
		var ret string
		return ret
	}
	return *o.CreditCouponCode
}

// GetCreditCouponCodeOk returns a tuple with the CreditCouponCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoicelineitems) GetCreditCouponCodeOk() (*string, bool) {
	if o == nil || o.CreditCouponCode == nil {
		return nil, false
	}
	return o.CreditCouponCode, true
}

// HasCreditCouponCode returns a boolean if a field has been set.
func (o *Invoicelineitems) HasCreditCouponCode() bool {
	if o != nil && o.CreditCouponCode != nil {
		return true
	}

	return false
}

// SetCreditCouponCode gets a reference to the given string and assigns it to the CreditCouponCode field.
func (o *Invoicelineitems) SetCreditCouponCode(v string) {
	o.CreditCouponCode = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *Invoicelineitems) GetRate() float32 {
	if o == nil || o.Rate == nil {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoicelineitems) GetRateOk() (*float32, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *Invoicelineitems) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *Invoicelineitems) SetRate(v float32) {
	o.Rate = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *Invoicelineitems) GetUnits() float32 {
	if o == nil || o.Units == nil {
		var ret float32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoicelineitems) GetUnitsOk() (*float32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *Invoicelineitems) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given float32 and assigns it to the Units field.
func (o *Invoicelineitems) SetUnits(v float32) {
	o.Units = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *Invoicelineitems) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoicelineitems) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *Invoicelineitems) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *Invoicelineitems) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductGroup returns the ProductGroup field value if set, zero value otherwise.
func (o *Invoicelineitems) GetProductGroup() string {
	if o == nil || o.ProductGroup == nil {
		var ret string
		return ret
	}
	return *o.ProductGroup
}

// GetProductGroupOk returns a tuple with the ProductGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoicelineitems) GetProductGroupOk() (*string, bool) {
	if o == nil || o.ProductGroup == nil {
		return nil, false
	}
	return o.ProductGroup, true
}

// HasProductGroup returns a boolean if a field has been set.
func (o *Invoicelineitems) HasProductGroup() bool {
	if o != nil && o.ProductGroup != nil {
		return true
	}

	return false
}

// SetProductGroup gets a reference to the given string and assigns it to the ProductGroup field.
func (o *Invoicelineitems) SetProductGroup(v string) {
	o.ProductGroup = &v
}

// GetProductLine returns the ProductLine field value if set, zero value otherwise.
func (o *Invoicelineitems) GetProductLine() string {
	if o == nil || o.ProductLine == nil {
		var ret string
		return ret
	}
	return *o.ProductLine
}

// GetProductLineOk returns a tuple with the ProductLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoicelineitems) GetProductLineOk() (*string, bool) {
	if o == nil || o.ProductLine == nil {
		return nil, false
	}
	return o.ProductLine, true
}

// HasProductLine returns a boolean if a field has been set.
func (o *Invoicelineitems) HasProductLine() bool {
	if o != nil && o.ProductLine != nil {
		return true
	}

	return false
}

// SetProductLine gets a reference to the given string and assigns it to the ProductLine field.
func (o *Invoicelineitems) SetProductLine(v string) {
	o.ProductLine = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Invoicelineitems) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoicelineitems) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Invoicelineitems) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Invoicelineitems) SetRegion(v string) {
	o.Region = &v
}

// GetUsageType returns the UsageType field value if set, zero value otherwise.
func (o *Invoicelineitems) GetUsageType() string {
	if o == nil || o.UsageType == nil {
		var ret string
		return ret
	}
	return *o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoicelineitems) GetUsageTypeOk() (*string, bool) {
	if o == nil || o.UsageType == nil {
		return nil, false
	}
	return o.UsageType, true
}

// HasUsageType returns a boolean if a field has been set.
func (o *Invoicelineitems) HasUsageType() bool {
	if o != nil && o.UsageType != nil {
		return true
	}

	return false
}

// SetUsageType gets a reference to the given string and assigns it to the UsageType field.
func (o *Invoicelineitems) SetUsageType(v string) {
	o.UsageType = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Invoicelineitems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.CreditCouponCode != nil {
		toSerialize["credit_coupon_code"] = o.CreditCouponCode
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
func (o *Invoicelineitems) UnmarshalJSON(bytes []byte) (err error) {
	varInvoicelineitems := _Invoicelineitems{}

	if err = json.Unmarshal(bytes, &varInvoicelineitems); err == nil {
		*o = Invoicelineitems(varInvoicelineitems)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "credit_coupon_code")
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

// NullableInvoicelineitems is a helper abstraction for handling nullable invoicelineitems types. 
type NullableInvoicelineitems struct {
	value *Invoicelineitems
	isSet bool
}

// Get returns the value.
func (v NullableInvoicelineitems) Get() *Invoicelineitems {
	return v.value
}

// Set modifies the value.
func (v *NullableInvoicelineitems) Set(val *Invoicelineitems) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInvoicelineitems) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInvoicelineitems) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInvoicelineitems returns a pointer to a new instance of NullableInvoicelineitems.
func NewNullableInvoicelineitems(val *Invoicelineitems) *NullableInvoicelineitems {
	return &NullableInvoicelineitems{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInvoicelineitems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableInvoicelineitems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
