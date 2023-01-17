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
	"time"
)

// ServiceVersionDetail struct for ServiceVersionDetail
type ServiceVersionDetail struct {
	// Whether this is the active version or not.
	Active *bool `json:"active,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// Unused at this time.
	Deployed *bool `json:"deployed,omitempty"`
	// Whether this version is locked or not. Objects can not be added or edited on locked versions.
	Locked *bool `json:"locked,omitempty"`
	// The number of this version.
	Number *int32 `json:"number,omitempty"`
	// Unused at this time.
	Staging *bool `json:"staging,omitempty"`
	// Unused at this time.
	Testing *bool `json:"testing,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
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
	Settings *VersionDetailSettings `json:"settings,omitempty"`
	// List of VCL snippets for this service.
	Snippets []SchemasSnippetResponse `json:"snippets,omitempty"`
	// List of VCL files for this service.
	Vcls []SchemasVclResponse `json:"vcls,omitempty"`
	// A list of Wordpress rules with this service.
	Wordpress []*map[string]any `json:"wordpress,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceVersionDetail ServiceVersionDetail

// NewServiceVersionDetail instantiates a new ServiceVersionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVersionDetail() *ServiceVersionDetail {
	this := ServiceVersionDetail{}
	var active bool = false
	this.Active = &active
	var locked bool = false
	this.Locked = &locked
	var staging bool = false
	this.Staging = &staging
	var testing bool = false
	this.Testing = &testing
	return &this
}

// NewServiceVersionDetailWithDefaults instantiates a new ServiceVersionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVersionDetailWithDefaults() *ServiceVersionDetail {
	this := ServiceVersionDetail{}
	var active bool = false
	this.Active = &active
	var locked bool = false
	this.Locked = &locked
	var staging bool = false
	this.Staging = &staging
	var testing bool = false
	this.Testing = &testing
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ServiceVersionDetail) SetActive(v bool) {
	o.Active = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceVersionDetail) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceVersionDetail) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ServiceVersionDetail) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ServiceVersionDetail) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ServiceVersionDetail) UnsetComment() {
	o.Comment.Unset()
}

// GetDeployed returns the Deployed field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetDeployed() bool {
	if o == nil || o.Deployed == nil {
		var ret bool
		return ret
	}
	return *o.Deployed
}

// GetDeployedOk returns a tuple with the Deployed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetDeployedOk() (*bool, bool) {
	if o == nil || o.Deployed == nil {
		return nil, false
	}
	return o.Deployed, true
}

// HasDeployed returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasDeployed() bool {
	if o != nil && o.Deployed != nil {
		return true
	}

	return false
}

// SetDeployed gets a reference to the given bool and assigns it to the Deployed field.
func (o *ServiceVersionDetail) SetDeployed(v bool) {
	o.Deployed = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *ServiceVersionDetail) SetLocked(v bool) {
	o.Locked = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *ServiceVersionDetail) SetNumber(v int32) {
	o.Number = &v
}

// GetStaging returns the Staging field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetStaging() bool {
	if o == nil || o.Staging == nil {
		var ret bool
		return ret
	}
	return *o.Staging
}

// GetStagingOk returns a tuple with the Staging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetStagingOk() (*bool, bool) {
	if o == nil || o.Staging == nil {
		return nil, false
	}
	return o.Staging, true
}

// HasStaging returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasStaging() bool {
	if o != nil && o.Staging != nil {
		return true
	}

	return false
}

// SetStaging gets a reference to the given bool and assigns it to the Staging field.
func (o *ServiceVersionDetail) SetStaging(v bool) {
	o.Staging = &v
}

// GetTesting returns the Testing field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetTesting() bool {
	if o == nil || o.Testing == nil {
		var ret bool
		return ret
	}
	return *o.Testing
}

// GetTestingOk returns a tuple with the Testing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetTestingOk() (*bool, bool) {
	if o == nil || o.Testing == nil {
		return nil, false
	}
	return o.Testing, true
}

// HasTesting returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasTesting() bool {
	if o != nil && o.Testing != nil {
		return true
	}

	return false
}

// SetTesting gets a reference to the given bool and assigns it to the Testing field.
func (o *ServiceVersionDetail) SetTesting(v bool) {
	o.Testing = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceVersionDetail) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceVersionDetail) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *ServiceVersionDetail) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ServiceVersionDetail) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ServiceVersionDetail) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceVersionDetail) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceVersionDetail) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *ServiceVersionDetail) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *ServiceVersionDetail) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *ServiceVersionDetail) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceVersionDetail) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceVersionDetail) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *ServiceVersionDetail) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *ServiceVersionDetail) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *ServiceVersionDetail) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *ServiceVersionDetail) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetBackends returns the Backends field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetBackends() []BackendResponse {
	if o == nil || o.Backends == nil {
		var ret []BackendResponse
		return ret
	}
	return o.Backends
}

// GetBackendsOk returns a tuple with the Backends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetBackendsOk() ([]BackendResponse, bool) {
	if o == nil || o.Backends == nil {
		return nil, false
	}
	return o.Backends, true
}

// HasBackends returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasBackends() bool {
	if o != nil && o.Backends != nil {
		return true
	}

	return false
}

// SetBackends gets a reference to the given []BackendResponse and assigns it to the Backends field.
func (o *ServiceVersionDetail) SetBackends(v []BackendResponse) {
	o.Backends = v
}

// GetCacheSettings returns the CacheSettings field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetCacheSettings() []CacheSettingResponse {
	if o == nil || o.CacheSettings == nil {
		var ret []CacheSettingResponse
		return ret
	}
	return o.CacheSettings
}

// GetCacheSettingsOk returns a tuple with the CacheSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetCacheSettingsOk() ([]CacheSettingResponse, bool) {
	if o == nil || o.CacheSettings == nil {
		return nil, false
	}
	return o.CacheSettings, true
}

// HasCacheSettings returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasCacheSettings() bool {
	if o != nil && o.CacheSettings != nil {
		return true
	}

	return false
}

// SetCacheSettings gets a reference to the given []CacheSettingResponse and assigns it to the CacheSettings field.
func (o *ServiceVersionDetail) SetCacheSettings(v []CacheSettingResponse) {
	o.CacheSettings = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetConditions() []ConditionResponse {
	if o == nil || o.Conditions == nil {
		var ret []ConditionResponse
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetConditionsOk() ([]ConditionResponse, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ConditionResponse and assigns it to the Conditions field.
func (o *ServiceVersionDetail) SetConditions(v []ConditionResponse) {
	o.Conditions = v
}

// GetDirectors returns the Directors field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetDirectors() []Director {
	if o == nil || o.Directors == nil {
		var ret []Director
		return ret
	}
	return o.Directors
}

// GetDirectorsOk returns a tuple with the Directors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetDirectorsOk() ([]Director, bool) {
	if o == nil || o.Directors == nil {
		return nil, false
	}
	return o.Directors, true
}

// HasDirectors returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasDirectors() bool {
	if o != nil && o.Directors != nil {
		return true
	}

	return false
}

// SetDirectors gets a reference to the given []Director and assigns it to the Directors field.
func (o *ServiceVersionDetail) SetDirectors(v []Director) {
	o.Directors = v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetDomains() []DomainResponse {
	if o == nil || o.Domains == nil {
		var ret []DomainResponse
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetDomainsOk() ([]DomainResponse, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []DomainResponse and assigns it to the Domains field.
func (o *ServiceVersionDetail) SetDomains(v []DomainResponse) {
	o.Domains = v
}

// GetGzips returns the Gzips field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetGzips() []GzipResponse {
	if o == nil || o.Gzips == nil {
		var ret []GzipResponse
		return ret
	}
	return o.Gzips
}

// GetGzipsOk returns a tuple with the Gzips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetGzipsOk() ([]GzipResponse, bool) {
	if o == nil || o.Gzips == nil {
		return nil, false
	}
	return o.Gzips, true
}

// HasGzips returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasGzips() bool {
	if o != nil && o.Gzips != nil {
		return true
	}

	return false
}

// SetGzips gets a reference to the given []GzipResponse and assigns it to the Gzips field.
func (o *ServiceVersionDetail) SetGzips(v []GzipResponse) {
	o.Gzips = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetHeaders() []HeaderResponse {
	if o == nil || o.Headers == nil {
		var ret []HeaderResponse
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetHeadersOk() ([]HeaderResponse, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []HeaderResponse and assigns it to the Headers field.
func (o *ServiceVersionDetail) SetHeaders(v []HeaderResponse) {
	o.Headers = v
}

// GetHealthchecks returns the Healthchecks field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetHealthchecks() []HealthcheckResponse {
	if o == nil || o.Healthchecks == nil {
		var ret []HealthcheckResponse
		return ret
	}
	return o.Healthchecks
}

// GetHealthchecksOk returns a tuple with the Healthchecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetHealthchecksOk() ([]HealthcheckResponse, bool) {
	if o == nil || o.Healthchecks == nil {
		return nil, false
	}
	return o.Healthchecks, true
}

// HasHealthchecks returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasHealthchecks() bool {
	if o != nil && o.Healthchecks != nil {
		return true
	}

	return false
}

// SetHealthchecks gets a reference to the given []HealthcheckResponse and assigns it to the Healthchecks field.
func (o *ServiceVersionDetail) SetHealthchecks(v []HealthcheckResponse) {
	o.Healthchecks = v
}

// GetRequestSettings returns the RequestSettings field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetRequestSettings() []RequestSettingsResponse {
	if o == nil || o.RequestSettings == nil {
		var ret []RequestSettingsResponse
		return ret
	}
	return o.RequestSettings
}

// GetRequestSettingsOk returns a tuple with the RequestSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetRequestSettingsOk() ([]RequestSettingsResponse, bool) {
	if o == nil || o.RequestSettings == nil {
		return nil, false
	}
	return o.RequestSettings, true
}

// HasRequestSettings returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasRequestSettings() bool {
	if o != nil && o.RequestSettings != nil {
		return true
	}

	return false
}

// SetRequestSettings gets a reference to the given []RequestSettingsResponse and assigns it to the RequestSettings field.
func (o *ServiceVersionDetail) SetRequestSettings(v []RequestSettingsResponse) {
	o.RequestSettings = v
}

// GetResponseObjects returns the ResponseObjects field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetResponseObjects() []ResponseObjectResponse {
	if o == nil || o.ResponseObjects == nil {
		var ret []ResponseObjectResponse
		return ret
	}
	return o.ResponseObjects
}

// GetResponseObjectsOk returns a tuple with the ResponseObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetResponseObjectsOk() ([]ResponseObjectResponse, bool) {
	if o == nil || o.ResponseObjects == nil {
		return nil, false
	}
	return o.ResponseObjects, true
}

// HasResponseObjects returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasResponseObjects() bool {
	if o != nil && o.ResponseObjects != nil {
		return true
	}

	return false
}

// SetResponseObjects gets a reference to the given []ResponseObjectResponse and assigns it to the ResponseObjects field.
func (o *ServiceVersionDetail) SetResponseObjects(v []ResponseObjectResponse) {
	o.ResponseObjects = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetSettings() VersionDetailSettings {
	if o == nil || o.Settings == nil {
		var ret VersionDetailSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetSettingsOk() (*VersionDetailSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given VersionDetailSettings and assigns it to the Settings field.
func (o *ServiceVersionDetail) SetSettings(v VersionDetailSettings) {
	o.Settings = &v
}

// GetSnippets returns the Snippets field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetSnippets() []SchemasSnippetResponse {
	if o == nil || o.Snippets == nil {
		var ret []SchemasSnippetResponse
		return ret
	}
	return o.Snippets
}

// GetSnippetsOk returns a tuple with the Snippets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetSnippetsOk() ([]SchemasSnippetResponse, bool) {
	if o == nil || o.Snippets == nil {
		return nil, false
	}
	return o.Snippets, true
}

// HasSnippets returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasSnippets() bool {
	if o != nil && o.Snippets != nil {
		return true
	}

	return false
}

// SetSnippets gets a reference to the given []SchemasSnippetResponse and assigns it to the Snippets field.
func (o *ServiceVersionDetail) SetSnippets(v []SchemasSnippetResponse) {
	o.Snippets = v
}

// GetVcls returns the Vcls field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetVcls() []SchemasVclResponse {
	if o == nil || o.Vcls == nil {
		var ret []SchemasVclResponse
		return ret
	}
	return o.Vcls
}

// GetVclsOk returns a tuple with the Vcls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetVclsOk() ([]SchemasVclResponse, bool) {
	if o == nil || o.Vcls == nil {
		return nil, false
	}
	return o.Vcls, true
}

// HasVcls returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasVcls() bool {
	if o != nil && o.Vcls != nil {
		return true
	}

	return false
}

// SetVcls gets a reference to the given []SchemasVclResponse and assigns it to the Vcls field.
func (o *ServiceVersionDetail) SetVcls(v []SchemasVclResponse) {
	o.Vcls = v
}

// GetWordpress returns the Wordpress field value if set, zero value otherwise.
func (o *ServiceVersionDetail) GetWordpress() []*map[string]any {
	if o == nil || o.Wordpress == nil {
		var ret []*map[string]any
		return ret
	}
	return o.Wordpress
}

// GetWordpressOk returns a tuple with the Wordpress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersionDetail) GetWordpressOk() ([]*map[string]any, bool) {
	if o == nil || o.Wordpress == nil {
		return nil, false
	}
	return o.Wordpress, true
}

// HasWordpress returns a boolean if a field has been set.
func (o *ServiceVersionDetail) HasWordpress() bool {
	if o != nil && o.Wordpress != nil {
		return true
	}

	return false
}

// SetWordpress gets a reference to the given []*map[string]any and assigns it to the Wordpress field.
func (o *ServiceVersionDetail) SetWordpress(v []*map[string]any) {
	o.Wordpress = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceVersionDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Deployed != nil {
		toSerialize["deployed"] = o.Deployed
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Staging != nil {
		toSerialize["staging"] = o.Staging
	}
	if o.Testing != nil {
		toSerialize["testing"] = o.Testing
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
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
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
func (o *ServiceVersionDetail) UnmarshalJSON(bytes []byte) (err error) {
	varServiceVersionDetail := _ServiceVersionDetail{}

	if err = json.Unmarshal(bytes, &varServiceVersionDetail); err == nil {
		*o = ServiceVersionDetail(varServiceVersionDetail)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "deployed")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "number")
		delete(additionalProperties, "staging")
		delete(additionalProperties, "testing")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
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

// NullableServiceVersionDetail is a helper abstraction for handling nullable serviceversiondetail types. 
type NullableServiceVersionDetail struct {
	value *ServiceVersionDetail
	isSet bool
}

// Get returns the value.
func (v NullableServiceVersionDetail) Get() *ServiceVersionDetail {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceVersionDetail) Set(val *ServiceVersionDetail) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceVersionDetail) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceVersionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceVersionDetail returns a pointer to a new instance of NullableServiceVersionDetail.
func NewNullableServiceVersionDetail(val *ServiceVersionDetail) *NullableServiceVersionDetail {
	return &NullableServiceVersionDetail{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceVersionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceVersionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
