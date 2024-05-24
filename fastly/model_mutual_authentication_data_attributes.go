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

// MutualAuthenticationDataAttributes struct for MutualAuthenticationDataAttributes
type MutualAuthenticationDataAttributes struct {
	// One or more certificates. Enter each individual certificate blob on a new line. Must be PEM-formatted. Required on create. You may optionally rotate the cert_bundle on update.
	CertBundle *string `json:"cert_bundle,omitempty"`
	// Determines whether Mutual TLS will fail closed (enforced) or fail open. A true value will require a successful Mutual TLS handshake for the connection to continue and will fail closed if unsuccessful. A false value will fail open and allow the connection to proceed. Optional. Defaults to true.
	Enforced *bool `json:"enforced,omitempty"`
	// A custom name for your mutual authentication. Optional. If name is not supplied we will auto-generate one.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _MutualAuthenticationDataAttributes MutualAuthenticationDataAttributes

// NewMutualAuthenticationDataAttributes instantiates a new MutualAuthenticationDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualAuthenticationDataAttributes() *MutualAuthenticationDataAttributes {
	this := MutualAuthenticationDataAttributes{}
	return &this
}

// NewMutualAuthenticationDataAttributesWithDefaults instantiates a new MutualAuthenticationDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualAuthenticationDataAttributesWithDefaults() *MutualAuthenticationDataAttributes {
	this := MutualAuthenticationDataAttributes{}
	return &this
}

// GetCertBundle returns the CertBundle field value if set, zero value otherwise.
func (o *MutualAuthenticationDataAttributes) GetCertBundle() string {
	if o == nil || o.CertBundle == nil {
		var ret string
		return ret
	}
	return *o.CertBundle
}

// GetCertBundleOk returns a tuple with the CertBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationDataAttributes) GetCertBundleOk() (*string, bool) {
	if o == nil || o.CertBundle == nil {
		return nil, false
	}
	return o.CertBundle, true
}

// HasCertBundle returns a boolean if a field has been set.
func (o *MutualAuthenticationDataAttributes) HasCertBundle() bool {
	if o != nil && o.CertBundle != nil {
		return true
	}

	return false
}

// SetCertBundle gets a reference to the given string and assigns it to the CertBundle field.
func (o *MutualAuthenticationDataAttributes) SetCertBundle(v string) {
	o.CertBundle = &v
}

// GetEnforced returns the Enforced field value if set, zero value otherwise.
func (o *MutualAuthenticationDataAttributes) GetEnforced() bool {
	if o == nil || o.Enforced == nil {
		var ret bool
		return ret
	}
	return *o.Enforced
}

// GetEnforcedOk returns a tuple with the Enforced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationDataAttributes) GetEnforcedOk() (*bool, bool) {
	if o == nil || o.Enforced == nil {
		return nil, false
	}
	return o.Enforced, true
}

// HasEnforced returns a boolean if a field has been set.
func (o *MutualAuthenticationDataAttributes) HasEnforced() bool {
	if o != nil && o.Enforced != nil {
		return true
	}

	return false
}

// SetEnforced gets a reference to the given bool and assigns it to the Enforced field.
func (o *MutualAuthenticationDataAttributes) SetEnforced(v bool) {
	o.Enforced = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MutualAuthenticationDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MutualAuthenticationDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MutualAuthenticationDataAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o MutualAuthenticationDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CertBundle != nil {
		toSerialize["cert_bundle"] = o.CertBundle
	}
	if o.Enforced != nil {
		toSerialize["enforced"] = o.Enforced
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *MutualAuthenticationDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varMutualAuthenticationDataAttributes := _MutualAuthenticationDataAttributes{}

	if err = json.Unmarshal(bytes, &varMutualAuthenticationDataAttributes); err == nil {
		*o = MutualAuthenticationDataAttributes(varMutualAuthenticationDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cert_bundle")
		delete(additionalProperties, "enforced")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableMutualAuthenticationDataAttributes is a helper abstraction for handling nullable mutualauthenticationdataattributes types. 
type NullableMutualAuthenticationDataAttributes struct {
	value *MutualAuthenticationDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableMutualAuthenticationDataAttributes) Get() *MutualAuthenticationDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableMutualAuthenticationDataAttributes) Set(val *MutualAuthenticationDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableMutualAuthenticationDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableMutualAuthenticationDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableMutualAuthenticationDataAttributes returns a pointer to a new instance of NullableMutualAuthenticationDataAttributes.
func NewNullableMutualAuthenticationDataAttributes(val *MutualAuthenticationDataAttributes) *NullableMutualAuthenticationDataAttributes {
	return &NullableMutualAuthenticationDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableMutualAuthenticationDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableMutualAuthenticationDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
