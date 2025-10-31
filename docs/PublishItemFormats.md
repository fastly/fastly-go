# PublishItemFormats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpResponse** | Pointer to [**HttpResponseFormat**](HttpResponseFormat.md) |  | [optional] 
**HttpStream** | Pointer to [**HttpStreamFormat**](HttpStreamFormat.md) |  | [optional] 
**WsMessage** | Pointer to [**WsMessageFormat**](WsMessageFormat.md) |  | [optional] 

## Methods

### NewPublishItemFormats

`func NewPublishItemFormats() *PublishItemFormats`

NewPublishItemFormats instantiates a new PublishItemFormats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishItemFormatsWithDefaults

`func NewPublishItemFormatsWithDefaults() *PublishItemFormats`

NewPublishItemFormatsWithDefaults instantiates a new PublishItemFormats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpResponse

`func (o *PublishItemFormats) GetHttpResponse() HttpResponseFormat`

GetHttpResponse returns the HttpResponse field if non-nil, zero value otherwise.

### GetHttpResponseOk

`func (o *PublishItemFormats) GetHttpResponseOk() (*HttpResponseFormat, bool)`

GetHttpResponseOk returns a tuple with the HttpResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponse

`func (o *PublishItemFormats) SetHttpResponse(v HttpResponseFormat)`

SetHttpResponse sets HttpResponse field to given value.

### HasHttpResponse

`func (o *PublishItemFormats) HasHttpResponse() bool`

HasHttpResponse returns a boolean if a field has been set.

### GetHttpStream

`func (o *PublishItemFormats) GetHttpStream() HttpStreamFormat`

GetHttpStream returns the HttpStream field if non-nil, zero value otherwise.

### GetHttpStreamOk

`func (o *PublishItemFormats) GetHttpStreamOk() (*HttpStreamFormat, bool)`

GetHttpStreamOk returns a tuple with the HttpStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStream

`func (o *PublishItemFormats) SetHttpStream(v HttpStreamFormat)`

SetHttpStream sets HttpStream field to given value.

### HasHttpStream

`func (o *PublishItemFormats) HasHttpStream() bool`

HasHttpStream returns a boolean if a field has been set.

### GetWsMessage

`func (o *PublishItemFormats) GetWsMessage() WsMessageFormat`

GetWsMessage returns the WsMessage field if non-nil, zero value otherwise.

### GetWsMessageOk

`func (o *PublishItemFormats) GetWsMessageOk() (*WsMessageFormat, bool)`

GetWsMessageOk returns a tuple with the WsMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsMessage

`func (o *PublishItemFormats) SetWsMessage(v WsMessageFormat)`

SetWsMessage sets WsMessage field to given value.

### HasWsMessage

`func (o *PublishItemFormats) HasWsMessage() bool`

HasWsMessage returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


