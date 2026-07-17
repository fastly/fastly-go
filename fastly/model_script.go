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

// Script struct for Script
type Script struct {
	// Unique script identifier
	Id *string `json:"id,omitempty"`
	// Parent page ID
	PageId *string `json:"page_id,omitempty"`
	// Script source (inline or external URL)
	Source *string `json:"source,omitempty"`
	// URLs where this script was observed
	Urls        []string   `json:"urls,omitempty"`
	FirstSeenAt *time.Time `json:"first_seen_at,omitempty"`
	LastSeenAt  *time.Time `json:"last_seen_at,omitempty"`
	// Reason for authorization decision
	Justification *string `json:"justification,omitempty"`
	// Current script content hash
	CurrentHash *string `json:"current_hash,omitempty"`
	// Hash of authorized script content
	AuthorizedHash *string `json:"authorized_hash,omitempty"`
	// Script authorization status
	AuthorizationStatus  *string    `json:"authorization_status,omitempty"`
	AuthorizedAt         *time.Time `json:"authorized_at,omitempty"`
	AdditionalProperties map[string]any
}

type _Script Script

// NewScript instantiates a new Script object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScript() *Script {
	this := Script{}
	return &this
}

// NewScriptWithDefaults instantiates a new Script object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptWithDefaults() *Script {
	this := Script{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Script) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Script) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Script) SetId(v string) {
	o.Id = &v
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *Script) GetPageId() string {
	if o == nil || o.PageId == nil {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetPageIdOk() (*string, bool) {
	if o == nil || o.PageId == nil {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *Script) HasPageId() bool {
	if o != nil && o.PageId != nil {
		return true
	}

	return false
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *Script) SetPageId(v string) {
	o.PageId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Script) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Script) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *Script) SetSource(v string) {
	o.Source = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *Script) GetUrls() []string {
	if o == nil || o.Urls == nil {
		var ret []string
		return ret
	}
	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetUrlsOk() ([]string, bool) {
	if o == nil || o.Urls == nil {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *Script) HasUrls() bool {
	if o != nil && o.Urls != nil {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []string and assigns it to the Urls field.
func (o *Script) SetUrls(v []string) {
	o.Urls = v
}

// GetFirstSeenAt returns the FirstSeenAt field value if set, zero value otherwise.
func (o *Script) GetFirstSeenAt() time.Time {
	if o == nil || o.FirstSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetFirstSeenAtOk() (*time.Time, bool) {
	if o == nil || o.FirstSeenAt == nil {
		return nil, false
	}
	return o.FirstSeenAt, true
}

// HasFirstSeenAt returns a boolean if a field has been set.
func (o *Script) HasFirstSeenAt() bool {
	if o != nil && o.FirstSeenAt != nil {
		return true
	}

	return false
}

// SetFirstSeenAt gets a reference to the given time.Time and assigns it to the FirstSeenAt field.
func (o *Script) SetFirstSeenAt(v time.Time) {
	o.FirstSeenAt = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *Script) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *Script) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt != nil {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *Script) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

// GetJustification returns the Justification field value if set, zero value otherwise.
func (o *Script) GetJustification() string {
	if o == nil || o.Justification == nil {
		var ret string
		return ret
	}
	return *o.Justification
}

// GetJustificationOk returns a tuple with the Justification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetJustificationOk() (*string, bool) {
	if o == nil || o.Justification == nil {
		return nil, false
	}
	return o.Justification, true
}

// HasJustification returns a boolean if a field has been set.
func (o *Script) HasJustification() bool {
	if o != nil && o.Justification != nil {
		return true
	}

	return false
}

// SetJustification gets a reference to the given string and assigns it to the Justification field.
func (o *Script) SetJustification(v string) {
	o.Justification = &v
}

// GetCurrentHash returns the CurrentHash field value if set, zero value otherwise.
func (o *Script) GetCurrentHash() string {
	if o == nil || o.CurrentHash == nil {
		var ret string
		return ret
	}
	return *o.CurrentHash
}

// GetCurrentHashOk returns a tuple with the CurrentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetCurrentHashOk() (*string, bool) {
	if o == nil || o.CurrentHash == nil {
		return nil, false
	}
	return o.CurrentHash, true
}

// HasCurrentHash returns a boolean if a field has been set.
func (o *Script) HasCurrentHash() bool {
	if o != nil && o.CurrentHash != nil {
		return true
	}

	return false
}

// SetCurrentHash gets a reference to the given string and assigns it to the CurrentHash field.
func (o *Script) SetCurrentHash(v string) {
	o.CurrentHash = &v
}

// GetAuthorizedHash returns the AuthorizedHash field value if set, zero value otherwise.
func (o *Script) GetAuthorizedHash() string {
	if o == nil || o.AuthorizedHash == nil {
		var ret string
		return ret
	}
	return *o.AuthorizedHash
}

// GetAuthorizedHashOk returns a tuple with the AuthorizedHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetAuthorizedHashOk() (*string, bool) {
	if o == nil || o.AuthorizedHash == nil {
		return nil, false
	}
	return o.AuthorizedHash, true
}

// HasAuthorizedHash returns a boolean if a field has been set.
func (o *Script) HasAuthorizedHash() bool {
	if o != nil && o.AuthorizedHash != nil {
		return true
	}

	return false
}

// SetAuthorizedHash gets a reference to the given string and assigns it to the AuthorizedHash field.
func (o *Script) SetAuthorizedHash(v string) {
	o.AuthorizedHash = &v
}

// GetAuthorizationStatus returns the AuthorizationStatus field value if set, zero value otherwise.
func (o *Script) GetAuthorizationStatus() string {
	if o == nil || o.AuthorizationStatus == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationStatus
}

// GetAuthorizationStatusOk returns a tuple with the AuthorizationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetAuthorizationStatusOk() (*string, bool) {
	if o == nil || o.AuthorizationStatus == nil {
		return nil, false
	}
	return o.AuthorizationStatus, true
}

// HasAuthorizationStatus returns a boolean if a field has been set.
func (o *Script) HasAuthorizationStatus() bool {
	if o != nil && o.AuthorizationStatus != nil {
		return true
	}

	return false
}

// SetAuthorizationStatus gets a reference to the given string and assigns it to the AuthorizationStatus field.
func (o *Script) SetAuthorizationStatus(v string) {
	o.AuthorizationStatus = &v
}

// GetAuthorizedAt returns the AuthorizedAt field value if set, zero value otherwise.
func (o *Script) GetAuthorizedAt() time.Time {
	if o == nil || o.AuthorizedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AuthorizedAt
}

// GetAuthorizedAtOk returns a tuple with the AuthorizedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetAuthorizedAtOk() (*time.Time, bool) {
	if o == nil || o.AuthorizedAt == nil {
		return nil, false
	}
	return o.AuthorizedAt, true
}

// HasAuthorizedAt returns a boolean if a field has been set.
func (o *Script) HasAuthorizedAt() bool {
	if o != nil && o.AuthorizedAt != nil {
		return true
	}

	return false
}

// SetAuthorizedAt gets a reference to the given time.Time and assigns it to the AuthorizedAt field.
func (o *Script) SetAuthorizedAt(v time.Time) {
	o.AuthorizedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Script) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PageId != nil {
		toSerialize["page_id"] = o.PageId
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}
	if o.FirstSeenAt != nil {
		toSerialize["first_seen_at"] = o.FirstSeenAt
	}
	if o.LastSeenAt != nil {
		toSerialize["last_seen_at"] = o.LastSeenAt
	}
	if o.Justification != nil {
		toSerialize["justification"] = o.Justification
	}
	if o.CurrentHash != nil {
		toSerialize["current_hash"] = o.CurrentHash
	}
	if o.AuthorizedHash != nil {
		toSerialize["authorized_hash"] = o.AuthorizedHash
	}
	if o.AuthorizationStatus != nil {
		toSerialize["authorization_status"] = o.AuthorizationStatus
	}
	if o.AuthorizedAt != nil {
		toSerialize["authorized_at"] = o.AuthorizedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Script) UnmarshalJSON(bytes []byte) (err error) {
	varScript := _Script{}

	if err = json.Unmarshal(bytes, &varScript); err == nil {
		*o = Script(varScript)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "page_id")
		delete(additionalProperties, "source")
		delete(additionalProperties, "urls")
		delete(additionalProperties, "first_seen_at")
		delete(additionalProperties, "last_seen_at")
		delete(additionalProperties, "justification")
		delete(additionalProperties, "current_hash")
		delete(additionalProperties, "authorized_hash")
		delete(additionalProperties, "authorization_status")
		delete(additionalProperties, "authorized_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableScript is a helper abstraction for handling nullable script types.
type NullableScript struct {
	value *Script
	isSet bool
}

// Get returns the value.
func (v NullableScript) Get() *Script {
	return v.value
}

// Set modifies the value.
func (v *NullableScript) Set(val *Script) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableScript) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableScript) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableScript returns a pointer to a new instance of NullableScript.
func NewNullableScript(val *Script) *NullableScript {
	return &NullableScript{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
