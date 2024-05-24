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

// SchemasContactResponse struct for SchemasContactResponse
type SchemasContactResponse struct {
	// The alphanumeric string representing the user for this customer contact.
	UserID NullableString `json:"user_id,omitempty"`
	// The type of contact.
	ContactType *string `json:"contact_type,omitempty"`
	// The name of this contact, when user_id is not provided.
	Name NullableString `json:"name,omitempty"`
	// The email of this contact, when a user_id is not provided.
	Email NullableString `json:"email,omitempty"`
	// The phone number for this contact. Required for primary, technical, and security contact types.
	Phone NullableString `json:"phone,omitempty"`
	// The alphanumeric string representing the customer for this customer contact.
	CustomerID NullableString `json:"customer_id,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _SchemasContactResponse SchemasContactResponse

// NewSchemasContactResponse instantiates a new SchemasContactResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasContactResponse() *SchemasContactResponse {
	this := SchemasContactResponse{}
	return &this
}

// NewSchemasContactResponseWithDefaults instantiates a new SchemasContactResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasContactResponseWithDefaults() *SchemasContactResponse {
	this := SchemasContactResponse{}
	return &this
}

// GetUserID returns the UserID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasContactResponse) GetUserID() string {
	if o == nil || o.UserID.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserID.Get()
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasContactResponse) GetUserIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserID.Get(), o.UserID.IsSet()
}

// HasUserID returns a boolean if a field has been set.
func (o *SchemasContactResponse) HasUserID() bool {
	if o != nil && o.UserID.IsSet() {
		return true
	}

	return false
}

// SetUserID gets a reference to the given NullableString and assigns it to the UserID field.
func (o *SchemasContactResponse) SetUserID(v string) {
	o.UserID.Set(&v)
}
// SetUserIDNil sets the value for UserID to be an explicit nil
func (o *SchemasContactResponse) SetUserIDNil() {
	o.UserID.Set(nil)
}

// UnsetUserID ensures that no value is present for UserID, not even an explicit nil
func (o *SchemasContactResponse) UnsetUserID() {
	o.UserID.Unset()
}

// GetContactType returns the ContactType field value if set, zero value otherwise.
func (o *SchemasContactResponse) GetContactType() string {
	if o == nil || o.ContactType == nil {
		var ret string
		return ret
	}
	return *o.ContactType
}

// GetContactTypeOk returns a tuple with the ContactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasContactResponse) GetContactTypeOk() (*string, bool) {
	if o == nil || o.ContactType == nil {
		return nil, false
	}
	return o.ContactType, true
}

// HasContactType returns a boolean if a field has been set.
func (o *SchemasContactResponse) HasContactType() bool {
	if o != nil && o.ContactType != nil {
		return true
	}

	return false
}

// SetContactType gets a reference to the given string and assigns it to the ContactType field.
func (o *SchemasContactResponse) SetContactType(v string) {
	o.ContactType = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasContactResponse) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasContactResponse) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SchemasContactResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SchemasContactResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SchemasContactResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SchemasContactResponse) UnsetName() {
	o.Name.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasContactResponse) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasContactResponse) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *SchemasContactResponse) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *SchemasContactResponse) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *SchemasContactResponse) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *SchemasContactResponse) UnsetEmail() {
	o.Email.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasContactResponse) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasContactResponse) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *SchemasContactResponse) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *SchemasContactResponse) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *SchemasContactResponse) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *SchemasContactResponse) UnsetPhone() {
	o.Phone.Unset()
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasContactResponse) GetCustomerID() string {
	if o == nil || o.CustomerID.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomerID.Get()
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasContactResponse) GetCustomerIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomerID.Get(), o.CustomerID.IsSet()
}

// HasCustomerID returns a boolean if a field has been set.
func (o *SchemasContactResponse) HasCustomerID() bool {
	if o != nil && o.CustomerID.IsSet() {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given NullableString and assigns it to the CustomerID field.
func (o *SchemasContactResponse) SetCustomerID(v string) {
	o.CustomerID.Set(&v)
}
// SetCustomerIDNil sets the value for CustomerID to be an explicit nil
func (o *SchemasContactResponse) SetCustomerIDNil() {
	o.CustomerID.Set(nil)
}

// UnsetCustomerID ensures that no value is present for CustomerID, not even an explicit nil
func (o *SchemasContactResponse) UnsetCustomerID() {
	o.CustomerID.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasContactResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasContactResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SchemasContactResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *SchemasContactResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *SchemasContactResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *SchemasContactResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasContactResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasContactResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *SchemasContactResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *SchemasContactResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *SchemasContactResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *SchemasContactResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasContactResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasContactResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SchemasContactResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *SchemasContactResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *SchemasContactResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *SchemasContactResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *SchemasContactResponse) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasContactResponse) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *SchemasContactResponse) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *SchemasContactResponse) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SchemasContactResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.UserID.IsSet() {
		toSerialize["user_id"] = o.UserID.Get()
	}
	if o.ContactType != nil {
		toSerialize["contact_type"] = o.ContactType
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.CustomerID.IsSet() {
		toSerialize["customer_id"] = o.CustomerID.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
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
func (o *SchemasContactResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSchemasContactResponse := _SchemasContactResponse{}

	if err = json.Unmarshal(bytes, &varSchemasContactResponse); err == nil {
		*o = SchemasContactResponse(varSchemasContactResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "contact_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSchemasContactResponse is a helper abstraction for handling nullable schemascontactresponse types. 
type NullableSchemasContactResponse struct {
	value *SchemasContactResponse
	isSet bool
}

// Get returns the value.
func (v NullableSchemasContactResponse) Get() *SchemasContactResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableSchemasContactResponse) Set(val *SchemasContactResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSchemasContactResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSchemasContactResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSchemasContactResponse returns a pointer to a new instance of NullableSchemasContactResponse.
func NewNullableSchemasContactResponse(val *SchemasContactResponse) *NullableSchemasContactResponse {
	return &NullableSchemasContactResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSchemasContactResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSchemasContactResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
