# LoggingKafkaResponsePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to **NullableString** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  | [optional] 
**ResponseCondition** | Pointer to **NullableString** | The name of an existing condition in the configured endpoint, or leave blank to always execute. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [optional] [default to "%h %l %u %t \"%r\" %&gt;s %b"]
**FormatVersion** | Pointer to **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to 2]
**TLSCaCert** | Pointer to **NullableString** | A secure certificate to authenticate a server with. Must be in PEM format. | [optional] [default to "null"]
**TLSClientCert** | Pointer to **NullableString** | The client certificate used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TLSClientKey** | Pointer to **NullableString** | The client private key used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TLSHostname** | Pointer to **NullableString** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [optional] [default to "null"]
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**Topic** | Pointer to **string** | The Kafka topic to send logs to. Required. | [optional] 
**Brokers** | Pointer to **string** | A comma-separated list of IP addresses or hostnames of Kafka brokers. Required. | [optional] 
**CompressionCodec** | Pointer to **NullableString** | The codec used for compression of your logs. | [optional] 
**RequiredAcks** | Pointer to **int32** | The number of acknowledgements a leader must receive before a write is considered successful. | [optional] [default to 1]
**RequestMaxBytes** | Pointer to **int32** | The maximum number of bytes sent in one request. Defaults `0` (no limit). | [optional] [default to 0]
**ParseLogKeyvals** | Pointer to **bool** | Enables parsing of key&#x3D;value tuples from the beginning of a logline, turning them into [record headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers). | [optional] 
**AuthMethod** | Pointer to **string** | SASL authentication method. | [optional] 
**User** | Pointer to **string** | SASL user. | [optional] 
**Password** | Pointer to **string** | SASL password. | [optional] 
**UseTLS** | Pointer to [**LoggingUseTLSString**](LoggingUseTLSString.md) |  | [optional] [default to LOGGINGUSETLSSTRING_no_tls]

## Methods

### NewLoggingKafkaResponsePost

`func NewLoggingKafkaResponsePost() *LoggingKafkaResponsePost`

NewLoggingKafkaResponsePost instantiates a new LoggingKafkaResponsePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingKafkaResponsePostWithDefaults

`func NewLoggingKafkaResponsePostWithDefaults() *LoggingKafkaResponsePost`

NewLoggingKafkaResponsePostWithDefaults instantiates a new LoggingKafkaResponsePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingKafkaResponsePost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingKafkaResponsePost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingKafkaResponsePost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingKafkaResponsePost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingKafkaResponsePost) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingKafkaResponsePost) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingKafkaResponsePost) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingKafkaResponsePost) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingKafkaResponsePost) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingKafkaResponsePost) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetResponseCondition

`func (o *LoggingKafkaResponsePost) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingKafkaResponsePost) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingKafkaResponsePost) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingKafkaResponsePost) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingKafkaResponsePost) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingKafkaResponsePost) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingKafkaResponsePost) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingKafkaResponsePost) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingKafkaResponsePost) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingKafkaResponsePost) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatVersion

`func (o *LoggingKafkaResponsePost) GetFormatVersion() int32`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingKafkaResponsePost) GetFormatVersionOk() (*int32, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingKafkaResponsePost) SetFormatVersion(v int32)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingKafkaResponsePost) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetTLSCaCert

`func (o *LoggingKafkaResponsePost) GetTLSCaCert() string`

GetTLSCaCert returns the TLSCaCert field if non-nil, zero value otherwise.

### GetTLSCaCertOk

`func (o *LoggingKafkaResponsePost) GetTLSCaCertOk() (*string, bool)`

GetTLSCaCertOk returns a tuple with the TLSCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCaCert

`func (o *LoggingKafkaResponsePost) SetTLSCaCert(v string)`

SetTLSCaCert sets TLSCaCert field to given value.

### HasTLSCaCert

`func (o *LoggingKafkaResponsePost) HasTLSCaCert() bool`

HasTLSCaCert returns a boolean if a field has been set.

### SetTLSCaCertNil

`func (o *LoggingKafkaResponsePost) SetTLSCaCertNil(b bool)`

 SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil

### UnsetTLSCaCert
`func (o *LoggingKafkaResponsePost) UnsetTLSCaCert()`

UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
### GetTLSClientCert

`func (o *LoggingKafkaResponsePost) GetTLSClientCert() string`

GetTLSClientCert returns the TLSClientCert field if non-nil, zero value otherwise.

### GetTLSClientCertOk

`func (o *LoggingKafkaResponsePost) GetTLSClientCertOk() (*string, bool)`

GetTLSClientCertOk returns a tuple with the TLSClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientCert

`func (o *LoggingKafkaResponsePost) SetTLSClientCert(v string)`

SetTLSClientCert sets TLSClientCert field to given value.

### HasTLSClientCert

`func (o *LoggingKafkaResponsePost) HasTLSClientCert() bool`

HasTLSClientCert returns a boolean if a field has been set.

### SetTLSClientCertNil

`func (o *LoggingKafkaResponsePost) SetTLSClientCertNil(b bool)`

 SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil

### UnsetTLSClientCert
`func (o *LoggingKafkaResponsePost) UnsetTLSClientCert()`

UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
### GetTLSClientKey

`func (o *LoggingKafkaResponsePost) GetTLSClientKey() string`

GetTLSClientKey returns the TLSClientKey field if non-nil, zero value otherwise.

### GetTLSClientKeyOk

`func (o *LoggingKafkaResponsePost) GetTLSClientKeyOk() (*string, bool)`

GetTLSClientKeyOk returns a tuple with the TLSClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientKey

`func (o *LoggingKafkaResponsePost) SetTLSClientKey(v string)`

SetTLSClientKey sets TLSClientKey field to given value.

### HasTLSClientKey

`func (o *LoggingKafkaResponsePost) HasTLSClientKey() bool`

HasTLSClientKey returns a boolean if a field has been set.

### SetTLSClientKeyNil

`func (o *LoggingKafkaResponsePost) SetTLSClientKeyNil(b bool)`

 SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil

### UnsetTLSClientKey
`func (o *LoggingKafkaResponsePost) UnsetTLSClientKey()`

UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
### GetTLSHostname

`func (o *LoggingKafkaResponsePost) GetTLSHostname() string`

GetTLSHostname returns the TLSHostname field if non-nil, zero value otherwise.

### GetTLSHostnameOk

`func (o *LoggingKafkaResponsePost) GetTLSHostnameOk() (*string, bool)`

GetTLSHostnameOk returns a tuple with the TLSHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSHostname

`func (o *LoggingKafkaResponsePost) SetTLSHostname(v string)`

SetTLSHostname sets TLSHostname field to given value.

### HasTLSHostname

`func (o *LoggingKafkaResponsePost) HasTLSHostname() bool`

HasTLSHostname returns a boolean if a field has been set.

### SetTLSHostnameNil

`func (o *LoggingKafkaResponsePost) SetTLSHostnameNil(b bool)`

 SetTLSHostnameNil sets the value for TLSHostname to be an explicit nil

### UnsetTLSHostname
`func (o *LoggingKafkaResponsePost) UnsetTLSHostname()`

UnsetTLSHostname ensures that no value is present for TLSHostname, not even an explicit nil
### GetCreatedAt

`func (o *LoggingKafkaResponsePost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoggingKafkaResponsePost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoggingKafkaResponsePost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoggingKafkaResponsePost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoggingKafkaResponsePost) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoggingKafkaResponsePost) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *LoggingKafkaResponsePost) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *LoggingKafkaResponsePost) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *LoggingKafkaResponsePost) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *LoggingKafkaResponsePost) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *LoggingKafkaResponsePost) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *LoggingKafkaResponsePost) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoggingKafkaResponsePost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoggingKafkaResponsePost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoggingKafkaResponsePost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoggingKafkaResponsePost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoggingKafkaResponsePost) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoggingKafkaResponsePost) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *LoggingKafkaResponsePost) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *LoggingKafkaResponsePost) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *LoggingKafkaResponsePost) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *LoggingKafkaResponsePost) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *LoggingKafkaResponsePost) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoggingKafkaResponsePost) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoggingKafkaResponsePost) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LoggingKafkaResponsePost) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTopic

`func (o *LoggingKafkaResponsePost) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *LoggingKafkaResponsePost) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *LoggingKafkaResponsePost) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *LoggingKafkaResponsePost) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetBrokers

`func (o *LoggingKafkaResponsePost) GetBrokers() string`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *LoggingKafkaResponsePost) GetBrokersOk() (*string, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *LoggingKafkaResponsePost) SetBrokers(v string)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *LoggingKafkaResponsePost) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.

### GetCompressionCodec

`func (o *LoggingKafkaResponsePost) GetCompressionCodec() string`

GetCompressionCodec returns the CompressionCodec field if non-nil, zero value otherwise.

### GetCompressionCodecOk

`func (o *LoggingKafkaResponsePost) GetCompressionCodecOk() (*string, bool)`

GetCompressionCodecOk returns a tuple with the CompressionCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionCodec

`func (o *LoggingKafkaResponsePost) SetCompressionCodec(v string)`

SetCompressionCodec sets CompressionCodec field to given value.

### HasCompressionCodec

`func (o *LoggingKafkaResponsePost) HasCompressionCodec() bool`

HasCompressionCodec returns a boolean if a field has been set.

### SetCompressionCodecNil

`func (o *LoggingKafkaResponsePost) SetCompressionCodecNil(b bool)`

 SetCompressionCodecNil sets the value for CompressionCodec to be an explicit nil

### UnsetCompressionCodec
`func (o *LoggingKafkaResponsePost) UnsetCompressionCodec()`

UnsetCompressionCodec ensures that no value is present for CompressionCodec, not even an explicit nil
### GetRequiredAcks

`func (o *LoggingKafkaResponsePost) GetRequiredAcks() int32`

GetRequiredAcks returns the RequiredAcks field if non-nil, zero value otherwise.

### GetRequiredAcksOk

`func (o *LoggingKafkaResponsePost) GetRequiredAcksOk() (*int32, bool)`

GetRequiredAcksOk returns a tuple with the RequiredAcks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAcks

`func (o *LoggingKafkaResponsePost) SetRequiredAcks(v int32)`

SetRequiredAcks sets RequiredAcks field to given value.

### HasRequiredAcks

`func (o *LoggingKafkaResponsePost) HasRequiredAcks() bool`

HasRequiredAcks returns a boolean if a field has been set.

### GetRequestMaxBytes

`func (o *LoggingKafkaResponsePost) GetRequestMaxBytes() int32`

GetRequestMaxBytes returns the RequestMaxBytes field if non-nil, zero value otherwise.

### GetRequestMaxBytesOk

`func (o *LoggingKafkaResponsePost) GetRequestMaxBytesOk() (*int32, bool)`

GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxBytes

`func (o *LoggingKafkaResponsePost) SetRequestMaxBytes(v int32)`

SetRequestMaxBytes sets RequestMaxBytes field to given value.

### HasRequestMaxBytes

`func (o *LoggingKafkaResponsePost) HasRequestMaxBytes() bool`

HasRequestMaxBytes returns a boolean if a field has been set.

### GetParseLogKeyvals

`func (o *LoggingKafkaResponsePost) GetParseLogKeyvals() bool`

GetParseLogKeyvals returns the ParseLogKeyvals field if non-nil, zero value otherwise.

### GetParseLogKeyvalsOk

`func (o *LoggingKafkaResponsePost) GetParseLogKeyvalsOk() (*bool, bool)`

GetParseLogKeyvalsOk returns a tuple with the ParseLogKeyvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseLogKeyvals

`func (o *LoggingKafkaResponsePost) SetParseLogKeyvals(v bool)`

SetParseLogKeyvals sets ParseLogKeyvals field to given value.

### HasParseLogKeyvals

`func (o *LoggingKafkaResponsePost) HasParseLogKeyvals() bool`

HasParseLogKeyvals returns a boolean if a field has been set.

### GetAuthMethod

`func (o *LoggingKafkaResponsePost) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *LoggingKafkaResponsePost) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *LoggingKafkaResponsePost) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *LoggingKafkaResponsePost) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetUser

`func (o *LoggingKafkaResponsePost) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingKafkaResponsePost) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingKafkaResponsePost) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingKafkaResponsePost) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *LoggingKafkaResponsePost) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoggingKafkaResponsePost) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoggingKafkaResponsePost) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoggingKafkaResponsePost) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseTLS

`func (o *LoggingKafkaResponsePost) GetUseTLS() LoggingUseTLSString`

GetUseTLS returns the UseTLS field if non-nil, zero value otherwise.

### GetUseTLSOk

`func (o *LoggingKafkaResponsePost) GetUseTLSOk() (*LoggingUseTLSString, bool)`

GetUseTLSOk returns a tuple with the UseTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTLS

`func (o *LoggingKafkaResponsePost) SetUseTLS(v LoggingUseTLSString)`

SetUseTLS sets UseTLS field to given value.

### HasUseTLS

`func (o *LoggingKafkaResponsePost) HasUseTLS() bool`

HasUseTLS returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
