# LoggingElasticsearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to **NullableString** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  | [optional] 
**ResponseCondition** | Pointer to **NullableString** | The name of an existing condition in the configured endpoint, or leave blank to always execute. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that Elasticsearch can ingest. | [optional] 
**LogProcessingRegion** | Pointer to **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [optional] [default to "none"]
**FormatVersion** | Pointer to **string** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to "2"]
**TlsCaCert** | Pointer to **NullableString** | A secure certificate to authenticate a server with. Must be in PEM format. | [optional] [default to "null"]
**TlsClientCert** | Pointer to **NullableString** | The client certificate used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TlsClientKey** | Pointer to **NullableString** | The client private key used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TlsHostname** | Pointer to **NullableString** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [optional] [default to "null"]
**RequestMaxEntries** | Pointer to **int32** | The maximum number of logs sent in one request. Defaults `0` for unbounded. | [optional] [default to 0]
**RequestMaxBytes** | Pointer to **int32** | The maximum number of bytes sent in one request. Defaults `0` for unbounded. | [optional] [default to 0]
**Index** | Pointer to **string** | The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, `#{%F}` will interpolate as `YYYY-MM-DD` with today&#39;s date. | [optional] 
**Url** | Pointer to **string** | The URL to stream logs to. Must use HTTPS. | [optional] 
**Pipeline** | Pointer to **NullableString** | The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html). | [optional] 
**User** | Pointer to **NullableString** | Basic Auth username. | [optional] 
**Password** | Pointer to **NullableString** | Basic Auth password. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceId** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLoggingElasticsearchResponse

`func NewLoggingElasticsearchResponse() *LoggingElasticsearchResponse`

NewLoggingElasticsearchResponse instantiates a new LoggingElasticsearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingElasticsearchResponseWithDefaults

`func NewLoggingElasticsearchResponseWithDefaults() *LoggingElasticsearchResponse`

NewLoggingElasticsearchResponseWithDefaults instantiates a new LoggingElasticsearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingElasticsearchResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingElasticsearchResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingElasticsearchResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingElasticsearchResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingElasticsearchResponse) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingElasticsearchResponse) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingElasticsearchResponse) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingElasticsearchResponse) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingElasticsearchResponse) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingElasticsearchResponse) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetResponseCondition

`func (o *LoggingElasticsearchResponse) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingElasticsearchResponse) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingElasticsearchResponse) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingElasticsearchResponse) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingElasticsearchResponse) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingElasticsearchResponse) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingElasticsearchResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingElasticsearchResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingElasticsearchResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingElasticsearchResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetLogProcessingRegion

`func (o *LoggingElasticsearchResponse) GetLogProcessingRegion() string`

GetLogProcessingRegion returns the LogProcessingRegion field if non-nil, zero value otherwise.

### GetLogProcessingRegionOk

`func (o *LoggingElasticsearchResponse) GetLogProcessingRegionOk() (*string, bool)`

GetLogProcessingRegionOk returns a tuple with the LogProcessingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogProcessingRegion

`func (o *LoggingElasticsearchResponse) SetLogProcessingRegion(v string)`

SetLogProcessingRegion sets LogProcessingRegion field to given value.

### HasLogProcessingRegion

`func (o *LoggingElasticsearchResponse) HasLogProcessingRegion() bool`

HasLogProcessingRegion returns a boolean if a field has been set.

### GetFormatVersion

`func (o *LoggingElasticsearchResponse) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingElasticsearchResponse) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingElasticsearchResponse) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingElasticsearchResponse) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetTlsCaCert

`func (o *LoggingElasticsearchResponse) GetTlsCaCert() string`

GetTlsCaCert returns the TlsCaCert field if non-nil, zero value otherwise.

### GetTlsCaCertOk

`func (o *LoggingElasticsearchResponse) GetTlsCaCertOk() (*string, bool)`

GetTlsCaCertOk returns a tuple with the TlsCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCaCert

`func (o *LoggingElasticsearchResponse) SetTlsCaCert(v string)`

SetTlsCaCert sets TlsCaCert field to given value.

### HasTlsCaCert

`func (o *LoggingElasticsearchResponse) HasTlsCaCert() bool`

HasTlsCaCert returns a boolean if a field has been set.

### SetTlsCaCertNil

`func (o *LoggingElasticsearchResponse) SetTlsCaCertNil(b bool)`

 SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil

### UnsetTlsCaCert
`func (o *LoggingElasticsearchResponse) UnsetTlsCaCert()`

UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
### GetTlsClientCert

`func (o *LoggingElasticsearchResponse) GetTlsClientCert() string`

GetTlsClientCert returns the TlsClientCert field if non-nil, zero value otherwise.

### GetTlsClientCertOk

`func (o *LoggingElasticsearchResponse) GetTlsClientCertOk() (*string, bool)`

GetTlsClientCertOk returns a tuple with the TlsClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCert

`func (o *LoggingElasticsearchResponse) SetTlsClientCert(v string)`

SetTlsClientCert sets TlsClientCert field to given value.

### HasTlsClientCert

`func (o *LoggingElasticsearchResponse) HasTlsClientCert() bool`

HasTlsClientCert returns a boolean if a field has been set.

### SetTlsClientCertNil

`func (o *LoggingElasticsearchResponse) SetTlsClientCertNil(b bool)`

 SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil

### UnsetTlsClientCert
`func (o *LoggingElasticsearchResponse) UnsetTlsClientCert()`

UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
### GetTlsClientKey

`func (o *LoggingElasticsearchResponse) GetTlsClientKey() string`

GetTlsClientKey returns the TlsClientKey field if non-nil, zero value otherwise.

### GetTlsClientKeyOk

`func (o *LoggingElasticsearchResponse) GetTlsClientKeyOk() (*string, bool)`

GetTlsClientKeyOk returns a tuple with the TlsClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientKey

`func (o *LoggingElasticsearchResponse) SetTlsClientKey(v string)`

SetTlsClientKey sets TlsClientKey field to given value.

### HasTlsClientKey

`func (o *LoggingElasticsearchResponse) HasTlsClientKey() bool`

HasTlsClientKey returns a boolean if a field has been set.

### SetTlsClientKeyNil

`func (o *LoggingElasticsearchResponse) SetTlsClientKeyNil(b bool)`

 SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil

### UnsetTlsClientKey
`func (o *LoggingElasticsearchResponse) UnsetTlsClientKey()`

UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
### GetTlsHostname

`func (o *LoggingElasticsearchResponse) GetTlsHostname() string`

GetTlsHostname returns the TlsHostname field if non-nil, zero value otherwise.

### GetTlsHostnameOk

`func (o *LoggingElasticsearchResponse) GetTlsHostnameOk() (*string, bool)`

GetTlsHostnameOk returns a tuple with the TlsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsHostname

`func (o *LoggingElasticsearchResponse) SetTlsHostname(v string)`

SetTlsHostname sets TlsHostname field to given value.

### HasTlsHostname

`func (o *LoggingElasticsearchResponse) HasTlsHostname() bool`

HasTlsHostname returns a boolean if a field has been set.

### SetTlsHostnameNil

`func (o *LoggingElasticsearchResponse) SetTlsHostnameNil(b bool)`

 SetTlsHostnameNil sets the value for TlsHostname to be an explicit nil

### UnsetTlsHostname
`func (o *LoggingElasticsearchResponse) UnsetTlsHostname()`

UnsetTlsHostname ensures that no value is present for TlsHostname, not even an explicit nil
### GetRequestMaxEntries

`func (o *LoggingElasticsearchResponse) GetRequestMaxEntries() int32`

GetRequestMaxEntries returns the RequestMaxEntries field if non-nil, zero value otherwise.

### GetRequestMaxEntriesOk

`func (o *LoggingElasticsearchResponse) GetRequestMaxEntriesOk() (*int32, bool)`

GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxEntries

`func (o *LoggingElasticsearchResponse) SetRequestMaxEntries(v int32)`

SetRequestMaxEntries sets RequestMaxEntries field to given value.

### HasRequestMaxEntries

`func (o *LoggingElasticsearchResponse) HasRequestMaxEntries() bool`

HasRequestMaxEntries returns a boolean if a field has been set.

### GetRequestMaxBytes

`func (o *LoggingElasticsearchResponse) GetRequestMaxBytes() int32`

GetRequestMaxBytes returns the RequestMaxBytes field if non-nil, zero value otherwise.

### GetRequestMaxBytesOk

`func (o *LoggingElasticsearchResponse) GetRequestMaxBytesOk() (*int32, bool)`

GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxBytes

`func (o *LoggingElasticsearchResponse) SetRequestMaxBytes(v int32)`

SetRequestMaxBytes sets RequestMaxBytes field to given value.

### HasRequestMaxBytes

`func (o *LoggingElasticsearchResponse) HasRequestMaxBytes() bool`

HasRequestMaxBytes returns a boolean if a field has been set.

### GetIndex

`func (o *LoggingElasticsearchResponse) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *LoggingElasticsearchResponse) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *LoggingElasticsearchResponse) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *LoggingElasticsearchResponse) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetUrl

`func (o *LoggingElasticsearchResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LoggingElasticsearchResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LoggingElasticsearchResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LoggingElasticsearchResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPipeline

`func (o *LoggingElasticsearchResponse) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *LoggingElasticsearchResponse) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *LoggingElasticsearchResponse) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *LoggingElasticsearchResponse) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### SetPipelineNil

`func (o *LoggingElasticsearchResponse) SetPipelineNil(b bool)`

 SetPipelineNil sets the value for Pipeline to be an explicit nil

### UnsetPipeline
`func (o *LoggingElasticsearchResponse) UnsetPipeline()`

UnsetPipeline ensures that no value is present for Pipeline, not even an explicit nil
### GetUser

`func (o *LoggingElasticsearchResponse) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingElasticsearchResponse) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingElasticsearchResponse) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingElasticsearchResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *LoggingElasticsearchResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *LoggingElasticsearchResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *LoggingElasticsearchResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoggingElasticsearchResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoggingElasticsearchResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoggingElasticsearchResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *LoggingElasticsearchResponse) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *LoggingElasticsearchResponse) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetCreatedAt

`func (o *LoggingElasticsearchResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoggingElasticsearchResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoggingElasticsearchResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoggingElasticsearchResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoggingElasticsearchResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoggingElasticsearchResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *LoggingElasticsearchResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *LoggingElasticsearchResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *LoggingElasticsearchResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *LoggingElasticsearchResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *LoggingElasticsearchResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *LoggingElasticsearchResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoggingElasticsearchResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoggingElasticsearchResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoggingElasticsearchResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoggingElasticsearchResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoggingElasticsearchResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoggingElasticsearchResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceId

`func (o *LoggingElasticsearchResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *LoggingElasticsearchResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *LoggingElasticsearchResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *LoggingElasticsearchResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetVersion

`func (o *LoggingElasticsearchResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoggingElasticsearchResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoggingElasticsearchResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LoggingElasticsearchResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


