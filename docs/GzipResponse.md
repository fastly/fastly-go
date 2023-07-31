# GzipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheCondition** | Pointer to **NullableString** | Name of the cache condition controlling when this configuration applies. | [optional] 
**ContentTypes** | Pointer to **NullableString** | Space-separated list of content types to compress. If you omit this field a default list will be used. | [optional] 
**Extensions** | Pointer to **NullableString** | Space-separated list of file extensions to compress. If you omit this field a default list will be used. | [optional] 
**Name** | Pointer to **string** | Name of the gzip configuration. | [optional] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 

## Methods

### NewGzipResponse

`func NewGzipResponse() *GzipResponse`

NewGzipResponse instantiates a new GzipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGzipResponseWithDefaults

`func NewGzipResponseWithDefaults() *GzipResponse`

NewGzipResponseWithDefaults instantiates a new GzipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCondition

`func (o *GzipResponse) GetCacheCondition() string`

GetCacheCondition returns the CacheCondition field if non-nil, zero value otherwise.

### GetCacheConditionOk

`func (o *GzipResponse) GetCacheConditionOk() (*string, bool)`

GetCacheConditionOk returns a tuple with the CacheCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCondition

`func (o *GzipResponse) SetCacheCondition(v string)`

SetCacheCondition sets CacheCondition field to given value.

### HasCacheCondition

`func (o *GzipResponse) HasCacheCondition() bool`

HasCacheCondition returns a boolean if a field has been set.

### SetCacheConditionNil

`func (o *GzipResponse) SetCacheConditionNil(b bool)`

 SetCacheConditionNil sets the value for CacheCondition to be an explicit nil

### UnsetCacheCondition
`func (o *GzipResponse) UnsetCacheCondition()`

UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
### GetContentTypes

`func (o *GzipResponse) GetContentTypes() string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *GzipResponse) GetContentTypesOk() (*string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *GzipResponse) SetContentTypes(v string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *GzipResponse) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### SetContentTypesNil

`func (o *GzipResponse) SetContentTypesNil(b bool)`

 SetContentTypesNil sets the value for ContentTypes to be an explicit nil

### UnsetContentTypes
`func (o *GzipResponse) UnsetContentTypes()`

UnsetContentTypes ensures that no value is present for ContentTypes, not even an explicit nil
### GetExtensions

`func (o *GzipResponse) GetExtensions() string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *GzipResponse) GetExtensionsOk() (*string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *GzipResponse) SetExtensions(v string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *GzipResponse) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *GzipResponse) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *GzipResponse) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetName

`func (o *GzipResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GzipResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GzipResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GzipResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceID

`func (o *GzipResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *GzipResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *GzipResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *GzipResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *GzipResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GzipResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GzipResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GzipResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GzipResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GzipResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GzipResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GzipResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GzipResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GzipResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *GzipResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GzipResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GzipResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GzipResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *GzipResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *GzipResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GzipResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GzipResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GzipResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GzipResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GzipResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GzipResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
