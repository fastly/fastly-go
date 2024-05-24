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

// ServiceCreate struct for ServiceCreate
type ServiceCreate struct {
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// The name of the service.
	Name *string `json:"name,omitempty"`
	// Alphanumeric string identifying the customer.
	CustomerID *string `json:"customer_id,omitempty"`
	// The type of this service.
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceCreate ServiceCreate

// NewServiceCreate instantiates a new ServiceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceCreate() *ServiceCreate {
	this := ServiceCreate{}
	return &this
}

// NewServiceCreateWithDefaults instantiates a new ServiceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceCreateWithDefaults() *ServiceCreate {
	this := ServiceCreate{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceCreate) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceCreate) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ServiceCreate) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ServiceCreate) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ServiceCreate) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ServiceCreate) UnsetComment() {
	o.Comment.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceCreate) SetName(v string) {
	o.Name = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *ServiceCreate) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreate) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *ServiceCreate) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *ServiceCreate) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceCreate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceCreate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceCreate) SetType(v string) {
	o.Type = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServiceCreate) UnmarshalJSON(bytes []byte) (err error) {
	varServiceCreate := _ServiceCreate{}

	if err = json.Unmarshal(bytes, &varServiceCreate); err == nil {
		*o = ServiceCreate(varServiceCreate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "name")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceCreate is a helper abstraction for handling nullable servicecreate types. 
type NullableServiceCreate struct {
	value *ServiceCreate
	isSet bool
}

// Get returns the value.
func (v NullableServiceCreate) Get() *ServiceCreate {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceCreate) Set(val *ServiceCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceCreate returns a pointer to a new instance of NullableServiceCreate.
func NewNullableServiceCreate(val *ServiceCreate) *NullableServiceCreate {
	return &NullableServiceCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
