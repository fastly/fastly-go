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

// Header struct for Header
type Header struct {
	// Accepts a string value.
	Action *string `json:"action,omitempty"`
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition NullableString `json:"cache_condition,omitempty"`
	// Header to set.
	Dst *string `json:"dst,omitempty"`
	// Don't add the header if it is added already. Only applies to 'set' action.
	IgnoreIfSet *int32 `json:"ignore_if_set,omitempty"`
	// A handle to refer to this Header object.
	Name *string `json:"name,omitempty"`
	// Priority determines execution order. Lower numbers execute first.
	Priority *int32 `json:"priority,omitempty"`
	// Regular expression to use. Only applies to `regex` and `regex_repeat` actions.
	Regex NullableString `json:"regex,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition NullableString `json:"request_condition,omitempty"`
	// Optional name of a response condition to apply.
	ResponseCondition NullableString `json:"response_condition,omitempty"`
	// Variable to be used as a source for the header content. Does not apply to `delete` action.
	Src NullableString `json:"src,omitempty"`
	// Value to substitute in place of regular expression. Only applies to `regex` and `regex_repeat` actions.
	Substitution NullableString `json:"substitution,omitempty"`
	// Accepts a string value.
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]any
}

type _Header Header

// NewHeader instantiates a new Header object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeader() *Header {
	this := Header{}
	var priority int32 = 100
	this.Priority = &priority
	return &this
}

// NewHeaderWithDefaults instantiates a new Header object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderWithDefaults() *Header {
	this := Header{}
	var priority int32 = 100
	this.Priority = &priority
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Header) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Header) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Header) SetAction(v string) {
	o.Action = &v
}

// GetCacheCondition returns the CacheCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Header) GetCacheCondition() string {
	if o == nil || o.CacheCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.CacheCondition.Get()
}

// GetCacheConditionOk returns a tuple with the CacheCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Header) GetCacheConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CacheCondition.Get(), o.CacheCondition.IsSet()
}

// HasCacheCondition returns a boolean if a field has been set.
func (o *Header) HasCacheCondition() bool {
	if o != nil && o.CacheCondition.IsSet() {
		return true
	}

	return false
}

// SetCacheCondition gets a reference to the given NullableString and assigns it to the CacheCondition field.
func (o *Header) SetCacheCondition(v string) {
	o.CacheCondition.Set(&v)
}
// SetCacheConditionNil sets the value for CacheCondition to be an explicit nil
func (o *Header) SetCacheConditionNil() {
	o.CacheCondition.Set(nil)
}

// UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
func (o *Header) UnsetCacheCondition() {
	o.CacheCondition.Unset()
}

// GetDst returns the Dst field value if set, zero value otherwise.
func (o *Header) GetDst() string {
	if o == nil || o.Dst == nil {
		var ret string
		return ret
	}
	return *o.Dst
}

// GetDstOk returns a tuple with the Dst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetDstOk() (*string, bool) {
	if o == nil || o.Dst == nil {
		return nil, false
	}
	return o.Dst, true
}

// HasDst returns a boolean if a field has been set.
func (o *Header) HasDst() bool {
	if o != nil && o.Dst != nil {
		return true
	}

	return false
}

// SetDst gets a reference to the given string and assigns it to the Dst field.
func (o *Header) SetDst(v string) {
	o.Dst = &v
}

// GetIgnoreIfSet returns the IgnoreIfSet field value if set, zero value otherwise.
func (o *Header) GetIgnoreIfSet() int32 {
	if o == nil || o.IgnoreIfSet == nil {
		var ret int32
		return ret
	}
	return *o.IgnoreIfSet
}

// GetIgnoreIfSetOk returns a tuple with the IgnoreIfSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetIgnoreIfSetOk() (*int32, bool) {
	if o == nil || o.IgnoreIfSet == nil {
		return nil, false
	}
	return o.IgnoreIfSet, true
}

// HasIgnoreIfSet returns a boolean if a field has been set.
func (o *Header) HasIgnoreIfSet() bool {
	if o != nil && o.IgnoreIfSet != nil {
		return true
	}

	return false
}

// SetIgnoreIfSet gets a reference to the given int32 and assigns it to the IgnoreIfSet field.
func (o *Header) SetIgnoreIfSet(v int32) {
	o.IgnoreIfSet = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Header) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Header) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Header) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Header) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Header) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *Header) SetPriority(v int32) {
	o.Priority = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Header) GetRegex() string {
	if o == nil || o.Regex.Get() == nil {
		var ret string
		return ret
	}
	return *o.Regex.Get()
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Header) GetRegexOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Regex.Get(), o.Regex.IsSet()
}

// HasRegex returns a boolean if a field has been set.
func (o *Header) HasRegex() bool {
	if o != nil && o.Regex.IsSet() {
		return true
	}

	return false
}

// SetRegex gets a reference to the given NullableString and assigns it to the Regex field.
func (o *Header) SetRegex(v string) {
	o.Regex.Set(&v)
}
// SetRegexNil sets the value for Regex to be an explicit nil
func (o *Header) SetRegexNil() {
	o.Regex.Set(nil)
}

// UnsetRegex ensures that no value is present for Regex, not even an explicit nil
func (o *Header) UnsetRegex() {
	o.Regex.Unset()
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Header) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Header) GetRequestConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *Header) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *Header) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}
// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *Header) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *Header) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Header) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Header) GetResponseConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *Header) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *Header) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}
// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *Header) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *Header) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetSrc returns the Src field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Header) GetSrc() string {
	if o == nil || o.Src.Get() == nil {
		var ret string
		return ret
	}
	return *o.Src.Get()
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Header) GetSrcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Src.Get(), o.Src.IsSet()
}

// HasSrc returns a boolean if a field has been set.
func (o *Header) HasSrc() bool {
	if o != nil && o.Src.IsSet() {
		return true
	}

	return false
}

// SetSrc gets a reference to the given NullableString and assigns it to the Src field.
func (o *Header) SetSrc(v string) {
	o.Src.Set(&v)
}
// SetSrcNil sets the value for Src to be an explicit nil
func (o *Header) SetSrcNil() {
	o.Src.Set(nil)
}

// UnsetSrc ensures that no value is present for Src, not even an explicit nil
func (o *Header) UnsetSrc() {
	o.Src.Unset()
}

// GetSubstitution returns the Substitution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Header) GetSubstitution() string {
	if o == nil || o.Substitution.Get() == nil {
		var ret string
		return ret
	}
	return *o.Substitution.Get()
}

// GetSubstitutionOk returns a tuple with the Substitution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Header) GetSubstitutionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Substitution.Get(), o.Substitution.IsSet()
}

// HasSubstitution returns a boolean if a field has been set.
func (o *Header) HasSubstitution() bool {
	if o != nil && o.Substitution.IsSet() {
		return true
	}

	return false
}

// SetSubstitution gets a reference to the given NullableString and assigns it to the Substitution field.
func (o *Header) SetSubstitution(v string) {
	o.Substitution.Set(&v)
}
// SetSubstitutionNil sets the value for Substitution to be an explicit nil
func (o *Header) SetSubstitutionNil() {
	o.Substitution.Set(nil)
}

// UnsetSubstitution ensures that no value is present for Substitution, not even an explicit nil
func (o *Header) UnsetSubstitution() {
	o.Substitution.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Header) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Header) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Header) SetType(v string) {
	o.Type = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Header) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.CacheCondition.IsSet() {
		toSerialize["cache_condition"] = o.CacheCondition.Get()
	}
	if o.Dst != nil {
		toSerialize["dst"] = o.Dst
	}
	if o.IgnoreIfSet != nil {
		toSerialize["ignore_if_set"] = o.IgnoreIfSet
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Regex.IsSet() {
		toSerialize["regex"] = o.Regex.Get()
	}
	if o.RequestCondition.IsSet() {
		toSerialize["request_condition"] = o.RequestCondition.Get()
	}
	if o.ResponseCondition.IsSet() {
		toSerialize["response_condition"] = o.ResponseCondition.Get()
	}
	if o.Src.IsSet() {
		toSerialize["src"] = o.Src.Get()
	}
	if o.Substitution.IsSet() {
		toSerialize["substitution"] = o.Substitution.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Header) UnmarshalJSON(bytes []byte) (err error) {
	varHeader := _Header{}

	if err = json.Unmarshal(bytes, &varHeader); err == nil {
		*o = Header(varHeader)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "cache_condition")
		delete(additionalProperties, "dst")
		delete(additionalProperties, "ignore_if_set")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "regex")
		delete(additionalProperties, "request_condition")
		delete(additionalProperties, "response_condition")
		delete(additionalProperties, "src")
		delete(additionalProperties, "substitution")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHeader is a helper abstraction for handling nullable header types. 
type NullableHeader struct {
	value *Header
	isSet bool
}

// Get returns the value.
func (v NullableHeader) Get() *Header {
	return v.value
}

// Set modifies the value.
func (v *NullableHeader) Set(val *Header) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHeader) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHeader) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHeader returns a pointer to a new instance of NullableHeader.
func NewNullableHeader(val *Header) *NullableHeader {
	return &NullableHeader{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
