# ResponseObject

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

## Methods

### NewResponseObject

`func NewResponseObject() *ResponseObject`

NewResponseObject instantiates a new ResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseObjectWithDefaults

`func NewResponseObjectWithDefaults() *ResponseObject`

NewResponseObjectWithDefaults instantiates a new ResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCondition

`func (o *ResponseObject) GetCacheCondition() string`

GetCacheCondition returns the CacheCondition field if non-nil, zero value otherwise.

### GetCacheConditionOk

`func (o *ResponseObject) GetCacheConditionOk() (*string, bool)`

GetCacheConditionOk returns a tuple with the CacheCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCondition

`func (o *ResponseObject) SetCacheCondition(v string)`

SetCacheCondition sets CacheCondition field to given value.

### HasCacheCondition

`func (o *ResponseObject) HasCacheCondition() bool`

HasCacheCondition returns a boolean if a field has been set.

### SetCacheConditionNil

`func (o *ResponseObject) SetCacheConditionNil(b bool)`

 SetCacheConditionNil sets the value for CacheCondition to be an explicit nil

### UnsetCacheCondition
`func (o *ResponseObject) UnsetCacheCondition()`

UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
### GetContent

`func (o *ResponseObject) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ResponseObject) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ResponseObject) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ResponseObject) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentType

`func (o *ResponseObject) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ResponseObject) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ResponseObject) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ResponseObject) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ResponseObject) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ResponseObject) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetName

`func (o *ResponseObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResponse

`func (o *ResponseObject) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ResponseObject) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ResponseObject) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ResponseObject) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetRequestCondition

`func (o *ResponseObject) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *ResponseObject) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *ResponseObject) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *ResponseObject) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *ResponseObject) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *ResponseObject) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


