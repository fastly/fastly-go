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

// Snippet struct for Snippet
type Snippet struct {
	// The name for the snippet.
	Name *string `json:"name,omitempty"`
	// The location in generated VCL where the snippet should be placed.
	Type *string `json:"type,omitempty"`
	// The VCL code that specifies exactly what the snippet does.
	Content *string `json:"content,omitempty"`
	// Priority determines execution order. Lower numbers execute first.
	Priority *string `json:"priority,omitempty"`
	// Sets the snippet version.
	Dynamic *string `json:"dynamic,omitempty"`
	AdditionalProperties map[string]any
}

type _Snippet Snippet

// NewSnippet instantiates a new Snippet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnippet() *Snippet {
	this := Snippet{}
	var priority string = "100"
	this.Priority = &priority
	return &this
}

// NewSnippetWithDefaults instantiates a new Snippet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnippetWithDefaults() *Snippet {
	this := Snippet{}
	var priority string = "100"
	this.Priority = &priority
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Snippet) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snippet) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Snippet) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Snippet) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Snippet) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snippet) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Snippet) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Snippet) SetType(v string) {
	o.Type = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Snippet) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snippet) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Snippet) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Snippet) SetContent(v string) {
	o.Content = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Snippet) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snippet) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Snippet) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *Snippet) SetPriority(v string) {
	o.Priority = &v
}

// GetDynamic returns the Dynamic field value if set, zero value otherwise.
func (o *Snippet) GetDynamic() string {
	if o == nil || o.Dynamic == nil {
		var ret string
		return ret
	}
	return *o.Dynamic
}

// GetDynamicOk returns a tuple with the Dynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snippet) GetDynamicOk() (*string, bool) {
	if o == nil || o.Dynamic == nil {
		return nil, false
	}
	return o.Dynamic, true
}

// HasDynamic returns a boolean if a field has been set.
func (o *Snippet) HasDynamic() bool {
	if o != nil && o.Dynamic != nil {
		return true
	}

	return false
}

// SetDynamic gets a reference to the given string and assigns it to the Dynamic field.
func (o *Snippet) SetDynamic(v string) {
	o.Dynamic = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Snippet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Dynamic != nil {
		toSerialize["dynamic"] = o.Dynamic
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Snippet) UnmarshalJSON(bytes []byte) (err error) {
	varSnippet := _Snippet{}

	if err = json.Unmarshal(bytes, &varSnippet); err == nil {
		*o = Snippet(varSnippet)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "content")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "dynamic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSnippet is a helper abstraction for handling nullable snippet types. 
type NullableSnippet struct {
	value *Snippet
	isSet bool
}

// Get returns the value.
func (v NullableSnippet) Get() *Snippet {
	return v.value
}

// Set modifies the value.
func (v *NullableSnippet) Set(val *Snippet) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSnippet) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSnippet) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSnippet returns a pointer to a new instance of NullableSnippet.
func NewNullableSnippet(val *Snippet) *NullableSnippet {
	return &NullableSnippet{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSnippet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSnippet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
