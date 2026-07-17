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

// WafSimulateRequest Request body for simulating a WAF request. The total request body must not exceed 200 KB.
type WafSimulateRequest struct {
	// The raw HTTP request in wire format to simulate through the WAF. Must include the request line, headers, and optionally a body, separated by CRLF sequences.
	Request string `json:"request"`
	// The raw HTTP response in wire format. The WAF engine inspects response headers during its PostRequest phase and may generate signals from them. When omitted, a default response of `HTTP/1.1 200 OK\\r\\n\\r\\n` is used.
	Response             *string `json:"response,omitempty"`
	AdditionalProperties map[string]any
}

type _WafSimulateRequest WafSimulateRequest

// NewWafSimulateRequest instantiates a new WafSimulateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafSimulateRequest(request string) *WafSimulateRequest {
	this := WafSimulateRequest{}
	this.Request = request
	return &this
}

// NewWafSimulateRequestWithDefaults instantiates a new WafSimulateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafSimulateRequestWithDefaults() *WafSimulateRequest {
	this := WafSimulateRequest{}
	return &this
}

// GetRequest returns the Request field value
func (o *WafSimulateRequest) GetRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *WafSimulateRequest) GetRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *WafSimulateRequest) SetRequest(v string) {
	o.Request = v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *WafSimulateRequest) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSimulateRequest) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *WafSimulateRequest) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *WafSimulateRequest) SetResponse(v string) {
	o.Response = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafSimulateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafSimulateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varWafSimulateRequest := _WafSimulateRequest{}

	if err = json.Unmarshal(bytes, &varWafSimulateRequest); err == nil {
		*o = WafSimulateRequest(varWafSimulateRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request")
		delete(additionalProperties, "response")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafSimulateRequest is a helper abstraction for handling nullable wafsimulaterequest types.
type NullableWafSimulateRequest struct {
	value *WafSimulateRequest
	isSet bool
}

// Get returns the value.
func (v NullableWafSimulateRequest) Get() *WafSimulateRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableWafSimulateRequest) Set(val *WafSimulateRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafSimulateRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafSimulateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafSimulateRequest returns a pointer to a new instance of NullableWafSimulateRequest.
func NewNullableWafSimulateRequest(val *WafSimulateRequest) *NullableWafSimulateRequest {
	return &NullableWafSimulateRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafSimulateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafSimulateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
