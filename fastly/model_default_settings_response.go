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

// DefaultSettingsResponse struct for DefaultSettingsResponse
type DefaultSettingsResponse struct {
	// The type of filter to use while resizing an image.
	ResizeFilter string `json:"resize_filter"`
	// Controls whether or not to default to WebP output when the client supports it. This is equivalent to adding \"auto=webp\" to all image optimizer requests.
	Webp bool `json:"webp"`
	// The default quality to use with WebP output. This can be overridden with the second option in the \"quality\" URL parameter on specific image optimizer requests.
	WebpQuality int32 `json:"webp_quality"`
	// The default type of JPEG output to use. This can be overridden with \"format=bjpeg\" and \"format=pjpeg\" on specific image optimizer requests.
	JpegType string `json:"jpeg_type"`
	// The default quality to use with JPEG output. This can be overridden with the \"quality\" parameter on specific image optimizer requests.
	JpegQuality int32 `json:"jpeg_quality"`
	// Whether or not we should allow output images to render at sizes larger than input.
	Upscale bool `json:"upscale"`
	// Enables GIF to MP4 transformations on this service.
	AllowVideo           bool `json:"allow_video"`
	AdditionalProperties map[string]any
}

type _DefaultSettingsResponse DefaultSettingsResponse

// NewDefaultSettingsResponse instantiates a new DefaultSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultSettingsResponse(resizeFilter string, webp bool, webpQuality int32, jpegType string, jpegQuality int32, upscale bool, allowVideo bool) *DefaultSettingsResponse {
	this := DefaultSettingsResponse{}
	this.ResizeFilter = resizeFilter
	this.Webp = webp
	this.WebpQuality = webpQuality
	this.JpegType = jpegType
	this.JpegQuality = jpegQuality
	this.Upscale = upscale
	this.AllowVideo = allowVideo
	return &this
}

// NewDefaultSettingsResponseWithDefaults instantiates a new DefaultSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultSettingsResponseWithDefaults() *DefaultSettingsResponse {
	this := DefaultSettingsResponse{}
	var resizeFilter string = "lanczos3"
	this.ResizeFilter = resizeFilter
	var webp bool = false
	this.Webp = webp
	var webpQuality int32 = 85
	this.WebpQuality = webpQuality
	var jpegType string = "auto"
	this.JpegType = jpegType
	var jpegQuality int32 = 85
	this.JpegQuality = jpegQuality
	var upscale bool = false
	this.Upscale = upscale
	var allowVideo bool = false
	this.AllowVideo = allowVideo
	return &this
}

// GetResizeFilter returns the ResizeFilter field value
func (o *DefaultSettingsResponse) GetResizeFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResizeFilter
}

// GetResizeFilterOk returns a tuple with the ResizeFilter field value
// and a boolean to check if the value has been set.
func (o *DefaultSettingsResponse) GetResizeFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResizeFilter, true
}

// SetResizeFilter sets field value
func (o *DefaultSettingsResponse) SetResizeFilter(v string) {
	o.ResizeFilter = v
}

// GetWebp returns the Webp field value
func (o *DefaultSettingsResponse) GetWebp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Webp
}

// GetWebpOk returns a tuple with the Webp field value
// and a boolean to check if the value has been set.
func (o *DefaultSettingsResponse) GetWebpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Webp, true
}

// SetWebp sets field value
func (o *DefaultSettingsResponse) SetWebp(v bool) {
	o.Webp = v
}

// GetWebpQuality returns the WebpQuality field value
func (o *DefaultSettingsResponse) GetWebpQuality() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WebpQuality
}

// GetWebpQualityOk returns a tuple with the WebpQuality field value
// and a boolean to check if the value has been set.
func (o *DefaultSettingsResponse) GetWebpQualityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebpQuality, true
}

// SetWebpQuality sets field value
func (o *DefaultSettingsResponse) SetWebpQuality(v int32) {
	o.WebpQuality = v
}

// GetJpegType returns the JpegType field value
func (o *DefaultSettingsResponse) GetJpegType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JpegType
}

// GetJpegTypeOk returns a tuple with the JpegType field value
// and a boolean to check if the value has been set.
func (o *DefaultSettingsResponse) GetJpegTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JpegType, true
}

// SetJpegType sets field value
func (o *DefaultSettingsResponse) SetJpegType(v string) {
	o.JpegType = v
}

// GetJpegQuality returns the JpegQuality field value
func (o *DefaultSettingsResponse) GetJpegQuality() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.JpegQuality
}

// GetJpegQualityOk returns a tuple with the JpegQuality field value
// and a boolean to check if the value has been set.
func (o *DefaultSettingsResponse) GetJpegQualityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JpegQuality, true
}

// SetJpegQuality sets field value
func (o *DefaultSettingsResponse) SetJpegQuality(v int32) {
	o.JpegQuality = v
}

// GetUpscale returns the Upscale field value
func (o *DefaultSettingsResponse) GetUpscale() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Upscale
}

// GetUpscaleOk returns a tuple with the Upscale field value
// and a boolean to check if the value has been set.
func (o *DefaultSettingsResponse) GetUpscaleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Upscale, true
}

// SetUpscale sets field value
func (o *DefaultSettingsResponse) SetUpscale(v bool) {
	o.Upscale = v
}

// GetAllowVideo returns the AllowVideo field value
func (o *DefaultSettingsResponse) GetAllowVideo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowVideo
}

// GetAllowVideoOk returns a tuple with the AllowVideo field value
// and a boolean to check if the value has been set.
func (o *DefaultSettingsResponse) GetAllowVideoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowVideo, true
}

// SetAllowVideo sets field value
func (o *DefaultSettingsResponse) SetAllowVideo(v bool) {
	o.AllowVideo = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DefaultSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["resize_filter"] = o.ResizeFilter
	}
	if true {
		toSerialize["webp"] = o.Webp
	}
	if true {
		toSerialize["webp_quality"] = o.WebpQuality
	}
	if true {
		toSerialize["jpeg_type"] = o.JpegType
	}
	if true {
		toSerialize["jpeg_quality"] = o.JpegQuality
	}
	if true {
		toSerialize["upscale"] = o.Upscale
	}
	if true {
		toSerialize["allow_video"] = o.AllowVideo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DefaultSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDefaultSettingsResponse := _DefaultSettingsResponse{}

	if err = json.Unmarshal(bytes, &varDefaultSettingsResponse); err == nil {
		*o = DefaultSettingsResponse(varDefaultSettingsResponse)
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

// NullableDefaultSettingsResponse is a helper abstraction for handling nullable defaultsettingsresponse types.
type NullableDefaultSettingsResponse struct {
	value *DefaultSettingsResponse
	isSet bool
}

// Get returns the value.
func (v NullableDefaultSettingsResponse) Get() *DefaultSettingsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableDefaultSettingsResponse) Set(val *DefaultSettingsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDefaultSettingsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDefaultSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDefaultSettingsResponse returns a pointer to a new instance of NullableDefaultSettingsResponse.
func NewNullableDefaultSettingsResponse(val *DefaultSettingsResponse) *NullableDefaultSettingsResponse {
	return &NullableDefaultSettingsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDefaultSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDefaultSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
