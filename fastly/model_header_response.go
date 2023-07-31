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

// HeaderResponse struct for HeaderResponse
type HeaderResponse struct {
	// Accepts a string value.
	Action *string `json:"action,omitempty"`
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition NullableString `json:"cache_condition,omitempty"`
	// Header to set.
	Dst *string `json:"dst,omitempty"`
	// A handle to refer to this Header object.
	Name *string `json:"name,omitempty"`
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
	// Don't add the header if it is added already. Only applies to 'set' action. Numerical value (\"0\" = false, \"1\" = true)
	IgnoreIfSet *string `json:"ignore_if_set,omitempty"`
	// Priority determines execution order. Lower numbers execute first.
	Priority *string `json:"priority,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	Version *string `json:"version,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _HeaderResponse HeaderResponse

// NewHeaderResponse instantiates a new HeaderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaderResponse() *HeaderResponse {
	this := HeaderResponse{}
	var priority string = "100"
	this.Priority = &priority
	return &this
}

// NewHeaderResponseWithDefaults instantiates a new HeaderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderResponseWithDefaults() *HeaderResponse {
	this := HeaderResponse{}
	var priority string = "100"
	this.Priority = &priority
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *HeaderResponse) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderResponse) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *HeaderResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *HeaderResponse) SetAction(v string) {
	o.Action = &v
}

// GetCacheCondition returns the CacheCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderResponse) GetCacheCondition() string {
	if o == nil || o.CacheCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.CacheCondition.Get()
}

// GetCacheConditionOk returns a tuple with the CacheCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderResponse) GetCacheConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CacheCondition.Get(), o.CacheCondition.IsSet()
}

// HasCacheCondition returns a boolean if a field has been set.
func (o *HeaderResponse) HasCacheCondition() bool {
	if o != nil && o.CacheCondition.IsSet() {
		return true
	}

	return false
}

// SetCacheCondition gets a reference to the given NullableString and assigns it to the CacheCondition field.
func (o *HeaderResponse) SetCacheCondition(v string) {
	o.CacheCondition.Set(&v)
}
// SetCacheConditionNil sets the value for CacheCondition to be an explicit nil
func (o *HeaderResponse) SetCacheConditionNil() {
	o.CacheCondition.Set(nil)
}

// UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
func (o *HeaderResponse) UnsetCacheCondition() {
	o.CacheCondition.Unset()
}

// GetDst returns the Dst field value if set, zero value otherwise.
func (o *HeaderResponse) GetDst() string {
	if o == nil || o.Dst == nil {
		var ret string
		return ret
	}
	return *o.Dst
}

// GetDstOk returns a tuple with the Dst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderResponse) GetDstOk() (*string, bool) {
	if o == nil || o.Dst == nil {
		return nil, false
	}
	return o.Dst, true
}

// HasDst returns a boolean if a field has been set.
func (o *HeaderResponse) HasDst() bool {
	if o != nil && o.Dst != nil {
		return true
	}

	return false
}

// SetDst gets a reference to the given string and assigns it to the Dst field.
func (o *HeaderResponse) SetDst(v string) {
	o.Dst = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HeaderResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HeaderResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HeaderResponse) SetName(v string) {
	o.Name = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderResponse) GetRegex() string {
	if o == nil || o.Regex.Get() == nil {
		var ret string
		return ret
	}
	return *o.Regex.Get()
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderResponse) GetRegexOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Regex.Get(), o.Regex.IsSet()
}

// HasRegex returns a boolean if a field has been set.
func (o *HeaderResponse) HasRegex() bool {
	if o != nil && o.Regex.IsSet() {
		return true
	}

	return false
}

// SetRegex gets a reference to the given NullableString and assigns it to the Regex field.
func (o *HeaderResponse) SetRegex(v string) {
	o.Regex.Set(&v)
}
// SetRegexNil sets the value for Regex to be an explicit nil
func (o *HeaderResponse) SetRegexNil() {
	o.Regex.Set(nil)
}

// UnsetRegex ensures that no value is present for Regex, not even an explicit nil
func (o *HeaderResponse) UnsetRegex() {
	o.Regex.Unset()
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderResponse) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderResponse) GetRequestConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *HeaderResponse) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *HeaderResponse) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}
// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *HeaderResponse) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *HeaderResponse) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// GetResponseCondition returns the ResponseCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderResponse) GetResponseCondition() string {
	if o == nil || o.ResponseCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCondition.Get()
}

// GetResponseConditionOk returns a tuple with the ResponseCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderResponse) GetResponseConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseCondition.Get(), o.ResponseCondition.IsSet()
}

// HasResponseCondition returns a boolean if a field has been set.
func (o *HeaderResponse) HasResponseCondition() bool {
	if o != nil && o.ResponseCondition.IsSet() {
		return true
	}

	return false
}

// SetResponseCondition gets a reference to the given NullableString and assigns it to the ResponseCondition field.
func (o *HeaderResponse) SetResponseCondition(v string) {
	o.ResponseCondition.Set(&v)
}
// SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil
func (o *HeaderResponse) SetResponseConditionNil() {
	o.ResponseCondition.Set(nil)
}

// UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
func (o *HeaderResponse) UnsetResponseCondition() {
	o.ResponseCondition.Unset()
}

// GetSrc returns the Src field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderResponse) GetSrc() string {
	if o == nil || o.Src.Get() == nil {
		var ret string
		return ret
	}
	return *o.Src.Get()
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderResponse) GetSrcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Src.Get(), o.Src.IsSet()
}

// HasSrc returns a boolean if a field has been set.
func (o *HeaderResponse) HasSrc() bool {
	if o != nil && o.Src.IsSet() {
		return true
	}

	return false
}

// SetSrc gets a reference to the given NullableString and assigns it to the Src field.
func (o *HeaderResponse) SetSrc(v string) {
	o.Src.Set(&v)
}
// SetSrcNil sets the value for Src to be an explicit nil
func (o *HeaderResponse) SetSrcNil() {
	o.Src.Set(nil)
}

// UnsetSrc ensures that no value is present for Src, not even an explicit nil
func (o *HeaderResponse) UnsetSrc() {
	o.Src.Unset()
}

// GetSubstitution returns the Substitution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderResponse) GetSubstitution() string {
	if o == nil || o.Substitution.Get() == nil {
		var ret string
		return ret
	}
	return *o.Substitution.Get()
}

// GetSubstitutionOk returns a tuple with the Substitution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderResponse) GetSubstitutionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Substitution.Get(), o.Substitution.IsSet()
}

// HasSubstitution returns a boolean if a field has been set.
func (o *HeaderResponse) HasSubstitution() bool {
	if o != nil && o.Substitution.IsSet() {
		return true
	}

	return false
}

// SetSubstitution gets a reference to the given NullableString and assigns it to the Substitution field.
func (o *HeaderResponse) SetSubstitution(v string) {
	o.Substitution.Set(&v)
}
// SetSubstitutionNil sets the value for Substitution to be an explicit nil
func (o *HeaderResponse) SetSubstitutionNil() {
	o.Substitution.Set(nil)
}

// UnsetSubstitution ensures that no value is present for Substitution, not even an explicit nil
func (o *HeaderResponse) UnsetSubstitution() {
	o.Substitution.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HeaderResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HeaderResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HeaderResponse) SetType(v string) {
	o.Type = &v
}

// GetIgnoreIfSet returns the IgnoreIfSet field value if set, zero value otherwise.
func (o *HeaderResponse) GetIgnoreIfSet() string {
	if o == nil || o.IgnoreIfSet == nil {
		var ret string
		return ret
	}
	return *o.IgnoreIfSet
}

// GetIgnoreIfSetOk returns a tuple with the IgnoreIfSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderResponse) GetIgnoreIfSetOk() (*string, bool) {
	if o == nil || o.IgnoreIfSet == nil {
		return nil, false
	}
	return o.IgnoreIfSet, true
}

// HasIgnoreIfSet returns a boolean if a field has been set.
func (o *HeaderResponse) HasIgnoreIfSet() bool {
	if o != nil && o.IgnoreIfSet != nil {
		return true
	}

	return false
}

// SetIgnoreIfSet gets a reference to the given string and assigns it to the IgnoreIfSet field.
func (o *HeaderResponse) SetIgnoreIfSet(v string) {
	o.IgnoreIfSet = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *HeaderResponse) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderResponse) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *HeaderResponse) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *HeaderResponse) SetPriority(v string) {
	o.Priority = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *HeaderResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *HeaderResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *HeaderResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HeaderResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HeaderResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HeaderResponse) SetVersion(v string) {
	o.Version = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *HeaderResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *HeaderResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *HeaderResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *HeaderResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *HeaderResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *HeaderResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *HeaderResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *HeaderResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *HeaderResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *HeaderResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *HeaderResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *HeaderResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HeaderResponse) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.IgnoreIfSet != nil {
		toSerialize["ignore_if_set"] = o.IgnoreIfSet
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *HeaderResponse) UnmarshalJSON(bytes []byte) (err error) {
	varHeaderResponse := _HeaderResponse{}

	if err = json.Unmarshal(bytes, &varHeaderResponse); err == nil {
		*o = HeaderResponse(varHeaderResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "cache_condition")
		delete(additionalProperties, "dst")
		delete(additionalProperties, "name")
		delete(additionalProperties, "regex")
		delete(additionalProperties, "request_condition")
		delete(additionalProperties, "response_condition")
		delete(additionalProperties, "src")
		delete(additionalProperties, "substitution")
		delete(additionalProperties, "type")
		delete(additionalProperties, "ignore_if_set")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHeaderResponse is a helper abstraction for handling nullable headerresponse types. 
type NullableHeaderResponse struct {
	value *HeaderResponse
	isSet bool
}

// Get returns the value.
func (v NullableHeaderResponse) Get() *HeaderResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableHeaderResponse) Set(val *HeaderResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHeaderResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHeaderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHeaderResponse returns a pointer to a new instance of NullableHeaderResponse.
func NewNullableHeaderResponse(val *HeaderResponse) *NullableHeaderResponse {
	return &NullableHeaderResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHeaderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHeaderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
