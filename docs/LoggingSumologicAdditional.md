# LoggingSumologicAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageType** | Pointer to [**LoggingMessageType**](LoggingMessageType.md) |  | [optional] [default to LOGGINGMESSAGETYPE_CLASSIC]
**Url** | Pointer to **string** | The URL to post logs to. | [optional] 

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

### GetUrl

`func (o *LoggingSumologicAdditional) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LoggingSumologicAdditional) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LoggingSumologicAdditional) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LoggingSumologicAdditional) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


