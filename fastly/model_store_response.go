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

// StoreResponse struct for StoreResponse
type StoreResponse struct {
	// ID of the store.
	ID *string `json:"id,omitempty"`
	// A human-readable name for the store.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _StoreResponse StoreResponse

// NewStoreResponse instantiates a new StoreResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreResponse() *StoreResponse {
	this := StoreResponse{}
	return &this
}

// NewStoreResponseWithDefaults instantiates a new StoreResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreResponseWithDefaults() *StoreResponse {
	this := StoreResponse{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *StoreResponse) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponse) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *StoreResponse) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *StoreResponse) SetID(v string) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StoreResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StoreResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StoreResponse) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o StoreResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
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
func (o *StoreResponse) UnmarshalJSON(bytes []byte) (err error) {
	varStoreResponse := _StoreResponse{}

	if err = json.Unmarshal(bytes, &varStoreResponse); err == nil {
		*o = StoreResponse(varStoreResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableStoreResponse is a helper abstraction for handling nullable storeresponse types. 
type NullableStoreResponse struct {
	value *StoreResponse
	isSet bool
}

// Get returns the value.
func (v NullableStoreResponse) Get() *StoreResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableStoreResponse) Set(val *StoreResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableStoreResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableStoreResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableStoreResponse returns a pointer to a new instance of NullableStoreResponse.
func NewNullableStoreResponse(val *StoreResponse) *NullableStoreResponse {
	return &NullableStoreResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableStoreResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableStoreResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
