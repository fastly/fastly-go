# RateLimiterResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | HTTP status code for custom limit enforcement response. | [optional] 
**ContentType** | Pointer to **string** | MIME type for custom limit enforcement response. | [optional] 
**Content** | Pointer to **string** | Response body for custom limit enforcement response. | [optional] 

## Methods

### NewRateLimiterResponse1

`func NewRateLimiterResponse1() *RateLimiterResponse1`

NewRateLimiterResponse1 instantiates a new RateLimiterResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimiterResponse1WithDefaults

`func NewRateLimiterResponse1WithDefaults() *RateLimiterResponse1`

NewRateLimiterResponse1WithDefaults instantiates a new RateLimiterResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RateLimiterResponse1) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RateLimiterResponse1) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RateLimiterResponse1) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RateLimiterResponse1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContentType

`func (o *RateLimiterResponse1) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *RateLimiterResponse1) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *RateLimiterResponse1) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *RateLimiterResponse1) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContent

`func (o *RateLimiterResponse1) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RateLimiterResponse1) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RateLimiterResponse1) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *RateLimiterResponse1) HasContent() bool`

HasContent returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
