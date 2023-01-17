# LoggingHTTPSAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** | The URL to send logs to. Must use HTTPS. Required. | [optional] 
**RequestMaxEntries** | Pointer to **int32** | The maximum number of logs sent in one request. Defaults `0` (10k). | [optional] [default to 0]
**RequestMaxBytes** | Pointer to **int32** | The maximum number of bytes sent in one request. Defaults `0` (100MB). | [optional] [default to 0]
**ContentType** | Pointer to **NullableString** | Content type of the header sent with the request. | [optional] [default to "null"]
**HeaderName** | Pointer to **NullableString** | Name of the custom header sent with the request. | [optional] [default to "null"]
**MessageType** | Pointer to [**LoggingMessageType**](LoggingMessageType.md) |  | [optional] [default to LOGGINGMESSAGETYPE_CLASSIC]
**HeaderValue** | Pointer to **NullableString** | Value of the custom header sent with the request. | [optional] [default to "null"]
**Method** | Pointer to **string** | HTTP method used for request. | [optional] [default to "POST"]
**JSONFormat** | Pointer to **string** | Enforces valid JSON formatting for log entries. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [optional] [default to "%h %l %u %t \"%r\" %&gt;s %b"]

## Methods

### NewLoggingHTTPSAllOf

`func NewLoggingHTTPSAllOf() *LoggingHTTPSAllOf`

NewLoggingHTTPSAllOf instantiates a new LoggingHTTPSAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingHTTPSAllOfWithDefaults

`func NewLoggingHTTPSAllOfWithDefaults() *LoggingHTTPSAllOf`

NewLoggingHTTPSAllOfWithDefaults instantiates a new LoggingHTTPSAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *LoggingHTTPSAllOf) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingHTTPSAllOf) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingHTTPSAllOf) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingHTTPSAllOf) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetRequestMaxEntries

`func (o *LoggingHTTPSAllOf) GetRequestMaxEntries() int32`

GetRequestMaxEntries returns the RequestMaxEntries field if non-nil, zero value otherwise.

### GetRequestMaxEntriesOk

`func (o *LoggingHTTPSAllOf) GetRequestMaxEntriesOk() (*int32, bool)`

GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxEntries

`func (o *LoggingHTTPSAllOf) SetRequestMaxEntries(v int32)`

SetRequestMaxEntries sets RequestMaxEntries field to given value.

### HasRequestMaxEntries

`func (o *LoggingHTTPSAllOf) HasRequestMaxEntries() bool`

HasRequestMaxEntries returns a boolean if a field has been set.

### GetRequestMaxBytes

`func (o *LoggingHTTPSAllOf) GetRequestMaxBytes() int32`

GetRequestMaxBytes returns the RequestMaxBytes field if non-nil, zero value otherwise.

### GetRequestMaxBytesOk

`func (o *LoggingHTTPSAllOf) GetRequestMaxBytesOk() (*int32, bool)`

GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxBytes

`func (o *LoggingHTTPSAllOf) SetRequestMaxBytes(v int32)`

SetRequestMaxBytes sets RequestMaxBytes field to given value.

### HasRequestMaxBytes

`func (o *LoggingHTTPSAllOf) HasRequestMaxBytes() bool`

HasRequestMaxBytes returns a boolean if a field has been set.

### GetContentType

`func (o *LoggingHTTPSAllOf) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LoggingHTTPSAllOf) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LoggingHTTPSAllOf) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *LoggingHTTPSAllOf) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *LoggingHTTPSAllOf) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *LoggingHTTPSAllOf) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetHeaderName

`func (o *LoggingHTTPSAllOf) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *LoggingHTTPSAllOf) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *LoggingHTTPSAllOf) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *LoggingHTTPSAllOf) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### SetHeaderNameNil

`func (o *LoggingHTTPSAllOf) SetHeaderNameNil(b bool)`

 SetHeaderNameNil sets the value for HeaderName to be an explicit nil

### UnsetHeaderName
`func (o *LoggingHTTPSAllOf) UnsetHeaderName()`

UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil
### GetMessageType

`func (o *LoggingHTTPSAllOf) GetMessageType() LoggingMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingHTTPSAllOf) GetMessageTypeOk() (*LoggingMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingHTTPSAllOf) SetMessageType(v LoggingMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingHTTPSAllOf) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetHeaderValue

`func (o *LoggingHTTPSAllOf) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *LoggingHTTPSAllOf) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *LoggingHTTPSAllOf) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.

### HasHeaderValue

`func (o *LoggingHTTPSAllOf) HasHeaderValue() bool`

HasHeaderValue returns a boolean if a field has been set.

### SetHeaderValueNil

`func (o *LoggingHTTPSAllOf) SetHeaderValueNil(b bool)`

 SetHeaderValueNil sets the value for HeaderValue to be an explicit nil

### UnsetHeaderValue
`func (o *LoggingHTTPSAllOf) UnsetHeaderValue()`

UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
### GetMethod

`func (o *LoggingHTTPSAllOf) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoggingHTTPSAllOf) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoggingHTTPSAllOf) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoggingHTTPSAllOf) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetJSONFormat

`func (o *LoggingHTTPSAllOf) GetJSONFormat() string`

GetJSONFormat returns the JSONFormat field if non-nil, zero value otherwise.

### GetJSONFormatOk

`func (o *LoggingHTTPSAllOf) GetJSONFormatOk() (*string, bool)`

GetJSONFormatOk returns a tuple with the JSONFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJSONFormat

`func (o *LoggingHTTPSAllOf) SetJSONFormat(v string)`

SetJSONFormat sets JSONFormat field to given value.

### HasJSONFormat

`func (o *LoggingHTTPSAllOf) HasJSONFormat() bool`

HasJSONFormat returns a boolean if a field has been set.

### GetFormat

`func (o *LoggingHTTPSAllOf) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingHTTPSAllOf) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingHTTPSAllOf) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingHTTPSAllOf) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
