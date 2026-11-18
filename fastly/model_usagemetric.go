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

// Usagemetric struct for Usagemetric
type Usagemetric struct {
	// The year and month of the usage element.
	Month *string `json:"month,omitempty"`
	// The usage type identifier for the usage. This is a single, billable metric for the product.
	UsageType *string `json:"usage_type,omitempty"`
	// Full name of the product usage type as it might appear on a customer's invoice.
	Name *string `json:"name,omitempty"`
	// The geographical area applicable for regionally based products.
	Region *string `json:"region,omitempty"`
	// The unit for the usage as shown on an invoice. If there is no explicit unit, this field will be \"unit\" (e.g., a request with `product_id` of 'cdn_usage' and `usage_type` of 'North America Requests' has no unit, and will return \"unit\").
	Unit *string `json:"unit,omitempty"`
	// The quantity of the usage for the product.
	Quantity *float32 `json:"quantity,omitempty"`
	// The raw units measured for the product.
	RawQuantity *float32 `json:"raw_quantity,omitempty"`
	// The product identifier associated with the usage type. This corresponds to a Fastly product offering.
	ProductID *string `json:"product_id,omitempty"`
	// The date when the usage metric was last updated.
	LastUpdatedAt        *string `json:"last_updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _Usagemetric Usagemetric

// NewUsagemetric instantiates a new Usagemetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsagemetric() *Usagemetric {
	this := Usagemetric{}
	return &this
}

// NewUsagemetricWithDefaults instantiates a new Usagemetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsagemetricWithDefaults() *Usagemetric {
	this := Usagemetric{}
	return &this
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *Usagemetric) GetMonth() string {
	if o == nil || o.Month == nil {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usagemetric) GetMonthOk() (*string, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *Usagemetric) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *Usagemetric) SetMonth(v string) {
	o.Month = &v
}

// GetUsageType returns the UsageType field value if set, zero value otherwise.
func (o *Usagemetric) GetUsageType() string {
	if o == nil || o.UsageType == nil {
		var ret string
		return ret
	}
	return *o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usagemetric) GetUsageTypeOk() (*string, bool) {
	if o == nil || o.UsageType == nil {
		return nil, false
	}
	return o.UsageType, true
}

// HasUsageType returns a boolean if a field has been set.
func (o *Usagemetric) HasUsageType() bool {
	if o != nil && o.UsageType != nil {
		return true
	}

	return false
}

// SetUsageType gets a reference to the given string and assigns it to the UsageType field.
func (o *Usagemetric) SetUsageType(v string) {
	o.UsageType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Usagemetric) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usagemetric) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Usagemetric) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Usagemetric) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Usagemetric) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usagemetric) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Usagemetric) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Usagemetric) SetRegion(v string) {
	o.Region = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Usagemetric) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usagemetric) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Usagemetric) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Usagemetric) SetUnit(v string) {
	o.Unit = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Usagemetric) GetQuantity() float32 {
	if o == nil || o.Quantity == nil {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usagemetric) GetQuantityOk() (*float32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Usagemetric) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *Usagemetric) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetRawQuantity returns the RawQuantity field value if set, zero value otherwise.
func (o *Usagemetric) GetRawQuantity() float32 {
	if o == nil || o.RawQuantity == nil {
		var ret float32
		return ret
	}
	return *o.RawQuantity
}

// GetRawQuantityOk returns a tuple with the RawQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usagemetric) GetRawQuantityOk() (*float32, bool) {
	if o == nil || o.RawQuantity == nil {
		return nil, false
	}
	return o.RawQuantity, true
}

// HasRawQuantity returns a boolean if a field has been set.
func (o *Usagemetric) HasRawQuantity() bool {
	if o != nil && o.RawQuantity != nil {
		return true
	}

	return false
}

// SetRawQuantity gets a reference to the given float32 and assigns it to the RawQuantity field.
func (o *Usagemetric) SetRawQuantity(v float32) {
	o.RawQuantity = &v
}

// GetProductID returns the ProductID field value if set, zero value otherwise.
func (o *Usagemetric) GetProductID() string {
	if o == nil || o.ProductID == nil {
		var ret string
		return ret
	}
	return *o.ProductID
}

// GetProductIDOk returns a tuple with the ProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usagemetric) GetProductIDOk() (*string, bool) {
	if o == nil || o.ProductID == nil {
		return nil, false
	}
	return o.ProductID, true
}

// HasProductID returns a boolean if a field has been set.
func (o *Usagemetric) HasProductID() bool {
	if o != nil && o.ProductID != nil {
		return true
	}

	return false
}

// SetProductID gets a reference to the given string and assigns it to the ProductID field.
func (o *Usagemetric) SetProductID(v string) {
	o.ProductID = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *Usagemetric) GetLastUpdatedAt() string {
	if o == nil || o.LastUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usagemetric) GetLastUpdatedAtOk() (*string, bool) {
	if o == nil || o.LastUpdatedAt == nil {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *Usagemetric) HasLastUpdatedAt() bool {
	if o != nil && o.LastUpdatedAt != nil {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given string and assigns it to the LastUpdatedAt field.
func (o *Usagemetric) SetLastUpdatedAt(v string) {
	o.LastUpdatedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Usagemetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.UsageType != nil {
		toSerialize["usage_type"] = o.UsageType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.RawQuantity != nil {
		toSerialize["raw_quantity"] = o.RawQuantity
	}
	if o.ProductID != nil {
		toSerialize["product_id"] = o.ProductID
	}
	if o.LastUpdatedAt != nil {
		toSerialize["last_updated_at"] = o.LastUpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Usagemetric) UnmarshalJSON(bytes []byte) (err error) {
	varUsagemetric := _Usagemetric{}

	if err = json.Unmarshal(bytes, &varUsagemetric); err == nil {
		*o = Usagemetric(varUsagemetric)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "month")
		delete(additionalProperties, "usage_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "region")
		delete(additionalProperties, "unit")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "raw_quantity")
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "last_updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableUsagemetric is a helper abstraction for handling nullable usagemetric types.
type NullableUsagemetric struct {
	value *Usagemetric
	isSet bool
}

// Get returns the value.
func (v NullableUsagemetric) Get() *Usagemetric {
	return v.value
}

// Set modifies the value.
func (v *NullableUsagemetric) Set(val *Usagemetric) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableUsagemetric) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableUsagemetric) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUsagemetric returns a pointer to a new instance of NullableUsagemetric.
func NewNullableUsagemetric(val *Usagemetric) *NullableUsagemetric {
	return &NullableUsagemetric{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableUsagemetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableUsagemetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
