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

// LoggingScalyrAllOf struct for LoggingScalyrAllOf
type LoggingScalyrAllOf struct {
	// The region that log data will be sent to.
	Region *string `json:"region,omitempty"`
	// The token to use for authentication ([https://www.scalyr.com/keys](https://www.scalyr.com/keys)).
	Token *string `json:"token,omitempty"`
	// The name of the logfile within Scalyr.
	ProjectID *string `json:"project_id,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingScalyrAllOf LoggingScalyrAllOf

// NewLoggingScalyrAllOf instantiates a new LoggingScalyrAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingScalyrAllOf() *LoggingScalyrAllOf {
	this := LoggingScalyrAllOf{}
	var region string = "US"
	this.Region = &region
	var projectID string = "logplex"
	this.ProjectID = &projectID
	return &this
}

// NewLoggingScalyrAllOfWithDefaults instantiates a new LoggingScalyrAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingScalyrAllOfWithDefaults() *LoggingScalyrAllOf {
	this := LoggingScalyrAllOf{}
	var region string = "US"
	this.Region = &region
	var projectID string = "logplex"
	this.ProjectID = &projectID
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *LoggingScalyrAllOf) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingScalyrAllOf) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *LoggingScalyrAllOf) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *LoggingScalyrAllOf) SetRegion(v string) {
	o.Region = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoggingScalyrAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingScalyrAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoggingScalyrAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoggingScalyrAllOf) SetToken(v string) {
	o.Token = &v
}

// GetProjectID returns the ProjectID field value if set, zero value otherwise.
func (o *LoggingScalyrAllOf) GetProjectID() string {
	if o == nil || o.ProjectID == nil {
		var ret string
		return ret
	}
	return *o.ProjectID
}

// GetProjectIDOk returns a tuple with the ProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingScalyrAllOf) GetProjectIDOk() (*string, bool) {
	if o == nil || o.ProjectID == nil {
		return nil, false
	}
	return o.ProjectID, true
}

// HasProjectID returns a boolean if a field has been set.
func (o *LoggingScalyrAllOf) HasProjectID() bool {
	if o != nil && o.ProjectID != nil {
		return true
	}

	return false
}

// SetProjectID gets a reference to the given string and assigns it to the ProjectID field.
func (o *LoggingScalyrAllOf) SetProjectID(v string) {
	o.ProjectID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingScalyrAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.ProjectID != nil {
		toSerialize["project_id"] = o.ProjectID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingScalyrAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingScalyrAllOf := _LoggingScalyrAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingScalyrAllOf); err == nil {
		*o = LoggingScalyrAllOf(varLoggingScalyrAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "region")
		delete(additionalProperties, "token")
		delete(additionalProperties, "project_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingScalyrAllOf is a helper abstraction for handling nullable loggingscalyrallof types. 
type NullableLoggingScalyrAllOf struct {
	value *LoggingScalyrAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingScalyrAllOf) Get() *LoggingScalyrAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingScalyrAllOf) Set(val *LoggingScalyrAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingScalyrAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingScalyrAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingScalyrAllOf returns a pointer to a new instance of NullableLoggingScalyrAllOf.
func NewNullableLoggingScalyrAllOf(val *LoggingScalyrAllOf) *NullableLoggingScalyrAllOf {
	return &NullableLoggingScalyrAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingScalyrAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingScalyrAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
