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

// Content struct for Content
type Content struct {
	Hash *string `json:"hash,omitempty"`
	Request map[string]any `json:"request,omitempty"`
	Response map[string]any `json:"response,omitempty"`
	ResponseTime *float32 `json:"response_time,omitempty"`
	Server *string `json:"server,omitempty"`
	Pop *string `json:"pop,omitempty"`
	AdditionalProperties map[string]any
}

type _Content Content

// NewContent instantiates a new Content object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContent() *Content {
	this := Content{}
	return &this
}

// NewContentWithDefaults instantiates a new Content object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentWithDefaults() *Content {
	this := Content{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *Content) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *Content) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *Content) SetHash(v string) {
	o.Hash = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *Content) GetRequest() map[string]any {
	if o == nil || o.Request == nil {
		var ret map[string]any
		return ret
	}
	return o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetRequestOk() (map[string]any, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *Content) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given map[string]any and assigns it to the Request field.
func (o *Content) SetRequest(v map[string]any) {
	o.Request = v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *Content) GetResponse() map[string]any {
	if o == nil || o.Response == nil {
		var ret map[string]any
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetResponseOk() (map[string]any, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *Content) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given map[string]any and assigns it to the Response field.
func (o *Content) SetResponse(v map[string]any) {
	o.Response = v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *Content) GetResponseTime() float32 {
	if o == nil || o.ResponseTime == nil {
		var ret float32
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetResponseTimeOk() (*float32, bool) {
	if o == nil || o.ResponseTime == nil {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *Content) HasResponseTime() bool {
	if o != nil && o.ResponseTime != nil {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given float32 and assigns it to the ResponseTime field.
func (o *Content) SetResponseTime(v float32) {
	o.ResponseTime = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *Content) GetServer() string {
	if o == nil || o.Server == nil {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetServerOk() (*string, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *Content) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *Content) SetServer(v string) {
	o.Server = &v
}

// GetPop returns the Pop field value if set, zero value otherwise.
func (o *Content) GetPop() string {
	if o == nil || o.Pop == nil {
		var ret string
		return ret
	}
	return *o.Pop
}

// GetPopOk returns a tuple with the Pop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetPopOk() (*string, bool) {
	if o == nil || o.Pop == nil {
		return nil, false
	}
	return o.Pop, true
}

// HasPop returns a boolean if a field has been set.
func (o *Content) HasPop() bool {
	if o != nil && o.Pop != nil {
		return true
	}

	return false
}

// SetPop gets a reference to the given string and assigns it to the Pop field.
func (o *Content) SetPop(v string) {
	o.Pop = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Content) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.ResponseTime != nil {
		toSerialize["response_time"] = o.ResponseTime
	}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	if o.Pop != nil {
		toSerialize["pop"] = o.Pop
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Content) UnmarshalJSON(bytes []byte) (err error) {
	varContent := _Content{}

	if err = json.Unmarshal(bytes, &varContent); err == nil {
		*o = Content(varContent)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "hash")
		delete(additionalProperties, "request")
		delete(additionalProperties, "response")
		delete(additionalProperties, "response_time")
		delete(additionalProperties, "server")
		delete(additionalProperties, "pop")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableContent is a helper abstraction for handling nullable content types. 
type NullableContent struct {
	value *Content
	isSet bool
}

// Get returns the value.
func (v NullableContent) Get() *Content {
	return v.value
}

// Set modifies the value.
func (v *NullableContent) Set(val *Content) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableContent) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableContent) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableContent returns a pointer to a new instance of NullableContent.
func NewNullableContent(val *Content) *NullableContent {
	return &NullableContent{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
