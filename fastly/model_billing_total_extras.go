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

// BillingTotalExtras struct for BillingTotalExtras
type BillingTotalExtras struct {
	// The name of this extra cost.
	Name *string `json:"name,omitempty"`
	// Recurring monthly cost in USD. Not present if $0.0.
	Recurring *float32 `json:"recurring,omitempty"`
	// Initial set up cost in USD. Not present if $0.0 or this is not the month the extra was added.
	Setup *float32 `json:"setup,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingTotalExtras BillingTotalExtras

// NewBillingTotalExtras instantiates a new BillingTotalExtras object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingTotalExtras() *BillingTotalExtras {
	this := BillingTotalExtras{}
	return &this
}

// NewBillingTotalExtrasWithDefaults instantiates a new BillingTotalExtras object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingTotalExtrasWithDefaults() *BillingTotalExtras {
	this := BillingTotalExtras{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingTotalExtras) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotalExtras) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillingTotalExtras) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingTotalExtras) SetName(v string) {
	o.Name = &v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *BillingTotalExtras) GetRecurring() float32 {
	if o == nil || o.Recurring == nil {
		var ret float32
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotalExtras) GetRecurringOk() (*float32, bool) {
	if o == nil || o.Recurring == nil {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *BillingTotalExtras) HasRecurring() bool {
	if o != nil && o.Recurring != nil {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given float32 and assigns it to the Recurring field.
func (o *BillingTotalExtras) SetRecurring(v float32) {
	o.Recurring = &v
}

// GetSetup returns the Setup field value if set, zero value otherwise.
func (o *BillingTotalExtras) GetSetup() float32 {
	if o == nil || o.Setup == nil {
		var ret float32
		return ret
	}
	return *o.Setup
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTotalExtras) GetSetupOk() (*float32, bool) {
	if o == nil || o.Setup == nil {
		return nil, false
	}
	return o.Setup, true
}

// HasSetup returns a boolean if a field has been set.
func (o *BillingTotalExtras) HasSetup() bool {
	if o != nil && o.Setup != nil {
		return true
	}

	return false
}

// SetSetup gets a reference to the given float32 and assigns it to the Setup field.
func (o *BillingTotalExtras) SetSetup(v float32) {
	o.Setup = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingTotalExtras) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Recurring != nil {
		toSerialize["recurring"] = o.Recurring
	}
	if o.Setup != nil {
		toSerialize["setup"] = o.Setup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingTotalExtras) UnmarshalJSON(bytes []byte) (err error) {
	varBillingTotalExtras := _BillingTotalExtras{}

	if err = json.Unmarshal(bytes, &varBillingTotalExtras); err == nil {
		*o = BillingTotalExtras(varBillingTotalExtras)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "recurring")
		delete(additionalProperties, "setup")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingTotalExtras is a helper abstraction for handling nullable billingtotalextras types. 
type NullableBillingTotalExtras struct {
	value *BillingTotalExtras
	isSet bool
}

// Get returns the value.
func (v NullableBillingTotalExtras) Get() *BillingTotalExtras {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingTotalExtras) Set(val *BillingTotalExtras) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingTotalExtras) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingTotalExtras) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingTotalExtras returns a pointer to a new instance of NullableBillingTotalExtras.
func NewNullableBillingTotalExtras(val *BillingTotalExtras) *NullableBillingTotalExtras {
	return &NullableBillingTotalExtras{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingTotalExtras) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingTotalExtras) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
