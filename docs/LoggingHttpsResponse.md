# LoggingHttpsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to **NullableString** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  | [optional] 
**ResponseCondition** | Pointer to **NullableString** | The name of an existing condition in the configured endpoint, or leave blank to always execute. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [optional] [default to "%h %l %u %t \"%r\" %&gt;s %b"]
**LogProcessingRegion** | Pointer to **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [optional] [default to "none"]
**FormatVersion** | Pointer to **string** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to "2"]
**TlsCaCert** | Pointer to **NullableString** | A secure certificate to authenticate a server with. Must be in PEM format. | [optional] [default to "null"]
**TlsClientCert** | Pointer to **NullableString** | The client certificate used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TlsClientKey** | Pointer to **NullableString** | The client private key used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TlsHostname** | Pointer to **NullableString** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [optional] [default to "null"]
**RequestMaxEntries** | Pointer to **int32** | The maximum number of logs sent in one request. Defaults `0` (10k). | [optional] [default to 0]
**RequestMaxBytes** | Pointer to **int32** | The maximum number of bytes sent in one request. Defaults `0` (100MB). | [optional] [default to 0]
**Url** | Pointer to **string** | The URL to send logs to. Must use HTTPS. Required. | [optional] 
**ContentType** | Pointer to **NullableString** | Content type of the header sent with the request. | [optional] [default to "null"]
**HeaderName** | Pointer to **NullableString** | Name of the custom header sent with the request. | [optional] [default to "null"]
**MessageType** | Pointer to [**LoggingMessageType**](LoggingMessageType.md) |  | [optional] [default to LOGGINGMESSAGETYPE_CLASSIC]
**HeaderValue** | Pointer to **NullableString** | Value of the custom header sent with the request. | [optional] [default to "null"]
**Method** | Pointer to **string** | HTTP method used for request. | [optional] [default to "POST"]
**JsonFormat** | Pointer to **string** | Enforces valid JSON formatting for log entries. | [optional] 
**Period** | Pointer to **int32** | How frequently, in seconds, batches of log data are sent to the HTTPS endpoint. A value of `0` sends logs at the same interval as the default, which is `5` seconds. | [optional] [default to 5]
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceId** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLoggingHttpsResponse

`func NewLoggingHttpsResponse() *LoggingHttpsResponse`

NewLoggingHttpsResponse instantiates a new LoggingHttpsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingHttpsResponseWithDefaults

`func NewLoggingHttpsResponseWithDefaults() *LoggingHttpsResponse`

NewLoggingHttpsResponseWithDefaults instantiates a new LoggingHttpsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingHttpsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingHttpsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingHttpsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingHttpsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingHttpsResponse) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingHttpsResponse) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingHttpsResponse) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingHttpsResponse) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingHttpsResponse) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingHttpsResponse) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetResponseCondition

`func (o *LoggingHttpsResponse) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingHttpsResponse) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingHttpsResponse) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingHttpsResponse) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingHttpsResponse) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingHttpsResponse) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingHttpsResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingHttpsResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingHttpsResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingHttpsResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetLogProcessingRegion

`func (o *LoggingHttpsResponse) GetLogProcessingRegion() string`

GetLogProcessingRegion returns the LogProcessingRegion field if non-nil, zero value otherwise.

### GetLogProcessingRegionOk

`func (o *LoggingHttpsResponse) GetLogProcessingRegionOk() (*string, bool)`

GetLogProcessingRegionOk returns a tuple with the LogProcessingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogProcessingRegion

`func (o *LoggingHttpsResponse) SetLogProcessingRegion(v string)`

SetLogProcessingRegion sets LogProcessingRegion field to given value.

### HasLogProcessingRegion

`func (o *LoggingHttpsResponse) HasLogProcessingRegion() bool`

HasLogProcessingRegion returns a boolean if a field has been set.

### GetFormatVersion

`func (o *LoggingHttpsResponse) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingHttpsResponse) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingHttpsResponse) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingHttpsResponse) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetTlsCaCert

`func (o *LoggingHttpsResponse) GetTlsCaCert() string`

GetTlsCaCert returns the TlsCaCert field if non-nil, zero value otherwise.

### GetTlsCaCertOk

`func (o *LoggingHttpsResponse) GetTlsCaCertOk() (*string, bool)`

GetTlsCaCertOk returns a tuple with the TlsCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCaCert

`func (o *LoggingHttpsResponse) SetTlsCaCert(v string)`

SetTlsCaCert sets TlsCaCert field to given value.

### HasTlsCaCert

`func (o *LoggingHttpsResponse) HasTlsCaCert() bool`

HasTlsCaCert returns a boolean if a field has been set.

### SetTlsCaCertNil

`func (o *LoggingHttpsResponse) SetTlsCaCertNil(b bool)`

 SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil

### UnsetTlsCaCert
`func (o *LoggingHttpsResponse) UnsetTlsCaCert()`

UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
### GetTlsClientCert

`func (o *LoggingHttpsResponse) GetTlsClientCert() string`

GetTlsClientCert returns the TlsClientCert field if non-nil, zero value otherwise.

### GetTlsClientCertOk

`func (o *LoggingHttpsResponse) GetTlsClientCertOk() (*string, bool)`

GetTlsClientCertOk returns a tuple with the TlsClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCert

`func (o *LoggingHttpsResponse) SetTlsClientCert(v string)`

SetTlsClientCert sets TlsClientCert field to given value.

### HasTlsClientCert

`func (o *LoggingHttpsResponse) HasTlsClientCert() bool`

HasTlsClientCert returns a boolean if a field has been set.

### SetTlsClientCertNil

`func (o *LoggingHttpsResponse) SetTlsClientCertNil(b bool)`

 SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil

### UnsetTlsClientCert
`func (o *LoggingHttpsResponse) UnsetTlsClientCert()`

UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
### GetTlsClientKey

`func (o *LoggingHttpsResponse) GetTlsClientKey() string`

GetTlsClientKey returns the TlsClientKey field if non-nil, zero value otherwise.

### GetTlsClientKeyOk

`func (o *LoggingHttpsResponse) GetTlsClientKeyOk() (*string, bool)`

GetTlsClientKeyOk returns a tuple with the TlsClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientKey

`func (o *LoggingHttpsResponse) SetTlsClientKey(v string)`

SetTlsClientKey sets TlsClientKey field to given value.

### HasTlsClientKey

`func (o *LoggingHttpsResponse) HasTlsClientKey() bool`

HasTlsClientKey returns a boolean if a field has been set.

### SetTlsClientKeyNil

`func (o *LoggingHttpsResponse) SetTlsClientKeyNil(b bool)`

 SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil

### UnsetTlsClientKey
`func (o *LoggingHttpsResponse) UnsetTlsClientKey()`

UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
### GetTlsHostname

`func (o *LoggingHttpsResponse) GetTlsHostname() string`

GetTlsHostname returns the TlsHostname field if non-nil, zero value otherwise.

### GetTlsHostnameOk

`func (o *LoggingHttpsResponse) GetTlsHostnameOk() (*string, bool)`

GetTlsHostnameOk returns a tuple with the TlsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsHostname

`func (o *LoggingHttpsResponse) SetTlsHostname(v string)`

SetTlsHostname sets TlsHostname field to given value.

### HasTlsHostname

`func (o *LoggingHttpsResponse) HasTlsHostname() bool`

HasTlsHostname returns a boolean if a field has been set.

### SetTlsHostnameNil

`func (o *LoggingHttpsResponse) SetTlsHostnameNil(b bool)`

 SetTlsHostnameNil sets the value for TlsHostname to be an explicit nil

### UnsetTlsHostname
`func (o *LoggingHttpsResponse) UnsetTlsHostname()`

UnsetTlsHostname ensures that no value is present for TlsHostname, not even an explicit nil
### GetRequestMaxEntries

`func (o *LoggingHttpsResponse) GetRequestMaxEntries() int32`

GetRequestMaxEntries returns the RequestMaxEntries field if non-nil, zero value otherwise.

### GetRequestMaxEntriesOk

`func (o *LoggingHttpsResponse) GetRequestMaxEntriesOk() (*int32, bool)`

GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxEntries

`func (o *LoggingHttpsResponse) SetRequestMaxEntries(v int32)`

SetRequestMaxEntries sets RequestMaxEntries field to given value.

### HasRequestMaxEntries

`func (o *LoggingHttpsResponse) HasRequestMaxEntries() bool`

HasRequestMaxEntries returns a boolean if a field has been set.

### GetRequestMaxBytes

`func (o *LoggingHttpsResponse) GetRequestMaxBytes() int32`

GetRequestMaxBytes returns the RequestMaxBytes field if non-nil, zero value otherwise.

### GetRequestMaxBytesOk

`func (o *LoggingHttpsResponse) GetRequestMaxBytesOk() (*int32, bool)`

GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxBytes

`func (o *LoggingHttpsResponse) SetRequestMaxBytes(v int32)`

SetRequestMaxBytes sets RequestMaxBytes field to given value.

### HasRequestMaxBytes

`func (o *LoggingHttpsResponse) HasRequestMaxBytes() bool`

HasRequestMaxBytes returns a boolean if a field has been set.

### GetUrl

`func (o *LoggingHttpsResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LoggingHttpsResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LoggingHttpsResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LoggingHttpsResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContentType

`func (o *LoggingHttpsResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LoggingHttpsResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LoggingHttpsResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *LoggingHttpsResponse) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *LoggingHttpsResponse) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *LoggingHttpsResponse) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetHeaderName

`func (o *LoggingHttpsResponse) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *LoggingHttpsResponse) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *LoggingHttpsResponse) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *LoggingHttpsResponse) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### SetHeaderNameNil

`func (o *LoggingHttpsResponse) SetHeaderNameNil(b bool)`

 SetHeaderNameNil sets the value for HeaderName to be an explicit nil

### UnsetHeaderName
`func (o *LoggingHttpsResponse) UnsetHeaderName()`

UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil
### GetMessageType

`func (o *LoggingHttpsResponse) GetMessageType() LoggingMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingHttpsResponse) GetMessageTypeOk() (*LoggingMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingHttpsResponse) SetMessageType(v LoggingMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingHttpsResponse) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetHeaderValue

`func (o *LoggingHttpsResponse) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *LoggingHttpsResponse) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *LoggingHttpsResponse) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.

### HasHeaderValue

`func (o *LoggingHttpsResponse) HasHeaderValue() bool`

HasHeaderValue returns a boolean if a field has been set.

### SetHeaderValueNil

`func (o *LoggingHttpsResponse) SetHeaderValueNil(b bool)`

 SetHeaderValueNil sets the value for HeaderValue to be an explicit nil

### UnsetHeaderValue
`func (o *LoggingHttpsResponse) UnsetHeaderValue()`

UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
### GetMethod

`func (o *LoggingHttpsResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoggingHttpsResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoggingHttpsResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoggingHttpsResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetJsonFormat

`func (o *LoggingHttpsResponse) GetJsonFormat() string`

GetJsonFormat returns the JsonFormat field if non-nil, zero value otherwise.

### GetJsonFormatOk

`func (o *LoggingHttpsResponse) GetJsonFormatOk() (*string, bool)`

GetJsonFormatOk returns a tuple with the JsonFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonFormat

`func (o *LoggingHttpsResponse) SetJsonFormat(v string)`

SetJsonFormat sets JsonFormat field to given value.

### HasJsonFormat

`func (o *LoggingHttpsResponse) HasJsonFormat() bool`

HasJsonFormat returns a boolean if a field has been set.

### GetPeriod

`func (o *LoggingHttpsResponse) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LoggingHttpsResponse) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LoggingHttpsResponse) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LoggingHttpsResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoggingHttpsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoggingHttpsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoggingHttpsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoggingHttpsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoggingHttpsResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoggingHttpsResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *LoggingHttpsResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *LoggingHttpsResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *LoggingHttpsResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *LoggingHttpsResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *LoggingHttpsResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *LoggingHttpsResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoggingHttpsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoggingHttpsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoggingHttpsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoggingHttpsResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoggingHttpsResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoggingHttpsResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceId

`func (o *LoggingHttpsResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *LoggingHttpsResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *LoggingHttpsResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *LoggingHttpsResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetVersion

`func (o *LoggingHttpsResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoggingHttpsResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoggingHttpsResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LoggingHttpsResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


