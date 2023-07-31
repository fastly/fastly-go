# LoggingHTTPSResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to **NullableString** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  | [optional] 
**ResponseCondition** | Pointer to **NullableString** | The name of an existing condition in the configured endpoint, or leave blank to always execute. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [optional] [default to "%h %l %u %t \"%r\" %&gt;s %b"]
**FormatVersion** | Pointer to **string** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to "2"]
**TLSCaCert** | Pointer to **NullableString** | A secure certificate to authenticate a server with. Must be in PEM format. | [optional] [default to "null"]
**TLSClientCert** | Pointer to **NullableString** | The client certificate used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TLSClientKey** | Pointer to **NullableString** | The client private key used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TLSHostname** | Pointer to **NullableString** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [optional] [default to "null"]
**RequestMaxEntries** | Pointer to **int32** | The maximum number of logs sent in one request. Defaults `0` (10k). | [optional] [default to 0]
**RequestMaxBytes** | Pointer to **int32** | The maximum number of bytes sent in one request. Defaults `0` (100MB). | [optional] [default to 0]
**URL** | Pointer to **string** | The URL to send logs to. Must use HTTPS. Required. | [optional] 
**ContentType** | Pointer to **NullableString** | Content type of the header sent with the request. | [optional] [default to "null"]
**HeaderName** | Pointer to **NullableString** | Name of the custom header sent with the request. | [optional] [default to "null"]
**MessageType** | Pointer to [**LoggingMessageType**](LoggingMessageType.md) |  | [optional] [default to LOGGINGMESSAGETYPE_CLASSIC]
**HeaderValue** | Pointer to **NullableString** | Value of the custom header sent with the request. | [optional] [default to "null"]
**Method** | Pointer to **string** | HTTP method used for request. | [optional] [default to "POST"]
**JSONFormat** | Pointer to **string** | Enforces valid JSON formatting for log entries. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLoggingHTTPSResponse

`func NewLoggingHTTPSResponse() *LoggingHTTPSResponse`

NewLoggingHTTPSResponse instantiates a new LoggingHTTPSResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingHTTPSResponseWithDefaults

`func NewLoggingHTTPSResponseWithDefaults() *LoggingHTTPSResponse`

NewLoggingHTTPSResponseWithDefaults instantiates a new LoggingHTTPSResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingHTTPSResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingHTTPSResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingHTTPSResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingHTTPSResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingHTTPSResponse) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingHTTPSResponse) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingHTTPSResponse) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingHTTPSResponse) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingHTTPSResponse) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingHTTPSResponse) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetResponseCondition

`func (o *LoggingHTTPSResponse) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingHTTPSResponse) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingHTTPSResponse) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingHTTPSResponse) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingHTTPSResponse) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingHTTPSResponse) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingHTTPSResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingHTTPSResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingHTTPSResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingHTTPSResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatVersion

`func (o *LoggingHTTPSResponse) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingHTTPSResponse) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingHTTPSResponse) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingHTTPSResponse) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetTLSCaCert

`func (o *LoggingHTTPSResponse) GetTLSCaCert() string`

GetTLSCaCert returns the TLSCaCert field if non-nil, zero value otherwise.

### GetTLSCaCertOk

`func (o *LoggingHTTPSResponse) GetTLSCaCertOk() (*string, bool)`

GetTLSCaCertOk returns a tuple with the TLSCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCaCert

`func (o *LoggingHTTPSResponse) SetTLSCaCert(v string)`

SetTLSCaCert sets TLSCaCert field to given value.

### HasTLSCaCert

`func (o *LoggingHTTPSResponse) HasTLSCaCert() bool`

HasTLSCaCert returns a boolean if a field has been set.

### SetTLSCaCertNil

`func (o *LoggingHTTPSResponse) SetTLSCaCertNil(b bool)`

 SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil

### UnsetTLSCaCert
`func (o *LoggingHTTPSResponse) UnsetTLSCaCert()`

UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
### GetTLSClientCert

`func (o *LoggingHTTPSResponse) GetTLSClientCert() string`

GetTLSClientCert returns the TLSClientCert field if non-nil, zero value otherwise.

### GetTLSClientCertOk

`func (o *LoggingHTTPSResponse) GetTLSClientCertOk() (*string, bool)`

GetTLSClientCertOk returns a tuple with the TLSClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientCert

`func (o *LoggingHTTPSResponse) SetTLSClientCert(v string)`

SetTLSClientCert sets TLSClientCert field to given value.

### HasTLSClientCert

`func (o *LoggingHTTPSResponse) HasTLSClientCert() bool`

HasTLSClientCert returns a boolean if a field has been set.

### SetTLSClientCertNil

`func (o *LoggingHTTPSResponse) SetTLSClientCertNil(b bool)`

 SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil

### UnsetTLSClientCert
`func (o *LoggingHTTPSResponse) UnsetTLSClientCert()`

UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
### GetTLSClientKey

`func (o *LoggingHTTPSResponse) GetTLSClientKey() string`

GetTLSClientKey returns the TLSClientKey field if non-nil, zero value otherwise.

### GetTLSClientKeyOk

`func (o *LoggingHTTPSResponse) GetTLSClientKeyOk() (*string, bool)`

GetTLSClientKeyOk returns a tuple with the TLSClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientKey

`func (o *LoggingHTTPSResponse) SetTLSClientKey(v string)`

SetTLSClientKey sets TLSClientKey field to given value.

### HasTLSClientKey

`func (o *LoggingHTTPSResponse) HasTLSClientKey() bool`

HasTLSClientKey returns a boolean if a field has been set.

### SetTLSClientKeyNil

`func (o *LoggingHTTPSResponse) SetTLSClientKeyNil(b bool)`

 SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil

### UnsetTLSClientKey
`func (o *LoggingHTTPSResponse) UnsetTLSClientKey()`

UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
### GetTLSHostname

`func (o *LoggingHTTPSResponse) GetTLSHostname() string`

GetTLSHostname returns the TLSHostname field if non-nil, zero value otherwise.

### GetTLSHostnameOk

`func (o *LoggingHTTPSResponse) GetTLSHostnameOk() (*string, bool)`

GetTLSHostnameOk returns a tuple with the TLSHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSHostname

`func (o *LoggingHTTPSResponse) SetTLSHostname(v string)`

SetTLSHostname sets TLSHostname field to given value.

### HasTLSHostname

`func (o *LoggingHTTPSResponse) HasTLSHostname() bool`

HasTLSHostname returns a boolean if a field has been set.

### SetTLSHostnameNil

`func (o *LoggingHTTPSResponse) SetTLSHostnameNil(b bool)`

 SetTLSHostnameNil sets the value for TLSHostname to be an explicit nil

### UnsetTLSHostname
`func (o *LoggingHTTPSResponse) UnsetTLSHostname()`

UnsetTLSHostname ensures that no value is present for TLSHostname, not even an explicit nil
### GetRequestMaxEntries

`func (o *LoggingHTTPSResponse) GetRequestMaxEntries() int32`

GetRequestMaxEntries returns the RequestMaxEntries field if non-nil, zero value otherwise.

### GetRequestMaxEntriesOk

`func (o *LoggingHTTPSResponse) GetRequestMaxEntriesOk() (*int32, bool)`

GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxEntries

`func (o *LoggingHTTPSResponse) SetRequestMaxEntries(v int32)`

SetRequestMaxEntries sets RequestMaxEntries field to given value.

### HasRequestMaxEntries

`func (o *LoggingHTTPSResponse) HasRequestMaxEntries() bool`

HasRequestMaxEntries returns a boolean if a field has been set.

### GetRequestMaxBytes

`func (o *LoggingHTTPSResponse) GetRequestMaxBytes() int32`

GetRequestMaxBytes returns the RequestMaxBytes field if non-nil, zero value otherwise.

### GetRequestMaxBytesOk

`func (o *LoggingHTTPSResponse) GetRequestMaxBytesOk() (*int32, bool)`

GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxBytes

`func (o *LoggingHTTPSResponse) SetRequestMaxBytes(v int32)`

SetRequestMaxBytes sets RequestMaxBytes field to given value.

### HasRequestMaxBytes

`func (o *LoggingHTTPSResponse) HasRequestMaxBytes() bool`

HasRequestMaxBytes returns a boolean if a field has been set.

### GetURL

`func (o *LoggingHTTPSResponse) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingHTTPSResponse) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingHTTPSResponse) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingHTTPSResponse) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetContentType

`func (o *LoggingHTTPSResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LoggingHTTPSResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LoggingHTTPSResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *LoggingHTTPSResponse) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *LoggingHTTPSResponse) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *LoggingHTTPSResponse) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetHeaderName

`func (o *LoggingHTTPSResponse) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *LoggingHTTPSResponse) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *LoggingHTTPSResponse) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *LoggingHTTPSResponse) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### SetHeaderNameNil

`func (o *LoggingHTTPSResponse) SetHeaderNameNil(b bool)`

 SetHeaderNameNil sets the value for HeaderName to be an explicit nil

### UnsetHeaderName
`func (o *LoggingHTTPSResponse) UnsetHeaderName()`

UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil
### GetMessageType

`func (o *LoggingHTTPSResponse) GetMessageType() LoggingMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingHTTPSResponse) GetMessageTypeOk() (*LoggingMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingHTTPSResponse) SetMessageType(v LoggingMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingHTTPSResponse) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetHeaderValue

`func (o *LoggingHTTPSResponse) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *LoggingHTTPSResponse) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *LoggingHTTPSResponse) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.

### HasHeaderValue

`func (o *LoggingHTTPSResponse) HasHeaderValue() bool`

HasHeaderValue returns a boolean if a field has been set.

### SetHeaderValueNil

`func (o *LoggingHTTPSResponse) SetHeaderValueNil(b bool)`

 SetHeaderValueNil sets the value for HeaderValue to be an explicit nil

### UnsetHeaderValue
`func (o *LoggingHTTPSResponse) UnsetHeaderValue()`

UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
### GetMethod

`func (o *LoggingHTTPSResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoggingHTTPSResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoggingHTTPSResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoggingHTTPSResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetJSONFormat

`func (o *LoggingHTTPSResponse) GetJSONFormat() string`

GetJSONFormat returns the JSONFormat field if non-nil, zero value otherwise.

### GetJSONFormatOk

`func (o *LoggingHTTPSResponse) GetJSONFormatOk() (*string, bool)`

GetJSONFormatOk returns a tuple with the JSONFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJSONFormat

`func (o *LoggingHTTPSResponse) SetJSONFormat(v string)`

SetJSONFormat sets JSONFormat field to given value.

### HasJSONFormat

`func (o *LoggingHTTPSResponse) HasJSONFormat() bool`

HasJSONFormat returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoggingHTTPSResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoggingHTTPSResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoggingHTTPSResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoggingHTTPSResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoggingHTTPSResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoggingHTTPSResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *LoggingHTTPSResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *LoggingHTTPSResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *LoggingHTTPSResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *LoggingHTTPSResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *LoggingHTTPSResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *LoggingHTTPSResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoggingHTTPSResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoggingHTTPSResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoggingHTTPSResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoggingHTTPSResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoggingHTTPSResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoggingHTTPSResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *LoggingHTTPSResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *LoggingHTTPSResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *LoggingHTTPSResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *LoggingHTTPSResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *LoggingHTTPSResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoggingHTTPSResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoggingHTTPSResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LoggingHTTPSResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
