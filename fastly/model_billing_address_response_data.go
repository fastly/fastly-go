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

// BillingAddressResponseData struct for BillingAddressResponseData
type BillingAddressResponseData struct {
	// Alphanumeric string identifying the billing address.
	ID                   *string                   `json:"id,omitempty"`
	Attributes           *BillingAddressAttributes `json:"attributes,omitempty"`
	Type                 *TypeBillingAddress       `json:"type,omitempty"`
	Relationships        *RelationshipCustomer     `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingAddressResponseData BillingAddressResponseData

// NewBillingAddressResponseData instantiates a new BillingAddressResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAddressResponseData() *BillingAddressResponseData {
	this := BillingAddressResponseData{}
	var resourceType TypeBillingAddress = TYPEBILLINGADDRESS_BILLING_ADDRESS
	this.Type = &resourceType
	return &this
}

// NewBillingAddressResponseDataWithDefaults instantiates a new BillingAddressResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAddressResponseDataWithDefaults() *BillingAddressResponseData {
	this := BillingAddressResponseData{}
	var resourceType TypeBillingAddress = TYPEBILLINGADDRESS_BILLING_ADDRESS
	this.Type = &resourceType
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *BillingAddressResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *BillingAddressResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *BillingAddressResponseData) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BillingAddressResponseData) GetAttributes() BillingAddressAttributes {
	if o == nil || o.Attributes == nil {
		var ret BillingAddressAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressResponseData) GetAttributesOk() (*BillingAddressAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BillingAddressResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BillingAddressAttributes and assigns it to the Attributes field.
func (o *BillingAddressResponseData) SetAttributes(v BillingAddressAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingAddressResponseData) GetType() TypeBillingAddress {
	if o == nil || o.Type == nil {
		var ret TypeBillingAddress
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressResponseData) GetTypeOk() (*TypeBillingAddress, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingAddressResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeBillingAddress and assigns it to the Type field.
func (o *BillingAddressResponseData) SetType(v TypeBillingAddress) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BillingAddressResponseData) GetRelationships() RelationshipCustomer {
	if o == nil || o.Relationships == nil {
		var ret RelationshipCustomer
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddressResponseData) GetRelationshipsOk() (*RelationshipCustomer, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BillingAddressResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipCustomer and assigns it to the Relationships field.
func (o *BillingAddressResponseData) SetRelationships(v RelationshipCustomer) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingAddressResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BillingAddressResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varBillingAddressResponseData := _BillingAddressResponseData{}

	if err = json.Unmarshal(bytes, &varBillingAddressResponseData); err == nil {
		*o = BillingAddressResponseData(varBillingAddressResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "type")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingAddressResponseData is a helper abstraction for handling nullable billingaddressresponsedata types.
type NullableBillingAddressResponseData struct {
	value *BillingAddressResponseData
	isSet bool
}

// Get returns the value.
func (v NullableBillingAddressResponseData) Get() *BillingAddressResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingAddressResponseData) Set(val *BillingAddressResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingAddressResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingAddressResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingAddressResponseData returns a pointer to a new instance of NullableBillingAddressResponseData.
func NewNullableBillingAddressResponseData(val *BillingAddressResponseData) *NullableBillingAddressResponseData {
	return &NullableBillingAddressResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingAddressResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBillingAddressResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
