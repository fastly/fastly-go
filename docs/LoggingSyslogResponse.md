# LoggingSyslogResponse

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
**Address** | Pointer to **string** | A hostname or IPv4 address. | [optional] 
**Port** | Pointer to **int32** | The port number. | [optional] [default to 514]
**MessageType** | Pointer to [**LoggingMessageType**](LoggingMessageType.md) |  | [optional] [default to LOGGINGMESSAGETYPE_CLASSIC]
**Hostname** | Pointer to **string** | The hostname used for the syslog endpoint. | [optional] 
**Ipv4** | Pointer to **NullableString** | The IPv4 address used for the syslog endpoint. | [optional] 
**Token** | Pointer to **NullableString** | Whether to prepend each message with a specific token. | [optional] [default to "null"]
**UseTLS** | Pointer to [**LoggingUseTLS**](LoggingUseTLS.md) |  | [optional] [default to LOGGINGUSETLS_no_tls]
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLoggingSyslogResponse

`func NewLoggingSyslogResponse() *LoggingSyslogResponse`

NewLoggingSyslogResponse instantiates a new LoggingSyslogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSyslogResponseWithDefaults

`func NewLoggingSyslogResponseWithDefaults() *LoggingSyslogResponse`

NewLoggingSyslogResponseWithDefaults instantiates a new LoggingSyslogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingSyslogResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingSyslogResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingSyslogResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingSyslogResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingSyslogResponse) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingSyslogResponse) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingSyslogResponse) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingSyslogResponse) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingSyslogResponse) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingSyslogResponse) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetResponseCondition

`func (o *LoggingSyslogResponse) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingSyslogResponse) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingSyslogResponse) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingSyslogResponse) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingSyslogResponse) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingSyslogResponse) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingSyslogResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingSyslogResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingSyslogResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingSyslogResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatVersion

`func (o *LoggingSyslogResponse) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingSyslogResponse) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingSyslogResponse) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingSyslogResponse) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetTLSCaCert

`func (o *LoggingSyslogResponse) GetTLSCaCert() string`

GetTLSCaCert returns the TLSCaCert field if non-nil, zero value otherwise.

### GetTLSCaCertOk

`func (o *LoggingSyslogResponse) GetTLSCaCertOk() (*string, bool)`

GetTLSCaCertOk returns a tuple with the TLSCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCaCert

`func (o *LoggingSyslogResponse) SetTLSCaCert(v string)`

SetTLSCaCert sets TLSCaCert field to given value.

### HasTLSCaCert

`func (o *LoggingSyslogResponse) HasTLSCaCert() bool`

HasTLSCaCert returns a boolean if a field has been set.

### SetTLSCaCertNil

`func (o *LoggingSyslogResponse) SetTLSCaCertNil(b bool)`

 SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil

### UnsetTLSCaCert
`func (o *LoggingSyslogResponse) UnsetTLSCaCert()`

UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
### GetTLSClientCert

`func (o *LoggingSyslogResponse) GetTLSClientCert() string`

GetTLSClientCert returns the TLSClientCert field if non-nil, zero value otherwise.

### GetTLSClientCertOk

`func (o *LoggingSyslogResponse) GetTLSClientCertOk() (*string, bool)`

GetTLSClientCertOk returns a tuple with the TLSClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientCert

`func (o *LoggingSyslogResponse) SetTLSClientCert(v string)`

SetTLSClientCert sets TLSClientCert field to given value.

### HasTLSClientCert

`func (o *LoggingSyslogResponse) HasTLSClientCert() bool`

HasTLSClientCert returns a boolean if a field has been set.

### SetTLSClientCertNil

`func (o *LoggingSyslogResponse) SetTLSClientCertNil(b bool)`

 SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil

### UnsetTLSClientCert
`func (o *LoggingSyslogResponse) UnsetTLSClientCert()`

UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
### GetTLSClientKey

`func (o *LoggingSyslogResponse) GetTLSClientKey() string`

GetTLSClientKey returns the TLSClientKey field if non-nil, zero value otherwise.

### GetTLSClientKeyOk

`func (o *LoggingSyslogResponse) GetTLSClientKeyOk() (*string, bool)`

GetTLSClientKeyOk returns a tuple with the TLSClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientKey

`func (o *LoggingSyslogResponse) SetTLSClientKey(v string)`

SetTLSClientKey sets TLSClientKey field to given value.

### HasTLSClientKey

`func (o *LoggingSyslogResponse) HasTLSClientKey() bool`

HasTLSClientKey returns a boolean if a field has been set.

### SetTLSClientKeyNil

`func (o *LoggingSyslogResponse) SetTLSClientKeyNil(b bool)`

 SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil

### UnsetTLSClientKey
`func (o *LoggingSyslogResponse) UnsetTLSClientKey()`

UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
### GetTLSHostname

`func (o *LoggingSyslogResponse) GetTLSHostname() string`

GetTLSHostname returns the TLSHostname field if non-nil, zero value otherwise.

### GetTLSHostnameOk

`func (o *LoggingSyslogResponse) GetTLSHostnameOk() (*string, bool)`

GetTLSHostnameOk returns a tuple with the TLSHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSHostname

`func (o *LoggingSyslogResponse) SetTLSHostname(v string)`

SetTLSHostname sets TLSHostname field to given value.

### HasTLSHostname

`func (o *LoggingSyslogResponse) HasTLSHostname() bool`

HasTLSHostname returns a boolean if a field has been set.

### SetTLSHostnameNil

`func (o *LoggingSyslogResponse) SetTLSHostnameNil(b bool)`

 SetTLSHostnameNil sets the value for TLSHostname to be an explicit nil

### UnsetTLSHostname
`func (o *LoggingSyslogResponse) UnsetTLSHostname()`

UnsetTLSHostname ensures that no value is present for TLSHostname, not even an explicit nil
### GetAddress

`func (o *LoggingSyslogResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LoggingSyslogResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LoggingSyslogResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LoggingSyslogResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *LoggingSyslogResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoggingSyslogResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoggingSyslogResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoggingSyslogResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetMessageType

`func (o *LoggingSyslogResponse) GetMessageType() LoggingMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LoggingSyslogResponse) GetMessageTypeOk() (*LoggingMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LoggingSyslogResponse) SetMessageType(v LoggingMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LoggingSyslogResponse) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetHostname

`func (o *LoggingSyslogResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LoggingSyslogResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LoggingSyslogResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LoggingSyslogResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpv4

`func (o *LoggingSyslogResponse) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *LoggingSyslogResponse) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *LoggingSyslogResponse) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *LoggingSyslogResponse) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### SetIpv4Nil

`func (o *LoggingSyslogResponse) SetIpv4Nil(b bool)`

 SetIpv4Nil sets the value for Ipv4 to be an explicit nil

### UnsetIpv4
`func (o *LoggingSyslogResponse) UnsetIpv4()`

UnsetIpv4 ensures that no value is present for Ipv4, not even an explicit nil
### GetToken

`func (o *LoggingSyslogResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingSyslogResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingSyslogResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingSyslogResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *LoggingSyslogResponse) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *LoggingSyslogResponse) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetUseTLS

`func (o *LoggingSyslogResponse) GetUseTLS() LoggingUseTLS`

GetUseTLS returns the UseTLS field if non-nil, zero value otherwise.

### GetUseTLSOk

`func (o *LoggingSyslogResponse) GetUseTLSOk() (*LoggingUseTLS, bool)`

GetUseTLSOk returns a tuple with the UseTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTLS

`func (o *LoggingSyslogResponse) SetUseTLS(v LoggingUseTLS)`

SetUseTLS sets UseTLS field to given value.

### HasUseTLS

`func (o *LoggingSyslogResponse) HasUseTLS() bool`

HasUseTLS returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoggingSyslogResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoggingSyslogResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoggingSyslogResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoggingSyslogResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoggingSyslogResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoggingSyslogResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *LoggingSyslogResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *LoggingSyslogResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *LoggingSyslogResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *LoggingSyslogResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *LoggingSyslogResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *LoggingSyslogResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoggingSyslogResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoggingSyslogResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoggingSyslogResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoggingSyslogResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoggingSyslogResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoggingSyslogResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *LoggingSyslogResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *LoggingSyslogResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *LoggingSyslogResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *LoggingSyslogResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *LoggingSyslogResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoggingSyslogResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoggingSyslogResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LoggingSyslogResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
