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

// BillingAddressAttributes struct for BillingAddressAttributes
type BillingAddressAttributes struct {
	// The first address line.
	Address1 *string `json:"address_1,omitempty"`
	// The second address line.
	Address2 NullableString `json:"address_2,omitempty"`
	// The city name.
	City NullableString `json:"city,omitempty"`
	// ISO 3166-1 two-letter country code.
	Country *string `json:"country,omitempty"`
	// Other locality.
	Locality NullableString `json:"locality,omitempty"`
	// Postal code (ZIP code for US addresses).
	PostalCode *string `json:"postal_code,omitempty"`
	// The state or province name.
	State NullableString `json:"state,omitempty"`
	CustomerID *string `json:"customer_id,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingAddressAttributes BillingAddressAttributes

// NewBillingAddressAttributes instantiates a new BillingAddressAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAddressAttributes() *BillingAddressAttributes {
	this := BillingAddressAttributes{}
	return &this
}

// NewBillingAddressAttributesWithDefaults instantiates a new BillingAddressAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAddressAttributesWithDefaults() *BillingAddressAttributes {
	this := BillingAddressAttributes{}
	return &this
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *BillingAddressAttributes) GetAddress1() string {
	if o == nil || o.Address1 == nil {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressAttributes) GetAddress1Ok() (*string, bool) {
	if o == nil || o.Address1 == nil {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *BillingAddressAttributes) HasAddress1() bool {
	if o != nil && o.Address1 != nil {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *BillingAddressAttributes) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingAddressAttributes) GetAddress2() string {
	if o == nil || o.Address2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Address2.Get()
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingAddressAttributes) GetAddress2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address2.Get(), o.Address2.IsSet()
}

// HasAddress2 returns a boolean if a field has been set.
func (o *BillingAddressAttributes) HasAddress2() bool {
	if o != nil && o.Address2.IsSet() {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given NullableString and assigns it to the Address2 field.
func (o *BillingAddressAttributes) SetAddress2(v string) {
	o.Address2.Set(&v)
}
// SetAddress2Nil sets the value for Address2 to be an explicit nil
func (o *BillingAddressAttributes) SetAddress2Nil() {
	o.Address2.Set(nil)
}

// UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
func (o *BillingAddressAttributes) UnsetAddress2() {
	o.Address2.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingAddressAttributes) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingAddressAttributes) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *BillingAddressAttributes) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *BillingAddressAttributes) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *BillingAddressAttributes) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *BillingAddressAttributes) UnsetCity() {
	o.City.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *BillingAddressAttributes) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressAttributes) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *BillingAddressAttributes) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *BillingAddressAttributes) SetCountry(v string) {
	o.Country = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingAddressAttributes) GetLocality() string {
	if o == nil || o.Locality.Get() == nil {
		var ret string
		return ret
	}
	return *o.Locality.Get()
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingAddressAttributes) GetLocalityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Locality.Get(), o.Locality.IsSet()
}

// HasLocality returns a boolean if a field has been set.
func (o *BillingAddressAttributes) HasLocality() bool {
	if o != nil && o.Locality.IsSet() {
		return true
	}

	return false
}

// SetLocality gets a reference to the given NullableString and assigns it to the Locality field.
func (o *BillingAddressAttributes) SetLocality(v string) {
	o.Locality.Set(&v)
}
// SetLocalityNil sets the value for Locality to be an explicit nil
func (o *BillingAddressAttributes) SetLocalityNil() {
	o.Locality.Set(nil)
}

// UnsetLocality ensures that no value is present for Locality, not even an explicit nil
func (o *BillingAddressAttributes) UnsetLocality() {
	o.Locality.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *BillingAddressAttributes) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressAttributes) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *BillingAddressAttributes) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *BillingAddressAttributes) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingAddressAttributes) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingAddressAttributes) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *BillingAddressAttributes) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *BillingAddressAttributes) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *BillingAddressAttributes) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *BillingAddressAttributes) UnsetState() {
	o.State.Unset()
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *BillingAddressAttributes) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressAttributes) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *BillingAddressAttributes) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *BillingAddressAttributes) SetCustomerID(v string) {
	o.CustomerID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingAddressAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Address1 != nil {
		toSerialize["address_1"] = o.Address1
	}
	if o.Address2.IsSet() {
		toSerialize["address_2"] = o.Address2.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Locality.IsSet() {
		toSerialize["locality"] = o.Locality.Get()
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingAddressAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varBillingAddressAttributes := _BillingAddressAttributes{}

	if err = json.Unmarshal(bytes, &varBillingAddressAttributes); err == nil {
		*o = BillingAddressAttributes(varBillingAddressAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address_1")
		delete(additionalProperties, "address_2")
		delete(additionalProperties, "city")
		delete(additionalProperties, "country")
		delete(additionalProperties, "locality")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "state")
		delete(additionalProperties, "customer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingAddressAttributes is a helper abstraction for handling nullable billingaddressattributes types. 
type NullableBillingAddressAttributes struct {
	value *BillingAddressAttributes
	isSet bool
}

// Get returns the value.
func (v NullableBillingAddressAttributes) Get() *BillingAddressAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingAddressAttributes) Set(val *BillingAddressAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingAddressAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingAddressAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingAddressAttributes returns a pointer to a new instance of NullableBillingAddressAttributes.
func NewNullableBillingAddressAttributes(val *BillingAddressAttributes) *NullableBillingAddressAttributes {
	return &NullableBillingAddressAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingAddressAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingAddressAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
