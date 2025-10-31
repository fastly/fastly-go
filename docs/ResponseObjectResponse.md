# ResponseObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheCondition** | Pointer to **NullableString** | Name of the cache condition controlling when this configuration applies. | [optional] 
**Content** | Pointer to **string** | The content to deliver for the response object, can be empty. | [optional] 
**ContentType** | Pointer to **NullableString** | The MIME type of the content, can be empty. | [optional] 
**Name** | Pointer to **string** | Name for the request settings. | [optional] 
**Status** | Pointer to **string** | The HTTP status code. | [optional] [default to "200"]
**Response** | Pointer to **string** | The HTTP response. | [optional] [default to "Ok"]
**RequestCondition** | Pointer to **NullableString** | Condition which, if met, will select this configuration during a request. Optional. | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 

## Methods

### NewResponseObjectResponse

`func NewResponseObjectResponse() *ResponseObjectResponse`

NewResponseObjectResponse instantiates a new ResponseObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseObjectResponseWithDefaults

`func NewResponseObjectResponseWithDefaults() *ResponseObjectResponse`

NewResponseObjectResponseWithDefaults instantiates a new ResponseObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCondition

`func (o *ResponseObjectResponse) GetCacheCondition() string`

GetCacheCondition returns the CacheCondition field if non-nil, zero value otherwise.

### GetCacheConditionOk

`func (o *ResponseObjectResponse) GetCacheConditionOk() (*string, bool)`

GetCacheConditionOk returns a tuple with the CacheCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCondition

`func (o *ResponseObjectResponse) SetCacheCondition(v string)`

SetCacheCondition sets CacheCondition field to given value.

### HasCacheCondition

`func (o *ResponseObjectResponse) HasCacheCondition() bool`

HasCacheCondition returns a boolean if a field has been set.

### SetCacheConditionNil

`func (o *ResponseObjectResponse) SetCacheConditionNil(b bool)`

 SetCacheConditionNil sets the value for CacheCondition to be an explicit nil

### UnsetCacheCondition
`func (o *ResponseObjectResponse) UnsetCacheCondition()`

UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
### GetContent

`func (o *ResponseObjectResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ResponseObjectResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ResponseObjectResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ResponseObjectResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentType

`func (o *ResponseObjectResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ResponseObjectResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ResponseObjectResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ResponseObjectResponse) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ResponseObjectResponse) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ResponseObjectResponse) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetName

`func (o *ResponseObjectResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseObjectResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseObjectResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseObjectResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseObjectResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseObjectResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseObjectResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseObjectResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResponse

`func (o *ResponseObjectResponse) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ResponseObjectResponse) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ResponseObjectResponse) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ResponseObjectResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetRequestCondition

`func (o *ResponseObjectResponse) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *ResponseObjectResponse) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *ResponseObjectResponse) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *ResponseObjectResponse) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *ResponseObjectResponse) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *ResponseObjectResponse) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetServiceId

`func (o *ResponseObjectResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ResponseObjectResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ResponseObjectResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ResponseObjectResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetVersion

`func (o *ResponseObjectResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponseObjectResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponseObjectResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponseObjectResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponseObjectResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseObjectResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseObjectResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResponseObjectResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ResponseObjectResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ResponseObjectResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *ResponseObjectResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ResponseObjectResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ResponseObjectResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ResponseObjectResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ResponseObjectResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ResponseObjectResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ResponseObjectResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResponseObjectResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResponseObjectResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResponseObjectResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ResponseObjectResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ResponseObjectResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


