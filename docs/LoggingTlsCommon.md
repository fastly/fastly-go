# LoggingTlsCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsCaCert** | Pointer to **NullableString** | A secure certificate to authenticate a server with. Must be in PEM format. | [optional] [default to "null"]
**TlsClientCert** | Pointer to **NullableString** | The client certificate used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TlsClientKey** | Pointer to **NullableString** | The client private key used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TlsHostname** | Pointer to **NullableString** | The hostname to verify the server&#39;s certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported. | [optional] [default to "null"]

## Methods

### NewLoggingTlsCommon

`func NewLoggingTlsCommon() *LoggingTlsCommon`

NewLoggingTlsCommon instantiates a new LoggingTlsCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingTlsCommonWithDefaults

`func NewLoggingTlsCommonWithDefaults() *LoggingTlsCommon`

NewLoggingTlsCommonWithDefaults instantiates a new LoggingTlsCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsCaCert

`func (o *LoggingTlsCommon) GetTlsCaCert() string`

GetTlsCaCert returns the TlsCaCert field if non-nil, zero value otherwise.

### GetTlsCaCertOk

`func (o *LoggingTlsCommon) GetTlsCaCertOk() (*string, bool)`

GetTlsCaCertOk returns a tuple with the TlsCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCaCert

`func (o *LoggingTlsCommon) SetTlsCaCert(v string)`

SetTlsCaCert sets TlsCaCert field to given value.

### HasTlsCaCert

`func (o *LoggingTlsCommon) HasTlsCaCert() bool`

HasTlsCaCert returns a boolean if a field has been set.

### SetTlsCaCertNil

`func (o *LoggingTlsCommon) SetTlsCaCertNil(b bool)`

 SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil

### UnsetTlsCaCert
`func (o *LoggingTlsCommon) UnsetTlsCaCert()`

UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
### GetTlsClientCert

`func (o *LoggingTlsCommon) GetTlsClientCert() string`

GetTlsClientCert returns the TlsClientCert field if non-nil, zero value otherwise.

### GetTlsClientCertOk

`func (o *LoggingTlsCommon) GetTlsClientCertOk() (*string, bool)`

GetTlsClientCertOk returns a tuple with the TlsClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCert

`func (o *LoggingTlsCommon) SetTlsClientCert(v string)`

SetTlsClientCert sets TlsClientCert field to given value.

### HasTlsClientCert

`func (o *LoggingTlsCommon) HasTlsClientCert() bool`

HasTlsClientCert returns a boolean if a field has been set.

### SetTlsClientCertNil

`func (o *LoggingTlsCommon) SetTlsClientCertNil(b bool)`

 SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil

### UnsetTlsClientCert
`func (o *LoggingTlsCommon) UnsetTlsClientCert()`

UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
### GetTlsClientKey

`func (o *LoggingTlsCommon) GetTlsClientKey() string`

GetTlsClientKey returns the TlsClientKey field if non-nil, zero value otherwise.

### GetTlsClientKeyOk

`func (o *LoggingTlsCommon) GetTlsClientKeyOk() (*string, bool)`

GetTlsClientKeyOk returns a tuple with the TlsClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientKey

`func (o *LoggingTlsCommon) SetTlsClientKey(v string)`

SetTlsClientKey sets TlsClientKey field to given value.

### HasTlsClientKey

`func (o *LoggingTlsCommon) HasTlsClientKey() bool`

HasTlsClientKey returns a boolean if a field has been set.

### SetTlsClientKeyNil

`func (o *LoggingTlsCommon) SetTlsClientKeyNil(b bool)`

 SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil

### UnsetTlsClientKey
`func (o *LoggingTlsCommon) UnsetTlsClientKey()`

UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
### GetTlsHostname

`func (o *LoggingTlsCommon) GetTlsHostname() string`

GetTlsHostname returns the TlsHostname field if non-nil, zero value otherwise.

### GetTlsHostnameOk

`func (o *LoggingTlsCommon) GetTlsHostnameOk() (*string, bool)`

GetTlsHostnameOk returns a tuple with the TlsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsHostname

`func (o *LoggingTlsCommon) SetTlsHostname(v string)`

SetTlsHostname sets TlsHostname field to given value.

### HasTlsHostname

`func (o *LoggingTlsCommon) HasTlsHostname() bool`

HasTlsHostname returns a boolean if a field has been set.

### SetTlsHostnameNil

`func (o *LoggingTlsCommon) SetTlsHostnameNil(b bool)`

 SetTlsHostnameNil sets the value for TlsHostname to be an explicit nil

### UnsetTlsHostname
`func (o *LoggingTlsCommon) UnsetTlsHostname()`

UnsetTlsHostname ensures that no value is present for TlsHostname, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


