# LoggingSumologicAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageType** | Pointer to [**LoggingMessageType**](LoggingMessageType.md) |  | [optional] [default to LOGGINGMESSAGETYPE_CLASSIC]
**URL** | Pointer to **string** | The URL to post logs to. | [optional] 

## Methods

### NewLoggingSumologicAllOf

`func NewLoggingSumologicAllOf() *LoggingSumologicAllOf`

NewLoggingSumologicAllOf instantiates a new LoggingSumologicAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSumologicAllOfWithDefaults

`func NewLoggingSumologicAllOfWithDefaults() *LoggingSumologicAllOf`

NewLoggingSumologicAllOfWithDefaults instantiates a new LoggingSumologicAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageType

`func (o *LoggingSumologicAllOf) GetMessageType() LoggingMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingSumologicAllOf) GetMessageTypeOk() (*LoggingMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingSumologicAllOf) SetMessageType(v LoggingMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingSumologicAllOf) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetURL

`func (o *LoggingSumologicAllOf) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingSumologicAllOf) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingSumologicAllOf) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingSumologicAllOf) HasURL() bool`

HasURL returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
