# LoggingHttpsAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL to send logs to. Must use HTTPS. Required. | [optional] 
**RequestMaxEntries** | Pointer to **int32** | The maximum number of logs sent in one request. Defaults `0` (10k). | [optional] [default to 0]
**RequestMaxBytes** | Pointer to **int32** | The maximum number of bytes sent in one request. Defaults `0` (100MB). | [optional] [default to 0]
**ContentType** | Pointer to **NullableString** | Content type of the header sent with the request. | [optional] [default to "null"]
**HeaderName** | Pointer to **NullableString** | Name of the custom header sent with the request. | [optional] [default to "null"]
**MessageType** | Pointer to [**LoggingMessageType**](LoggingMessageType.md) |  | [optional] [default to LOGGINGMESSAGETYPE_CLASSIC]
**HeaderValue** | Pointer to **NullableString** | Value of the custom header sent with the request. | [optional] [default to "null"]
**Method** | Pointer to **string** | HTTP method used for request. | [optional] [default to "POST"]
**JsonFormat** | Pointer to **string** | Enforces valid JSON formatting for log entries. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [optional] [default to "%h %l %u %t \"%r\" %&gt;s %b"]
**Period** | Pointer to **int32** | How frequently, in seconds, batches of log data are sent to the HTTPS endpoint. A value of `0` sends logs at the same interval as the default, which is `5` seconds. | [optional] [default to 5]

## Methods

### NewLoggingHttpsAdditional

`func NewLoggingHttpsAdditional() *LoggingHttpsAdditional`

NewLoggingHttpsAdditional instantiates a new LoggingHttpsAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingHttpsAdditionalWithDefaults

`func NewLoggingHttpsAdditionalWithDefaults() *LoggingHttpsAdditional`

NewLoggingHttpsAdditionalWithDefaults instantiates a new LoggingHttpsAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *LoggingHttpsAdditional) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LoggingHttpsAdditional) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LoggingHttpsAdditional) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LoggingHttpsAdditional) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequestMaxEntries

`func (o *LoggingHttpsAdditional) GetRequestMaxEntries() int32`

GetRequestMaxEntries returns the RequestMaxEntries field if non-nil, zero value otherwise.

### GetRequestMaxEntriesOk

`func (o *LoggingHttpsAdditional) GetRequestMaxEntriesOk() (*int32, bool)`

GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxEntries

`func (o *LoggingHttpsAdditional) SetRequestMaxEntries(v int32)`

SetRequestMaxEntries sets RequestMaxEntries field to given value.

### HasRequestMaxEntries

`func (o *LoggingHttpsAdditional) HasRequestMaxEntries() bool`

HasRequestMaxEntries returns a boolean if a field has been set.

### GetRequestMaxBytes

`func (o *LoggingHttpsAdditional) GetRequestMaxBytes() int32`

GetRequestMaxBytes returns the RequestMaxBytes field if non-nil, zero value otherwise.

### GetRequestMaxBytesOk

`func (o *LoggingHttpsAdditional) GetRequestMaxBytesOk() (*int32, bool)`

GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxBytes

`func (o *LoggingHttpsAdditional) SetRequestMaxBytes(v int32)`

SetRequestMaxBytes sets RequestMaxBytes field to given value.

### HasRequestMaxBytes

`func (o *LoggingHttpsAdditional) HasRequestMaxBytes() bool`

HasRequestMaxBytes returns a boolean if a field has been set.

### GetContentType

`func (o *LoggingHttpsAdditional) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LoggingHttpsAdditional) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LoggingHttpsAdditional) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *LoggingHttpsAdditional) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *LoggingHttpsAdditional) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *LoggingHttpsAdditional) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetHeaderName

`func (o *LoggingHttpsAdditional) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *LoggingHttpsAdditional) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *LoggingHttpsAdditional) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *LoggingHttpsAdditional) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### SetHeaderNameNil

`func (o *LoggingHttpsAdditional) SetHeaderNameNil(b bool)`

 SetHeaderNameNil sets the value for HeaderName to be an explicit nil

### UnsetHeaderName
`func (o *LoggingHttpsAdditional) UnsetHeaderName()`

UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil
### GetMessageType

`func (o *LoggingHttpsAdditional) GetMessageType() LoggingMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingHttpsAdditional) GetMessageTypeOk() (*LoggingMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingHttpsAdditional) SetMessageType(v LoggingMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingHttpsAdditional) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetHeaderValue

`func (o *LoggingHttpsAdditional) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *LoggingHttpsAdditional) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *LoggingHttpsAdditional) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.

### HasHeaderValue

`func (o *LoggingHttpsAdditional) HasHeaderValue() bool`

HasHeaderValue returns a boolean if a field has been set.

### SetHeaderValueNil

`func (o *LoggingHttpsAdditional) SetHeaderValueNil(b bool)`

 SetHeaderValueNil sets the value for HeaderValue to be an explicit nil

### UnsetHeaderValue
`func (o *LoggingHttpsAdditional) UnsetHeaderValue()`

UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
### GetMethod

`func (o *LoggingHttpsAdditional) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoggingHttpsAdditional) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoggingHttpsAdditional) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoggingHttpsAdditional) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetJsonFormat

`func (o *LoggingHttpsAdditional) GetJsonFormat() string`

GetJsonFormat returns the JsonFormat field if non-nil, zero value otherwise.

### GetJsonFormatOk

`func (o *LoggingHttpsAdditional) GetJsonFormatOk() (*string, bool)`

GetJsonFormatOk returns a tuple with the JsonFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonFormat

`func (o *LoggingHttpsAdditional) SetJsonFormat(v string)`

SetJsonFormat sets JsonFormat field to given value.

### HasJsonFormat

`func (o *LoggingHttpsAdditional) HasJsonFormat() bool`

HasJsonFormat returns a boolean if a field has been set.

### GetFormat

`func (o *LoggingHttpsAdditional) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingHttpsAdditional) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingHttpsAdditional) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingHttpsAdditional) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPeriod

`func (o *LoggingHttpsAdditional) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LoggingHttpsAdditional) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LoggingHttpsAdditional) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LoggingHttpsAdditional) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


