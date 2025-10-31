# DefaultSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResizeFilter** | Pointer to **string** | The type of filter to use while resizing an image. | [optional] [default to "lanczos3"]
**Webp** | Pointer to **bool** | Controls whether or not to default to WebP output when the client supports it. This is equivalent to adding \&quot;auto&#x3D;webp\&quot; to all image optimizer requests.  | [optional] [default to false]
**WebpQuality** | Pointer to **int32** | The default quality to use with WebP output. This can be overridden with the second option in the \&quot;quality\&quot; URL parameter on specific image optimizer requests.  | [optional] [default to 85]
**JpegType** | Pointer to **string** | The default type of JPEG output to use. This can be overridden with \&quot;format&#x3D;bjpeg\&quot; and \&quot;format&#x3D;pjpeg\&quot; on specific image optimizer requests.  | [optional] [default to "auto"]
**JpegQuality** | Pointer to **int32** | The default quality to use with JPEG output. This can be overridden with the \&quot;quality\&quot; parameter on specific image optimizer requests.  | [optional] [default to 85]
**Upscale** | Pointer to **bool** | Whether or not we should allow output images to render at sizes larger than input.  | [optional] [default to false]
**AllowVideo** | Pointer to **bool** | Enables GIF to MP4 transformations on this service. | [optional] [default to false]

## Methods

### NewDefaultSettings

`func NewDefaultSettings() *DefaultSettings`

NewDefaultSettings instantiates a new DefaultSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultSettingsWithDefaults

`func NewDefaultSettingsWithDefaults() *DefaultSettings`

NewDefaultSettingsWithDefaults instantiates a new DefaultSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResizeFilter

`func (o *DefaultSettings) GetResizeFilter() string`

GetResizeFilter returns the ResizeFilter field if non-nil, zero value otherwise.

### GetResizeFilterOk

`func (o *DefaultSettings) GetResizeFilterOk() (*string, bool)`

GetResizeFilterOk returns a tuple with the ResizeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeFilter

`func (o *DefaultSettings) SetResizeFilter(v string)`

SetResizeFilter sets ResizeFilter field to given value.

### HasResizeFilter

`func (o *DefaultSettings) HasResizeFilter() bool`

HasResizeFilter returns a boolean if a field has been set.

### GetWebp

`func (o *DefaultSettings) GetWebp() bool`

GetWebp returns the Webp field if non-nil, zero value otherwise.

### GetWebpOk

`func (o *DefaultSettings) GetWebpOk() (*bool, bool)`

GetWebpOk returns a tuple with the Webp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebp

`func (o *DefaultSettings) SetWebp(v bool)`

SetWebp sets Webp field to given value.

### HasWebp

`func (o *DefaultSettings) HasWebp() bool`

HasWebp returns a boolean if a field has been set.

### GetWebpQuality

`func (o *DefaultSettings) GetWebpQuality() int32`

GetWebpQuality returns the WebpQuality field if non-nil, zero value otherwise.

### GetWebpQualityOk

`func (o *DefaultSettings) GetWebpQualityOk() (*int32, bool)`

GetWebpQualityOk returns a tuple with the WebpQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebpQuality

`func (o *DefaultSettings) SetWebpQuality(v int32)`

SetWebpQuality sets WebpQuality field to given value.

### HasWebpQuality

`func (o *DefaultSettings) HasWebpQuality() bool`

HasWebpQuality returns a boolean if a field has been set.

### GetJpegType

`func (o *DefaultSettings) GetJpegType() string`

GetJpegType returns the JpegType field if non-nil, zero value otherwise.

### GetJpegTypeOk

`func (o *DefaultSettings) GetJpegTypeOk() (*string, bool)`

GetJpegTypeOk returns a tuple with the JpegType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpegType

`func (o *DefaultSettings) SetJpegType(v string)`

SetJpegType sets JpegType field to given value.

### HasJpegType

`func (o *DefaultSettings) HasJpegType() bool`

HasJpegType returns a boolean if a field has been set.

### GetJpegQuality

`func (o *DefaultSettings) GetJpegQuality() int32`

GetJpegQuality returns the JpegQuality field if non-nil, zero value otherwise.

### GetJpegQualityOk

`func (o *DefaultSettings) GetJpegQualityOk() (*int32, bool)`

GetJpegQualityOk returns a tuple with the JpegQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpegQuality

`func (o *DefaultSettings) SetJpegQuality(v int32)`

SetJpegQuality sets JpegQuality field to given value.

### HasJpegQuality

`func (o *DefaultSettings) HasJpegQuality() bool`

HasJpegQuality returns a boolean if a field has been set.

### GetUpscale

`func (o *DefaultSettings) GetUpscale() bool`

GetUpscale returns the Upscale field if non-nil, zero value otherwise.

### GetUpscaleOk

`func (o *DefaultSettings) GetUpscaleOk() (*bool, bool)`

GetUpscaleOk returns a tuple with the Upscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpscale

`func (o *DefaultSettings) SetUpscale(v bool)`

SetUpscale sets Upscale field to given value.

### HasUpscale

`func (o *DefaultSettings) HasUpscale() bool`

HasUpscale returns a boolean if a field has been set.

### GetAllowVideo

`func (o *DefaultSettings) GetAllowVideo() bool`

GetAllowVideo returns the AllowVideo field if non-nil, zero value otherwise.

### GetAllowVideoOk

`func (o *DefaultSettings) GetAllowVideoOk() (*bool, bool)`

GetAllowVideoOk returns a tuple with the AllowVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVideo

`func (o *DefaultSettings) SetAllowVideo(v bool)`

SetAllowVideo sets AllowVideo field to given value.

### HasAllowVideo

`func (o *DefaultSettings) HasAllowVideo() bool`

HasAllowVideo returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


