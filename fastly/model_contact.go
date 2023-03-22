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

// Contact struct for Contact
type Contact struct {
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
	AdditionalProperties map[string]any
}

type _Contact Contact

// NewContact instantiates a new Contact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContact() *Contact {
	this := Contact{}
	return &this
}

// NewContactWithDefaults instantiates a new Contact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactWithDefaults() *Contact {
	this := Contact{}
	return &this
}

// GetUserID returns the UserID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contact) GetUserID() string {
	if o == nil || o.UserID.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserID.Get()
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contact) GetUserIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserID.Get(), o.UserID.IsSet()
}

// HasUserID returns a boolean if a field has been set.
func (o *Contact) HasUserID() bool {
	if o != nil && o.UserID.IsSet() {
		return true
	}

	return false
}

// SetUserID gets a reference to the given NullableString and assigns it to the UserID field.
func (o *Contact) SetUserID(v string) {
	o.UserID.Set(&v)
}
// SetUserIDNil sets the value for UserID to be an explicit nil
func (o *Contact) SetUserIDNil() {
	o.UserID.Set(nil)
}

// UnsetUserID ensures that no value is present for UserID, not even an explicit nil
func (o *Contact) UnsetUserID() {
	o.UserID.Unset()
}

// GetContactType returns the ContactType field value if set, zero value otherwise.
func (o *Contact) GetContactType() string {
	if o == nil || o.ContactType == nil {
		var ret string
		return ret
	}
	return *o.ContactType
}

// GetContactTypeOk returns a tuple with the ContactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetContactTypeOk() (*string, bool) {
	if o == nil || o.ContactType == nil {
		return nil, false
	}
	return o.ContactType, true
}

// HasContactType returns a boolean if a field has been set.
func (o *Contact) HasContactType() bool {
	if o != nil && o.ContactType != nil {
		return true
	}

	return false
}

// SetContactType gets a reference to the given string and assigns it to the ContactType field.
func (o *Contact) SetContactType(v string) {
	o.ContactType = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contact) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contact) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Contact) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Contact) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Contact) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Contact) UnsetName() {
	o.Name.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contact) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contact) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *Contact) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *Contact) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *Contact) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *Contact) UnsetEmail() {
	o.Email.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contact) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contact) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *Contact) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *Contact) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *Contact) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *Contact) UnsetPhone() {
	o.Phone.Unset()
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contact) GetCustomerID() string {
	if o == nil || o.CustomerID.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomerID.Get()
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contact) GetCustomerIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomerID.Get(), o.CustomerID.IsSet()
}

// HasCustomerID returns a boolean if a field has been set.
func (o *Contact) HasCustomerID() bool {
	if o != nil && o.CustomerID.IsSet() {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given NullableString and assigns it to the CustomerID field.
func (o *Contact) SetCustomerID(v string) {
	o.CustomerID.Set(&v)
}
// SetCustomerIDNil sets the value for CustomerID to be an explicit nil
func (o *Contact) SetCustomerIDNil() {
	o.CustomerID.Set(nil)
}

// UnsetCustomerID ensures that no value is present for CustomerID, not even an explicit nil
func (o *Contact) UnsetCustomerID() {
	o.CustomerID.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Contact) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Contact) UnmarshalJSON(bytes []byte) (err error) {
	varContact := _Contact{}

	if err = json.Unmarshal(bytes, &varContact); err == nil {
		*o = Contact(varContact)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "contact_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "customer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableContact is a helper abstraction for handling nullable contact types. 
type NullableContact struct {
	value *Contact
	isSet bool
}

// Get returns the value.
func (v NullableContact) Get() *Contact {
	return v.value
}

// Set modifies the value.
func (v *NullableContact) Set(val *Contact) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableContact) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableContact) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableContact returns a pointer to a new instance of NullableContact.
func NewNullableContact(val *Contact) *NullableContact {
	return &NullableContact{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
