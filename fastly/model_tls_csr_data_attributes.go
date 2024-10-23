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

// TLSCsrDataAttributes struct for TLSCsrDataAttributes
type TLSCsrDataAttributes struct {
	// Subject Alternate Names - An array of one or more fully qualified domain names or public IP addresses to be secured by this certificate. Required.
	Sans []string `json:"sans"`
	// Common Name (CN) - The fully qualified domain name (FQDN) to be secured by this certificate. The common name should be one of the entries in the SANs parameter.
	CommonName *string `json:"common_name,omitempty"`
	// Country (C) - The two-letter ISO country code where the organization is located.
	Country *string `json:"country,omitempty"`
	// State (S) - The state, province, region, or county where the organization is located. This should not be abbreviated.
	State *string `json:"state,omitempty"`
	// Locality (L) - The locality, city, town, or village where the organization is located.
	City *string `json:"city,omitempty"`
	// Postal Code - The postal code where the organization is located.
	PostalCode *string `json:"postal_code,omitempty"`
	// Street Address - The street address where the organization is located.
	StreetAddress *string `json:"street_address,omitempty"`
	// Organization (O) - The legal name of the organization, including any suffixes. This should not be abbreviated.
	Organization *string `json:"organization,omitempty"`
	// Organizational Unit (OU) - The internal division of the organization managing the certificate.
	OrganizationalUnit *string `json:"organizational_unit,omitempty"`
	// Email Address (EMAIL) - The organizational contact for this.
	Email *string `json:"email,omitempty"`
	// CSR Key Type.
	KeyType              *string `json:"key_type,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCsrDataAttributes TLSCsrDataAttributes

// NewTLSCsrDataAttributes instantiates a new TLSCsrDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCsrDataAttributes(sans []string) *TLSCsrDataAttributes {
	this := TLSCsrDataAttributes{}
	this.Sans = sans
	return &this
}

// NewTLSCsrDataAttributesWithDefaults instantiates a new TLSCsrDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCsrDataAttributesWithDefaults() *TLSCsrDataAttributes {
	this := TLSCsrDataAttributes{}
	return &this
}

// GetSans returns the Sans field value
func (o *TLSCsrDataAttributes) GetSans() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Sans
}

// GetSansOk returns a tuple with the Sans field value
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetSansOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sans, true
}

// SetSans sets field value
func (o *TLSCsrDataAttributes) SetSans(v []string) {
	o.Sans = v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *TLSCsrDataAttributes) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *TLSCsrDataAttributes) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *TLSCsrDataAttributes) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *TLSCsrDataAttributes) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *TLSCsrDataAttributes) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *TLSCsrDataAttributes) SetCountry(v string) {
	o.Country = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TLSCsrDataAttributes) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TLSCsrDataAttributes) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TLSCsrDataAttributes) SetState(v string) {
	o.State = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *TLSCsrDataAttributes) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *TLSCsrDataAttributes) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *TLSCsrDataAttributes) SetCity(v string) {
	o.City = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *TLSCsrDataAttributes) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *TLSCsrDataAttributes) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *TLSCsrDataAttributes) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *TLSCsrDataAttributes) GetStreetAddress() string {
	if o == nil || o.StreetAddress == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetStreetAddressOk() (*string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *TLSCsrDataAttributes) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *TLSCsrDataAttributes) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *TLSCsrDataAttributes) GetOrganization() string {
	if o == nil || o.Organization == nil {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetOrganizationOk() (*string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *TLSCsrDataAttributes) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *TLSCsrDataAttributes) SetOrganization(v string) {
	o.Organization = &v
}

// GetOrganizationalUnit returns the OrganizationalUnit field value if set, zero value otherwise.
func (o *TLSCsrDataAttributes) GetOrganizationalUnit() string {
	if o == nil || o.OrganizationalUnit == nil {
		var ret string
		return ret
	}
	return *o.OrganizationalUnit
}

// GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetOrganizationalUnitOk() (*string, bool) {
	if o == nil || o.OrganizationalUnit == nil {
		return nil, false
	}
	return o.OrganizationalUnit, true
}

// HasOrganizationalUnit returns a boolean if a field has been set.
func (o *TLSCsrDataAttributes) HasOrganizationalUnit() bool {
	if o != nil && o.OrganizationalUnit != nil {
		return true
	}

	return false
}

// SetOrganizationalUnit gets a reference to the given string and assigns it to the OrganizationalUnit field.
func (o *TLSCsrDataAttributes) SetOrganizationalUnit(v string) {
	o.OrganizationalUnit = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *TLSCsrDataAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *TLSCsrDataAttributes) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *TLSCsrDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *TLSCsrDataAttributes) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrDataAttributes) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *TLSCsrDataAttributes) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *TLSCsrDataAttributes) SetKeyType(v string) {
	o.KeyType = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCsrDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["sans"] = o.Sans
	}
	if o.CommonName != nil {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.StreetAddress != nil {
		toSerialize["street_address"] = o.StreetAddress
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.OrganizationalUnit != nil {
		toSerialize["organizational_unit"] = o.OrganizationalUnit
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.KeyType != nil {
		toSerialize["key_type"] = o.KeyType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TLSCsrDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCsrDataAttributes := _TLSCsrDataAttributes{}

	if err = json.Unmarshal(bytes, &varTLSCsrDataAttributes); err == nil {
		*o = TLSCsrDataAttributes(varTLSCsrDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sans")
		delete(additionalProperties, "common_name")
		delete(additionalProperties, "country")
		delete(additionalProperties, "state")
		delete(additionalProperties, "city")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "street_address")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "organizational_unit")
		delete(additionalProperties, "email")
		delete(additionalProperties, "key_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSCsrDataAttributes is a helper abstraction for handling nullable tlscsrdataattributes types.
type NullableTLSCsrDataAttributes struct {
	value *TLSCsrDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTLSCsrDataAttributes) Get() *TLSCsrDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCsrDataAttributes) Set(val *TLSCsrDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCsrDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCsrDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCsrDataAttributes returns a pointer to a new instance of NullableTLSCsrDataAttributes.
func NewNullableTLSCsrDataAttributes(val *TLSCsrDataAttributes) *NullableTLSCsrDataAttributes {
	return &NullableTLSCsrDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCsrDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSCsrDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
