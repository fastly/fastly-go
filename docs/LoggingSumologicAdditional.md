# LoggingSumologicAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageType** | Pointer to [**LoggingMessageType**](LoggingMessageType.md) |  | [optional] [default to LOGGINGMESSAGETYPE_CLASSIC]
**URL** | Pointer to **string** | The URL to post logs to. | [optional] 

## Methods

### NewLoggingSumologicAdditional

`func NewLoggingSumologicAdditional() *LoggingSumologicAdditional`

NewLoggingSumologicAdditional instantiates a new LoggingSumologicAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSumologicAdditionalWithDefaults

`func NewLoggingSumologicAdditionalWithDefaults() *LoggingSumologicAdditional`

NewLoggingSumologicAdditionalWithDefaults instantiates a new LoggingSumologicAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageType

`func (o *LoggingSumologicAdditional) GetMessageType() LoggingMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingSumologicAdditional) GetMessageTypeOk() (*LoggingMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingSumologicAdditional) SetMessageType(v LoggingMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingSumologicAdditional) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetURL

`func (o *LoggingSumologicAdditional) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingSumologicAdditional) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingSumologicAdditional) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingSumologicAdditional) HasURL() bool`

HasURL returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
