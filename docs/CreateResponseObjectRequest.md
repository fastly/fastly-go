# CreateResponseObjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the response object to create. | [optional] 
**Status** | Pointer to **string** | The status code the response will have. Defaults to 200. | [optional] 
**Response** | Pointer to **string** | The status text the response will have. Defaults to &#39;OK&#39;. | [optional] 
**Content** | Pointer to **string** | The content the response will deliver. | [optional] 
**ContentType** | Pointer to **NullableString** | The MIME type of your response content. | [optional] 
**RequestCondition** | Pointer to **NullableString** | Condition which, if met, will select this configuration during a request. Optional. | [optional] 
**CacheCondition** | Pointer to **NullableString** | Name of the cache condition controlling when this configuration applies. | [optional] 

## Methods

### NewCreateResponseObjectRequest

`func NewCreateResponseObjectRequest() *CreateResponseObjectRequest`

NewCreateResponseObjectRequest instantiates a new CreateResponseObjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResponseObjectRequestWithDefaults

`func NewCreateResponseObjectRequestWithDefaults() *CreateResponseObjectRequest`

NewCreateResponseObjectRequestWithDefaults instantiates a new CreateResponseObjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateResponseObjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResponseObjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResponseObjectRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateResponseObjectRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *CreateResponseObjectRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateResponseObjectRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateResponseObjectRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateResponseObjectRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResponse

`func (o *CreateResponseObjectRequest) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *CreateResponseObjectRequest) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *CreateResponseObjectRequest) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *CreateResponseObjectRequest) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetContent

`func (o *CreateResponseObjectRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateResponseObjectRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateResponseObjectRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CreateResponseObjectRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentType

`func (o *CreateResponseObjectRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CreateResponseObjectRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CreateResponseObjectRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CreateResponseObjectRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *CreateResponseObjectRequest) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *CreateResponseObjectRequest) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetRequestCondition

`func (o *CreateResponseObjectRequest) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *CreateResponseObjectRequest) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *CreateResponseObjectRequest) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *CreateResponseObjectRequest) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *CreateResponseObjectRequest) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *CreateResponseObjectRequest) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetCacheCondition

`func (o *CreateResponseObjectRequest) GetCacheCondition() string`

GetCacheCondition returns the CacheCondition field if non-nil, zero value otherwise.

### GetCacheConditionOk

`func (o *CreateResponseObjectRequest) GetCacheConditionOk() (*string, bool)`

GetCacheConditionOk returns a tuple with the CacheCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCondition

`func (o *CreateResponseObjectRequest) SetCacheCondition(v string)`

SetCacheCondition sets CacheCondition field to given value.

### HasCacheCondition

`func (o *CreateResponseObjectRequest) HasCacheCondition() bool`

HasCacheCondition returns a boolean if a field has been set.

### SetCacheConditionNil

`func (o *CreateResponseObjectRequest) SetCacheConditionNil(b bool)`

 SetCacheConditionNil sets the value for CacheCondition to be an explicit nil

### UnsetCacheCondition
`func (o *CreateResponseObjectRequest) UnsetCacheCondition()`

UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
