# PublishItemFormats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HTTPResponse** | Pointer to [**HTTPResponseFormat**](HttpResponseFormat.md) |  | [optional] 
**HTTPStream** | Pointer to [**HTTPStreamFormat**](HttpStreamFormat.md) |  | [optional] 
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

### GetHTTPResponse

`func (o *PublishItemFormats) GetHTTPResponse() HTTPResponseFormat`

GetHTTPResponse returns the HTTPResponse field if non-nil, zero value otherwise.

### GetHTTPResponseOk

`func (o *PublishItemFormats) GetHTTPResponseOk() (*HTTPResponseFormat, bool)`

GetHTTPResponseOk returns a tuple with the HTTPResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPResponse

`func (o *PublishItemFormats) SetHTTPResponse(v HTTPResponseFormat)`

SetHTTPResponse sets HTTPResponse field to given value.

### HasHTTPResponse

`func (o *PublishItemFormats) HasHTTPResponse() bool`

HasHTTPResponse returns a boolean if a field has been set.

### GetHTTPStream

`func (o *PublishItemFormats) GetHTTPStream() HTTPStreamFormat`

GetHTTPStream returns the HTTPStream field if non-nil, zero value otherwise.

### GetHTTPStreamOk

`func (o *PublishItemFormats) GetHTTPStreamOk() (*HTTPStreamFormat, bool)`

GetHTTPStreamOk returns a tuple with the HTTPStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPStream

`func (o *PublishItemFormats) SetHTTPStream(v HTTPStreamFormat)`

SetHTTPStream sets HTTPStream field to given value.

### HasHTTPStream

`func (o *PublishItemFormats) HasHTTPStream() bool`

HasHTTPStream returns a boolean if a field has been set.

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
