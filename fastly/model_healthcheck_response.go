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
	"time"
)

// HealthcheckResponse struct for HealthcheckResponse
type HealthcheckResponse struct {
	// How often to run the health check in milliseconds.
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
	HTTPVersion *string `json:"http_version,omitempty"`
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
	Window *int32 `json:"window,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	Version *int32 `json:"version,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _HealthcheckResponse HealthcheckResponse

// NewHealthcheckResponse instantiates a new HealthcheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthcheckResponse() *HealthcheckResponse {
	this := HealthcheckResponse{}
	return &this
}

// NewHealthcheckResponseWithDefaults instantiates a new HealthcheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthcheckResponseWithDefaults() *HealthcheckResponse {
	this := HealthcheckResponse{}
	return &this
}

// GetCheckInterval returns the CheckInterval field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetCheckInterval() int32 {
	if o == nil || o.CheckInterval == nil {
		var ret int32
		return ret
	}
	return *o.CheckInterval
}

// GetCheckIntervalOk returns a tuple with the CheckInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetCheckIntervalOk() (*int32, bool) {
	if o == nil || o.CheckInterval == nil {
		return nil, false
	}
	return o.CheckInterval, true
}

// HasCheckInterval returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasCheckInterval() bool {
	if o != nil && o.CheckInterval != nil {
		return true
	}

	return false
}

// SetCheckInterval gets a reference to the given int32 and assigns it to the CheckInterval field.
func (o *HealthcheckResponse) SetCheckInterval(v int32) {
	o.CheckInterval = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthcheckResponse) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthcheckResponse) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *HealthcheckResponse) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *HealthcheckResponse) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *HealthcheckResponse) UnsetComment() {
	o.Comment.Unset()
}

// GetExpectedResponse returns the ExpectedResponse field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetExpectedResponse() int32 {
	if o == nil || o.ExpectedResponse == nil {
		var ret int32
		return ret
	}
	return *o.ExpectedResponse
}

// GetExpectedResponseOk returns a tuple with the ExpectedResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetExpectedResponseOk() (*int32, bool) {
	if o == nil || o.ExpectedResponse == nil {
		return nil, false
	}
	return o.ExpectedResponse, true
}

// HasExpectedResponse returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasExpectedResponse() bool {
	if o != nil && o.ExpectedResponse != nil {
		return true
	}

	return false
}

// SetExpectedResponse gets a reference to the given int32 and assigns it to the ExpectedResponse field.
func (o *HealthcheckResponse) SetExpectedResponse(v int32) {
	o.ExpectedResponse = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetHeaders() []string {
	if o == nil || o.Headers == nil {
		var ret []string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetHeadersOk() ([]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []string and assigns it to the Headers field.
func (o *HealthcheckResponse) SetHeaders(v []string) {
	o.Headers = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *HealthcheckResponse) SetHost(v string) {
	o.Host = &v
}

// GetHTTPVersion returns the HTTPVersion field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetHTTPVersion() string {
	if o == nil || o.HTTPVersion == nil {
		var ret string
		return ret
	}
	return *o.HTTPVersion
}

// GetHTTPVersionOk returns a tuple with the HTTPVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetHTTPVersionOk() (*string, bool) {
	if o == nil || o.HTTPVersion == nil {
		return nil, false
	}
	return o.HTTPVersion, true
}

// HasHTTPVersion returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasHTTPVersion() bool {
	if o != nil && o.HTTPVersion != nil {
		return true
	}

	return false
}

// SetHTTPVersion gets a reference to the given string and assigns it to the HTTPVersion field.
func (o *HealthcheckResponse) SetHTTPVersion(v string) {
	o.HTTPVersion = &v
}

// GetInitial returns the Initial field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetInitial() int32 {
	if o == nil || o.Initial == nil {
		var ret int32
		return ret
	}
	return *o.Initial
}

// GetInitialOk returns a tuple with the Initial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetInitialOk() (*int32, bool) {
	if o == nil || o.Initial == nil {
		return nil, false
	}
	return o.Initial, true
}

// HasInitial returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasInitial() bool {
	if o != nil && o.Initial != nil {
		return true
	}

	return false
}

// SetInitial gets a reference to the given int32 and assigns it to the Initial field.
func (o *HealthcheckResponse) SetInitial(v int32) {
	o.Initial = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *HealthcheckResponse) SetMethod(v string) {
	o.Method = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HealthcheckResponse) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HealthcheckResponse) SetPath(v string) {
	o.Path = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetThreshold() int32 {
	if o == nil || o.Threshold == nil {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetThresholdOk() (*int32, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *HealthcheckResponse) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *HealthcheckResponse) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetWindow() int32 {
	if o == nil || o.Window == nil {
		var ret int32
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetWindowOk() (*int32, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasWindow() bool {
	if o != nil && o.Window != nil {
		return true
	}

	return false
}

// SetWindow gets a reference to the given int32 and assigns it to the Window field.
func (o *HealthcheckResponse) SetWindow(v int32) {
	o.Window = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *HealthcheckResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HealthcheckResponse) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckResponse) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *HealthcheckResponse) SetVersion(v int32) {
	o.Version = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthcheckResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthcheckResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *HealthcheckResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *HealthcheckResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *HealthcheckResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthcheckResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthcheckResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *HealthcheckResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *HealthcheckResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *HealthcheckResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthcheckResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthcheckResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *HealthcheckResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *HealthcheckResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *HealthcheckResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *HealthcheckResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HealthcheckResponse) MarshalJSON() ([]byte, error) {
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
	if o.HTTPVersion != nil {
		toSerialize["http_version"] = o.HTTPVersion
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
func (o *HealthcheckResponse) UnmarshalJSON(bytes []byte) (err error) {
	varHealthcheckResponse := _HealthcheckResponse{}

	if err = json.Unmarshal(bytes, &varHealthcheckResponse); err == nil {
		*o = HealthcheckResponse(varHealthcheckResponse)
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
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHealthcheckResponse is a helper abstraction for handling nullable healthcheckresponse types. 
type NullableHealthcheckResponse struct {
	value *HealthcheckResponse
	isSet bool
}

// Get returns the value.
func (v NullableHealthcheckResponse) Get() *HealthcheckResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableHealthcheckResponse) Set(val *HealthcheckResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHealthcheckResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHealthcheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHealthcheckResponse returns a pointer to a new instance of NullableHealthcheckResponse.
func NewNullableHealthcheckResponse(val *HealthcheckResponse) *NullableHealthcheckResponse {
	return &NullableHealthcheckResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHealthcheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHealthcheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
