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

// SudoRequest struct for SudoRequest
type SudoRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ExpiryTime *string `json:"expiry_time,omitempty"`
	AdditionalProperties map[string]any
}

type _SudoRequest SudoRequest

// NewSudoRequest instantiates a new SudoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSudoRequest(username string, password string) *SudoRequest {
	this := SudoRequest{}
	this.Username = username
	this.Password = password
	return &this
}

// NewSudoRequestWithDefaults instantiates a new SudoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSudoRequestWithDefaults() *SudoRequest {
	this := SudoRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *SudoRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SudoRequest) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SudoRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *SudoRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SudoRequest) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SudoRequest) SetPassword(v string) {
	o.Password = v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *SudoRequest) GetExpiryTime() string {
	if o == nil || o.ExpiryTime == nil {
		var ret string
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SudoRequest) GetExpiryTimeOk() (*string, bool) {
	if o == nil || o.ExpiryTime == nil {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *SudoRequest) HasExpiryTime() bool {
	if o != nil && o.ExpiryTime != nil {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given string and assigns it to the ExpiryTime field.
func (o *SudoRequest) SetExpiryTime(v string) {
	o.ExpiryTime = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SudoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.ExpiryTime != nil {
		toSerialize["expiry_time"] = o.ExpiryTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *SudoRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSudoRequest := _SudoRequest{}

	if err = json.Unmarshal(bytes, &varSudoRequest); err == nil {
		*o = SudoRequest(varSudoRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "expiry_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSudoRequest is a helper abstraction for handling nullable sudorequest types. 
type NullableSudoRequest struct {
	value *SudoRequest
	isSet bool
}

// Get returns the value.
func (v NullableSudoRequest) Get() *SudoRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableSudoRequest) Set(val *SudoRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSudoRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSudoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSudoRequest returns a pointer to a new instance of NullableSudoRequest.
func NewNullableSudoRequest(val *SudoRequest) *NullableSudoRequest {
	return &NullableSudoRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSudoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSudoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
