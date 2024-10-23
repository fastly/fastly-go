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

// Server struct for Server
type Server struct {
	// Weight (`1-100`) used to load balance this server against others.
	Weight *int32 `json:"weight,omitempty"`
	// Maximum number of connections. If the value is `0`, it inherits the value from pool's `max_conn_default`.
	MaxConn *int32 `json:"max_conn,omitempty"`
	// Port number. Setting port `443` does not force TLS. Set `use_tls` in pool to force TLS.
	Port *int32 `json:"port,omitempty"`
	// A hostname, IPv4, or IPv6 address for the server. Required.
	Address *string `json:"address,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// Allows servers to be enabled and disabled in a pool.
	Disabled *bool `json:"disabled,omitempty"`
	// The hostname to override the Host header. Defaults to `null` meaning no override of the Host header if not set. This setting can also be added to a Pool definition. However, the server setting will override the Pool setting.
	OverrideHost         NullableString `json:"override_host,omitempty"`
	AdditionalProperties map[string]any
}

type _Server Server

// NewServer instantiates a new Server object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServer() *Server {
	this := Server{}
	var weight int32 = 100
	this.Weight = &weight
	var maxConn int32 = 0
	this.MaxConn = &maxConn
	var port int32 = 80
	this.Port = &port
	var disabled bool = false
	this.Disabled = &disabled
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	return &this
}

// NewServerWithDefaults instantiates a new Server object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerWithDefaults() *Server {
	this := Server{}
	var weight int32 = 100
	this.Weight = &weight
	var maxConn int32 = 0
	this.MaxConn = &maxConn
	var port int32 = 80
	this.Port = &port
	var disabled bool = false
	this.Disabled = &disabled
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	return &this
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *Server) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Server) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *Server) SetWeight(v int32) {
	o.Weight = &v
}

// GetMaxConn returns the MaxConn field value if set, zero value otherwise.
func (o *Server) GetMaxConn() int32 {
	if o == nil || o.MaxConn == nil {
		var ret int32
		return ret
	}
	return *o.MaxConn
}

// GetMaxConnOk returns a tuple with the MaxConn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetMaxConnOk() (*int32, bool) {
	if o == nil || o.MaxConn == nil {
		return nil, false
	}
	return o.MaxConn, true
}

// HasMaxConn returns a boolean if a field has been set.
func (o *Server) HasMaxConn() bool {
	if o != nil && o.MaxConn != nil {
		return true
	}

	return false
}

// SetMaxConn gets a reference to the given int32 and assigns it to the MaxConn field.
func (o *Server) SetMaxConn(v int32) {
	o.MaxConn = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Server) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Server) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *Server) SetPort(v int32) {
	o.Port = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Server) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Server) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Server) SetAddress(v string) {
	o.Address = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Server) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Server) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *Server) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *Server) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *Server) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *Server) UnsetComment() {
	o.Comment.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *Server) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *Server) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *Server) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetOverrideHost returns the OverrideHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Server) GetOverrideHost() string {
	if o == nil || o.OverrideHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideHost.Get()
}

// GetOverrideHostOk returns a tuple with the OverrideHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Server) GetOverrideHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverrideHost.Get(), o.OverrideHost.IsSet()
}

// HasOverrideHost returns a boolean if a field has been set.
func (o *Server) HasOverrideHost() bool {
	if o != nil && o.OverrideHost.IsSet() {
		return true
	}

	return false
}

// SetOverrideHost gets a reference to the given NullableString and assigns it to the OverrideHost field.
func (o *Server) SetOverrideHost(v string) {
	o.OverrideHost.Set(&v)
}

// SetOverrideHostNil sets the value for OverrideHost to be an explicit nil
func (o *Server) SetOverrideHostNil() {
	o.OverrideHost.Set(nil)
}

// UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
func (o *Server) UnsetOverrideHost() {
	o.OverrideHost.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Server) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.MaxConn != nil {
		toSerialize["max_conn"] = o.MaxConn
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.OverrideHost.IsSet() {
		toSerialize["override_host"] = o.OverrideHost.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Server) UnmarshalJSON(bytes []byte) (err error) {
	varServer := _Server{}

	if err = json.Unmarshal(bytes, &varServer); err == nil {
		*o = Server(varServer)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "weight")
		delete(additionalProperties, "max_conn")
		delete(additionalProperties, "port")
		delete(additionalProperties, "address")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "override_host")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServer is a helper abstraction for handling nullable server types.
type NullableServer struct {
	value *Server
	isSet bool
}

// Get returns the value.
func (v NullableServer) Get() *Server {
	return v.value
}

// Set modifies the value.
func (v *NullableServer) Set(val *Server) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServer) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServer) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServer returns a pointer to a new instance of NullableServer.
func NewNullableServer(val *Server) *NullableServer {
	return &NullableServer{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
