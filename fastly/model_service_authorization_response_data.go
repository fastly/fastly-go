// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// ServiceAuthorizationResponseData struct for ServiceAuthorizationResponseData
type ServiceAuthorizationResponseData struct {
	Type *TypeServiceAuthorization `json:"type,omitempty"`
	Attributes *Timestamps `json:"attributes,omitempty"`
	Relationships *ServiceAuthorizationDataRelationships `json:"relationships,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceAuthorizationResponseData ServiceAuthorizationResponseData

// NewServiceAuthorizationResponseData instantiates a new ServiceAuthorizationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAuthorizationResponseData() *ServiceAuthorizationResponseData {
	this := ServiceAuthorizationResponseData{}
	var resourceType TypeServiceAuthorization = TYPESERVICEAUTHORIZATION_SERVICE_AUTHORIZATION
	this.Type = &resourceType
	return &this
}

// NewServiceAuthorizationResponseDataWithDefaults instantiates a new ServiceAuthorizationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAuthorizationResponseDataWithDefaults() *ServiceAuthorizationResponseData {
	this := ServiceAuthorizationResponseData{}
	var resourceType TypeServiceAuthorization = TYPESERVICEAUTHORIZATION_SERVICE_AUTHORIZATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceAuthorizationResponseData) GetType() TypeServiceAuthorization {
	if o == nil || o.Type == nil {
		var ret TypeServiceAuthorization
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationResponseData) GetTypeOk() (*TypeServiceAuthorization, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceAuthorizationResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeServiceAuthorization and assigns it to the Type field.
func (o *ServiceAuthorizationResponseData) SetType(v TypeServiceAuthorization) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ServiceAuthorizationResponseData) GetAttributes() Timestamps {
	if o == nil || o.Attributes == nil {
		var ret Timestamps
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationResponseData) GetAttributesOk() (*Timestamps, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ServiceAuthorizationResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Timestamps and assigns it to the Attributes field.
func (o *ServiceAuthorizationResponseData) SetAttributes(v Timestamps) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ServiceAuthorizationResponseData) GetRelationships() ServiceAuthorizationDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ServiceAuthorizationDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationResponseData) GetRelationshipsOk() (*ServiceAuthorizationDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ServiceAuthorizationResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ServiceAuthorizationDataRelationships and assigns it to the Relationships field.
func (o *ServiceAuthorizationResponseData) SetRelationships(v ServiceAuthorizationDataRelationships) {
	o.Relationships = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ServiceAuthorizationResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ServiceAuthorizationResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ServiceAuthorizationResponseData) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceAuthorizationResponseData) MarshalJSON() ([]byte, error) {
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
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServiceAuthorizationResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varServiceAuthorizationResponseData := _ServiceAuthorizationResponseData{}

	if err = json.Unmarshal(bytes, &varServiceAuthorizationResponseData); err == nil {
		*o = ServiceAuthorizationResponseData(varServiceAuthorizationResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceAuthorizationResponseData is a helper abstraction for handling nullable serviceauthorizationresponsedata types. 
type NullableServiceAuthorizationResponseData struct {
	value *ServiceAuthorizationResponseData
	isSet bool
}

// Get returns the value.
func (v NullableServiceAuthorizationResponseData) Get() *ServiceAuthorizationResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceAuthorizationResponseData) Set(val *ServiceAuthorizationResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceAuthorizationResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceAuthorizationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceAuthorizationResponseData returns a pointer to a new instance of NullableServiceAuthorizationResponseData.
func NewNullableServiceAuthorizationResponseData(val *ServiceAuthorizationResponseData) *NullableServiceAuthorizationResponseData {
	return &NullableServiceAuthorizationResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceAuthorizationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceAuthorizationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
