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

// ArcCreateVirtualKeyResponse The virtual key that was created, including its access token.
type ArcCreateVirtualKeyResponse struct {
	// The unique identifier of the virtual key.
	ID string `json:"id"`
	// The ID of the customer account that owns the key.
	CustomerID string `json:"customer_id"`
	// The display name of the virtual key.
	Name string `json:"name"`
	// The ID of the user who created the key.
	UserID string `json:"user_id"`
	// The date and time the key expires, in RFC 3339 format. `null` if the key doesn't expire.
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	// The requests-per-minute limit for this key. `null` if there is no limit.
	RpmLimit NullableInt32 `json:"rpm_limit,omitempty"`
	// The tokens-per-minute limit for this key. `null` if there is no limit.
	TpmLimit NullableInt32 `json:"tpm_limit,omitempty"`
	// The date and time the key was created, in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the key was last updated, in RFC 3339 format.
	UpdatedAt time.Time `json:"updated_at"`
	// The access token for this virtual key. This is only returned when the key is created or rotated, and cannot be retrieved again afterward.
	AccessToken          string `json:"access_token"`
	AdditionalProperties map[string]any
}

type _ArcCreateVirtualKeyResponse ArcCreateVirtualKeyResponse

// NewArcCreateVirtualKeyResponse instantiates a new ArcCreateVirtualKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcCreateVirtualKeyResponse(id string, customerID string, name string, userID string, createdAt time.Time, updatedAt time.Time, accessToken string) *ArcCreateVirtualKeyResponse {
	this := ArcCreateVirtualKeyResponse{}
	this.ID = id
	this.CustomerID = customerID
	this.Name = name
	this.UserID = userID
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.AccessToken = accessToken
	return &this
}

// NewArcCreateVirtualKeyResponseWithDefaults instantiates a new ArcCreateVirtualKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcCreateVirtualKeyResponseWithDefaults() *ArcCreateVirtualKeyResponse {
	this := ArcCreateVirtualKeyResponse{}
	return &this
}

// GetID returns the ID field value
func (o *ArcCreateVirtualKeyResponse) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *ArcCreateVirtualKeyResponse) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *ArcCreateVirtualKeyResponse) SetID(v string) {
	o.ID = v
}

// GetCustomerID returns the CustomerID field value
func (o *ArcCreateVirtualKeyResponse) GetCustomerID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value
// and a boolean to check if the value has been set.
func (o *ArcCreateVirtualKeyResponse) GetCustomerIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerID, true
}

// SetCustomerID sets field value
func (o *ArcCreateVirtualKeyResponse) SetCustomerID(v string) {
	o.CustomerID = v
}

// GetName returns the Name field value
func (o *ArcCreateVirtualKeyResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArcCreateVirtualKeyResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArcCreateVirtualKeyResponse) SetName(v string) {
	o.Name = v
}

// GetUserID returns the UserID field value
func (o *ArcCreateVirtualKeyResponse) GetUserID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value
// and a boolean to check if the value has been set.
func (o *ArcCreateVirtualKeyResponse) GetUserIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserID, true
}

// SetUserID sets field value
func (o *ArcCreateVirtualKeyResponse) SetUserID(v string) {
	o.UserID = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcCreateVirtualKeyResponse) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcCreateVirtualKeyResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ArcCreateVirtualKeyResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *ArcCreateVirtualKeyResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *ArcCreateVirtualKeyResponse) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *ArcCreateVirtualKeyResponse) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetRpmLimit returns the RpmLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcCreateVirtualKeyResponse) GetRpmLimit() int32 {
	if o == nil || o.RpmLimit.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RpmLimit.Get()
}

// GetRpmLimitOk returns a tuple with the RpmLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcCreateVirtualKeyResponse) GetRpmLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RpmLimit.Get(), o.RpmLimit.IsSet()
}

// HasRpmLimit returns a boolean if a field has been set.
func (o *ArcCreateVirtualKeyResponse) HasRpmLimit() bool {
	if o != nil && o.RpmLimit.IsSet() {
		return true
	}

	return false
}

// SetRpmLimit gets a reference to the given NullableInt32 and assigns it to the RpmLimit field.
func (o *ArcCreateVirtualKeyResponse) SetRpmLimit(v int32) {
	o.RpmLimit.Set(&v)
}

// SetRpmLimitNil sets the value for RpmLimit to be an explicit nil
func (o *ArcCreateVirtualKeyResponse) SetRpmLimitNil() {
	o.RpmLimit.Set(nil)
}

// UnsetRpmLimit ensures that no value is present for RpmLimit, not even an explicit nil
func (o *ArcCreateVirtualKeyResponse) UnsetRpmLimit() {
	o.RpmLimit.Unset()
}

// GetTpmLimit returns the TpmLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcCreateVirtualKeyResponse) GetTpmLimit() int32 {
	if o == nil || o.TpmLimit.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TpmLimit.Get()
}

// GetTpmLimitOk returns a tuple with the TpmLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcCreateVirtualKeyResponse) GetTpmLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TpmLimit.Get(), o.TpmLimit.IsSet()
}

// HasTpmLimit returns a boolean if a field has been set.
func (o *ArcCreateVirtualKeyResponse) HasTpmLimit() bool {
	if o != nil && o.TpmLimit.IsSet() {
		return true
	}

	return false
}

// SetTpmLimit gets a reference to the given NullableInt32 and assigns it to the TpmLimit field.
func (o *ArcCreateVirtualKeyResponse) SetTpmLimit(v int32) {
	o.TpmLimit.Set(&v)
}

// SetTpmLimitNil sets the value for TpmLimit to be an explicit nil
func (o *ArcCreateVirtualKeyResponse) SetTpmLimitNil() {
	o.TpmLimit.Set(nil)
}

// UnsetTpmLimit ensures that no value is present for TpmLimit, not even an explicit nil
func (o *ArcCreateVirtualKeyResponse) UnsetTpmLimit() {
	o.TpmLimit.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArcCreateVirtualKeyResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcCreateVirtualKeyResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArcCreateVirtualKeyResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArcCreateVirtualKeyResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcCreateVirtualKeyResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArcCreateVirtualKeyResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetAccessToken returns the AccessToken field value
func (o *ArcCreateVirtualKeyResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ArcCreateVirtualKeyResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ArcCreateVirtualKeyResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcCreateVirtualKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.ID
	}
	if true {
		toSerialize["customer_id"] = o.CustomerID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["user_id"] = o.UserID
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.RpmLimit.IsSet() {
		toSerialize["rpm_limit"] = o.RpmLimit.Get()
	}
	if o.TpmLimit.IsSet() {
		toSerialize["tpm_limit"] = o.TpmLimit.Get()
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
func (o *ArcCreateVirtualKeyResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcCreateVirtualKeyResponse := _ArcCreateVirtualKeyResponse{}

	if err = json.Unmarshal(bytes, &varArcCreateVirtualKeyResponse); err == nil {
		*o = ArcCreateVirtualKeyResponse(varArcCreateVirtualKeyResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "rpm_limit")
		delete(additionalProperties, "tpm_limit")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "access_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcCreateVirtualKeyResponse is a helper abstraction for handling nullable arccreatevirtualkeyresponse types.
type NullableArcCreateVirtualKeyResponse struct {
	value *ArcCreateVirtualKeyResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcCreateVirtualKeyResponse) Get() *ArcCreateVirtualKeyResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcCreateVirtualKeyResponse) Set(val *ArcCreateVirtualKeyResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcCreateVirtualKeyResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcCreateVirtualKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcCreateVirtualKeyResponse returns a pointer to a new instance of NullableArcCreateVirtualKeyResponse.
func NewNullableArcCreateVirtualKeyResponse(val *ArcCreateVirtualKeyResponse) *NullableArcCreateVirtualKeyResponse {
	return &NullableArcCreateVirtualKeyResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcCreateVirtualKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcCreateVirtualKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
