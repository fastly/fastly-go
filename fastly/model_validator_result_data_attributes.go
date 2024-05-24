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

// ValidatorResultDataAttributes struct for ValidatorResultDataAttributes
type ValidatorResultDataAttributes struct {
	Msg NullableString `json:"msg,omitempty"`
	Status *string `json:"status,omitempty"`
	Errors []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	Messages []ValidatorResultDataAttributesMessages `json:"messages,omitempty"`
	AdditionalProperties map[string]any
}

type _ValidatorResultDataAttributes ValidatorResultDataAttributes

// NewValidatorResultDataAttributes instantiates a new ValidatorResultDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatorResultDataAttributes() *ValidatorResultDataAttributes {
	this := ValidatorResultDataAttributes{}
	return &this
}

// NewValidatorResultDataAttributesWithDefaults instantiates a new ValidatorResultDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatorResultDataAttributesWithDefaults() *ValidatorResultDataAttributes {
	this := ValidatorResultDataAttributes{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValidatorResultDataAttributes) GetMsg() string {
	if o == nil || o.Msg.Get() == nil {
		var ret string
		return ret
	}
	return *o.Msg.Get()
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValidatorResultDataAttributes) GetMsgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Msg.Get(), o.Msg.IsSet()
}

// HasMsg returns a boolean if a field has been set.
func (o *ValidatorResultDataAttributes) HasMsg() bool {
	if o != nil && o.Msg.IsSet() {
		return true
	}

	return false
}

// SetMsg gets a reference to the given NullableString and assigns it to the Msg field.
func (o *ValidatorResultDataAttributes) SetMsg(v string) {
	o.Msg.Set(&v)
}
// SetMsgNil sets the value for Msg to be an explicit nil
func (o *ValidatorResultDataAttributes) SetMsgNil() {
	o.Msg.Set(nil)
}

// UnsetMsg ensures that no value is present for Msg, not even an explicit nil
func (o *ValidatorResultDataAttributes) UnsetMsg() {
	o.Msg.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ValidatorResultDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatorResultDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ValidatorResultDataAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ValidatorResultDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ValidatorResultDataAttributes) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatorResultDataAttributes) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ValidatorResultDataAttributes) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ValidatorResultDataAttributes) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ValidatorResultDataAttributes) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatorResultDataAttributes) GetWarningsOk() ([]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ValidatorResultDataAttributes) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *ValidatorResultDataAttributes) SetWarnings(v []string) {
	o.Warnings = v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *ValidatorResultDataAttributes) GetMessages() []ValidatorResultDataAttributesMessages {
	if o == nil || o.Messages == nil {
		var ret []ValidatorResultDataAttributesMessages
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatorResultDataAttributes) GetMessagesOk() ([]ValidatorResultDataAttributesMessages, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ValidatorResultDataAttributes) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ValidatorResultDataAttributesMessages and assigns it to the Messages field.
func (o *ValidatorResultDataAttributes) SetMessages(v []ValidatorResultDataAttributesMessages) {
	o.Messages = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValidatorResultDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Msg.IsSet() {
		toSerialize["msg"] = o.Msg.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ValidatorResultDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varValidatorResultDataAttributes := _ValidatorResultDataAttributes{}

	if err = json.Unmarshal(bytes, &varValidatorResultDataAttributes); err == nil {
		*o = ValidatorResultDataAttributes(varValidatorResultDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "msg")
		delete(additionalProperties, "status")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "messages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValidatorResultDataAttributes is a helper abstraction for handling nullable validatorresultdataattributes types. 
type NullableValidatorResultDataAttributes struct {
	value *ValidatorResultDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableValidatorResultDataAttributes) Get() *ValidatorResultDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableValidatorResultDataAttributes) Set(val *ValidatorResultDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValidatorResultDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValidatorResultDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValidatorResultDataAttributes returns a pointer to a new instance of NullableValidatorResultDataAttributes.
func NewNullableValidatorResultDataAttributes(val *ValidatorResultDataAttributes) *NullableValidatorResultDataAttributes {
	return &NullableValidatorResultDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValidatorResultDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableValidatorResultDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
