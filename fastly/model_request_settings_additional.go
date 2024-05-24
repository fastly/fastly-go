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

// RequestSettingsAdditional struct for RequestSettingsAdditional
type RequestSettingsAdditional struct {
	// Allows you to terminate request handling and immediately perform an action.
	Action NullableString `json:"action,omitempty"`
	// Sets the host header.
	DefaultHost NullableString `json:"default_host,omitempty"`
	// Comma separated list of varnish request object fields that should be in the hash key.
	HashKeys NullableString `json:"hash_keys,omitempty"`
	// Name for the request settings.
	Name *string `json:"name,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition NullableString `json:"request_condition,omitempty"`
	// Short for X-Forwarded-For.
	Xff NullableString `json:"xff,omitempty"`
	AdditionalProperties map[string]any
}

type _RequestSettingsAdditional RequestSettingsAdditional

// NewRequestSettingsAdditional instantiates a new RequestSettingsAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSettingsAdditional() *RequestSettingsAdditional {
	this := RequestSettingsAdditional{}
	return &this
}

// NewRequestSettingsAdditionalWithDefaults instantiates a new RequestSettingsAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSettingsAdditionalWithDefaults() *RequestSettingsAdditional {
	this := RequestSettingsAdditional{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsAdditional) GetAction() string {
	if o == nil || o.Action.Get() == nil {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsAdditional) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *RequestSettingsAdditional) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *RequestSettingsAdditional) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *RequestSettingsAdditional) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *RequestSettingsAdditional) UnsetAction() {
	o.Action.Unset()
}

// GetDefaultHost returns the DefaultHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsAdditional) GetDefaultHost() string {
	if o == nil || o.DefaultHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultHost.Get()
}

// GetDefaultHostOk returns a tuple with the DefaultHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsAdditional) GetDefaultHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultHost.Get(), o.DefaultHost.IsSet()
}

// HasDefaultHost returns a boolean if a field has been set.
func (o *RequestSettingsAdditional) HasDefaultHost() bool {
	if o != nil && o.DefaultHost.IsSet() {
		return true
	}

	return false
}

// SetDefaultHost gets a reference to the given NullableString and assigns it to the DefaultHost field.
func (o *RequestSettingsAdditional) SetDefaultHost(v string) {
	o.DefaultHost.Set(&v)
}
// SetDefaultHostNil sets the value for DefaultHost to be an explicit nil
func (o *RequestSettingsAdditional) SetDefaultHostNil() {
	o.DefaultHost.Set(nil)
}

// UnsetDefaultHost ensures that no value is present for DefaultHost, not even an explicit nil
func (o *RequestSettingsAdditional) UnsetDefaultHost() {
	o.DefaultHost.Unset()
}

// GetHashKeys returns the HashKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsAdditional) GetHashKeys() string {
	if o == nil || o.HashKeys.Get() == nil {
		var ret string
		return ret
	}
	return *o.HashKeys.Get()
}

// GetHashKeysOk returns a tuple with the HashKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsAdditional) GetHashKeysOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HashKeys.Get(), o.HashKeys.IsSet()
}

// HasHashKeys returns a boolean if a field has been set.
func (o *RequestSettingsAdditional) HasHashKeys() bool {
	if o != nil && o.HashKeys.IsSet() {
		return true
	}

	return false
}

// SetHashKeys gets a reference to the given NullableString and assigns it to the HashKeys field.
func (o *RequestSettingsAdditional) SetHashKeys(v string) {
	o.HashKeys.Set(&v)
}
// SetHashKeysNil sets the value for HashKeys to be an explicit nil
func (o *RequestSettingsAdditional) SetHashKeysNil() {
	o.HashKeys.Set(nil)
}

// UnsetHashKeys ensures that no value is present for HashKeys, not even an explicit nil
func (o *RequestSettingsAdditional) UnsetHashKeys() {
	o.HashKeys.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestSettingsAdditional) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSettingsAdditional) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestSettingsAdditional) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestSettingsAdditional) SetName(v string) {
	o.Name = &v
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsAdditional) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsAdditional) GetRequestConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *RequestSettingsAdditional) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *RequestSettingsAdditional) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}
// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *RequestSettingsAdditional) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *RequestSettingsAdditional) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// GetXff returns the Xff field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSettingsAdditional) GetXff() string {
	if o == nil || o.Xff.Get() == nil {
		var ret string
		return ret
	}
	return *o.Xff.Get()
}

// GetXffOk returns a tuple with the Xff field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSettingsAdditional) GetXffOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Xff.Get(), o.Xff.IsSet()
}

// HasXff returns a boolean if a field has been set.
func (o *RequestSettingsAdditional) HasXff() bool {
	if o != nil && o.Xff.IsSet() {
		return true
	}

	return false
}

// SetXff gets a reference to the given NullableString and assigns it to the Xff field.
func (o *RequestSettingsAdditional) SetXff(v string) {
	o.Xff.Set(&v)
}
// SetXffNil sets the value for Xff to be an explicit nil
func (o *RequestSettingsAdditional) SetXffNil() {
	o.Xff.Set(nil)
}

// UnsetXff ensures that no value is present for Xff, not even an explicit nil
func (o *RequestSettingsAdditional) UnsetXff() {
	o.Xff.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RequestSettingsAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.DefaultHost.IsSet() {
		toSerialize["default_host"] = o.DefaultHost.Get()
	}
	if o.HashKeys.IsSet() {
		toSerialize["hash_keys"] = o.HashKeys.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RequestCondition.IsSet() {
		toSerialize["request_condition"] = o.RequestCondition.Get()
	}
	if o.Xff.IsSet() {
		toSerialize["xff"] = o.Xff.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RequestSettingsAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varRequestSettingsAdditional := _RequestSettingsAdditional{}

	if err = json.Unmarshal(bytes, &varRequestSettingsAdditional); err == nil {
		*o = RequestSettingsAdditional(varRequestSettingsAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "default_host")
		delete(additionalProperties, "hash_keys")
		delete(additionalProperties, "name")
		delete(additionalProperties, "request_condition")
		delete(additionalProperties, "xff")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRequestSettingsAdditional is a helper abstraction for handling nullable requestsettingsadditional types. 
type NullableRequestSettingsAdditional struct {
	value *RequestSettingsAdditional
	isSet bool
}

// Get returns the value.
func (v NullableRequestSettingsAdditional) Get() *RequestSettingsAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableRequestSettingsAdditional) Set(val *RequestSettingsAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRequestSettingsAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRequestSettingsAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRequestSettingsAdditional returns a pointer to a new instance of NullableRequestSettingsAdditional.
func NewNullableRequestSettingsAdditional(val *RequestSettingsAdditional) *NullableRequestSettingsAdditional {
	return &NullableRequestSettingsAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRequestSettingsAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRequestSettingsAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
