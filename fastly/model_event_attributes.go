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
	"time"
)

// EventAttributes struct for EventAttributes
type EventAttributes struct {
	// Indicates if event was performed by Fastly.
	Admin *bool `json:"admin,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	CustomerID *string `json:"customer_id,omitempty"`
	// Description of the event.
	Description *string `json:"description,omitempty"`
	// Type of event. Can be used with `filter[event_type]`
	EventType *string `json:"event_type,omitempty"`
	// IP addresses that the event was requested from.
	IP *string `json:"ip,omitempty"`
	// Hash of key value pairs of additional information.
	Metadata map[string]map[string]any `json:"metadata,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	UserID *string `json:"user_id,omitempty"`
	TokenID *string `json:"token_id,omitempty"`
	AdditionalProperties map[string]any
}

type _EventAttributes EventAttributes

// NewEventAttributes instantiates a new EventAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventAttributes() *EventAttributes {
	this := EventAttributes{}
	return &this
}

// NewEventAttributesWithDefaults instantiates a new EventAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventAttributesWithDefaults() *EventAttributes {
	this := EventAttributes{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *EventAttributes) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *EventAttributes) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *EventAttributes) SetAdmin(v bool) {
	o.Admin = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EventAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *EventAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *EventAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *EventAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *EventAttributes) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *EventAttributes) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *EventAttributes) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EventAttributes) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EventAttributes) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *EventAttributes) SetEventType(v string) {
	o.EventType = &v
}

// GetIP returns the IP field value if set, zero value otherwise.
func (o *EventAttributes) GetIP() string {
	if o == nil || o.IP == nil {
		var ret string
		return ret
	}
	return *o.IP
}

// GetIPOk returns a tuple with the IP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetIPOk() (*string, bool) {
	if o == nil || o.IP == nil {
		return nil, false
	}
	return o.IP, true
}

// HasIP returns a boolean if a field has been set.
func (o *EventAttributes) HasIP() bool {
	if o != nil && o.IP != nil {
		return true
	}

	return false
}

// SetIP gets a reference to the given string and assigns it to the IP field.
func (o *EventAttributes) SetIP(v string) {
	o.IP = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *EventAttributes) GetMetadata() map[string]map[string]any {
	if o == nil || o.Metadata == nil {
		var ret map[string]map[string]any
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetMetadataOk() (map[string]map[string]any, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *EventAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]any and assigns it to the Metadata field.
func (o *EventAttributes) SetMetadata(v map[string]map[string]any) {
	o.Metadata = v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *EventAttributes) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *EventAttributes) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *EventAttributes) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *EventAttributes) GetUserID() string {
	if o == nil || o.UserID == nil {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetUserIDOk() (*string, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *EventAttributes) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *EventAttributes) SetUserID(v string) {
	o.UserID = &v
}

// GetTokenID returns the TokenID field value if set, zero value otherwise.
func (o *EventAttributes) GetTokenID() string {
	if o == nil || o.TokenID == nil {
		var ret string
		return ret
	}
	return *o.TokenID
}

// GetTokenIDOk returns a tuple with the TokenID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAttributes) GetTokenIDOk() (*string, bool) {
	if o == nil || o.TokenID == nil {
		return nil, false
	}
	return o.TokenID, true
}

// HasTokenID returns a boolean if a field has been set.
func (o *EventAttributes) HasTokenID() bool {
	if o != nil && o.TokenID != nil {
		return true
	}

	return false
}

// SetTokenID gets a reference to the given string and assigns it to the TokenID field.
func (o *EventAttributes) SetTokenID(v string) {
	o.TokenID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o EventAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EventType != nil {
		toSerialize["event_type"] = o.EventType
	}
	if o.IP != nil {
		toSerialize["ip"] = o.IP
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.UserID != nil {
		toSerialize["user_id"] = o.UserID
	}
	if o.TokenID != nil {
		toSerialize["token_id"] = o.TokenID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *EventAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varEventAttributes := _EventAttributes{}

	if err = json.Unmarshal(bytes, &varEventAttributes); err == nil {
		*o = EventAttributes(varEventAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "admin")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "event_type")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableEventAttributes is a helper abstraction for handling nullable eventattributes types. 
type NullableEventAttributes struct {
	value *EventAttributes
	isSet bool
}

// Get returns the value.
func (v NullableEventAttributes) Get() *EventAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableEventAttributes) Set(val *EventAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableEventAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableEventAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableEventAttributes returns a pointer to a new instance of NullableEventAttributes.
func NewNullableEventAttributes(val *EventAttributes) *NullableEventAttributes {
	return &NullableEventAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableEventAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableEventAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
