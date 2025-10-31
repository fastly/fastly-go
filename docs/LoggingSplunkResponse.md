# LoggingSplunkResponse

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
**RequestMaxEntries** | Pointer to **int32** | The maximum number of logs sent in one request. Defaults `0` for unbounded. | [optional] [default to 0]
**RequestMaxBytes** | Pointer to **int32** | The maximum number of bytes sent in one request. Defaults `0` for unbounded. | [optional] [default to 0]
**Url** | Pointer to **string** | The URL to post logs to. | [optional] 
**Token** | Pointer to **string** | A Splunk token for use in posting logs over HTTP to your collector. | [optional] 
**UseTls** | Pointer to [**LoggingUseTlsString**](LoggingUseTlsString.md) |  | [optional] [default to LOGGINGUSETLSSTRING_no_tls]
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceId** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLoggingSplunkResponse

`func NewLoggingSplunkResponse() *LoggingSplunkResponse`

NewLoggingSplunkResponse instantiates a new LoggingSplunkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSplunkResponseWithDefaults

`func NewLoggingSplunkResponseWithDefaults() *LoggingSplunkResponse`

NewLoggingSplunkResponseWithDefaults instantiates a new LoggingSplunkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingSplunkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingSplunkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingSplunkResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingSplunkResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingSplunkResponse) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingSplunkResponse) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingSplunkResponse) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingSplunkResponse) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingSplunkResponse) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingSplunkResponse) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetResponseCondition

`func (o *LoggingSplunkResponse) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingSplunkResponse) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingSplunkResponse) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingSplunkResponse) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingSplunkResponse) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingSplunkResponse) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingSplunkResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingSplunkResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingSplunkResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingSplunkResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetLogProcessingRegion

`func (o *LoggingSplunkResponse) GetLogProcessingRegion() string`

GetLogProcessingRegion returns the LogProcessingRegion field if non-nil, zero value otherwise.

### GetLogProcessingRegionOk

`func (o *LoggingSplunkResponse) GetLogProcessingRegionOk() (*string, bool)`

GetLogProcessingRegionOk returns a tuple with the LogProcessingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogProcessingRegion

`func (o *LoggingSplunkResponse) SetLogProcessingRegion(v string)`

SetLogProcessingRegion sets LogProcessingRegion field to given value.

### HasLogProcessingRegion

`func (o *LoggingSplunkResponse) HasLogProcessingRegion() bool`

HasLogProcessingRegion returns a boolean if a field has been set.

### GetFormatVersion

`func (o *LoggingSplunkResponse) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingSplunkResponse) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingSplunkResponse) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingSplunkResponse) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetTlsCaCert

`func (o *LoggingSplunkResponse) GetTlsCaCert() string`

GetTlsCaCert returns the TlsCaCert field if non-nil, zero value otherwise.

### GetTlsCaCertOk

`func (o *LoggingSplunkResponse) GetTlsCaCertOk() (*string, bool)`

GetTlsCaCertOk returns a tuple with the TlsCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCaCert

`func (o *LoggingSplunkResponse) SetTlsCaCert(v string)`

SetTlsCaCert sets TlsCaCert field to given value.

### HasTlsCaCert

`func (o *LoggingSplunkResponse) HasTlsCaCert() bool`

HasTlsCaCert returns a boolean if a field has been set.

### SetTlsCaCertNil

`func (o *LoggingSplunkResponse) SetTlsCaCertNil(b bool)`

 SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil

### UnsetTlsCaCert
`func (o *LoggingSplunkResponse) UnsetTlsCaCert()`

UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
### GetTlsClientCert

`func (o *LoggingSplunkResponse) GetTlsClientCert() string`

GetTlsClientCert returns the TlsClientCert field if non-nil, zero value otherwise.

### GetTlsClientCertOk

`func (o *LoggingSplunkResponse) GetTlsClientCertOk() (*string, bool)`

GetTlsClientCertOk returns a tuple with the TlsClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCert

`func (o *LoggingSplunkResponse) SetTlsClientCert(v string)`

SetTlsClientCert sets TlsClientCert field to given value.

### HasTlsClientCert

`func (o *LoggingSplunkResponse) HasTlsClientCert() bool`

HasTlsClientCert returns a boolean if a field has been set.

### SetTlsClientCertNil

`func (o *LoggingSplunkResponse) SetTlsClientCertNil(b bool)`

 SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil

### UnsetTlsClientCert
`func (o *LoggingSplunkResponse) UnsetTlsClientCert()`

UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
### GetTlsClientKey

`func (o *LoggingSplunkResponse) GetTlsClientKey() string`

GetTlsClientKey returns the TlsClientKey field if non-nil, zero value otherwise.

### GetTlsClientKeyOk

`func (o *LoggingSplunkResponse) GetTlsClientKeyOk() (*string, bool)`

GetTlsClientKeyOk returns a tuple with the TlsClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientKey

`func (o *LoggingSplunkResponse) SetTlsClientKey(v string)`

SetTlsClientKey sets TlsClientKey field to given value.

### HasTlsClientKey

`func (o *LoggingSplunkResponse) HasTlsClientKey() bool`

HasTlsClientKey returns a boolean if a field has been set.

### SetTlsClientKeyNil

`func (o *LoggingSplunkResponse) SetTlsClientKeyNil(b bool)`

 SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil

### UnsetTlsClientKey
`func (o *LoggingSplunkResponse) UnsetTlsClientKey()`

UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
### GetTlsHostname

`func (o *LoggingSplunkResponse) GetTlsHostname() string`

GetTlsHostname returns the TlsHostname field if non-nil, zero value otherwise.

### GetTlsHostnameOk

`func (o *LoggingSplunkResponse) GetTlsHostnameOk() (*string, bool)`

GetTlsHostnameOk returns a tuple with the TlsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsHostname

`func (o *LoggingSplunkResponse) SetTlsHostname(v string)`

SetTlsHostname sets TlsHostname field to given value.

### HasTlsHostname

`func (o *LoggingSplunkResponse) HasTlsHostname() bool`

HasTlsHostname returns a boolean if a field has been set.

### SetTlsHostnameNil

`func (o *LoggingSplunkResponse) SetTlsHostnameNil(b bool)`

 SetTlsHostnameNil sets the value for TlsHostname to be an explicit nil

### UnsetTlsHostname
`func (o *LoggingSplunkResponse) UnsetTlsHostname()`

UnsetTlsHostname ensures that no value is present for TlsHostname, not even an explicit nil
### GetRequestMaxEntries

`func (o *LoggingSplunkResponse) GetRequestMaxEntries() int32`

GetRequestMaxEntries returns the RequestMaxEntries field if non-nil, zero value otherwise.

### GetRequestMaxEntriesOk

`func (o *LoggingSplunkResponse) GetRequestMaxEntriesOk() (*int32, bool)`

GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxEntries

`func (o *LoggingSplunkResponse) SetRequestMaxEntries(v int32)`

SetRequestMaxEntries sets RequestMaxEntries field to given value.

### HasRequestMaxEntries

`func (o *LoggingSplunkResponse) HasRequestMaxEntries() bool`

HasRequestMaxEntries returns a boolean if a field has been set.

### GetRequestMaxBytes

`func (o *LoggingSplunkResponse) GetRequestMaxBytes() int32`

GetRequestMaxBytes returns the RequestMaxBytes field if non-nil, zero value otherwise.

### GetRequestMaxBytesOk

`func (o *LoggingSplunkResponse) GetRequestMaxBytesOk() (*int32, bool)`

GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxBytes

`func (o *LoggingSplunkResponse) SetRequestMaxBytes(v int32)`

SetRequestMaxBytes sets RequestMaxBytes field to given value.

### HasRequestMaxBytes

`func (o *LoggingSplunkResponse) HasRequestMaxBytes() bool`

HasRequestMaxBytes returns a boolean if a field has been set.

### GetUrl

`func (o *LoggingSplunkResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LoggingSplunkResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LoggingSplunkResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LoggingSplunkResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetToken

`func (o *LoggingSplunkResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingSplunkResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingSplunkResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingSplunkResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUseTls

`func (o *LoggingSplunkResponse) GetUseTls() LoggingUseTlsString`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *LoggingSplunkResponse) GetUseTlsOk() (*LoggingUseTlsString, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *LoggingSplunkResponse) SetUseTls(v LoggingUseTlsString)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *LoggingSplunkResponse) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoggingSplunkResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoggingSplunkResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoggingSplunkResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoggingSplunkResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoggingSplunkResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoggingSplunkResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *LoggingSplunkResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *LoggingSplunkResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *LoggingSplunkResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *LoggingSplunkResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *LoggingSplunkResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *LoggingSplunkResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoggingSplunkResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoggingSplunkResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoggingSplunkResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoggingSplunkResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoggingSplunkResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoggingSplunkResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceId

`func (o *LoggingSplunkResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *LoggingSplunkResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *LoggingSplunkResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *LoggingSplunkResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetVersion

`func (o *LoggingSplunkResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoggingSplunkResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoggingSplunkResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LoggingSplunkResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


