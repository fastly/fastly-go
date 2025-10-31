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

// Healthcheck struct for Healthcheck
type Healthcheck struct {
	// How often to run the health check in milliseconds. Minimum 1 second, maximum 1 hour.
	CheckInterval *int32 `json:"check_interval,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// The status code expected from the host.
	ExpectedResponse *int32 `json:"expected_response,omitempty"`
	// Array of custom headers that will be added to the health check probes.
	Headers []string `json:"headers,omitempty"`
	// Which host to check.
	Host *string `json:"host,omitempty"`
	// Whether to use version 1.0 or 1.1 HTTP.
	HttpVersion *string `json:"http_version,omitempty"`
	// When loading a config, the initial number of probes to be seen as OK.
	Initial *int32 `json:"initial,omitempty"`
	// Which HTTP method to use.
	Method *string `json:"method,omitempty"`
	// The name of the health check.
	Name *string `json:"name,omitempty"`
	// The path to check.
	Path *string `json:"path,omitempty"`
	// How many health checks must succeed to be considered healthy.
	Threshold *int32 `json:"threshold,omitempty"`
	// Timeout in milliseconds.
	Timeout *int32 `json:"timeout,omitempty"`
	// The number of most recent health check queries to keep for this health check.
	Window               *int32 `json:"window,omitempty"`
	AdditionalProperties map[string]any
}

type _Healthcheck Healthcheck

// NewHealthcheck instantiates a new Healthcheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthcheck() *Healthcheck {
	this := Healthcheck{}
	return &this
}

// NewHealthcheckWithDefaults instantiates a new Healthcheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthcheckWithDefaults() *Healthcheck {
	this := Healthcheck{}
	return &this
}

// GetCheckInterval returns the CheckInterval field value if set, zero value otherwise.
func (o *Healthcheck) GetCheckInterval() int32 {
	if o == nil || o.CheckInterval == nil {
		var ret int32
		return ret
	}
	return *o.CheckInterval
}

// GetCheckIntervalOk returns a tuple with the CheckInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetCheckIntervalOk() (*int32, bool) {
	if o == nil || o.CheckInterval == nil {
		return nil, false
	}
	return o.CheckInterval, true
}

// HasCheckInterval returns a boolean if a field has been set.
func (o *Healthcheck) HasCheckInterval() bool {
	if o != nil && o.CheckInterval != nil {
		return true
	}

	return false
}

// SetCheckInterval gets a reference to the given int32 and assigns it to the CheckInterval field.
func (o *Healthcheck) SetCheckInterval(v int32) {
	o.CheckInterval = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Healthcheck) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Healthcheck) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *Healthcheck) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *Healthcheck) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *Healthcheck) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *Healthcheck) UnsetComment() {
	o.Comment.Unset()
}

// GetExpectedResponse returns the ExpectedResponse field value if set, zero value otherwise.
func (o *Healthcheck) GetExpectedResponse() int32 {
	if o == nil || o.ExpectedResponse == nil {
		var ret int32
		return ret
	}
	return *o.ExpectedResponse
}

// GetExpectedResponseOk returns a tuple with the ExpectedResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetExpectedResponseOk() (*int32, bool) {
	if o == nil || o.ExpectedResponse == nil {
		return nil, false
	}
	return o.ExpectedResponse, true
}

// HasExpectedResponse returns a boolean if a field has been set.
func (o *Healthcheck) HasExpectedResponse() bool {
	if o != nil && o.ExpectedResponse != nil {
		return true
	}

	return false
}

// SetExpectedResponse gets a reference to the given int32 and assigns it to the ExpectedResponse field.
func (o *Healthcheck) SetExpectedResponse(v int32) {
	o.ExpectedResponse = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *Healthcheck) GetHeaders() []string {
	if o == nil || o.Headers == nil {
		var ret []string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetHeadersOk() ([]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *Healthcheck) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []string and assigns it to the Headers field.
func (o *Healthcheck) SetHeaders(v []string) {
	o.Headers = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Healthcheck) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Healthcheck) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Healthcheck) SetHost(v string) {
	o.Host = &v
}

// GetHttpVersion returns the HttpVersion field value if set, zero value otherwise.
func (o *Healthcheck) GetHttpVersion() string {
	if o == nil || o.HttpVersion == nil {
		var ret string
		return ret
	}
	return *o.HttpVersion
}

// GetHttpVersionOk returns a tuple with the HttpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetHttpVersionOk() (*string, bool) {
	if o == nil || o.HttpVersion == nil {
		return nil, false
	}
	return o.HttpVersion, true
}

// HasHttpVersion returns a boolean if a field has been set.
func (o *Healthcheck) HasHttpVersion() bool {
	if o != nil && o.HttpVersion != nil {
		return true
	}

	return false
}

// SetHttpVersion gets a reference to the given string and assigns it to the HttpVersion field.
func (o *Healthcheck) SetHttpVersion(v string) {
	o.HttpVersion = &v
}

// GetInitial returns the Initial field value if set, zero value otherwise.
func (o *Healthcheck) GetInitial() int32 {
	if o == nil || o.Initial == nil {
		var ret int32
		return ret
	}
	return *o.Initial
}

// GetInitialOk returns a tuple with the Initial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetInitialOk() (*int32, bool) {
	if o == nil || o.Initial == nil {
		return nil, false
	}
	return o.Initial, true
}

// HasInitial returns a boolean if a field has been set.
func (o *Healthcheck) HasInitial() bool {
	if o != nil && o.Initial != nil {
		return true
	}

	return false
}

// SetInitial gets a reference to the given int32 and assigns it to the Initial field.
func (o *Healthcheck) SetInitial(v int32) {
	o.Initial = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *Healthcheck) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *Healthcheck) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *Healthcheck) SetMethod(v string) {
	o.Method = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Healthcheck) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Healthcheck) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Healthcheck) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Healthcheck) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Healthcheck) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Healthcheck) SetPath(v string) {
	o.Path = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *Healthcheck) GetThreshold() int32 {
	if o == nil || o.Threshold == nil {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetThresholdOk() (*int32, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *Healthcheck) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *Healthcheck) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *Healthcheck) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *Healthcheck) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *Healthcheck) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *Healthcheck) GetWindow() int32 {
	if o == nil || o.Window == nil {
		var ret int32
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Healthcheck) GetWindowOk() (*int32, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *Healthcheck) HasWindow() bool {
	if o != nil && o.Window != nil {
		return true
	}

	return false
}

// SetWindow gets a reference to the given int32 and assigns it to the Window field.
func (o *Healthcheck) SetWindow(v int32) {
	o.Window = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Healthcheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CheckInterval != nil {
		toSerialize["check_interval"] = o.CheckInterval
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.ExpectedResponse != nil {
		toSerialize["expected_response"] = o.ExpectedResponse
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.HttpVersion != nil {
		toSerialize["http_version"] = o.HttpVersion
	}
	if o.Initial != nil {
		toSerialize["initial"] = o.Initial
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Window != nil {
		toSerialize["window"] = o.Window
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Healthcheck) UnmarshalJSON(bytes []byte) (err error) {
	varHealthcheck := _Healthcheck{}

	if err = json.Unmarshal(bytes, &varHealthcheck); err == nil {
		*o = Healthcheck(varHealthcheck)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "check_interval")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "expected_response")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "host")
		delete(additionalProperties, "http_version")
		delete(additionalProperties, "initial")
		delete(additionalProperties, "method")
		delete(additionalProperties, "name")
		delete(additionalProperties, "path")
		delete(additionalProperties, "threshold")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "window")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHealthcheck is a helper abstraction for handling nullable healthcheck types.
type NullableHealthcheck struct {
	value *Healthcheck
	isSet bool
}

// Get returns the value.
func (v NullableHealthcheck) Get() *Healthcheck {
	return v.value
}

// Set modifies the value.
func (v *NullableHealthcheck) Set(val *Healthcheck) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHealthcheck) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHealthcheck) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHealthcheck returns a pointer to a new instance of NullableHealthcheck.
func NewNullableHealthcheck(val *Healthcheck) *NullableHealthcheck {
	return &NullableHealthcheck{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHealthcheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableHealthcheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
