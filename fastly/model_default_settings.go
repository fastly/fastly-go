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

// DefaultSettings struct for DefaultSettings
type DefaultSettings struct {
	// The type of filter to use while resizing an image.
	ResizeFilter *string `json:"resize_filter,omitempty"`
	// Controls whether or not to default to WebP output when the client supports it. This is equivalent to adding \"auto=webp\" to all image optimizer requests.
	Webp *bool `json:"webp,omitempty"`
	// The default quality to use with WebP output. This can be overridden with the second option in the \"quality\" URL parameter on specific image optimizer requests.
	WebpQuality *int32 `json:"webp_quality,omitempty"`
	// The default type of JPEG output to use. This can be overridden with \"format=bjpeg\" and \"format=pjpeg\" on specific image optimizer requests.
	JpegType *string `json:"jpeg_type,omitempty"`
	// The default quality to use with JPEG output. This can be overridden with the \"quality\" parameter on specific image optimizer requests.
	JpegQuality *int32 `json:"jpeg_quality,omitempty"`
	// Whether or not we should allow output images to render at sizes larger than input.
	Upscale *bool `json:"upscale,omitempty"`
	// Enables GIF to MP4 transformations on this service.
	AllowVideo           *bool `json:"allow_video,omitempty"`
	AdditionalProperties map[string]any
}

type _DefaultSettings DefaultSettings

// NewDefaultSettings instantiates a new DefaultSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultSettings() *DefaultSettings {
	this := DefaultSettings{}
	var resizeFilter string = "lanczos3"
	this.ResizeFilter = &resizeFilter
	var webp bool = false
	this.Webp = &webp
	var webpQuality int32 = 85
	this.WebpQuality = &webpQuality
	var jpegType string = "auto"
	this.JpegType = &jpegType
	var jpegQuality int32 = 85
	this.JpegQuality = &jpegQuality
	var upscale bool = false
	this.Upscale = &upscale
	var allowVideo bool = false
	this.AllowVideo = &allowVideo
	return &this
}

// NewDefaultSettingsWithDefaults instantiates a new DefaultSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultSettingsWithDefaults() *DefaultSettings {
	this := DefaultSettings{}
	var resizeFilter string = "lanczos3"
	this.ResizeFilter = &resizeFilter
	var webp bool = false
	this.Webp = &webp
	var webpQuality int32 = 85
	this.WebpQuality = &webpQuality
	var jpegType string = "auto"
	this.JpegType = &jpegType
	var jpegQuality int32 = 85
	this.JpegQuality = &jpegQuality
	var upscale bool = false
	this.Upscale = &upscale
	var allowVideo bool = false
	this.AllowVideo = &allowVideo
	return &this
}

// GetResizeFilter returns the ResizeFilter field value if set, zero value otherwise.
func (o *DefaultSettings) GetResizeFilter() string {
	if o == nil || o.ResizeFilter == nil {
		var ret string
		return ret
	}
	return *o.ResizeFilter
}

// GetResizeFilterOk returns a tuple with the ResizeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultSettings) GetResizeFilterOk() (*string, bool) {
	if o == nil || o.ResizeFilter == nil {
		return nil, false
	}
	return o.ResizeFilter, true
}

// HasResizeFilter returns a boolean if a field has been set.
func (o *DefaultSettings) HasResizeFilter() bool {
	if o != nil && o.ResizeFilter != nil {
		return true
	}

	return false
}

// SetResizeFilter gets a reference to the given string and assigns it to the ResizeFilter field.
func (o *DefaultSettings) SetResizeFilter(v string) {
	o.ResizeFilter = &v
}

// GetWebp returns the Webp field value if set, zero value otherwise.
func (o *DefaultSettings) GetWebp() bool {
	if o == nil || o.Webp == nil {
		var ret bool
		return ret
	}
	return *o.Webp
}

// GetWebpOk returns a tuple with the Webp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultSettings) GetWebpOk() (*bool, bool) {
	if o == nil || o.Webp == nil {
		return nil, false
	}
	return o.Webp, true
}

// HasWebp returns a boolean if a field has been set.
func (o *DefaultSettings) HasWebp() bool {
	if o != nil && o.Webp != nil {
		return true
	}

	return false
}

// SetWebp gets a reference to the given bool and assigns it to the Webp field.
func (o *DefaultSettings) SetWebp(v bool) {
	o.Webp = &v
}

// GetWebpQuality returns the WebpQuality field value if set, zero value otherwise.
func (o *DefaultSettings) GetWebpQuality() int32 {
	if o == nil || o.WebpQuality == nil {
		var ret int32
		return ret
	}
	return *o.WebpQuality
}

// GetWebpQualityOk returns a tuple with the WebpQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultSettings) GetWebpQualityOk() (*int32, bool) {
	if o == nil || o.WebpQuality == nil {
		return nil, false
	}
	return o.WebpQuality, true
}

// HasWebpQuality returns a boolean if a field has been set.
func (o *DefaultSettings) HasWebpQuality() bool {
	if o != nil && o.WebpQuality != nil {
		return true
	}

	return false
}

// SetWebpQuality gets a reference to the given int32 and assigns it to the WebpQuality field.
func (o *DefaultSettings) SetWebpQuality(v int32) {
	o.WebpQuality = &v
}

// GetJpegType returns the JpegType field value if set, zero value otherwise.
func (o *DefaultSettings) GetJpegType() string {
	if o == nil || o.JpegType == nil {
		var ret string
		return ret
	}
	return *o.JpegType
}

// GetJpegTypeOk returns a tuple with the JpegType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultSettings) GetJpegTypeOk() (*string, bool) {
	if o == nil || o.JpegType == nil {
		return nil, false
	}
	return o.JpegType, true
}

// HasJpegType returns a boolean if a field has been set.
func (o *DefaultSettings) HasJpegType() bool {
	if o != nil && o.JpegType != nil {
		return true
	}

	return false
}

// SetJpegType gets a reference to the given string and assigns it to the JpegType field.
func (o *DefaultSettings) SetJpegType(v string) {
	o.JpegType = &v
}

// GetJpegQuality returns the JpegQuality field value if set, zero value otherwise.
func (o *DefaultSettings) GetJpegQuality() int32 {
	if o == nil || o.JpegQuality == nil {
		var ret int32
		return ret
	}
	return *o.JpegQuality
}

// GetJpegQualityOk returns a tuple with the JpegQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultSettings) GetJpegQualityOk() (*int32, bool) {
	if o == nil || o.JpegQuality == nil {
		return nil, false
	}
	return o.JpegQuality, true
}

// HasJpegQuality returns a boolean if a field has been set.
func (o *DefaultSettings) HasJpegQuality() bool {
	if o != nil && o.JpegQuality != nil {
		return true
	}

	return false
}

// SetJpegQuality gets a reference to the given int32 and assigns it to the JpegQuality field.
func (o *DefaultSettings) SetJpegQuality(v int32) {
	o.JpegQuality = &v
}

// GetUpscale returns the Upscale field value if set, zero value otherwise.
func (o *DefaultSettings) GetUpscale() bool {
	if o == nil || o.Upscale == nil {
		var ret bool
		return ret
	}
	return *o.Upscale
}

// GetUpscaleOk returns a tuple with the Upscale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultSettings) GetUpscaleOk() (*bool, bool) {
	if o == nil || o.Upscale == nil {
		return nil, false
	}
	return o.Upscale, true
}

// HasUpscale returns a boolean if a field has been set.
func (o *DefaultSettings) HasUpscale() bool {
	if o != nil && o.Upscale != nil {
		return true
	}

	return false
}

// SetUpscale gets a reference to the given bool and assigns it to the Upscale field.
func (o *DefaultSettings) SetUpscale(v bool) {
	o.Upscale = &v
}

// GetAllowVideo returns the AllowVideo field value if set, zero value otherwise.
func (o *DefaultSettings) GetAllowVideo() bool {
	if o == nil || o.AllowVideo == nil {
		var ret bool
		return ret
	}
	return *o.AllowVideo
}

// GetAllowVideoOk returns a tuple with the AllowVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultSettings) GetAllowVideoOk() (*bool, bool) {
	if o == nil || o.AllowVideo == nil {
		return nil, false
	}
	return o.AllowVideo, true
}

// HasAllowVideo returns a boolean if a field has been set.
func (o *DefaultSettings) HasAllowVideo() bool {
	if o != nil && o.AllowVideo != nil {
		return true
	}

	return false
}

// SetAllowVideo gets a reference to the given bool and assigns it to the AllowVideo field.
func (o *DefaultSettings) SetAllowVideo(v bool) {
	o.AllowVideo = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DefaultSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ResizeFilter != nil {
		toSerialize["resize_filter"] = o.ResizeFilter
	}
	if o.Webp != nil {
		toSerialize["webp"] = o.Webp
	}
	if o.WebpQuality != nil {
		toSerialize["webp_quality"] = o.WebpQuality
	}
	if o.JpegType != nil {
		toSerialize["jpeg_type"] = o.JpegType
	}
	if o.JpegQuality != nil {
		toSerialize["jpeg_quality"] = o.JpegQuality
	}
	if o.Upscale != nil {
		toSerialize["upscale"] = o.Upscale
	}
	if o.AllowVideo != nil {
		toSerialize["allow_video"] = o.AllowVideo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DefaultSettings) UnmarshalJSON(bytes []byte) (err error) {
	varDefaultSettings := _DefaultSettings{}

	if err = json.Unmarshal(bytes, &varDefaultSettings); err == nil {
		*o = DefaultSettings(varDefaultSettings)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "resize_filter")
		delete(additionalProperties, "webp")
		delete(additionalProperties, "webp_quality")
		delete(additionalProperties, "jpeg_type")
		delete(additionalProperties, "jpeg_quality")
		delete(additionalProperties, "upscale")
		delete(additionalProperties, "allow_video")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDefaultSettings is a helper abstraction for handling nullable defaultsettings types.
type NullableDefaultSettings struct {
	value *DefaultSettings
	isSet bool
}

// Get returns the value.
func (v NullableDefaultSettings) Get() *DefaultSettings {
	return v.value
}

// Set modifies the value.
func (v *NullableDefaultSettings) Set(val *DefaultSettings) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDefaultSettings) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDefaultSettings) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDefaultSettings returns a pointer to a new instance of NullableDefaultSettings.
func NewNullableDefaultSettings(val *DefaultSettings) *NullableDefaultSettings {
	return &NullableDefaultSettings{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDefaultSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDefaultSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
