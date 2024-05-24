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

// ServiceInvitationData struct for ServiceInvitationData
type ServiceInvitationData struct {
	Type *TypeServiceInvitation `json:"type,omitempty"`
	Attributes *ServiceInvitationDataAttributes `json:"attributes,omitempty"`
	Relationships *ServiceInvitationDataRelationships `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceInvitationData ServiceInvitationData

// NewServiceInvitationData instantiates a new ServiceInvitationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceInvitationData() *ServiceInvitationData {
	this := ServiceInvitationData{}
	var resourceType TypeServiceInvitation = TYPESERVICEINVITATION_SERVICE_INVITATION
	this.Type = &resourceType
	return &this
}

// NewServiceInvitationDataWithDefaults instantiates a new ServiceInvitationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceInvitationDataWithDefaults() *ServiceInvitationData {
	this := ServiceInvitationData{}
	var resourceType TypeServiceInvitation = TYPESERVICEINVITATION_SERVICE_INVITATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceInvitationData) GetType() TypeServiceInvitation {
	if o == nil || o.Type == nil {
		var ret TypeServiceInvitation
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInvitationData) GetTypeOk() (*TypeServiceInvitation, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceInvitationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeServiceInvitation and assigns it to the Type field.
func (o *ServiceInvitationData) SetType(v TypeServiceInvitation) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ServiceInvitationData) GetAttributes() ServiceInvitationDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ServiceInvitationDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInvitationData) GetAttributesOk() (*ServiceInvitationDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ServiceInvitationData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ServiceInvitationDataAttributes and assigns it to the Attributes field.
func (o *ServiceInvitationData) SetAttributes(v ServiceInvitationDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ServiceInvitationData) GetRelationships() ServiceInvitationDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ServiceInvitationDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInvitationData) GetRelationshipsOk() (*ServiceInvitationDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ServiceInvitationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ServiceInvitationDataRelationships and assigns it to the Relationships field.
func (o *ServiceInvitationData) SetRelationships(v ServiceInvitationDataRelationships) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceInvitationData) MarshalJSON() ([]byte, error) {
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
func (o *ServiceInvitationData) UnmarshalJSON(bytes []byte) (err error) {
	varServiceInvitationData := _ServiceInvitationData{}

	if err = json.Unmarshal(bytes, &varServiceInvitationData); err == nil {
		*o = ServiceInvitationData(varServiceInvitationData)
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

// NullableServiceInvitationData is a helper abstraction for handling nullable serviceinvitationdata types. 
type NullableServiceInvitationData struct {
	value *ServiceInvitationData
	isSet bool
}

// Get returns the value.
func (v NullableServiceInvitationData) Get() *ServiceInvitationData {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceInvitationData) Set(val *ServiceInvitationData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceInvitationData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceInvitationData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceInvitationData returns a pointer to a new instance of NullableServiceInvitationData.
func NewNullableServiceInvitationData(val *ServiceInvitationData) *NullableServiceInvitationData {
	return &NullableServiceInvitationData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceInvitationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceInvitationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
