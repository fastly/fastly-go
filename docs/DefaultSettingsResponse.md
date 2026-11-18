# DefaultSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResizeFilter** | **string** | The type of filter to use while resizing an image. | [default to "lanczos3"]
**Webp** | **bool** | Controls whether or not to default to WebP output when the client supports it. This is equivalent to adding \&quot;auto&#x3D;webp\&quot; to all image optimizer requests.  | [default to false]
**WebpQuality** | **int32** | The default quality to use with WebP output. This can be overridden with the second option in the \&quot;quality\&quot; URL parameter on specific image optimizer requests.  | [default to 85]
**JpegType** | **string** | The default type of JPEG output to use. This can be overridden with \&quot;format&#x3D;bjpeg\&quot; and \&quot;format&#x3D;pjpeg\&quot; on specific image optimizer requests.  | [default to "auto"]
**JpegQuality** | **int32** | The default quality to use with JPEG output. This can be overridden with the \&quot;quality\&quot; parameter on specific image optimizer requests.  | [default to 85]
**Upscale** | **bool** | Whether or not we should allow output images to render at sizes larger than input.  | [default to false]
**AllowVideo** | **bool** | Enables GIF to MP4 transformations on this service. | [default to false]

## Methods

### NewDefaultSettingsResponse

`func NewDefaultSettingsResponse(resizeFilter string, webp bool, webpQuality int32, jpegType string, jpegQuality int32, upscale bool, allowVideo bool, ) *DefaultSettingsResponse`

NewDefaultSettingsResponse instantiates a new DefaultSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultSettingsResponseWithDefaults

`func NewDefaultSettingsResponseWithDefaults() *DefaultSettingsResponse`

NewDefaultSettingsResponseWithDefaults instantiates a new DefaultSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResizeFilter

`func (o *DefaultSettingsResponse) GetResizeFilter() string`

GetResizeFilter returns the ResizeFilter field if non-nil, zero value otherwise.

### GetResizeFilterOk

`func (o *DefaultSettingsResponse) GetResizeFilterOk() (*string, bool)`

GetResizeFilterOk returns a tuple with the ResizeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeFilter

`func (o *DefaultSettingsResponse) SetResizeFilter(v string)`

SetResizeFilter sets ResizeFilter field to given value.


### GetWebp

`func (o *DefaultSettingsResponse) GetWebp() bool`

GetWebp returns the Webp field if non-nil, zero value otherwise.

### GetWebpOk

`func (o *DefaultSettingsResponse) GetWebpOk() (*bool, bool)`

GetWebpOk returns a tuple with the Webp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebp

`func (o *DefaultSettingsResponse) SetWebp(v bool)`

SetWebp sets Webp field to given value.


### GetWebpQuality

`func (o *DefaultSettingsResponse) GetWebpQuality() int32`

GetWebpQuality returns the WebpQuality field if non-nil, zero value otherwise.

### GetWebpQualityOk

`func (o *DefaultSettingsResponse) GetWebpQualityOk() (*int32, bool)`

GetWebpQualityOk returns a tuple with the WebpQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebpQuality

`func (o *DefaultSettingsResponse) SetWebpQuality(v int32)`

SetWebpQuality sets WebpQuality field to given value.


### GetJpegType

`func (o *DefaultSettingsResponse) GetJpegType() string`

GetJpegType returns the JpegType field if non-nil, zero value otherwise.

### GetJpegTypeOk

`func (o *DefaultSettingsResponse) GetJpegTypeOk() (*string, bool)`

GetJpegTypeOk returns a tuple with the JpegType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpegType

`func (o *DefaultSettingsResponse) SetJpegType(v string)`

SetJpegType sets JpegType field to given value.


### GetJpegQuality

`func (o *DefaultSettingsResponse) GetJpegQuality() int32`

GetJpegQuality returns the JpegQuality field if non-nil, zero value otherwise.

### GetJpegQualityOk

`func (o *DefaultSettingsResponse) GetJpegQualityOk() (*int32, bool)`

GetJpegQualityOk returns a tuple with the JpegQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpegQuality

`func (o *DefaultSettingsResponse) SetJpegQuality(v int32)`

SetJpegQuality sets JpegQuality field to given value.


### GetUpscale

`func (o *DefaultSettingsResponse) GetUpscale() bool`

GetUpscale returns the Upscale field if non-nil, zero value otherwise.

### GetUpscaleOk

`func (o *DefaultSettingsResponse) GetUpscaleOk() (*bool, bool)`

GetUpscaleOk returns a tuple with the Upscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpscale

`func (o *DefaultSettingsResponse) SetUpscale(v bool)`

SetUpscale sets Upscale field to given value.


### GetAllowVideo

`func (o *DefaultSettingsResponse) GetAllowVideo() bool`

GetAllowVideo returns the AllowVideo field if non-nil, zero value otherwise.

### GetAllowVideoOk

`func (o *DefaultSettingsResponse) GetAllowVideoOk() (*bool, bool)`

GetAllowVideoOk returns a tuple with the AllowVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVideo

`func (o *DefaultSettingsResponse) SetAllowVideo(v bool)`

SetAllowVideo sets AllowVideo field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
