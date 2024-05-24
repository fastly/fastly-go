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

// PackageMetadata [Package metadata](#metadata-model) that has been extracted from the uploaded package. 
type PackageMetadata struct {
	// Name of the Compute package.
	Name *string `json:"name,omitempty"`
	// Description of the Compute package.
	Description *string `json:"description,omitempty"`
	// A list of package authors' email addresses.
	Authors []string `json:"authors,omitempty"`
	// The language of the Compute package.
	Language *string `json:"language,omitempty"`
	// Size of the Compute package in bytes.
	Size *int32 `json:"size,omitempty"`
	// Hash of the Compute package.
	Hashsum *string `json:"hashsum,omitempty"`
	// Hash of the files within the Compute package.
	FilesHash *string `json:"files_hash,omitempty"`
	AdditionalProperties map[string]any
}

type _PackageMetadata PackageMetadata

// NewPackageMetadata instantiates a new PackageMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageMetadata() *PackageMetadata {
	this := PackageMetadata{}
	return &this
}

// NewPackageMetadataWithDefaults instantiates a new PackageMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageMetadataWithDefaults() *PackageMetadata {
	this := PackageMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PackageMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PackageMetadata) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PackageMetadata) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PackageMetadata) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PackageMetadata) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PackageMetadata) SetDescription(v string) {
	o.Description = &v
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *PackageMetadata) GetAuthors() []string {
	if o == nil || o.Authors == nil {
		var ret []string
		return ret
	}
	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetAuthorsOk() ([]string, bool) {
	if o == nil || o.Authors == nil {
		return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *PackageMetadata) HasAuthors() bool {
	if o != nil && o.Authors != nil {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given []string and assigns it to the Authors field.
func (o *PackageMetadata) SetAuthors(v []string) {
	o.Authors = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PackageMetadata) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PackageMetadata) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *PackageMetadata) SetLanguage(v string) {
	o.Language = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PackageMetadata) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PackageMetadata) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PackageMetadata) SetSize(v int32) {
	o.Size = &v
}

// GetHashsum returns the Hashsum field value if set, zero value otherwise.
func (o *PackageMetadata) GetHashsum() string {
	if o == nil || o.Hashsum == nil {
		var ret string
		return ret
	}
	return *o.Hashsum
}

// GetHashsumOk returns a tuple with the Hashsum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetHashsumOk() (*string, bool) {
	if o == nil || o.Hashsum == nil {
		return nil, false
	}
	return o.Hashsum, true
}

// HasHashsum returns a boolean if a field has been set.
func (o *PackageMetadata) HasHashsum() bool {
	if o != nil && o.Hashsum != nil {
		return true
	}

	return false
}

// SetHashsum gets a reference to the given string and assigns it to the Hashsum field.
func (o *PackageMetadata) SetHashsum(v string) {
	o.Hashsum = &v
}

// GetFilesHash returns the FilesHash field value if set, zero value otherwise.
func (o *PackageMetadata) GetFilesHash() string {
	if o == nil || o.FilesHash == nil {
		var ret string
		return ret
	}
	return *o.FilesHash
}

// GetFilesHashOk returns a tuple with the FilesHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageMetadata) GetFilesHashOk() (*string, bool) {
	if o == nil || o.FilesHash == nil {
		return nil, false
	}
	return o.FilesHash, true
}

// HasFilesHash returns a boolean if a field has been set.
func (o *PackageMetadata) HasFilesHash() bool {
	if o != nil && o.FilesHash != nil {
		return true
	}

	return false
}

// SetFilesHash gets a reference to the given string and assigns it to the FilesHash field.
func (o *PackageMetadata) SetFilesHash(v string) {
	o.FilesHash = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PackageMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Authors != nil {
		toSerialize["authors"] = o.Authors
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Hashsum != nil {
		toSerialize["hashsum"] = o.Hashsum
	}
	if o.FilesHash != nil {
		toSerialize["files_hash"] = o.FilesHash
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *PackageMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varPackageMetadata := _PackageMetadata{}

	if err = json.Unmarshal(bytes, &varPackageMetadata); err == nil {
		*o = PackageMetadata(varPackageMetadata)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "authors")
		delete(additionalProperties, "language")
		delete(additionalProperties, "size")
		delete(additionalProperties, "hashsum")
		delete(additionalProperties, "files_hash")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePackageMetadata is a helper abstraction for handling nullable packagemetadata types. 
type NullablePackageMetadata struct {
	value *PackageMetadata
	isSet bool
}

// Get returns the value.
func (v NullablePackageMetadata) Get() *PackageMetadata {
	return v.value
}

// Set modifies the value.
func (v *NullablePackageMetadata) Set(val *PackageMetadata) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePackageMetadata) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePackageMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePackageMetadata returns a pointer to a new instance of NullablePackageMetadata.
func NewNullablePackageMetadata(val *PackageMetadata) *NullablePackageMetadata {
	return &NullablePackageMetadata{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePackageMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePackageMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
