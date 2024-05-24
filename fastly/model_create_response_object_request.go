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

// CreateResponseObjectRequest struct for CreateResponseObjectRequest
type CreateResponseObjectRequest struct {
	// The name of the response object to create.
	Name *string `json:"name,omitempty"`
	// The status code the response will have. Defaults to 200.
	Status *string `json:"status,omitempty"`
	// The status text the response will have. Defaults to 'OK'.
	Response *string `json:"response,omitempty"`
	// The content the response will deliver.
	Content *string `json:"content,omitempty"`
	// The MIME type of your response content.
	ContentType NullableString `json:"content_type,omitempty"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition NullableString `json:"request_condition,omitempty"`
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition NullableString `json:"cache_condition,omitempty"`
	AdditionalProperties map[string]any
}

type _CreateResponseObjectRequest CreateResponseObjectRequest

// NewCreateResponseObjectRequest instantiates a new CreateResponseObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateResponseObjectRequest() *CreateResponseObjectRequest {
	this := CreateResponseObjectRequest{}
	return &this
}

// NewCreateResponseObjectRequestWithDefaults instantiates a new CreateResponseObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateResponseObjectRequestWithDefaults() *CreateResponseObjectRequest {
	this := CreateResponseObjectRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateResponseObjectRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResponseObjectRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateResponseObjectRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateResponseObjectRequest) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateResponseObjectRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResponseObjectRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateResponseObjectRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateResponseObjectRequest) SetStatus(v string) {
	o.Status = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CreateResponseObjectRequest) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResponseObjectRequest) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CreateResponseObjectRequest) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *CreateResponseObjectRequest) SetResponse(v string) {
	o.Response = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CreateResponseObjectRequest) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResponseObjectRequest) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CreateResponseObjectRequest) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *CreateResponseObjectRequest) SetContent(v string) {
	o.Content = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateResponseObjectRequest) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateResponseObjectRequest) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *CreateResponseObjectRequest) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *CreateResponseObjectRequest) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *CreateResponseObjectRequest) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *CreateResponseObjectRequest) UnsetContentType() {
	o.ContentType.Unset()
}

// GetRequestCondition returns the RequestCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateResponseObjectRequest) GetRequestCondition() string {
	if o == nil || o.RequestCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestCondition.Get()
}

// GetRequestConditionOk returns a tuple with the RequestCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateResponseObjectRequest) GetRequestConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestCondition.Get(), o.RequestCondition.IsSet()
}

// HasRequestCondition returns a boolean if a field has been set.
func (o *CreateResponseObjectRequest) HasRequestCondition() bool {
	if o != nil && o.RequestCondition.IsSet() {
		return true
	}

	return false
}

// SetRequestCondition gets a reference to the given NullableString and assigns it to the RequestCondition field.
func (o *CreateResponseObjectRequest) SetRequestCondition(v string) {
	o.RequestCondition.Set(&v)
}
// SetRequestConditionNil sets the value for RequestCondition to be an explicit nil
func (o *CreateResponseObjectRequest) SetRequestConditionNil() {
	o.RequestCondition.Set(nil)
}

// UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
func (o *CreateResponseObjectRequest) UnsetRequestCondition() {
	o.RequestCondition.Unset()
}

// GetCacheCondition returns the CacheCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateResponseObjectRequest) GetCacheCondition() string {
	if o == nil || o.CacheCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.CacheCondition.Get()
}

// GetCacheConditionOk returns a tuple with the CacheCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateResponseObjectRequest) GetCacheConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CacheCondition.Get(), o.CacheCondition.IsSet()
}

// HasCacheCondition returns a boolean if a field has been set.
func (o *CreateResponseObjectRequest) HasCacheCondition() bool {
	if o != nil && o.CacheCondition.IsSet() {
		return true
	}

	return false
}

// SetCacheCondition gets a reference to the given NullableString and assigns it to the CacheCondition field.
func (o *CreateResponseObjectRequest) SetCacheCondition(v string) {
	o.CacheCondition.Set(&v)
}
// SetCacheConditionNil sets the value for CacheCondition to be an explicit nil
func (o *CreateResponseObjectRequest) SetCacheConditionNil() {
	o.CacheCondition.Set(nil)
}

// UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
func (o *CreateResponseObjectRequest) UnsetCacheCondition() {
	o.CacheCondition.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o CreateResponseObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if o.RequestCondition.IsSet() {
		toSerialize["request_condition"] = o.RequestCondition.Get()
	}
	if o.CacheCondition.IsSet() {
		toSerialize["cache_condition"] = o.CacheCondition.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *CreateResponseObjectRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateResponseObjectRequest := _CreateResponseObjectRequest{}

	if err = json.Unmarshal(bytes, &varCreateResponseObjectRequest); err == nil {
		*o = CreateResponseObjectRequest(varCreateResponseObjectRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "response")
		delete(additionalProperties, "content")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "request_condition")
		delete(additionalProperties, "cache_condition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableCreateResponseObjectRequest is a helper abstraction for handling nullable createresponseobjectrequest types. 
type NullableCreateResponseObjectRequest struct {
	value *CreateResponseObjectRequest
	isSet bool
}

// Get returns the value.
func (v NullableCreateResponseObjectRequest) Get() *CreateResponseObjectRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableCreateResponseObjectRequest) Set(val *CreateResponseObjectRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableCreateResponseObjectRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableCreateResponseObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCreateResponseObjectRequest returns a pointer to a new instance of NullableCreateResponseObjectRequest.
func NewNullableCreateResponseObjectRequest(val *CreateResponseObjectRequest) *NullableCreateResponseObjectRequest {
	return &NullableCreateResponseObjectRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableCreateResponseObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableCreateResponseObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
