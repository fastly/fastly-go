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

// InvitationResponseData struct for InvitationResponseData
type InvitationResponseData struct {
	Type *TypeInvitation `json:"type,omitempty"`
	Attributes *Timestamps `json:"attributes,omitempty"`
	Relationships *RelationshipsForInvitation `json:"relationships,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _InvitationResponseData InvitationResponseData

// NewInvitationResponseData instantiates a new InvitationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationResponseData() *InvitationResponseData {
	this := InvitationResponseData{}
	var resourceType TypeInvitation = TYPEINVITATION_INVITATION
	this.Type = &resourceType
	return &this
}

// NewInvitationResponseDataWithDefaults instantiates a new InvitationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationResponseDataWithDefaults() *InvitationResponseData {
	this := InvitationResponseData{}
	var resourceType TypeInvitation = TYPEINVITATION_INVITATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InvitationResponseData) GetType() TypeInvitation {
	if o == nil || o.Type == nil {
		var ret TypeInvitation
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationResponseData) GetTypeOk() (*TypeInvitation, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InvitationResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeInvitation and assigns it to the Type field.
func (o *InvitationResponseData) SetType(v TypeInvitation) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InvitationResponseData) GetAttributes() Timestamps {
	if o == nil || o.Attributes == nil {
		var ret Timestamps
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationResponseData) GetAttributesOk() (*Timestamps, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InvitationResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Timestamps and assigns it to the Attributes field.
func (o *InvitationResponseData) SetAttributes(v Timestamps) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InvitationResponseData) GetRelationships() RelationshipsForInvitation {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForInvitation
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationResponseData) GetRelationshipsOk() (*RelationshipsForInvitation, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InvitationResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForInvitation and assigns it to the Relationships field.
func (o *InvitationResponseData) SetRelationships(v RelationshipsForInvitation) {
	o.Relationships = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *InvitationResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *InvitationResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *InvitationResponseData) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InvitationResponseData) MarshalJSON() ([]byte, error) {
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
func (o *InvitationResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varInvitationResponseData := _InvitationResponseData{}

	if err = json.Unmarshal(bytes, &varInvitationResponseData); err == nil {
		*o = InvitationResponseData(varInvitationResponseData)
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

// NullableInvitationResponseData is a helper abstraction for handling nullable invitationresponsedata types. 
type NullableInvitationResponseData struct {
	value *InvitationResponseData
	isSet bool
}

// Get returns the value.
func (v NullableInvitationResponseData) Get() *InvitationResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableInvitationResponseData) Set(val *InvitationResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInvitationResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInvitationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInvitationResponseData returns a pointer to a new instance of NullableInvitationResponseData.
func NewNullableInvitationResponseData(val *InvitationResponseData) *NullableInvitationResponseData {
	return &NullableInvitationResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInvitationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableInvitationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
