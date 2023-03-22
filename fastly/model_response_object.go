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

// ResponseObject struct for ResponseObject
type ResponseObject struct {
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition NullableString `json:"cache_condition,omitempty"`
	// The content to deliver for the response object, can be empty.
	Content *string `json:"content,omitempty"`
	// The MIME type of the content, can be empty.
	ContentType NullableString `json:"content_type,omitempty"`
	// Name for the request settings.
	Name *string `json:"name,omitempty"`
	// The HTTP status code.
	Status *int32 `json:"status,omitempty"`
	// The HTTP response.
	Response *string `json:"response,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition NullableString `json:"request_condition,omitempty"`
	AdditionalProperties map[string]any
}

type _ResponseObject ResponseObject

// NewResponseObject instantiates a new ResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseObject() *ResponseObject {
	this := ResponseObject{}
	var status int32 = 200
	this.Status = &status
	var response string = "Ok"
	this.Response = &response
	return &this
}

// NewResponseObjectWithDefaults instantiates a new ResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseObjectWithDefaults() *ResponseObject {
	this := ResponseObject{}
	var status int32 = 200
	this.Status = &status
	var response string = "Ok"
	this.Response = &response
	return &this
}

// GetCacheCondition returns the CacheCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseObject) GetCacheCondition() string {
	if o == nil || o.CacheCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.CacheCondition.Get()
}

// GetCacheConditionOk returns a tuple with the CacheCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseObject) GetCacheConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CacheCondition.Get(), o.CacheCondition.IsSet()
}

// HasCacheCondition returns a boolean if a field has been set.
func (o *ResponseObject) HasCacheCondition() bool {
	if o != nil && o.CacheCondition.IsSet() {
		return true
	}

	return false
}

// SetCacheCondition gets a reference to the given NullableString and assigns it to the CacheCondition field.
func (o *ResponseObject) SetCacheCondition(v string) {
	o.CacheCondition.Set(&v)
}
// SetCacheConditionNil sets the value for CacheCondition to be an explicit nil
func (o *ResponseObject) SetCacheConditionNil() {
	o.CacheCondition.Set(nil)
}

// UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
func (o *ResponseObject) UnsetCacheCondition() {
	o.CacheCondition.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ResponseObject) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseObject) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ResponseObject) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ResponseObject) SetContent(v string) {
	o.Content = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseObject) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseObject) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *ResponseObject) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *ResponseObject) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *ResponseObject) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *ResponseObject) UnsetContentType() {
	o.ContentType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponseObject) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseObject) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponseObject) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResponseObject) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseObject) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseObject) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseObject) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ResponseObject) SetStatus(v int32) {
	o.Status = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ResponseObject) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseObject) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ResponseObject) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *ResponseObject) SetResponse(v string) {
	o.Response = &v
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseObject) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseObject) GetRequestConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *ResponseObject) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *ResponseObject) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}
// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *ResponseObject) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *ResponseObject) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CacheCondition.IsSet() {
		toSerialize["cache_condition"] = o.CacheCondition.Get()
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.RequestCondition.IsSet() {
		toSerialize["request_condition"] = o.RequestCondition.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ResponseObject) UnmarshalJSON(bytes []byte) (err error) {
	varResponseObject := _ResponseObject{}

	if err = json.Unmarshal(bytes, &varResponseObject); err == nil {
		*o = ResponseObject(varResponseObject)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cache_condition")
		delete(additionalProperties, "content")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "response")
		delete(additionalProperties, "request_condition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableResponseObject is a helper abstraction for handling nullable responseobject types. 
type NullableResponseObject struct {
	value *ResponseObject
	isSet bool
}

// Get returns the value.
func (v NullableResponseObject) Get() *ResponseObject {
	return v.value
}

// Set modifies the value.
func (v *NullableResponseObject) Set(val *ResponseObject) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableResponseObject) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableResponseObject returns a pointer to a new instance of NullableResponseObject.
func NewNullableResponseObject(val *ResponseObject) *NullableResponseObject {
	return &NullableResponseObject{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
