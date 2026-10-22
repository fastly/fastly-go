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
	"time"
)

// ArcRotateVirtualKeyResponse The virtual key that was rotated, including its new access token.
type ArcRotateVirtualKeyResponse struct {
	// The unique identifier of the virtual key.
	Id string `json:"id"`
	// The ID of the customer account that owns the key.
	CustomerId string `json:"customer_id"`
	// The display name of the virtual key.
	Name string `json:"name"`
	// The ID of the user who created the key.
	UserId string `json:"user_id"`
	// The date and time the key expires, in RFC 3339 format. `null` if the key doesn't expire.
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	// The date and time the key was created, in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the key was last updated, in RFC 3339 format.
	UpdatedAt time.Time `json:"updated_at"`
	// The new access token for this virtual key. This is only returned when the key is created or rotated, and cannot be retrieved again afterward.
	AccessToken          string `json:"access_token"`
	AdditionalProperties map[string]any
}

type _ArcRotateVirtualKeyResponse ArcRotateVirtualKeyResponse

// NewArcRotateVirtualKeyResponse instantiates a new ArcRotateVirtualKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcRotateVirtualKeyResponse(id string, customerId string, name string, userId string, createdAt time.Time, updatedAt time.Time, accessToken string) *ArcRotateVirtualKeyResponse {
	this := ArcRotateVirtualKeyResponse{}
	this.Id = id
	this.CustomerId = customerId
	this.Name = name
	this.UserId = userId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.AccessToken = accessToken
	return &this
}

// NewArcRotateVirtualKeyResponseWithDefaults instantiates a new ArcRotateVirtualKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcRotateVirtualKeyResponseWithDefaults() *ArcRotateVirtualKeyResponse {
	this := ArcRotateVirtualKeyResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ArcRotateVirtualKeyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArcRotateVirtualKeyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArcRotateVirtualKeyResponse) SetId(v string) {
	o.Id = v
}

// GetCustomerId returns the CustomerId field value
func (o *ArcRotateVirtualKeyResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ArcRotateVirtualKeyResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ArcRotateVirtualKeyResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetName returns the Name field value
func (o *ArcRotateVirtualKeyResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArcRotateVirtualKeyResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArcRotateVirtualKeyResponse) SetName(v string) {
	o.Name = v
}

// GetUserId returns the UserId field value
func (o *ArcRotateVirtualKeyResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ArcRotateVirtualKeyResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ArcRotateVirtualKeyResponse) SetUserId(v string) {
	o.UserId = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcRotateVirtualKeyResponse) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcRotateVirtualKeyResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ArcRotateVirtualKeyResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *ArcRotateVirtualKeyResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *ArcRotateVirtualKeyResponse) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *ArcRotateVirtualKeyResponse) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArcRotateVirtualKeyResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcRotateVirtualKeyResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArcRotateVirtualKeyResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArcRotateVirtualKeyResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcRotateVirtualKeyResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArcRotateVirtualKeyResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetAccessToken returns the AccessToken field value
func (o *ArcRotateVirtualKeyResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ArcRotateVirtualKeyResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ArcRotateVirtualKeyResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcRotateVirtualKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcRotateVirtualKeyResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcRotateVirtualKeyResponse := _ArcRotateVirtualKeyResponse{}

	if err = json.Unmarshal(bytes, &varArcRotateVirtualKeyResponse); err == nil {
		*o = ArcRotateVirtualKeyResponse(varArcRotateVirtualKeyResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "access_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcRotateVirtualKeyResponse is a helper abstraction for handling nullable arcrotatevirtualkeyresponse types.
type NullableArcRotateVirtualKeyResponse struct {
	value *ArcRotateVirtualKeyResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcRotateVirtualKeyResponse) Get() *ArcRotateVirtualKeyResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcRotateVirtualKeyResponse) Set(val *ArcRotateVirtualKeyResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcRotateVirtualKeyResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcRotateVirtualKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcRotateVirtualKeyResponse returns a pointer to a new instance of NullableArcRotateVirtualKeyResponse.
func NewNullableArcRotateVirtualKeyResponse(val *ArcRotateVirtualKeyResponse) *NullableArcRotateVirtualKeyResponse {
	return &NullableArcRotateVirtualKeyResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcRotateVirtualKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcRotateVirtualKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
