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

// Service struct for Service
type Service struct {
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// The name of the service.
	Name *string `json:"name,omitempty"`
	// Alphanumeric string identifying the customer.
	CustomerID *string `json:"customer_id,omitempty"`
	AdditionalProperties map[string]any
}

type _Service Service

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService() *Service {
	this := Service{}
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Service) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Service) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *Service) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *Service) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *Service) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *Service) UnsetComment() {
	o.Comment.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Service) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Service) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Service) SetName(v string) {
	o.Name = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *Service) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *Service) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *Service) SetCustomerID(v string) {
	o.CustomerID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Service) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Service) UnmarshalJSON(bytes []byte) (err error) {
	varService := _Service{}

	if err = json.Unmarshal(bytes, &varService); err == nil {
		*o = Service(varService)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "name")
		delete(additionalProperties, "customer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableService is a helper abstraction for handling nullable service types. 
type NullableService struct {
	value *Service
	isSet bool
}

// Get returns the value.
func (v NullableService) Get() *Service {
	return v.value
}

// Set modifies the value.
func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableService) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableService returns a pointer to a new instance of NullableService.
func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
