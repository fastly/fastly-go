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

// LoggingKafkaAdditional struct for LoggingKafkaAdditional
type LoggingKafkaAdditional struct {
	// The Kafka topic to send logs to. Required.
	Topic *string `json:"topic,omitempty"`
	// A comma-separated list of IP addresses or hostnames of Kafka brokers. Required.
	Brokers *string `json:"brokers,omitempty"`
	// The codec used for compression of your logs.
	CompressionCodec NullableString `json:"compression_codec,omitempty"`
	// The number of acknowledgements a leader must receive before a write is considered successful.
	RequiredAcks *int32 `json:"required_acks,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` (no limit).
	RequestMaxBytes *int32 `json:"request_max_bytes,omitempty"`
	// Enables parsing of key=value tuples from the beginning of a logline, turning them into [record headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers).
	ParseLogKeyvals *bool `json:"parse_log_keyvals,omitempty"`
	// SASL authentication method.
	AuthMethod *string `json:"auth_method,omitempty"`
	// SASL user.
	User *string `json:"user,omitempty"`
	// SASL password.
	Password *string `json:"password,omitempty"`
	UseTLS *LoggingUseTLSString `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingKafkaAdditional LoggingKafkaAdditional

// NewLoggingKafkaAdditional instantiates a new LoggingKafkaAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingKafkaAdditional() *LoggingKafkaAdditional {
	this := LoggingKafkaAdditional{}
	var requiredAcks int32 = 1
	this.RequiredAcks = &requiredAcks
	var requestMaxBytes int32 = 0
	this.RequestMaxBytes = &requestMaxBytes
	var useTLS LoggingUseTLSString = LOGGINGUSETLSSTRING_no_tls
	this.UseTLS = &useTLS
	return &this
}

// NewLoggingKafkaAdditionalWithDefaults instantiates a new LoggingKafkaAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingKafkaAdditionalWithDefaults() *LoggingKafkaAdditional {
	this := LoggingKafkaAdditional{}
	var requiredAcks int32 = 1
	this.RequiredAcks = &requiredAcks
	var requestMaxBytes int32 = 0
	this.RequestMaxBytes = &requestMaxBytes
	var useTLS LoggingUseTLSString = LOGGINGUSETLSSTRING_no_tls
	this.UseTLS = &useTLS
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *LoggingKafkaAdditional) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaAdditional) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *LoggingKafkaAdditional) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *LoggingKafkaAdditional) SetTopic(v string) {
	o.Topic = &v
}

// GetBrokers returns the Brokers field value if set, zero value otherwise.
func (o *LoggingKafkaAdditional) GetBrokers() string {
	if o == nil || o.Brokers == nil {
		var ret string
		return ret
	}
	return *o.Brokers
}

// GetBrokersOk returns a tuple with the Brokers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaAdditional) GetBrokersOk() (*string, bool) {
	if o == nil || o.Brokers == nil {
		return nil, false
	}
	return o.Brokers, true
}

// HasBrokers returns a boolean if a field has been set.
func (o *LoggingKafkaAdditional) HasBrokers() bool {
	if o != nil && o.Brokers != nil {
		return true
	}

	return false
}

// SetBrokers gets a reference to the given string and assigns it to the Brokers field.
func (o *LoggingKafkaAdditional) SetBrokers(v string) {
	o.Brokers = &v
}

// GetCompressionCodec returns the CompressionCodec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingKafkaAdditional) GetCompressionCodec() string {
	if o == nil || o.CompressionCodec.Get() == nil {
		var ret string
		return ret
	}
	return *o.CompressionCodec.Get()
}

// GetCompressionCodecOk returns a tuple with the CompressionCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingKafkaAdditional) GetCompressionCodecOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompressionCodec.Get(), o.CompressionCodec.IsSet()
}

// HasCompressionCodec returns a boolean if a field has been set.
func (o *LoggingKafkaAdditional) HasCompressionCodec() bool {
	if o != nil && o.CompressionCodec.IsSet() {
		return true
	}

	return false
}

// SetCompressionCodec gets a reference to the given NullableString and assigns it to the CompressionCodec field.
func (o *LoggingKafkaAdditional) SetCompressionCodec(v string) {
	o.CompressionCodec.Set(&v)
}
// SetCompressionCodecNil sets the value for CompressionCodec to be an explicit nil
func (o *LoggingKafkaAdditional) SetCompressionCodecNil() {
	o.CompressionCodec.Set(nil)
}

// UnsetCompressionCodec ensures that no value is present for CompressionCodec, not even an explicit nil
func (o *LoggingKafkaAdditional) UnsetCompressionCodec() {
	o.CompressionCodec.Unset()
}

// GetRequiredAcks returns the RequiredAcks field value if set, zero value otherwise.
func (o *LoggingKafkaAdditional) GetRequiredAcks() int32 {
	if o == nil || o.RequiredAcks == nil {
		var ret int32
		return ret
	}
	return *o.RequiredAcks
}

// GetRequiredAcksOk returns a tuple with the RequiredAcks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaAdditional) GetRequiredAcksOk() (*int32, bool) {
	if o == nil || o.RequiredAcks == nil {
		return nil, false
	}
	return o.RequiredAcks, true
}

// HasRequiredAcks returns a boolean if a field has been set.
func (o *LoggingKafkaAdditional) HasRequiredAcks() bool {
	if o != nil && o.RequiredAcks != nil {
		return true
	}

	return false
}

// SetRequiredAcks gets a reference to the given int32 and assigns it to the RequiredAcks field.
func (o *LoggingKafkaAdditional) SetRequiredAcks(v int32) {
	o.RequiredAcks = &v
}

// GetRequestMaxBytes returns the RequestMaxBytes field value if set, zero value otherwise.
func (o *LoggingKafkaAdditional) GetRequestMaxBytes() int32 {
	if o == nil || o.RequestMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxBytes
}

// GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaAdditional) GetRequestMaxBytesOk() (*int32, bool) {
	if o == nil || o.RequestMaxBytes == nil {
		return nil, false
	}
	return o.RequestMaxBytes, true
}

// HasRequestMaxBytes returns a boolean if a field has been set.
func (o *LoggingKafkaAdditional) HasRequestMaxBytes() bool {
	if o != nil && o.RequestMaxBytes != nil {
		return true
	}

	return false
}

// SetRequestMaxBytes gets a reference to the given int32 and assigns it to the RequestMaxBytes field.
func (o *LoggingKafkaAdditional) SetRequestMaxBytes(v int32) {
	o.RequestMaxBytes = &v
}

// GetParseLogKeyvals returns the ParseLogKeyvals field value if set, zero value otherwise.
func (o *LoggingKafkaAdditional) GetParseLogKeyvals() bool {
	if o == nil || o.ParseLogKeyvals == nil {
		var ret bool
		return ret
	}
	return *o.ParseLogKeyvals
}

// GetParseLogKeyvalsOk returns a tuple with the ParseLogKeyvals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaAdditional) GetParseLogKeyvalsOk() (*bool, bool) {
	if o == nil || o.ParseLogKeyvals == nil {
		return nil, false
	}
	return o.ParseLogKeyvals, true
}

// HasParseLogKeyvals returns a boolean if a field has been set.
func (o *LoggingKafkaAdditional) HasParseLogKeyvals() bool {
	if o != nil && o.ParseLogKeyvals != nil {
		return true
	}

	return false
}

// SetParseLogKeyvals gets a reference to the given bool and assigns it to the ParseLogKeyvals field.
func (o *LoggingKafkaAdditional) SetParseLogKeyvals(v bool) {
	o.ParseLogKeyvals = &v
}

// GetAuthMethod returns the AuthMethod field value if set, zero value otherwise.
func (o *LoggingKafkaAdditional) GetAuthMethod() string {
	if o == nil || o.AuthMethod == nil {
		var ret string
		return ret
	}
	return *o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaAdditional) GetAuthMethodOk() (*string, bool) {
	if o == nil || o.AuthMethod == nil {
		return nil, false
	}
	return o.AuthMethod, true
}

// HasAuthMethod returns a boolean if a field has been set.
func (o *LoggingKafkaAdditional) HasAuthMethod() bool {
	if o != nil && o.AuthMethod != nil {
		return true
	}

	return false
}

// SetAuthMethod gets a reference to the given string and assigns it to the AuthMethod field.
func (o *LoggingKafkaAdditional) SetAuthMethod(v string) {
	o.AuthMethod = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoggingKafkaAdditional) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaAdditional) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoggingKafkaAdditional) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *LoggingKafkaAdditional) SetUser(v string) {
	o.User = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LoggingKafkaAdditional) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaAdditional) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LoggingKafkaAdditional) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LoggingKafkaAdditional) SetPassword(v string) {
	o.Password = &v
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *LoggingKafkaAdditional) GetUseTLS() LoggingUseTLSString {
	if o == nil || o.UseTLS == nil {
		var ret LoggingUseTLSString
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingKafkaAdditional) GetUseTLSOk() (*LoggingUseTLSString, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *LoggingKafkaAdditional) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given LoggingUseTLSString and assigns it to the UseTLS field.
func (o *LoggingKafkaAdditional) SetUseTLS(v LoggingUseTLSString) {
	o.UseTLS = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingKafkaAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.Brokers != nil {
		toSerialize["brokers"] = o.Brokers
	}
	if o.CompressionCodec.IsSet() {
		toSerialize["compression_codec"] = o.CompressionCodec.Get()
	}
	if o.RequiredAcks != nil {
		toSerialize["required_acks"] = o.RequiredAcks
	}
	if o.RequestMaxBytes != nil {
		toSerialize["request_max_bytes"] = o.RequestMaxBytes
	}
	if o.ParseLogKeyvals != nil {
		toSerialize["parse_log_keyvals"] = o.ParseLogKeyvals
	}
	if o.AuthMethod != nil {
		toSerialize["auth_method"] = o.AuthMethod
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.UseTLS != nil {
		toSerialize["use_tls"] = o.UseTLS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingKafkaAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingKafkaAdditional := _LoggingKafkaAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingKafkaAdditional); err == nil {
		*o = LoggingKafkaAdditional(varLoggingKafkaAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "topic")
		delete(additionalProperties, "brokers")
		delete(additionalProperties, "compression_codec")
		delete(additionalProperties, "required_acks")
		delete(additionalProperties, "request_max_bytes")
		delete(additionalProperties, "parse_log_keyvals")
		delete(additionalProperties, "auth_method")
		delete(additionalProperties, "user")
		delete(additionalProperties, "password")
		delete(additionalProperties, "use_tls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingKafkaAdditional is a helper abstraction for handling nullable loggingkafkaadditional types. 
type NullableLoggingKafkaAdditional struct {
	value *LoggingKafkaAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingKafkaAdditional) Get() *LoggingKafkaAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingKafkaAdditional) Set(val *LoggingKafkaAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingKafkaAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingKafkaAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingKafkaAdditional returns a pointer to a new instance of NullableLoggingKafkaAdditional.
func NewNullableLoggingKafkaAdditional(val *LoggingKafkaAdditional) *NullableLoggingKafkaAdditional {
	return &NullableLoggingKafkaAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingKafkaAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingKafkaAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
