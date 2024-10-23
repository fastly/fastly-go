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

// CustomerAddress struct for CustomerAddress
type CustomerAddress struct {
	// The type of the address.
	Type *string `json:"type,omitempty"`
	// The street number and name of the address.
	Address1 *string `json:"address_1,omitempty"`
	// Additional address line for unit number, P.O. Box, etc.
	Address2 *string `json:"address_2,omitempty"`
	// City, town, or locality name the address is located.
	Locality *string `json:"locality,omitempty"`
	// State, province, or region of the address.
	Region *string `json:"region,omitempty"`
	// ISO 3166-1 alpha-2 country code (e.g., \"US\", \"CA\", \"NZ\")
	Country *string `json:"country,omitempty"`
	// Postal or Zip code of the address.
	PostalCode           *string `json:"postal_code,omitempty"`
	AdditionalProperties map[string]any
}

type _CustomerAddress CustomerAddress

// NewCustomerAddress instantiates a new CustomerAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddress() *CustomerAddress {
	this := CustomerAddress{}
	return &this
}

// NewCustomerAddressWithDefaults instantiates a new CustomerAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressWithDefaults() *CustomerAddress {
	this := CustomerAddress{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomerAddress) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomerAddress) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomerAddress) SetType(v string) {
	o.Type = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *CustomerAddress) GetAddress1() string {
	if o == nil || o.Address1 == nil {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetAddress1Ok() (*string, bool) {
	if o == nil || o.Address1 == nil {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *CustomerAddress) HasAddress1() bool {
	if o != nil && o.Address1 != nil {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *CustomerAddress) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *CustomerAddress) GetAddress2() string {
	if o == nil || o.Address2 == nil {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetAddress2Ok() (*string, bool) {
	if o == nil || o.Address2 == nil {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *CustomerAddress) HasAddress2() bool {
	if o != nil && o.Address2 != nil {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *CustomerAddress) SetAddress2(v string) {
	o.Address2 = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *CustomerAddress) GetLocality() string {
	if o == nil || o.Locality == nil {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetLocalityOk() (*string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *CustomerAddress) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *CustomerAddress) SetLocality(v string) {
	o.Locality = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CustomerAddress) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CustomerAddress) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CustomerAddress) SetRegion(v string) {
	o.Region = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerAddress) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerAddress) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerAddress) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *CustomerAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CustomerAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *CustomerAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o CustomerAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Address1 != nil {
		toSerialize["address_1"] = o.Address1
	}
	if o.Address2 != nil {
		toSerialize["address_2"] = o.Address2
	}
	if o.Locality != nil {
		toSerialize["locality"] = o.Locality
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *CustomerAddress) UnmarshalJSON(bytes []byte) (err error) {
	varCustomerAddress := _CustomerAddress{}

	if err = json.Unmarshal(bytes, &varCustomerAddress); err == nil {
		*o = CustomerAddress(varCustomerAddress)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "address_1")
		delete(additionalProperties, "address_2")
		delete(additionalProperties, "locality")
		delete(additionalProperties, "region")
		delete(additionalProperties, "country")
		delete(additionalProperties, "postal_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableCustomerAddress is a helper abstraction for handling nullable customeraddress types.
type NullableCustomerAddress struct {
	value *CustomerAddress
	isSet bool
}

// Get returns the value.
func (v NullableCustomerAddress) Get() *CustomerAddress {
	return v.value
}

// Set modifies the value.
func (v *NullableCustomerAddress) Set(val *CustomerAddress) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableCustomerAddress) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableCustomerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCustomerAddress returns a pointer to a new instance of NullableCustomerAddress.
func NewNullableCustomerAddress(val *CustomerAddress) *NullableCustomerAddress {
	return &NullableCustomerAddress{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableCustomerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableCustomerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
