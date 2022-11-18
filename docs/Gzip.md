# Gzip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheCondition** | Pointer to **NullableString** | Name of the cache condition controlling when this configuration applies. | [optional] 
**ContentTypes** | Pointer to **NullableString** | Space-separated list of content types to compress. If you omit this field a default list will be used. | [optional] 
**Extensions** | Pointer to **NullableString** | Space-separated list of file extensions to compress. If you omit this field a default list will be used. | [optional] 
**Name** | Pointer to **string** | Name of the gzip configuration. | [optional] 

## Methods

### NewGzip

`func NewGzip() *Gzip`

NewGzip instantiates a new Gzip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGzipWithDefaults

`func NewGzipWithDefaults() *Gzip`

NewGzipWithDefaults instantiates a new Gzip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCondition

`func (o *Gzip) GetCacheCondition() string`

GetCacheCondition returns the CacheCondition field if non-nil, zero value otherwise.

### GetCacheConditionOk

`func (o *Gzip) GetCacheConditionOk() (*string, bool)`

GetCacheConditionOk returns a tuple with the CacheCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCondition

`func (o *Gzip) SetCacheCondition(v string)`

SetCacheCondition sets CacheCondition field to given value.

### HasCacheCondition

`func (o *Gzip) HasCacheCondition() bool`

HasCacheCondition returns a boolean if a field has been set.

### SetCacheConditionNil

`func (o *Gzip) SetCacheConditionNil(b bool)`

 SetCacheConditionNil sets the value for CacheCondition to be an explicit nil

### UnsetCacheCondition
`func (o *Gzip) UnsetCacheCondition()`

UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
### GetContentTypes

`func (o *Gzip) GetContentTypes() string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *Gzip) GetContentTypesOk() (*string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *Gzip) SetContentTypes(v string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *Gzip) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### SetContentTypesNil

`func (o *Gzip) SetContentTypesNil(b bool)`

 SetContentTypesNil sets the value for ContentTypes to be an explicit nil

### UnsetContentTypes
`func (o *Gzip) UnsetContentTypes()`

UnsetContentTypes ensures that no value is present for ContentTypes, not even an explicit nil
### GetExtensions

`func (o *Gzip) GetExtensions() string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Gzip) GetExtensionsOk() (*string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Gzip) SetExtensions(v string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Gzip) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *Gzip) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *Gzip) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetName

`func (o *Gzip) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gzip) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gzip) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Gzip) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
