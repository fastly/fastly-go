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

// ServiceDetail struct for ServiceDetail
type ServiceDetail struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// The name of the service.
	Name *string `json:"name,omitempty"`
	// Alphanumeric string identifying the customer.
	CustomerID *string `json:"customer_id,omitempty"`
	// The type of this service.
	Type *string `json:"type,omitempty"`
	ID *string `json:"id,omitempty"`
	// Unused at this time.
	PublishKey *string `json:"publish_key,omitempty"`
	// Whether the service is paused. Services are paused due to a lack of traffic for an extended period of time. Services are resumed either when a draft version is activated or a locked version is cloned and reactivated.
	Paused *bool `json:"paused,omitempty"`
	// A list of [versions](https://www.fastly.com/documentation/reference/api/services/version/) associated with the service.
	Versions []SchemasVersionResponse `json:"versions,omitempty"`
	ActiveVersion NullableServiceVersionDetailOrNull `json:"active_version,omitempty"`
	Version *ServiceVersionDetail `json:"version,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceDetail ServiceDetail

// NewServiceDetail instantiates a new ServiceDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDetail() *ServiceDetail {
	this := ServiceDetail{}
	return &this
}

// NewServiceDetailWithDefaults instantiates a new ServiceDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDetailWithDefaults() *ServiceDetail {
	this := ServiceDetail{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDetail) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDetail) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceDetail) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *ServiceDetail) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ServiceDetail) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ServiceDetail) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDetail) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDetail) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ServiceDetail) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *ServiceDetail) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *ServiceDetail) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *ServiceDetail) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDetail) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDetail) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServiceDetail) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *ServiceDetail) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *ServiceDetail) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *ServiceDetail) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDetail) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDetail) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ServiceDetail) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ServiceDetail) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ServiceDetail) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ServiceDetail) UnsetComment() {
	o.Comment.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceDetail) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDetail) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceDetail) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceDetail) SetName(v string) {
	o.Name = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *ServiceDetail) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDetail) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *ServiceDetail) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *ServiceDetail) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceDetail) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDetail) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceDetail) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceDetail) SetType(v string) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ServiceDetail) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDetail) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ServiceDetail) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ServiceDetail) SetID(v string) {
	o.ID = &v
}

// GetPublishKey returns the PublishKey field value if set, zero value otherwise.
func (o *ServiceDetail) GetPublishKey() string {
	if o == nil || o.PublishKey == nil {
		var ret string
		return ret
	}
	return *o.PublishKey
}

// GetPublishKeyOk returns a tuple with the PublishKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDetail) GetPublishKeyOk() (*string, bool) {
	if o == nil || o.PublishKey == nil {
		return nil, false
	}
	return o.PublishKey, true
}

// HasPublishKey returns a boolean if a field has been set.
func (o *ServiceDetail) HasPublishKey() bool {
	if o != nil && o.PublishKey != nil {
		return true
	}

	return false
}

// SetPublishKey gets a reference to the given string and assigns it to the PublishKey field.
func (o *ServiceDetail) SetPublishKey(v string) {
	o.PublishKey = &v
}

// GetPaused returns the Paused field value if set, zero value otherwise.
func (o *ServiceDetail) GetPaused() bool {
	if o == nil || o.Paused == nil {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDetail) GetPausedOk() (*bool, bool) {
	if o == nil || o.Paused == nil {
		return nil, false
	}
	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *ServiceDetail) HasPaused() bool {
	if o != nil && o.Paused != nil {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *ServiceDetail) SetPaused(v bool) {
	o.Paused = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ServiceDetail) GetVersions() []SchemasVersionResponse {
	if o == nil || o.Versions == nil {
		var ret []SchemasVersionResponse
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDetail) GetVersionsOk() ([]SchemasVersionResponse, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ServiceDetail) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []SchemasVersionResponse and assigns it to the Versions field.
func (o *ServiceDetail) SetVersions(v []SchemasVersionResponse) {
	o.Versions = v
}

// GetActiveVersion returns the ActiveVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDetail) GetActiveVersion() ServiceVersionDetailOrNull {
	if o == nil || o.ActiveVersion.Get() == nil {
		var ret ServiceVersionDetailOrNull
		return ret
	}
	return *o.ActiveVersion.Get()
}

// GetActiveVersionOk returns a tuple with the ActiveVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDetail) GetActiveVersionOk() (*ServiceVersionDetailOrNull, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActiveVersion.Get(), o.ActiveVersion.IsSet()
}

// HasActiveVersion returns a boolean if a field has been set.
func (o *ServiceDetail) HasActiveVersion() bool {
	if o != nil && o.ActiveVersion.IsSet() {
		return true
	}

	return false
}

// SetActiveVersion gets a reference to the given NullableServiceVersionDetailOrNull and assigns it to the ActiveVersion field.
func (o *ServiceDetail) SetActiveVersion(v ServiceVersionDetailOrNull) {
	o.ActiveVersion.Set(&v)
}
// SetActiveVersionNil sets the value for ActiveVersion to be an explicit nil
func (o *ServiceDetail) SetActiveVersionNil() {
	o.ActiveVersion.Set(nil)
}

// UnsetActiveVersion ensures that no value is present for ActiveVersion, not even an explicit nil
func (o *ServiceDetail) UnsetActiveVersion() {
	o.ActiveVersion.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServiceDetail) GetVersion() ServiceVersionDetail {
	if o == nil || o.Version == nil {
		var ret ServiceVersionDetail
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDetail) GetVersionOk() (*ServiceVersionDetail, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServiceDetail) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given ServiceVersionDetail and assigns it to the Version field.
func (o *ServiceDetail) SetVersion(v ServiceVersionDetail) {
	o.Version = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.PublishKey != nil {
		toSerialize["publish_key"] = o.PublishKey
	}
	if o.Paused != nil {
		toSerialize["paused"] = o.Paused
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	if o.ActiveVersion.IsSet() {
		toSerialize["active_version"] = o.ActiveVersion.Get()
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServiceDetail) UnmarshalJSON(bytes []byte) (err error) {
	varServiceDetail := _ServiceDetail{}

	if err = json.Unmarshal(bytes, &varServiceDetail); err == nil {
		*o = ServiceDetail(varServiceDetail)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "name")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "publish_key")
		delete(additionalProperties, "paused")
		delete(additionalProperties, "versions")
		delete(additionalProperties, "active_version")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceDetail is a helper abstraction for handling nullable servicedetail types. 
type NullableServiceDetail struct {
	value *ServiceDetail
	isSet bool
}

// Get returns the value.
func (v NullableServiceDetail) Get() *ServiceDetail {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceDetail) Set(val *ServiceDetail) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceDetail) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceDetail) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceDetail returns a pointer to a new instance of NullableServiceDetail.
func NewNullableServiceDetail(val *ServiceDetail) *NullableServiceDetail {
	return &NullableServiceDetail{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
