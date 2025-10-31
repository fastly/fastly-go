# PoolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsCaCert** | Pointer to **NullableString** | A secure certificate to authenticate a server with. Must be in PEM format. | [optional] [default to "null"]
**TlsClientCert** | Pointer to **NullableString** | The client certificate used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TlsClientKey** | Pointer to **NullableString** | The client private key used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TlsCertHostname** | Pointer to **NullableString** | The hostname used to verify a server&#39;s certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN). | [optional] [default to "null"]
**UseTls** | Pointer to **string** | Whether to use TLS. | [optional] [default to "0"]
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceId** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Name for the Pool. | [optional] 
**Shield** | Pointer to **NullableString** | Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding. | [optional] [default to "null"]
**RequestCondition** | Pointer to **NullableString** | Condition which, if met, will select this configuration during a request. Optional. | [optional] 
**TlsCiphers** | Pointer to **NullableString** | List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional. | [optional] 
**TlsSniHostname** | Pointer to **NullableString** | SNI hostname. Optional. | [optional] 
**MinTlsVersion** | Pointer to **NullableInt32** | Minimum allowed TLS version on connections to this server. Optional. | [optional] 
**MaxTlsVersion** | Pointer to **NullableInt32** | Maximum allowed TLS version on connections to this server. Optional. | [optional] 
**Healthcheck** | Pointer to **NullableString** | Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools. | [optional] 
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Type** | Pointer to **string** | What type of load balance group to use. | [optional] 
**OverrideHost** | Pointer to **NullableString** | The hostname to [override the Host header](https://www.fastly.com/documentation/guides/full-site-delivery/domains-and-origins/specifying-an-override-host/). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting. | [optional] [default to "null"]
**BetweenBytesTimeout** | Pointer to **string** | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, for Delivery services, the response received so far will be considered complete and the fetch will end. For Compute services, timeout expiration is treated as a failure of the backend connection, and an error is generated. May be set at runtime using `bereq.between_bytes_timeout`. | [optional] 
**ConnectTimeout** | Pointer to **string** | How long to wait for a timeout in milliseconds. | [optional] 
**FirstByteTimeout** | Pointer to **string** | How long to wait for the first byte in milliseconds. | [optional] 
**MaxConnDefault** | Pointer to **string** | Maximum number of connections. | [optional] [default to "200"]
**TlsCheckCert** | Pointer to **NullableString** | Be strict on checking TLS certs. | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Quorum** | Pointer to **string** | Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up. | [optional] [default to "75"]

## Methods

### NewPoolResponse

`func NewPoolResponse() *PoolResponse`

NewPoolResponse instantiates a new PoolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolResponseWithDefaults

`func NewPoolResponseWithDefaults() *PoolResponse`

NewPoolResponseWithDefaults instantiates a new PoolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsCaCert

`func (o *PoolResponse) GetTlsCaCert() string`

GetTlsCaCert returns the TlsCaCert field if non-nil, zero value otherwise.

### GetTlsCaCertOk

`func (o *PoolResponse) GetTlsCaCertOk() (*string, bool)`

GetTlsCaCertOk returns a tuple with the TlsCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCaCert

`func (o *PoolResponse) SetTlsCaCert(v string)`

SetTlsCaCert sets TlsCaCert field to given value.

### HasTlsCaCert

`func (o *PoolResponse) HasTlsCaCert() bool`

HasTlsCaCert returns a boolean if a field has been set.

### SetTlsCaCertNil

`func (o *PoolResponse) SetTlsCaCertNil(b bool)`

 SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil

### UnsetTlsCaCert
`func (o *PoolResponse) UnsetTlsCaCert()`

UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
### GetTlsClientCert

`func (o *PoolResponse) GetTlsClientCert() string`

GetTlsClientCert returns the TlsClientCert field if non-nil, zero value otherwise.

### GetTlsClientCertOk

`func (o *PoolResponse) GetTlsClientCertOk() (*string, bool)`

GetTlsClientCertOk returns a tuple with the TlsClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCert

`func (o *PoolResponse) SetTlsClientCert(v string)`

SetTlsClientCert sets TlsClientCert field to given value.

### HasTlsClientCert

`func (o *PoolResponse) HasTlsClientCert() bool`

HasTlsClientCert returns a boolean if a field has been set.

### SetTlsClientCertNil

`func (o *PoolResponse) SetTlsClientCertNil(b bool)`

 SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil

### UnsetTlsClientCert
`func (o *PoolResponse) UnsetTlsClientCert()`

UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
### GetTlsClientKey

`func (o *PoolResponse) GetTlsClientKey() string`

GetTlsClientKey returns the TlsClientKey field if non-nil, zero value otherwise.

### GetTlsClientKeyOk

`func (o *PoolResponse) GetTlsClientKeyOk() (*string, bool)`

GetTlsClientKeyOk returns a tuple with the TlsClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientKey

`func (o *PoolResponse) SetTlsClientKey(v string)`

SetTlsClientKey sets TlsClientKey field to given value.

### HasTlsClientKey

`func (o *PoolResponse) HasTlsClientKey() bool`

HasTlsClientKey returns a boolean if a field has been set.

### SetTlsClientKeyNil

`func (o *PoolResponse) SetTlsClientKeyNil(b bool)`

 SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil

### UnsetTlsClientKey
`func (o *PoolResponse) UnsetTlsClientKey()`

UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
### GetTlsCertHostname

`func (o *PoolResponse) GetTlsCertHostname() string`

GetTlsCertHostname returns the TlsCertHostname field if non-nil, zero value otherwise.

### GetTlsCertHostnameOk

`func (o *PoolResponse) GetTlsCertHostnameOk() (*string, bool)`

GetTlsCertHostnameOk returns a tuple with the TlsCertHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertHostname

`func (o *PoolResponse) SetTlsCertHostname(v string)`

SetTlsCertHostname sets TlsCertHostname field to given value.

### HasTlsCertHostname

`func (o *PoolResponse) HasTlsCertHostname() bool`

HasTlsCertHostname returns a boolean if a field has been set.

### SetTlsCertHostnameNil

`func (o *PoolResponse) SetTlsCertHostnameNil(b bool)`

 SetTlsCertHostnameNil sets the value for TlsCertHostname to be an explicit nil

### UnsetTlsCertHostname
`func (o *PoolResponse) UnsetTlsCertHostname()`

UnsetTlsCertHostname ensures that no value is present for TlsCertHostname, not even an explicit nil
### GetUseTls

`func (o *PoolResponse) GetUseTls() string`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *PoolResponse) GetUseTlsOk() (*string, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *PoolResponse) SetUseTls(v string)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *PoolResponse) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PoolResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PoolResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PoolResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PoolResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PoolResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PoolResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *PoolResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PoolResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PoolResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PoolResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *PoolResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *PoolResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *PoolResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PoolResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PoolResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PoolResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *PoolResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *PoolResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceId

`func (o *PoolResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PoolResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PoolResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PoolResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetVersion

`func (o *PoolResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PoolResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PoolResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PoolResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *PoolResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShield

`func (o *PoolResponse) GetShield() string`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *PoolResponse) GetShieldOk() (*string, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *PoolResponse) SetShield(v string)`

SetShield sets Shield field to given value.

### HasShield

`func (o *PoolResponse) HasShield() bool`

HasShield returns a boolean if a field has been set.

### SetShieldNil

`func (o *PoolResponse) SetShieldNil(b bool)`

 SetShieldNil sets the value for Shield to be an explicit nil

### UnsetShield
`func (o *PoolResponse) UnsetShield()`

UnsetShield ensures that no value is present for Shield, not even an explicit nil
### GetRequestCondition

`func (o *PoolResponse) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *PoolResponse) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *PoolResponse) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *PoolResponse) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *PoolResponse) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *PoolResponse) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetTlsCiphers

`func (o *PoolResponse) GetTlsCiphers() string`

GetTlsCiphers returns the TlsCiphers field if non-nil, zero value otherwise.

### GetTlsCiphersOk

`func (o *PoolResponse) GetTlsCiphersOk() (*string, bool)`

GetTlsCiphersOk returns a tuple with the TlsCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCiphers

`func (o *PoolResponse) SetTlsCiphers(v string)`

SetTlsCiphers sets TlsCiphers field to given value.

### HasTlsCiphers

`func (o *PoolResponse) HasTlsCiphers() bool`

HasTlsCiphers returns a boolean if a field has been set.

### SetTlsCiphersNil

`func (o *PoolResponse) SetTlsCiphersNil(b bool)`

 SetTlsCiphersNil sets the value for TlsCiphers to be an explicit nil

### UnsetTlsCiphers
`func (o *PoolResponse) UnsetTlsCiphers()`

UnsetTlsCiphers ensures that no value is present for TlsCiphers, not even an explicit nil
### GetTlsSniHostname

`func (o *PoolResponse) GetTlsSniHostname() string`

GetTlsSniHostname returns the TlsSniHostname field if non-nil, zero value otherwise.

### GetTlsSniHostnameOk

`func (o *PoolResponse) GetTlsSniHostnameOk() (*string, bool)`

GetTlsSniHostnameOk returns a tuple with the TlsSniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSniHostname

`func (o *PoolResponse) SetTlsSniHostname(v string)`

SetTlsSniHostname sets TlsSniHostname field to given value.

### HasTlsSniHostname

`func (o *PoolResponse) HasTlsSniHostname() bool`

HasTlsSniHostname returns a boolean if a field has been set.

### SetTlsSniHostnameNil

`func (o *PoolResponse) SetTlsSniHostnameNil(b bool)`

 SetTlsSniHostnameNil sets the value for TlsSniHostname to be an explicit nil

### UnsetTlsSniHostname
`func (o *PoolResponse) UnsetTlsSniHostname()`

UnsetTlsSniHostname ensures that no value is present for TlsSniHostname, not even an explicit nil
### GetMinTlsVersion

`func (o *PoolResponse) GetMinTlsVersion() int32`

GetMinTlsVersion returns the MinTlsVersion field if non-nil, zero value otherwise.

### GetMinTlsVersionOk

`func (o *PoolResponse) GetMinTlsVersionOk() (*int32, bool)`

GetMinTlsVersionOk returns a tuple with the MinTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTlsVersion

`func (o *PoolResponse) SetMinTlsVersion(v int32)`

SetMinTlsVersion sets MinTlsVersion field to given value.

### HasMinTlsVersion

`func (o *PoolResponse) HasMinTlsVersion() bool`

HasMinTlsVersion returns a boolean if a field has been set.

### SetMinTlsVersionNil

`func (o *PoolResponse) SetMinTlsVersionNil(b bool)`

 SetMinTlsVersionNil sets the value for MinTlsVersion to be an explicit nil

### UnsetMinTlsVersion
`func (o *PoolResponse) UnsetMinTlsVersion()`

UnsetMinTlsVersion ensures that no value is present for MinTlsVersion, not even an explicit nil
### GetMaxTlsVersion

`func (o *PoolResponse) GetMaxTlsVersion() int32`

GetMaxTlsVersion returns the MaxTlsVersion field if non-nil, zero value otherwise.

### GetMaxTlsVersionOk

`func (o *PoolResponse) GetMaxTlsVersionOk() (*int32, bool)`

GetMaxTlsVersionOk returns a tuple with the MaxTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTlsVersion

`func (o *PoolResponse) SetMaxTlsVersion(v int32)`

SetMaxTlsVersion sets MaxTlsVersion field to given value.

### HasMaxTlsVersion

`func (o *PoolResponse) HasMaxTlsVersion() bool`

HasMaxTlsVersion returns a boolean if a field has been set.

### SetMaxTlsVersionNil

`func (o *PoolResponse) SetMaxTlsVersionNil(b bool)`

 SetMaxTlsVersionNil sets the value for MaxTlsVersion to be an explicit nil

### UnsetMaxTlsVersion
`func (o *PoolResponse) UnsetMaxTlsVersion()`

UnsetMaxTlsVersion ensures that no value is present for MaxTlsVersion, not even an explicit nil
### GetHealthcheck

`func (o *PoolResponse) GetHealthcheck() string`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *PoolResponse) GetHealthcheckOk() (*string, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *PoolResponse) SetHealthcheck(v string)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *PoolResponse) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### SetHealthcheckNil

`func (o *PoolResponse) SetHealthcheckNil(b bool)`

 SetHealthcheckNil sets the value for Healthcheck to be an explicit nil

### UnsetHealthcheck
`func (o *PoolResponse) UnsetHealthcheck()`

UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
### GetComment

`func (o *PoolResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PoolResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PoolResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PoolResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *PoolResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *PoolResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetType

`func (o *PoolResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoolResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoolResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PoolResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOverrideHost

`func (o *PoolResponse) GetOverrideHost() string`

GetOverrideHost returns the OverrideHost field if non-nil, zero value otherwise.

### GetOverrideHostOk

`func (o *PoolResponse) GetOverrideHostOk() (*string, bool)`

GetOverrideHostOk returns a tuple with the OverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideHost

`func (o *PoolResponse) SetOverrideHost(v string)`

SetOverrideHost sets OverrideHost field to given value.

### HasOverrideHost

`func (o *PoolResponse) HasOverrideHost() bool`

HasOverrideHost returns a boolean if a field has been set.

### SetOverrideHostNil

`func (o *PoolResponse) SetOverrideHostNil(b bool)`

 SetOverrideHostNil sets the value for OverrideHost to be an explicit nil

### UnsetOverrideHost
`func (o *PoolResponse) UnsetOverrideHost()`

UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
### GetBetweenBytesTimeout

`func (o *PoolResponse) GetBetweenBytesTimeout() string`

GetBetweenBytesTimeout returns the BetweenBytesTimeout field if non-nil, zero value otherwise.

### GetBetweenBytesTimeoutOk

`func (o *PoolResponse) GetBetweenBytesTimeoutOk() (*string, bool)`

GetBetweenBytesTimeoutOk returns a tuple with the BetweenBytesTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenBytesTimeout

`func (o *PoolResponse) SetBetweenBytesTimeout(v string)`

SetBetweenBytesTimeout sets BetweenBytesTimeout field to given value.

### HasBetweenBytesTimeout

`func (o *PoolResponse) HasBetweenBytesTimeout() bool`

HasBetweenBytesTimeout returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *PoolResponse) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *PoolResponse) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *PoolResponse) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *PoolResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetFirstByteTimeout

`func (o *PoolResponse) GetFirstByteTimeout() string`

GetFirstByteTimeout returns the FirstByteTimeout field if non-nil, zero value otherwise.

### GetFirstByteTimeoutOk

`func (o *PoolResponse) GetFirstByteTimeoutOk() (*string, bool)`

GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstByteTimeout

`func (o *PoolResponse) SetFirstByteTimeout(v string)`

SetFirstByteTimeout sets FirstByteTimeout field to given value.

### HasFirstByteTimeout

`func (o *PoolResponse) HasFirstByteTimeout() bool`

HasFirstByteTimeout returns a boolean if a field has been set.

### GetMaxConnDefault

`func (o *PoolResponse) GetMaxConnDefault() string`

GetMaxConnDefault returns the MaxConnDefault field if non-nil, zero value otherwise.

### GetMaxConnDefaultOk

`func (o *PoolResponse) GetMaxConnDefaultOk() (*string, bool)`

GetMaxConnDefaultOk returns a tuple with the MaxConnDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnDefault

`func (o *PoolResponse) SetMaxConnDefault(v string)`

SetMaxConnDefault sets MaxConnDefault field to given value.

### HasMaxConnDefault

`func (o *PoolResponse) HasMaxConnDefault() bool`

HasMaxConnDefault returns a boolean if a field has been set.

### GetTlsCheckCert

`func (o *PoolResponse) GetTlsCheckCert() string`

GetTlsCheckCert returns the TlsCheckCert field if non-nil, zero value otherwise.

### GetTlsCheckCertOk

`func (o *PoolResponse) GetTlsCheckCertOk() (*string, bool)`

GetTlsCheckCertOk returns a tuple with the TlsCheckCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCheckCert

`func (o *PoolResponse) SetTlsCheckCert(v string)`

SetTlsCheckCert sets TlsCheckCert field to given value.

### HasTlsCheckCert

`func (o *PoolResponse) HasTlsCheckCert() bool`

HasTlsCheckCert returns a boolean if a field has been set.

### SetTlsCheckCertNil

`func (o *PoolResponse) SetTlsCheckCertNil(b bool)`

 SetTlsCheckCertNil sets the value for TlsCheckCert to be an explicit nil

### UnsetTlsCheckCert
`func (o *PoolResponse) UnsetTlsCheckCert()`

UnsetTlsCheckCert ensures that no value is present for TlsCheckCert, not even an explicit nil
### GetId

`func (o *PoolResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PoolResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuorum

`func (o *PoolResponse) GetQuorum() string`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *PoolResponse) GetQuorumOk() (*string, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *PoolResponse) SetQuorum(v string)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *PoolResponse) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


