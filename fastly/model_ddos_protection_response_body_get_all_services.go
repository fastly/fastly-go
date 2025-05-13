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

// DdosProtectionResponseBodyGetAllServices struct for DdosProtectionResponseBodyGetAllServices
type DdosProtectionResponseBodyGetAllServices struct {
	Product  *DdosProtectionResponseProductProduct  `json:"product,omitempty"`
	Customer *BotManagementResponseCustomerCustomer `json:"customer,omitempty"`
	// A list of services with DDoS Protection enabled.
	Services             []string                                        `json:"services,omitempty"`
	Links                *DdosProtectionResponseLinksGetAllServicesLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionResponseBodyGetAllServices DdosProtectionResponseBodyGetAllServices

// NewDdosProtectionResponseBodyGetAllServices instantiates a new DdosProtectionResponseBodyGetAllServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionResponseBodyGetAllServices() *DdosProtectionResponseBodyGetAllServices {
	this := DdosProtectionResponseBodyGetAllServices{}
	return &this
}

// NewDdosProtectionResponseBodyGetAllServicesWithDefaults instantiates a new DdosProtectionResponseBodyGetAllServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionResponseBodyGetAllServicesWithDefaults() *DdosProtectionResponseBodyGetAllServices {
	this := DdosProtectionResponseBodyGetAllServices{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *DdosProtectionResponseBodyGetAllServices) GetProduct() DdosProtectionResponseProductProduct {
	if o == nil || o.Product == nil {
		var ret DdosProtectionResponseProductProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseBodyGetAllServices) GetProductOk() (*DdosProtectionResponseProductProduct, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *DdosProtectionResponseBodyGetAllServices) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given DdosProtectionResponseProductProduct and assigns it to the Product field.
func (o *DdosProtectionResponseBodyGetAllServices) SetProduct(v DdosProtectionResponseProductProduct) {
	o.Product = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *DdosProtectionResponseBodyGetAllServices) GetCustomer() BotManagementResponseCustomerCustomer {
	if o == nil || o.Customer == nil {
		var ret BotManagementResponseCustomerCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseBodyGetAllServices) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *DdosProtectionResponseBodyGetAllServices) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given BotManagementResponseCustomerCustomer and assigns it to the Customer field.
func (o *DdosProtectionResponseBodyGetAllServices) SetCustomer(v BotManagementResponseCustomerCustomer) {
	o.Customer = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *DdosProtectionResponseBodyGetAllServices) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseBodyGetAllServices) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *DdosProtectionResponseBodyGetAllServices) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *DdosProtectionResponseBodyGetAllServices) SetServices(v []string) {
	o.Services = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DdosProtectionResponseBodyGetAllServices) GetLinks() DdosProtectionResponseLinksGetAllServicesLinks {
	if o == nil || o.Links == nil {
		var ret DdosProtectionResponseLinksGetAllServicesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseBodyGetAllServices) GetLinksOk() (*DdosProtectionResponseLinksGetAllServicesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DdosProtectionResponseBodyGetAllServices) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DdosProtectionResponseLinksGetAllServicesLinks and assigns it to the Links field.
func (o *DdosProtectionResponseBodyGetAllServices) SetLinks(v DdosProtectionResponseLinksGetAllServicesLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionResponseBodyGetAllServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Product != nil {
		toSerialize["product"] = o.Product
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionResponseBodyGetAllServices) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionResponseBodyGetAllServices := _DdosProtectionResponseBodyGetAllServices{}

	if err = json.Unmarshal(bytes, &varDdosProtectionResponseBodyGetAllServices); err == nil {
		*o = DdosProtectionResponseBodyGetAllServices(varDdosProtectionResponseBodyGetAllServices)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		delete(additionalProperties, "customer")
		delete(additionalProperties, "services")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionResponseBodyGetAllServices is a helper abstraction for handling nullable ddosprotectionresponsebodygetallservices types.
type NullableDdosProtectionResponseBodyGetAllServices struct {
	value *DdosProtectionResponseBodyGetAllServices
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionResponseBodyGetAllServices) Get() *DdosProtectionResponseBodyGetAllServices {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionResponseBodyGetAllServices) Set(val *DdosProtectionResponseBodyGetAllServices) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionResponseBodyGetAllServices) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionResponseBodyGetAllServices) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionResponseBodyGetAllServices returns a pointer to a new instance of NullableDdosProtectionResponseBodyGetAllServices.
func NewNullableDdosProtectionResponseBodyGetAllServices(val *DdosProtectionResponseBodyGetAllServices) *NullableDdosProtectionResponseBodyGetAllServices {
	return &NullableDdosProtectionResponseBodyGetAllServices{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionResponseBodyGetAllServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionResponseBodyGetAllServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
