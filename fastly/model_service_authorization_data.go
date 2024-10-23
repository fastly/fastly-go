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

// ServiceAuthorizationData struct for ServiceAuthorizationData
type ServiceAuthorizationData struct {
	Type                 *TypeServiceAuthorization              `json:"type,omitempty"`
	Attributes           *ServiceAuthorizationDataAttributes    `json:"attributes,omitempty"`
	Relationships        *ServiceAuthorizationDataRelationships `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceAuthorizationData ServiceAuthorizationData

// NewServiceAuthorizationData instantiates a new ServiceAuthorizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAuthorizationData() *ServiceAuthorizationData {
	this := ServiceAuthorizationData{}
	var resourceType TypeServiceAuthorization = TYPESERVICEAUTHORIZATION_SERVICE_AUTHORIZATION
	this.Type = &resourceType
	return &this
}

// NewServiceAuthorizationDataWithDefaults instantiates a new ServiceAuthorizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAuthorizationDataWithDefaults() *ServiceAuthorizationData {
	this := ServiceAuthorizationData{}
	var resourceType TypeServiceAuthorization = TYPESERVICEAUTHORIZATION_SERVICE_AUTHORIZATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceAuthorizationData) GetType() TypeServiceAuthorization {
	if o == nil || o.Type == nil {
		var ret TypeServiceAuthorization
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationData) GetTypeOk() (*TypeServiceAuthorization, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceAuthorizationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeServiceAuthorization and assigns it to the Type field.
func (o *ServiceAuthorizationData) SetType(v TypeServiceAuthorization) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ServiceAuthorizationData) GetAttributes() ServiceAuthorizationDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ServiceAuthorizationDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationData) GetAttributesOk() (*ServiceAuthorizationDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ServiceAuthorizationData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ServiceAuthorizationDataAttributes and assigns it to the Attributes field.
func (o *ServiceAuthorizationData) SetAttributes(v ServiceAuthorizationDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ServiceAuthorizationData) GetRelationships() ServiceAuthorizationDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ServiceAuthorizationDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationData) GetRelationshipsOk() (*ServiceAuthorizationDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ServiceAuthorizationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ServiceAuthorizationDataRelationships and assigns it to the Relationships field.
func (o *ServiceAuthorizationData) SetRelationships(v ServiceAuthorizationDataRelationships) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceAuthorizationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
func (o *ServiceAuthorizationData) UnmarshalJSON(bytes []byte) (err error) {
	varServiceAuthorizationData := _ServiceAuthorizationData{}

	if err = json.Unmarshal(bytes, &varServiceAuthorizationData); err == nil {
		*o = ServiceAuthorizationData(varServiceAuthorizationData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceAuthorizationData is a helper abstraction for handling nullable serviceauthorizationdata types.
type NullableServiceAuthorizationData struct {
	value *ServiceAuthorizationData
	isSet bool
}

// Get returns the value.
func (v NullableServiceAuthorizationData) Get() *ServiceAuthorizationData {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceAuthorizationData) Set(val *ServiceAuthorizationData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceAuthorizationData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceAuthorizationData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceAuthorizationData returns a pointer to a new instance of NullableServiceAuthorizationData.
func NewNullableServiceAuthorizationData(val *ServiceAuthorizationData) *NullableServiceAuthorizationData {
	return &NullableServiceAuthorizationData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceAuthorizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableServiceAuthorizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
