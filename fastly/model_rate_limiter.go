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

// RateLimiter struct for RateLimiter
type RateLimiter struct {
	// A human readable name for the rate limiting rule.
	Name *string `json:"name,omitempty"`
	// The name of an Edge Dictionary containing URIs as keys. If not defined or `null`, all origin URIs will be rate limited.
	URIDictionaryName NullableString `json:"uri_dictionary_name,omitempty"`
	// Array of HTTP methods to apply rate limiting to.
	HTTPMethods []string `json:"http_methods,omitempty"`
	// Upper limit of requests per second allowed by the rate limiter.
	RpsLimit *int32 `json:"rps_limit,omitempty"`
	// Number of seconds during which the RPS limit must be exceeded in order to trigger a violation.
	WindowSize *int32 `json:"window_size,omitempty"`
	// Array of VCL variables used to generate a counter key to identify a client. Example variables include `req.http.Fastly-Client-IP`, `req.http.User-Agent`, or a custom header like `req.http.API-Key`.
	ClientKey []string `json:"client_key,omitempty"`
	// Length of time in minutes that the rate limiter is in effect after the initial violation is detected.
	PenaltyBoxDuration *int32 `json:"penalty_box_duration,omitempty"`
	// The action to take when a rate limiter violation is detected.
	Action *string `json:"action,omitempty"`
	// Custom response to be sent when the rate limit is exceeded. Required if `action` is `response`.
	Response map[string]string `json:"response,omitempty"`
	// Name of existing response object. Required if `action` is `response_object`. Note that the rate limiter response is only updated to reflect the response object content when saving the rate limiter configuration.
	ResponseObjectName NullableString `json:"response_object_name,omitempty"`
	// Name of the type of logging endpoint to be used when action is `log_only`. The logging endpoint type is used to determine the appropriate log format to use when emitting log entries.
	LoggerType *string `json:"logger_type,omitempty"`
	// Revision number of the rate limiting feature implementation. Defaults to the most recent revision.
	FeatureRevision *int32 `json:"feature_revision,omitempty"`
	AdditionalProperties map[string]any
}

type _RateLimiter RateLimiter

// NewRateLimiter instantiates a new RateLimiter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimiter() *RateLimiter {
	this := RateLimiter{}
	return &this
}

// NewRateLimiterWithDefaults instantiates a new RateLimiter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimiterWithDefaults() *RateLimiter {
	this := RateLimiter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RateLimiter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RateLimiter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RateLimiter) SetName(v string) {
	o.Name = &v
}

// GetURIDictionaryName returns the URIDictionaryName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RateLimiter) GetURIDictionaryName() string {
	if o == nil || o.URIDictionaryName.Get() == nil {
		var ret string
		return ret
	}
	return *o.URIDictionaryName.Get()
}

// GetURIDictionaryNameOk returns a tuple with the URIDictionaryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RateLimiter) GetURIDictionaryNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.URIDictionaryName.Get(), o.URIDictionaryName.IsSet()
}

// HasURIDictionaryName returns a boolean if a field has been set.
func (o *RateLimiter) HasURIDictionaryName() bool {
	if o != nil && o.URIDictionaryName.IsSet() {
		return true
	}

	return false
}

// SetURIDictionaryName gets a reference to the given NullableString and assigns it to the URIDictionaryName field.
func (o *RateLimiter) SetURIDictionaryName(v string) {
	o.URIDictionaryName.Set(&v)
}
// SetURIDictionaryNameNil sets the value for URIDictionaryName to be an explicit nil
func (o *RateLimiter) SetURIDictionaryNameNil() {
	o.URIDictionaryName.Set(nil)
}

// UnsetURIDictionaryName ensures that no value is present for URIDictionaryName, not even an explicit nil
func (o *RateLimiter) UnsetURIDictionaryName() {
	o.URIDictionaryName.Unset()
}

// GetHTTPMethods returns the HTTPMethods field value if set, zero value otherwise.
func (o *RateLimiter) GetHTTPMethods() []string {
	if o == nil || o.HTTPMethods == nil {
		var ret []string
		return ret
	}
	return o.HTTPMethods
}

// GetHTTPMethodsOk returns a tuple with the HTTPMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiter) GetHTTPMethodsOk() ([]string, bool) {
	if o == nil || o.HTTPMethods == nil {
		return nil, false
	}
	return o.HTTPMethods, true
}

// HasHTTPMethods returns a boolean if a field has been set.
func (o *RateLimiter) HasHTTPMethods() bool {
	if o != nil && o.HTTPMethods != nil {
		return true
	}

	return false
}

// SetHTTPMethods gets a reference to the given []string and assigns it to the HTTPMethods field.
func (o *RateLimiter) SetHTTPMethods(v []string) {
	o.HTTPMethods = v
}

// GetRpsLimit returns the RpsLimit field value if set, zero value otherwise.
func (o *RateLimiter) GetRpsLimit() int32 {
	if o == nil || o.RpsLimit == nil {
		var ret int32
		return ret
	}
	return *o.RpsLimit
}

// GetRpsLimitOk returns a tuple with the RpsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiter) GetRpsLimitOk() (*int32, bool) {
	if o == nil || o.RpsLimit == nil {
		return nil, false
	}
	return o.RpsLimit, true
}

// HasRpsLimit returns a boolean if a field has been set.
func (o *RateLimiter) HasRpsLimit() bool {
	if o != nil && o.RpsLimit != nil {
		return true
	}

	return false
}

// SetRpsLimit gets a reference to the given int32 and assigns it to the RpsLimit field.
func (o *RateLimiter) SetRpsLimit(v int32) {
	o.RpsLimit = &v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *RateLimiter) GetWindowSize() int32 {
	if o == nil || o.WindowSize == nil {
		var ret int32
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiter) GetWindowSizeOk() (*int32, bool) {
	if o == nil || o.WindowSize == nil {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *RateLimiter) HasWindowSize() bool {
	if o != nil && o.WindowSize != nil {
		return true
	}

	return false
}

// SetWindowSize gets a reference to the given int32 and assigns it to the WindowSize field.
func (o *RateLimiter) SetWindowSize(v int32) {
	o.WindowSize = &v
}

// GetClientKey returns the ClientKey field value if set, zero value otherwise.
func (o *RateLimiter) GetClientKey() []string {
	if o == nil || o.ClientKey == nil {
		var ret []string
		return ret
	}
	return o.ClientKey
}

// GetClientKeyOk returns a tuple with the ClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiter) GetClientKeyOk() ([]string, bool) {
	if o == nil || o.ClientKey == nil {
		return nil, false
	}
	return o.ClientKey, true
}

// HasClientKey returns a boolean if a field has been set.
func (o *RateLimiter) HasClientKey() bool {
	if o != nil && o.ClientKey != nil {
		return true
	}

	return false
}

// SetClientKey gets a reference to the given []string and assigns it to the ClientKey field.
func (o *RateLimiter) SetClientKey(v []string) {
	o.ClientKey = v
}

// GetPenaltyBoxDuration returns the PenaltyBoxDuration field value if set, zero value otherwise.
func (o *RateLimiter) GetPenaltyBoxDuration() int32 {
	if o == nil || o.PenaltyBoxDuration == nil {
		var ret int32
		return ret
	}
	return *o.PenaltyBoxDuration
}

// GetPenaltyBoxDurationOk returns a tuple with the PenaltyBoxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiter) GetPenaltyBoxDurationOk() (*int32, bool) {
	if o == nil || o.PenaltyBoxDuration == nil {
		return nil, false
	}
	return o.PenaltyBoxDuration, true
}

// HasPenaltyBoxDuration returns a boolean if a field has been set.
func (o *RateLimiter) HasPenaltyBoxDuration() bool {
	if o != nil && o.PenaltyBoxDuration != nil {
		return true
	}

	return false
}

// SetPenaltyBoxDuration gets a reference to the given int32 and assigns it to the PenaltyBoxDuration field.
func (o *RateLimiter) SetPenaltyBoxDuration(v int32) {
	o.PenaltyBoxDuration = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RateLimiter) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiter) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RateLimiter) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RateLimiter) SetAction(v string) {
	o.Action = &v
}

// GetResponse returns the Response field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RateLimiter) GetResponse() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RateLimiter) GetResponseOk() (*map[string]string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return &o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *RateLimiter) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given map[string]string and assigns it to the Response field.
func (o *RateLimiter) SetResponse(v map[string]string) {
	o.Response = v
}

// GetResponseObjectName returns the ResponseObjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RateLimiter) GetResponseObjectName() string {
	if o == nil || o.ResponseObjectName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseObjectName.Get()
}

// GetResponseObjectNameOk returns a tuple with the ResponseObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RateLimiter) GetResponseObjectNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseObjectName.Get(), o.ResponseObjectName.IsSet()
}

// HasResponseObjectName returns a boolean if a field has been set.
func (o *RateLimiter) HasResponseObjectName() bool {
	if o != nil && o.ResponseObjectName.IsSet() {
		return true
	}

	return false
}

// SetResponseObjectName gets a reference to the given NullableString and assigns it to the ResponseObjectName field.
func (o *RateLimiter) SetResponseObjectName(v string) {
	o.ResponseObjectName.Set(&v)
}
// SetResponseObjectNameNil sets the value for ResponseObjectName to be an explicit nil
func (o *RateLimiter) SetResponseObjectNameNil() {
	o.ResponseObjectName.Set(nil)
}

// UnsetResponseObjectName ensures that no value is present for ResponseObjectName, not even an explicit nil
func (o *RateLimiter) UnsetResponseObjectName() {
	o.ResponseObjectName.Unset()
}

// GetLoggerType returns the LoggerType field value if set, zero value otherwise.
func (o *RateLimiter) GetLoggerType() string {
	if o == nil || o.LoggerType == nil {
		var ret string
		return ret
	}
	return *o.LoggerType
}

// GetLoggerTypeOk returns a tuple with the LoggerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiter) GetLoggerTypeOk() (*string, bool) {
	if o == nil || o.LoggerType == nil {
		return nil, false
	}
	return o.LoggerType, true
}

// HasLoggerType returns a boolean if a field has been set.
func (o *RateLimiter) HasLoggerType() bool {
	if o != nil && o.LoggerType != nil {
		return true
	}

	return false
}

// SetLoggerType gets a reference to the given string and assigns it to the LoggerType field.
func (o *RateLimiter) SetLoggerType(v string) {
	o.LoggerType = &v
}

// GetFeatureRevision returns the FeatureRevision field value if set, zero value otherwise.
func (o *RateLimiter) GetFeatureRevision() int32 {
	if o == nil || o.FeatureRevision == nil {
		var ret int32
		return ret
	}
	return *o.FeatureRevision
}

// GetFeatureRevisionOk returns a tuple with the FeatureRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiter) GetFeatureRevisionOk() (*int32, bool) {
	if o == nil || o.FeatureRevision == nil {
		return nil, false
	}
	return o.FeatureRevision, true
}

// HasFeatureRevision returns a boolean if a field has been set.
func (o *RateLimiter) HasFeatureRevision() bool {
	if o != nil && o.FeatureRevision != nil {
		return true
	}

	return false
}

// SetFeatureRevision gets a reference to the given int32 and assigns it to the FeatureRevision field.
func (o *RateLimiter) SetFeatureRevision(v int32) {
	o.FeatureRevision = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RateLimiter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.URIDictionaryName.IsSet() {
		toSerialize["uri_dictionary_name"] = o.URIDictionaryName.Get()
	}
	if o.HTTPMethods != nil {
		toSerialize["http_methods"] = o.HTTPMethods
	}
	if o.RpsLimit != nil {
		toSerialize["rps_limit"] = o.RpsLimit
	}
	if o.WindowSize != nil {
		toSerialize["window_size"] = o.WindowSize
	}
	if o.ClientKey != nil {
		toSerialize["client_key"] = o.ClientKey
	}
	if o.PenaltyBoxDuration != nil {
		toSerialize["penalty_box_duration"] = o.PenaltyBoxDuration
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.ResponseObjectName.IsSet() {
		toSerialize["response_object_name"] = o.ResponseObjectName.Get()
	}
	if o.LoggerType != nil {
		toSerialize["logger_type"] = o.LoggerType
	}
	if o.FeatureRevision != nil {
		toSerialize["feature_revision"] = o.FeatureRevision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RateLimiter) UnmarshalJSON(bytes []byte) (err error) {
	varRateLimiter := _RateLimiter{}

	if err = json.Unmarshal(bytes, &varRateLimiter); err == nil {
		*o = RateLimiter(varRateLimiter)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "uri_dictionary_name")
		delete(additionalProperties, "http_methods")
		delete(additionalProperties, "rps_limit")
		delete(additionalProperties, "window_size")
		delete(additionalProperties, "client_key")
		delete(additionalProperties, "penalty_box_duration")
		delete(additionalProperties, "action")
		delete(additionalProperties, "response")
		delete(additionalProperties, "response_object_name")
		delete(additionalProperties, "logger_type")
		delete(additionalProperties, "feature_revision")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRateLimiter is a helper abstraction for handling nullable ratelimiter types. 
type NullableRateLimiter struct {
	value *RateLimiter
	isSet bool
}

// Get returns the value.
func (v NullableRateLimiter) Get() *RateLimiter {
	return v.value
}

// Set modifies the value.
func (v *NullableRateLimiter) Set(val *RateLimiter) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRateLimiter) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRateLimiter) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRateLimiter returns a pointer to a new instance of NullableRateLimiter.
func NewNullableRateLimiter(val *RateLimiter) *NullableRateLimiter {
	return &NullableRateLimiter{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRateLimiter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRateLimiter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
