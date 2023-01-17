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

// InvitationDataAttributes struct for InvitationDataAttributes
type InvitationDataAttributes struct {
	// The email address of the invitee.
	Email *string `json:"email,omitempty"`
	// Indicates the user has limited access to the customer's services.
	LimitServices *bool `json:"limit_services,omitempty"`
	Role *RoleUser `json:"role,omitempty"`
	// Indicates whether or not the invitation is active.
	StatusCode *int32 `json:"status_code,omitempty"`
	AdditionalProperties map[string]any
}

type _InvitationDataAttributes InvitationDataAttributes

// NewInvitationDataAttributes instantiates a new InvitationDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationDataAttributes() *InvitationDataAttributes {
	this := InvitationDataAttributes{}
	return &this
}

// NewInvitationDataAttributesWithDefaults instantiates a new InvitationDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationDataAttributesWithDefaults() *InvitationDataAttributes {
	this := InvitationDataAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InvitationDataAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InvitationDataAttributes) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InvitationDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetLimitServices returns the LimitServices field value if set, zero value otherwise.
func (o *InvitationDataAttributes) GetLimitServices() bool {
	if o == nil || o.LimitServices == nil {
		var ret bool
		return ret
	}
	return *o.LimitServices
}

// GetLimitServicesOk returns a tuple with the LimitServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationDataAttributes) GetLimitServicesOk() (*bool, bool) {
	if o == nil || o.LimitServices == nil {
		return nil, false
	}
	return o.LimitServices, true
}

// HasLimitServices returns a boolean if a field has been set.
func (o *InvitationDataAttributes) HasLimitServices() bool {
	if o != nil && o.LimitServices != nil {
		return true
	}

	return false
}

// SetLimitServices gets a reference to the given bool and assigns it to the LimitServices field.
func (o *InvitationDataAttributes) SetLimitServices(v bool) {
	o.LimitServices = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *InvitationDataAttributes) GetRole() RoleUser {
	if o == nil || o.Role == nil {
		var ret RoleUser
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationDataAttributes) GetRoleOk() (*RoleUser, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *InvitationDataAttributes) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given RoleUser and assigns it to the Role field.
func (o *InvitationDataAttributes) SetRole(v RoleUser) {
	o.Role = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *InvitationDataAttributes) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationDataAttributes) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *InvitationDataAttributes) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *InvitationDataAttributes) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InvitationDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.LimitServices != nil {
		toSerialize["limit_services"] = o.LimitServices
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *InvitationDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varInvitationDataAttributes := _InvitationDataAttributes{}

	if err = json.Unmarshal(bytes, &varInvitationDataAttributes); err == nil {
		*o = InvitationDataAttributes(varInvitationDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "limit_services")
		delete(additionalProperties, "role")
		delete(additionalProperties, "status_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInvitationDataAttributes is a helper abstraction for handling nullable invitationdataattributes types. 
type NullableInvitationDataAttributes struct {
	value *InvitationDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableInvitationDataAttributes) Get() *InvitationDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableInvitationDataAttributes) Set(val *InvitationDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInvitationDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInvitationDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInvitationDataAttributes returns a pointer to a new instance of NullableInvitationDataAttributes.
func NewNullableInvitationDataAttributes(val *InvitationDataAttributes) *NullableInvitationDataAttributes {
	return &NullableInvitationDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInvitationDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableInvitationDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
