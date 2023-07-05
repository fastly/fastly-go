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

// ValidatorResultDataAttributesMessages struct for ValidatorResultDataAttributesMessages
type ValidatorResultDataAttributesMessages struct {
	Type string `json:"type"`
	Warning bool `json:"warning"`
	Message string `json:"message"`
	Tokens []map[string]TokensAdditionalProps `json:"tokens"`
	AdditionalProperties map[string]any
}

type _ValidatorResultDataAttributesMessages ValidatorResultDataAttributesMessages

// NewValidatorResultDataAttributesMessages instantiates a new ValidatorResultDataAttributesMessages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatorResultDataAttributesMessages(resourceType string, warning bool, message string, tokens []map[string]TokensAdditionalProps) *ValidatorResultDataAttributesMessages {
	this := ValidatorResultDataAttributesMessages{}
	this.Type = resourceType
	this.Warning = warning
	this.Message = message
	this.Tokens = tokens
	return &this
}

// NewValidatorResultDataAttributesMessagesWithDefaults instantiates a new ValidatorResultDataAttributesMessages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatorResultDataAttributesMessagesWithDefaults() *ValidatorResultDataAttributesMessages {
	this := ValidatorResultDataAttributesMessages{}
	return &this
}

// GetType returns the Type field value
func (o *ValidatorResultDataAttributesMessages) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ValidatorResultDataAttributesMessages) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ValidatorResultDataAttributesMessages) SetType(v string) {
	o.Type = v
}

// GetWarning returns the Warning field value
func (o *ValidatorResultDataAttributesMessages) GetWarning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Warning
}

// GetWarningOk returns a tuple with the Warning field value
// and a boolean to check if the value has been set.
func (o *ValidatorResultDataAttributesMessages) GetWarningOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Warning, true
}

// SetWarning sets field value
func (o *ValidatorResultDataAttributesMessages) SetWarning(v bool) {
	o.Warning = v
}

// GetMessage returns the Message field value
func (o *ValidatorResultDataAttributesMessages) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ValidatorResultDataAttributesMessages) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ValidatorResultDataAttributesMessages) SetMessage(v string) {
	o.Message = v
}

// GetTokens returns the Tokens field value
func (o *ValidatorResultDataAttributesMessages) GetTokens() []map[string]TokensAdditionalProps {
	if o == nil {
		var ret []map[string]TokensAdditionalProps
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *ValidatorResultDataAttributesMessages) GetTokensOk() ([]map[string]TokensAdditionalProps, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Tokens, true
}

// SetTokens sets field value
func (o *ValidatorResultDataAttributesMessages) SetTokens(v []map[string]TokensAdditionalProps) {
	o.Tokens = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValidatorResultDataAttributesMessages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["warning"] = o.Warning
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["tokens"] = o.Tokens
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ValidatorResultDataAttributesMessages) UnmarshalJSON(bytes []byte) (err error) {
	varValidatorResultDataAttributesMessages := _ValidatorResultDataAttributesMessages{}

	if err = json.Unmarshal(bytes, &varValidatorResultDataAttributesMessages); err == nil {
		*o = ValidatorResultDataAttributesMessages(varValidatorResultDataAttributesMessages)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "warning")
		delete(additionalProperties, "message")
		delete(additionalProperties, "tokens")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValidatorResultDataAttributesMessages is a helper abstraction for handling nullable validatorresultdataattributesmessages types. 
type NullableValidatorResultDataAttributesMessages struct {
	value *ValidatorResultDataAttributesMessages
	isSet bool
}

// Get returns the value.
func (v NullableValidatorResultDataAttributesMessages) Get() *ValidatorResultDataAttributesMessages {
	return v.value
}

// Set modifies the value.
func (v *NullableValidatorResultDataAttributesMessages) Set(val *ValidatorResultDataAttributesMessages) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValidatorResultDataAttributesMessages) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValidatorResultDataAttributesMessages) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValidatorResultDataAttributesMessages returns a pointer to a new instance of NullableValidatorResultDataAttributesMessages.
func NewNullableValidatorResultDataAttributesMessages(val *ValidatorResultDataAttributesMessages) *NullableValidatorResultDataAttributesMessages {
	return &NullableValidatorResultDataAttributesMessages{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValidatorResultDataAttributesMessages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableValidatorResultDataAttributesMessages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
