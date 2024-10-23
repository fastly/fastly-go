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

// VersionDetail struct for VersionDetail
type VersionDetail struct {
	// List of backends associated to this service.
	Backends []BackendResponse `json:"backends,omitempty"`
	// List of cache settings associated to this service.
	CacheSettings []CacheSettingResponse `json:"cache_settings,omitempty"`
	// List of conditions associated to this service.
	Conditions []ConditionResponse `json:"conditions,omitempty"`
	// List of directors associated to this service.
	Directors []Director `json:"directors,omitempty"`
	// List of domains associated to this service.
	Domains []DomainResponse `json:"domains,omitempty"`
	// List of gzip rules associated to this service.
	Gzips []GzipResponse `json:"gzips,omitempty"`
	// List of headers associated to this service.
	Headers []HeaderResponse `json:"headers,omitempty"`
	// List of healthchecks associated to this service.
	Healthchecks []HealthcheckResponse `json:"healthchecks,omitempty"`
	// List of request settings for this service.
	RequestSettings []RequestSettingsResponse `json:"request_settings,omitempty"`
	// List of response objects for this service.
	ResponseObjects []ResponseObjectResponse `json:"response_objects,omitempty"`
	Settings        *VersionDetailSettings   `json:"settings,omitempty"`
	// List of VCL snippets for this service.
	Snippets []SchemasSnippetResponse `json:"snippets,omitempty"`
	// List of VCL files for this service.
	Vcls []SchemasVclResponse `json:"vcls,omitempty"`
	// A list of Wordpress rules with this service.
	Wordpress            []*map[string]any `json:"wordpress,omitempty"`
	AdditionalProperties map[string]any
}

type _VersionDetail VersionDetail

// NewVersionDetail instantiates a new VersionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionDetail() *VersionDetail {
	this := VersionDetail{}
	return &this
}

// NewVersionDetailWithDefaults instantiates a new VersionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionDetailWithDefaults() *VersionDetail {
	this := VersionDetail{}
	return &this
}

// GetBackends returns the Backends field value if set, zero value otherwise.
func (o *VersionDetail) GetBackends() []BackendResponse {
	if o == nil || o.Backends == nil {
		var ret []BackendResponse
		return ret
	}
	return o.Backends
}

// GetBackendsOk returns a tuple with the Backends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetBackendsOk() ([]BackendResponse, bool) {
	if o == nil || o.Backends == nil {
		return nil, false
	}
	return o.Backends, true
}

// HasBackends returns a boolean if a field has been set.
func (o *VersionDetail) HasBackends() bool {
	if o != nil && o.Backends != nil {
		return true
	}

	return false
}

// SetBackends gets a reference to the given []BackendResponse and assigns it to the Backends field.
func (o *VersionDetail) SetBackends(v []BackendResponse) {
	o.Backends = v
}

// GetCacheSettings returns the CacheSettings field value if set, zero value otherwise.
func (o *VersionDetail) GetCacheSettings() []CacheSettingResponse {
	if o == nil || o.CacheSettings == nil {
		var ret []CacheSettingResponse
		return ret
	}
	return o.CacheSettings
}

// GetCacheSettingsOk returns a tuple with the CacheSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetCacheSettingsOk() ([]CacheSettingResponse, bool) {
	if o == nil || o.CacheSettings == nil {
		return nil, false
	}
	return o.CacheSettings, true
}

// HasCacheSettings returns a boolean if a field has been set.
func (o *VersionDetail) HasCacheSettings() bool {
	if o != nil && o.CacheSettings != nil {
		return true
	}

	return false
}

// SetCacheSettings gets a reference to the given []CacheSettingResponse and assigns it to the CacheSettings field.
func (o *VersionDetail) SetCacheSettings(v []CacheSettingResponse) {
	o.CacheSettings = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *VersionDetail) GetConditions() []ConditionResponse {
	if o == nil || o.Conditions == nil {
		var ret []ConditionResponse
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetConditionsOk() ([]ConditionResponse, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *VersionDetail) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ConditionResponse and assigns it to the Conditions field.
func (o *VersionDetail) SetConditions(v []ConditionResponse) {
	o.Conditions = v
}

// GetDirectors returns the Directors field value if set, zero value otherwise.
func (o *VersionDetail) GetDirectors() []Director {
	if o == nil || o.Directors == nil {
		var ret []Director
		return ret
	}
	return o.Directors
}

// GetDirectorsOk returns a tuple with the Directors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetDirectorsOk() ([]Director, bool) {
	if o == nil || o.Directors == nil {
		return nil, false
	}
	return o.Directors, true
}

// HasDirectors returns a boolean if a field has been set.
func (o *VersionDetail) HasDirectors() bool {
	if o != nil && o.Directors != nil {
		return true
	}

	return false
}

// SetDirectors gets a reference to the given []Director and assigns it to the Directors field.
func (o *VersionDetail) SetDirectors(v []Director) {
	o.Directors = v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *VersionDetail) GetDomains() []DomainResponse {
	if o == nil || o.Domains == nil {
		var ret []DomainResponse
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetDomainsOk() ([]DomainResponse, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *VersionDetail) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []DomainResponse and assigns it to the Domains field.
func (o *VersionDetail) SetDomains(v []DomainResponse) {
	o.Domains = v
}

// GetGzips returns the Gzips field value if set, zero value otherwise.
func (o *VersionDetail) GetGzips() []GzipResponse {
	if o == nil || o.Gzips == nil {
		var ret []GzipResponse
		return ret
	}
	return o.Gzips
}

// GetGzipsOk returns a tuple with the Gzips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetGzipsOk() ([]GzipResponse, bool) {
	if o == nil || o.Gzips == nil {
		return nil, false
	}
	return o.Gzips, true
}

// HasGzips returns a boolean if a field has been set.
func (o *VersionDetail) HasGzips() bool {
	if o != nil && o.Gzips != nil {
		return true
	}

	return false
}

// SetGzips gets a reference to the given []GzipResponse and assigns it to the Gzips field.
func (o *VersionDetail) SetGzips(v []GzipResponse) {
	o.Gzips = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *VersionDetail) GetHeaders() []HeaderResponse {
	if o == nil || o.Headers == nil {
		var ret []HeaderResponse
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetHeadersOk() ([]HeaderResponse, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *VersionDetail) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []HeaderResponse and assigns it to the Headers field.
func (o *VersionDetail) SetHeaders(v []HeaderResponse) {
	o.Headers = v
}

// GetHealthchecks returns the Healthchecks field value if set, zero value otherwise.
func (o *VersionDetail) GetHealthchecks() []HealthcheckResponse {
	if o == nil || o.Healthchecks == nil {
		var ret []HealthcheckResponse
		return ret
	}
	return o.Healthchecks
}

// GetHealthchecksOk returns a tuple with the Healthchecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetHealthchecksOk() ([]HealthcheckResponse, bool) {
	if o == nil || o.Healthchecks == nil {
		return nil, false
	}
	return o.Healthchecks, true
}

// HasHealthchecks returns a boolean if a field has been set.
func (o *VersionDetail) HasHealthchecks() bool {
	if o != nil && o.Healthchecks != nil {
		return true
	}

	return false
}

// SetHealthchecks gets a reference to the given []HealthcheckResponse and assigns it to the Healthchecks field.
func (o *VersionDetail) SetHealthchecks(v []HealthcheckResponse) {
	o.Healthchecks = v
}

// GetRequestSettings returns the RequestSettings field value if set, zero value otherwise.
func (o *VersionDetail) GetRequestSettings() []RequestSettingsResponse {
	if o == nil || o.RequestSettings == nil {
		var ret []RequestSettingsResponse
		return ret
	}
	return o.RequestSettings
}

// GetRequestSettingsOk returns a tuple with the RequestSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetRequestSettingsOk() ([]RequestSettingsResponse, bool) {
	if o == nil || o.RequestSettings == nil {
		return nil, false
	}
	return o.RequestSettings, true
}

// HasRequestSettings returns a boolean if a field has been set.
func (o *VersionDetail) HasRequestSettings() bool {
	if o != nil && o.RequestSettings != nil {
		return true
	}

	return false
}

// SetRequestSettings gets a reference to the given []RequestSettingsResponse and assigns it to the RequestSettings field.
func (o *VersionDetail) SetRequestSettings(v []RequestSettingsResponse) {
	o.RequestSettings = v
}

// GetResponseObjects returns the ResponseObjects field value if set, zero value otherwise.
func (o *VersionDetail) GetResponseObjects() []ResponseObjectResponse {
	if o == nil || o.ResponseObjects == nil {
		var ret []ResponseObjectResponse
		return ret
	}
	return o.ResponseObjects
}

// GetResponseObjectsOk returns a tuple with the ResponseObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetResponseObjectsOk() ([]ResponseObjectResponse, bool) {
	if o == nil || o.ResponseObjects == nil {
		return nil, false
	}
	return o.ResponseObjects, true
}

// HasResponseObjects returns a boolean if a field has been set.
func (o *VersionDetail) HasResponseObjects() bool {
	if o != nil && o.ResponseObjects != nil {
		return true
	}

	return false
}

// SetResponseObjects gets a reference to the given []ResponseObjectResponse and assigns it to the ResponseObjects field.
func (o *VersionDetail) SetResponseObjects(v []ResponseObjectResponse) {
	o.ResponseObjects = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *VersionDetail) GetSettings() VersionDetailSettings {
	if o == nil || o.Settings == nil {
		var ret VersionDetailSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetSettingsOk() (*VersionDetailSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *VersionDetail) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given VersionDetailSettings and assigns it to the Settings field.
func (o *VersionDetail) SetSettings(v VersionDetailSettings) {
	o.Settings = &v
}

// GetSnippets returns the Snippets field value if set, zero value otherwise.
func (o *VersionDetail) GetSnippets() []SchemasSnippetResponse {
	if o == nil || o.Snippets == nil {
		var ret []SchemasSnippetResponse
		return ret
	}
	return o.Snippets
}

// GetSnippetsOk returns a tuple with the Snippets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetSnippetsOk() ([]SchemasSnippetResponse, bool) {
	if o == nil || o.Snippets == nil {
		return nil, false
	}
	return o.Snippets, true
}

// HasSnippets returns a boolean if a field has been set.
func (o *VersionDetail) HasSnippets() bool {
	if o != nil && o.Snippets != nil {
		return true
	}

	return false
}

// SetSnippets gets a reference to the given []SchemasSnippetResponse and assigns it to the Snippets field.
func (o *VersionDetail) SetSnippets(v []SchemasSnippetResponse) {
	o.Snippets = v
}

// GetVcls returns the Vcls field value if set, zero value otherwise.
func (o *VersionDetail) GetVcls() []SchemasVclResponse {
	if o == nil || o.Vcls == nil {
		var ret []SchemasVclResponse
		return ret
	}
	return o.Vcls
}

// GetVclsOk returns a tuple with the Vcls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetVclsOk() ([]SchemasVclResponse, bool) {
	if o == nil || o.Vcls == nil {
		return nil, false
	}
	return o.Vcls, true
}

// HasVcls returns a boolean if a field has been set.
func (o *VersionDetail) HasVcls() bool {
	if o != nil && o.Vcls != nil {
		return true
	}

	return false
}

// SetVcls gets a reference to the given []SchemasVclResponse and assigns it to the Vcls field.
func (o *VersionDetail) SetVcls(v []SchemasVclResponse) {
	o.Vcls = v
}

// GetWordpress returns the Wordpress field value if set, zero value otherwise.
func (o *VersionDetail) GetWordpress() []*map[string]any {
	if o == nil || o.Wordpress == nil {
		var ret []*map[string]any
		return ret
	}
	return o.Wordpress
}

// GetWordpressOk returns a tuple with the Wordpress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetail) GetWordpressOk() ([]*map[string]any, bool) {
	if o == nil || o.Wordpress == nil {
		return nil, false
	}
	return o.Wordpress, true
}

// HasWordpress returns a boolean if a field has been set.
func (o *VersionDetail) HasWordpress() bool {
	if o != nil && o.Wordpress != nil {
		return true
	}

	return false
}

// SetWordpress gets a reference to the given []*map[string]any and assigns it to the Wordpress field.
func (o *VersionDetail) SetWordpress(v []*map[string]any) {
	o.Wordpress = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o VersionDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Backends != nil {
		toSerialize["backends"] = o.Backends
	}
	if o.CacheSettings != nil {
		toSerialize["cache_settings"] = o.CacheSettings
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Directors != nil {
		toSerialize["directors"] = o.Directors
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.Gzips != nil {
		toSerialize["gzips"] = o.Gzips
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Healthchecks != nil {
		toSerialize["healthchecks"] = o.Healthchecks
	}
	if o.RequestSettings != nil {
		toSerialize["request_settings"] = o.RequestSettings
	}
	if o.ResponseObjects != nil {
		toSerialize["response_objects"] = o.ResponseObjects
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Snippets != nil {
		toSerialize["snippets"] = o.Snippets
	}
	if o.Vcls != nil {
		toSerialize["vcls"] = o.Vcls
	}
	if o.Wordpress != nil {
		toSerialize["wordpress"] = o.Wordpress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *VersionDetail) UnmarshalJSON(bytes []byte) (err error) {
	varVersionDetail := _VersionDetail{}

	if err = json.Unmarshal(bytes, &varVersionDetail); err == nil {
		*o = VersionDetail(varVersionDetail)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "backends")
		delete(additionalProperties, "cache_settings")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "directors")
		delete(additionalProperties, "domains")
		delete(additionalProperties, "gzips")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "healthchecks")
		delete(additionalProperties, "request_settings")
		delete(additionalProperties, "response_objects")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "snippets")
		delete(additionalProperties, "vcls")
		delete(additionalProperties, "wordpress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableVersionDetail is a helper abstraction for handling nullable versiondetail types.
type NullableVersionDetail struct {
	value *VersionDetail
	isSet bool
}

// Get returns the value.
func (v NullableVersionDetail) Get() *VersionDetail {
	return v.value
}

// Set modifies the value.
func (v *NullableVersionDetail) Set(val *VersionDetail) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableVersionDetail) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableVersionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableVersionDetail returns a pointer to a new instance of NullableVersionDetail.
func NewNullableVersionDetail(val *VersionDetail) *NullableVersionDetail {
	return &NullableVersionDetail{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableVersionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableVersionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
